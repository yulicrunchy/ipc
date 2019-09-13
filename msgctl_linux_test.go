package ipc_test

import (
	"bufio"
	"os"
	"strconv"
	"testing"

	"github.com/siadat/ipc"
)

func TestMsgmax(t *testing.T) {
	var msginfo ipc.Msginfo

	err := ipc.Msgctl(0, ipc.IPC_INFO, &msginfo)
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Open("/proc/sys/kernel/msgmax")
	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		systemval, err := strconv.Atoi(scanner.Text())
		if err != nil {
			t.Fatal(err)
		}

		if msginfo.Msgmax != systemval {
			t.Fatalf("msgmax is %d and should be %d", msginfo.Msgmax, systemval)
		}

		break
	}
}
