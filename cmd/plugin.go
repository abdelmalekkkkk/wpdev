/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pluginCmd represents the plugin command
var pluginCmd = &cobra.Command{
	Use:  "plugin",
	Long: `The plugin utility allows you to create new plugins with a standard and unified structure. Please use wpdev plugin -h to see what you can do.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("plugin called")
	},
}

func init() {
	rootCmd.AddCommand(pluginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pluginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pluginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
