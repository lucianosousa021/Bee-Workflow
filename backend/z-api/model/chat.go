package model

type ChatResponse struct {
	Archived         string `json:"archived"`
	Pinned           string `json:"pinned"`
	MessagesUnread   int    `json:"messagesUnread"`
	Phone            string `json:"phone"`
	Unread           string `json:"unread"`
	Name             string `json:"name"`
	LastMessageTime  string `json:"lastMessageTime"`
	IsMuted          string `json:"isMuted"`
	IsMarkedSpam     string `json:"isMarkedSpam"`
	ProfileThumbnail string `json:"profileThumbnail"`
	About            string `json:"about"`
}

type Chat struct {
	Archived        string `json:"archived"`
	Pinned          string `json:"pinned"`
	MessagesUnread  int    `json:"messagesUnread"`
	Phone           string `json:"phone"`
	Unread          string `json:"unread"`
	Name            string `json:"name"`
	LastMessageTime string `json:"lastMessageTime"`
	MuteEndTime     int64  `json:"muteEndTime,omitempty"`
	IsMuted         string `json:"isMuted"`
	IsMarkedSpam    string `json:"isMarkedSpam"`
}

type ChatMetadata struct {
	Phone            string `json:"phone"`
	Unread           string `json:"unread"`
	LastMessageTime  string `json:"lastMessageTime"`
	IsMuted          string `json:"isMuted"`
	IsMarkedSpam     string `json:"isMarkedSpam"`
	ProfileThumbnail string `json:"profileThumbnail"`
	MessagesUnread   int    `json:"messagesUnread"`
	About            string `json:"about"`
}

type ModifyChatStatus struct {
	Phone  string `json:"phone"`
	Action string `json:"action"`
}
