package code

import (
	"fmt"
	"os"
	"encoding/json"
	"bufio"
	"strings"
	pf "github.com/oltolm/go-pandocfilters"
)

const (
	BLOCKOUT = iota
	BLOCKOPEN
	BLOCKCLOSE
)

type Code struct {
	Path     string `json:path`
	Lang     string `json:lang`
	Segment  string `json:segment`
}

func (c *Code) Block(class string, target string, content string) interface{} {
	var code Code
	err := json.Unmarshal([]byte(content), &code); if err != nil {
		return nil
	}

	f, err := os.Open(code.Path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s doesn't exist\n", code.Path)
		return nil
	}
	defer f.Close()

	var codeStr string
	state := BLOCKOUT
	input := bufio.NewScanner(f)

ScanLoop:
	for input.Scan() {
		// TODO(luke): Deal with nested block with the same block name properly.
		//             For now letting the outer block win sounds not so bad.
		switch state {
		case BLOCKOUT:
			if strings.HasPrefix(input.Text(), "//!+ " + code.Segment) {
				// migrate to next state, ignore the opening line
				state = BLOCKOPEN
			}
		case BLOCKOPEN:
			if strings.HasPrefix(input.Text(), "//!- " + code.Segment) {
				state = BLOCKCLOSE
				break ScanLoop
			}
			codeStr = codeStr + input.Text() + "\n"
		}
	}
	if state != BLOCKCLOSE {
		fmt.Fprintf(os.Stderr, "block is not closed!")
		codeStr = codeStr + content
	}

	return pf.CodeBlock([]interface{} {
		"",
		[]string{code.Lang},
		[]interface{}{},
	}, codeStr)
}
