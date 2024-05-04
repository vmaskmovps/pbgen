package pbgen

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"
)

type Example struct {
	Name        string
	Input       string
	Output      string
	Explanation string
}

type Problem struct {
	Metadata     ProblemMetadata
	Statement    string
	InputData    string
	OutputData   string
	Restrictions string
	Examples     []Example
}

func ParseIntoProblem(problem *ProblemDetails) *Problem {
	p := Problem{
		Metadata:     *NewProblemMetadata(problem),
		Statement:    problem.Statement,
		InputData:    "",
		OutputData:   "",
		Restrictions: "",
		Examples:     make([]Example, 0),
	}

	return &p
}

//go:embed problem.tmpl
var problemTemplate string

func ConvertProblemToMarkdown(details *ProblemDetails) (string, error) {
	pb := ParseIntoProblem(details)

	tmpl := template.New("details")
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
	err = tmpl.Execute(&buf, pb)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return "", err
	}

	return string(buf.String()), nil
}
