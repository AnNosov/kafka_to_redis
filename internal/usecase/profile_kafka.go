package usecase

import (
	"context"
	"encoding/json"
	"ktd/internal/entity"
	"ktd/pkg/kfk"
	"log"
)

type KafkaProfile struct {
	*kfk.KafkaReader
}

func NewKafkaReader(k *kfk.KafkaReader) *KafkaProfile {
	return &KafkaProfile{k}
}

func (kr *KafkaProfile) GetProfile(profileChan chan entity.Profile) {
	msg, err := kr.KafkaReader.KfkReader.ReadMessage(context.Background())
	if err != nil {
		log.Println("Kafka GetProfile: ", err)
		return
	}
	var profile entity.Profile
	if err := json.Unmarshal(msg.Value, &profile); err != nil {
		log.Println("value: ", string(msg.Value), "error: ", err)
		return
	}
	profileChan <- profile
}
