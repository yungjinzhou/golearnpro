package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var s vbStateRes
	args := `{"hostname":"", "state":1, "message":"successful", "hostAddr":"localhost", "machinePort":"9009", "machineIp":"172.0.19.2"}`
	err := json.Unmarshal([]byte(args), &s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

type vbStateRes struct {
	Hostname    mys    `json:"hostname"`
	State       int    `json:"state"`
	Message     string `json:"message"`
	HostAddr    string `json:"hostAddr"`
	MachinePort string `json:"machinePort"`
	MachineIp   string `json:"machineIp"`
}

type mys struct {
	Id string
}

func (rt *mys) UnmarshalJSON(b []byte) error {
	// Do nothing if data is the empty string.
	rt.Id = string(b)
	return nil
}