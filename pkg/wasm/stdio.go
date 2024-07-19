package wasm

import (
	"fmt"
	"os"
)

func WriteStdErr(code byte, message string) {
	body := make([]byte, len(message)+1)
	body[0] = code
	copy(body[1:], message)
	os.Stderr.Write(body)
}

func Logf(format string, args ...any) {
	WriteStdErr(0, fmt.Sprintf(format, args...))
}

func Exit(err error) {
	WriteStdErr(1, err.Error())
}
