package builder

import (
	"flag"
	"os"
	"path/filepath"
)

const (
	defaultTemplatePath = "example/template.html"
	defaultDataPath     = "example/config.json"
	defaultWatch        = true
	defaultPort         = "5500"
)

// ParseFlags reads command-line arguments and returns a configured Builder
func ParseFlags() *Builder {
	templateFlag := flag.String("template", defaultTemplatePath, "Path to HTML template file")
	dataFlag := flag.String("data", defaultDataPath, "Path to JSON data file")
	outputFlag := flag.String("output", "", "Path to output HTML file (default: temp dir in watch mode, output.html otherwise)")
	watchFlag := flag.Bool("watch", defaultWatch, "Enable hot reload mode")
	portFlag := flag.String("port", defaultPort, "Port for HTTP server")
	flag.Parse()

	outputPath := resolveOutputPath(*outputFlag, *watchFlag)

	return &Builder{
		TemplatePath: resolvePath(*templateFlag),
		DataPath:     resolvePath(*dataFlag),
		OutputPath:   resolvePath(outputPath),
		Port:         *portFlag,
		Watch:        *watchFlag,
	}
}

// resolveOutputPath determines the output path based on flags
func resolveOutputPath(outputFlag string, watchMode bool) string {
	if outputFlag != "" {
		return outputFlag
	}

	if watchMode {
		// Watch mode: use temp directory (auto-cleanup on exit)
		return filepath.Join(os.TempDir(), "template-renderer-output.html")
	}

	// Single render: use local file (permanent)
	return "output.html"
}

// resolvePath converts a path to absolute, exits on error
func resolvePath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fatalError("resolving path", err)
	}
	return absPath
}
