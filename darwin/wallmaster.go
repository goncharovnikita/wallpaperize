// +build darwin

package darwin

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

// DarwinWallMaster implementation
type DarwinWallMaster struct{}

// Get returns the path to the current wallpaper.
func (DarwinWallMaster) Get() (string, error) {
	stdout, err := exec.Command("osascript", "-e", `tell application "Finder" to get POSIX path of (get desktop picture as alias)`).Output()
	if err != nil {
		return "", err
	}

	// is calling strings.TrimSpace() necessary?
	return strings.TrimSpace(string(stdout)), nil
}

// SetFromFile uses AppleScript to tell Finder to set the desktop wallpaper to specified file.
func (DarwinWallMaster) SetFromFile(file string) error {
	cmd := exec.Command("osascript", "-e", `tell application "Finder" to set desktop picture to POSIX file `+strconv.Quote(file))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getCacheDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Join(usr.HomeDir, "Library", "Caches"), nil
}
