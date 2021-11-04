// +build darwin,!ios

package giohyperlink

import (
	"gioui.org/io/event"
	"net/url"
	"os/exec"
)

func listenEvents(_ event.Event) {
	// NO-OP
}

func open(u *url.URL) error {
	return exec.Command("open", u.String()).Run()
}
