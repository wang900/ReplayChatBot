package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	entity "github.com/wang900/ReplayChatBot/Bot/entities"
)

const token string = ""

var BotID = ""

var anime []entity.Anime

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

	BotID = u.ID
	generateTestdata()

	dg.AddHandler(messageHandler)

	err = dg.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}
	fmt.Println(m.Content)
	var newMessage entity.ChatMessage
	var err error
	newMessage.Message = m.Content
	newMessage.AuthorID = m.Author.ID
	newMessage.TimeStamp, err = m.Timestamp.Parse()
	if err != nil {
		fmt.Println(err.Error())
	}
	anime[0].Episodes[0].Chat.ChatMessages = append(anime[0].Episodes[0].Chat.ChatMessages, newMessage)
	for i := 0; i < len(anime[0].Episodes[0].Chat.ChatMessages); i++ {
		currentChatMessage := anime[0].Episodes[0].Chat.ChatMessages[i]
		_, _ = s.ChannelMessageSend("938515064221208657", currentChatMessage.Message+" "+currentChatMessage.TimeStamp.String())
	}

}

func generateTestdata() {
	var testChatMessages []entity.ChatMessage

	var testChat entity.Chat
	testChat.ID = 1
	testChat.ChatMessages = testChatMessages

	var testEpisodes []entity.Episode
	var testEpisode entity.Episode
	testEpisode.ID = 1
	testEpisode.Value = 1
	testEpisode.Chat = testChat
	testEpisodes = append(testEpisodes, testEpisode)

	var testAnime entity.Anime
	testAnime.ID = 1
	testAnime.Name = "Tim anime"
	testAnime.Episodes = testEpisodes

	anime = append(anime, testAnime)
}
