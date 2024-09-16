package main

import (
	"fmt"
	"os"

	"mrg"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "wrong input")
		return
	}
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	if arg1 != "sample.txt" {
		fmt.Fprintln(os.Stderr, "wrong input")

		return
	} else if arg2 != "result.txt" {
		fmt.Println("wrong input")
		return
	}
	file, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		return
	}

	result := mrg.UpLowCap(string(file))
	result = mrg.HexBin(result)
	result = mrg.UpLowCapNbr(string(result))
	result = mrg.Punc(result)
	result = mrg.Quote(result)
	result = mrg.Aandan(result)
	result = result[:len(result)-4]
	err2 := os.WriteFile(os.Args[2], []byte(result), 0o644)
	if err2 != nil {
		fmt.Fprintln(os.Stderr, "Error writing file:", err)
		return
	}
}
