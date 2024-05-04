package pbgen

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"
)

type Problem struct {
	Metadata     ProblemMetadata
	Statement    string
	InputData    string
	OutputData   string
	Restrictions string
	Examples     []struct {
		Name        string
		Input       string
		Output      string
		Explanation string
	}
}

func NewProblem(details *ProblemDetails) *Problem {
	return &Problem{
		Metadata:     *NewProblemMetadata(details),
		Statement:    details.Statement,
		InputData:    "",
		OutputData:   "",
		Restrictions: "",
		Examples: make([]struct {
			Name        string
			Input       string
			Output      string
			Explanation string
		}, 0),
	}
}

func NewProblemFromId(id int) (*Problem, error) {
	details, err := NewProblemDetails(id)
	if err != nil {
		return nil, err
	}

	return NewProblem(details), nil
}

//go:embed problem.tmpl
var problemTemplate string

func (p *Problem) ToMarkdown() (string, error) {
	tmpl := template.New("problem")
	tmpl = tmpl.Funcs(template.FuncMap{
		"MetadataToMarkdown": func(metadata ProblemMetadata) (string, error) {
			md, err := metadata.ToMarkdown()
			if err != nil {
				return "", err
			}
			return md, nil
		},
	})

	tmpl, err := tmpl.Parse(problemTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, p)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return "", err
	}

	return buf.String(), nil
}
