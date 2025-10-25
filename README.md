# Go Template Renderer

A powerful and flexible Go template rendering tool with hot reload capabilities and built-in web server for live development.

🆓 **Open Source** - MIT License. Free to use, modify and distribute in your projects!

## Features

✨ **Template Rendering**: Render HTML templates with JSON data  
🔄 **Hot Reload**: Automatic re-rendering when files change  
⚡ **Fast**: Optimized with debouncing to prevent excessive renders  
📦 **Standalone**: Single binary, easy to distribute

## Quick Start

### Using the Pre-built Binary

The easiest way to use this tool is with the pre-built `tmp-render` binary:

```bash
# Watch mode with hot reload (default behavior)
./tmp-render
```

Then open your browser at [http://localhost:5500](http://localhost:5500)

```bash
# Single render without watch mode
./tmp-render --watch=false
```

### Using with Go

Alternatively, you can run directly with Go:

```bash
# Watch mode (default)
go run main.go

# Single render without watch mode
go run main.go --watch=false
```

## Command Line Options

| Flag         | Default                   | Description                                                                                      |
| ------------ | ------------------------- | ------------------------------------------------------------------------------------------------ |
| `--template` | `example/template.html`   | Path to HTML template file                                                                       |
| `--data`     | `example/config.json`     | Path to JSON data file                                                                           |
| `--output`   | (temp file / output.html) | Path to output HTML file. Default: temp file in watch mode, `output.html` when watch is disabled |
| `--port`     | `5500`                    | Port for HTTP server (watch mode only)                                                           |
| `--watch`    | `true`                    | Enable hot reload mode with web server                                                           |

## Template Syntax

This tool uses Go's standard `html/template` package. Here's a quick reference:

### Basic Variable

```html
<p>Hello, {{.name}}!</p>
```

### Conditionals

```html
{{if .active}}
<span>Active</span>
{{else}}
<span>Inactive</span>
{{end}}
```

### Loops

```html
{{range .items}}
<li>{{.}}</li>
{{end}}
```

### Nested Data

```html
<p>{{.user.email}}</p>
```

For more details, see [Go template documentation](https://pkg.go.dev/html/template).

## Data Format

JSON data files should contain an object with your template variables:

```json
{
  "customer_name": "John Doe",
  "email": "john@example.com",
  "items": ["item1", "item2"],
  "active": true
}
```

## Use Cases

- 📧 **Email Template Development**: Build and preview email templates
- 📄 **Report Generation**: Generate HTML reports from data
- 🎨 **Static Site Building**: Create static HTML pages
- 🧪 **Template Testing**: Test templates with different data sets

## Troubleshooting

### Port Already in Use

If port 5500 is already in use:

```bash
./tmp-render --port=8080
```

### Template Not Found

Ensure your template path is correct relative to where you run the command:

```bash
./tmp-render --template=path/to/your/template.html
```

### Changes Not Reflecting

1. Check that you're running in watch mode (it's enabled by default)
2. Verify the browser console for SSE connection
3. Look for error messages in the terminal

## Tips

1. **Keep your browser's DevTools open** to see SSE connection status in watch mode
2. **Use the debounce feature** by saving files normally - multiple rapid saves are handled automatically
3. **Check the console output** for rendering errors and timing information
4. **Use `--watch=false`** if you only need a single render without the server

---

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Installation for Development

```bash
# Clone the repository
git clone <your-repo-url>
cd templates

# Download dependencies
go mod download

# Build the binary
go build -o tmp-render
```

### Project Structure

```
.
├── main.go           # Entry point and main flow
├── tmp-render        # Pre-built binary
├── src/              # Source code package
│   └── builder/
│       ├── main.go      # Template rendering logic
│       ├── watcher.go   # File watching and change detection
│       ├── server.go    # HTTP server and live reload
│       ├── sse.go       # Server-Sent Events implementation
│       ├── config.go    # Configuration and flag parsing
│       └── utils.go     # Utility functions
├── example/          # Example template and data
│   ├── template.html
│   └── config.json
├── img/              # Static assets
├── go.mod            # Go module definition
├── go.sum            # Dependency checksums
└── README.md         # This file
```

### How It Works

#### Watch Mode (Default)

1. Perform initial render to temporary file
2. Start HTTP server on specified port (default: 5500)
3. Watch template and data files for changes
4. On file change:
   - Re-render the template
   - Notify connected browsers via Server-Sent Events (SSE)
   - Browser automatically reloads to show new content

#### Single Render Mode (`--watch=false`)

1. Parse command-line flags
2. Load and parse the HTML template
3. Read and parse JSON data
4. Execute template with data
5. Write output to `output.html` (or custom path with `--output`)

### Development

#### Running Tests

```bash
go test ./...
```

#### Building for Production

```bash
# Build for current platform
go build -o tmp-render

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o tmp-render-linux

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o tmp-render.exe

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o tmp-render-mac
```

## Author

Created with ❤️ for developers who love fast feedback loops.
