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

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -llogging_demo__rosidl_typesupport_c -llogging_demo__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <logging_demo/srv/config_logger.h>
*/
import "C"

import (
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	"github.com/TIERS/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("logging_demo/ConfigLogger", ConfigLoggerTypeSupport)
}

type _ConfigLoggerTypeSupport struct {}

func (s _ConfigLoggerTypeSupport) Request() types.MessageTypeSupport {
	return ConfigLogger_RequestTypeSupport
}

func (s _ConfigLoggerTypeSupport) Response() types.MessageTypeSupport {
	return ConfigLogger_ResponseTypeSupport
}

func (s _ConfigLoggerTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__logging_demo__srv__ConfigLogger())
}

// Modifying this variable is undefined behavior.
var ConfigLoggerTypeSupport types.ServiceTypeSupport = _ConfigLoggerTypeSupport{}
