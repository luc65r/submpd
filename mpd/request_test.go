package mpd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequests(t *testing.T) {
	var (
		r   Request
		err error
	)

	r, err = ParseRequest([]byte(""))
	assert.NotNil(t, err, "empty request without a newline should fail")

	r, err = ParseRequest([]byte("\n"))
	assert.NotNil(t, err, "empty request should fail")

	r, err = ParseRequest([]byte("ping\n"))
	assert.Nil(t, err)
	assert.Equal(t, Request{"ping", nil}, r)

	r, err = ParseRequest([]byte(`find "(Artist == \"foo\\'bar\\\"\")"` + "\n"))
	assert.Nil(t, err)
	assert.Equal(t, Request{"find", []string{`(Artist == "foo\'bar\"")`}}, r)
}
