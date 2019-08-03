package model

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var (
	UserDaoInstance *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (dao *UserDao) {
	dao = &UserDao{
		pool: pool,
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	conn.Do("auth", "baijin")
	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOT_EXEISTS
		}
		return nil, err
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return
	}
	return
}

func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		fmt.Println("this.getUserById err ", err)
		return
	}

	if user.Password != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
