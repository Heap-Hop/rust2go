package main

import (
	"fmt"
	"time"
)

type Demo struct {
	lastUserName string
}

func init() {
	DemoCallImpl = &Demo{}
}

func (demo *Demo) demo_oneway(req DemoUser) {
	demo.lastUserName = req.name
	fmt.Printf("[Go-oneway] Golang received name: %s, age: %d\n", req.name, req.age)
}

func (demo *Demo) demo_check(req DemoComplicatedRequest) DemoResponse {
	fmt.Printf("[Go-call] Golang received req\n")
	fmt.Printf("[Go-call] Golang returned result\n")
	return DemoResponse{pass: true, last_request_user_name: demo.lastUserName}
}

func (demo *Demo) demo_check_async(req DemoComplicatedRequest) DemoResponse {
	fmt.Printf("[Go-call async] Golang received req, will sleep 1s\n")
	time.Sleep(1 * time.Second)
	fmt.Printf("[Go-call async] Golang returned result\n")
	return DemoResponse{pass: true}
}

func (demo *Demo) demo_check_async_safe(req DemoComplicatedRequest) DemoResponse {
	fmt.Printf("[Go-call async drop_safe] Golang received req, will sleep 1s\n")
	time.Sleep(1 * time.Second)
	resp := DemoResponse{pass: req.balabala[0] == 1}
	fmt.Printf("[Go-call async drop_safe] Golang returned result, pass: %v\n", req.balabala[0] == 1)
	return resp
}
