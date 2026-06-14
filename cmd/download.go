/*
Copyright © 2026 JOHN AMOS idoroyen33@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/idorocodes/warpget/internal"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download [url]",
	Short: "Downloads a file",
	Long: `Download initiates a high-performance concurrent download of the file 
provided via the URL. It splits the file into chunks and downloads them 
in parallel to maximize network throughput.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		targetURL := args[0]
	 
		size, supportsRange, err := internal.FileInfo(targetURL)
		if err != nil {
			return fmt.Errorf("failed to get file info: %w", err)
		}

		if !supportsRange {
			fmt.Println("⚠️ Server does not support range requests. Downloading sequentially...")
		} else {
			fmt.Printf("🚀 Starting concurrent download. Size: %d bytes\n", size)
		}

	 

		return nil
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
