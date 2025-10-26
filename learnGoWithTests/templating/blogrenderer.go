package blogrenderer

import (
	"html/template"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

const (
	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
)

func Render(w io.Writer, p Post) error {
	templ, err := template.New("blog").Parse(postTemplate)
	if err != nil {
		return err
	}
	if err := templ.Execute(w, p); err != nil {
		return err
	}
	return nil

	/*
		_, err := fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>", p.Title, p.Description)
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(w, "Tags: <ul>")
		if err != nil {
			return err
		}
		for _, tag := range p.Tags {
			_, err = fmt.Fprintf(w, "<li>%s</li>", tag)
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprint(w, "</ul>")
		if err != nil {
			return err
		}
		return nil

	*/
}
