package handler

import (
	"dous/biz/model/query"
	"strings"
)

func ParseToken(token string) int64 {
	res := strings.Split(token, "&")
	u := query.User

	id, _ := u.Where(u.Username.Eq(res[0])).First()

	return id.ID
}
