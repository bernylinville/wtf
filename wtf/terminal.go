package wtf

import "github.com/olebedev/config"

// SetTerminal sets the TERM environment variable, defaulting to whatever the OS
// has as the current value if none is specified.
// See https://www.gnu.org/software/gettext/manual/html_node/The-TERM-variable.html for
// more details.
func SetTerminal(config *config.Config) {}
