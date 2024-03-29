/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package turtlesim_srv
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lturtlesim__rosidl_typesupport_c -lturtlesim__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <turtlesim/srv/spawn.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("turtlesim/Spawn_Request", Spawn_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewSpawn_Request
// function instead.
type Spawn_Request struct {
	X float32 `yaml:"x"`
	Y float32 `yaml:"y"`
	Theta float32 `yaml:"theta"`
	Name string `yaml:"name"`// Optional.  A unique name will be created and returned if this is empty
}

// NewSpawn_Request creates a new Spawn_Request with default values.
func NewSpawn_Request() *Spawn_Request {
	self := Spawn_Request{}
	self.SetDefaults()
	return &self
}

func (t *Spawn_Request) Clone() *Spawn_Request {
	c := &Spawn_Request{}
	c.X = t.X
	c.Y = t.Y
	c.Theta = t.Theta
	c.Name = t.Name
	return c
}

func (t *Spawn_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Spawn_Request) SetDefaults() {
	t.X = 0
	t.Y = 0
	t.Theta = 0
	t.Name = ""
}

// CloneSpawn_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSpawn_RequestSlice(dst, src []Spawn_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Spawn_RequestTypeSupport types.MessageTypeSupport = _Spawn_RequestTypeSupport{}

type _Spawn_RequestTypeSupport struct{}

func (t _Spawn_RequestTypeSupport) New() types.Message {
	return NewSpawn_Request()
}

func (t _Spawn_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.turtlesim__srv__Spawn_Request
	return (unsafe.Pointer)(C.turtlesim__srv__Spawn_Request__create())
}

func (t _Spawn_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.turtlesim__srv__Spawn_Request__destroy((*C.turtlesim__srv__Spawn_Request)(pointer_to_free))
}

func (t _Spawn_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Spawn_Request)
	mem := (*C.turtlesim__srv__Spawn_Request)(dst)
	mem.x = C.float(m.X)
	mem.y = C.float(m.Y)
	mem.theta = C.float(m.Theta)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.name), m.Name)
}

func (t _Spawn_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Spawn_Request)
	mem := (*C.turtlesim__srv__Spawn_Request)(ros2_message_buffer)
	m.X = float32(mem.x)
	m.Y = float32(mem.y)
	m.Theta = float32(mem.theta)
	primitives.StringAsGoStruct(&m.Name, unsafe.Pointer(&mem.name))
}

func (t _Spawn_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__turtlesim__srv__Spawn_Request())
}

type CSpawn_Request = C.turtlesim__srv__Spawn_Request
type CSpawn_Request__Sequence = C.turtlesim__srv__Spawn_Request__Sequence

func Spawn_Request__Sequence_to_Go(goSlice *[]Spawn_Request, cSlice CSpawn_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Spawn_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.turtlesim__srv__Spawn_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__srv__Spawn_Request * uintptr(i)),
		))
		Spawn_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Spawn_Request__Sequence_to_C(cSlice *CSpawn_Request__Sequence, goSlice []Spawn_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.turtlesim__srv__Spawn_Request)(C.malloc((C.size_t)(C.sizeof_struct_turtlesim__srv__Spawn_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.turtlesim__srv__Spawn_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_turtlesim__srv__Spawn_Request * uintptr(i)),
		))
		Spawn_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Spawn_Request__Array_to_Go(goSlice []Spawn_Request, cSlice []CSpawn_Request) {
	for i := 0; i < len(cSlice); i++ {
		Spawn_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Spawn_Request__Array_to_C(cSlice []CSpawn_Request, goSlice []Spawn_Request) {
	for i := 0; i < len(goSlice); i++ {
		Spawn_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
