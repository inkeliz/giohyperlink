// +build windows

package giohyperlink

import (
	"gioui.org/io/event"
	"golang.org/x/sys/windows"
	"net/url"
)

func listenEvents(_ event.Event) {
	// NO-OP
}

func open(u *url.URL) error {
	return windows.ShellExecute(0, nil, windows.StringToUTF16Ptr(u.String()), nil, nil, windows.SW_SHOWNORMAL)
}
