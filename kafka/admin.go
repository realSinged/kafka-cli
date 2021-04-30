package kafka

import "github.com/Shopify/sarama"

func NewAdmin(addr []string, config *sarama.Config) (sarama.ClusterAdmin, error) {
	//config.Net.SASL.Enable = true
	//config.Net.SASL.User = "admin"
	//config.Net.SASL.Password = ""
	a, err := sarama.NewClusterAdmin(addr, config)
	return a, err
}
