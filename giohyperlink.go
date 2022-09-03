package giohyperlink

import (
	"errors"
	"net/url"

	"gioui.org/io/event"
)

var (
	// ErrNotReady may occur when try to open a URL before the initialization is done.
	ErrNotReady = errors.New("some needed library was not loaded yet, make use that you are using ListenEvents()")
	// ErrInvalidURL occur when provide an invalid URL, like a non http/https URL.
	ErrInvalidURL = errors.New("given url is invalid")
)

var (
	// InsecureIgnoreScheme will remove any attempt to validate the URL
	// It's "false" by default. Set it to "true" if you are using a custom scheme (like "myapp://").
	InsecureIgnoreScheme bool
)

// ListenEvents must get all the events from Gio, in order to get the GioView once it's ready. You need
// to include that function where you listen for Gio events.
//
// Similar as:
//
//	select {
//	case e := <-w.Events():
//		giohyperlink.ListenEvents(e)
//
//		switch e := e.(type) {
//	     (( ... your code ...  ))
func ListenEvents(event event.Event) {
	listenEvents(event)
}

// OpenURL opens the given url.URL in the browser (or equivalent app)
func OpenURL(u *url.URL) error {
	if u == nil || u.Scheme == "" || ((u.Scheme != "http" && u.Scheme != "https") && InsecureIgnoreScheme == false) {
		return ErrInvalidURL
	}

	return open(u)
}

// Open opens the given string url in the browser (or equivalent app).
func Open(uri string) error {
	if uri == "" {
		return ErrInvalidURL
	}

	u, err := url.Parse(uri)
	if err != nil {
		return ErrInvalidURL
	}

	return OpenURL(u)
}
