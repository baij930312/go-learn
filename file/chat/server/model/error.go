package model

import "errors"

var (
	ERROR_USER_NOT_EXEISTS = errors.New("不存在 ")
	ERROR_USER_EXEISTS     = errors.New("已存在")
	ERROR_USER_PWD         = errors.New("密码无效")
	ERROR_UNKNOW           = errors.New("位置错误")
)
