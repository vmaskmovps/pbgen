package pbgen

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type MetadataTable struct {
	Headers []string
	Rows    [][]string
}

func ConvertMetadataToMarkdown(metadata ProblemMetadata) (string, error) {
	table := &MetadataTable{
		Headers: []string{"Postată de", "Clasă", "Intrare/Ieșire", "Limită timp", "Limită memorie", "Sursă", "Autor", "Dificultate"},
		Rows: [][]string{
			{metadata.PostedBy,
				fmt.Sprintf("%d", metadata.Grade),
				getInputOutput(metadata.IsInputFile),
				fmt.Sprintf("%s s", strconv.FormatFloat(float64(metadata.TimeLimit), 'f', -2, 32)),
				fmt.Sprintf("%s / %s", formatMemory(metadata.MemoryLimit), formatMemory(metadata.StackLimit)),
				metadata.Source,
				metadata.Author,
				metadata.Difficulty},
		},
	}

	tableMarkdown, err := ConvertTableToMarkdown(table)
	if err != nil {
		return "", err
	}

	return tableMarkdown, nil
}

func ConvertTableToMarkdown(table *MetadataTable) (string, error) {
	var buf bytes.Buffer

	headerWidths := make([]int, len(table.Headers))
	skipColumns := make([]bool, len(table.Headers))
	hasContent := make([]bool, len(table.Headers))
	for i, header := range table.Headers {
		headerWidths[i] = len(header)
		hasContent[i] = false
	}
	for _, row := range table.Rows {
		for i, cell := range row {
			if len(cell) > headerWidths[i] {
				headerWidths[i] = len(cell)
			}
			if cell != "" {
				hasContent[i] = true
			}
		}
	}
	for i := range skipColumns {
		if !hasContent[i] {
			skipColumns[i] = true
		}
	}

	buf.WriteString("|")
	for i, header := range table.Headers {
		if !skipColumns[i] {
			buf.WriteString(fmt.Sprintf(" %-*s |", headerWidths[i], header))
		}
	}
	buf.WriteString("\n|")

	for i, width := range headerWidths {
		if !skipColumns[i] {
			buf.WriteString(fmt.Sprintf("-%s-|", strings.Repeat("-", width)))
		}
	}
	buf.WriteString("\n")

	for _, row := range table.Rows {
		buf.WriteString("|")
		for i, cell := range row {
			if !skipColumns[i] {
				buf.WriteString(fmt.Sprintf(" %-*s |", headerWidths[i], cell))
			}
		}
		buf.WriteString("\n")
	}

	return buf.String(), nil
}

func getInputOutput(isInputFile bool) string {
	if isInputFile {
		return "Fișier"
	}
	return "Tastatură/ecran"
}

func formatMemory(memory float32) string {
	if memory >= 1.0 {
		return fmt.Sprintf("%.0fMB", memory)
	}
	return fmt.Sprintf("%.0fKB", memory*1024)
}

func formatTime(timeLimit float32) string {
	if timeLimit == float32(int(timeLimit)) {
		return fmt.Sprintf("%.0fs", timeLimit)
	}
	return fmt.Sprintf("%.1f s", timeLimit)
}
