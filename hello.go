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
		switch cmdList[i].name {
		case "Users":
			userList = Users(string(output))
			fmt.Printf("%+q\r\n", userList)
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
}
func Apps(raw_output string) string {
	fmt.Println("9")
	return (raw_output)
}
func Services(raw_output string) string {
	fmt.Println("10")
	return (raw_output)
}
