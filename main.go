package macosversion

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/Masterminds/semver"
)

const isMacOS = runtime.GOOS == "darwin"
var version string

func clean(version string) string {
	items := strings.Split(version, ".")

	if len(items) == 1 {
		return fmt.Sprintf("%s.0.0", version)
	}

	if len(items) == 2 {
		return fmt.Sprintf("%s.0", version)
	}

	return version
}

func parseVersion(plist string) string {
	re := regexp.MustCompile(`<key>ProductVersion<\/key>\s*<string>([\d.]+)<\/string>`)

	if len(re.FindStringIndex(plist)) == 0 {
		panic("Unable to fetch macOS version")
	}

	match := re.FindAllStringSubmatch(plist, -1)[0][1]
	match = strings.Replace(match, "10.16", "11", -1)
	return match
}

func macOSVersion() (string, error) {
	if !isMacOS {
		return "", errors.New("requires macOS")
	}

	if (len(version) == 0) {
		content, _ := os.ReadFile("/System/Library/CoreServices/SystemVersion.plist")
		match := parseVersion(string(content))
		
		version = clean(match)
	}

	return version, nil
}

func isMacOSVersion(semverRange string) (bool, error) {
	if !isMacOS {
		return false, errors.New("requires macOS")
	}

	semverRange = strings.Replace(semverRange, "10.16", "11", -1)

	c, err := semver.NewConstraint(clean(semverRange))

	if err != nil {
		return false, errors.New("unable to parse constraint")
	}

	macV, _ := macOSVersion()
	v, err := semver.NewVersion(macV)
	if err != nil {
		return false, errors.New("unable to parse macOS version")
	}

	return c.Check(v), nil
}

func assertMacOSVersion(semverRange string) {
	semverRange = strings.Replace(semverRange, "10.16", "11", -1)

	r, err := isMacOSVersion(semverRange)
	if err != nil || !r {
		panic(fmt.Sprintf("Requires macOS %s", semverRange))
	}
}

func assertMacOS() {
	if !isMacOS {
		panic("Requires macOS")
	}
}