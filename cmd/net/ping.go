/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping is subcommand of net pallet. It pings a host",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	pingCmd.Flags().StringP("host", "H", "", "Host to ping")
	if err := pingCmd.MarkFlagRequired("host"); err != nil {
		panic(err)
	}

	// add commands to pallet
	NetCmd.AddCommand(pingCmd)
}
