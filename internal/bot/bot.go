package bot

import (
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/bot/msg"
	"github.com/Tnze/go-mc/bot/playerlist"
	"github.com/Tnze/go-mc/chat"
)

type Bot struct {
	Client      *bot.Client
	Player      *basic.Player
	PlayerList  *playerlist.PlayerList
	ChatHandler *msg.Manager
}

type DisconnectErr struct {
	Reason chat.Message
}

func NewBot(botName string) *Bot {
	c := bot.NewClient()
	c.Auth = bot.Auth{Name: botName}
	playerList := playerlist.New(c)

	newBot := &Bot{Client: c, PlayerList: playerList}
	// set handlers
	newBot.setPlayerHandler()
	newBot.setChatHandler()

	return newBot
}

func (b *Bot) setPlayerHandler() {
	p := basic.NewPlayer(b.Client, basic.DefaultSettings, basic.EventsListener{
		Disconnect: b.onDisconnect,
		Death:      b.onDeath,
	})
	b.Player = p
}
func (b *Bot) setChatHandler() {
	chatHandler := msg.New(b.Client, b.Player, b.PlayerList, msg.EventsHandler{
		SystemChat:        b.onSystemChat,
		PlayerChatMessage: b.onPlayerChat,
		DisguisedChat:     b.onDisguisedChat,
	})
	b.ChatHandler = chatHandler
}
