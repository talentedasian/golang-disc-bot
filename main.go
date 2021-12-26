package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := ""
	discord, _ := discordgo.New("Bot " + token)

	discord.AddHandler(ready)

	discord.AddHandler(onMessageLfd)

	discord.AddHandler(lfdMessage)

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
	if isBotChat(s, m) {
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
	msg := m.Message.Content
	gld, _ := FindGuild(s, m)

	role := LfdRole(gld)
	if role == nil {
		return
	}

	if strings.Contains(msg, "lfd") {
		InviteFriends(s, m, role)
	}
}
