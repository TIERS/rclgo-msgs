/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package logging_demo_srv
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -llogging_demo__rosidl_typesupport_c -llogging_demo__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <logging_demo/srv/config_logger.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("logging_demo/ConfigLogger_Request", ConfigLogger_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewConfigLogger_Request
// function instead.
type ConfigLogger_Request struct {
	LoggerName string `yaml:"logger_name"`
	Level string `yaml:"level"`
}

// NewConfigLogger_Request creates a new ConfigLogger_Request with default values.
func NewConfigLogger_Request() *ConfigLogger_Request {
	self := ConfigLogger_Request{}
	self.SetDefaults()
	return &self
}

func (t *ConfigLogger_Request) Clone() *ConfigLogger_Request {
	c := &ConfigLogger_Request{}
	c.LoggerName = t.LoggerName
	c.Level = t.Level
	return c
}

func (t *ConfigLogger_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ConfigLogger_Request) SetDefaults() {
	t.LoggerName = ""
	t.Level = ""
}

// CloneConfigLogger_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneConfigLogger_RequestSlice(dst, src []ConfigLogger_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ConfigLogger_RequestTypeSupport types.MessageTypeSupport = _ConfigLogger_RequestTypeSupport{}

type _ConfigLogger_RequestTypeSupport struct{}

func (t _ConfigLogger_RequestTypeSupport) New() types.Message {
	return NewConfigLogger_Request()
}

func (t _ConfigLogger_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.logging_demo__srv__ConfigLogger_Request
	return (unsafe.Pointer)(C.logging_demo__srv__ConfigLogger_Request__create())
}

func (t _ConfigLogger_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.logging_demo__srv__ConfigLogger_Request__destroy((*C.logging_demo__srv__ConfigLogger_Request)(pointer_to_free))
}

func (t _ConfigLogger_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ConfigLogger_Request)
	mem := (*C.logging_demo__srv__ConfigLogger_Request)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.logger_name), m.LoggerName)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.level), m.Level)
}

func (t _ConfigLogger_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ConfigLogger_Request)
	mem := (*C.logging_demo__srv__ConfigLogger_Request)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.LoggerName, unsafe.Pointer(&mem.logger_name))
	primitives.StringAsGoStruct(&m.Level, unsafe.Pointer(&mem.level))
}

func (t _ConfigLogger_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__logging_demo__srv__ConfigLogger_Request())
}

type CConfigLogger_Request = C.logging_demo__srv__ConfigLogger_Request
type CConfigLogger_Request__Sequence = C.logging_demo__srv__ConfigLogger_Request__Sequence

func ConfigLogger_Request__Sequence_to_Go(goSlice *[]ConfigLogger_Request, cSlice CConfigLogger_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ConfigLogger_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.logging_demo__srv__ConfigLogger_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_logging_demo__srv__ConfigLogger_Request * uintptr(i)),
		))
		ConfigLogger_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ConfigLogger_Request__Sequence_to_C(cSlice *CConfigLogger_Request__Sequence, goSlice []ConfigLogger_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.logging_demo__srv__ConfigLogger_Request)(C.malloc((C.size_t)(C.sizeof_struct_logging_demo__srv__ConfigLogger_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.logging_demo__srv__ConfigLogger_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_logging_demo__srv__ConfigLogger_Request * uintptr(i)),
		))
		ConfigLogger_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ConfigLogger_Request__Array_to_Go(goSlice []ConfigLogger_Request, cSlice []CConfigLogger_Request) {
	for i := 0; i < len(cSlice); i++ {
		ConfigLogger_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ConfigLogger_Request__Array_to_C(cSlice []CConfigLogger_Request, goSlice []ConfigLogger_Request) {
	for i := 0; i < len(goSlice); i++ {
		ConfigLogger_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
