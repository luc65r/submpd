package l

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

const (
	debugLevel int = iota
	infoLevel
	warningLevel
	errorLevel
	fatalLevel
)

var levelNames = []string{
	debugLevel:   "DEBUG",
	infoLevel:    "INFO",
	warningLevel: "WARNING",
	errorLevel:   "ERROR",
	fatalLevel:   "FATAL",
}

var levelColors = []string{
	debugLevel:   "",
	infoLevel:    "",
	warningLevel: "33",
	errorLevel:   "31",
	fatalLevel:   "31",
}

var Level int = debugLevel

var (
	out = os.Stderr
	mu  sync.Mutex
	buf []byte
)

func lprint(ll int, s string) {
	if ll >= Level {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		} else {
			slash := strings.LastIndex(file, "/")
			if slash >= 0 {
				file = file[slash+1:]
			}
		}

		mu.Lock()
		defer mu.Unlock()

		buf = buf[:0]
		buf = append(buf, "\033["...)
		buf = append(buf, levelColors[ll]...)
		buf = append(buf, 'm')
		buf = append(buf, levelNames[ll]...)
		buf = append(buf, ": "...)
		buf = append(buf, file...)
		buf = append(buf, ':')
		buf = append(buf, strconv.Itoa(line)...)
		buf = append(buf, ": "...)
		buf = append(buf, s...)
		if s[len(s)-1] != '\n' {
			buf = append(buf, '\n')
		}
		buf = append(buf, "\033[m"...)
		out.Write(buf)
	}
}

func Debug(a ...interface{}) {
	lprint(debugLevel, fmt.Sprint(a...))
}

func Debugf(f string, a ...interface{}) {
	lprint(debugLevel, fmt.Sprintf(f, a...))
}

func Info(a ...interface{}) {
	lprint(infoLevel, fmt.Sprint(a...))
}

func Infof(f string, a ...interface{}) {
	lprint(infoLevel, fmt.Sprintf(f, a...))
}

func Warning(a ...interface{}) {
	lprint(warningLevel, fmt.Sprint(a...))
}

func Warningf(f string, a ...interface{}) {
	lprint(warningLevel, fmt.Sprintf(f, a...))
}

func Error(a ...interface{}) {
	lprint(errorLevel, fmt.Sprint(a...))
}

func Errorf(f string, a ...interface{}) {
	lprint(errorLevel, fmt.Sprintf(f, a...))
}

func Fatal(a ...interface{}) {
	lprint(fatalLevel, fmt.Sprint(a...))
	os.Exit(1)
}

func Fatalf(f string, a ...interface{}) {
	lprint(fatalLevel, fmt.Sprintf(f, a...))
	os.Exit(1)
}
