package ipc

import (
	"fmt"
	"os"
)

var (
	BufSize = 8192
)

func init() {
	var msginfo Msginfo

	err := Msgctl(0, IPC_INFO, &msginfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get msginfo: (%v)\n", err)
		os.Exit(1)
	}

	BufSize = msginfo.Msgmax
}
