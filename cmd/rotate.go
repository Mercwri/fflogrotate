package cmd

import (
	"github.com/Mercwri/fflogrotate/rotate"
	"github.com/spf13/cobra"
)

var rotateCmd = &cobra.Command{
	Use:   "rotate",
	Short: "compress and delete logs and archives",
	Long:  "Rotates short-term logs into archives and deletes long-lived archives",
	RunE: func(cmd *cobra.Command, args []string) error {
		short_term, _ := cmd.Flags().GetInt("short-term")
		long_term, _ := cmd.Flags().GetInt("long-term")
		rotate.Rotate(short_term, long_term)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(rotateCmd)
	rotateCmd.PersistentFlags().Int("short-term", 24, "time in hours before archiving")
	rotateCmd.PersistentFlags().Int("long-term", 120, "time in hours before deleting an archive")
}
