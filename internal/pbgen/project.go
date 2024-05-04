package pbgen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateProject(lang string, id int, basedir *os.File) error {
	problem, err := NewProblemFromId(id)
	if err != nil {
		log.Fatal(err)
	}

	projectDir, _ := filepath.Abs(filepath.Join(basedir.Name(), fmt.Sprintf("%04d-%s/%s", id, problem.Metadata.Name, lang)))
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		return err
	}

	markdown, err := problem.ToMarkdown()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(filepath.Join(projectDir, "README.md"))
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)

	switch lang {
	case "c":
		fmt.Println("Created C project!")
	case "cpp":
		fmt.Println("Created C++ project!")
	case "pas":
		fmt.Println("Created Pascal project!")
	default:
		return fmt.Errorf("unsupported language: %s", lang)
	}

	_, err = f.Write([]byte(markdown))
	if err != nil {
		return err
	}
	return nil
}
