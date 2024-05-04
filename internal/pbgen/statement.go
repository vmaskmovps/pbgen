package pbgen

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"
)

type Example struct {
	Name        string
	Input       string
	Output      string
	Explanation string
}

type ProblemMetadata struct {
	Id           uint16
	Name         string
	IsInputFile  bool
	IsOutputFile bool
	MemoryLimit  float32
	TimeLimit    float32
	StackLimit   float32
	Grade        uint8
	Source       string
	Author       string
	PostedBy     string
	Difficulty   string
}

func NewProblemMetadata(details *ProblemDetails) *ProblemMetadata {
	memoryLimit, _ := strconv.ParseFloat(details.MemoryLimit, 32)
	timeLimit, _ := strconv.ParseFloat(details.TimeLimit, 32)
	stackLimit, _ := strconv.ParseFloat(details.StackLimit, 32)

	var difficulty string
	switch details.Difficulty {
	case 1:
		difficulty = "Ușoară"
	case 2:
		difficulty = "Medie"
	case 3:
		difficulty = "Dificilă"
	default:
		difficulty = "Concurs"
	}

	metadata := ProblemMetadata{
		Id:           uint16(details.ID),
		Name:         details.Name,
		IsInputFile:  details.UseConsole == "1",
		IsOutputFile: details.UseConsole == "1",
		TimeLimit:    float32(timeLimit),
		MemoryLimit:  float32(memoryLimit),
		StackLimit:   float32(stackLimit),
		Grade:        uint8(details.Grade),
		Source:       details.ProblemSource,
		Author:       details.Author,
		Difficulty:   difficulty,
		PostedBy:     fmt.Sprintf("%s %s", details.User.Prenume, details.User.Nume),
	}

	return &metadata
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

func ConvertProblemToMarkdown(problem *ProblemDetails) (string, error) {
	p := ParseIntoProblem(problem)
	headerTemplate :=
		`# [{{.Metadata.Name}} #{{.Metadata.Id}}](https://new.pbinfo.ro/probleme/{{.Metadata.Id}}/{{.Metadata.Name}})
	
{{ ConvertMetadataToMarkdown .Metadata }}
`
	tmpl := template.New("header")
	tmpl = tmpl.Funcs(template.FuncMap{
		"ConvertMetadataToMarkdown": ConvertMetadataToMarkdown,
	})
	tmpl, err := tmpl.Parse(headerTemplate)
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

	return string(buf.String()), nil
}
