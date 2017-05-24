package core

import (
	"errors"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

// CommandCallback is a function or something
type CommandCallback func(dargs CommandArgs, args []string) (success bool)

// CommandArgs is arguments for the command provided by discordgo
type CommandArgs struct {
	Session *discordgo.Session
	Event   *discordgo.MessageCreate
	UsedTag string
}

func (c CommandArgs) SendMessage(text string) *discordgo.Message {
	msg, err := c.Session.ChannelMessageSend(c.Event.ChannelID, text)
	if err != nil {
		fmt.Println("Failed sending message! @"+c.Event.ChannelID+", err:", err)
		return &discordgo.Message{}
	}
	return msg
}

func (c CommandArgs) SendEmbed(embed *discordgo.MessageEmbed) *discordgo.Message {
	msg, err := c.Session.ChannelMessageSendEmbed(c.Event.ChannelID, embed)
	if err != nil {
		fmt.Println("Failed sending embed! @"+c.Event.ChannelID+", err:", err)
		return &discordgo.Message{}
	}
	return msg
}

func (c CommandArgs) GetGuild() *discordgo.Guild {
	ch, err := c.Session.Channel(c.Event.ChannelID)
	if err != nil {
		LogError(err.Error(), "Celestibot CMD")
		return &discordgo.Guild{}
	}
	guild, err := c.Session.Guild(ch.GuildID)
	if err != nil {
		LogError(err.Error(), "Celestibot CMD")
		return &discordgo.Guild{}
	}
	return guild
}

func (c CommandArgs) GetMe() *discordgo.User {
	user, err := c.Session.User("@me")
	if err != nil {
		return nil
	}
	return user
}

func applyEmoji(c *discordgo.Guild, text string) string {
	output := ""

	inEmoji := false
	current := ""
	for _, char := range text {
		if !inEmoji {
			if char == ':' {
				inEmoji = true
				output += string(char)
			} else {
				output += string(char)
			}
		} else {
			current += string(char)
			if char == ' ' || char == '\n' {
				output += current
				inEmoji = false
				current = ""
			} else if char == ':' {
				inEmoji = false
				output += getEmoji(c, current[:len(current)-1])
				current = ""
			}
		}
	}
	if current != "" {
		output += current
	}
	current = ""
	return output
}

func getEmoji(c *discordgo.Guild, name string) string {
	for _, emoji := range c.Emojis {
		if emoji.Name == name {
			return "<:" + emoji.Name + ":" + emoji.ID + ">"
		}
	}
	return ":" + name + ":"
}

func (c CommandArgs) SendMessageFormatted(text string) {
	ch, err := c.Session.Channel(c.Event.ChannelID)
	if err != nil {
		LogError(err.Error(), "Celestibot CMD")
		return
	}
	guild, err := c.Session.Guild(ch.GuildID)
	if err != nil {
		LogError(err.Error(), "Celestibot CMD")
		return
	}
	c.Session.ChannelMessageSend(c.Event.ChannelID, applyEmoji(guild, text))
}

func (c CommandArgs) SendFile(name, file string) error {
	fl, err := os.Open(file)
	if err != nil {
		return err
	}
	c.Session.ChannelFileSend(c.Event.ChannelID, name, fl)
	fl.Close()
	return nil
}

func (c CommandArgs) SendTyping() {
	c.Session.ChannelTyping(c.Event.ChannelID)
}

func (c CommandArgs) SetGame(game string) {
	c.Session.UpdateStatus(0, game)
}

func (c CommandArgs) GetUserMention(ID string) string {
	return ("<@" + ID + ">")
}

func (c CommandArgs) GetChannelMention(ID string) string {
	return ("<#" + ID + ">")
}

// Command is a command.
type Command struct {
	Callback      CommandCallback
	CommandTag    string
	AlternateTags []string
}

// Commands list.
var commands = make(map[string]Command)

func GetCommandsLength() int {
	return len(commands)
}

// AddCommand adds a command to Commands.
func AddCommand(key string, callback CommandCallback, tag string, alttags []string) {
	commands[key] = Command{callback, tag, alttags}
	LogInfo("Added command "+tag+"...", "Command Handler")
}

func hasTag(tag string, tags []string) bool {
	for _, tagb := range tags {
		if tag == tagb {
			return true
		}
	}
	return false
}

// GetCommandByTag gets an command by its tag.
func GetCommandByTag(tag string) (Command, error) {
	for _, cmd := range commands {
		if cmd.CommandTag == tag {
			return cmd, nil
		}
		if cmd.AlternateTags != nil {
			if hasTag(tag, cmd.AlternateTags) {
				return cmd, nil
			}
		}
	}
	return Command{}, errors.New("Command " + tag + " not found!")
}

//GetCommand gets an command via its name.
func GetCommand(key string) (Command, error) {
	if val, ok := commands[key]; ok {
		return val, nil
	}

	return Command{}, errors.New("Could not find command with key " + key)
}
