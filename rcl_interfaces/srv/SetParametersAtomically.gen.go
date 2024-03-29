/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package rcl_interfaces_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lrcl_interfaces__rosidl_typesupport_c -lrcl_interfaces__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <rcl_interfaces/srv/set_parameters_atomically.h>
*/
import "C"

import (
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	"github.com/TIERS/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("rcl_interfaces/SetParametersAtomically", SetParametersAtomicallyTypeSupport)
}

type _SetParametersAtomicallyTypeSupport struct {}

func (s _SetParametersAtomicallyTypeSupport) Request() types.MessageTypeSupport {
	return SetParametersAtomically_RequestTypeSupport
}

func (s _SetParametersAtomicallyTypeSupport) Response() types.MessageTypeSupport {
	return SetParametersAtomically_ResponseTypeSupport
}

func (s _SetParametersAtomicallyTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__rcl_interfaces__srv__SetParametersAtomically())
}

// Modifying this variable is undefined behavior.
var SetParametersAtomicallyTypeSupport types.ServiceTypeSupport = _SetParametersAtomicallyTypeSupport{}
