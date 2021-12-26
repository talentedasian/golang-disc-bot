package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := "OTI0NTQwOTM5MTQ0Mzk2ODIx.YcgD2Q.YmHJ35hQa79VYFwJ2hTCXr-yDoU"
	discord, _ := discordgo.New("Bot " + token)

	discord.AddHandler(ready)
	discord.AddHandler(onMessageLfd)

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

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}

	g, gErr := s.Guild(m.GuildID)
	if gErr != nil {
		return
	}

	roles := g.Roles

	var role *discordgo.Role
	for _, rl := range roles {
		if rl.Name == "hl2" {
			role = rl
		}
	}

	if role == nil {
		fmt.Println("Role \"hl2\" not found")
	}

	if m.Message.Content != "!lfd" {
		return
	}

	mentionedRole := role.Mention()
	msg := mentionedRole + " G na potangina ang tagal"
	s.ChannelMessageSend(c.ID, msg)
}

func isBotChat(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return s.State.User.ID == m.Author.ID
}
