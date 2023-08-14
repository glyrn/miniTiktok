package pojo

type Comment struct {
	Id           int64  `json:"id,omitempty"`
	User_service User   `json:"user,omitempty"`
	Content      string `json:"content"`
	CreateData   string `json:"create_data"`
}
