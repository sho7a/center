package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func ExecuteCenter(cmd *cobra.Command, args []string) {
	fileInfo, _ := os.Stdin.Stat()
	piped := fileInfo.Mode()&os.ModeCharDevice == 0

	if piped {
		scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
		for scanner.Scan() {
			PrintCenter(scanner.Text())
		}
	}

	for _, s := range args {
		PrintCenter(s)
	}

	if !piped && len(args) == 0 {
		_ = cmd.Help()
	}
}

func PrintCenter(c string) {
	scanner := bufio.NewScanner(strings.NewReader(c))
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		width, _ := GetTerminalSize()
		fmt.Println(fmt.Sprintf("%*s", -width, fmt.Sprintf("%*s", (width+len(s))/2, s)))
	}
}
