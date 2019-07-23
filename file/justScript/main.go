package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

//Config Config
type Config struct {
	Projects []string `json:"projects"`
	Scripts  []string `json:"scripts"`
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

func runScript(path string, config Config) {
	fmt.Println("begin update for path:", path)
	//cd到目标目录
	cdErr := os.Chdir(path)
	if cdErr != nil {
		fmt.Println("cd err path:", path, cdErr)
	}
	for _, script := range config.Scripts {
		_, execErr := execShell(script)
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
		runScript(path, config)
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
