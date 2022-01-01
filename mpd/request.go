package mpd

import (
	"fmt"
)

type Request struct {
	Command string
	Args    []string
}

func ParseRequest(bs []byte) (rq Request, err error) {
	if len(bs) == 0 {
		return rq, fmt.Errorf("empty request")
	}
	if bs[len(bs)-1] != '\n' {
		return rq, fmt.Errorf("request must end by a newline character")
	}

	i := 0
	for isAlpha(bs[i]) {
		i++
	}
	if i == 0 {
		return rq, fmt.Errorf("empty command")
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
				return nil, fmt.Errorf("character %c cannot be escaped", bs[*i])
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
