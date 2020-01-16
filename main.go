package main

import (
	"log"
	"time"

	"github.com/Inkp/cian-fetch/db"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func main() {
	var xlsURL = "<xls_url>"
	var tgToken = "<token>"
	var chatID int64 = 1

	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	//

	db, err := db.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	for {
		offers, err := fetch(xlsURL)
		if err != nil {
			panic(err)
		}
		for _, offer := range offers {
			if !db.Exists(offer.ID) {
				log.Printf("Found new!! %v\n", offer.URL)
				msg := tgbotapi.NewMessage(chatID, offer.URL)
				_, err := bot.Send(msg)
				if err != nil {
					log.Fatalln(err)
					continue
				}
				db.Save(offer.ID)
			}
		}
		time.Sleep(5 * time.Minute)
	}
}
