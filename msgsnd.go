package ipc

import (
	"fmt"
	"syscall"
	"unsafe"
)

// Msgsnd calls the msgsnd() syscall.
func Msgsnd(qid uint64, msg *Msgbuf, flags uint64) error {
	if len(msg.Mtext) > BufSize {
		return fmt.Errorf("mtext is too large, %d > %d", len(msg.Mtext), BufSize)
	}
	qbuf := msgbufInternal{
		Mtype: msg.Mtype,
	}
	copy(qbuf.Mtext[:], msg.Mtext)

	_, _, err := syscall.Syscall6(syscall.SYS_MSGSND,
		uintptr(qid),
		uintptr(unsafe.Pointer(&qbuf)),
		uintptr(len(msg.Mtext)),
		uintptr(flags),
		0,
		0,
	)
	if err != 0 {
		return err
	}
	return nil
}
