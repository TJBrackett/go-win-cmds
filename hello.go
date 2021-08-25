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
	User_List []User
}

type User struct {
	username string
	id       string
}

func main() {
	// userList := []User{}
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

		if err != nil {
			log.Fatal(err)
		}

		switch cmdList[i].name {
		case "Users":
			Users(string(output))
			fmt.Println(string(output))
		}
	}
}

func SystemInfo(raw_output string) string {
	fmt.Println("0")
	return (raw_output)
}
func Users(raw_output string) string {
	str_split := strings.Split(raw_output, "\r\n")
	fmt.Println(str_split)
	for i := 1; i < len(str_split)-1; i++ {
		// str_split[i] = str_split[i].Trim(str_split[i], "\r\n")
		if str_split[0] == "Name" {
			fmt.Println(str_split)
			// userList = append(userList, User{"username:", str_split[i]})
		} else if str_split[i] != "" && str_split[0] == "SID" {
			// userList = append(userList, User{"id:", str_split[i]})
			fmt.Println(str_split[i])
		}
	}
	return ("")
}
func Apps(raw_output string) string {
	fmt.Println("9")
	return (raw_output)
}
func Services(raw_output string) string {
	fmt.Println("10")
	return (raw_output)
}
