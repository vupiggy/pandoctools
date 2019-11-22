// Author: Luke Huang <lukehuang.ca@me.com>
// Copyright: Luke Huang <lukehuang.ca@me.com>
// License: BSD3

package main

import (
	"fmt"
	"os"
	"strconv"
	"image"
	"image/png"
	pf "github.com/oltolm/go-pandocfilters"
)

// TODO(luke): Only png supported now, other formats later
func imageDimensions(path string) (width int64, height int64, err error) {
	file, err := os.Open(path)
	defer file.Close()

	if err == nil {
		var img image.Config
		img, err = png.DecodeConfig(file)
		if err == nil {
			width, height = int64(img.Width), int64(img.Height)
		}
	}
	return width, height, err
}

func adjustSizes(width int64, height int64, meta interface{}, options []interface{}) ([]interface{}) {
	var real_width int64
	new_options := []interface{} {}

	m := meta.(map[string]interface{})
	textwidth, err := strconv.ParseInt(pf.Stringify(m["latexTextWidth"]), 10, 32)
	if err != nil {
		fmt.Fprintf(os.Stderr, "meta[latexTextWidth]: %s\n", pf.Stringify(m["latexTextWidth"]))
		return options
	}

	// ``ratio'' suppresses other dimension options
	var opt_ratio interface{}
	opt_ratio, options = pf.GetValue(options, "ratio")

	if opt_ratio != nil {
		str, ok := opt_ratio.(string)
		if ok {
			ratio, err := strconv.ParseFloat(str, 64)
			if err == nil {
				real_width =  int64(ratio * float64(textwidth))
				fmt.Fprintf(os.Stderr, "%f %d %d\n", float64(textwidth), width, real_width)
			}
		}
	}

	var opt_width interface{}
	opt_width, options = pf.GetValue(options, "width")
	if real_width == 0 && opt_width != nil {
		// for simplicity, ``width'' determines dimensions of the image
		real_width, err = strconv.ParseInt(opt_width.(string), 10, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ParseInt error: %s\n", opt_width.(string))
		}
	}

	if real_width == 0 {
		real_width = width
	}

	width, height = real_width, real_width * height / width
	width_str  := []interface{} {"width",  strconv.FormatInt(width, 10)  + "pt"}
	height_str := []interface{} {"height", strconv.FormatInt(height, 10) + "pt"}

	for _, option := range options {
		new_options = append(new_options, option)
	}

	new_options = append(new_options, width_str, height_str)
	return new_options
}

func processImage(key string, value interface{}, target string, meta interface{}) interface{} {
	if key == "Image" {
		v := value.([]interface{})

		// attrs might be adjusted
		attrs   := v[0].([]interface{})
		// label   := attrs[0]
		// id      := attrs[1]
		options := attrs[2].([]interface{})
		var new_options []interface{}

		// caption shouldn't be modified!
		caption := v[1].([]interface{})

		// path ( dest[0] ) of the image
		dest := v[2].([]interface{})
		path := dest[0].(string)
		img_width, img_height, err := imageDimensions(path)

		// only adjust the image's size for PDF
		if err == nil && target == "latex" {
			new_options = adjustSizes(img_width, img_height, meta, options)
			attrs[2] = new_options
		}

		t := []string{}
		for _, str := range dest {
			s, ok := str.(string)
			if ok {
				t = append(t, s)
			} else {
				fmt.Fprintf(os.Stderr, "%V is not a string\n", s)
			}
		}
		return pf.Image(attrs, caption, t)
	}
	return nil
}

func main() {
	pf.ToJSONFilter(processImage)
}
