package mpd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequests(t *testing.T) {
	var (
		rs  []Request
		err error
	)

	rs, err = ParseRequests("")
	assert.NotNil(t, err, "empty request without a newline should fail")

	rs, err = ParseRequests("\n")
	assert.NotNil(t, err, "empty request should fail")

	rs, err = ParseRequests("command_list_begin\n")
	assert.NotNil(t, err, "`command_list_begin` as the only command should fail")

	rs, err = ParseRequests("command_list_end\n")
	assert.NotNil(t, err, "`command_list_end` as the only command should fail")

	rs, err = ParseRequests("ping\n")
	assert.Nil(t, err)
	assert.Equal(t, []Request{{"ping", nil}}, rs)

	rs, err = ParseRequests(`find "(Artist == \"foo\\'bar\\\"\")"` + "\n")
	assert.Nil(t, err)
	assert.Equal(t, []Request{{"find", []string{`(Artist == "foo\'bar\"")`}}}, rs)

	rs, err = ParseRequests(`command_list_begin
volume 86
play 10240
status
command_list_end
`)
	assert.Nil(t, err)
	assert.Equal(t, []Request{
		{"volume", []string{"86"}},
		{"play", []string{"10240"}},
		{"status", nil},
	}, rs)
}
