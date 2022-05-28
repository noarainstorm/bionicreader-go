package main

import (
	"log"
	"math"
	"strings"
	utf "unicode/utf8"
)

const (
	pre    = "\033[1m"
	out    = "\033[0m"
	italic = "\033[37;2m"
)

type bionic struct {
	Out string
}

func (tk *bionic) getBionic(text string) {
	build := new(strings.Builder)
	splitted := strings.Split(text, " ")
	for _, em := range splitted {
		if utf.RuneCountInString(em) != 1 {
			bWord := pre + em[0:makelen(em)] + out + italic + em[makelen(em):] + out
			_, err := build.WriteString(bWord + " ")
			if err != nil {
				log.Fatal(err)
			}
		}
		if utf.RuneCountInString(em) == 1 {
			_, err := build.WriteString(em + " ")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	tk.Out = build.String()
}

func makelen(word string) int {
	var count int
	clen := int(math.Round(float64(utf.RuneCountInString(word) / 2)))
	var cl int
	for _, em := range word {
		if cl == clen {
			break
		}
		count += len(string(em))
		cl++

	}
	return count
}
