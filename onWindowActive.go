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

	println(title, class, title == "Wine")
	// if strings.Contains(class, "mathbreakers.exe") {
	if title == "Wine Desktop" {
		wt.engines[0].SetMods(keyModifierLib.ParseModifyArgs([]string{
			"--modify", "space", "turbo", "downFor", "20ms", "delay", "20ms",
			"--modify", "space", "maxPressTime", "600ms",
			"--modify", "e", "replace", "r",
			"--modify", "2", "replace", "6",
			"--modify", "3", "replace", "6",
			"--modify", "4", "replace", "j",
			"--modify", "4", "turbo",
			"--modify", "e", "turbo", "downFor", "5ms", "delay", "5ms",
			"--modify", "e", "maxPressTime", "20ms",
			"--modify", "r", "turbo", "downFor", "5ms", "delay", "5ms",
			"--modify", "r", "maxPressTime", "20ms",
			"--modify", "rbutton", "turbo", "downFor", "5ms", "delay", "5ms",
			"--modify", "rbutton", "maxPressTime", "20ms",
			"--modify", "f", "replace", "j",
			"--modify", "rbutton", "replace", "r",
		}))
	} else {
		wt.engines[0].SetMods(keyModifierLib.ParseModifyArgs(getRules(0)))
	}
}

func getRules(i int) []string {
	rules := [][]string{
		{
			"--modify", "5", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "bs",
			"--modify", "6", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "del",
			"--modify", "4", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "browserforward",
			"--modify", "7", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "browserback",
			"--modify", "0", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "combo", "takeover", "<ctrl", "w",
			"--modify", "kpsub", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "combo", "takeover", "<ctrl", "r",
			"--modify", "kpadd", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "combo", "takeover", "<ctrl", "<shift", "t",
		},
	}

	if int(i) >= len(rules) {
		return nil // Return nil or handle out-of-bounds index
	}

	return rules[i]
}
