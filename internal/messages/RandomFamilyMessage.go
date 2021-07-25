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
}

func RandomFamilyMessage() string {
	rand.Seed(time.Now().UnixNano())
	return FamilyMessages[rand.Intn(len(FamilyMessages))]
}
