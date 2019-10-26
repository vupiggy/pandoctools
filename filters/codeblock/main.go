// Author: Luke Huang <lukehuang.ca@me.com>
// Copyright: Luke Huang <lukehuang.ca@me.com>
// License: BSD3

package main

import (
	"github.com/vupiggy/pandoc-filter/codeblock/codeblock"
	"github.com/vupiggy/pandoc-filter/codeblock/figure"
	"github.com/vupiggy/pandoc-filter/codeblock/amsthm"
	"github.com/vupiggy/pandoc-filter/codeblock/code"
	pf "github.com/oltolm/go-pandocfilters"
)

func Insert(cb codeblock.CodeBlock,
	class string, target string, content string, keyvals []interface{}) interface{} {
	if cb != nil {
		return cb.Block(class, target, content, keyvals)
	}
	return nil
}

var fig figure.Figure
var thm amsthm.Theorem
var cod code.Code

var cbMap = map[string]codeblock.CodeBlock {
	"figure":  &fig,
	"theorem": &thm,
	"Go":      &cod,
	"C":       &cod,
	"C++":     &cod,
}

func processCB(key string, value interface{}, target string, meta interface{}) interface{} {
	if key == "CodeBlock" {
		cb		:= value.([]interface{})
		attrs	:= cb[0].([]interface{})
		classes	:= attrs[1].([]interface{})
		keyvals	:= attrs[2].([]interface{})
		content := cb[1].(string)

		if len(classes) > 0 {
			t   := cbMap[classes[0].(string)]
			if t != nil {
				return Insert(cbMap[classes[0].(string)], classes[0].(string), target, content, keyvals)
			}
		}
	}
	return nil
}

func main() {
	pf.ToJSONFilter(processCB)
}
