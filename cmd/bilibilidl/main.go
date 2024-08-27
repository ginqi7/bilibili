package main

import (
	"fmt"
	"os"

	bilibili "github.com/ginqi7/bilibili/pkg/client"
)

var (
	client *bilibili.Client
)

func init() {
	client = &bilibili.Client{}
}

func main() {
	exitOnError(rootCmd.Execute())
}

func exitOnError(err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
