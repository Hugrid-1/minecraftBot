package main

import (
	"errors"
	"github.com/Hugrid-1/minecraftBot/config"
	myBot "github.com/Hugrid-1/minecraftBot/internal/bot"
	"github.com/Hugrid-1/minecraftBot/internal/router"
	"github.com/Hugrid-1/minecraftBot/internal/server/httpserver"
	"github.com/Tnze/go-mc/bot"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appConfig := config.NewAppConfig()

	appRouter := router.NewRouter()
	server := httpserver.NewHTTPServer(appConfig.ServerSettings, appRouter)
	_ = server

	nBot := myBot.NewBot("GigaChadBot")

	// Login
	err = nBot.Client.JoinServer("cheshuiki.aternos.me:25781")
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
