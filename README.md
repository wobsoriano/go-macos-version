# go-macos-version

Get or check the current macOS version in Go.

## Install

```bash
$ go get github.com/wobsoriano/go-macos-version
```

## Usage

```go
package main

import (
	"fmt"

	mac "github.com/wobsoriano/go-macos-version"
)

func main() {
	version, err := mac.MacOSVersion()
	// => "10.2.3"

	matches, err := mac.IsMacOSVersion(">10.10")
	// => true

	mac.AssertMacOSVersion(">=10.12.5")
	// Error: Requires macOS >=10.12.5

	mac.assertMacOS()
	// Error: Requires macOS

	if mac.IsMacOS {
		fmt.Println("macOS")
	}
}
```

## Credits

https://github.com/sindresorhus/macos-version

## License

MIT
