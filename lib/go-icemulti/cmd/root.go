package cmd

import (
  "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:   "go-icemulti",
    Short: "go-icemulti is an extended version of icestorm/icemulti",
    Long: `A multi-bitstream packager for iCE40 devices. This is a golang-based version of icestorm/icemulti <github.com/cliffordwolf/icestorm>, extended to optimize memory footprint when packaging large sets.`,
    //Run: func(cmd *cobra.Command, args []string) {
        // Do Stuff Here
    //},
}

func init() {
//    cobra.OnInitialize(initConfig)
//    RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
//    RootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
//    RootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
//    RootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
//    RootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
//    viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
//    viper.BindPFlag("projectbase", RootCmd.PersistentFlags().Lookup("projectbase"))
//    viper.BindPFlag("useViper", RootCmd.PersistentFlags().Lookup("viper"))
//    viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
//    viper.SetDefault("license", "apache")
}
