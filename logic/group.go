package logic

import (
	"encoding/json"

	"chative-server-go/models"
)

func GetGroupName(ctx *Context, gid string) (string, error) {
	jsonString, err := ctx.redisCmd.Get(ctx.ctx, "Group_Group_1_"+gid).Result()
	if err != nil || jsonString == "" {
		// 数据库中取
		var group models.Group
		err = ctx.db.Where("id = ?", gid).First(&group).Error

		return group.Name, err
	}
	var groupInfo struct {
		Name string `json:"name"`
	}
	err = json.Unmarshal([]byte(jsonString), &groupInfo)
	if err != nil {
		return "", err
	}
	return groupInfo.Name, nil
}
