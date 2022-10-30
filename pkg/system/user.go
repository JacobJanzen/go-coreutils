package system

import (
	/*
		#include <pwd.h>
		#include <sys/types.h>
	*/
	"C"
)
import "fmt"

func GetUsername(uid uint) (string, error) {
	pw := C.getpwuid(C.uint(uid))

	if pw == nil || pw.pw_name == nil {
		return "", fmt.Errorf("error getting user with uid %d\n", uid)
	}

	return C.GoString(pw.pw_name), nil
}
