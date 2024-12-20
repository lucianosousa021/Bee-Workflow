package usecase

import (
	"zapi/model"
)

func (uc *Usecase) GetChats(userInstance, userToken, accountToken string) ([]model.ChatResponse, error) {
	var chatResponses []model.ChatResponse
	chats, err := uc.Repo.GetChat(userInstance, userToken, accountToken)
	if err != nil {
		return []model.ChatResponse{}, err
	}

	for _, chat := range chats {
		metadata, err := uc.Repo.GetChatMetadata(chat.Phone, userInstance, userToken, accountToken)
		if err != nil {
			return []model.ChatResponse{}, err
		}

		chatResponse := model.ChatResponse{
			Archived:         chat.Archived,
			Pinned:           chat.Pinned,
			MessagesUnread:   chat.MessagesUnread,
			Phone:            chat.Phone,
			Unread:           chat.Unread,
			Name:             chat.Name,
			LastMessageTime:  chat.LastMessageTime,
			IsMuted:          chat.IsMuted,
			IsMarkedSpam:     chat.IsMarkedSpam,
			ProfileThumbnail: metadata.ProfileThumbnail,
			About:            metadata.About,
		}

		chatResponses = append(chatResponses, chatResponse)
	}

	return chatResponses, nil
}
