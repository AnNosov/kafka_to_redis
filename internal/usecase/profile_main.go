package usecase

import "ktd/internal/entity"

type PrflUseCase struct {
	KafkaProfile
	RedisProfile
}

func New(k KafkaProfile, r RedisProfile) *PrflUseCase {
	return &PrflUseCase{
		KafkaProfile: k,
		RedisProfile: r,
	}
}

func (uc *PrflUseCase) TransportData() {
	c := make(chan entity.Profile, 10)
	for {
		go uc.KafkaProfile.GetProfile(c)
		go uc.RedisProfile.PutProfiles(c)
	}
}
