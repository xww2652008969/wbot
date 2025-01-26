package Message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xww2652008969/wbot/client/utils"
	"io"
	"strconv"
)

// 提供用于处理发送信息的api
// AddText 添加string文本
func (m *Message) AddText(text string) *Message {
	var da upd
	da.Type = "text"
	da.Data.Text = text
	m.updata = append(m.updata, da)
	return m
}

// Sendat 设置@某人  传入qq号 如果为0则@全体
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

// AddImage 添加图片
func (m *Message) AddImage(url string) *Message {
	var da upd
	da.Type = "image"
	da.Data.File = url
	m.updata = append(m.updata, da)
	return m
}

// AddFace 添加系统表情
func (m *Message) AddFace(id int) *Message {
	var da upd
	da.Type = "face"
	da.Data.Id = id
	m.updata = append(m.updata, da)
	return m
}

// AddCopy 复制
func (m *Message) AddCopy() *Message {
	m.updata = m.Message
	return m
}

// AddRecord 添加语音
func (m *Message) AddRecord(url string) *Message {
	var da upd
	da.Type = "record"
	da.Data.File = url
	m.updata = append(m.updata, da)
	return m
}

// AddVideo
func (m *Message) AddVideo(url string) *Message {
	var da upd
	da.Type = "video"
	da.Data.File = url
	m.updata = append(m.updata, da)
	return m
}

// Addreply 添加回复消息的信息 传入Messageid
func (m *Message) Addreply(id int64) *Message {
	var da upd
	da.Type = "reply"
	da.Data.Id = int(id)
	m.updata = append([]upd{da}, m.updata...)
	return m
}

// AddMusicCard 添加音乐卡片（不知道如何获取id和只能用qq？？？）
func (m *Message) AddMusicCard(t string, id int) *Message {
	var da upd
	da.Type = "music"
	da.Data.Id = id
	da.Type = t
	m.updata = append(m.updata, da)
	return m
}

// AddDice  发送骰子表情(??直接face就可以了的？？)
func (m *Message) AddDice() *Message {
	var da upd
	da.Type = "dice"
	m.updata = append(m.updata, da)
	return m
}
func (m *Message) AddFile(url string) *Message {
	var da upd
	da.Type = "file"
	da.Data.File = url
	m.updata = append(m.updata, da)
	return m
}

// SendGroupMsg 发送私聊群聊信息
func (m *Message) SendGroupMsg(id int64, clean ...bool) {
	var u uPMessage
	u.Group_id = strconv.FormatInt(id, 10)
	u.Message = m.updata
	a, _ := json.Marshal(u)
	res, err := utils.Httppost(m.Httpclient+"/send_group_msg", nil, bytes.NewBuffer(a))
	if err != nil {
		fmt.Println(err)
	}
	d, _ := io.ReadAll(res.Body)
	fmt.Println(string(d))
	if len(clean) > 0 && !clean[0] {

		return
	}
	m.updata = nil
}

// SendPrivateMsg 发送私聊消息
func (m *Message) SendPrivateMsg(id int64, clean ...bool) {
	var u uPMessage
	u.UserId = strconv.FormatInt(id, 10)
	u.Message = m.updata
	a, _ := json.Marshal(u)
	res, err := utils.Httppost(m.Httpclient+"/send_private_msg", nil, bytes.NewBuffer(a))
	if err != nil {
		fmt.Println(err)
	}
	d, _ := io.ReadAll(res.Body)
	fmt.Println(string(d))
	if len(clean) > 0 && !clean[0] {

		return
	}
	m.updata = nil
}

// SendGrouppoke 戳戳群成员
func (m *Message) SendGrouppoke(groupid int64, userid int64) {
	p := gpokeData{
		GroupId: groupid,
		UserId:  userid,
	}
	a, err := json.Marshal(p)
	fmt.Print(err)
	fmt.Println(string(a))
	res, err := utils.Httppost(m.Httpclient+"/group_poke", nil, bytes.NewBuffer(a))
	if err != nil {
		fmt.Println(err)
	}
	d, _ := io.ReadAll(res.Body)
	fmt.Println(string(d))
}

// SendPrivatepoke 私聊戳戳
func (m *Message) SendPrivatepoke(userid int64) {
	p := ppokedata{UserId: userid}
	a, _ := json.Marshal(p)
	res, err := utils.Httppost(m.Httpclient+"/friend_poke", nil, bytes.NewBuffer([]byte(a)))
	if err != nil {
		fmt.Println(err)
	}
	d, _ := io.ReadAll(res.Body)
	fmt.Println(string(d))
}

// DeleMsg 撤回
func (m *Message) DeleMsg(msgid int64) {
	data := deletemsgdata{MessageId: msgid}
	a, _ := json.Marshal(data)
	res, err := utils.Httppost(m.Httpclient+"/delete_msg", nil, bytes.NewBuffer(a))
	if err != nil {
		fmt.Println(err)
	}
	d, _ := io.ReadAll(res.Body)
	fmt.Println(string(d))
}
