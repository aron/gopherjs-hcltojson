package main

import (
	"github.com/tmccombs/hcl2json/convert"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Module.Get("exports").Set("hcltojson", hcltojson);
}

func hcltojson(input string) string {
	jsonBytes, err := convert.Bytes([]byte(input), "", convert.Options{})
	if err != nil {
		panic(err.Error())
	}
	return string(jsonBytes)
}
