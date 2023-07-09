package main

import (
	"github.com/vzvu3k6k/testcasesmap"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(testcasesmap.Analyzer) }
