package figure

import (
	"os"
	"fmt"
	"bytes"
	"text/template"
	"encoding/json"
	pf "github.com/oltolm/go-pandocfilters"
)

var suffixMap = map[string]string{
	"html":  ".svg",
	"html5": ".svg",
	"latex": ".pdf",
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
  \includegraphics[%
      {{.Options}}]%
      {{"{"}}{{.Path}}{{index .Suffix .Target}}{{"}"}}
  \caption{{"{"}}{{.Caption}}{{"}"}}
  \label{{"{fig:"}}{{.Label}}{{"}"}}
\end{figure}
`
//!- figure

func (figure *Figure) Block(format string, content string) interface{} {
	var fig Figure
	err := json.Unmarshal([]byte(content), &fig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "figure issue\n");
		return nil
	}

	// Pandoc image block automatically set width and height,
	// don't know how to suppress yet. Use rawblock instead.
	if format == "latex" {
		var tpl *template.Template
		// var err error
		var output bytes.Buffer

		fig.Target = format;
		fig.Suffix = suffixMap;

		tpl, err = template.New("figure").
			Parse(latexTemplate); if err != nil {
			return nil
		}
		tpl.Execute(&output, fig)
		return pf.RawBlock(format, output.String())
	}

	options := [][]string{}
	return pf.Para([]interface{} {
		pf.Image(
			[]interface{} {"fig:" + fig.Label, []interface{}{}, options},
			[]interface{} {map[string]string {"c":fig.Caption, "t":"Str"}},
			[]string {fig.Path + suffixMap[format], "fig:"})})
}
