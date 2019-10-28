// Author: Luke Huang <lukehuang.ca@me.com>
// Copyright: Luke Huang <lukehuang.ca@me.com>
// License: BSD3

package amsthm

import (
	"bytes"
	"encoding/json"
	"text/template"
	pf "github.com/oltolm/go-pandocfilters"
)

var templateMap = map[string]string{
	"html":		htmlTemplate,
	"html5":	htmlTemplate,
	"latex":	latexTemplate,
}

type Theorem struct {
	Type	string `json:type`
	Item	string `json:item`
	Text	string `json:text`
}

const latexTemplate = `
\begin{{"{"}}{{.Type}}{{"}"}}[{{.Item}}]
{{.Text}}
\end{{"{"}}{{.Type}}{{"}"}}
`

const htmlTemplate = `
<div class="{{.Type}}">
({{.Item}}) {{.Text}}
</div>
`

func (theorem *Theorem) Block(target string, content string) interface{} {
	var tpl *template.Template
	var output bytes.Buffer
	var thm Theorem
	err := json.Unmarshal([]byte(content), &thm); if err != nil {
		return nil
	}
	tpl, err = template.New("theorem").
			Parse(templateMap[target]); if err != nil {
			return nil
		}
	tpl.Execute(&output, thm)
	return pf.RawBlock(target, output.String())
}
