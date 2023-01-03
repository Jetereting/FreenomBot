package wechat

import (
	"github.com/tidwall/gjson"
	"time"

	"github.com/astaxie/beego/httplib"
)

// Send message
func Send(corpID, corpSecret, agentID, msg string) {
	token, e := httplib.Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpID + "&corpsecret=" + corpSecret).String()
	if e != nil {
		return
	}
	token = gjson.Get(token, "access_token").String()

	msg += time.Now().Format("\n\n2006-01-02 15:04")
	_, _ = httplib.Post("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + token).Body(`
{
   "touser" : "@all",
   "msgtype" : "text",
   "agentid" : ` + agentID + `,
   "text" : {
       "content" : "` + msg + `"
   },
   "enable_duplicate_check": 1,
}`).DoRequest()
}
