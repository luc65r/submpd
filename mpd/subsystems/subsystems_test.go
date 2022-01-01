package subsystems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	var (
		s   Subsystems
		err error
	)

	s, err = Parse("")
	assert.Nil(t, err)
	assert.Equal(t, Subsystems(0), s)

	s, err = Parse("  ")
	assert.Nil(t, err)
	assert.Equal(t, Subsystems(0), s)

	s, err = Parse("mixer")
	assert.Nil(t, err)
	assert.Equal(t, Subsystems(Mixer), s)

	s, err = Parse("unknown")
	assert.NotNil(t, err)

	s, err = Parse("stored_playlist playlist")
	assert.Nil(t, err)
	assert.Equal(t, Subsystems(StoredPlaylist|Playlist), s)

	s, err = Parse(" mount        options player  ")
	assert.Nil(t, err)
	assert.Equal(t, Subsystems(Mount|Options|Player), s)
}

func TestString(t *testing.T) {
	assert.Equal(t, "", Subsystems(0).String())
	assert.Equal(t, "update mixer neighbor", (Update | Mixer | Neighbor).String())
}
