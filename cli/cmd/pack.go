// Copyright © 2018 1138-4EB <@github>
//
// Unless required by applicable law or agreed to in writing, this
// software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// packCmd represents the serve command
var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Make a pack out of multiple image-bitstreams",
	Long: `A set of bitstream images are appended and an index
	is generated. Pointers for five of them are added to the
	first positions of the resulting pack. This resemble the
	'applet' expected by iCE40 devices when cold/warm boot is
	used.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pack called")
		fmt.Println(cmd)
		fmt.Println(len(args))
	},
}

var Mode bool
var Compress string
var Base int
var PointersStr string
var Pointers [4]int
var Align int

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.PersistentFlags().BoolVarP(&Mode, "mode", "m", false, "enable cold/warm boot mode (pointer selected by CBSEL0:1)")
	packCmd.PersistentFlags().IntVarP(&Base, "base", "b", 0, "power on reset image when not using cold/warm boot")
	packCmd.PersistentFlags().StringVarP(&PointersStr, "pointers", "p", "1,2,3,4", "four default-addressable pointers")
	packCmd.PersistentFlags().IntVarP(&Align, "align", "a", 0, "align images at 2^<a> bytes")
	/*
	   split := strings.Split(PointersStr,",")
	   for i, v := range split {
	     a, err := strconv.Atoi(v)
	     if err != nil { os.Exit(1); }
	     Pointers[i] = a
	   }
	   if len(split)>3 {
	     fmt.Printf("WARNING: 'pointers' contains too many indexes. Since iCE40 devices support up to four, only %d are used.\n", Pointers)
	   }
	*/
}
