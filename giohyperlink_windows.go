// +build windows

package giohyperlink

import (
	"gioui.org/io/event"
	"net/url"
	"os/exec"
	"strings"
	"syscall"
)

func listenEvents(_ event.Event) {
	// NO-OP
}

func open(u *url.URL) error {
	cmd := exec.Command("cmd", `/C`, "start", strings.NewReplacer("&", "^&").Replace(u.String()))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return cmd.Run()
}
