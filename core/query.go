package core

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

type ticket struct {
	Parent *CommandArgs

	// hooks here
	Data_hook chan DataRecieve

	// vendor info here
	Attachedvendor *Vendor
	ID uint
}


type DataRecieve struct {
	EventType int // from const
	A interface{}
}

const (
	TERM = -1
	emojiAdd = iota
	emojiDel
	messAdd
	messDel
	messEdit
)

type Vendor struct {
	T map[uint]*ticket
	sync.Mutex
	latestID uint
} // ID -> ticket associated

func (v *Vendor) Handle(s *discordgo.Session, a interface{}) { // as this is an interface, discord will send *all* events through here
	switch a.(type) {
	case *discordgo.MessageReactionAdd:
		v.spread(DataRecieve{emojiAdd, a})
	case *discordgo.MessageReactionRemove:
		v.spread(DataRecieve{emojiDel, a})
	case *discordgo.MessageCreate:
		v.spread(DataRecieve{messAdd, a})
	case *discordgo.MessageDelete:
		v.spread(DataRecieve{messDel, a})
	case *discordgo.MessageEdit:
		v.spread(DataRecieve{messEdit, a})
	}
}

func (v *Vendor) TERMINATE() {
	v.spread(DataRecieve{TERM, nil})
}

func (v *Vendor) spread(d DataRecieve) {
	for _, t := range v.T {
		go func() {t.Data_hook <- d}()
	}
}

func (v *Vendor) Request(t *ticket) {

		v.Lock()
		defer v.Unlock()
	v.latestID+=1
		v.T[v.latestID] = t
		v.T[v.latestID].Attachedvendor = v
	v.T[v.latestID].ID = v.latestID
}

func (t *ticket) Done() {
	v := t.Attachedvendor
	v.Lock()
	defer v.Unlock()
	delete(v.T, t.ID)
}
