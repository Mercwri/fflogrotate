package cmd

import (
	"github.com/Mercwri/fflogrotate/schedule"
	"github.com/spf13/cobra"
)

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "create a task to run on a schedule",
	Long:  "Rotates short-term logs into archives and deletes long-lived archives",
	RunE: func(cmd *cobra.Command, args []string) error {
		schedule.Schedule()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
}
