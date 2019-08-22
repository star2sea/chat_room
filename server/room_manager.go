package main

import (
	"fmt"
)

var roomMgr = &RoomManager{make(map[string]*Room)}

type RoomManager struct{
	rooms map[string]*Room
}

func (rm *RoomManager) CreateRoom(name string) (*Room, error){
	if rm.RoomExisted(name){
		return nil, fmt.Errorf("room %s existed, create failed", name)
	}
	r := NewRoom(name)
	rm.rooms[name] = r
	return r, nil
}

func (rm *RoomManager) DestroyRoom(name string){
	delete(rm.rooms, name)
}

func (rm *RoomManager) GetRoom(name string) *Room{
	return rm.rooms[name]
}

func (rm *RoomManager) RoomExisted(name string) bool{
	if _, ok := rm.rooms[name]; ok{
		return true
	}else{
		return false
	}
}