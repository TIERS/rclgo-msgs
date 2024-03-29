/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package nav_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <nav_msgs/srv/get_map.h>
*/
import "C"

import (
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	"github.com/TIERS/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("nav_msgs/GetMap", GetMapTypeSupport)
}

type _GetMapTypeSupport struct {}

func (s _GetMapTypeSupport) Request() types.MessageTypeSupport {
	return GetMap_RequestTypeSupport
}

func (s _GetMapTypeSupport) Response() types.MessageTypeSupport {
	return GetMap_ResponseTypeSupport
}

func (s _GetMapTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__nav_msgs__srv__GetMap())
}

// Modifying this variable is undefined behavior.
var GetMapTypeSupport types.ServiceTypeSupport = _GetMapTypeSupport{}
