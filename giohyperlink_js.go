// +build js

package giohyperlink

import (
	"gioui.org/io/event"
	"net/url"
	"syscall/js"
)

func listenEvents(_ event.Event) {
	// NO-OP
}

func open(u *url.URL) error {
	js.Global().Call("open", u.String(), "_blank", "noreferrer,noopener")
	return nil
}
