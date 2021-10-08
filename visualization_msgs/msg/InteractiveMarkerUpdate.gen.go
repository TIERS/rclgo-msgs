/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package visualization_msgs_msg
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lvisualization_msgs__rosidl_typesupport_c -lvisualization_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <visualization_msgs/msg/interactive_marker_update.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("visualization_msgs/InteractiveMarkerUpdate", InteractiveMarkerUpdateTypeSupport)
}
const (
	InteractiveMarkerUpdate_KEEP_ALIVE uint8 = 0// Type holds the purpose of this message.  It must be one of UPDATE or KEEP_ALIVE.UPDATE: Incremental update to previous state.The sequence number must be 1 higher than forthe previous update.KEEP_ALIVE: Indicates the that the server is still living.The sequence number does not increase.No payload data should be filled out (markers, poses, or erases).
	InteractiveMarkerUpdate_UPDATE uint8 = 1// Type holds the purpose of this message.  It must be one of UPDATE or KEEP_ALIVE.UPDATE: Incremental update to previous state.The sequence number must be 1 higher than forthe previous update.KEEP_ALIVE: Indicates the that the server is still living.The sequence number does not increase.No payload data should be filled out (markers, poses, or erases).
)

// Do not create instances of this type directly. Always use NewInteractiveMarkerUpdate
// function instead.
type InteractiveMarkerUpdate struct {
	ServerId string `yaml:"server_id"`// Identifying string. Must be unique in the topic namespacethat this server works on.
	SeqNum uint64 `yaml:"seq_num"`// Sequence number.The client will use this to detect if it has missed an update.
	Type uint8 `yaml:"type"`
	Markers []InteractiveMarker `yaml:"markers"`// Markers to be added or updated
	Poses []InteractiveMarkerPose `yaml:"poses"`// Poses of markers that should be moved
	Erases []string `yaml:"erases"`// Names of markers to be erased
}

// NewInteractiveMarkerUpdate creates a new InteractiveMarkerUpdate with default values.
func NewInteractiveMarkerUpdate() *InteractiveMarkerUpdate {
	self := InteractiveMarkerUpdate{}
	self.SetDefaults()
	return &self
}

func (t *InteractiveMarkerUpdate) Clone() *InteractiveMarkerUpdate {
	c := &InteractiveMarkerUpdate{}
	c.ServerId = t.ServerId
	c.SeqNum = t.SeqNum
	c.Type = t.Type
	if t.Markers != nil {
		c.Markers = make([]InteractiveMarker, len(t.Markers))
		CloneInteractiveMarkerSlice(c.Markers, t.Markers)
	}
	if t.Poses != nil {
		c.Poses = make([]InteractiveMarkerPose, len(t.Poses))
		CloneInteractiveMarkerPoseSlice(c.Poses, t.Poses)
	}
	if t.Erases != nil {
		c.Erases = make([]string, len(t.Erases))
		copy(c.Erases, t.Erases)
	}
	return c
}

func (t *InteractiveMarkerUpdate) CloneMsg() types.Message {
	return t.Clone()
}

func (t *InteractiveMarkerUpdate) SetDefaults() {
	t.ServerId = ""
	t.SeqNum = 0
	t.Type = 0
	t.Markers = nil
	t.Poses = nil
	t.Erases = nil
}

// CloneInteractiveMarkerUpdateSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInteractiveMarkerUpdateSlice(dst, src []InteractiveMarkerUpdate) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var InteractiveMarkerUpdateTypeSupport types.MessageTypeSupport = _InteractiveMarkerUpdateTypeSupport{}

type _InteractiveMarkerUpdateTypeSupport struct{}

func (t _InteractiveMarkerUpdateTypeSupport) New() types.Message {
	return NewInteractiveMarkerUpdate()
}

func (t _InteractiveMarkerUpdateTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__msg__InteractiveMarkerUpdate
	return (unsafe.Pointer)(C.visualization_msgs__msg__InteractiveMarkerUpdate__create())
}

func (t _InteractiveMarkerUpdateTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__msg__InteractiveMarkerUpdate__destroy((*C.visualization_msgs__msg__InteractiveMarkerUpdate)(pointer_to_free))
}

func (t _InteractiveMarkerUpdateTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*InteractiveMarkerUpdate)
	mem := (*C.visualization_msgs__msg__InteractiveMarkerUpdate)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.server_id), m.ServerId)
	mem.seq_num = C.uint64_t(m.SeqNum)
	mem._type = C.uint8_t(m.Type)
	InteractiveMarker__Sequence_to_C(&mem.markers, m.Markers)
	InteractiveMarkerPose__Sequence_to_C(&mem.poses, m.Poses)
	primitives.String__Sequence_to_C((*primitives.CString__Sequence)(unsafe.Pointer(&mem.erases)), m.Erases)
}

func (t _InteractiveMarkerUpdateTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*InteractiveMarkerUpdate)
	mem := (*C.visualization_msgs__msg__InteractiveMarkerUpdate)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.ServerId, unsafe.Pointer(&mem.server_id))
	m.SeqNum = uint64(mem.seq_num)
	m.Type = uint8(mem._type)
	InteractiveMarker__Sequence_to_Go(&m.Markers, mem.markers)
	InteractiveMarkerPose__Sequence_to_Go(&m.Poses, mem.poses)
	primitives.String__Sequence_to_Go(&m.Erases, *(*primitives.CString__Sequence)(unsafe.Pointer(&mem.erases)))
}

func (t _InteractiveMarkerUpdateTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__msg__InteractiveMarkerUpdate())
}

type CInteractiveMarkerUpdate = C.visualization_msgs__msg__InteractiveMarkerUpdate
type CInteractiveMarkerUpdate__Sequence = C.visualization_msgs__msg__InteractiveMarkerUpdate__Sequence

func InteractiveMarkerUpdate__Sequence_to_Go(goSlice *[]InteractiveMarkerUpdate, cSlice CInteractiveMarkerUpdate__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]InteractiveMarkerUpdate, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.visualization_msgs__msg__InteractiveMarkerUpdate__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerUpdate * uintptr(i)),
		))
		InteractiveMarkerUpdateTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func InteractiveMarkerUpdate__Sequence_to_C(cSlice *CInteractiveMarkerUpdate__Sequence, goSlice []InteractiveMarkerUpdate) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.visualization_msgs__msg__InteractiveMarkerUpdate)(C.malloc((C.size_t)(C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerUpdate * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.visualization_msgs__msg__InteractiveMarkerUpdate)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerUpdate * uintptr(i)),
		))
		InteractiveMarkerUpdateTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func InteractiveMarkerUpdate__Array_to_Go(goSlice []InteractiveMarkerUpdate, cSlice []CInteractiveMarkerUpdate) {
	for i := 0; i < len(cSlice); i++ {
		InteractiveMarkerUpdateTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func InteractiveMarkerUpdate__Array_to_C(cSlice []CInteractiveMarkerUpdate, goSlice []InteractiveMarkerUpdate) {
	for i := 0; i < len(goSlice); i++ {
		InteractiveMarkerUpdateTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
