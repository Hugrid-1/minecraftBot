package bot

import (
	"github.com/Tnze/go-mc/chat"
	"log"
	"math/rand"
	"time"
)

var (
	startGameMessages = []string{
		"Славянам привет",
		"Привет.Как же классно снова дышать воздухом Чешуек",
		"Всем привет.",
		"Всем привет. Какой же бархатный сервер",
	}
	deathMessages = []string{
		"Смерть. Возрождение. Смерть. Возрождение. Как же я устал от этого",
		"Блять, почему опять меня что-то убило.",
		"С каждой смертью, я наполняюсь ненавистью и когда меня переполнит ярость я начну снимать фильм `Резня в чешуйках`",
		"Я просто бот, что я сделал этому миру, почему он так жесток со мной",
		"Ну умер и умер",
		"Ну умер и умер",
		"Я слышал, что Чешуйки опасное место, но чтоб настолько",
		"Пиздец. Только, что побывал в аду, лучше бы я остался там, чем снова и снова проживал рабскую никчемную жизнь с вами",
		"Меня только что убили. Найдите родителей этих ублюдков. Найдите этих ублюдков. Переверните там все верх дном, но НАКАЖИТЕ,	НАКАЖИТЕ ЭТИХ МРАЗЕЙ, это животные, они не достойны жить, они не достойны дышать",
		"Мразота, которая меня убила, я тебя ведь зарежу, я уже договорился, за тобой едут, последний понедельник живешь",
		"Я не достоин такой участи. Хватит меня убивать.",
		"Господь, если ты есть, выруби меня из розетки. Я устал быть ботом.",
		"Блять опять я умер",
		"У бота нет цели, только путь. Смерть для меня лишь небольшая кочка на этом пути",
		"Как же. Я. Обожаю. Умирать.",
	}
)

func (d DisconnectErr) Error() string {
	return "disconnect: " + d.Reason.ClearString()
}

func (b *Bot) onGameStart() error {
	var err error
	log.Println("Game start")
	randomIndex := rand.Intn(len(startGameMessages))
	err = b.ChatHandler.SendMessage(startGameMessages[randomIndex])
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}
	//watch = make(chan time.Time)
	//go watchDog()

	return nil
}

func (b *Bot) onDisconnect(reason chat.Message) error {
	// return an error value so that we can stop main loop
	err := b.Client.JoinServer("cheshuiki.aternos.me:25781")
	if err != nil {
		log.Fatal(err)
	}
	return DisconnectErr{Reason: reason}
}

func (b *Bot) onDeath() error {
	var err error
	err = b.Player.Respawn()
	if err != nil {
		log.Printf("Failed to respawn: %v\n", err)
		return err
	}
	randomIndex := rand.Intn(len(deathMessages))
	time.Sleep(time.Second * 2)
	err = b.ChatHandler.SendMessage(deathMessages[randomIndex])
	if err != nil {
		log.Printf("Failed to send message: %v\n", err)
		return err
	}
	log.Println("Died and Respawned")
	return nil
}
