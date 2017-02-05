package colog

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

// getting the color functions from
// fatih/color. These function can be used like Sprintf
// so cWarn("This is a warning!") produced a yellow string
var (
	cWarn    = color.New(color.FgYellow).SprintFunc()
	cError   = color.New(color.FgRed).SprintFunc()
	cSuccess = color.New(color.FgGreen).SprintFunc()
	cInfo    = color.New(color.FgBlue).SprintFunc()
)

// initialize the config map
var config = make(map[string]string)

// Struct Colog is the heart of Colog and provides all
// String formats needed to create write colorful logs
type Colog struct {
	LogFormat    string
	TimeFormat   string
	InfoLabel    string
	ErrorLabel   string
	WarnLabel    string
	SuccessLabel string
}

// Initialize default config
func init() {
	config["LogFormat"] = "[%s][%s] %s"
	config["TimeFormat"] = "2006-Jan-2 15:04:05"
	config["InfoLabel"] = "INFO"
	config["ErrorLabel"] = "ERROR"
	config["WarnLabel"] = "WARN"
	config["SuccessLabel"] = "SUCCESS"
}

// Returns the log string.  levelStr is the log level string, e.g. "WARN" or "ERROR"
func (c *Colog) createColorLog(str string, levelStr string) string {
	t := time.Now()
	tFormat := t.Format(c.TimeFormat)

	return fmt.Sprintf(c.LogFormat, levelStr, tFormat, str)
}

// Logs a warning using the format specified as "LogFormat" and the string
// specified as "WarnLabel"
func (c *Colog) Warn(str string) string {
	coloredLog := c.createColorLog(str, cWarn(c.WarnLabel))
	return coloredLog
}

// Logs a error using the format specified as "LogFormat" and the string
// specified as "ErrorLabel"
func (c *Colog) Error(str string) string {
	coloredLog := c.createColorLog(str, cError(c.ErrorLabel))
	return coloredLog
}

// Logs a success using the format specified as "LogFormat" and the string
// specified as "SuccessLabel"
func (c *Colog) Success(str string) string {
	coloredLog := c.createColorLog(str, cSuccess(c.SuccessLabel))
	return coloredLog
}

// Logs an info using the format specified as "LogFormat" and the string
// specified as "InfoLabel"
func (c *Colog) Info(str string) string {
	coloredLog := c.createColorLog(str, cInfo(c.InfoLabel))
	return coloredLog
}

// Sets a Colog config. Valid options are all fields from the
// Colog struct.
func (c *Colog) Set(key string, value string) {
	// TODO: I am sure this can be re-written to be more effecient.
	switch varKey := key; varKey {
	case "InfoLabel":
		c.InfoLabel = value
	case "ErrorLabel":
		c.ErrorLabel = value
	case "WarnLabel":
		c.WarnLabel = value
	case "SuccessLabel":
		c.SuccessLabel = value
	case "LogFormat":
		c.LogFormat = value
	case "TimeFormat":
		c.TimeFormat = value
	default:
		fmt.Printf("Unknown settings key '%s' specified.\n", varKey)
	}
}

func extendDefaultConfig(settings map[string]string) map[string]string {
	for key, value := range settings {
		config[key] = value
	}
	return config
}

// Create a new Colog instance by providing settings as map[string]string
// All fields of the Colog struct are valid parameters
//
// Example
// 		config = make(map[string]string)
// 	 	config["ErrorLabel"] = "ERROR"
//		config["TimeFormat"] = "2006-Jan-2"
//		Log := colog.NewColog(config)
//
// All not configured values will use the default values.
func NewColog(settings map[string]string) Colog {
	fullConfig := extendDefaultConfig(settings)
	log := Colog{}
	for key, value := range fullConfig {
		log.Set(key, value)
	}
	return log
}
