package main

import (
	"fmt"
	"github.com/MathisBurger/dom-toretto/internal/handler"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if os.Getenv("mode") != "prod" {
		err := godotenv.Load()
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("Starting Dom Toretto")
	fmt.Println("Family first!")
	fmt.Println("token:", os.Getenv("botToken"))
	fmt.Println("-------------------------------------------------------------------")

	session, err := discordgo.New("Bot " + os.Getenv("botToken"))
	if err != nil {
		fmt.Println("Error while creating discord session")
		return
	}

	session.AddHandler(handler.MessageHandler)

	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	err = session.Open()
	if err != nil {
		fmt.Println("Cannot connect to discord websocket")
		return
	}

	err = session.UpdateListeningStatus("family")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("The bot is running now. Terminate by using CTRL-C")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	session.Close()
}
