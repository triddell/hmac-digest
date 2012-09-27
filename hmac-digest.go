package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"flag"
	"fmt"
	"hash"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: hmac-digest [textToEncode] [secretKey]\n")
		os.Exit(1)
	}

	var h hash.Hash = hmac.New(sha1.New, []byte(args[1]))
	h.Write([]byte(args[0]))
	fmt.Println(base64.URLEncoding.EncodeToString(h.Sum(nil)))
}
