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

	mac.AssertMacOS()
	// Error: Requires macOS

	if mac.IsMacOS {
		fmt.Println("macOS")
	}
}
```

## API

### MacOSVersion()

Returns the macOS version or an error if the platform is not macOS.

### IsMacOSVersion(semverRange)

Returns a `bool` if whether the specified [semver range](https://github.com/Masterminds/semver#basic-comparisons) matches the macOS version or an error if there is an issue parsing the version.

### AssertMacOSVersion(semverRange)

Throws an error if the specified [semver range](https://github.com/Masterminds/semver#basic-comparisons) does not match the macOS version.

### AssertMacOS()

Throws an error if the platform is not macOS.

### IsMacOS

Type: `bool`

Whether the platform is macOS.

## Credits

https://github.com/sindresorhus/macos-version

## License

MIT
