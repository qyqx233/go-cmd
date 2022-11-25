package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func sha256_(path, raw string) {
	h := sha256.New()
	fmt.Printf("path=[%s], raw=[%s]\n", path, raw)
	var fd *os.File
	var err error
	var n int
	if path != "" {
		fd, err = os.Open(path)
		if err != nil {
			fmt.Println("")
			return
		}
		var data = make([]byte, 4096)
		r := bufio.NewReader(fd)
		for {
			n, err = r.Read(data)
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("sha256 failed", err)
				return
			}
			h.Write(data[:n])
		}
		defer fd.Close()
	} else {
		h.Write([]byte(raw))
	}
	fmt.Printf("%x\n", h.Sum(nil))
}
