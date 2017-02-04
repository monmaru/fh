package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/ripemd160"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please input the file path.")
		fmt.Println("Example: fh.exe hoge/moge.txt")
		return
	}

	f := os.Args[1]
	fmt.Println("MD5       : ", compute(f, md5.New()))
	fmt.Println("SHA1      : ", compute(f, sha1.New()))
	fmt.Println("SHA256    : ", compute(f, sha256.New()))
	fmt.Println("SHA-384   : ", compute(f, sha512.New384()))
	fmt.Println("SHA-512   : ", compute(f, sha512.New()))
	fmt.Println("RIPEMD160 : ", compute(f, ripemd160.New()))
}

func compute(filePath string, hash hash.Hash) string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := io.Copy(hash, f); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hash.Sum(nil))
}
