package main

import (
	"encoding/json"
	"fmt"
	"irctc-telegram-bot/model"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
)

func main() {
	errorsFile, err := os.OpenFile("errors.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	InfoFile, err := os.OpenFile("info.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	ErrorLogger = log.New(errorsFile, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(InfoFile, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	bot, err := tgbotapi.NewBotAPI(TelegramAccessKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		InfoLogger.Printf("chatID: %d, MessageID: %v, FirstName : %s, LastName : %s,UserName : %s, Text: %s ", update.Message.Chat.ID, update.Message.MessageID, update.Message.From.FirstName, update.Message.From.LastName, update.Message.From.UserName, update.Message.Text)
		if update.Message.Text == "/start" {
			update.Message.Text = "Please Enter your PNR Number"
		} else {
			update.Message.Text = getPNRInfo(update)

		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		InfoLogger.Printf("chatID: %d, MessageID: %v, FirstName : %s, LastName : %s,UserName : %s,SentText: %s ", update.Message.Chat.ID, msg.ReplyToMessageID, update.Message.From.FirstName, update.Message.From.LastName, update.Message.From.UserName, update.Message.Text)
		InfoLogger.Printf("msg :%#v", msg)

		bot.Send(msg)
	}

}

func getPNRInfo(update tgbotapi.Update) string {
	client := &http.Client{}
	fmt.Printf("update is----------------------- %#v", update.Message)
	req, err := http.NewRequest("GET", "https://pnr-status-indian-railway.p.rapidapi.com/rail/"+update.Message.Text, nil)
	req.Header.Add("x-rapidapi-key", RapidAPIKey)
	req.Header.Add("x-rapidapi-host", RapidAPIHost)
	resp, err := client.Do(req)
	if err != nil {
		ErrorLogger.Printf("chatID: %d, MessageID: %v, FirstName : %s, LastName : %s,UserName : %s, Text: %s ", update.Message.Chat.ID, update.Message.MessageID, update.Message.From.FirstName, update.Message.From.LastName, update.Message.From.UserName, update.Message.Text)
	}
	InfoLogger.Println("resp", resp)
	var response model.ResponseModel
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		ErrorLogger.Printf("chatID: %d, MessageID: %v, FirstName : %s, LastName : %s,UserName : %s, Text: %s ", update.Message.Chat.ID, update.Message.MessageID, update.Message.From.FirstName, update.Message.From.LastName, update.Message.From.UserName, update.Message.Text)
	}
	data := fmt.Sprintf("ChartStatus:%s, train name: %s", response.ChartStatus, response.TrainName)
	InfoLogger.Println("response", response)
	return data
}
