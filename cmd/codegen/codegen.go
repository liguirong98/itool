package main

import (
	"fmt"
	"itool/tools/codegen"
	"os"
)

func main() {

	if err := codegen.CodeGen.Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
