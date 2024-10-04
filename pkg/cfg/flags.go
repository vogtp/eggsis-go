package cfg

import (
	"fmt"

	"github.com/spf13/pflag"
)

const (
	// CfgFile
	CfgFile = "config.file"
	// CfgSave triggers periodic config saves
	CfgSave = "config.save"

	// LogLevel error warn info debug
	LogLevel = "log.level"
	// LogSource should we log the source
	LogSource = "log.source"
	// LogJson log in json
	LogJson = "log.json"

	// PlayerDeath indicates if the player can die
	PlayerDeath = "player.death"

	PlayerLP = "player.lp"
)

func init() {
	pflag.Bool(CfgSave, false, "Should the configs be written to file periodically")
	pflag.String(CfgFile, fmt.Sprintf("%s.yml", APP_NAME), "File with the config to load")
	pflag.String(LogLevel, "warn", "Set the loglevel: error warn info debug trace off")
	pflag.Bool(LogSource, false, "Log the source line")
	pflag.Bool(LogJson, false, "Log in json")
	pflag.Bool(PlayerDeath, true, "Player death possible?  (debugging only)")
	pflag.Int(PlayerLP, 100, "Player health")
}
