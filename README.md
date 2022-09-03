GioHyperlink
--------

Opens a hyperlink in the default browser. 🤩

## Setup

First, you need to `go get github.com/inkeliz/giohyperlink`, then you need to provide `giohyperlink` access to the Window events, so you need to add the following to your main loop function:

```rust

```diff
 for evt := range w.Events() { // Gio main event loop
+    giohyperlink.ListenEvents(e)

    switch evt := evt.(type) {
        // ...
    }
}
```

> _⚠️In some OSes (Windows, macOS...) this setup is optional, but it's recommended to do it anyway._

## Usage

To open one link, you can use the `Open` function:

```go
giohyperlink.Open("https://github.com")
```

That will open the link in the default browser. You can use `OpenURL` to open a `*url.URL`:

```go
giohyperlink.OpenURL(&url.URL{
    Scheme: "https",
    Host:   "github.com",
})
```

By default only HTTP and HTTPS links are allowed, but you can change that by changing `InsecureIgnoreScheme` to `true`,
you should validate the URL and scheme on your own.
