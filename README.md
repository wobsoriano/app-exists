# app-exists

Check if an app exists on macOS in Go.

## Install

```bash
$ go get github.com/wobsoriano/app-exists
```

## Usage

```go
package main

import (
	"fmt"

	exists "github.com/wobsoriano/app-exists"
)

func main() {
	fmt.Println("App exists: ", exists.AppExists())
}
```

## Credits

https://github.com/sindresorhus/app-exists

## License

MIT
