package main

import (
    "runtime"
    "fmt"
    "os"

    "./cmd"
    _"github.com/spf13/cobra/doc"
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    if err := cmd.RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }

  	//doc.GenMarkdownTree(cmd.RootCmd, "./doc")
}

//./go-icemulti.exe pack -mb 0 -p 1,2,3,4 ../dynreconfig/bitstreams/img0*.bin
