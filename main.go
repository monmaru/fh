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
	fmt.Print("MD5       : ")
	fmt.Println(compute(f, md5.New()))
	fmt.Print("SHA1      : ")
	fmt.Println(compute(f, sha1.New()))
	fmt.Print("SHA256    : ")
	fmt.Println(compute(f, sha256.New()))
	fmt.Print("SHA-384   : ")
	fmt.Println(compute(f, sha512.New384()))
	fmt.Print("SHA-512   : ")
	fmt.Println(compute(f, sha512.New()))
	fmt.Print("RIPEMD160 : ")
	fmt.Println(compute(f, ripemd160.New()))
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
