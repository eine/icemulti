// Copyright © 2018 eine <@github>
//
// Unless required by applicable law or agreed to in writing, this
// software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.

package cmd

import (
	"runtime"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"

	"github.com/eine/icemulti/cli/helper"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of DTD",
	Long:  `All software has versions. This is DTD's.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		printICEmultiVersion()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func printICEmultiVersion() {
	//	if hugolib.CommitHash == "" {
	buildDate := "2018/03/01" //hugolib.BuildDate
	jww.FEEDBACK.Printf("Go(lang) tools to exploit warm/cold boot in iCE40 FPGAs v%s %s/%s BuildDate: %s\n", helper.CurrentICEmultiVersion, runtime.GOOS, runtime.GOARCH, buildDate)
	//	} else {
	//		jww.FEEDBACK.Printf("Hugo Static Site Generator v%s-%s %s/%s BuildDate: %s\n", helpers.CurrentHugoVersion, strings.ToUpper(hugolib.CommitHash), runtime.GOOS, runtime.GOARCH, hugolib.BuildDate)
	//  }
}
