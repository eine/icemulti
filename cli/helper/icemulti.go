// Copyright © 2018 1138-4EB <@github>
//
// Unless required by applicable law or agreed to in writing, this
// software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.

package helper

import (
	"fmt"
)

// ICEmultiVersion represents the ICEmulti build version.
type ICEmultiVersion struct {
	// Major and minor version.
	Number float32
	// Increment this for bug releases
	PatchLevel int
	// ICEmultiVersionSuffix is the suffix used in the ICEmulti version string.
	// It will be blank for release versions.
	Suffix string
}

// CurrentICEmultiVersion represents the current build version.
// This should be the only one.
var CurrentICEmultiVersion = ICEmultiVersion{
	Number:     0.0,
	PatchLevel: 0,
	Suffix:     "-dev",
}

func (v ICEmultiVersion) String() string {
	return ICEmultiVersionString(v.Number, v.PatchLevel, v.Suffix)
}

func ICEmultiVersionString(version float32, patchVersion int, suffix string) string {
	if patchVersion > 0 {
		return fmt.Sprintf("%.2f.%d%s", version, patchVersion, suffix)
	}
	return fmt.Sprintf("%.2f%s", version, suffix)
}
