package messages

import (
	"strings"
)

var FamilyDefinitions = []string{
	"familie",
	"family",
	"la familia",
	"fam",
	"alabama",
	"sweet home",
	"stepsis",
	"stepsister",
	"step sister",
	"stepbro",
	"stepbrother",
	"step brother",
	"step bro",
	"step sis",
}

func FamilyChecker(msg string) bool {
	lowerCase := strings.ToLower(msg)
	for _, el := range FamilyDefinitions {
		if strings.Contains(lowerCase, el) {
			return true
		}
	}
	return false
}
