package app

import "github.com/olebedev/config"

// WtfAppManager handles the instances of WtfApp, ensuring that they're displayed as requested
type WtfAppManager struct {
	WtfApps []*WtfApp

	selected int
}

// NewAppManager creates and returns an instance of AppManager
func NewAppManager() WtfAppManager {
	appMan := WtfAppManager{
		WtfApps: []*WtfApp{},
	}

	return appMan
}

// MakeNewWtfApp creates and starts a new instance of WtfApp from a set of configuration params
func (appMan *WtfAppManager) MakeNewWtfApp(config *config.Config, configFilePath string) {}

// Current returns the currently-displaying instance of WtfApp
func (appMan *WtfAppManager) Current() (*WtfApp, error) {
	return appMan.WtfApps[appMan.selected], nil
}
