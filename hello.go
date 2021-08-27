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
	UserList    []UserStruct
	AppsList    []AppsStruct
	SvcList     []SvcStruct
	PatchesList []PatchesStruct
	StartupList []StartupStruct
}

type UserStruct struct {
	username string
	id       string
}

type AppsStruct struct {
	installDate string
	location    string
	name        string
	publisher   string
	version     string
}

type SvcStruct struct {
	displayName string
	name        string
	path        string
}

type PatchesStruct struct {
	caption     string
	description string
	installDate string
	hotfixId    string
}

type StartupStruct struct {
	name     string
	location string
}

func main() {
	cmdb := CMDB{}
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
		{
			name: "Startup",
			cmd:  "wmic startup get location,name /format:csv",
		},
		{
			name: "Users",
			cmd:  "wmic useraccount get name,sid /format:csv",
		},
		{
			name: "Patches",
			cmd:  "wmic qfe get hotfixid,description,installdate,caption /format:csv",
		},
		{
			name: "Apps",
			cmd:  "wmic product get name,version,vendor,installdate,installlocation /format:csv",
		},
		{
			name: "Services",
			cmd:  "wmic service get displayname, name, pathname /format:csv",
		},
	}

	for i := 0; i < len(cmdList); i++ {
		output, err := exec.Command("cmd", "/c", cmdList[i].cmd).Output()

		if err != nil {
			log.Fatal(err)
		}

		switch cmdList[i].name {
		case "Users":
			cmdb.UserList = Users(string(output))
			fmt.Printf("%+q\r\n", cmdb.UserList)
		case "Apps":
			cmdb.AppsList = Apps(string(output))
			fmt.Printf("%+q\r\n", cmdb.AppsList)
		case "Services":
			cmdb.SvcList = Services(string(output))
			fmt.Printf("%+q\r\n", cmdb.SvcList)
		case "Patches":
			cmdb.PatchesList = Patches(string(output))
			fmt.Printf("%+q\r\n", cmdb.PatchesList)
		case "Startup":
			cmdb.StartupList = Startup(string(output))
			fmt.Printf("%+q\r\n", cmdb.StartupList)
		}
	}
}

func SystemInfo(raw_output string) string {
	fmt.Println("0")
	return (raw_output)
}
func Users(rawOutput string) []UserStruct {
	strSplit := strings.Split(rawOutput, "\r\n")
	userData := make([]UserStruct, len(strSplit))

	for i, instance := range strSplit {
		if i != 0 {
			if instance != "" {
				instance = strings.Trim(instance, "\r\n ")
				splitInstance := strings.Split(instance, ",")

				if splitInstance[1] != "Name" && splitInstance[1] != "" {
					userData[i].username = splitInstance[1]
					userData[i].id = splitInstance[2]
				}
				// The above if excludes the "Name" line but still creates an empty "" in username
				// The below if removes that element from the slice
				// if userData[i].username == "" {
				// 	fmt.Println("test")
				// 	userData = append(userData[:i], userData[i+1:]...)
				// 	i--
				// 	continue
				// }
				// fmt.Printf("Username: %v\nSID: %v\n", userData[i].username, userData[i].id)
			}
		}
	}
	return (userData)
}
func Apps(rawOutput string) []AppsStruct {
	strSplit := strings.Split(rawOutput, "\r\n")
	appsData := make([]AppsStruct, len(strSplit))

	for i, instance := range strSplit {
		if i != 0 {
			if instance != "" {
				instance = strings.Trim(instance, "\r\n ")
				splitInstance := strings.Split(instance, ",")

				if splitInstance[1] != "InstallDate" {
					appsData[i].installDate = splitInstance[1]
					appsData[i].location = splitInstance[2]
					appsData[i].name = splitInstance[3]
					appsData[i].publisher = splitInstance[4]
					appsData[i].version = splitInstance[5]
				}
			}
		}
	}
	return (appsData)
}
func Services(rawOutput string) []SvcStruct {
	strSplit := strings.Split(rawOutput, "\r\n")
	svcData := make([]SvcStruct, len(strSplit))

	for i, instance := range strSplit {
		if i != 0 {
			if instance != "" {
				instance = strings.Trim(instance, "\r\n ")
				splitInstance := strings.Split(instance, ",")

				if splitInstance[1] != "Name" {
					svcData[i].displayName = splitInstance[1]
					svcData[i].name = splitInstance[2]
					svcData[i].path = splitInstance[3]
				}
			}
		}
	}
	return (svcData)
}

func Patches(rawOutput string) []PatchesStruct {
	strSplit := strings.Split(rawOutput, "\r\n")
	patchesData := make([]PatchesStruct, len(strSplit))

	for i, instance := range strSplit {
		if i != 0 {
			if instance != "" {
				instance = strings.Trim(instance, "\r\n ")
				splitInstance := strings.Split(instance, ",")

				if splitInstance[1] != "InstallDate" {
					patchesData[i].caption = splitInstance[1]
					patchesData[i].description = splitInstance[2]
					patchesData[i].hotfixId = splitInstance[3]
					patchesData[i].installDate = splitInstance[4]
				}
			}
		}
	}
	return (patchesData)
}
func Startup(rawOutput string) []StartupStruct {
	strSplit := strings.Split(rawOutput, "\r\n")
	startupData := make([]StartupStruct, len(strSplit))

	for i, instance := range strSplit {
		if i != 0 {
			if instance != "" {
				instance = strings.Trim(instance, "\r\n ")
				splitInstance := strings.Split(instance, ",")

				if splitInstance[1] != "Name" && splitInstance[1] != "" {
					startupData[i].location = splitInstance[1]
					startupData[i].name = splitInstance[2]
				}
			}
		}
	}
	return (startupData)
}
