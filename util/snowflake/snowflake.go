package snowflake

import (
	"errors"
	"github.com/sony/sonyflake"
)

var snowFlakeInstance *sonyflake.Sonyflake

func InitSnowFlake() (err error) {
	settings := sonyflake.Settings{}
	snowFlakeInstance = sonyflake.NewSonyflake(settings)
	if snowFlakeInstance == nil {
		err := errors.New("生成snowflake实例失败，请重试")
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateID() (id uint64, err error) {
	id, err = snowFlakeInstance.NextID()
	return
}
