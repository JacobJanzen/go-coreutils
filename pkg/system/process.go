package system

import (
	/*
		#include <sys/sysinfo.h>
	*/
	"C"
	"fmt"
)

func checkProcCount(nproc int) (int, error) {
	if nproc < 1 {
		return nproc, fmt.Errorf("error getting processor count, invalid number %d", nproc)
	}

	return nproc, nil
}

func NProc() (int, error) {
	nproc := int(C.get_nprocs())

	return checkProcCount(nproc)
}

func NProcAll() (int, error) {
	nproc := int(C.get_nprocs_conf())

	return checkProcCount(nproc)
}
