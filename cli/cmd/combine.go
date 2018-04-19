package cmd

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/hashicorp/vault/shamir"
	"github.com/spf13/cobra"
)

var combineCmd = &cobra.Command{
	Use:   "combine",
	Short: "Reconstruct a secret from the parts read from stdin",
	Run:   runCombine,
}

func init() {
	RootCmd.AddCommand(combineCmd)
}

func runCombine(cmd *cobra.Command, args []string) {
	var hexParts []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		hexParts = append(hexParts, scanner.Text())
	}
	var byteParts [][]byte
	for _, hexPart := range hexParts {
		b, err := hex.DecodeString(hexPart)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to decode %q: %v\n", hexPart, err)
			os.Exit(1)
		}
		byteParts = append(byteParts, b)
	}
	secretBytes, err := shamir.Combine(byteParts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to combine secret: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(secretBytes))
}
