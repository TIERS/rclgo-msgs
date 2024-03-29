/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package map_msgs_srv
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/TIERS/rclgo-msgs/std_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lmap_msgs__rosidl_typesupport_c -lmap_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <map_msgs/srv/save_map.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("map_msgs/SaveMap_Request", SaveMap_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewSaveMap_Request
// function instead.
type SaveMap_Request struct {
	Filename std_msgs_msg.String `yaml:"filename"`// Save the map to the filesystem
}

// NewSaveMap_Request creates a new SaveMap_Request with default values.
func NewSaveMap_Request() *SaveMap_Request {
	self := SaveMap_Request{}
	self.SetDefaults()
	return &self
}

func (t *SaveMap_Request) Clone() *SaveMap_Request {
	c := &SaveMap_Request{}
	c.Filename = *t.Filename.Clone()
	return c
}

func (t *SaveMap_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SaveMap_Request) SetDefaults() {
	t.Filename.SetDefaults()
}

// CloneSaveMap_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSaveMap_RequestSlice(dst, src []SaveMap_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SaveMap_RequestTypeSupport types.MessageTypeSupport = _SaveMap_RequestTypeSupport{}

type _SaveMap_RequestTypeSupport struct{}

func (t _SaveMap_RequestTypeSupport) New() types.Message {
	return NewSaveMap_Request()
}

func (t _SaveMap_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.map_msgs__srv__SaveMap_Request
	return (unsafe.Pointer)(C.map_msgs__srv__SaveMap_Request__create())
}

func (t _SaveMap_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.map_msgs__srv__SaveMap_Request__destroy((*C.map_msgs__srv__SaveMap_Request)(pointer_to_free))
}

func (t _SaveMap_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SaveMap_Request)
	mem := (*C.map_msgs__srv__SaveMap_Request)(dst)
	std_msgs_msg.StringTypeSupport.AsCStruct(unsafe.Pointer(&mem.filename), &m.Filename)
}

func (t _SaveMap_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SaveMap_Request)
	mem := (*C.map_msgs__srv__SaveMap_Request)(ros2_message_buffer)
	std_msgs_msg.StringTypeSupport.AsGoStruct(&m.Filename, unsafe.Pointer(&mem.filename))
}

func (t _SaveMap_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__map_msgs__srv__SaveMap_Request())
}

type CSaveMap_Request = C.map_msgs__srv__SaveMap_Request
type CSaveMap_Request__Sequence = C.map_msgs__srv__SaveMap_Request__Sequence

func SaveMap_Request__Sequence_to_Go(goSlice *[]SaveMap_Request, cSlice CSaveMap_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SaveMap_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.map_msgs__srv__SaveMap_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__srv__SaveMap_Request * uintptr(i)),
		))
		SaveMap_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SaveMap_Request__Sequence_to_C(cSlice *CSaveMap_Request__Sequence, goSlice []SaveMap_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.map_msgs__srv__SaveMap_Request)(C.malloc((C.size_t)(C.sizeof_struct_map_msgs__srv__SaveMap_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.map_msgs__srv__SaveMap_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__srv__SaveMap_Request * uintptr(i)),
		))
		SaveMap_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SaveMap_Request__Array_to_Go(goSlice []SaveMap_Request, cSlice []CSaveMap_Request) {
	for i := 0; i < len(cSlice); i++ {
		SaveMap_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SaveMap_Request__Array_to_C(cSlice []CSaveMap_Request, goSlice []SaveMap_Request) {
	for i := 0; i < len(goSlice); i++ {
		SaveMap_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
