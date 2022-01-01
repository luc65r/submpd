package subsystems

import (
	"fmt"
	"strings"
)

type Subsystems int16

const (
	Database Subsystems = 1 << iota
	Update
	StoredPlaylist
	Playlist
	Player
	Mixer
	Output
	Options
	Partition
	Sticker
	Subscription
	Message
	Neighbor
	Mount
)

var FromString = map[string]Subsystems{
	"database":        Database,
	"update":          Update,
	"stored_playlist": StoredPlaylist,
	"playlist":        Playlist,
	"player":          Player,
	"mixer":           Mixer,
	"output":          Output,
	"options":         Options,
	"partition":       Partition,
	"sticker":         Sticker,
	"subscription":    Subscription,
	"message":         Message,
	"neighbor":        Neighbor,
	"mount":           Mount,
}

var ToString = map[Subsystems]string{
	Database:       "database",
	Update:         "update",
	StoredPlaylist: "stored_playlist",
	Playlist:       "playlist",
	Player:         "player",
	Mixer:          "mixer",
	Output:         "output",
	Options:        "options",
	Partition:      "partition",
	Sticker:        "sticker",
	Subscription:   "subscription",
	Message:        "message",
	Neighbor:       "neighbor",
	Mount:          "mount",
}

var All = []Subsystems{
	Database, Update, StoredPlaylist, Playlist,
	Player, Mixer, Output, Options, Partition,
	Sticker, Subscription, Message, Neighbor,
	Mount,
}

func Parse(str string) (sub Subsystems, err error) {
	for _, s := range strings.Fields(str) {
		a, ok := FromString[s]
		if !ok {
			return sub, fmt.Errorf("unknown subsystem: `%s`", s)
		}
		sub |= a
	}
	return
}

func (sub Subsystems) String() (res string) {
	for i, s := range All {
		if sub&(1<<i) != 0 {
			res += ToString[s] + " "
		}
	}
	if len(res) > 0 {
		res = res[:len(res)-1]
	}
	return
}
