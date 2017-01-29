# Colog
> Colorful terminal logs

Create colorful logs in terminals. Wrapper around [fatih/color/](https://github.com/fatih/color/).

### Example

```go
package main

import (
	"fmt"
	"github.com/kevingimbel/colog"
)

// var logConfig = &colog.Colog{LogFormat: "%s"}
var Log = colog.Colog{LogFormat: "[%s][%s] -- %s", TimeFormat: "2006"}

func main() {
	fmt.Println(Log.Success("Hello, Success!"))
	fmt.Println(Log.Info("Hello, Info!"))
	fmt.Println(Log.Error("Hello, Error!"))
	fmt.Println(Log.Warn("Hello, Warning!"))
}
```
