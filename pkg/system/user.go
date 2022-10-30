package system

import (
	/*
		#include <pwd.h>
		#include <sys/types.h>
		#include <unistd.h>
	*/
	"C"
	"fmt"
)

func GetUsername(uid uint) (string, error) {
	pw := C.getpwuid(C.uint(uid))

	if pw == nil || pw.pw_name == nil {
		return "", fmt.Errorf("error getting user with uid %d", uid)
	}

	return C.GoString(pw.pw_name), nil
}

func GetLogin() (string, error) {
	login := C.getlogin()

	if login == nil {
		return "", fmt.Errorf("error getting userid")
	}

	return C.GoString(login), nil
}
