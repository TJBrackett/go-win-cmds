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

type CMDB struct {
	User_List []UserStruct
}

type UserStruct struct {
	username string
	id       string
}

type ReturnData struct {
	key   string
	value string
}

func main() {
	// cmdb := CMDB{}
	userList := []UserStruct{}
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
			cmd:  "wmic useraccount get name,sid /format:csv",
		},
		// {
		// 	name: "Users",
		// 	cmd:  "wmic useraccount get sid",
		// },
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

		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(string(output))
		switch cmdList[i].name {
		case "Users":
			userList = Users(string(output))
			// fmt.Printf("%+q\r\n", userList)
			// Users(string(output))
			// for i, data := range tmpData {
			// 	if data.key == "Name" {
			// 		userList[i].username = data.value
			// 	} else if data.key == "SID" {
			// 		userList[i].id = data.value
			// 	}
			// 	fmt.Println(userList[0].username)
			// 	fmt.Println(userList[0].id)
			// }
		}
	}
}

func SystemInfo(raw_output string) string {
	fmt.Println("0")
	return (raw_output)
}
func Users(rawOutput string) []UserStruct {
	strSplit := strings.Split(rawOutput, "\r\n")
	tmpUserData := make([]UserStruct, len(strSplit))

	for i, instance := range strSplit {
		if i != 0 {
			if instance != "" {
				splitInstance := strings.Split(instance, ",")

				tmpUserData[i].username = splitInstance[1]
				tmpUserData[i].id = splitInstance[2]
				// fmt.Printf("Username: %v\nSID: %v\n", splitInstance[1], splitInstance[2])
			}
		}
	}
	return (tmpUserData)
	// str_split := strings.Split(raw_output, "\r\n")
	// var tmpData []ReturnData
	// fmt.Println(str_split)
	// for i, instance := range str_split {
	// 	str_split[i] = strings.Trim(string(str_split[i]), " \r\t")
	// 	fmt.Println(str_split[i])

	// 	if str_split[i] != "" && str_split[0] == "Name" && i > 0 {
	// 		jsonData := ReturnData{str_split[0], instance}
	// 		tmpData = append(tmpData, jsonData)

	// 	} else if str_split[i] != "" && str_split[0] == "SID" && i > 0 {
	// 		jsonData := ReturnData{str_split[0], instance}
	// 		tmpData = append(tmpData, jsonData)

	// 	}
	// }
}
func Apps(raw_output string) string {
	fmt.Println("9")
	return (raw_output)
}
func Services(raw_output string) string {
	fmt.Println("10")
	return (raw_output)
}
