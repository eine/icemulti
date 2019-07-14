// Copyright © 2018 1138-4EB <@github>
//
// Unless required by applicable law or agreed to in writing, this
// software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.

package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/1138-4EB/icemulti/lib"
)

// listCmd represents the serve command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Scan and list binary and metadata files",
	Long: `A list of paths and/or files are scanned to find binary files (either
bitstreams or packs) and metadata (JSON) files. Metadata is retained only
related to found binary files.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose != 0 {
			fmt.Println("[icemulti] list", len(args), "args")
			if Verbose > 1 {
				for _, a := range args {
					fmt.Println(a)
				}
			}
		}
		bins, err := lib.List(args, Recursive)
		if err != nil {
			fmt.Printf("list error [%v]\n", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Size", "ModTime", "Design", "Device", "Board"})

		for _, v := range bins {
			table.Append([]string{
				v.Name,
				fmt.Sprintf("%d", v.Size),
				v.ModTime.Format("06/01/02 15:04"),
				//v.Meta.Name,
				v.Meta.Device,
				v.Meta.Board,
			})
		}

		/*
			table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
				tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
				tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor})

			table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor})
		*/
		table.Render() // Send output

		if Output != "" {
			fmt.Println("SHOULD WRITE TO FILE", Output)
		}
	},
}

var Recursive bool

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")
	listCmd.PersistentFlags().BoolVarP(&Recursive, "recursive", "r", false, "scan directories recursively")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
