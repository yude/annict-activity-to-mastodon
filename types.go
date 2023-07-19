package main

import "time"

type CredentialsConfig struct {
	AnnictCredentials   `toml:"annict"`
	MastodonCredentials `toml:"mastodon"`
}

type AnnictCredentials struct {
	AnnictKey      string `toml:"key"`
	AnnictUsername string `toml:"username"`
}

type MastodonCredentials struct {
	MastodonUrl   string `toml:"domain"`
	MastodonToken string `toml:"access_token"`
}

type Config struct {
	Credentials CredentialsConfig `toml:"credentials"`
}

type AnnictActivityBody struct {
	ID   int `json:"id"`
	User struct {
		ID                 int       `json:"id"`
		Username           string    `json:"username"`
		Name               string    `json:"name"`
		Description        string    `json:"description"`
		URL                string    `json:"url"`
		AvatarURL          string    `json:"avatar_url"`
		BackgroundImageURL string    `json:"background_image_url"`
		RecordsCount       int       `json:"records_count"`
		FollowingsCount    int       `json:"followings_count"`
		FollowersCount     int       `json:"followers_count"`
		WannaWatchCount    int       `json:"wanna_watch_count"`
		WatchingCount      int       `json:"watching_count"`
		WatchedCount       int       `json:"watched_count"`
		OnHoldCount        int       `json:"on_hold_count"`
		StopWatchingCount  int       `json:"stop_watching_count"`
		CreatedAt          time.Time `json:"created_at"`
	} `json:"user"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
	Work      struct {
		ID              int    `json:"id"`
		Title           string `json:"title"`
		TitleKana       string `json:"title_kana"`
		TitleEn         string `json:"title_en"`
		Media           string `json:"media"`
		MediaText       string `json:"media_text"`
		ReleasedOn      string `json:"released_on"`
		ReleasedOnAbout string `json:"released_on_about"`
		OfficialSiteURL string `json:"official_site_url"`
		WikipediaURL    string `json:"wikipedia_url"`
		TwitterUsername string `json:"twitter_username"`
		TwitterHashtag  string `json:"twitter_hashtag"`
		SyobocalTid     string `json:"syobocal_tid"`
		MalAnimeID      string `json:"mal_anime_id"`
		Images          struct {
			RecommendedURL string `json:"recommended_url"`
			Facebook       struct {
				OgImageURL string `json:"og_image_url"`
			} `json:"facebook"`
			Twitter struct {
				MiniAvatarURL     string `json:"mini_avatar_url"`
				NormalAvatarURL   string `json:"normal_avatar_url"`
				BiggerAvatarURL   string `json:"bigger_avatar_url"`
				OriginalAvatarURL string `json:"original_avatar_url"`
				ImageURL          string `json:"image_url"`
			} `json:"twitter"`
		} `json:"images"`
		EpisodesCount  int    `json:"episodes_count"`
		WatchersCount  int    `json:"watchers_count"`
		ReviewsCount   int    `json:"reviews_count"`
		NoEpisodes     bool   `json:"no_episodes"`
		SeasonName     string `json:"season_name"`
		SeasonNameText string `json:"season_name_text"`
	} `json:"work"`
	Status struct {
		Kind string `json:"kind"`
	} `json:"status,omitempty"`
	Episode struct {
		ID                  int     `json:"id"`
		Number              float64 `json:"number"`
		NumberText          string  `json:"number_text"`
		SortNumber          int     `json:"sort_number"`
		Title               string  `json:"title"`
		RecordsCount        int     `json:"records_count"`
		RecordCommentsCount int     `json:"record_comments_count"`
	} `json:"episode,omitempty"`
	Record struct {
		ID            int       `json:"id"`
		Comment       string    `json:"comment"`
		Rating        any       `json:"rating"`
		RatingState   string    `json:"rating_state"`
		IsModified    bool      `json:"is_modified"`
		LikesCount    int       `json:"likes_count"`
		CommentsCount int       `json:"comments_count"`
		CreatedAt     time.Time `json:"created_at"`
	} `json:"record,omitempty"`
}

type AnnictActivity struct {
	Activities []AnnictActivityBody `json:"activities"`
	TotalCount int                  `json:"total_count"`
	NextPage   int                  `json:"next_page"`
	PrevPage   any                  `json:"prev_page"`
}
