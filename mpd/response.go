package mpd

import "fmt"

type Response interface {
	Format() []byte
}

type NormalResponse struct {
	Data map[string]string
}

func (r NormalResponse) Format() (bs []byte) {
	for k, v := range r.Data {
		bs = append(bs, (k + ": " + v + "\n")...)
	}
	bs = append(bs, "OK\n"...)
	return
}

type BinaryResponse struct {
	Data map[string]string
	BinaryData []byte
}

func (r BinaryResponse) Format() (bs []byte) {
	for k, v := range r.Data {
		bs = append(bs, (k + ": " + v + "\n")...)
	}
	bs = append(bs, fmt.Sprintf("binary: %d\n", len(r.BinaryData))...)
	bs = append(bs, r.BinaryData...)
	bs = append(bs, "\nOK\n"...)
	return
}

type Ack int

const (
	AckErrorNotList Ack = 1
	AckErrorArg Ack = 2
	AckErrorPassword Ack = 3
	AckErrorPermission Ack = 4
	AckErrorUnknown Ack = 5

	AckErrorNoExist = 50
	AckErrorPlaylistMax = 51
	AckErrorSystem = 52
	AckErrorPlaylistLoad = 53
	AckErrorUpdateAlready = 54
	AckErrorPlayerSync = 55
	AckErrorExist = 56
)

type FailureResponse struct {
	Error Ack
	CommandNb int
	Command string
	Message string
}

func (r FailureResponse) Format() []byte {
	return []byte(fmt.Sprintf(
		"ACK [%d@%d] {%s} %s\n",
		r.Error, r.CommandNb, r.Command, r.Message,
	))
}
