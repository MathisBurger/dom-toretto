package handler

import (
	"github.com/MathisBurger/dom-toretto/internal/messages"
	"github.com/bwmarrin/discordgo"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID != s.State.User.ID && messages.FamilyChecker(m.Content) {
		_, _ = s.ChannelMessageSendReply(
			m.ChannelID,
			messages.RandomFamilyMessage(),
			m.MessageReference,
		)
	}
}
