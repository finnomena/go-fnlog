# FINNOMENA LOG <img src="https://img.shields.io/github/license/finnomena/go-fnlog">

FINNOMENA LOG is a Golang library for logging that compatible with Stackdriver.

## Installation

Use the go get or go mod.

```bash
go get github.com/finnomena/go-fnlog
```

## Usage

```go
package main

import (
	"github.com/finnomena/go-fnlog"
)


func main() {
	fnlog.Info("global")

	logger := fnlog.NewLogger()
	logger.SetLevel(fnlog.TraceLevel)

	logger.Debug("logging with struct")

	var obj = &test{
		logger: logger,
	}

	obj.logger.Warn("log attribute")
	obj.print()

	custom := fnlog.NewLoggerWithOptions(fnlog.Options{
		Writer: os.Stdout,
		Formatter: &fnlog.JSONFormatter{
			Timeformat: time.RFC822Z,
		},
	})

	custom.Info("custom log")
}

```

Result will be

```json
{"serverity":"info","timestamp":"2020-02-12T01:40:40.058983+07:00","caller":"github.com/finnomena/go-fnlog.(*standard).print","message":"global"}
{"serverity":"debug","timestamp":"2020-02-12T01:40:40.059154+07:00","caller":"github.com/finnomena/go-fnlog.(*standard).print","message":"logging with struct"}
{"serverity":"warn","timestamp":"2020-02-12T01:40:40.05916+07:00","caller":"github.com/finnomena/go-fnlog.(*standard).print","message":"log attribute"}
{"serverity":"error","timestamp":"2020-02-12T01:40:40.059164+07:00","caller":"github.com/finnomena/go-fnlog.(*standard).print","message":"log with method"}
{"serverity":"info","timestamp":"12 Feb 20 01:40 +0700","caller":"github.com/finnomena/go-fnlog.(*standard).print","message":"custom log"}
```

### Using with [Echo](https://echo.labstack.com/)

```go
import (
    "github.com/labstack/echo/v4"
    "github.com/finnomena/go-fnlog"
)

func main() {
    e := echo.New()

    e.Use(echo.WrapMiddleware(fnlog.LoggingMiddleware()))
    e.GET("/", func(c echo.Context) error {
        fnlo.Info("test")
        return c.String(http.StatusOK, "Hello, World!")
    })

    e.Logger.Fatal(e.Start(":1323"))
}
```

Use text formatter for reading easier

```go
import (
	"os"
	"github.com/finnomena/go-fnlog"
)

func main() {
	var logger fnlog.Logger

    if os.Getenv("environment") == "develop" {
		logger = fnlog.NewLoggerWithOptions(fnlog.Options{
			Writer: os.Stdout,
			Formatter: &fnlog.TextFormatter{
				Timeformat: "15:04:05",
			},
		})
	} else {
		logger = fnlog.NewLogger()
	}
}
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)