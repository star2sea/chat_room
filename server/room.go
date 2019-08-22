package main

import (
	"fmt"
	"strings"
)

type Room struct{
	name string
	members map[string]Connection
}

func NewRoom(name string) *Room{
	r := &Room{name:name, members:make(map[string]Connection)}
	return r
}

func (r *Room) EnterRoom(name string, conn Connection){
	r.CallMembers(fmt.Sprintf("%s enter room", name))
	r.members[name] = conn
	r.CallMember(name, r.GetMemberList())
}

func (r *Room) LeaveRoom(name string){
	delete(r.members, name)
	r.CallMembers(fmt.Sprintf("%s leave room", name))
}

func (r *Room) CallMember(name string, msg string){
	if conn, ok := r.members[name]; ok{
		conn.send(msg)
	}

}

func (r *Room) CallMembers(msg string){
	for name := range r.members{
		r.CallMember(name, msg)
	}
}

func (r *Room) GetMemberList() string{
	m := make([]string, len(r.members))
	for name := range r.members{
		m = append(m, name)
	}
	return strings.Join(m, "\r\n")
}

