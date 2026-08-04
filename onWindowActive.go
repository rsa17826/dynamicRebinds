package main

import (
	keyModifierLib "github.com/rsa17826/key-modifier/lib"
)

func (wt *WindowTracker) windowChanged(window []string) {
	class := window[0]
	title := window[1]
	// fmt.Printf("Window length: %d, Content: %v\n", len(window), window[1])
	if class == "" || title == "" {
	}
	if len(window) == 0 {
		return
	}

	// if strings.Contains(title, "http://127.0.0.1:1533/MathQuest/play.html") {
	// 	wt.engines[0].SetMods(keyModifierLib.ParseModifyArgs([]string{
	// 		"--modify", "d", "turbo", "downfor", "1ms", "delay", "1ms",
	// 	}))
	// } else {
	// 	wt.engines[0].SetMods(nil)
	// }

	wt.engines[1].SetMods(keyModifierLib.ParseModifyArgs([]string{
		"--modify", "5", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "bs",
		"--modify", "6", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "del",
		"--modify", "4", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "browserforward",
		"--modify", "7", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "browserback",
		"--modify", "0", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "combo", "takeover", "<ctrl", "w",
	}))
}
