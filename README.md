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
}

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)