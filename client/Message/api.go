package Message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xww2652008969/wbot/client/utils"
	"io"
	"strconv"
)

// SendLike 点赞
func (m *Message) SendLike(id int64, times int) *Message {
	data := sendlikedata{
		UserId: strconv.FormatInt(id, 10),
		Times:  times,
	}
	jsonBytes, _ := json.Marshal(data)
	res, err := utils.Httppost(m.Httpclient+"/send_like", nil, bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	return m
}

// GetGroupMemberList 返回群成员列表
func (m *Message) GetGroupMemberList(id int64) (GroupMemberList, error) {
	var g GroupMemberList
	req := reqGroupMemberList{
		GroupId: id,
		NoCache: false,
	}
	d, _ := json.Marshal(req)
	res, err := utils.Httppost(m.Httpclient+"/get_group_member_list", nil, bytes.NewReader(d))
	if err != nil {
		return g, err
	}
	data, _ := io.ReadAll(res.Body)
	json.Unmarshal(data, &g)
	return g, nil
}

// GetGroupNemberInfo 获取群成员信息
func (m *Message) GetGroupNemberInfo(groupid, user int64) (groupMemberInfo, error) {
	var g groupMemberInfo
	req := reqgroupMemberInfo{
		GroupId: groupid,
		UserId:  user,
		NoCache: true,
	}
	d, _ := json.Marshal(req)
	res, err := utils.Httppost(m.Httpclient+"/get_group_member_info", nil, bytes.NewBuffer(d))
	if err != nil {
		return g, err
	}
	data, _ := io.ReadAll(res.Body)
	json.Unmarshal(data, &g)
	return g, nil
}
func (m *Message) GetMsginfo(msgid int) (MsgInfo, error) {
	var g MsgInfo
	req := MsgInforeq{MessageId: msgid}
	d, _ := json.Marshal(req)
	res, err := utils.Httppost(m.Httpclient+"/get_msg_info", nil, bytes.NewBuffer(d))
	if err != nil {
		return g, err
	}
	data, _ := io.ReadAll(res.Body)
	json.Unmarshal(data, &g)
	return g, nil
}
