package builder

import (
	"encoding/json"
	"html/template"
	"os"
)

// Builder handles template rendering
type Builder struct {
	TemplatePath string
	DataPath     string
	OutputPath   string
	Port         string
	Watch        bool

	tmpl *template.Template
}

// Render performs a full template rendering cycle
func (b *Builder) Render() {
	b.loadTemplate()
	b.execute()
	NotifyClients("reload")
}

// loadTemplate reads and parses the template file
func (b *Builder) loadTemplate() {
	content, err := os.ReadFile(b.TemplatePath)
	if err != nil {
		fatalError("reading template", err)
	}

	b.tmpl, err = template.New("email").Parse(string(content))
	if err != nil {
		fatalError("parsing template", err)
	}
}

// execute renders the template with data and writes to output file
func (b *Builder) execute() {
	data := b.loadData()

	outputFile, err := os.Create(b.OutputPath)
	if err != nil {
		fatalError("creating output file", err)
	}
	defer outputFile.Close()

	err = b.tmpl.Execute(outputFile, data)
	if err != nil {
		fatalError("executing template", err)
	}
}

// loadData reads and parses the JSON data file
func (b *Builder) loadData() map[string]interface{} {
	content, err := os.ReadFile(b.DataPath)
	if err != nil {
		fatalError("reading data file", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		fatalError("parsing JSON", err)
	}

	return data
}

func (b *Builder) Cleanup() {
	if b.Watch {
		os.Remove(b.OutputPath)
	}
}
