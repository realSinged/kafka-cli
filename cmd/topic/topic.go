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
	"strings"
)

var (
	topicExample = ``
)

type topicOptions struct {
	bootstrapServers string
	list bool
	describe string
	create string
	delete string
	addPartition string
	numPartition int32 //创建topic时指定的partition
	numReplica int16 //创建topic时指定的副本数
}

func newTopicOptions() *topicOptions{
	return &topicOptions{}
}

func (o *topicOptions) run(cmd *cobra.Command, args []string) {
	config := sarama.NewConfig()
	servers := strings.Split(o.bootstrapServers, ",")
	admin, err := kafka.NewAdmin(servers, config)
	if err != nil {
		log.Error("err", zap.Error(err))
		os.Exit(1)
	}
	if o.list {
		topics, err := admin.ListTopics()
		if err != nil {
			log.Error("err", zap.Error(err))
			os.Exit(1)
		}
		for k, v := range topics {
			s, _ :=  json.MarshalIndent(v, "", "")
			fmt.Printf("%s: %v\n", k, string(s))
		}
	} else if o.describe != "" {
		topics, err := admin.DescribeTopics(strings.Split(o.describe, ","))
		if err != nil {
			log.Error("err", zap.Error(err))
			os.Exit(1)
		}
		for k,v := range topics {
			s, _ :=  json.MarshalIndent(v, "", " ")
			fmt.Printf("%s: %v\n", k, string(s))
		}
	}else if o.create != "" {
		err := admin.CreateTopic(o.create, &sarama.TopicDetail{NumPartitions: o.numPartition, ReplicationFactor: o.numReplica}, false)
		if err != nil{
			log.Error("err", zap.Error(err))
			os.Exit(1)
		}
	} else if o.delete  != ""{
		err := admin.DeleteTopic(o.delete)
		if err != nil {
			log.Error("err", zap.Error(err))
			os.Exit(1)
		}
		log.Info("topic delete success", zap.String("topic", o.delete))
	} else if o.addPartition != "" {
		err := admin.CreatePartitions(o.addPartition, o.numPartition, [][]int32{}, false)
		if err != nil {
			log.Error("err", zap.Error(err))
			os.Exit(1)
		}
	} else {
		cmd.Help()
	}
}

func NewCmdTopic() *cobra.Command {
	o := newTopicOptions()

	cmd := &cobra.Command{
		Use: "topic",
		Short: "Kafka topic operations",
		Long: "Topic operations, include topic create、list、delete、detail, topic partition create",
		Example: topicExample,
		Run: o.run,
	}
	cmd.Flags().StringVarP(&o.bootstrapServers, "bootstrap-server", "b", "localhost:9092", "The Kafka server to connect to.more than one should be separated by commas")
	cmd.Flags().BoolVarP(&o.list, "list", "l", o.list, "List all available topics.")
	cmd.Flags().StringVar(&o.describe, "describe", o.describe, "List details for the given topics.more than one should be separated by commas")
	cmd.Flags().StringVarP(&o.create, "create", "c", o.create, "Create a new topic.")
	cmd.Flags().Int32Var(&o.numPartition, "partition-num", 1, "The specified partition when create topic")
	cmd.Flags().Int16Var(&o.numReplica, "replica-num", 1, "The specified replica when create topic")
	cmd.Flags().StringVarP(&o.delete, "delete", "d", o.delete,"Delete a topic.")
	cmd.Flags().StringVar(&o.addPartition, "add-partition", o.addPartition, "The Topic which need to create partition, partition num must higher than which already exists")
	return cmd
}