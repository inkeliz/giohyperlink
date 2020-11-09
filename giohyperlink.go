package giohyperlink

import (
	"errors"
	"gioui.org/io/event"
	"net/url"
)

var (
	// ErrNotReady may ocurr when try to open an URL too earlier
	ErrNotReady = errors.New("some needed library was not loaded yet, make use that you are using ListenEvents()")
	// ErrInvalidURL occur when provide an invalid URL (only HTTPS/HTTP is accepted)
	ErrInvalidURL = errors.New("given url is invalid")
)

var (
	// InsecureIgnoreVerification will remove any attempt to validate the URL
	// However, even when "false" (which means that the verification is enabled) it's NOT RECOMMENDED to open an
	// user-supplied URL.
	InsecureIgnoreScheme bool
)

// ListenEvents must get all the events from Gio, in order to get the GioView once it's ready. You need
// to include that function where you listen for Gio events.
//
// Similar as:
//
// select {
// case e := <-w.Events():
// giohyperlink.ListenEvents(e)
//
// switch e := e.(type) {
// (( ... your code ...  ))
func ListenEvents(event event.Event) {
	listenEvents(event)
}

// OpenURL opens the given url.URL in the browser (or equivalent app)
func OpenURL(u *url.URL) error {
	if u == nil || ((u.Scheme != "http" && u.Scheme != "https") && InsecureIgnoreScheme == false) {
		return ErrInvalidURL
	}

	return open(u)
}

// Open opens the given string url in the browser (or equivalent app).
func Open(uri string) error {
	u, err := url.Parse(uri)
	if err != nil {
		return ErrInvalidURL
	}

	return OpenURL(u)
}