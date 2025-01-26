package Message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xww2652008969/wbot/client/utils"
	"strconv"
)

// 提供获取数据的api
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
