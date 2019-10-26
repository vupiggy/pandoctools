package code

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	pf "github.com/oltolm/go-pandocfilters"
)

const (
	BLOCKOUT = iota
	BLOCKOPEN
	BLOCKCLOSE
)

type Code string

func (c *Code) Block(class string, target string, content string, keyvals []interface{}) interface{} {
	path, keyvals := pf.GetValue(keyvals, "path")
	if path == nil {
		// Embedded code block, leave it as is
		return nil
	}

	block, keyvals := pf.GetValue(keyvals, "block")
	if block == nil {
		fmt.Fprintf(os.Stderr, "which block to insert?\n")
		return nil
	}

	f, err := os.Open(path.(string))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s doesn't exist\n", path.(string))
		return nil
	}
	defer f.Close()

	var code string
	state := BLOCKOUT
	input := bufio.NewScanner(f)

ScanLoop:
	for input.Scan() {
		// TODO(luke): Deal with nested block with the same block name properly.
		//             For now letting the outer block win sounds not so bad.
		switch state {
		case BLOCKOUT:
			if strings.HasPrefix(input.Text(), "//!+ " + block.(string)) {
				// migrate to next state, ignore the opening line
				state = BLOCKOPEN
			}
		case BLOCKOPEN:
			if strings.HasPrefix(input.Text(), "//!- " + block.(string)) {
				state = BLOCKCLOSE
				break ScanLoop
			}
			code = code + input.Text() + "\n"
		}
	}
	if state != BLOCKCLOSE {
		fmt.Fprintf(os.Stderr, "block is not closed!")
		code = code + content
	}

	return pf.CodeBlock([]interface{} {
		"",
		[]string{class},
		[]interface{}{},
	}, code)
}
