package topic

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/realSinged/kafka-cli/kafka"
	"github.com/realSinged/kafka-cli/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

var (
	topicExample = ``
)

func NewCmdTopic() *cobra.Command {
	cmd := &cobra.Command{
		Use: "topic",
		Short: "topic create/list/delete/detail etc",
		Long: "",
		Example: topicExample,
		Run: func(cmd *cobra.Command, args []string) {
			config := sarama.NewConfig()
			admin, err := kafka.NewAdmin([]string{"localhost:9092"}, config)
			if err != nil {
				log.Error("err", zap.Error(err))
				os.Exit(1)
			}
			topics, err := admin.ListTopics()
			if err != nil {
				log.Error("err", zap.Error(err))
				os.Exit(1)
			}
			for k, v := range topics {
				s, _ :=  json.Marshal(v)
				fmt.Printf("%s: %v\n", k, string(s))
			}
		},
	}
	cmd.AddCommand(
		NewCmdTopicList(),
		NewCmdTopicCreate(),
		NewCmdTopicDescribe(),
		NewCmdTopicDelete(),
		NewCmdTopicPartitionCreate(),
		NewCmdTopicPartitionReassign(),
		NewCmdTopicPartitionReassignList(),
		NewCmdTopicRecordDelete(),
		)
	return cmd
}

func NewCmdTopicList() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
	}
	return cmd
}

func NewCmdTopicCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create",
	}
	return cmd
}

func NewCmdTopicDescribe() *cobra.Command {
	cmd := &cobra.Command{
		Use: "describe",
	}
	return cmd
}

func NewCmdTopicDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete",
	}
	return cmd
}

func NewCmdTopicPartitionCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use: "partitionCreate",
	}
	return cmd
}

func NewCmdTopicPartitionReassign() *cobra.Command {
	cmd := &cobra.Command{
		Use: "partitionReassign",
	}
	return cmd
}

func NewCmdTopicPartitionReassignList() *cobra.Command {
	cmd := &cobra.Command{
		Use: "partitionReassignList",
	}
	return cmd
}

func NewCmdTopicRecordDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use: "recordDelete",
	}
	return cmd
}











