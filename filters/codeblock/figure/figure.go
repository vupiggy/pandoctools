package figure

import (
	"bytes"
	"text/template"
	pf "github.com/oltolm/go-pandocfilters"
)

type Figure struct {
	Target	string
	Path	string
	Caption	interface{}
	Label	interface{}				// LaTeX only
	Options	interface{}				// LaTeX only
	Place	interface{}				// LaTeX only
	Suffix  map[string]string		// not externally used
}

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

const latexTemplate = `
\begin{figure}[{{.Place}}]
  \centering
  \includegraphics[{{.Options | stringify}}]%
      {{"{"}}{{.Path}}.{{index .Suffix .Target}}{{"}"}}
  \caption{{"{"}}{{.Caption | stringify}}{{"}"}}
  \label{{"{fig:"}}{{.Label | stringify}}{{"}"}}
\end{figure}
`
//!- figure

const htmlTemplate = `
<figure class="fullwidth">
    <img src="{{.Path}}.{{index .Suffix .Target}}" alt="" />
    <figcaption>
        {{.Caption | stringify}}
    </figcaption>
</figure>
`

func (figure *Figure) Block(class string, target string, content string, keyvals []interface{}) interface{} {
	var tpl *template.Template
	var err error
	var output bytes.Buffer
	fig := Figure{
		Path:	content,
		Target:	target,
		Suffix: suffixMap,
	}
	if len(keyvals) > 0 {
		fig.Caption, keyvals	= pf.GetValue(keyvals, "caption")
		fig.Label,   keyvals	= pf.GetValue(keyvals, "label")
		fig.Options, keyvals	= pf.GetValue(keyvals, "options", "width=\\textwidth")
		fig.Place,   keyvals	= pf.GetValue(keyvals, "place", "ht")
	}
	tpl, err = template.New("figure").
		Funcs(template.FuncMap{
		"stringify": func (x interface{}) string { return x.(string) }}).
			Parse(templateMap[target]); if err != nil {
			return ""
		}
	tpl.Execute(&output, fig)
	return pf.RawBlock(target, output.String())
}
