package builder

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

// StartWatching monitors file changes and triggers re-renders
func (b *Builder) StartWatching() {
	watcher := b.createWatcher()
	defer watcher.Close()

	b.addFilesToWatch(watcher)
	b.printWatchInfo()
	b.watchLoop(watcher)
}

// createWatcher initializes a new file system watcher
func (b *Builder) createWatcher() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fatalError("creating watcher", err)
	}
	return watcher
}

// addFilesToWatch registers template and data files with the watcher
func (b *Builder) addFilesToWatch(watcher *fsnotify.Watcher) {
	// Paths are already absolute from config.go
	if err := watcher.Add(b.TemplatePath); err != nil {
		fatalError("watching template", err)
	}
	if err := watcher.Add(b.DataPath); err != nil {
		fatalError("watching data file", err)
	}
}

// printWatchInfo displays watching status information
func (b *Builder) printWatchInfo() {
	fmt.Printf("👀 Watching for changes...\n")
	fmt.Printf("   - Template: %s\n", b.TemplatePath)
	fmt.Printf("   - Data: %s\n", b.DataPath)
	fmt.Printf("\nPress Ctrl+C to stop.\n\n")
}

// watchLoop handles file system events
func (b *Builder) watchLoop(watcher *fsnotify.Watcher) {
	var debounceTimer *time.Timer
	const debounceDelay = 100 * time.Millisecond

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if b.isWriteEvent(event) {
				debounceTimer = b.scheduleRender(debounceTimer, debounceDelay, event)
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Printf("⚠️  Watcher error: %v\n", err)
		}
	}
}

// isWriteEvent checks if the event is a write or create operation
func (b *Builder) isWriteEvent(event fsnotify.Event) bool {
	return event.Op.Has(fsnotify.Write) || event.Op.Has(fsnotify.Create)
}

// scheduleRender debounces render calls to avoid multiple renders
func (b *Builder) scheduleRender(timer *time.Timer, delay time.Duration, event fsnotify.Event) *time.Timer {
	if timer != nil {
		timer.Stop()
	}

	return time.AfterFunc(delay, func() {
		fmt.Printf("🔄 Change detected: %s\n", filepath.Base(event.Name))
		fmt.Printf("⏳ Re-rendering...\n")

		b.Render()

		fmt.Printf("✓ Template updated successfully: %s\n", b.OutputPath)
		fmt.Printf("  %s\n\n", time.Now().Format("15:04:05"))
	})
}
