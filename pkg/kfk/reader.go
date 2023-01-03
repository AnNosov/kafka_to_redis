package kfk

import (
	"ktd/config"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	KfkReader *kafka.Reader
}

func newConfig(cfg config.Kafka) *kafka.ReaderConfig {
	var brokers []string
	brokers[0] = cfg.Host + ":" + cfg.Port
	return &kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   cfg.Topic,
	}
}

func NewReader(cfg config.Kafka) *KafkaReader {
	r := kafka.NewReader(*newConfig(cfg))
	r.SetOffset(cfg.Offset)

	k := &KafkaReader{
		KfkReader: r,
	}
	return k
}

func (r *KafkaReader) Close() error {
	if err := r.KfkReader.Close(); err != nil {
		log.Println("close kafka connection: ", err)
		return err
	}
	return nil
}
