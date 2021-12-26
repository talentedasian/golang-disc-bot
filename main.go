package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("Token")
	discord, _ := discordgo.New("Bot " + token)

	discord.AddHandler(ready)

	discord.AddHandler(onMessageLfd)

	discord.AddHandler(lfdMessage)

	discord.AddHandler(lfdRole)

	opErr := discord.Open()
	if opErr != nil {
		fmt.Println("Could not start bot with error: ", opErr)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func ready(s *discordgo.Session, m *discordgo.Ready) {
	fmt.Println("Bot is ready with token: ", s.Token)
}

func onMessageLfd(s *discordgo.Session, m *discordgo.MessageCreate) {
	if isBotChat(s, m) && IsGiffRole(m.Content) {
		return
	}

	gld, gErr := FindGuild(s, m)
	if gErr != nil {
		return
	}

	role := LfdRole(gld)

	if role == nil {
		fmt.Println("Role \"hl2\" not found")
	}

	if m.Message.Content != "!lfd" {
		return
	}

	InviteFriends(s, m, role)
}

func isBotChat(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return s.State.User.ID == m.Author.ID
}

func lfdMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if IsGiffRole(m.Content) {
		return
	}

	msg := m.Content
	gld, _ := FindGuild(s, m)

	role := LfdRole(gld)
	if role == nil {
		return
	}

	if strings.Contains(msg, "lfd") {
		InviteFriends(s, m, role)
	}
}

func lfdRole(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !IsGiffRole(m.Message.Content) || isBotChat(s, m) {
		return
	}

	chId := m.ChannelID
	gld, gldErr := FindGuild(s, m)
	if gldErr != nil {
		log.Println("User did not invoke the command on the guild")

		s.ChannelMessageSend(chId, "Can only be used inside the guild")
		return
	}

	rlErr := s.GuildMemberRoleAdd(gld.ID, m.Author.ID, LfdRole(gld).ID)
	if rlErr != nil {
		log.Println(`Bot encountered an error adding a role to a user. The bot probably doesn't have
		the permission to do such an action`)

		s.ChannelMessageSend(chId, "Sorry, I can't add roles")
		return
	}

	s.ChannelMessageSend(chId, "Role added")
}
