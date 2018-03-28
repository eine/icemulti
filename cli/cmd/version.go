// Copyright © 2018 1138-4EB <@github>
//
// Unless required by applicable law or agreed to in writing, this
// software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.

package cmd

import (
	"runtime"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"

	"github.com/1138-4EB/dtd/cli/helper"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of DTD",
	Long: `All software has versions. This is DTD's.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		printDTDVersion()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func printDTDVersion() {
//	if hugolib.CommitHash == "" {
		buildDate := "2018/03/01" //hugolib.BuildDate
		jww.FEEDBACK.Printf("Yet another Digital Timing Diagram tool v%s %s/%s BuildDate: %s\n", helper.CurrentDTDVersion, runtime.GOOS, runtime.GOARCH, buildDate)
//	} else {
//		jww.FEEDBACK.Printf("Hugo Static Site Generator v%s-%s %s/%s BuildDate: %s\n", helpers.CurrentHugoVersion, strings.ToUpper(hugolib.CommitHash), runtime.GOOS, runtime.GOARCH, hugolib.BuildDate)
//  }
}
