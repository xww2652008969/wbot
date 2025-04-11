package client

import (
	"bytes"
	"encoding/json"
	"github.com/xww2652008969/wbot/client/utils"
	"io"
	"log"
	"strconv"
)

// SendLike 点赞
func (m *Client) SendLike(id int64, times int) *Client {
	data := sendlikedata{
		UserId: strconv.FormatInt(id, 10),
		Times:  times,
	}
	jsonBytes, _ := json.Marshal(data)
	res, err := utils.Httppost(m.Config.Clienthttp+"/send_like", nil, bytes.NewBuffer(jsonBytes))
	if err != nil {
		m.logdebug(err)
		return m
	}
	m.logdebug(res)
	return m
}

// GetGroupMemberList 返回群成员列表
func (m *Client) GetGroupMemberList(id int64) (GroupMemberList, error) {
	var g GroupMemberList
	req := reqGroupMemberList{
		GroupId: id,
		NoCache: false,
	}
	d, _ := json.Marshal(req)
	res, err := utils.Httppost(m.Config.Clienthttp+"/get_group_member_list", nil, bytes.NewReader(d))
	if err != nil {
		return g, err
	}
	data, _ := io.ReadAll(res.Body)
	json.Unmarshal(data, &g)
	return g, nil
}

// GetGroupNemberInfo 获取群成员信息
func (m *Client) GetGroupNemberInfo(groupid, user int64) (groupMemberInfo, error) {
	var g groupMemberInfo
	req := reqgroupMemberInfo{
		GroupId: groupid,
		UserId:  user,
		NoCache: true,
	}
	d, _ := json.Marshal(req)
	res, err := utils.Httppost(m.Config.Clienthttp+"/get_group_member_info", nil, bytes.NewBuffer(d))
	if err != nil {
		return g, err
	}
	data, _ := io.ReadAll(res.Body)
	json.Unmarshal(data, &g)
	return g, nil
}
func (m *Client) GetMsginfo(msgid int) (MsgInfo, error) {
	var g MsgInfo
	req := MsgInforeq{MessageId: msgid}
	d, _ := json.Marshal(req)
	res, err := utils.Httppost(m.Config.Clienthttp+"/get_msg", nil, bytes.NewBuffer(d))
	if err != nil {
		return g, err
	}
	data, _ := io.ReadAll(res.Body)
	json.Unmarshal(data, &g)
	return g, nil
}

// 用于处理发送消息
// AddText 添加string文本
func (m *Client) AddText(text string) *Client {
	var da upd
	da.Type = "text"
	da.Data.Text = text
	m.updata = append(m.updata, da)
	return m
}

