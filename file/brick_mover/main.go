package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

//Config Config
type Config struct {
	Projects        []string `json:"projects"`
	TargetPackages  []string `json:"targetPackages"`
	ScriptsForBegin []string `json:"scriptsForBegin"`
	ScriptsForEnd   []string `json:"scriptsForEnd"`
}

func loadConfig() (Config, error) {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("加载配置文件错误", err)
		return Config{}, errors.New("加载配置文件错误")
	}
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("配置解析错误", err)
		return Config{}, errors.New("配置解析错误")
	}
	fmt.Println(config)
	return config, nil
}

func runUpdate(path string, config Config) {
	fmt.Println("begin update for path:", path)
	//cd到目标目录
	cdErr := os.Chdir(path)
	if cdErr != nil {
		fmt.Println("cd err path:", path, cdErr)
	}

	//执行前置脚本
	for _, script := range config.ScriptsForBegin {
		_, startExecErr := execShell(script)
		if startExecErr != nil {
			fmt.Printf("start script err ,script:%s,  path:%s,   err:%v", script, path, startExecErr)
			return
		}
		fmt.Println("start script success : ", script)
	}

	//获取package.json原先的文件内容
	rawPackgeJson, err := ioutil.ReadFile("./package.json")
	if err != nil {
		fmt.Println("read file err. path :", path, err)
		return
	}
	str := string(rawPackgeJson)
	strArr := strings.Split(str, "\n")
	//查找到目标报名的行 删除之
	for _, key := range config.TargetPackages {
		for i, v := range strArr {
			if strings.Index(v, key) != -1 {
				strArr = append(strArr[:i], strArr[i+1:]...)
				break
			}
		}
	}
	//删除后的package.json 内容
	finalPackgeJson := strings.Join(strArr, "\n")
	//写入文件
	err = ioutil.WriteFile("./package.json", []byte(finalPackgeJson), 0666)
	if err != nil {
		fmt.Println("write file err .path :", path, err)
		return
	}
	fmt.Println("refactor file  success")
	//文件处理完成开始脚本执行
	_, execErr := execShell("yarn install")
	if execErr != nil {
		fmt.Println("first yarn install err . path", path, execErr)
		return
	}
	fmt.Println("first yarn install success")
	//恢复package.json
	_, execErr = execShell("git checkout package.json")
	if execErr != nil {
		fmt.Println("recover package.json err . path", path, execErr)
		return
	}
	fmt.Println("recover package.json success")
	//再次执行yarn install 下载最新的target 包
	_, execErr = execShell("yarn install")
	if execErr != nil {
		fmt.Println("second yarn install err . path", path, execErr)
	}
	fmt.Println("second yarn install success")

	//target包已经完成更新 然后执行config中的script
	for _, script := range config.ScriptsForEnd {
		_, execErr = execShell(script)
		if execErr != nil {
			fmt.Printf("end script err ,script:%s,  path:%s,   err:%v", script, path, execErr)
			return
		}
		fmt.Println("end script success : ", script)
	}
	fmt.Println("complete update for path:", path)
}

func main() {
	config, configErr := loadConfig()
	if configErr != nil {
		return
	}
	for _, path := range config.Projects {
		runUpdate(path, config)
	}

}

func execShell(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	checkErr(err)

	return out.String(), err
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)

	}
}
