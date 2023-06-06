package main

import (
	"errors"
	myBot "github.com/Hugrid-1/minecraftBot/internal/bot"
	"github.com/Tnze/go-mc/bot"
	"log"
)

func main() {
	nBot := myBot.NewBot("GigaChadBot")

	// Login
	err := nBot.Client.JoinServer("cheshuiki.aternos.me:25781")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	// JoinGame
	for {
		if err = nBot.Client.HandleGame(); err == nil {
			panic("HandleGame never return nil")
		}

		if err2 := new(bot.PacketHandlerError); errors.As(err, err2) {
			if err := new(myBot.DisconnectErr); errors.As(err2, err) {
				log.Print("Disconnect, reason: ", err.Reason)
				err := nBot.Client.JoinServer("cheshuiki.aternos.me:25781")
				if err != nil {
					log.Fatal(err)
				}
			} else {
				// normal packet handler error, ignore and continue.
				log.Print(err2)
			}
		} else {
			// if the error is not a PacketHandlerError, the connection is broken.
			// stop the program
			err = nBot.Client.JoinServer("cheshuiki.aternos.me:25781")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
