package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func ExecuteCenter(cmd *cobra.Command, args []string) {
	fileInfo, _ := os.Stdin.Stat()
	piped := fileInfo.Mode()&os.ModeCharDevice == 0

	if piped {
		scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
		for scanner.Scan() {
			printCenter(scanner.Text())
		}
	}

	for _, s := range args {
		printCenter(s)
	}

	if !piped && len(args) == 0 {
		_ = cmd.Help()
	}
}

func printCenter(c string) {
	c = strings.ReplaceAll(c, "\r\n", "\n")
	c = strings.ReplaceAll(c, "\r", "\n")
	a := strings.Split(c, `\n`)
	for _, s := range a {
		s = strings.TrimSpace(s)
		w, _ := GetTerminalSize()
		fmt.Println(fmt.Sprintf("%*s", -w, fmt.Sprintf("%*s", (w+len(s))/2, s)))
	}
}
