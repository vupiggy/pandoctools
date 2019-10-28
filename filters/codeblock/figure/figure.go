package figure

import (
	"bytes"
	"text/template"
	"encoding/json"
	pf "github.com/oltolm/go-pandocfilters"
)

var suffixMap = map[string]string{
	"html":  "svg",
	"html5": "svg",
	"latex": "pdf",
}

var templateMap = map[string]string{
	"html":		htmlTemplate,
	"html5":	htmlTemplate,
	"latex":	latexTemplate,
}

//!+ figure
type Figure struct {
	Target	string
	Path	string		`json:path`
	Caption	string		`json:caption`
	Label	string		`json:label`
	Options	string		`json:options`
	Place	string		`json:place`
	Suffix  map[string]string
}

const latexTemplate = `
\begin{figure}[{{.Place}}]
  \centering
  \includegraphics[{{.Options}}]%
      {{"{"}}{{.Path}}.{{index .Suffix .Target}}{{"}"}}
  \caption{{"{"}}{{.Caption}}{{"}"}}
  \label{{"{fig:"}}{{.Label}}{{"}"}}
\end{figure}
`
//!- figure

const htmlTemplate = `
<figure class="fullwidth">
    <img src="{{.Path}}.{{index .Suffix .Target}}" alt="" />
    <figcaption>
        {{.Caption}}
    </figcaption>
</figure>
`

func (figure *Figure) Block(target string, content string) interface{} {
	var tpl *template.Template
	var err error
	var output bytes.Buffer
	var fig Figure
	err = json.Unmarshal([]byte(content), &fig); if err != nil {
		return nil
	}
	fig.Target = target;
	fig.Suffix = suffixMap;
	tpl, err = template.New("figure").
		Parse(templateMap[target]); if err != nil {
		return nil
	}
	tpl.Execute(&output, fig)
	return pf.RawBlock(target, output.String())
}
