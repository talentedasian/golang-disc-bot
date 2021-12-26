package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func InviteFriends(s *discordgo.Session, m *discordgo.MessageCreate,
	role *discordgo.Role) {

	gld, gErr := FindGuild(s, m)

	if gErr != nil {
		fmt.Println("Could not find guild")
	}

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}

	mentionedRole := LfdRole(gld).Mention()

	msg := mentionedRole + " G na potangina ang tagal"
	s.ChannelMessageSend(c.ID, msg)
}

func LfdRole(gld *discordgo.Guild) *discordgo.Role {
	var role *discordgo.Role
	for _, rl := range gld.Roles {
		if rl.Name == "hl2" {
			role = rl
		}
	}

	if role == nil {
		log.Println("Role \"hl2\" couldn't be found on the guild")
	}

	return role
}

func FindGuild(s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.Guild, error) {
	return s.Guild(m.GuildID)
}
