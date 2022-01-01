package submpd

import (
	"bufio"
	"io"
	"net"

	"github.com/luc65r/submpd/l"
	"github.com/luc65r/submpd/mpd"
)

type Client struct {
	State *State
	Conn  net.Conn
}

func (c Client) Handle() {
	cnb := 0
	inCommandList := false

	l.Infof("new client: %v", c.Conn.RemoteAddr())

	r := bufio.NewReader(c.Conn)
	c.Conn.Write([]byte("OK MPD 0.24\n"))

	for {
		msg, err := r.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				l.Errorf("%v: %v", c.Conn.RemoteAddr(), err)
			}
			c.Conn.Close()
			return
		}

		rq, err := mpd.ParseRequest(msg)
		if err != nil {
			l.Errorf("%v: failed to parse request: %s", c.Conn.RemoteAddr(), msg)
			rp := mpd.FailureResponse{
				Error:     mpd.AckErrorUnknown,
				CommandNb: cnb,
				Command:   "",
				Message:   "failed to parse request",
			}
			c.Conn.Write(rp.Format())
		}
		l.Debugf("%v: %s", c.Conn.RemoteAddr(), msg)

		switch rq.Command {
		case "command_list_begin":
			inCommandList = true
			cnb = -1
		case "command_list_end":
			if !inCommandList {
				// TODO
			}
			inCommandList = false
		default:
			var rp mpd.Response
			if f, ok := Commands[rq.Command]; ok {
				rp = f(c.State, rq.Args)
				if r, ok := rp.(mpd.FailureResponse); ok {
					r.CommandNb = cnb
				}
			} else {
				rp = mpd.FailureResponse{
					Error:     mpd.AckErrorUnknown,
					CommandNb: cnb,
					Command:   rq.Command,
					Message:   "unknown command",
				}
			}
			c.Conn.Write(rp.Format())
		}

		if inCommandList {
			cnb++
		} else {
			cnb = 0
		}
	}
}
