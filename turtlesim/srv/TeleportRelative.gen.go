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
#include <turtlesim/srv/teleport_relative.h>
*/
import "C"

import (
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	"github.com/TIERS/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("turtlesim/TeleportRelative", TeleportRelativeTypeSupport)
}

type _TeleportRelativeTypeSupport struct {}

func (s _TeleportRelativeTypeSupport) Request() types.MessageTypeSupport {
	return TeleportRelative_RequestTypeSupport
}

func (s _TeleportRelativeTypeSupport) Response() types.MessageTypeSupport {
	return TeleportRelative_ResponseTypeSupport
}

func (s _TeleportRelativeTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__turtlesim__srv__TeleportRelative())
}

// Modifying this variable is undefined behavior.
var TeleportRelativeTypeSupport types.ServiceTypeSupport = _TeleportRelativeTypeSupport{}
