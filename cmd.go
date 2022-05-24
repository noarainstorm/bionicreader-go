package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func start() {
	tk := new(bionic)
	rootCmd := &cobra.Command{
		Use:   "bionical",
		Short: "Utility to make bionic text",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				log.Fatal("To few arguments...")
			}
			for _, em := range args {
				tk.get(em)
				fmt.Println(tk.Out)
			}
		},
	}
	rootCmd.Execute()
}
