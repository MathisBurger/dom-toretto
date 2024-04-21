package messages

import (
	"math/rand"
	"time"
)

var FamilyMessages = []string{
	"Family is very important!",
	"Nothing over family!",
	"Did you say family?",
	"I would kill for my family!",
	"I have no friends, I have family!",
	"My family is everything",
	"I would learn Java for my family",
	"Nothing over a corona beer",
	"My team is my family",
	"CSGO team is family",
	"I rush B with my family",
	"I rush A with my family",
	"My family is more important than anything else",
	"I only shit on family toilets",
	"Only people who drive Audi are family",
	"My family is the best thing I know",
}

func RandomFamilyMessage() string {
	rand.Seed(time.Now().UnixNano())
	return FamilyMessages[rand.Intn(len(FamilyMessages))]
}
