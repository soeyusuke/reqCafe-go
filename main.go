package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/soeyusuke/reqCafe-go/cafe"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, LINE Bot")
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	bot, err := linebot.New(
		os.Getenv("LINE_Channel_Secret"), //linechannelsecret
		os.Getenv("LINE_Channel_Token"),  //linechanneltoken
	)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(w, "Maybe goooood!")
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(111111)
		} else {
			w.WriteHeader(222222)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				switch message.Text {
				case "学食更新":
					cafe.UpdateCafe()
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("学食更新しました。")).Do(); err != nil {
						log.Print(err)
					}
				case "月":
					meshi := cafe.RequestCafeMon()
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(meshi)).Do(); err != nil {
						log.Print(err)
					}
				case "火":
					meshi := cafe.RequestCafeTue()
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(meshi)).Do(); err != nil {
						log.Print(err)
					}
				case "水":
					meshi := cafe.RequestCafeWen()
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(meshi)).Do(); err != nil {
						log.Print(err)
					}
				case "木":
					meshi := cafe.RequestCafeThu()
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(meshi)).Do(); err != nil {
						log.Print(err)
					}
				case "金":
					meshi := cafe.RequestCafeFri()
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(meshi)).Do(); err != nil {
						log.Print(err)
					}
				default:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("学食 と入力してください")).Do(); err != nil {
						log.Print(err)
					}
				}

			} //switch message...
		} //if event.Type...
	} //for _, event ...
} //callbackhandler...

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(
		os.Getenv("LINE_Channel_Secret"), //linechannelsecret
		os.Getenv("LINE_Channel_Token"),  //linechanneltoken
	)
	if err != nil {
		log.Println(err)
	}

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}
