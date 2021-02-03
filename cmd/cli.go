package main

import (
	"fmt"
	"github.com/realSinged/kafka-cli/cmd/consumer"
	"github.com/realSinged/kafka-cli/log"
	"math/rand"
	"os"
	"strings"
	"time"
	"github.com/spf13/cobra"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := NewKafkaCliCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}

}

func NewKafkaCliCommand() *cobra.Command{
	cmds := &cobra.Command{
		Use: "kafka-cli",
		Short: "a command line tools for apache kafka",
		Long: "",
	}
	logger := log.NewLogger()
	logger.Info("haha")

	cmdPrint := &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	cmds.AddCommand(cmdPrint)
	cmds.AddCommand(consumer.NewCmdConsume())
	return cmds
}
