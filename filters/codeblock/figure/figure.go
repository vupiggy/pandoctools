package figure

import (
	"fmt"
	"os"
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
	Options	map[string]string	`json:options`
	Place	string		`json:place`
	Suffix  map[string]string
}

const latexTemplate = `
\begin{figure}[{{.Place}}]
  \centering
  \includegraphics[%
      {{.Options | stringify}}]%
      {{"{"}}{{.Path}}{{index .Suffix .Target}}{{"}"}}
  \caption{{"{"}}{{.Caption}}{{"}"}}
  \label{{"{fig:"}}{{.Label}}{{"}"}}
\end{figure}
`
//!- figure

func stringify(options map[string]string) string {
	var res string
	i := 0
	for key, value := range options {
		if i > 0 {
			res += ","
		}
		res += key + "=" + value
		i++
	}
	fmt.Fprintf(os.Stderr, "<%s>\n", res)
	return res
}

func (figure *Figure) Block(format string, content string) interface{} {
	var fig Figure
	err := json.Unmarshal([]byte(content), &fig)
	if err != nil {
		return nil
	}
	options := [][]string{}

	// Pandoc image block automatically set width and height,
	// don't know how to suppress yet. Use rawblock instead.
	if format == "latex" {
		var tpl *template.Template
		// var err error
		var output bytes.Buffer

		fig.Target = format;
		fig.Suffix = suffixMap;
		tpl, err = template.New("figure").
			Funcs(template.FuncMap{"stringify":stringify}).
			Parse(latexTemplate); if err != nil {
			return nil
		}
		tpl.Execute(&output, fig)
		return pf.RawBlock(format, output.String())
	}

	/*
	for key, value := range fig.Options {
		options = append(options, []string{string(key), string(value)})
	}
	*/

	return pf.Para([]interface{} {
		pf.Image(
			[]interface{} {"fig:" + fig.Label, []interface{}{}, options},
			[]interface{} {map[string]string {"c":fig.Caption, "t":"Str"}},
			[]string {fig.Path + suffixMap[format], "fig:"})})
}
