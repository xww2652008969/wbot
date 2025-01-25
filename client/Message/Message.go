package Message

import (
	"bot/client/utils"
	"bytes"
	"encoding/json"
	"strconv"
)

// 添加string文本
func (m *Message) AddText(text string) *Message {
	var da upd
	da.Type = "text"
	da.Data.Text = text
	m.updata = append(m.updata, da)
	return m
}

// 设置@某人  传入qq号 如果为0则@全体
func (m *Message) Sendat(qq int64) *Message {
	var da upd
	da.Type = "at"
	if qq == 0 {
		da.Data.Qq = "all"
	} else {
		da.Data.Qq = strconv.FormatInt(qq, 10)
	}
	m.updata = append(m.updata, da)
	return m
}

// 添加图片
func (m *Message) AddImage(url string) *Message {
	var da upd
	da.Type = "image"
	da.Data.File = url
	m.updata = append(m.updata, da)
	return m
}

// 添加系统表情
func (m *Message) AddFace(id int) *Message {
	var da upd
	da.Type = "face"
	da.Data.Id = id
	m.updata = append(m.updata, da)
	return m
}

// 发送私聊群聊信息
func (m *Message) SendGroupMsg(id int64, clean ...bool) {
	var u uPMessage
	u.Group_id = strconv.FormatInt(id, 10)
	u.Message = m.updata
	a, _ := json.Marshal(u)
	println(string(a))
	utils.Httppost(m.Httpclient+"/send_group_msg", nil, bytes.NewBuffer(a))
	if len(clean) > 0 && !clean[0] {

		return
	}
	m.updata = nil
}
