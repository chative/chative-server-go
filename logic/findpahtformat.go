package logic

import (
	"strconv"
	"strings"
	"time"

	"chative-server-go/models"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/language"
)

// src -> dst , dst 查看 src怎么找到的dst
func FindPathFormat(ctx *Context, src, dst, acceptLang, sourceQueryType string) (string, error) {
	// 确定语言
	var lang = "en"
	tags, _, _ := language.ParseAcceptLanguage(acceptLang)
	for _, tag := range tags {
		tagStr := strings.ToLower(tag.String())
		if strings.HasPrefix(tagStr, "en") {
			lang = "en"
			break
		}
		if strings.Contains(tagStr, "zh") || strings.Contains(tagStr, "cn") {
			lang = "zh"
			break
		}
	}
	// 先用 src -> dst
	// 再用 dst -> src, 如果更早的话
	var path1, path2 models.UserFindPath
	ctx.db.First(&path1, "src = ? and dst = ?", src, dst)
	ctx.db.First(&path2, "src = ? and dst = ?", dst, src)
	var path models.UserFindPath
	var describe string
	if path1.ID != 0 { // The person found you
		path = path1
		var duStr = formatTimeDuration(time.Since(path.UpdatedAt), lang, sourceQueryType)
		switch path.Path.Type {
		case "fromGroup":
			gname, err := GetGroupName(ctx, path.Path.GroupID)
			if err != nil || gname == "" {
				ctx.Logger.Errorw("GetGroupName failed", logx.Field("err", err))
				if lang == "en" {
					describe = "The person found you via a group" + duStr + "."
				} else {
					describe = "他通过一个群找到您。" + duStr
				}
			} else {
				gname, _ = shortName(gname)
				if lang == "en" {
					describe = "The person found you via the group “" + gname + "“" + duStr + "."
				} else {
					describe = "他通过群“" + gname + "“找到您。" + duStr
				}
			}
		case "shareContact":
			userInfo, _ := GetUserBasicInfo(ctx.redisCmd, ctx.db, path.Path.UID)
			userName, _ := shortName(userInfo.PlainName)
			if lang == "en" {
				describe = `The person found you via the contact shared by "` + userName + `"` + duStr + "."
			} else {
				describe = `他通过"` + userName + `"分享的名片找到您。` + duStr
			}
		case "link":
			if lang == "en" {
				describe = "The person found you via your invite link or QR code" + duStr + "."
			} else {
				describe = "他通过您的邀请链接或二维码找到您。" + duStr
			}
		case "search":
			if lang == "en" {
				describe = "The person found you via searching your email/phone number" + duStr + "."
			} else {
				describe = "他通过搜索你的邮箱/手机号找到您。" + duStr
			}
		}
		if sourceQueryType != "met" {
			return describe, nil
		}
	}
	if path1.ID == 0 || path2.ID != 0 && path2.UpdatedAt.Before(path.UpdatedAt) { // You found the person
		path = path2
		var duStr = formatTimeDuration(time.Since(path.UpdatedAt), lang, sourceQueryType)
		switch path.Path.Type {
		case "fromGroup":
			gname, err := GetGroupName(ctx, path.Path.GroupID)
			if err != nil || gname == "" {
				ctx.Logger.Errorw("GetGroupName failed", logx.Field("err", err))
				gname = "Unknown"
			}
			gname, _ = shortName(gname)
			if lang == "en" {
				return "You found the person via the group “" + gname + "“" + duStr + ".", nil
			} else {
				return "你通过群“" + gname + "“找到他" + duStr, nil
			}
		case "shareContact":
			userInfo, _ := GetUserBasicInfo(ctx.redisCmd, ctx.db, path.Path.UID)
			userName, _ := shortName(userInfo.PlainName)
			if lang == "en" {
				return `You found the person via the contact shared by "` + userName + `"` + duStr + ".", nil
			} else {
				return `你通过"` + userName + `"分享的名片找到他` + duStr, nil
			}
		case "link":
			if lang == "en" {
				return "You found the person via invite link or QR code" + duStr + ".", nil
			} else {
				return "你通过邀请链接或二维码找到他" + duStr, nil
			}
		case "search":
			if lang == "en" {
				return "You found the person via searching email/phone number" + duStr + ".", nil
			} else {
				return "你通过搜索邮箱/手机找到他" + duStr, nil
			}
		}
	}

	if describe != "" {
		return describe, nil
	}
	// 先查找 src (inviter) -> dst
	var duStr string
	var youFoundThey = false
	var askRequest models.AskNewFriend
	err := ctx.db.Order("id").First(&askRequest,
		models.AskNewFriend{Inviter: src, Invitee: dst}).Error
	if err != nil || askRequest.ID == 0 { // 再查找 dst (inviter) -> src
		err = ctx.db.Order("id").First(&askRequest, models.AskNewFriend{Inviter: dst, Invitee: src}).Error
		youFoundThey = true
	}
	if err != nil || askRequest.ID == 0 {
		ctx.Logger.Errorw("First AskNewFriend failed", logx.Field("err", err), logx.Field("src", src), logx.Field("dst", dst))
		// 查好友关系
		var friendRelation = &models.FriendRelation{UserID1: src, UserID2: dst}
		if friendRelation.UserID1 > friendRelation.UserID2 {
			friendRelation.UserID1, friendRelation.UserID2 = friendRelation.UserID2, friendRelation.UserID1
		}
		err = ctx.db.First(friendRelation, friendRelation).Error
		if err != nil {
			return "", nil
		}
		youFoundThey = false
		duStr = formatTimeDuration(time.Since(friendRelation.UpdatedAt), lang, sourceQueryType)
	} else {
		duStr = formatTimeDuration(time.Since(askRequest.UpdatedAt), lang, sourceQueryType)
	}
	if youFoundThey {
		if lang == "en" {
			return "You added the person" + duStr + ".", nil
		} else {
			return "你添加他为好友" + duStr, nil
		}
	} else {
		if lang == "en" {
			return "The person added you" + duStr + ".", nil
		} else {
			return "他添加你为好友" + duStr, nil
		}
	}
}

