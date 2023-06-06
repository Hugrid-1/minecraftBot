package bot

import (
	"github.com/Tnze/go-mc/chat"
	"log"
)

func (b *Bot) onSystemChat(c chat.Message, overlay bool) error {
	log.Printf("System Chat: %v, Overlay: %v", c, overlay)
	return nil
}

func (b *Bot) onPlayerChat(c chat.Message, _ bool) error {
	log.Println("Player Chat:", c)
	return nil
}

func (b *Bot) onDisguisedChat(c chat.Message) error {
	log.Println("Disguised Chat:", c)
	return nil
}
