package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	keyModifierLib "github.com/rsa17826/key-modifier/lib"
)

func main() {
	wt := WindowTracker{}
	wt.engines = append(wt.engines, keyModifierLib.NewEngine())
	wt.engines = append(wt.engines, keyModifierLib.NewEngine())
	if err := wt.engines[0].Connect("key modifier"); err != nil {
		panic(err)
	}

	// Start the event loop once. It reads whatever mods are currently
	// active via SetMods and keeps running for the engine's whole
	// lifetime; window-focus changes never start a second competing loop,
	// they just call SetMods.
	wt.engines[0].EnsureRunning(func(err error) {
		if err != nil {
			fmt.Fprintln(os.Stderr, "engine stopped:", err)
		}
	})

	go wt.listenToHyprland()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT)
	<-sigChan
	wt.engines[0].Close()
	os.Exit(0)
}

type WindowTracker struct {
	mu         sync.Mutex
	LastActive bool
	engines    []*keyModifierLib.Engine
}

func (wt *WindowTracker) listenToHyprland() {
	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	instanceSig := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")
	if runtimeDir == "" || instanceSig == "" {
		fmt.Fprintln(os.Stderr, "Error: missing Hyprland environment variables (XDG_RUNTIME_DIR, HYPRLAND_INSTANCE_SIGNATURE)")
		return
	}

	socketPath := fmt.Sprintf("%s/hypr/%s/.socket2.sock", runtimeDir, instanceSig)
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to Hyprland socket: %v\n", err)
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()

		after, ok := strings.CutPrefix(line, "activewindow>>")
		if !ok {
			continue
		}

		parts := strings.SplitN(after, ",", 2)
		if len(parts) == 0 {
			continue
		}
		wt.mu.Lock()
		wt.windowChanged(parts)
		wt.mu.Unlock()

		// If the engine's loop ever died (e.g. IMan connection hiccup),
		// bring it back up so we don't silently stop modifying keys.
		for idx := range wt.engines {
			if !wt.engines[idx].IsRunning() {
				wt.engines[idx].EnsureRunning(func(err error) {
					if err != nil {
						fmt.Fprintln(os.Stderr, "engine stopped:", idx, err)
					}
				})
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from Hyprland socket: %v\n", err)
	}
}

// windowChanged updates the engine's active modifiers based on which
// window is now focused. Matching windows get their turbo modifier
// activated; anything else clears mods so keys pass through unmodified.
func (wt *WindowTracker) windowChanged(window []string) {
	if len(window) == 0 {
		return
	}

	if strings.HasPrefix(window[0], "Math Quest") {
		wt.engines[0].SetMods(keyModifierLib.ParseModifyArgs([]string{
			"--modify", "d", "turbo", "downfor", "1ms", "delay", "1ms",
		}))
	} else {
		// Condition unmet: stop modifying (releases any held/turbo'd keys).
		wt.engines[0].SetMods(nil)
	}
	wt.engines[1].SetMods(keyModifierLib.ParseModifyArgs([]string{
		"--modify", "5", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "bs",
		"--modify", "6", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "del",
		"--modify", "4", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "browserforward",
		"--modify", "7", "from", "id:usb-04d9_USB_Gaming_Mouse-if01-event-kbd", "replace", "browserback",
	}))
}
