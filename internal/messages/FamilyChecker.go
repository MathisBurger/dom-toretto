package messages

import (
	"fmt"
	"strings"
)

var FamilyDefinitions = []string{"familie", "family", "la familia"}

func FamilyChecker(msg string) bool {
	lowerCase := strings.ToLower(msg)
	fmt.Println(lowerCase)
	for _, el := range FamilyDefinitions {
		if strings.Contains(lowerCase, el) {
			return true
		}
	}
	return false
}
