package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type winCmds struct {
	name string
	cmd  string
}

type Host struct {
	User_List []UserStruct
}

type UserStruct struct {
	username string
	id       string
}

type ReturnData struct {
	index int
	key   string
	value string
}

func main() {
	userList := UserStruct{}
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
			name: "Users",
			cmd:  "wmic useraccount get name",
		},
		{
			name: "Users",
			cmd:  "wmic useraccount get sid",
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

	for i := 0; i < len(cmdList); i++ {
		output, err := exec.Command("cmd", "/c", cmdList[i].cmd).Output()
		// var tmpData string
		if err != nil {
			log.Fatal(err)
		}

		switch cmdList[i].name {
		case "Users":
			tmpData := Users(string(output))
			// jsonTest, _ := json.Marshal(tmpData)
			// fmt.Println(tmpData)
			for i, data := range tmpData {
				fmt.Printf("%-v", data)
				if data.key == "Home" {
					// fmt.Println("test")
				}
				userList = UserStruct{username: data.value}
				userList = UserStruct{id: data.value}
				fmt.Println(data)
				fmt.Println(userList.username)
			}
		}
	}
}

func SystemInfo(raw_output string) string {
	fmt.Println("0")
	return (raw_output)
}
func Users(raw_output string) []ReturnData {
	str_split := strings.Split(raw_output, "\r\n")
	var tmpData []ReturnData

	for i, instance := range str_split {
		str_split[i] = strings.Trim(string(str_split[i]), " \r")

		if str_split[i] != "" && str_split[0] == "Name" && i > 0 {
			jsonData := ReturnData{i, str_split[0], instance}
			tmpData = append(tmpData, jsonData)

		} else if str_split[i] != "" && str_split[0] == "SID" && i > 0 {
			jsonData := ReturnData{i, str_split[0], instance}
			tmpData = append(tmpData, jsonData)

		}
	}
	return (tmpData)
}
func Apps(raw_output string) string {
	fmt.Println("9")
	return (raw_output)
}
func Services(raw_output string) string {
	fmt.Println("10")
	return (raw_output)
}
