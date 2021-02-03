package kafka

import "github.com/Shopify/sarama"

/*
func NewConsumer() (sarama.Consumer, error){
	c, err := sarama.NewConsumer()
	if err != nil {
		return nil, err
	}
	return c, err
}
*/
func NewConsumerGroup(addrs []string, groupID string, config *sarama.Config)(sarama.ConsumerGroup, error) {
	g, err := sarama.NewConsumerGroup(addrs, groupID, config)
	if err != nil {
		return nil, err
	}
	return g, err
}
