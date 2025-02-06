package Message

type Message struct {
	SelfId      int64  `json:"self_id"`
	UserId      int64  `json:"user_id"`
	Time        int64  `json:"time"`
	MessageId   int64  `json:"message_id"`
	MessageSeq  int    `json:"message_seq"`
	RealId      int    `json:"real_id"`
	MessageType string `json:"message_type"`
	Sender      struct {
		UserId   int    `json:"user_id"`
		Nickname string `json:"nickname"`
		Card     string `json:"card"`
		Role     string `json:"role"`
	} `json:"sender"`
	RawMessage    string `json:"raw_message"`
	Font          int    `json:"font"`
	SubType       string `json:"sub_type"`
	Message       []upd  `json:"message"`
	MessageFormat string `json:"message_format"`
	PostType      string `json:"post_type"`
	GroupId       int64  `json:"group_id"`
	Httpclient    string `json:"_"`
	updata        []upd

	NoticeType string `json:"notice_type,omitempty"`
	TargetId   int64  `json:"target_id,omitempty"`
	SenderId   int    `json:"sender_id,omitempty"`
	RawInfo    []struct {
		Col  string `json:"col,omitempty"`
		Nm   string `json:"nm,omitempty"`
		Type string `json:"type"`
		Uid  string `json:"uid,omitempty"`
		Jp   string `json:"jp,omitempty"`
		Src  string `json:"src,omitempty"`
		Txt  string `json:"txt,omitempty"`
		Tp   string `json:"tp,omitempty"`
	} `json:"raw_info,omitempty"`
}

// 用于处理发送信息的数据结构
type upd struct {
	Type string `json:"type"`
	Data struct {
		Qq   string `json:"qq"`
		Name string `json:"name"`
		Text string `json:"text"`
		File string `json:"file"`
		Id   int    `json:"id"`
	} `json:"data"`
}

type uPMessage struct {
	Group_id string `json:"group_id"`
	UserId   string `json:"user_id"`
	Message  []upd  `json:"message"`
}

type Event func(message Message)

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

// 点赞数据处理 不导出
type sendlikedata struct {
	UserId string `json:"user_id"`
	Times  int    `json:"times"`
}

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
