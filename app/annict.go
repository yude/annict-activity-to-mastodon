package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fetch_annict() (*AnnictActivity, error) {
	// Parse config.toml
	conf := GetConfig()

	url := "https://api.annict.com/v1/activities?access_token=" + conf.Credentials.AnnictCredentials.AnnictKey + "&sort_id=desc&filter_username=" + conf.Credentials.AnnictCredentials.AnnictUsername + "&per_page=10"

	resp, err := http.Get(url)
	if err != nil {
		log.Print("Error: Failed to retrieve data from Annict's API.", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Print("Warning: API returned non-200 code. (", resp.StatusCode, ")")
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	jsonBytes := ([]byte)(byteArray)
	data := new(AnnictActivity)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		log.Print("Warning: Failed to parse Annict's JSON: ", err)
		log.Println("Raw response: ", data)
		return data, err
	}

	return data, nil
}

func format_data(data []AnnictActivityBody) []string {
	conf := GetConfig()

	var texts []string

	for i := 0; i < len(data); i++ {
		if data[i].Action == "create_record" {
			texts = append(texts, fmt.Sprintf("%s %s を見ました https://annict.com/@%s/records #%s", data[i].Work.Title, data[i].Episode.NumberText, conf.Credentials.AnnictUsername, data[i].Work.TwitterHashtag))
		}

		if data[i].Action == "create_status" {
			if data[i].Status.Kind != "no_select" {
				var status_kind_fmt string

				switch data[i].Status.Kind {
				case "wanna_watch":
					status_kind_fmt = "見たい"
				case "watching":
					status_kind_fmt = "見てる"
				case "watched":
					status_kind_fmt = "見た"
				case "on_hold":
					status_kind_fmt = "一時中断"
				case "stop_watching":
					status_kind_fmt = "視聴中止"
				}

				texts = append(texts, fmt.Sprintf("%s の視聴ステータスを「%s」にしました https://annict.com/@%s/%s", data[i].Work.Title, status_kind_fmt, conf.Credentials.AnnictUsername, data[i].Status.Kind))
			}
		}
	}

	return texts
}
