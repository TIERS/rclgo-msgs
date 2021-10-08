/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package fog_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lfog_msgs__rosidl_typesupport_c -lfog_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <fog_msgs/srv/waypoint_to_local.h>
*/
import "C"

import (
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	"github.com/TIERS/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("fog_msgs/WaypointToLocal", WaypointToLocalTypeSupport)
}

type _WaypointToLocalTypeSupport struct {}

func (s _WaypointToLocalTypeSupport) Request() types.MessageTypeSupport {
	return WaypointToLocal_RequestTypeSupport
}

func (s _WaypointToLocalTypeSupport) Response() types.MessageTypeSupport {
	return WaypointToLocal_ResponseTypeSupport
}

func (s _WaypointToLocalTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__fog_msgs__srv__WaypointToLocal())
}

// Modifying this variable is undefined behavior.
var WaypointToLocalTypeSupport types.ServiceTypeSupport = _WaypointToLocalTypeSupport{}
