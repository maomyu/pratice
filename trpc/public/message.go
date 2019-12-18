/*
* @Author: your name
 * @Date: 2019-11-23 10:35:1
 * @LastEditTime: 2019-11-23 10:36:07
 * @LastEditors your name
 * @Description: In UserSettings Edit
 * @FilePath: \GoWebd:\patice\trpc\public\message.go
*/
package public

type MessageAckReq struct {
	IsSignIn bool // 标记用户是否登录成功

	AppId    int64
	DeviceId int64  // 设备id
	UserId   int64  // 用户id
	Bytes    []byte // 字节数组
}

type MessageAckResp struct {
}

const (
	MessageTypeSync = 1 // 消息同步
	MessageTypeMail = 2 // 消息投递
)

type MessageReq struct {
	DeviceId int64  // 设备id
	Bytes    []byte // 消息投递字节包
}

type MessageResp struct {
}
