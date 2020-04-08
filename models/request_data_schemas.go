package models

type RegisterUserRequestSchema struct {
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LoginUserRequestSchema struct {
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type CreateChatRequestSchema struct {
	Title   string `json:"title"`
	AdminId uint   `json:"admin_id"`
	Members []uint `json:"members"`
}

type AddUserToChatRequestSchema struct {
	UserId uint `json:"user_id"`
	ChatId uint `json:"chat_id"`
}

type SendMessageRequestSchema struct {
	ChatId uint   `json:"chat_id"`
	Text   string `json:"text"`
}

type ListChatMembersRequestSchema struct {
	ChatId uint `json:"chat_id"`
}