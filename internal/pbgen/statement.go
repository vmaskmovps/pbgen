package pbgen

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strconv"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
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

type Problem struct {
	Metadata     ProblemMetadata
	Statement    string
	InputData    string
	OutputData   string
	Restrictions string
	Examples     []Example
}

func ConvertProblemToMarkdown(problem *ProblemDetails) (string, error) {
	id := uint16(problem.ID)
	usesConsole := (problem.UseConsole == "0")
	memoryLimit, _ := strconv.ParseFloat(problem.MemoryLimit, 32)
	timeLimit, _ := strconv.ParseFloat(problem.TimeLimit, 32)
	stackLimit, _ := strconv.ParseFloat(problem.StackLimit, 32)
	grade := uint8(problem.Grade)

	var difficulty string
	switch problem.Difficulty {
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
		Id:           id,
		Name:         problem.Name,
		IsInputFile:  !usesConsole,
		IsOutputFile: !usesConsole,
		TimeLimit:    float32(timeLimit),
		MemoryLimit:  float32(memoryLimit),
		StackLimit:   float32(stackLimit),
		Grade:        grade,
		Source:       problem.ProblemSource,
		Author:       problem.Author,
		Difficulty:   difficulty,
		PostedBy:     fmt.Sprintf("%s %s", problem.User.Prenume, problem.User.Nume),
	}

	p := Problem{
		Metadata: metadata,
	}
	htmlTemplate :=
		`<h1><a href="https://new.pbinfo.ro/probleme/{{.Metadata.Id}}/{{.Metadata.Name}}">{{.Metadata.Name}} #{{.Metadata.Id}}</a></h1>
		
<table>
	<tr>
		<th>Postată de</th>
		<th>Clasă</th>
		<th>Intrare/Ieșire</th>
		<th>Limită timp</th>
		<th>Limită memorie</th>
		{{if .Metadata.Source}}<th>Sursă</th>{{end}}
		{{if .Metadata.Author}}<th>Autor</th>{{end}}
		<th>Dificultate</th>
	</tr>
	<tr>
		<td>{{.Metadata.PostedBy}}</td>
		<td>{{.Metadata.Grade}}</td>
		<td>{{if .Metadata.IsInputFile}}Fișier{{else}}Tastatură/ecran{{end}}</td>
		<td>{{.Metadata.TimeLimit}} s</td>
		<td>{{.Metadata.MemoryLimit}}MB / {{.Metadata.StackLimit}}MB</td>
		{{if .Metadata.Source}}<td>{{.Metadata.Source}}</td>{{end}}
		{{if .Metadata.Author}}<td>{{.Metadata.Author}}</td>{{end}}
		<td>{{.Metadata.Difficulty}}</td>
	</tr>
</table>`
	tmpl, err := template.New("html").Parse(htmlTemplate)
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

	mark := md.NewConverter("new.pbinfo.ro", true, nil)
	mark.Use(plugin.Table())
	markdown, err := mark.ConvertString(buf.String())
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return markdown, nil
}
