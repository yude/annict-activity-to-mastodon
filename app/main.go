package main

import (
	"log"
	"time"

	"github.com/robfig/cron"
)

var last_updated time.Time

func main() {
	// Init Annict related things
	InitConfig()

	hasMastodon := conf.Credentials.MastodonCredentials.MastodonUrl != "" && conf.Credentials.MastodonCredentials.MastodonToken != ""
	hasMisskey := conf.Credentials.MisskeyCredentials.MisskeyUrl != "" && conf.Credentials.MisskeyCredentials.MisskeyToken != ""
	if !hasMastodon && !hasMisskey {
		log.Fatal("Error: Mastodon または Misskey の認証情報が設定されていません。config.toml を確認してください。")
	}

	data, err := fetch_annict()
	if err != nil {
		log.Fatal("Error: Something went wrong on startup. Exiting.")
	}

	last_updated = time.Now().UTC()
	log.Printf("✅ Annict に %s (ID: %d) としてログインしました。\n", data.Activities[0].User.Username, data.Activities[0].User.ID)
	if hasMastodon && hasMisskey {
		log.Printf("   %s (UTC) 以降のアクティビティを Mastodon / Misskey に投稿します。\n", last_updated.Format("2006/1/2 15:04:05"))
	} else if hasMastodon {
		log.Printf("   %s (UTC) 以降のアクティビティを Mastodon に投稿します。\n", last_updated.Format("2006/1/2 15:04:05"))
	} else {
		log.Printf("   %s (UTC) 以降のアクティビティを Misskey に投稿します。\n", last_updated.Format("2006/1/2 15:04:05"))
	}

	c := cron.New()
	c.AddFunc("@every 15m", func() {
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
			log.Println("📝 投稿します: " + formatted[i])
			if hasMastodon {
				err := PostToMastodon(formatted[i])
				if err != nil {
					log.Println("Mastodon Error:", err)
				}
			}
			if hasMisskey {
				err := PostToMisskey(formatted[i])
				if err != nil {
					log.Println("Misskey Error:", err)
				}
			}
		}

		last_updated = time.Now().UTC()
	})
	c.Start()

	for {
		time.Sleep(1138800 * time.Hour)
	}
}
