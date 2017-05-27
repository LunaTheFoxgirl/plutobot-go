package main

import (
	"github.com/Member1221/plutobot-go/cmds"
	"github.com/Member1221/plutobot-go/core"
	"github.com/Member1221/plutobot-go/db"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var dg *discordgo.Session
var err error

var V = &core.Vendor{}
var DB db.PlutoDB

func main() {
	DB, err = db.Open("plutobot")
	if len(os.Args[1:]) > 0 {
		args := os.Args[1:]
		if args[0] == "--registertoken" {
			err := RegisterToken(DB, args[1])
			if err != nil {
				core.LogFatal("Could not write token to databse, reason: "+err.Error(), "DATABASE_TOKEN_SET", 1)
			}
		}
	}
	if err != nil {
		core.LogFatal("Could not open database \"plutobot\", reason: "+err.Error(), "DATABASE_LOAD", 2)
		return
	}

	token, err := Token(DB)
	if err != nil {
		core.LogFatal("Token could not be loaded, reason: "+err.Error(), "DATABASE_TOKEN_GET", 4)
		return
	}
	dg, err = discordgo.New("Bot " + token)
	if err != nil {
		core.LogFatal("Discord could not connect, reason: "+err.Error(), "DISCORD_LOAD", 5)
		return
	}
	// Make sure to clear the token from memory.
	token = ""

	// Add message handler.
	dg.AddHandler(onMessageRecieve)
	dg.AddHandler(onMessageQue)
	dg.AddHandler(V.Handle)

	err = dg.Open()

	if err != nil {
		core.LogFatal("Discord could not connect, reason: "+err.Error(), "DISCORD_WS_LOAD", 6)
		return
	}
	AddCommands()

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
	DB.AddMessage(event.Message)
}

func onMessageRecieve(s *discordgo.Session, event *discordgo.MessageCreate) {

	go func(s *discordgo.Session, event *discordgo.MessageCreate) {
		if core.GetCommandsLength() > 0 {
			if strings.HasPrefix(event.Content, "$") {
				//s.ChannelMessageDelete(event.ChannelID, event.Message.ID)
				var substr = strings.Split(event.Content, " ")
				var dargs = core.SanitizeArgs(substr[1:])
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

	DB.Close()

	dg.Logout()
	dg.Close()
	return 0
}

//AddCommands adds commands.
func AddCommands() {
	//core.AddCommand("clearcmd", cmds.ClearCommand, "thepurge", nil)
	core.AddCommand("reddit", cmds.RedditCommand, "reddit", nil)
	core.AddCommand("vote", cmds.VoteCommand, "vote", nil)
}

func ExitCommand(a core.CommandArgs, v []string) bool {
	onExit()
	os.Exit(0)
	return true
}
