package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const token string = "OTM4NTE1NjU4MjUyMDk5NjM1.Yfra0Q.c3uTdFk3nS_3Qje0cID8cvlVOA0"

var BotId = ""

func main() {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotId = u.ID

	dg.AddHandler(messageHandler)

	err = dg.Open()

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(BotId)
		return
	}

	fmt.Println("Bot is running!")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotId {
		return
	}
	fmt.Println(m.Content)

	_, _ = s.ChannelMessageSend("938515064221208657", "KEKW")
}
