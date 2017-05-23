package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/Member1221/plutobot-go/core"
	"github.com/Member1221/plutobot-go/cmds"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"time"
	"sync"
)

var dg *discordgo.Session
var err error

var WEBHOOK_REDDIT *discordgo.Webhook = nil
var ASYNC_END bool = false
var LAST_REDDIT map[string]string = make(map[string]string)
var ASYNC_MUTEX sync.Mutex

func main() {
	dg, err = discordgo.New("Bot " + Token("prod"))
	if err != nil {
		core.LogFatal("Discord could not connect, reason: "+err.Error(), "DISCORD_LOAD", 1)
		return
	}

	//Add message handler.
	dg.AddHandler(onMessageRecieve)
	dg.AddHandler(onMessageQue)
	err = dg.Open()
	if err != nil {
		core.LogFatal("Discord could not connect, reason: "+err.Error(), "DISCORD_WS_LOAD", 1)
		return
	}
	AddCommands()
	WEBHOOK_REDDIT, err = dg.Webhook("316553534667751427/zgXPRSoXG__mYv6MaokHlHC1fCTsAp_-xxllNbF4aQUjyRe1jok4-iloNY__z5X7a3tO")
	go func () {
		for !ASYNC_END {
			RedditUpdate("spacex")
			time.Sleep(time.Minute)
		}
	}()

	core.LogInfoG("PlutoBot Connected! Ctrl-C to exit.", "DISCORD_LOAD")
	// Wait till Ctrl + C is pressed, then close the connection and exit.
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGINT|syscall.SIGHUP)
	go func() {
		<-c
		onExit()
		os.Exit(0)
	}()
	<-make(chan struct{})
	onExit()
}

func onMessageQue(s *discordgo.Session, event *discordgo.MessageCreate) {

}

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

func onMessageRecieve(s *discordgo.Session, event *discordgo.MessageCreate) {

	go func(s *discordgo.Session, event *discordgo.MessageCreate) {
			if core.GetCommandsLength() > 0 {
				if strings.HasPrefix(event.Content, "$") {
					//s.ChannelMessageDelete(event.ChannelID, event.Message.ID)
					var substr = strings.Split(event.Content, " ")
					var dargs = substr[1:]
					var cmdtag = substr[0][1:]
					var cmd, err = core.GetCommandByTag(cmdtag)
					if err != nil {
						s.ChannelMessageSend(event.ChannelID, "`[`<@"+event.Author.ID+">`] Command not found!`")
						return
					}
					core.LogInfo("Command: "+cmdtag, event.Author.Username)
					cmd.Callback(core.CommandArgs{Session: s, Event: event, UsedTag: cmdtag}, dargs)

				}
			} else {
				core.LogError("No commands have been implemented.", "CommandHandler")
			}
	}(s, event)
}

func onExit() int {
	ASYNC_END = true
	dg.Logout()
	dg.Close()
	return 0
}

//AddCommands adds commands.
func AddCommands() {
	//core.AddCommand("clearcmd", cmds.ClearCommand, "thepurge", nil)
	core.AddCommand("reddit", cmds.RedditCommand, "reddit", nil)
}

func ExitCommand(a core.CommandArgs, v []string) bool {
	onExit()
	os.Exit(0)
	return true
}
