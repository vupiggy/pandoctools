package codeblock

type CodeBlock interface {
	Block(class string, target string, content string) interface{}
}
