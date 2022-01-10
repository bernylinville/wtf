package flags

import "github.com/olebedev/config"

// Flags is the container for command line flag data
type Flags struct {
	Config  string `short:"c" long:"config" optional:"yes" description:"Path to config file"`
	Profile bool   `short:"p" long:"profile" optional:"yes" description:"Profile application memory usage"`

	hasCustom bool
}

// NewFlags creates an instance of Flags
func NewFlags() *Flags {
	flags := Flags{}
	return &flags
}

/* -------------------- Exported Functions -------------------- */

// ConfigFilePath returns the path to the currently-loaded config file
func (flags *Flags) ConfigFilePath() string {
	return flags.Config
}

// RenderIf displays special-case information based on the flags passed
// in, if any flags were passed in
func (flags *Flags) RenderIf(version, date string, config *config.Config) {}

// HasCustomConfig returns TRUE if a config path was passed in, FALSE if one was not
func (flags *Flags) HasCustomConfig() bool {
	return flags.hasCustom
}

// Parse parses the incoming flags
func (flags *Flags) Parse() {}
