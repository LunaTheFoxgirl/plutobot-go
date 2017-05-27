package cmds

import (
	"github.com/Member1221/plutobot-go/core"
	"github.com/Member1221/plutobot-go/rest"
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

// TODO: Reimplement this.
/*
func RedditUpdate(subreddit string) bool {
	req, err := http.NewRequest("GET", "https://www.reddit.com/r/"+subreddit+"/new.json", nil)
	if err != nil {
		fmt.Println(err)
		return false
	}
	req.Header.Set("User-Agent", "UGx1dG9Cb3Q=")

	// Handle the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return false
	}

	str, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	var posts cmds.SubredditPosts

	err = json.Unmarshal(str, &posts)
	if err != nil {

		fmt.Println(err.Error() + "\n: " + string(str))
		return false
	}
	ASYNC_MUTEX.Lock()
	if LAST_REDDIT[subreddit] == posts.Data.Children[0].Data.ID {
		return true
	}
	ASYNC_MUTEX.Unlock()

	var thumbnail *discordgo.MessageEmbedThumbnail = nil

	if posts.Data.Children[0].Data.Thumbnail != "self" {
		thumbnail = &discordgo.MessageEmbedThumbnail{posts.Data.Children[0].Data.Thumbnail, posts.Data.Children[0].Data.Thumbnail, 128, 128}
	}
	text := posts.Data.Children[0].Data.Selftext
	if len(text) > 250 {
		text = text[:247] + "..."
	}
	embed := discordgo.MessageEmbed{
		"https://reddit.com" + posts.Data.Children[0].Data.Permalink,
		"rich",
		posts.Data.Children[0].Data.Title,
		text,
		"",
		0xFF00FF,
		nil,
		nil,
		thumbnail,
		nil,
		nil,
		nil,
		[]*discordgo.MessageEmbedField{},
	}
	params := discordgo.WebhookParams{"", "PlutoBot->Reddit", "", false, "", []*discordgo.MessageEmbed{&embed}}

	err = dg.WebhookExecute(WEBHOOK_REDDIT.ID, WEBHOOK_REDDIT.Token, true, &params)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	ASYNC_MUTEX.Lock()
	LAST_REDDIT[subreddit] = posts.Data.Children[0].Data.ID
	ASYNC_MUTEX.Unlock()
	return true
}
*/
