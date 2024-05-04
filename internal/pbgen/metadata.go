package pbgen

import (
	"fmt"
	"strconv"
)

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

func (m *ProblemMetadata) ToTable() *MetadataTable {
	return &MetadataTable{
		Headers: []string{"Postată de", "Clasă", "Intrare/Ieșire", "Limită timp", "Limită memorie", "Sursă", "Autor", "Dificultate"},
		Rows: [][]string{
			{m.PostedBy,
				fmt.Sprintf("%d", m.Grade),
				getInputOutput(m.IsInputFile),
				fmt.Sprintf("%s s", strconv.FormatFloat(float64(m.TimeLimit), 'f', -2, 32)),
				fmt.Sprintf("%s / %s", formatMemory(m.MemoryLimit), formatMemory(m.StackLimit)),
				m.Source,
				m.Author,
				m.Difficulty},
		},
	}
}

func (m *ProblemMetadata) ToMarkdown() (string, error) {
	table := NewMetadataTable(m)
	md, err := table.ToMarkdown()
	if err != nil {
		return "", err
	}
	return md, nil
}
