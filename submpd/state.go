package submpd

import (
	"sync"

	"github.com/delucks/go-subsonic"
	"github.com/faiface/beep"
)

type State struct {
	Mu           sync.RWMutex
	Sub          subsonic.Client
	Streamer     beep.StreamCloser
	SampleRate   beep.SampleRate
	Ctrl         *beep.Ctrl
	CurrentQueue Queue
}
