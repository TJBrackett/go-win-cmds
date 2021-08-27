package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

type winCmds struct {
	name string
	cmd  string
}

type CMDB struct {
	UserList          []UserStruct
	AppsList          []AppsStruct
	SvcList           []SvcStruct
	PatchesList       []PatchesStruct
	StartupList       []StartupStruct
	InterfacesList    []InterfacesStruct
	ScheduledTaskList []ScheduledTask
	SystemInfoList    []SystemInfoStruct
}

type ScheduledTask struct {
	task_name   string
	task_time   string
	task_status string
	// task_action string
	// task_path   string
}

type UserStruct struct {
	username string
	id       string
}

type InterfacesStruct struct {
	index        string
	primaryIp    string
	macAddr      string
	configuredIp []string
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

type SystemInfoStruct struct {
	hostName       string
	osName         string
	osVersion      string
	cpuInfo        string
	cpuCount       int
	lastBootTime   string
	physicalMem    string
	hardwareVendor string
	hardwareModel  string
	hardwareSerial string
}

func main() {
	cmdb := CMDB{}
	var cmdList = []winCmds{
		{
			name: "SystemInfo",
			cmd:  "systeminfo /fo csv",
		},
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
		{
			name: "Interfaces",
			cmd:  "wmic nicconfig get index,macaddress,ipaddress /format:csv",
		},
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
		{
			name: "ScheduledTasks",
			cmd:  "schtasks /fo csv",
		},
	}

	for i := 0; i < len(cmdList); i++ {
		output, err := exec.Command("cmd", "/c", cmdList[i].cmd).Output()

		if err != nil {
			log.Fatal(err)
		}

		switch cmdList[i].name {
		case "Users":
			cmdb.Users(string(output))
			// fmt.Printf("%+q\r\n", cmdb.UserList)
		case "Apps":
			cmdb.Apps(string(output))
			// fmt.Printf("%+q\r\n", cmdb.AppsList)
		case "Services":
			cmdb.Services(string(output))
			// fmt.Printf("%+q\r\n", cmdb.SvcList)
		case "Patches":
			cmdb.Patches(string(output))
			// fmt.Printf("%+q\r\n", cmdb.PatchesList)
		case "Startup":
			cmdb.Startup(string(output))
			// fmt.Printf("%+q\r\n", cmdb.StartupList)
		case "Interfaces":
			// fmt.Println(string(output))
			cmdb.Interfaces(string(output))
			// fmt.Printf("%+q\r\n", cmdb.InterfacesList)
		case "ScheduledTasks":
			cmdb.ScheduledTask(string(output))
			// fmt.Printf("%+v\r\n", cmdb.ScheduledTaskList)
		case "SystemInfo":
			cmdb.SystemInfo(string(output))
			fmt.Printf("%+v\r\n", cmdb.SystemInfoList)
		}
	}
	// fmt.Println(runtime.NumCPU())
	// fmt.Printf("%+q\r\n", cmdb)
}

func (c *CMDB) SystemInfo(rawOutput string) error {
	strSplit := strings.Split(rawOutput, "\r\n")

	for i, instance := range strSplit {
		if i != 0 {
			if instance != "" {
				instance = strings.Trim(instance, "\r\n ")
				splitInstance := strings.Split(instance, ",")

				if splitInstance[1] != "Name" {
					tmpEntry := SystemInfoStruct{}

					tmpEntry.hostName = splitInstance[0]
					tmpEntry.osName = splitInstance[1]
					tmpEntry.osVersion = splitInstance[2]
					tmpEntry.lastBootTime = splitInstance[11] + " " + splitInstance[12]
					tmpEntry.hardwareVendor = splitInstance[13]
					tmpEntry.hardwareModel = splitInstance[14]
					tmpEntry.cpuInfo = splitInstance[17]
					tmpEntry.cpuCount = runtime.NumCPU()
					tmpEntry.physicalMem = splitInstance[26] + splitInstance[27]
					tmpEntry.hardwareSerial = splitInstance[8]

					c.SystemInfoList = append(c.SystemInfoList, tmpEntry)
				}
			}
		}
	}
	return (nil)
}
func (c *CMDB) Users(rawOutput string) error {
	strSplit := strings.Split(rawOutput, "\r\n")

	for i, instance := range strSplit {
		if i != 0 {
			if instance != "" {
				instance = strings.Trim(instance, "\r\n ")
				splitInstance := strings.Split(instance, ",")

				if splitInstance[1] != "Name" {
					tmpEntry := UserStruct{}

					tmpEntry.username = splitInstance[1]
					tmpEntry.id = splitInstance[2]

					c.UserList = append(c.UserList, tmpEntry)
				}
			}
		}
	}
	return (nil)
}
func (c *CMDB) Apps(rawOutput string) error {
	strSplit := strings.Split(rawOutput, "\r\n")

	for i, instance := range strSplit {
		if i != 0 {
			if instance != "" {
				instance = strings.Trim(instance, "\r\n ")
				splitInstance := strings.Split(instance, ",")

				if splitInstance[1] != "InstallDate" {
					tmpEntry := AppsStruct{}

					tmpEntry.installDate = splitInstance[1]
					tmpEntry.location = splitInstance[2]
					tmpEntry.name = splitInstance[3]
					tmpEntry.publisher = splitInstance[4]
					tmpEntry.version = splitInstance[5]

					c.AppsList = append(c.AppsList, tmpEntry)
				}
			}
		}
	}
	return (nil)
}
func (c *CMDB) Services(rawOutput string) error {
	strSplit := strings.Split(rawOutput, "\r\n")

	for i, instance := range strSplit {
		if i != 0 {
			if instance != "" {
				instance = strings.Trim(instance, "\r\n ")
				splitInstance := strings.Split(instance, ",")

				if splitInstance[1] != "DisplayName" {
					tmpEntry := SvcStruct{}

					tmpEntry.displayName = splitInstance[1]
					tmpEntry.name = splitInstance[2]
					tmpEntry.path = splitInstance[3]

					c.SvcList = append(c.SvcList, tmpEntry)
				}
			}
		}
	}
	return (nil)
}

func (c *CMDB) Patches(rawOutput string) error {
	strSplit := strings.Split(rawOutput, "\r\n")

	for i, instance := range strSplit {
		if i != 0 {
			if instance != "" {
				instance = strings.Trim(instance, "\r\n ")
				splitInstance := strings.Split(instance, ",")

				if splitInstance[1] != "Caption" {
					tmpEntry := PatchesStruct{}

					tmpEntry.caption = splitInstance[1]
					tmpEntry.description = splitInstance[2]
					tmpEntry.hotfixId = splitInstance[3]
					tmpEntry.installDate = splitInstance[4]

					c.PatchesList = append(c.PatchesList, tmpEntry)
				}
			}
		}
	}
	return (nil)
}
func (c *CMDB) Startup(rawOutput string) error {
	strSplit := strings.Split(rawOutput, "\r\n")

	for i, instance := range strSplit {
		if i != 0 {
			if instance != "" {
				instance = strings.Trim(instance, "\r\n ")
				splitInstance := strings.Split(instance, ",")

				if splitInstance[1] != "Location" {
					tmpEntry := StartupStruct{}

					tmpEntry.location = splitInstance[1]
					tmpEntry.name = splitInstance[2]

					c.StartupList = append(c.StartupList, tmpEntry)
				}
			}
		}
	}
	return (nil)
}

func (c *CMDB) Interfaces(rawOutput string) error {
	strSplit := strings.Split(rawOutput, "\r\n")

	for i, instance := range strSplit {
		if i != 0 {
			if instance != "" {
				instance = strings.Trim(instance, "\r\n ")
				splitInstance := strings.Split(instance, ",")

				if splitInstance[1] != "Index" {
					tmpEntry := InterfacesStruct{}
					tmpSplit := strings.Split(splitInstance[2], ";")

					if splitInstance[2] != "" || splitInstance[3] != "" {
						tmpEntry.index = splitInstance[1]
						tmpEntry.macAddr = splitInstance[3]

						for i := 0; i < len(tmpSplit); i++ {
							tmpSplit[i] = strings.Trim(tmpSplit[i], "{}")

							if i == 0 && len(tmpSplit) > 0 {
								tmpEntry.primaryIp = tmpSplit[0]
							} else if i > 1 && len(tmpSplit) > 0 {
								tmpEntry.configuredIp = append(tmpEntry.configuredIp, tmpSplit[i])
							} else if len(tmpSplit) == 0 {
								tmpEntry.primaryIp = splitInstance[2]
								tmpEntry.configuredIp = nil
							}
						}
						c.InterfacesList = append(c.InterfacesList, tmpEntry)
					}
				}
			}
		}
	}
	return (nil)
}
func (c *CMDB) ScheduledTask(rawOutput string) error {
	strSplit := strings.Split(rawOutput, "\r\n")

	for i, instance := range strSplit {
		if i != 0 {
			if instance != "" {
				instance = strings.Trim(instance, "\r\n ")
				splitInstance := strings.Split(instance, ",")

				if splitInstance[0] != "\"TaskName\"" {
					tmpEntry := ScheduledTask{}

					tmpEntry.task_name = splitInstance[0]
					tmpEntry.task_time = splitInstance[1]
					tmpEntry.task_status = splitInstance[2]

					c.ScheduledTaskList = append(c.ScheduledTaskList, tmpEntry)
				}
			}
		}
	}
	return (nil)
}
