package main

import (
	"github.com/gsdenys/pdgen/cmd"
	"github.com/gsdenys/pdgen/pkg/services/translate"
)

func main() {
	translate.Register()
	cmd.Execute()
}
