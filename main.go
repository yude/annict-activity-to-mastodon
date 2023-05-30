package main

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron"
	"github.com/sivchari/gotwtr"
)

var last_updated time.Time

func main() {
	// Init Annict related things
	data, err := fetch_annict()
	if err != nil {
		log.Fatal("Error: Something went wrong on startup. Exiting.")
	}

	last_updated = time.Now().UTC()
	log.Printf("✅ Annict に %s (ID: %d) としてログインしました。\n", data.Activities[0].User.Username, data.Activities[0].User.ID)
	log.Printf("   %s (UTC) 以降のアクティビティを Twitter に投稿します。\n", last_updated.Format("2006/1/2 15:04:05"))

	// Init Twitter API v2
	twtr_client := GetTwitterClient()

	c := cron.New()
	c.AddFunc("@every 5s", func() {
		data, err := fetch_annict()
		if err != nil {
			log.Fatal("Error: Something went wrong. Skipping the tasks.")
		}

		var target []AnnictActivityBody

		for i := 0; i < len(data.Activities); i++ {
			if data.Activities[i].CreatedAt.After(last_updated) {
				target = append(target, data.Activities[i])
			}
		}

		formatted := format_data(target)

		for i := 0; i < len(formatted); i++ {
			log.Println("📝 ツイートします: " + formatted[i])
			_, err := twtr_client.PostTweet(context.Background(), &gotwtr.PostTweetOption{
				Text: formatted[i],
			})
			if err != nil {
				log.Fatal(err)
			}
		}

		last_updated = time.Now().UTC()
	})
	c.Start()

	for {
		time.Sleep(1138800 * time.Hour)
	}
}
