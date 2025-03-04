package httputil

import (
	"fmt"
	"net/http"

	"github.com/guibressan/tmt/internal/schema"
)

func SendRedirect(w http.ResponseWriter, url string) error {
	b := `
<head>
	<meta http-equiv="Refresh" content="0; URL=%s"/>
</head>
	`
	_, err := fmt.Fprintf(w, b, url)
	return err
}

func ParseSchema(r *http.Request, target interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	return schema.NewDecoder().Decode(target, r.Form)
}
