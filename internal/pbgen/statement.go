package pbgen

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

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
	IsInputFile  bool
	IsOutputFile bool
	MemoryLimit  float32
	TimeLimit    float32
	StackLimit   float32
	Grade        uint8
	Source       string
	Author       string
	PostedBy     string
	Difficulty   uint8
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
	htmlTemplate :=
		`<h1><a href="https://new.pbinfo.ro/probleme/{{.ID}}/{{.Name}}">{{.Name}} #{{.ID}}</a></h1>
		
<table>
	<tr>
		<th>Postată de</th>
		<th>Clasă</th>
		<th>Intrare/Ieșire</th>
		<th>Limită timp</th>
		<th>Limită memorie</th>
		<th>Sursă</th>
		<th>Autor</th>
		<th>Dificultate</th>
	</tr>
	<tr>
		<td>{{.User.Prenume}} {{.User.Nume}}</td>
		<td>{{.Grade}}</td>
		{{if eq .UseConsole "0"}}
		<td>Fișier</td>
		{{else}}
		<td>Tastatură/Ecran</td>
		{{end}}
		<td>{{.TimeLimit}} s</td>
		<td>{{.MemoryLimit}}MB / {{.StackLimit}}MB</td>
		<td>{{.ProblemSource}}</td>
		<td>{{.Author}}</td>
		<td>
		{{if eq .Difficulty 1}}
		Ușoară
		{{else if eq .Difficulty 2}}
		Medie
		{{else if eq .Difficulty 3}}
		Dificilă
		{{else}}
		Concurs
		{{end}}</td>
	</tr>
</table>`
	tmpl, err := template.New("html").Parse(htmlTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, problem)
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
