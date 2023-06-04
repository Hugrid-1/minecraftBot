package main

import (
	"errors"
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/chat"
	"log"
)

var (
	c *bot.Client
	p *basic.Player
)

func main() {
	c = bot.NewClient()
	c.Auth = bot.Auth{
		Name: "GigaChadBot",
	}
	p = basic.NewPlayer(c, basic.DefaultSettings, basic.EventsListener{
		Disconnect: onDisconnect,
		Death:      onDeath,
	})

	// Login
	err := c.JoinServer("cheshuiki.aternos.me:25781")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	// JoinGame
	for {
		if err = c.HandleGame(); err == nil {
			panic("HandleGame never return nil")
		}

		if err2 := new(bot.PacketHandlerError); errors.As(err, err2) {
			if err := new(DisconnectErr); errors.As(err2, err) {
				log.Print("Disconnect, reason: ", err.Reason)
				err := c.JoinServer("cheshuiki.aternos.me:25781")
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
			//log.Fatal(err)
			err := c.JoinServer("cheshuiki.aternos.me:25781")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

type DisconnectErr struct {
	Reason chat.Message
}

func (d DisconnectErr) Error() string {
	return "disconnect: " + d.Reason.ClearString()
}

func onDisconnect(reason chat.Message) error {
	// return an error value so that we can stop main loop
	return DisconnectErr{Reason: reason}
}

func onDeath() error {
	log.Println("Died and Respawned")
	// If we exclude Respawn(...) then the player won't press the "Respawn" button upon death
	return p.Respawn()
}
