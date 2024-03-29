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

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lturtlesim__rosidl_typesupport_c -lturtlesim__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <turtlesim/srv/spawn.h>
*/
import "C"

import (
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	"github.com/TIERS/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("turtlesim/Spawn", SpawnTypeSupport)
}

type _SpawnTypeSupport struct {}

func (s _SpawnTypeSupport) Request() types.MessageTypeSupport {
	return Spawn_RequestTypeSupport
}

func (s _SpawnTypeSupport) Response() types.MessageTypeSupport {
	return Spawn_ResponseTypeSupport
}

func (s _SpawnTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__turtlesim__srv__Spawn())
}

// Modifying this variable is undefined behavior.
var SpawnTypeSupport types.ServiceTypeSupport = _SpawnTypeSupport{}
