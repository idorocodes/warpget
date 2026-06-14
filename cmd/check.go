 package cmd
 
 import (
	"fmt"
 	"github.com/spf13/cobra"
	"github.com/idorocodes/warpget/internal" 
 )
 
 // checkCmd represents the check command
 var checkCmd = &cobra.Command{
	Use:   "check [url]",
	Short: "Verify if a URL supports concurrent downloading",
	Long:  `Check makes an HTTP HEAD request to the target URL to inspect its headers and verify if parallel chunked downloading is supported.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		targetURL := args[0]
		fmt.Printf("🔍 Inspecting URL: %s\n\n", targetURL)
 
		
		size, supportsRange, err := internal.FileInfo(targetURL)
		if err != nil {
			return err
		}
 
		
		fmt.Printf("✅ File Size: %d bytes\n", size)
		fmt.Printf("🚀 Supports Concurrent Download: %v\n", supportsRange)
		
		return nil
	},
 }
 
 func init() {
	rootCmd.AddCommand(checkCmd)
 }