package submpd

import "github.com/luc65r/submpd/mpd"

var Commands = map[string]func(*Client, []string) mpd.Response{
	"ping": func(c *Client, args []string) mpd.Response {
		return mpd.NormalResponse{}
	},

	"status": func(c *Client, args []string) mpd.Response {
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

	"plchanges": func(c *Client, args []string) mpd.Response {
		return mpd.NormalResponse{}
	},

	"outputs": func(c *Client, args []string) mpd.Response {
		return mpd.NormalResponse{
			Data: map[string]string{
				"outputid":      "0",
				"outputname":    "default",
				"outputenabled": "1",
			},
		}
	},

	"decoders": func(c *Client, args []string) mpd.Response {
		return mpd.NormalResponse{}
	},

	"command_list_begin": func(c *Client, args []string) mpd.Response {
		c.inCommandList = true
		c.cnb = -1
		return nil
	},

	"command_list_end": func(c *Client, args []string) mpd.Response {
		c.inCommandList = false
		return nil
	},

	"idle": func(c *Client, args []string) mpd.Response {
		c.idle = true
		return nil
	},

	"noidle": func(c *Client, args []string) mpd.Response {
		if !c.idle {
			return mpd.NormalResponse{}
		}

		rp := mpd.NormalResponse{
			Data: map[string]string{
				"changed": c.subs.String(),
			},
		}
		c.subs = 0
		c.idle = false
		return rp
	},
}
