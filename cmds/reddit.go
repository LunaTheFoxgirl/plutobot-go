package cmds

import (
	"../core"
	"../rest"
)

type SubredditPosts struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				ContestMode           bool        `json:"contest_mode"`
				SubredditNamePrefixed string      `json:"subreddit_name_prefixed"`
				ThumbnailWidth        int         `json:"thumbnail_width"`
				Subreddit             string      `json:"subreddit"`
				SelftextHTML          string      `json:"selftext_html"`
				Selftext              string      `json:"selftext"`
				Likes                 int         `json:"likes,omitepmty"`
				SuggestedSort         string      `json:"suggested_sort"`
				ID                    string      `json:"id"`
				ViewCount             int         `json:"view_count,-"`
				Clicked               bool        `json:"clicked"`
				Author                string      `json:"author"`
				Saved                 bool        `json:"saved"`
				Name                  string      `json:"name"`
				Score                 int         `json:"score"`
				Over18                bool        `json:"over_18"`
				Domain                string      `json:"domain"`
				Hidden                bool        `json:"hidden"`
				Thumbnail             string      `json:"thumbnail"`
				SubredditID           string      `json:"subreddit_id"`
				Edited                interface{} `json:"edited"`
				AuthorFlairCSSClass   string      `json:"author_flair_css_class"`
				Gilded                int         `json:"gilded"`
				Downs                 int         `json:"downs"`
				BrandSafe             bool        `json:"brand_safe"`
				Archived              bool        `json:"archived"`
				CanGild               bool        `json:"can_gild"`
				HideScore             bool        `json:"hide_score"`
				Spoiler               bool        `json:"spoiler"`
				Permalink             string      `json:"permalink"`
				Locked                bool        `json:"locked"`
				Stickied              bool        `json:"stickied"`
				Created               float64     `json:"created"`
				URL                   string      `json:"url"`
				AuthorFlairText       string      `json:"author_flair_text"`
				Quarantine            bool        `json:"quarantine"`
				Title                 string      `json:"title"`
				CreatedUtc            float64     `json:"created_utc"`
				NumComments           int         `json:"num_comments"`
				IsSelf                bool        `json:"is_self"`
				Visited               bool        `json:"visited"`
				SubredditType         string      `json:"subreddit_type"`
				Ups                   int         `json:"ups"`
			} `json:"data"`
		} `json:"children"`
		After  string      `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

var Connection = prest.New()

func RedditCommand(a core.CommandArgs, v []string) bool {
	return true
}
