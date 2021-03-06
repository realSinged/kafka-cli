package consumer

import (
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/realSinged/kafka-cli/kafka"
	"github.com/realSinged/kafka-cli/utils"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	consumergExample = `
# Consume kafka messages of set of topic which certain group id, and will print those message in terminal
	kafka-cli consume -t test,singed -c default -b localhost:9092
		`
)

type consumerGOptions struct {
	bootstrapServers string
	groupID          string
	topics           string
	user, password   string
}

func newConsumerGOptions() *consumerGOptions {
	return &consumerGOptions{}
}

func (o *consumerGOptions) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (o *consumerGOptions) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (o *consumerGOptions) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		utils.PrintConsumerMessage(msg)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func (o *consumerGOptions) run(cmd *cobra.Command, args []string) {
	if o.topics != "" {
		config := sarama.NewConfig()
		//config.Consumer.Offsets.Initial = sarama.OffsetOldest
		config.Consumer.Offsets.Initial = sarama.OffsetNewest
		if o.user != "" && o.password != "" {
			config.Net.SASL.Enable = true
			config.Net.SASL.User = o.user
			config.Net.SASL.Password = o.password
		}

		fmt.Println(strings.Split(o.bootstrapServers, ","))
		c, err := kafka.NewConsumerGroup(strings.Split(o.bootstrapServers, ","), o.groupID, config)
		utils.CheckErr(err)
		defer func() {
			utils.CheckErr(c.Close())
		}()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err = c.Consume(ctx, strings.Split(o.topics, ","), o)
		utils.CheckErr(err)
	} else {
		cmd.Help()
	}
}

func NewCmdConsumeGroup() *cobra.Command {
	o := newConsumerGOptions()
	cmd := &cobra.Command{
		Use:     "consumerg",
		Short:   "Consume kafka message with given topics and group_id",
		Long:    "Consume kafka message with given topics and group_id, and message will be auto committed",
		Example: consumergExample,
		Run:     o.run,
	}

	cmd.Flags().StringVarP(&o.bootstrapServers, "bootstrap-servers", "b", "localhost:9092", "The Kafka server to connect to.more than one should be separated by commas")
	cmd.Flags().StringVar(&o.topics, "topics", o.topics, "The topics to consume,more than one should be separated by commas")
	cmd.Flags().StringVar(&o.groupID, "group-id", "kafka-cli", "The consumer group ID")
	cmd.Flags().StringVar(&o.user, "user", o.user, "auth user, if miss means no auth")
	cmd.Flags().StringVar(&o.password, "password", o.password, "auth password, if miss means no auth")
	return cmd
}
