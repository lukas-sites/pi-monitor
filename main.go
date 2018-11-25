package main

import "github.com/zserge/webview"

func main() {
	// Open wikipedia in a 800x600 resizable window
	settings := webview.Settings{
		Title: "Dashboard",
		URL:   "https://dakboard.com/app/screenPredefined?p=eeb986d731e9d0011c5948e88e1741f1",
	}
	view := webview.New(settings)

	view.SetFullscreen(true)
	view.Run()
}
