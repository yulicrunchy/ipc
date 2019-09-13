package ipc

import (
	"syscall"
	"unsafe"

	/*#include <sys/types.h>
	#include <sys/ipc.h>
	#include <sys/msg.h>
	*/
	"C"
)

// Msgctl calls the msgctl() syscall.
// Sets the msgmax value of msginfo on Linux.
func Msgctl(qid uint64, cmd int, msginfo *Msginfo) error {
	var buf2 C.struct_msginfo

	_, _, err := syscall.Syscall(syscall.SYS_MSGCTL, uintptr(qid), uintptr(cmd), uintptr(unsafe.Pointer(&buf2)))
	if err != 0 {
		return err
	}

	if msginfo != nil {
		msginfo.Msgmax = int(buf2.msgmax)
	}

	return nil
}
