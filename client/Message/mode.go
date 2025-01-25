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
	RawMessage string `json:"raw_message"`
	Font       int    `json:"font"`
	SubType    string `json:"sub_type"`
	Message    []struct {
		Type string `json:"type"`
		Data struct {
			Qq   string `json:"qq"`
			Name string `json:"name"`
			Text string `json:"text"`
		} `json:"data"`
	} `json:"message"`
	MessageFormat string `json:"message_format"`
	PostType      string `json:"post_type"`
	GroupId       int64  `json:"group_id"`
	Httpclient    string `json:"_"`
	updata        []upd  `json:"_"`
}

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
	Message  []upd  `json:"message"`
}

type MessageEvent func(message Message)
