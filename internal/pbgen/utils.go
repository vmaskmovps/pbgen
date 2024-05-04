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

func NewMetadataTable(metadata *ProblemMetadata) *MetadataTable {
	return metadata.ToTable()
}

func (mt *MetadataTable) ToMarkdown() (string, error) {
	var buf bytes.Buffer

	headerWidths := make([]int, len(mt.Headers))
	skipColumns := make([]bool, len(mt.Headers))
	hasContent := make([]bool, len(mt.Headers))
	for i, header := range mt.Headers {
		headerWidths[i] = len(header)
		hasContent[i] = false
	}
	for _, row := range mt.Rows {
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
	for i, header := range mt.Headers {
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

	for _, row := range mt.Rows {
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
	return fmt.Sprintf("%s MB", strconv.FormatFloat(float64(memory), 'f', -1, 32))
}
