package process

import "fmt"

type UserManager struct {
	onlineUsers map[int]*UserProcess
}

var (
	userManagerInstance *UserManager
)

func init() {
	userManagerInstance = &UserManager{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (this *UserManager) AddOnlineUser(user *UserProcess) {
	this.onlineUsers[user.userId] = user
}

func (this *UserManager) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

func (this *UserManager) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

func (this *UserManager) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err := fmt.Errorf("不存在 userId %d", userId)
		return up,err
	}
	return
}
