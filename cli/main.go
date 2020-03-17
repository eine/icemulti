// Copyright © 2018 eine <@github>
//
// Unless required by applicable law or agreed to in writing, this
// software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.

package main

import (
	"runtime"

	"github.com/eine/icemulti/cli/cmd"
	_ "github.com/spf13/cobra/doc"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
	//doc.GenMarkdownTree(cmd.RootCmd, "./doc")
}

//./go-icemulti.exe pack -mb 0 -p 1,2,3,4 ../dynreconfig/bitstreams/img0*.bin
