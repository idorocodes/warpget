/*
Copyright © 2026 JOHN AMOS idoroyen33@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/idorocodes/warpget/internal"
	"github.com/idorocodes/warpget/pkg"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download [url] [chunks]",
	Short: "Downloads a file",
	Long:  `Download initiates a high-performance concurrent download...`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		targetURL := args[0]
		numOfChunks := args[1]
		
		// 1. Get metadata
		size, supportsRange, err := internal.FileInfo(targetURL)
		if err != nil {
			return fmt.Errorf("failed to get file info: %w", err)
		}

	 
		outputFile := pkg.GetFileNameFromURL(targetURL)
		
		if !supportsRange {
			fmt.Println("⚠️ Server does not support range requests. Falling back to sequential.")
		} else {
			fmt.Printf("🚀 Starting concurrent download. Size: %d bytes\n", size)
		}

	 
		err = internal.DownloadFile(targetURL, numOfChunks, outputFile)
		if err != nil {
			return fmt.Errorf("download failed: %w", err)
		}

		fmt.Println("\n✅ Download successful!")
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
