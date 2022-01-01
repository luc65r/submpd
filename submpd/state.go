package submpd

import (
	"github.com/delucks/go-subsonic"
	"github.com/faiface/beep"
)

type State struct {
	Sub subsonic.Client
	Streamer beep.StreamCloser
	SampleRate beep.SampleRate
	Ctrl *beep.Ctrl
}
