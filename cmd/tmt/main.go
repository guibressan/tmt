package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/guibressan/tmt/internal/bstd/stackerr"
	"github.com/guibressan/tmt/internal/bstd/util"
	"github.com/guibressan/tmt/internal/captcha"
	"github.com/guibressan/tmt/internal/frontend"
	"github.com/guibressan/tmt/internal/godotenv"
	"github.com/guibressan/tmt/internal/httputil"
	"github.com/guibressan/tmt/internal/tmt"
)

const (
	minLogInterval     = time.Second
	minCaptchaInterval = time.Second * 10
)

var hf http.HandlerFunc = handler
var (
	lastLogAt time.Time
	reqLogMu  sync.Mutex
	reqCount  atomic.Int64

	root     string
	rootOnce sync.Once

	messageFile *os.File

	lastCaptchaSt time.Time
	lastCaptchaMu sync.Mutex
)

func main() {
	godotenv.Load()
	msgFile, err := os.OpenFile(
		util.MustEnv("MESSAGE_FILE"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o660,
	)
	messageFile = msgFile
	if err != nil {
		panic(stackerr.Wrap(err))
	}
	cstore := captcha.NewMemoryStore(10, time.Minute*10)
	captcha.SetCustomStore(cstore)
	addr := util.EnvOrDefault("LISTEN_ADDR", "127.0.0.1:8080")
	log.Println("to listen on", addr)
	http.ListenAndServe(addr, hf)
}

func handler(w http.ResponseWriter, r *http.Request) {
	maybeLog(r)
	switch r.URL.Path {
	case "/ping":
		io.WriteString(w, "pong")
	case "/public/style.css":
		if !assertMethod("GET", w, r) {
			return
		}
		w.Header().Set("Cache-Control", "max-age=600")
		w.Header().Set("Content-Type", "text/css")
		b := frontend.MustRead("style.css")
		buf := bytes.NewBuffer(b)
		buf.WriteTo(w)
	case "/submit":
		if !assertMethod("GET", w, r) {
			return
		}
		if !lastCaptchaMu.TryLock() {
			w.WriteHeader(429)
			io.WriteString(w, "TOO MANY REQUESTS")
			return
		}
		defer lastCaptchaMu.Unlock()
		now := time.Now()
		if now.Sub(lastCaptchaSt) < minCaptchaInterval {
			w.WriteHeader(429)
			io.WriteString(w, "TOO MANY REQUESTS")
			return
		}
		lastCaptchaSt = now
		captchaId := captcha.New()
		err := frontend.RenderSubmitMessage(w, "Submit Message", captchaId)
		if err != nil {
			log.Println(stackerr.Wrap(err))
		}
	case "/captcha":
		if !assertMethod("GET", w, r) {
			return
		}
		var p struct {
			CaptchaId string `schema:"i"`
		}
		err := httputil.ParseSchema(r, &p)
		if err != nil {
			log.Println(stackerr.Wrap(err))
			return
		}
		err = captcha.WriteImage(
			w, p.CaptchaId, captcha.StdWidth, captcha.StdHeight,
		)
	case "/sendmsg":
		if !assertMethod("POST", w, r) {
			return
		}
		var p struct {
			Name         string `schema:"n"`
			PersonalPage string `schema:"p"`
			MessageData  string `schema:"m"`
			CaptchaId    string `schema:"i"`
			CaptchaVal   string `schema:"v"`
		}
		err := httputil.ParseSchema(r, &p)
		if err != nil {
			log.Println(stackerr.Wrap(err))
			return
		}
		if p.CaptchaId == "" {
			w.WriteHeader(400)
			io.WriteString(w, "missing captcha id")
			return
		}
		if p.CaptchaVal == "" {
			w.WriteHeader(400)
			io.WriteString(w, "missing captcha value")
			return
		}
		if !captcha.VerifyString(p.CaptchaId, p.CaptchaVal) {
			w.WriteHeader(400)
			io.WriteString(w, "bad captcha")
			return
		}
		p.Name = sanitizeStr(p.Name)
		p.PersonalPage = sanitizeStr(p.PersonalPage)
		p.MessageData = sanitizeStr(p.MessageData)
		if p.Name == "" || p.MessageData == "" {
			w.WriteHeader(400)
			io.WriteString(w, "name and message are required fields")
			return
		}
		if len(p.Name) > 255 {
			w.WriteHeader(400)
			io.WriteString(w, "name exceeded max len")
			return
		}
		if len(p.PersonalPage) > 255 {
			w.WriteHeader(400)
			io.WriteString(w, "personal page exceeded max len")
			return
		}
		if len(p.MessageData) > 1024 {
			w.WriteHeader(400)
			io.WriteString(w, "message data exceeded max len")
			return
		}
		_, err = messageFile.WriteString(
			fmt.Sprintf("%s,%s,%s\n", p.Name, p.PersonalPage, p.MessageData),
		)
		if err != nil {
			w.WriteHeader(500)
			io.WriteString(w, "INTERNAL SERVER ERROR")
			return
		}
		w.WriteHeader(201)
		io.WriteString(w, "Message saved. Thank you!")
	case "/credits":
		if !assertMethod("GET", w, r) {
			return
		}
		err := frontend.RenderCredits(w, "Credits")
		if err != nil {
			log.Println(stackerr.Wrap(err))
		}
	case "/":
		if !assertMethod("GET", w, r) {
			return
		}
		rootOnce.Do(func() {
			buf := &bytes.Buffer{}
			err := frontend.RenderRoot(buf, "", tmt.Process())
			if err != nil {
				w.WriteHeader(500)
				io.WriteString(w, "INTERNAL SERVER ERROR")
				panic(stackerr.Wrap(err))
			}
			root = buf.String()
		})
		io.WriteString(w, root)
	default:
		w.WriteHeader(404)
		io.WriteString(w, "NOT FOUND")
	}
}

func maybeLog(r *http.Request) {
	defer reqCount.Add(1)
	ok := reqLogMu.TryLock()
	if !ok {
		return
	}
	defer reqLogMu.Unlock()
	now := time.Now()
	if now.Sub(lastLogAt) < minLogInterval {
		return
	}
	lastLogAt = now
	log.Printf("request %d on %s", reqCount.Load(), r.URL.Path)
}

func assertMethod(
	expected string, w http.ResponseWriter, request *http.Request,
) bool {
	if expected == request.Method {
		return true
	}
	w.WriteHeader(405)
	io.WriteString(w, "METHOD NOT ALLOWED")
	return false
}

func sanitizeStr(in string) string {
	r := make([]byte, 0, len(in))
	for _, v := range in {
		if v >= 'a' && v <= 'z' ||
			v >= 'A' && v <= 'Z' ||
			v >= '0' && v <= '9' ||
			v == ' ' ||
			v == '-' ||
			v == '_' ||
			v == ',' ||
			v == '/' ||
			v == ':' ||
			v == '.' {
			r = append(r, byte(v))
			continue
		}
		r = append(r, '?')
	}
	return string(r)
}
