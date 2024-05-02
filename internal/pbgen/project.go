package pbgen

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateProject(lang string, id int, basedir *os.File) error {
	projectDir := filepath.Join(basedir.Name(), fmt.Sprintf("%d/%s", id, lang))
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		return err
	}

	switch lang {
	case "c":
		fmt.Println("Created C project!")
		return nil
	case "cpp":
		fmt.Println("Created C++ project!")
		return nil
	case "pas":
		fmt.Println("Created Pascal project!")
		return nil
	default:
		return fmt.Errorf("unsupported language: %s", lang)
	}
}
