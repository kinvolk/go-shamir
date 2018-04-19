package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/vault/shamir"
	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split an arbitrarily long secret from stdin",
	Run:   runSplit,
}

var (
	parts     int
	threshold int
)

func init() {
	RootCmd.AddCommand(splitCmd)

	splitCmd.Flags().IntVarP(&parts, "parts", "p", 5, "number of shares")
	splitCmd.Flags().IntVarP(&threshold, "threshold", "t", 3, "threshold of shares required to reconstruct the secret")
}

func runSplit(cmd *cobra.Command, args []string) {
	secretBuf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read stdin: %v\n", err)
		os.Exit(1)
	}
	byteParts, err := shamir.Split(secretBuf, parts, threshold)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to split secret: %v\n", err)
		os.Exit(1)
	}
	for _, bytePart := range byteParts {
		fmt.Printf("%x\n", bytePart)
	}
}
