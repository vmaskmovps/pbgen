package pbgen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func CreateProject(lang string, id int, basedir *os.File) error {
	projectDir := filepath.Join(basedir.Name(), fmt.Sprintf("%d/%s", id, lang))
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		return err
	}

	problem, err := GetProblemDetails(id)
	if err != nil {
		log.Fatal(err)
	}

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertBytes([]byte(problem.Statement))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("md ->", string(markdown))

	f, err := os.Create(fmt.Sprintf("%d/%s/cerinta.md", id, lang))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

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

	f.Write(markdown)
	return nil
}
