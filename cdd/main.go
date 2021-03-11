package main

import (
	"time"

	"github.com/herryg91/cdd/cdd/commands/doctor"
	"github.com/herryg91/cdd/cdd/commands/gen"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "cdd", Short: "cdd", Long: "cdd"}
	rootCmd.AddCommand(gen.NewGenCmd().Command)
	rootCmd.AddCommand(doctor.NewDoctorCmd().Command)
	// execute
	rootCmd.Execute()

	time.Sleep(time.Second)
}
