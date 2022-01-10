package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bernylinville/wtf/app"
	"github.com/bernylinville/wtf/cfg"
	"github.com/bernylinville/wtf/flags"
	"github.com/bernylinville/wtf/utils"
	"github.com/bernylinville/wtf/wtf"
	"github.com/logrusorgru/aurora"
	"github.com/pkg/profile"
)

var (
	date    = "dev"
	version = "dev"
)

/* -------------------- Main -------------------- */
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Parse and handle flags
	flags := flags.NewFlags()
	flags.Parse()

	// Load the configuration file
	cfg.Initialize(flags.HasCustomConfig())
	config := cfg.LoadWtfConfigFile(flags.ConfigFilePath())

	wtf.SetTerminal(config)

	flags.RenderIf(version, date, config)

	if flags.Profile {
		defer profile.Start(profile.MemProfile).Stop()
	}

	openFileUtil := config.UString("wtf.openFileUtil", "open")
	openURLUtil := utils.ToStrs(config.UList("wtf.openURLUtil", []interface{}{}))
	utils.Init(openFileUtil, openURLUtil)

	/* Initialize the App Manager */
	appMan := app.NewAppManager()
	appMan.MakeNewWtfApp(config, flags.Config)

	currentApp, err := appMan.Current()
	if err != nil {
		fmt.Printf("\n%s %v\n", aurora.Red("ERROR"), err)
		os.Exit(1)
	}

	currentApp.Run()
}
