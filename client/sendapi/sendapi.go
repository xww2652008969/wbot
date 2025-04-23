package sendapi

import (
	"bytes"
	"encoding/json"
	"github.com/xww2652008969/wbot/client/utils"
	"io"
	"net/http"
	"strconv"
)

// SendGroupMsg 发送私聊群聊信息
func (sendapi *SendAPI) SendGroupMsg() (*http.Response, error) {
	if sendapi.chatmessage.UserId == 0 {
		return nil, nil
	}
	postdata, _ := json.Marshal(sendapi.chatmessage)
	return utils.Httppost(sendapi.httpurl+"/send_group_msg", nil, bytes.NewBuffer(postdata))
}

// SendPrivateMsg 发送私聊消息 可传入群号可以
func (sendapi *SendAPI) SendPrivateMsg() (*http.Response, error) {
	postdata, _ := json.Marshal(sendapi.chatmessage)
	return utils.Httppost(sendapi.httpurl+"/send_private_msg", nil, bytes.NewBuffer(postdata))
}

// 发送合并消息
func (sendapi *SendAPI) SendForwardMsg(data []byte) (*http.Response, error) {
	return utils.Httppost(sendapi.httpurl+"/send_forward_msg", nil, bytes.NewBuffer(data))
}

// SendGrouppoke 戳戳群成员
func (sendapi *SendAPI) SendGrouppoke(groupid int64, userid int64) (*http.Response, error) {
	p := gpokeData{
		GroupId: groupid,
		UserId:  userid,
	}
	a, _ := json.Marshal(p)
	return utils.Httppost(sendapi.httpurl+"/group_poke", nil, bytes.NewBuffer(a))
}

// SendPrivatepoke 私聊戳戳
func (sendapi *SendAPI) SendPrivatepoke(userid int64) (*http.Response, error) {
	p := ppokedata{UserId: userid}
	a, _ := json.Marshal(p)
	return utils.Httppost(sendapi.httpurl+"/friend_poke", nil, bytes.NewBuffer([]byte(a)))

}

// DeleMsg 撤回
func (sendapi *SendAPI) DeleMsg(msgid int64) (*http.Response, error) {
	data := deletemsgdata{MessageId: msgid}
	a, _ := json.Marshal(data)
	return utils.Httppost(sendapi.httpurl+"/delete_msg", nil, bytes.NewBuffer(a))
}

// SendLike 点赞
func (sendapi *SendAPI) SendLike(id int64, times int) (*http.Response, error) {
	data := sendlikedata{
		UserId: strconv.FormatInt(id, 10),
		Times:  times,
	}
	jsonBytes, _ := json.Marshal(data)
	return utils.Httppost(sendapi.httpurl+"/send_like", nil, bytes.NewBuffer(jsonBytes))
}

// GetGroupMemberList 返回群成员列表
func (sendapi *SendAPI) GetGroupMemberList(id int64) (GroupMemberList, error) {
	var g GroupMemberList
	req := reqGroupMemberList{
		GroupId: id,
		NoCache: false,
	}
	d, _ := json.Marshal(req)
	res, err := utils.Httppost(sendapi.httpurl+"/get_group_member_list", nil, bytes.NewReader(d))
	if err != nil {
		return g, err
	}
	data, _ := io.ReadAll(res.Body)
	json.Unmarshal(data, &g)
	return g, nil
}

// GetGroupNemberInfo 获取群成员信息
func (sendapi *SendAPI) GetGroupNemberInfo(groupid, user int64) (groupMemberInfo, error) {
	var g groupMemberInfo
	req := reqgroupMemberInfo{
		GroupId: groupid,
		UserId:  user,
		NoCache: true,
	}
	d, _ := json.Marshal(req)
	res, err := utils.Httppost(sendapi.httpurl+"/get_group_member_info", nil, bytes.NewBuffer(d))
	if err != nil {
		return g, err
	}
	data, _ := io.ReadAll(res.Body)
	json.Unmarshal(data, &g)
	return g, nil
}

//func (sendapi *SendAPI) GetMsginfo(msgid int) (MsgInfo, error) {
//	var g MsgInfo
//	req := MsgInforeq{MessageId: msgid}
//	d, _ := json.Marshal(req)
//	res, err := utils.Httppost(sendapi.httpurl+"/friend_poke", nil, bytes.NewBuffer(d))
//	if err != nil {
//		return g, err
//	}
//	data, _ := io.ReadAll(res.Body)
//	json.Unmarshal(data, &g)
//	return g, nil
//}
