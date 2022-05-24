package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	pre    = "\033[1m"
	out    = "\033[0m"
	italic = "\033[37;2m"
)

type bionic struct {
	Out string
}

func (tk *bionic) get(text string) {
	build := new(strings.Builder)
	splitted := strings.Split(text, " ")
	for _, em := range splitted {
		if len(em) != 1 {
			half := int(math.Round(float64(len(em)) / 2))
			bWord := pre + em[0:half] + out + italic + em[half:] + out
			fmt.Fprintf(build, "%s ", bWord)
		}
		if len(em) == 1 {
			fmt.Fprintf(build, "%s ", em)
		}
	}
	tk.Out = build.String()
}
