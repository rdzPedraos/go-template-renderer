package builder

import (
	"fmt"
	"net/http"
	"os"
)

// StartServer initializes and starts the HTTP server
func (b *Builder) StartServer() {
	b.registerHandlers()
	b.startListening()
}

// registerHandlers sets up HTTP routes
func (b *Builder) registerHandlers() {
	http.HandleFunc("/", b.serveHTML)
	http.HandleFunc("/events", HandleSSE)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
}

// startListening starts the HTTP server on the configured port
func (b *Builder) startListening() {
	addr := ":" + b.Port
	fmt.Printf("🌐 Server started at http://localhost%s\n", addr)
	fmt.Printf("   Open this URL in your browser.\n\n")

	if err := http.ListenAndServe(addr, nil); err != nil {
		fatalError("starting server", err)
	}
}

// serveHTML handles requests for the main HTML page
func (b *Builder) serveHTML(w http.ResponseWriter, r *http.Request) {
	content := b.readOutputFile()
	htmlContent := b.injectAutoReloadScript(content)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(htmlContent))
}

// readOutputFile reads the rendered output HTML file
func (b *Builder) readOutputFile() string {
	content, err := os.ReadFile(b.OutputPath)
	if err != nil {
		return fmt.Sprintf("<html><body><h1>Error reading output file: %v</h1></body></html>", err)
	}
	return string(content)
}

// injectAutoReloadScript adds live reload functionality to the HTML
func (b *Builder) injectAutoReloadScript(htmlContent string) string {
	reloadScript := `
<script>
const evtSource = new EventSource('/events');
evtSource.onmessage = function(event) {
	if (event.data === 'reload') {
		evtSource.close();
		setTimeout(() => location.reload(), 100);
	}
};
evtSource.onerror = function() {
	console.log('SSE connection lost, retrying...');
};
</script>
</body>`

	return replaceLastOccurrence(htmlContent, "</body>", reloadScript)
}
