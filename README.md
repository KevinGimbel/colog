# Colog
> Colorful terminal logs

Create colorful logs in terminals. Wrapper around [github.com/fatih/color](https://github.com/fatih/color/).

### Example

```go
package main

import (
	"fmt"
	"github.com/kevingimbel/colog"
)

// Create a new Log instance
var Log = colog.Colog{LogFormat: "[%s][%s] -- %s", TimeFormat: "2006"}

func main() {
// Use it.
	fmt.Println(Log.Success("Hello, Success!"))
	fmt.Println(Log.Info("Hello, Info!"))
	fmt.Println(Log.Error("Hello, Error!"))
	fmt.Println(Log.Warn("Hello, Warning!"))
}
```

!["Colog example"](./.github/colog.png)

## API

### `Colog` struct

```go
type Colog struct {
	LogFormat  string
	TimeFormat string
}
```

`LogFormat` is a string passed to `fmt.Sprintf` with three variable strings in the following order:

- Log type (WARN, SUCCESS, ERROR, INFO)
- Time format (based on `Colog.TimeFormat`)
- The string to log.

**Examples**

```go

var Log = colog.Colog{LogFormat: "%s - (%s) -- %s", TimeFormat: "2006"}
Log.Warn("Some string")
// => WARNING - 2017 - Some string

var Log = colog.Colog{LogFormat: "[%s]|(%s) >> %s", TimeFormat: "2006-Jan-2"}
Log.Error("Some string")
// => [WARNING]|(2017-Jan-29) >> Some string
```

## Functions

### `Info(str string) string`

Creates a blue colored log output. Makes use of `LogFormat` and `TimeFormat`.

### `Success(str string) string`

Creates a green colored log output. Makes use of `LogFormat` and `TimeFormat`.

### `Warn(str string) string`

Creates a yellow colored log output. Makes use of `LogFormat` and `TimeFormat`.

### `Error(str string) string`

Creates a red colored log output. Makes use of `LogFormat` and `TimeFormat`.
