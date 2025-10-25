package main

import (
	"fmt"
	"template-renderer/src/builder"
)

func main() {
	b := builder.ParseFlags()
	defer b.Cleanup()

	renderInitialTemplate(b)

	if b.Watch {
		runWatchMode(b)
	} else {
		printSuccessMessage()
	}
}

// renderInitialTemplate performs the first template rendering
func renderInitialTemplate(b *builder.Builder) {
	fmt.Printf("🚀 Rendering initial template...\n")
	b.Render()
	fmt.Printf("✓ Template rendered successfully: %s\n\n", b.OutputPath)
}

// runWatchMode starts the server and file watcher for live development
func runWatchMode(b *builder.Builder) {
	go b.StartServer()
	b.StartWatching()
}

// printSuccessMessage displays completion message for single render
func printSuccessMessage() {
	fmt.Printf("  You can open the file in your browser to see the result.\n")
	fmt.Printf("  Tip: use --watch for hot reload mode with web server.\n")
}
