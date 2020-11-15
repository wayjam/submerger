package version

import (
	"fmt"
)

var (
	BuildTime string = ""
	GitRev    string = ""
	Ver       string = ""
	GoVersion string = ""
)

func PrintVersion() {
	fmt.Printf("Version:\t%s\n", Ver)
	fmt.Printf("Git Rev:\t%s\n", GitRev)
	fmt.Printf("BuildTime:\t%s\n", BuildTime)
	fmt.Printf("GoVersion:\t%s\n", GoVersion)
}
