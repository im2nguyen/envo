package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
  "math/rand"
)

var maskMethod = flag.String("maskMethod", "trunc", "Masking method for environment variable value.")
var truncLength = flag.Int("truncLength", 0, "Length to truncate environment variable value.")
var r = regexp.MustCompile(`[^/-]`)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	flag.StringVar(maskMethod, "m", "trunc", "Masking method for environment variable value.")
	flag.IntVar(truncLength, "t", 0, "Length to truncate environment variable value.")
}

func main() {
	flag.Parse()

	for _, pair := range os.Environ() {
		processValue(pair, *maskMethod, *truncLength)
	}
}

func processValue(val string, method string, truncLength int) {
	pair := strings.SplitN(val, "=", 2)
	k := pair[0]
	v := pair[1]

	if truncLength > 0 {
		if truncLength >= len(v) {
			fmt.Printf("%s=%+v...\n", k, v[0:int64(len(v)/3)])
		} else if len(v) > truncLength {
			fmt.Printf("%s=%+v...\n", k, v[0:truncLength])
		} else {
			fmt.Printf("%s=%+v...\n", k, v[0])
		}
		return
	}

	switch method {
	case "trunc":
		fmt.Printf("%s=%+v...\n", k, v[0:int64(len(v)/3)])
	case "random":
		b := make([]byte, len(v))
		for i := range b {
      b[i] = charset[rand.Int63() % int64(len(charset))]
		}
		fmt.Printf("%s=%+v\n", k, string(b))
	default:
		fmt.Printf("%s=%+v\n", k, r.ReplaceAllString(v, method))
	}
}
