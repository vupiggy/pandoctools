package codeblock

type CodeBlock interface {
	Block(class string, target string, content string, keyval []interface{}) interface{}
}
