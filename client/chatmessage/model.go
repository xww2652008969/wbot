package chatmessage

type ChatMessageData struct {
	Type string         `json:"type"`
	Data messagePayload `json:"data"`
}
type messagePayload struct {
	Qq        string `json:"qq,omitempty"`
	Name      string `json:"name,omitempty"`
	Text      string `json:"text,omitempty"`
	File      string `json:"file,omitempty"`
	Id        string `json:"id,omitempty"`
	Url       string `json:"url,omitempty"`
	Sub_type  int    `json:"sub_Type,omitempty"`
	File_size string `json:"file_Size,omitempty"`
	Type      string `json:"type,omitempty"`
	Data      string `json:"data,omitempty"`
}
type ChatMessage struct {
	Group_id int64             `json:"group_id"`
	UserId   int64             `json:"user_id"`
	Message  []ChatMessageData `json:"message"`
}
