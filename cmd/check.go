package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check [url]",
	Short: "Verify if a URL supports concurrent downloading",
	Long:  `Check makes an HTTP HEAD request to the target URL to inspect its headers and verify if parallel chunked downloading is supported.`,
	// Force the CLI to require exactly 1 positional argument (the URL)
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		targetURL := args[0]
		fmt.Printf("🔍 Inspecting URL: %s\n\n", targetURL)

		// Make a lightweight HEAD request to read metadata headers without downloading the file
		resp, err := http.Head(targetURL)
		if err != nil {
			return fmt.Errorf("failed to connect to server: %w", err)
		}
		defer resp.Body.Close()

		// Print server connection status
		fmt.Printf("📡 Status:          %s\n", resp.Status)

		// Print file size if available
		if resp.ContentLength > 0 {
			fmt.Printf("📦 Content Length:  %d bytes (%.2f MB)\n", resp.ContentLength, float64(resp.ContentLength)/(1024*1024))
		} else {
			fmt.Println("📦 Content Length:  Unknown")
		}

		// Check the critical Accept-Ranges header for concurrent download capacity
		acceptRanges := resp.Header.Get("Accept-Ranges")
		if acceptRanges == "bytes" {
			fmt.Println("⚡ Concurrency:     Supported! (Server accepts partial byte ranges)")
		} else {
			fmt.Println("⚠️ Concurrency:     Not supported. (Downloads will fall back to a single thread)")
		}

		return nil
	},
}

func init() {
	// Register this subcommand to your main rootCmd
	rootCmd.AddCommand(checkCmd)
}
