package frontend

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"strings"
	"time"

	"github.com/guibressan/tmt/internal/bstd/stackerr"
	"github.com/guibressan/tmt/internal/tmt"
)

//go:embed static
var static embed.FS

func RenderRoot(w io.Writer, title string, data tmt.ProcessedData) error {
	historyBase := `
	<h3>%d</h3> 
	<ul>
		%s
	</ul>
	`
	historyRowBase := `
	<li><a href="%s">%s</a>: %s</li>
	`
	historyYears := &bytes.Buffer{}
	historyItems := &bytes.Buffer{}
	for _, v := range data.Year {
		for _, w := range v.Migrations {
			fmt.Fprintf(historyItems, historyRowBase, w.Uri, w.Company, w.Description)
		}
		fmt.Fprintf(historyYears, historyBase, v.Year, historyItems.String())
		historyItems.Reset()
	}
	top10FromBase := `
	<ul>
		%s
	</ul>
	`
	top10FromRowBase := `
	<li>%s: %d</li>
	`
	top10From := &bytes.Buffer{}
	top10FromItems := historyItems
	fromLen := len(data.From)
	for i := 0; i < 10 && i < fromLen; i++ {
		f := data.From[i]
		fmt.Fprintf(top10FromItems, top10FromRowBase, f.Name, f.Count)
	}
	fmt.Fprintf(top10From, top10FromBase, top10FromItems.String())
	top10FromItems.Reset()
	top10ToBase := top10FromBase
	top10ToRowBase := top10FromRowBase
	top10To := &bytes.Buffer{}
	top10ToItems := top10FromItems
	toLen := len(data.To)
	for i := 0; i < 10 && i < toLen; i++ {
		f := data.To[i]
		fmt.Fprintf(top10ToItems, top10ToRowBase, f.Name, f.Count)
	}
	fmt.Fprintf(top10To, top10ToBase, top10ToItems.String())
	top10ToItems.Reset()
	top10MigrationBase := top10ToBase
	top10MigrationRowBase := top10ToRowBase
	top10Migration := &bytes.Buffer{}
	top10MigrationItems := top10ToItems
	migrationLen := len(data.FromTo)
	for i := 0; i < 10 && i < migrationLen; i++ {
		f := data.FromTo[i]
		fmt.Fprintf(
			top10MigrationItems,
			top10MigrationRowBase,
			strings.Replace(f.Name, " ", " to ", 1),
			f.Count,
		)
	}
	fmt.Fprintf(top10Migration, top10MigrationBase, top10MigrationItems.String())
	top10MigrationItems.Reset()
	return renderBase(
		w,
		"",
		string(MustRead("root.html")),
		fmt.Sprintf(`{{define "history"}}%s{{end}}`, historyYears.String()),
		fmt.Sprintf(`{{define "from"}}%s{{end}}`, top10From.String()),
		fmt.Sprintf(`{{define "to"}}%s{{end}}`, top10To.String()),
		fmt.Sprintf(`{{define "migration"}}%s{{end}}`, top10Migration.String()),
	)
}

func RenderSubmitMessage(w io.Writer, title, captchaId string) error {
	return renderBase(
		w,
		"",
		string(MustRead("root.html")),
		string(MustRead("submit.html")),
		fmt.Sprintf(`{{define "captchaid"}}%s{{end}}`, captchaId),
	)
}

func RenderCredits(w io.Writer, title string) error {
	return renderBase(
		w,
		"",
		string(MustRead("root.html")),
		string(MustRead("credits.html")),
	)
}

func MustRead(fname string) []byte {
	f, err := static.ReadFile("static/" + fname)
	if err != nil {
		panic(stackerr.Wrap(err))
	}
	return f
}

func renderBase(
	w io.Writer, title string, templates ...string,
) error {
	tma := func(tmpl string) {
		templates = append(templates, tmpl)
	}
	t := template.New("")
	t, err := t.Parse(string(MustRead("base.html")))
	if err != nil {
		return stackerr.Wrap(err)
	}
	tma(string(MustRead("header.html")))
	tma(string(MustRead("footer.html")))
	if title != "" {
		tma(fmt.Sprintf(`{{define "title"}}%s{{end}}`, title))
	}
	tma(fmt.Sprintf(`{{define "year"}}%d{{end}}`, time.Now().Year()))
	for i, v := range templates {
		swap := t.New("swap")
		_, err = swap.Parse(v)
		if err != nil {
			return fmt.Errorf("error parsing template %d; %w", i, err)
		}
	}
	return t.Execute(w, nil)
}
