package appexists

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func AppExists(nameOrBundleId string) bool {
	isBundleId := strings.Contains(nameOrBundleId, ".")
	var query string

	if isBundleId {
		query = fmt.Sprintf("kMDItemContentType == 'com.apple.application-bundle' && kMDItemCFBundleIdentifier == '%s'", nameOrBundleId)
	} else {
		query = fmt.Sprintf("kMDItemKind == 'Application' && kMDItemFSName == '%s.app'", nameOrBundleId)
	}

	dirname, _ := os.UserHomeDir()
	
	output, _ := exec.Command("mdfind", query, "-onlyin", "/Applications", "-onlyin", filepath.Join(dirname, "Applications")).Output()

	return len(string(output)) > 0
}