package submpd

import "github.com/luc65r/submpd/mpd"

var Commands = map[string]func(*State, []string) mpd.Response{
	"ping": func(s *State, args []string) mpd.Response {
		return mpd.NormalResponse{}
	},

	"status": func(s *State, args []string) mpd.Response {
		return mpd.NormalResponse{
			Data: map[string]string{
				"repeat":         "0",
				"random":         "0",
				"single":         "0",
				"consume":        "0",
				"partition":      "default",
				"playlist":       "0",
				"playlistlength": "0",
				"mixrampdb":      "0",
				"state":          "stop",
			},
		}
	},

	"plchanges": func(s *State, args []string) mpd.Response {
		return mpd.NormalResponse{}
	},

	"outputs": func(s *State, args []string) mpd.Response {
		return mpd.NormalResponse{
			Data: map[string]string{
				"outputid": "0",
				"outputname": "default",
				"outputenabled": "1",
			},
		}
	},

	"decoders": func(s *State, args []string) mpd.Response {
		return mpd.NormalResponse{}
	},
}
