package main

import (
	"regexp"
)

// A source of news, a company that
// works with news.
type NewsSource string

const elNacional NewsSource = "El Nacional"
const cnn NewsSource = "CNN"
const elUniversal NewsSource = "El Universal"

var elNacionalHostRegex = regexp.MustCompile(`elnacional`)
var elUniversalHostRegex = regexp.MustCompile(`eluniversal`)
var cnnHostRegex = regexp.MustCompile(`cnn`)

// Parses a string to know if a URL host matches
// one of the supported news sources.
func getNewsSource(host string) NewsSource {
	if elNacionalHostRegex.MatchString(host) {
		return elNacional
	} else if elUniversalHostRegex.MatchString(host) {
		return elUniversal
	} else if cnnHostRegex.MatchString(host) {
		return cnn
	}
	panic("Impossible to determine news source for host: " + host)
}
