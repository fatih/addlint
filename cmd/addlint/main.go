package main

import (
	"github.com/fatih/addlint/addcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(addcheck.Analyzer)
}
