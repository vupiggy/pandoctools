# Tools and templates for Pandoc

This repository provides a set of tools and templates for documention with
[Pandoc](https://pandoc.org "Pandoc's homepage").
It includes Pandoc filters written in [Golang](https://golang.org "Golang's homepage"),
templates and style sheets for LaTeX, HTML output,
and a [GNU Make](https://www.gnu.org/software/make/ "GNU Make") framework for building.

## Filters

Thanks to [Pandoc filters for Go](https://github.com/oltolm/go-pandocfilters "Pandoc filters for Go")
on which all my filters are written.

The filters have to be built into binary first.
Go to `filters/codeblock` directory, 
then run:
```bash
go build -o codeblock-filter main.go
```

Then set the filter's path properly in the Makefile,
it might look like this:
```Makefile
DOCTOOLS	=	$(HOME)/Projects/pandoctools
PANDOCFILTERS	=	$(HOME)/Projects/pandoctools/filters/codeblock/codeblock-filter
# ...
include $(DOCTOOLS)/make/Makefile.in
```

Markdown files in `examples` directory shows how to insert code snippet, figures, and `amsthm` style theorems in the document.

## Make

`make/Makefile.in` defines rules to build a Pandoc file into PDF or HTMl.
It also defines how to generate figures from TikZ into PDF or SVG,
that can be inserted into PDF or HTML repectively.

## Styles

## Fonts

The users will have to put their own fonts in `fonts` directory since
almost all fonts are licensed so that can not be provided in an open sourced project.
The users might need to modify `styles/md2pdf_template.tex` to use their favorite fonts for PDF output.
See [fontspec](https://ctan.org/pkg/fontspec?lang=en) for more details.
It's unlikely the users need to do anything with HTML output.
