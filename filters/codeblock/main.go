// Author: Luke Huang <lukehuang.ca@me.com>
// Copyright: Luke Huang <lukehuang.ca@me.com>
// License: BSD3

package main

import (
	_ "os"
	_ "fmt"
	"github.com/vupiggy/pandoc-filter/codeblock/codeblock"
	"github.com/vupiggy/pandoc-filter/codeblock/figure"
	"github.com/vupiggy/pandoc-filter/codeblock/amsthm"
	"github.com/vupiggy/pandoc-filter/codeblock/code"
	pf "github.com/oltolm/go-pandocfilters"
)

func Insert(cb codeblock.CodeBlock, target string, content string) interface{} {
	if cb != nil {
		return cb.Block(target, content)
	}
	return nil
}

var fig figure.Figure
var thm amsthm.Theorem
var cod code.Code

var cbMap = map[string]codeblock.CodeBlock {
	"figure"  : &fig,
	"theorem" : &thm,
	"snippet" : &cod,
}

func processCB(key string, value interface{}, target string, meta interface{}) interface{} {
	var class string
	var content string

	if key == "CodeBlock" {
		cb		:= value.([]interface{})
		attrs	:= cb[0].([]interface{})

		if len(attrs[1].([]interface{})) > 0 {
			class	 = attrs[1].([]interface{})[0].(string)
			content  = cb[1].(string)
		}
	}
	if len(class) > 0 {
		t := cbMap[class]
		if t != nil {
			return Insert(t, target, content)
		}
	}
	return nil
}

func main() {
	pf.ToJSONFilter(processCB)
}
