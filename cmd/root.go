package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fflogrotate.exe rotate",
	Short: "Manages the lifecycle of FFLogfiles",
	Long:  `Rotate, Archive, Delete, and Backup Logfiles from ACT and IINACT`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}
