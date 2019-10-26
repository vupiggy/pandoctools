// Author: Luke Huang <lukehuang.ca@me.com>
// Copyright: Luke Huang <lukehuang.ca@me.com>
// License: BSD3

package amsthm

import (
	"bytes"
	"text/template"
	pf "github.com/oltolm/go-pandocfilters"
)

var templateMap = map[string]string{
	"html":		htmlTemplate,
	"html5":	htmlTemplate,
	"latex":	latexTemplate,
}

type Theorem struct {
	Type	interface{}
	Item	interface{}
	Content	string
}

const latexTemplate = `
\begin{{"{"}}{{.Type}}{{"}"}}[{{.Item}}]
{{.Content}}
\end{{"{"}}{{.Type}}{{"}"}}
`

const htmlTemplate = `
<div class="{{.Type}}">
({{.Item}}) {{.Content}}
</div>
`

func (theorem *Theorem) Block(class string, target string, content string, keyvals []interface{}) interface{} {
	var tpl *template.Template
	var err error
	var output bytes.Buffer
	thm := Theorem{
		Content: content,
	}
	if len(keyvals) > 0 {
		thm.Type, keyvals	= pf.GetValue(keyvals, "type")
		thm.Item, keyvals	= pf.GetValue(keyvals, "item")
	}
	tpl, err = template.New("theorem").
		Funcs(template.FuncMap{
		"stringify": func (x interface{}) string { return x.(string) }}).
			Parse(templateMap[target]); if err != nil {
			return ""
		}
	tpl.Execute(&output, thm)
	return pf.RawBlock(target, output.String())
}
