package mpd

import (
	"errors"
	"fmt"
	"strings"
)

type Request struct {
	Command string
	Args    []string
}

func ParseRequests(s string) ([]Request, error) {
	var err error
	if len(s) == 0 {
		return nil, errors.New("request must not be empty")
	}
	if s[len(s) - 1] != '\n' {
		return nil, errors.New("requests must end by a newline character")
	}
	ss := strings.SplitAfter(s, "\n")
	// We don't want the empty string
	ss = ss[:len(ss) - 1]
	rs := make([]Request, len(ss))

	for i, s := range ss {
		rs[i], err = parseRequest(s)
		if err != nil {
			return nil, err
		}
	}

	if len(rs) == 0 {
		return nil, errors.New("there must be at least one request")
	} else if len(rs) == 1 {
		if rs[0].Command == "command_list_begin" || rs[0].Command == "command_list_end" {
			return nil, errors.New(fmt.Sprintf("`%s` as the only command", rs[0].Command))
		} else {
			return rs, nil
		}
	} else {
		if rs[0].Command != "command_list_begin" {
			return nil, errors.New("command list does not begin with `command_list_begin`")
		} else if rs[len(rs)-1].Command != "command_list_end" {
			return nil, errors.New("command list does not end with `command_list_end`")
		} else {
			return rs[1 : len(rs)-1], nil
		}
	}
}

func parseRequest(s string) (rq Request, err error) {
	bs := []byte(s)

	i := 0
	for isAlpha(bs[i]) {
		i++
	}
	if i == 0 {
		return rq, errors.New("empty command")
	}
	rq.Command = string(bs[:i])

	for {
		switch bs[i] {
		case ' ', '\t':
			i++
		case '\n':
			return
		case '"':
			i++
			arg, err := parseQuotedArg(bs, &i)
			if err != nil {
				return rq, err
			}
			rq.Args = append(rq.Args, string(arg))
		default:
			start := i
			parseArg(bs, &i)
			rq.Args = append(rq.Args, string(bs[start:i]))
		}
	}
}

func parseQuotedArg(bs []byte, i *int) ([]byte, error) {
	res := make([]byte, 0, len(bs))
	escaped := false

	for {
		switch bs[*i] {
		case '"':
			if escaped {
				res = append(res, '"')
				escaped = false
			} else {
				(*i)++
				return res, nil
			}
		case '\\':
			if escaped {
				res = append(res, '\\')
				escaped = false
			} else {
				escaped = true
			}
		default:
			if escaped {
				return nil, errors.New(fmt.Sprintf("character %c cannot be escaped", bs[*i]))
			} else {
				res = append(res, bs[*i])
			}
		}
		(*i)++
	}
}

func parseArg(bs []byte, i *int) {
	for {
		switch bs[*i] {
		case ' ', '\t', '\n':
			return
		default:
			(*i)++
		}
	}
}

func isAlpha(c byte) bool {
	upper := c & 0b1101_1111
	return 'A' <= upper && upper <= 'Z' || c == '_'
}
