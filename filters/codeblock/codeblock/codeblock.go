package codeblock

type CodeBlock interface {
	Block(target string, content string) interface{}
}
