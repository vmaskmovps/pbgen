package pbgen

import (
	"fmt"

	"github.com/jxskiss/mcli"
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

var langId int

func Run() {
	app := &mcli.App{
		Description: "Create a project from a PBInfo statement",
		Options:     mcli.Options{AllowPosixSTMO: true},
	}
	app.SetGlobalFlags(&flags)
	mcli.Parse(&flags)
	switch flags.Lang {
	case "c":
		langId = LANG_C
	case "cpp":
		langId = LANG_CPP
	case "pas":
		langId = LANG_PAS
	default:
		fmt.Println("Language must be c, cpp or pas!")
		mcli.PrintHelp()
		return
	}
}
