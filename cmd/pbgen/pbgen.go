package pbgen

import (
	"fmt"
	"os"

	"github.com/jxskiss/mcli"
	"github.com/overanalytcl/pbgen/internal/pbgen"
)

var flags struct {
	Id      int    `cli:"#R, -i, --id, Problem ID"`
	Lang    string `cli:"-l, --lang, Programming language for the template [c, cpp, pas]" env:"PB_LANG" default:"cpp"`
	BaseDir string `cli:"-d, --dir, Base path for project" env:"PB_BASEDIR" default:"."`
}

const (
	LANG_C int = iota
	LANG_CPP
	LANG_PAS
)

func Run() {
	app := &mcli.App{
		Description: "Create a project from a PBInfo statement",
		Options:     mcli.Options{AllowPosixSTMO: true},
	}
	app.SetGlobalFlags(&flags)
	mcli.Parse(&flags)

	switch flags.Lang {
	case "c", "cpp", "pas":
		// Lasciate ogne speranza, voi ch'intrate
	default:
		fmt.Println("Language must be c, cpp or pas!")
		mcli.PrintHelp()
		return
	}

	baseDirFile, err := os.Open(flags.BaseDir)
	if err != nil {
		fmt.Printf("Error opening base directory: %v\n", err)
		return
	}
	defer baseDirFile.Close()

	err = pbgen.CreateProject(flags.Lang, flags.Id, baseDirFile)
	if err != nil {
		fmt.Printf("Error creating project: %v\n", err)
		return
	}

	fmt.Println("Project created successfully!")
}
