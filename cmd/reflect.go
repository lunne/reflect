package main

import "github.com/spf13/cobra"

func init() {
	Root.AddCommand(todayCmd)

}

// ...
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "what have you planned for today?",
}
