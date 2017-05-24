package core

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

type ticket struct {
	Parent *CommandArgs

	// hooks here
	Emoji_hook *emojihook

	// vendor info here
	Attachedvendor *Vendor
	vendorID       string
}

type emojihook struct {
	Add chan *discordgo.MessageReactionAdd
	Del chan *discordgo.MessageReactionRemove
}

type Vendor struct {
	T map[string]*ticket
	sync.Mutex
} // ID -> ticket associated

func (v *Vendor) Handle(s *discordgo.Session, a interface{}) { // as this is an interface, discord will send *all* events through here
	switch a.(type) {
	case *discordgo.MessageReactionAdd:
		Ra, _ := a.(*discordgo.MessageReactionAdd)
		v.exec_emoji_add(Ra)
	case *discordgo.MessageReactionRemove:
		Rr, _ := a.(*discordgo.MessageReactionRemove)
		v.exec_emoji_del(Rr)
	}
}

func (v *Vendor) exec_emoji_add(m *discordgo.MessageReactionAdd) {
	t, ok := v.T[m.MessageID]
	if ok {
		if t.Emoji_hook != nil {
			t.Emoji_hook.Add <- m
		}
	}
}

func (v *Vendor) exec_emoji_del(m *discordgo.MessageReactionRemove) {
	t, ok := v.T[m.MessageID]
	if ok {
		if t.Emoji_hook != nil { // check after, else we get dereference err
			t.Emoji_hook.Del <- m
		}
	}
}

func (v *Vendor) Request(t *ticket, ID string) {
	_, ok := v.T[ID]
	if !ok { // since we'll only add one ticket per message ID
		v.Lock()
		defer v.Unlock()
		v.T[ID] = t
		v.T[ID].Attachedvendor = v
		v.T[ID].vendorID = ID
	}
}

func (t *ticket) Done() {
	v := t.Attachedvendor
	v.Lock()
	defer v.Unlock()
	delete(v.T, t.vendorID)
}
