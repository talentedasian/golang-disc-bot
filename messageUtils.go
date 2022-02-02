package main

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func InviteFriends(s *discordgo.Session, m *discordgo.MessageCreate,
	role *discordgo.Role) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}

	mentionedRole := role.Mention()

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

func TekkenRole(gld *discordgo.Guild) *discordgo.Role {
	var role *discordgo.Role
	for _, rl := range gld.Roles {
		if rl.Name == "tekken" {
			role = rl
		}
	}

	if role == nil {
		log.Println("Role \"tekken\" couldn't be found on the guild")
	}

	return role
}

func FindGuild(s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.Guild, error) {
	return s.Guild(m.GuildID)
}

func IsGiffRole(msg string) bool {
	return strings.EqualFold("giff me lfd role", msg)
}

func IsGiffTekkenRole(msg string) bool {
	return strings.EqualFold("giff me tekken role", msg)
}
