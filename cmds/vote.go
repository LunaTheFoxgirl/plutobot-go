package cmds

import (
	"math/rand"
	"strconv"
	"time"

	"strings"

	"fmt"

	"github.com/Member1221/plutobot-go/core"
	"github.com/bwmarrin/discordgo"
)

var VoteEmotes = strings.Split(`0⃣ 1⃣ 2⃣ 3⃣ 4⃣ 5⃣ 6⃣ 7⃣ 8⃣ 9⃣ 🇦 🇧 🇨 🇩 🇪 🇫 🇬 🇭 🇮 🇯 🇰 🇱 🇲 🇳 🇴 🇵 🇶 🇷 🇸 🇹 🇺 🇻 🇼 🇽 🇾 🇿`, " ")
var VEWords = strings.Split(`:zero: :one: :two: :three: :four: :five: :six: :seven: :eight: :nine: :regional_indicator_a: :regional_indicator_b: :regional_indicator_c: :regional_indicator_d: :regional_indicator_e: :regional_indicator_f: :regional_indicator_g: :regional_indicator_h: :regional_indicator_i: :regional_indicator_j: :regional_indicator_k: :regional_indicator_l: :regional_indicator_m: :regional_indicator_n: :regional_indicator_o: :regional_indicator_p: :regional_indicator_q: :regional_indicator_r: :regional_indicator_s: :regional_indicator_t: :regional_indicator_u: :regional_indicator_v: :regional_indicator_w: :regional_indicator_x: :regional_indicator_y: :regional_indicator_z:`, " ")

func pickEmotes(amount int) []string {
	have := make(map[string]bool)
	for i := 0; i < amount; i++ {
		for {
			rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
			emote := VoteEmotes[rand.Intn(len(VoteEmotes))]
			_, ok := have[emote]
			if !ok {
				have[emote] = true
				break
			}
		}
	}
	var ret []string
	for s := range have {
		ret = append(ret, s)
	}
	return ret
}

func VoteCommand(a core.CommandArgs, v []string) bool {
	question := v[0]
	answers := v[1:]
	var e = &discordgo.MessageEmbed{}
	e.Title = ":grey_question: *" + question + "*"
	var emojis = make(map[string]string) // emoji -> answer
	for i, s := range answers {
		e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
			Value:  VoteEmotes[i],
			Name:   s,
			Inline: true,
		})
		emojis[VoteEmotes[i]] = s
	}
	msg := a.SendEmbed(e)
	for em := range emojis {
		go func(em string) {
			err := a.Session.MessageReactionAdd(msg.ChannelID, msg.ID, em)
			if err != nil {
				fmt.Println("Error posting emoji "+em+":", err)
			}
		}(em)
	}
	go func(msg *discordgo.Message, s *discordgo.Session, emojis map[string]string, question string, a core.CommandArgs) {
		time.Sleep(10 * time.Second)
		var results = make(map[string]int)
		for em, answer := range emojis {
			var newem string = em
			/*for i, E := range VoteEmotes {
				if E == em {
					//newem = strings.Trim(VEWords[i], ":")
					newem = VEWords[i]
					break
				}
			}
			if newem == "" {
				fmt.Println("Emoji not found:", em)
			}*/
			allReactions, err := s.MessageReactions(msg.ChannelID, msg.ID, newem, 100)
			if err != nil {
				fmt.Println("Error in getting emojis:", err, newem)
			}
			results[answer] = len(allReactions) - 1
		}
		s.ChannelMessageDelete(msg.ChannelID, msg.ID)

		var fields []*discordgo.MessageEmbedField

		for emoji, amount := range results {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:   emoji,
				Value:  strconv.Itoa(amount),
				Inline: false,
			})
		}

		a.SendEmbed(&discordgo.MessageEmbed{
			Title:  question,
			Fields: fields,
		})
	}(msg, a.Session, emojis, question, a)
	return true
}
