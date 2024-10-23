package main

import (
	"flag"
	"github.com/username/Adviser-Bot/Clients/Telegram"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	tgClient := Telegram.New(mustToken())
	// token = flags.Get(token)

	// fetcher = fetcher.New(tgClient)
	// processor = processor.New(tgClient)
	// consumer.Start(fetcher,processor)
}

func mustToken() string {
	var token = flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
