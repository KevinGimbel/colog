package colog

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

// color functions
var (
	cWarn    = color.New(color.BgBlack, color.FgYellow).SprintFunc()
	cError   = color.New(color.BgBlack, color.FgRed).SprintFunc()
	cSuccess = color.New(color.BgBlack, color.FgGreen).SprintFunc()
	cInfo    = color.New(color.BgBlack, color.FgBlue).SprintFunc()
)

// Prepared strings for colored output
var (
	sWarn    = cWarn("WARNING")
	sError   = cError("ERROR  ")
	sSuccess = cSuccess("SUCCESS")
	sInfo    = cInfo("INFO   ")
	// All dates are formated based on "Mon Jan 2 15:04:05 -0700 MST 2006"
	// So to get "2017-Jan-29" the formt must be "2006-Jan-2"
	sTime = "2006-Jan-2 15:04:05"
)

type Colog struct {
	LogFormat  string
	TimeFormat string
}

func (c *Colog) createColorLog(str string, coloredLog string) string {
	t := time.Now()
	tFormat := t.Format(c.TimeFormat)

	return fmt.Sprintf(c.LogFormat, coloredLog, tFormat, str)
}

func (c *Colog) Warn(str string) string {
	coloredLog := c.createColorLog(str, sWarn)
	return coloredLog
}

func (c *Colog) Error(str string) string {
	coloredLog := c.createColorLog(str, sError)
	return coloredLog
}

func (c *Colog) Success(str string) string {
	coloredLog := c.createColorLog(str, sSuccess)
	return coloredLog
}

func (c *Colog) Info(str string) string {
	coloredLog := c.createColorLog(str, sInfo)
	return coloredLog
}

// func main() {
// 	fmt.Println(Warn("OMG"))
// 	fmt.Println(Info("This is a information log"))
// 	fmt.Println(Error("This is an error!"))
// 	fmt.Println(Success("This is a success message"))
// }
