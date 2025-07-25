// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (mips64 || mips64le)

package linux

const (
	SYS_CLOSE         = 5003
	SYS_MPROTECT      = 5010
	SYS_FCNTL         = 5070
	SYS_PRCTL         = 5153
	SYS_EPOLL_CTL     = 5208
	SYS_EPOLL_PWAIT   = 5272
	SYS_EPOLL_CREATE1 = 5285
	SYS_EPOLL_PWAIT2  = 5441
	SYS_EVENTFD2      = 5284
	SYS_OPENAT        = 5247
	SYS_PREAD64       = 5016
	SYS_READ          = 5000

	EFD_NONBLOCK = 0x80

	O_LARGEFILE = 0x0
)

type EpollEvent struct {
	Events    uint32
	pad_cgo_0 [4]byte
	Data      [8]byte // unaligned uintptr
}
