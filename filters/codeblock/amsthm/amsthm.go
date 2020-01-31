// Author: Luke Huang <lukehuang.ca@me.com>
// Copyright: Luke Huang <lukehuang.ca@me.com>
// License: BSD3

package amsthm

import (
	"os"
	"fmt"
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
	PdfOnly	bool   `json:pdfonly`
}

const latexTemplate = `
\begin{{"{"}}{{.Type}}{{"}"}}[{{.Item}}]
{{.Text}}
\end{{"{"}}{{.Type}}{{"}"}}
`

const htmlTemplate = `
<div class="{{.Type}}">
({{.Item}}) {{"xxx"}} {{.Text}}
</div>
`

func (theorem *Theorem) Block(target string, content string) interface{} {
	var tpl *template.Template
	var output bytes.Buffer
	var thm Theorem
	bytes := []byte(content)
	for i, ch := range bytes {
		if ch == '\r' || ch == '\n' {
			bytes[i] = ' '
		}
	}
	err := json.Unmarshal(bytes, &thm); if err != nil {
		fmt.Fprintf(os.Stderr, "unmarshal %s\n", err);
		return nil
	}

	if (thm.PdfOnly && (target == "html" || target == "html5")) {
		return pf.Para([]interface{} {pf.Str("")})
	}

	tpl, err = template.New("theorem").
		Parse(templateMap[target]); if err != nil {
		return nil
	}
	tpl.Execute(&output, thm)
	return pf.RawBlock(target, output.String())
}
