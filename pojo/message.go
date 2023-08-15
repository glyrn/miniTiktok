package pojo

import "miniTiktok/entity"

type Message struct {
	entity.Message
	UserName string `json:"user_name"`
}
