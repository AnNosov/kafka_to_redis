package usecase

import "ktd/internal/entity"

type (
	// RedisProfile -
	RedisProfileI interface {
		PutProfiles(chan entity.Profile)
		putProfileKeys(string) error
		putProfileHash(entity.Profile) error
		deleteKeyProfile(string) error
	}

	// KafkaProfile -
	KafkaProfileI interface {
		GetProfile(chan entity.Profile)
	}

	// Profile -
	Profile interface {
		TransportData()
	}
)
