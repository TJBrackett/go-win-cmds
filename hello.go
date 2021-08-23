package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"reflect"
	"strings"
)

type winCmds struct {
	name     string
	funcName interface{}
	cmd      string
}

type cmdMap map[string]interface{}

var cmdNames = cmdMap{}

func main() {
	var cmdList = []winCmds{
		// {
		// 	name:     "SystemInfo",
		// 	funcName: SystemInfo,
		// 	cmd:      "systeminfo",
		// },
		// {
		// 	name:     "Uptime",
		// 	funcName: Uptime,
		// 	cmd:      "wmic path win32_operatingsystem get lastbootuptime",
		// },
		// {
		// 	name:     "OS",
		// 	funcName: OS,
		// 	cmd:      "wmic os get version",
		// },
		// {
		// 	name:     "IP",
		// 	funcName: IP,
		// 	cmd:      "netsh interface ip show address",
		// },
		// {
		// 	name:     "MAC",
		// 	funcName: MAC,
		// 	cmd:      "wmic nic",
		// },
		{
			name:     "Users",
			funcName: Users,
			cmd:      "wmic useraccount get name",
		},
		{
			name:     "Users",
			funcName: Users,
			cmd:      "wmic useraccount get sid",
		},
		// {
		// 	name:     "Apps",
		// 	funcName: Apps,
		// 	cmd:      "wmic product",
		// },
		// {
		// 	name:     "Services",
		// 	funcName: Services,
		// 	cmd:      "wmic service",
		// },
	}

	cmdNames = map[string]interface{}{}

	for _, v := range cmdList {
		MapCmdNames(v.name, v.funcName)
	}

	for i := 0; i < len(cmdList); i++ {
		output, err := exec.Command("cmd", "/c", cmdList[i].cmd).Output()

		if err != nil {
			log.Fatal(err)
		}
		// splitStr := strings.Split(string(output), ",")
		// fmt.Println(splitStr[0])
		testCall, _ := CallFunc(cmdList[i].name, string(output))
		var printCall string
		printCall = testCall.(string)
		fmt.Println(printCall)
	}

	// fmt.Println(runtime.GOOS)
}

func MapCmdNames(name string, funcName interface{}) {
	cmdNames[name] = funcName
}

func CallFunc(funcName string, params ...interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(cmdNames[funcName])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is out of index.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	var res []reflect.Value
	res = f.Call(in)
	result = res[0].Interface()
	return
}

func SystemInfo(raw_output string) string {
	fmt.Println("0")
	return (raw_output)
}
func Users(raw_output string) string {
	str_split := strings.Split(raw_output, " ")

	for i := 1; i < len(str_split)-1; i++ {
		if str_split[i] != "" && str_split[0] == "Name" {
			fmt.Println(str_split[i])
			// struct array push username @ i
		} else if str_split[i] != "" && str_split[0] == "SID" {
			fmt.Println(str_split[i])
			// struct array push id @ i
		}
	}

	return (str_split[0])
}
func Apps(raw_output string) string {
	fmt.Println("9")
	return (raw_output)
}
func Services(raw_output string) string {
	fmt.Println("10")
	return (raw_output)
}

func FormatArray(c string) {
	f := func(c rune) bool {
		return c == ','
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("a,,b,c", f))
}
