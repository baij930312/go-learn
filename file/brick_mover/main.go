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
	Projects            []string `json:"projects"`
	InvalidPackageNames []string `json:"InvalidPackageNames"`
}

func loadConfig() (Config, error) {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("加载配置文件错误err", err)
		return Config{}, errors.New("加载配置文件错误err")
	}
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("配置解析错误err", err)
		return Config{}, errors.New("配置解析错误err")
	}
	fmt.Println(config)
	return config, nil
}

func runUpdate(path string, config Config) {
	execErr := os.Chdir(path)
	if execErr != nil {
		fmt.Println("执行错误", execErr)
	}
	packageJSON, err := ioutil.ReadFile("./package.json")
	if err != nil {
		fmt.Println("读取文件", err)
	}
	var m map[string]interface{}
	err = json.Unmarshal(packageJSON, &m)
	if err != nil {
		return
	}
	//删除需要共享的项目的key
	if v := m["dependencies"].(map[string]interface{}); v != nil {
		for _, key := range config.InvalidPackageNames {
			delete(v, key)
		}
		fmt.Println(v)
	}
	file, err := json.Marshal(m)
	if err != nil {
		return
	}

	err = ioutil.WriteFile("./package.json", file, 0666)
	if err != nil {
		return
	}
	// result, execErr := execShell("yarn install")
	// if execErr != nil {
	// 	fmt.Println("执行错误", execErr)
	// }
	// fmt.Println(result)
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
		panic(err)
	}
}
