package rediscluster

import (
	"time"

	"github.com/go-redis/redis_rate/v9"
)

// func Check()  {

// }

var (
	AskFriendSingleLimit = redis_rate.Limit{
		Rate:   999,
		Period: time.Hour * 24,
		Burst:  999,
	}
	AskFriendMutualLimit = redis_rate.Limit{
		Rate:   999,
		Period: time.Hour * 24 * 7,
		Burst:  999,
	}

	VerificationCodeLimit = redis_rate.Limit{
		Rate:   3,
		Period: time.Minute,
		Burst:  3,
	}
)