func formatTimeDuration(du time.Duration, lang, sourceQueryType string) (duStr string) { //
	if sourceQueryType != "met" {
		return ""
	}
	// if du < time.Minute {
	// 	if lang == "en" {
	// 		duStr = " (just now)"
	// 	} else {
	// 		duStr = "（刚刚）"
	// 	}
	// 	return
	// }
	if du <= time.Hour*24 {
		if lang == "en" {
			duStr = " (today)"
		} else {
			duStr = "（今天）"
		}
		return
	}

	// if du < time.Hour {
	// 	if lang == "en" {
	// 		if du > time.Minute {
	// 			duStr = " (" + strconv.FormatInt(int64(du/time.Minute), 10) + " minutes ago)"
	// 		} else {
	// 			duStr = " (1 minute ago)"
	// 		}
	// 	} else {
	// 		duStr = "（" + strconv.FormatInt(int64(du/time.Minute), 10) + "分钟前）"
	// 	}
	// 	return
	// }

	// if du < time.Hour*24 {
	// 	if lang == "en" {
	// 		if du > time.Hour {
	// 			duStr = " (" + strconv.FormatInt(int64(du/time.Hour), 10) + " hours ago)"
	// 		} else {
	// 			duStr = " (1 hour ago)"
	// 		}
	// 	} else {
	// 		duStr = "（" + strconv.FormatInt(int64(du/time.Hour), 10) + "天前）"
	// 	}
	// 	return
	// }

	if lang == "en" {
		if du >= 2*time.Hour*24 {
			duStr = " (" + strconv.FormatInt(int64(du/(time.Hour*24)), 10) + " days ago)"
		} else {
			duStr = " (1 day ago)"
		}
	} else {
		duStr = "（" + strconv.FormatInt(int64(du/(time.Hour*24)), 10) + "天前）"
	}

	return
}

func shortName(name string) (newName string, cut bool) {
	runeName := []rune(name)
	if condition := len(runeName) > 20; condition {
		newName = string(runeName[0:20]) + ".."
		cut = true
	} else {
		newName = name
		cut = false
	}
	return
}
