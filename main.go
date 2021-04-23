package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/tkanos/gonfig"
)

// global config
var config Config

func main() {
	// parse config

	// check if env var
	env := os.Getenv("ELASTIC_PROXY_ENV")
	if env == "YES" {
		// pull config from env vars
		config.Port = 8080
		config.WebhookURL = os.Getenv("ELASTIC_PROXY_WEBHOOK_URL")
		rich := os.Getenv("ELASTIC_PROXY_RICH_MESSAGES")
		if rich == "NO" {
			config.RichMessages = false
		} else {
			config.RichMessages = true
		}
	} else {
		// don't pull from env, look for config.json
		err := gonfig.GetConf("config.json", &config)
		if err != nil {
			log.Fatal("Failed to parse config: " + err.Error())
		}
	}

	log.Printf("Using webhook URL: %s\n", config.WebhookURL)

	http.HandleFunc("/webhook", handleIncomingWebhook)

	log.Printf("Listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}

func handleIncomingWebhook(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to get incoming webhook body: " + err.Error())
		return
	}

	var webhook ElasticWebhook
	err = json.Unmarshal(buf, &webhook)
	if err != nil {
		log.Println("Failed to unmarshal Elastic webhook: " + err.Error())
		return
	}

	log.Println("Handling notifications for alert " + webhook.AlertID)
	sendDiscordMessage(webhook)
	log.Println("Finished handling notifications for alert " + webhook.AlertID)
}

func sendDiscordMessage(webhook ElasticWebhook) {
	red := 0xff0000
	green := 0x00ff00

	var discordMsg DiscordWebhook

	if config.RichMessages {
		var embed DiscordEmbed

		embed.Title = fmt.Sprintf("Elastic Alert: %s", webhook.AlertName)
		embed.Description = fmt.Sprintf("Triggered at: %s\n\n**%s**", webhook.State.LastTriggeredAt, webhook.Context.Message)

		t, err := strconv.ParseBool((webhook.State.IsTriggered))
		if err != nil {
			log.Println("Failed to parse state.isTriggered bool")
			return
		}

		if t {
			embed.Color = red
		} else {
			embed.Color = green
		}

		discordMsg.Embeds = []DiscordEmbed{embed}
	} else {
		discordMsg.Content = webhook.Context.Message
	}

	if !makeDiscordRequest(discordMsg) {
		log.Printf("Discord message failed to send for alert \"%s\"", webhook.AlertID)
	}
}

func makeDiscordRequest(msg DiscordWebhook) bool {
	jsonBody, err := json.Marshal(msg)
	if err != nil {
		log.Println("Failed to marshal Discord webhook: " + err.Error())
		return false
	}

	resp, err := http.Post(config.WebhookURL+"?wait=true", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Println("Failed to make Discord webhook request: " + err.Error())
		return false
	}

	defer resp.Body.Close()

	log.Println("Sent Discord webhook")
	return true
}
