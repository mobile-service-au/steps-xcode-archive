package utils

import (
	"path/filepath"

	"github.com/bitrise-io/go-utils/pathutil"
	"github.com/bitrise-tools/go-xcode/provisioningprofile"
	"github.com/pkg/errors"
)

const (
	provProfileSystemDirPath = "~/Library/MobileDevice/Provisioning Profiles"
)

// IOSProvProfileWalkFunc ...
type IOSProvProfileWalkFunc func(profile provisioningprofile.Profile) bool

// WalkIOSProvProfiles ...
func WalkIOSProvProfiles(walkFunc IOSProvProfileWalkFunc) error {
	absProvProfileDirPath, err := pathutil.AbsPath(provProfileSystemDirPath)
	if err != nil {
		return errors.Wrap(err, "failed to get Absolute path of Provisioning Profiles dir")
	}

	pths, err := filepath.Glob(absProvProfileDirPath + "/*.mobileprovision")
	if err != nil {
		return errors.Wrap(err, "failed to perform *.mobileprovision search")
	}

	for _, pth := range pths {
		profile, err := provisioningprofile.NewProfileFromFile(pth)
		if err != nil {
			return errors.Wrap(err, "failed to parse Provisioning Profile")
		}

		if breakWalk := walkFunc(profile); breakWalk {
			break
		}
	}

	return nil
}
