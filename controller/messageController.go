package controller

import "miniTiktok/pojo"

type MessageResponse struct {
	Response
	MessageList []pojo.Comment `json:"comment_list,omitempty"`
}
