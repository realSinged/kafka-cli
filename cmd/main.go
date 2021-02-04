package main

import (
	"github.com/realSinged/kafka-cli/cmd/admin"
	"github.com/realSinged/kafka-cli/cmd/consumer"
	"github.com/realSinged/kafka-cli/cmd/topic"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"time"
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
		Long: "a command line tools for apache kafka, include topic,consumer,producer, admin's operations",
		Run: runHelp,
	}
	cmds.AddCommand(consumer.NewCmdConsume())
	cmds.AddCommand(topic.NewCmdTopic())
	cmds.AddCommand(admin.NewCmdAdmin())
	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()

}