// Sendat 设置@某人  传入qq号 如果为0则@全体
func (m *Client) Sendat(qq int64) *Client {
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
func (m *Client) AddImage(url string) *Client {
	var da upd
	da.Type = "image"
	da.Data.File = url
	m.updata = append(m.updata, da)
	return m
}

// AddFace 添加系统表情
func (m *Client) AddFace(id int) *Client {
	var da upd
	da.Type = "face"
	da.Data.Id = strconv.Itoa(id)
	m.updata = append(m.updata, da)
	return m
}

// AddCopy 复制
func (m *Client) AddCopy(message Message) *Client {
	m.updata = message.Message
	return m
}

// AddRecord 添加语音
func (m *Client) AddRecord(url string) *Client {
	var da upd
	da.Type = "record"
	da.Data.File = url
	m.updata = append(m.updata, da)
	return m
}

// AddVideo
func (m *Client) AddVideo(url string) *Client {
	var da upd
	da.Type = "video"
	da.Data.File = url
	m.updata = append(m.updata, da)
	return m
}

// Addreply 添加回复消息的信息 传入Messageid
func (m *Client) Addreply(id int64) *Client {
	var da upd
	da.Type = "reply"
	da.Data.Id = strconv.Itoa(int(id))
	m.updata = append([]upd{da}, m.updata...)
	return m
}

// AddMusicCard 添加音乐卡片（不知道如何获取id和只能用qq？？？）
//func (m *Client) AddMusicCard(t string, id int) *Client {
//	var da upd
//	da.Type = "music"
//	da.Data.Id = strconv.Itoa(id)
//	da.Data.Type = t
//	fmt.Println(da)
//	m.updata = append(m.updata, da)
//	return m
//}

// AddDice  发送骰子表情(??直接face就可以了的？？)
func (m *Client) AddDice() *Client {
	var da upd
	da.Type = "dice"
	m.updata = append(m.updata, da)
	return m
}
func (m *Client) AddFile(url string) *Client {
	var da upd
	da.Type = "file"
	da.Data.File = url
	m.updata = append(m.updata, da)
	return m
}

// SendGroupMsg 发送私聊群聊信息
func (m *Client) SendGroupMsg(id int64) {
	var u uPMessage
	u.Group_id = strconv.FormatInt(id, 10)
	u.Message = m.updata
	a, _ := json.Marshal(u)
	res, err := utils.Httppost(m.Config.Clienthttp+"/send_group_msg", nil, bytes.NewBuffer(a))
	if err != nil {
		m.log.Println(err)
	}
	d, _ := io.ReadAll(res.Body)
	log.Println(string(d))
	m.updata = nil
}

// SendPrivateMsg 发送私聊消息 可传入群号可以给不
func (m *Client) SendPrivateMsg(userid int64, clean ...int64) {
	var u uPMessage
	u.UserId = strconv.FormatInt(userid, 10)
	if len(clean) > 0 {
		u.Group_id = strconv.FormatInt(clean[0], 10)
	}
	u.Message = m.updata
	a, _ := json.Marshal(u)
	m.log.Println(string(a))
	res, err := utils.Httppost(m.Config.Clienthttp+"/send_private_msg", nil, bytes.NewBuffer(a))
	if err != nil {
		m.log.Println(err)
	}
	d, _ := io.ReadAll(res.Body)
	m.log.Println(string(d))
	m.updata = nil
}

// SendGrouppoke 戳戳群成员
func (m *Client) SendGrouppoke(groupid int64, userid int64) {
	p := gpokeData{
		GroupId: groupid,
		UserId:  userid,
	}
	a, err := json.Marshal(p)
	m.log.Print(err)
	m.log.Println(string(a))
	res, err := utils.Httppost(m.Config.Clienthttp+"/group_poke", nil, bytes.NewBuffer(a))
	if err != nil {
		m.log.Println(err)
	}
	d, _ := io.ReadAll(res.Body)
	m.log.Println(string(d))
}

// SendPrivatepoke 私聊戳戳
func (m *Client) SendPrivatepoke(userid int64) {
	p := ppokedata{UserId: userid}
	a, _ := json.Marshal(p)
	res, err := utils.Httppost(m.Config.Clienthttp+"/friend_poke", nil, bytes.NewBuffer([]byte(a)))
	if err != nil {
		m.log.Println(err)
	}
	d, _ := io.ReadAll(res.Body)
	m.log.Println(string(d))
}

// DeleMsg 撤回
func (m *Client) DeleMsg(msgid int64) {
	data := deletemsgdata{MessageId: msgid}
	a, _ := json.Marshal(data)
	res, err := utils.Httppost(m.Config.Clienthttp+"/delete_msg", nil, bytes.NewBuffer(a))
	if err != nil {
		m.log.Println(err)
	}
	d, _ := io.ReadAll(res.Body)
	m.log.Println(string(d))
}

// 发送合并消息
func (m *Client) SendForwardMsg(data []byte) {
	res, err := utils.Httppost(m.Config.Clienthttp+"/send_forward_msg", nil, bytes.NewBuffer(data))
	if err != nil {
		m.log.Println(err)
	}
	d, _ := io.ReadAll(res.Body)
	m.log.Println(string(d))
}
