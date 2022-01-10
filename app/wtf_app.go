package app

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/rivo/tview"
	"os"
)

// WtfApp is the container for a collection of widgets that are all constructed from a single
// configuration file and displayed together
type WtfApp struct {
	TViewApp *tview.Application
}

/* -------------------- Exported Functions -------------------- */

// Run starts the underlying tview app
func (wtfApp *WtfApp) Run() {
	if err := wtfApp.TViewApp.Run(); err != nil {
		fmt.Printf("\n%s %v\n", aurora.Red("ERROR"), err)
		os.Exit(1)
	}
}
