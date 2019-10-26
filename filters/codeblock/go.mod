module main

go 1.13

require (
	github.com/oltolm/go-pandocfilters v0.0.0-20170923090435-277a35aeaa53
	github.com/vupiggy/pandoc-filter/codeblock/amsthm v0.0.1
	github.com/vupiggy/pandoc-filter/codeblock/code v0.0.1
	github.com/vupiggy/pandoc-filter/codeblock/codeblock v0.0.1
	github.com/vupiggy/pandoc-filter/codeblock/figure v0.0.1
)

replace github.com/vupiggy/pandoc-filter/codeblock/codeblock v0.0.1 => ./codeblock

replace github.com/vupiggy/pandoc-filter/codeblock/figure v0.0.1 => ./figure

replace github.com/vupiggy/pandoc-filter/codeblock/amsthm v0.0.1 => ./amsthm

replace github.com/vupiggy/pandoc-filter/codeblock/code v0.0.1 => ./code
