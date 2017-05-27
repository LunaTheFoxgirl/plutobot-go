package cmds

import (
	"fmt"
	"github.com/Member1221/plutobot-go/core"
	"time"
)

func ClearCommand(a core.CommandArgs, v []string) bool {
	loop := true
	msg := a.SendMessage("Please wait while the channel is purged of messages...")

	for loop {
		ar, err := a.Session.ChannelMessages(a.Event.ChannelID, 99, msg.ID, "", "")
		if len(ar) < 1 {
			loop = false
			return true
		}
		if err != nil {
			fmt.Println(err)
			return false
		}
		var ids []string
		for _, msg := range ar {
			ids = append(ids, msg.ID)
		}

		a.Session.ChannelMessagesBulkDelete(a.Event.ChannelID, ids)
		time.Sleep(time.Millisecond * 500)
	}
	a.Session.ChannelMessageDelete(a.Event.ChannelID, msg.ID)
	return true
}
