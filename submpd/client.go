package submpd

import (
	"bufio"
	"io"
	"net"

	"github.com/luc65r/submpd/l"
	"github.com/luc65r/submpd/mpd"
	"github.com/luc65r/submpd/mpd/subsystems"
)

type Client struct {
	State   *State
	Conn    net.Conn
	subChan <-chan subsystems.Subsystems
	subs    subsystems.Subsystems
	idle    bool
}

func (c Client) Handle() {
	l.Infof("new client: %v", c.Conn.RemoteAddr())
	c.Conn.Write([]byte("OK MPD 0.24\n"))

	data := make(chan []byte)
	go read(c.Conn, data)

	for {
		select {
		case msg, ok := <-data:
			if !ok {
				return
			}
			handleRequest(&c, msg)

		case sub := <-c.subChan:
			c.subs |= sub
			if c.idle {
				c.Conn.Write(Commands["noidle"](&c, []string{}).Format())
			}
		}
	}
}

func read(conn net.Conn, data chan<- []byte) {
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				l.Errorf("%v: %v", conn.RemoteAddr(), err)
			}
			conn.Close()
			close(data)
			break
		}
		data <- msg
	}
}

func handleRequest(c *Client, msg []byte) {
	rq, err := mpd.ParseRequest(msg)
	if err != nil {
		l.Warningf("%v: failed to parse request: %s", c.Conn.RemoteAddr(), msg)
		rp := mpd.FailureResponse{
			Error:   mpd.AckErrorUnknown,
			Command: "",
			Message: "failed to parse request",
		}
		c.Conn.Write(rp.Format())
	}
	l.Debugf("%v: %s", c.Conn.RemoteAddr(), msg)

	var rp mpd.Response
	if c.idle && rq.Command != "noidle" {
		l.Warningf("%v: expected noidle, got: %s", c.Conn.RemoteAddr(), rq.Command)
		rp = mpd.FailureResponse{
			Error:   mpd.AckErrorUnknown,
			Command: rq.Command,
			Message: "expected noidle",
		}
	} else if f, ok := Commands[rq.Command]; ok {
		l.Debugf("%v: handling command: %s", c.Conn.RemoteAddr(), rq.Command)
		rp = f(c, rq.Args)
	} else {
		l.Warningf("%v: unknown command: %s", c.Conn.RemoteAddr(), rq.Command)
		rp = mpd.FailureResponse{
			Error:   mpd.AckErrorUnknown,
			Command: rq.Command,
			Message: "unknown command",
		}
	}
	l.Debugf("%v: response: %v", c.Conn.RemoteAddr(), rp)
	if rp != nil {
		c.Conn.Write(rp.Format())
	}
}
