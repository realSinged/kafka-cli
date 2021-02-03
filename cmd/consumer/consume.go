package consumer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/realSinged/kafka-cli/kafka"
	"github.com/realSinged/kafka-cli/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"os"
)

var (
	consumeExample = `
        # Consume kafka messages of topic test and consumer default, and will print those message in terminal
        kafka-cli consume -t test -c default -b localhost:9092
		`
)
func NewCmdConsume() *cobra.Command {

	cmd := &cobra.Command{
		Use: "consume",
		Short: "Consume kafka message with given topic and group_id",
		Long: "",
		Example: consumeExample,
		Run: func(cmd *cobra.Command, args []string) {
			config := sarama.NewConfig()
			config.Consumer.Offsets.Initial = sarama.OffsetOldest

			c, err := kafka.NewConsumerGroup([]string{"127.0.0.1:9092"}, "default", config)
			if err != nil {
				log.Error("new consumer group failed", zap.Error(err))
				os.Exit(1)
			}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			log.Info("going to consume")
			err = c.Consume(ctx,[]string{"event"}, handler{})
			if err != nil {
				log.Error("consume failed", zap.Error(err))

			}
		},
	}
	return cmd
}

type handler struct{}

func (h handler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h handler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h handler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}
