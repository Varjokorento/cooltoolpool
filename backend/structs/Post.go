package structs

import (
	"html/template"
)

type Post struct {
	Title   string
	Content template.HTML
}
