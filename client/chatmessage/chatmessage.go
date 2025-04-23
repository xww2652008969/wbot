package chatmessage

import (
	"strconv"
)

const (
	MsgTypeText   = "text"
	MsgTypeAt     = "at"
	MsgTypeImage  = "image"
	MsgTypeFace   = "face"
	MsgTypeRecord = "record"
	MsgTypeVideo  = "video"
	MsgTypeReply  = "reply"
	MsgTypeMusic  = "music"
	MsgTypeDice   = "dice"
	MsgTypeFile   = "file"
)

func (m *ChatMessage) addMessage(msgType string, data messagePayload) *ChatMessage {
	m.Message = append(m.Message, ChatMessageData{
		Type: msgType,
		Data: data,
	})
	return m
}

// 用于处理发送消息

// AddImage 添加图片
func (m *ChatMessage) AddImage(url string) *ChatMessage {
	return m.addMessage(MsgTypeImage, messagePayload{File: url})
}

// AddFace 添加系统表情
func (m *ChatMessage) AddFace(id int) *ChatMessage {
	return m.addMessage(MsgTypeFace, messagePayload{Id: strconv.Itoa(id)})
}

// AddRecord 添加语音
func (m *ChatMessage) AddRecord(url string) *ChatMessage {
	return m.addMessage(MsgTypeRecord, messagePayload{File: url})
}

// AddVideo
func (m *ChatMessage) AddVideo(url string) *ChatMessage {
	return m.addMessage(MsgTypeVideo, messagePayload{File: url})
}

// Addreply 添加回复消息的信息 传入Messageid
func (m *ChatMessage) Addreply(id int64) *ChatMessage {
	return m.addMessage(MsgTypeReply, messagePayload{Id: strconv.FormatInt(id, 10)})
}

// AddMusicCard 添加音乐卡片（不知道如何获取id和只能用qq？？？）
func (m *ChatMessage) AddMusicCard(t string, id int) *ChatMessage {
	return m.addMessage(MsgTypeMusic, messagePayload{Type: t, Id: strconv.Itoa(id)})
}

// AddDice  发送骰子表情(??直接face就可以了的？？)
func (m *ChatMessage) AddDice() *ChatMessage {
	return m.addMessage(MsgTypeDice, messagePayload{Type: MsgTypeDice})
}
func (m *ChatMessage) AddFile(url string) *ChatMessage {
	return m.addMessage(MsgTypeFile, messagePayload{File: url})
}

// AddText 添加string文本
func (m *ChatMessage) AddText(text string) *ChatMessage {
	return m.addMessage(MsgTypeText, messagePayload{Text: text})
}

// Sendat 设置@某人  传入qq号 如果为0则@全体
func (m *ChatMessage) AddAt(qq int64) *ChatMessage {
	qqStr := "all"
	if qq != 0 {
		qqStr = strconv.FormatInt(qq, 10)
	}
	return m.addMessage(MsgTypeAt, messagePayload{Qq: qqStr})
}
