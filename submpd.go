package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/delucks/go-subsonic"
	"github.com/faiface/beep/speaker"
	"github.com/luc65r/submpd/l"
	"github.com/luc65r/submpd/submpd"
)

func main() {
	var err error

	port := flag.Int("port", 6600, "")
	subsonicUrl := flag.String("url", "https://demo.subsonic.org", "")
	username := flag.String("user", "", "")
	password := flag.String("password", "", "")
	flag.Parse()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		l.Fatal("Listen failed: ", err)
	}
	l.Info("server started")

	sub := subsonic.Client{
		Client: http.DefaultClient,
		BaseUrl: *subsonicUrl,
		User: *username,
		ClientName: "submpd",
	}

	err = sub.Authenticate(*password)
	if err != nil {
		l.Fatal(err)
	}
	l.Infof("authenticated to `%s` as `%s`", *subsonicUrl, *username)

	state := submpd.State{
		Sub: sub,
		SampleRate: 44100,
	}

	speaker.Init(state.SampleRate, state.SampleRate.N(time.Second / 30))

	for {
		c, err := ln.Accept()
		if err != nil {
			l.Fatal("Accept failed: ", err)
		}

		client := submpd.Client{
			State: &state,
			Conn: c,
		}

		go client.Handle()
	}
}
