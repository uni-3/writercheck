package main

import (
	"writercheck"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(writercheck.Analyzer) }
