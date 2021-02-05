package utils

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
)

func PrintTopic(topic string, detail sarama.TopicDetail) {
	s, _ := json.MarshalIndent(detail, "", "")
	printSeparator()
	fmt.Printf("TOPIC:%s\nDETAIL:%s\n", topic, string(s))
}

func PrintTopicMeta(meta *sarama.TopicMetadata) {
	printSeparator()
	s, _ := json.MarshalIndent(meta, "", " ")
	fmt.Printf("TOPIC:%s\nDETAIL:%10s\n", meta.Name, string(s))
}

func PrintConsumerMessage(msg *sarama.ConsumerMessage) {
	printSeparator()
	s, _ := json.MarshalIndent(msg, "", " ")
	fmt.Printf(string(s))
}

func printSeparator() {
	fmt.Println("****************************************")
}
