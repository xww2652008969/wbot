package sendapi

import "github.com/xww2652008969/wbot/client/chatmessage"

type SendAPI struct {
	httpurl     string
	chatmessage *chatmessage.ChatMessage
}

func (sendapi *SendAPI) Gethttpurl() string {
	return sendapi.httpurl
}
func (sendapi *SendAPI) Getchatmessage() *chatmessage.ChatMessage {
	return sendapi.chatmessage
}
func (sendapi *SendAPI) SetChatmessage(chatmessage *chatmessage.ChatMessage) {
	sendapi.chatmessage = chatmessage
}
func (sendapi *SendAPI) SetHttpUrl(url string) {
	sendapi.httpurl = url
}

// 点赞数据处理 不导出
type sendlikedata struct {
	UserId string `json:"user_id"`
	Times  int    `json:"times"`
}

// 处理获得某群成员信息的列表
type reqgroupMemberInfo struct {
	GroupId int64 `json:"group_id"`
	UserId  int64 `json:"user_id"`
	NoCache bool  `json:"no_cache"`
}
type groupMemberInfo struct {
	Status  string `json:"status"`
	Retcode int    `json:"retcode"`
	Data    struct {
		GroupId         int    `json:"group_id"`
		UserId          int    `json:"user_id"`
		Nickname        string `json:"nickname"`
		Card            string `json:"card"`
		Sex             string `json:"sex"`
		Age             int    `json:"age"`
		Area            string `json:"area"`
		Level           string `json:"level"`
		QqLevel         int    `json:"qq_level"`
		JoinTime        int    `json:"join_time"`
		LastSentTime    int    `json:"last_sent_time"`
		TitleExpireTime int    `json:"title_expire_time"`
		Unfriendly      bool   `json:"unfriendly"`
		CardChangeable  bool   `json:"card_changeable"`
		IsRobot         bool   `json:"is_robot"`
		ShutUpTimestamp int    `json:"shut_up_timestamp"`
		Role            string `json:"role"`
		Title           string `json:"title"`
	} `json:"data"`
	Message string      `json:"message"`
	Wording string      `json:"wording"`
	Echo    interface{} `json:"echo"`
}

//type MsgInfo struct {
//	Status  string      `json:"status"`
//	Retcode int         `json:"retcode"`
//	Data    Message     `json:"data"`
//	Message string      `json:"message"`
//	Wording string      `json:"wording"`
//	Echo    interface{} `json:"echo"`
//}

// 处理获得群成员信息的列表
type GroupMemberList struct {
	Status  string                `json:"status"`
	Retcode int                   `json:"retcode"`
	Data    []GroupMemberListData `json:"data"`
	Message string                `json:"message"`
	Wording string                `json:"wording"`
	Echo    interface{}           `json:"echo"`
}
type GroupMemberListData struct {
	GroupId         int    `json:"group_id"`
	UserId          int64  `json:"user_id"`
	Nickname        string `json:"nickname"`
	Card            string `json:"card"`
	Sex             string `json:"sex"`
	Age             int    `json:"age"`
	Area            string `json:"area"`
	Level           string `json:"level"`
	QqLevel         int    `json:"qq_level"`
	JoinTime        int    `json:"join_time"`
	LastSentTime    int    `json:"last_sent_time"`
	TitleExpireTime int    `json:"title_expire_time"`
	Unfriendly      bool   `json:"unfriendly"`
	CardChangeable  bool   `json:"card_changeable"`
	IsRobot         bool   `json:"is_robot"`
	ShutUpTimestamp int    `json:"shut_up_timestamp"`
	Role            string `json:"role"`
	Title           string `json:"title"`
}
type reqGroupMemberList struct {
	GroupId int64 `json:"group_id"`
	NoCache bool  `json:"no_cache"`
}
type MsgInforeq struct {
	MessageId int `json:"message_id"`
}

// 群聊戳戳数据处理 不导出
type gpokeData struct {
	GroupId int64 `json:"group_id"`
	UserId  int64 `json:"user_id"`
}

// 私聊戳戳数据处理 不导出
type ppokedata struct {
	UserId int64 `json:"user_id"`
}

// 撤回消息处理不导出
type deletemsgdata struct {
	MessageId int64 `json:"message_id"`
}
