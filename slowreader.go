package slowloris

import (
	"fmt"
	"io"
	"time"
)

type SlowReader struct {
	Content string
	index   int
	// SleepTime in milliseconds
	SleepTime time.Duration
	//controls printing content part to stdout
	Log bool
}

func (sr *SlowReader) Read(bytes []byte) (read int, err error) {
	if sr.index == len(sr.Content) {
		return 0, io.EOF
	}

	bytes[0] = byte(sr.Content[sr.index])
	if sr.Log {
		fmt.Printf("%s", string(sr.Content[sr.index]))
	}
	time.Sleep(sr.SleepTime)
	sr.index++
	return 1, nil
}
