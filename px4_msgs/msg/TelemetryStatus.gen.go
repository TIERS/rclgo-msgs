/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/telemetry_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/TelemetryStatus", TelemetryStatusTypeSupport)
}
const (
	TelemetryStatus_LINK_TYPE_GENERIC uint8 = 0
	TelemetryStatus_LINK_TYPE_UBIQUITY_BULLET uint8 = 1
	TelemetryStatus_LINK_TYPE_WIRE uint8 = 2
	TelemetryStatus_LINK_TYPE_USB uint8 = 3
	TelemetryStatus_LINK_TYPE_IRIDIUM uint8 = 4
	TelemetryStatus_HEARTBEAT_TIMEOUT_US uint64 = 1500000// Heartbeat timeout 1.5 seconds
)

// Do not create instances of this type directly. Always use NewTelemetryStatus
// function instead.
type TelemetryStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Type uint8 `yaml:"type"`// type of the radio hardware (LINK_TYPE_*)
	Mode uint8 `yaml:"mode"`
	FlowControl bool `yaml:"flow_control"`
	Forwarding bool `yaml:"forwarding"`
	MavlinkV2 bool `yaml:"mavlink_v2"`
	Ftp bool `yaml:"ftp"`
	Streams uint8 `yaml:"streams"`
	DataRate float32 `yaml:"data_rate"`// configured maximum data rate (Bytes/s)
	RateMultiplier float32 `yaml:"rate_multiplier"`
	TxRateAvg float32 `yaml:"tx_rate_avg"`// transmit rate average (Bytes/s)
	TxErrorRateAvg float32 `yaml:"tx_error_rate_avg"`// transmit error rate average (Bytes/s)
	TxMessageCount uint32 `yaml:"tx_message_count"`// total message sent count
	TxBufferOverruns uint32 `yaml:"tx_buffer_overruns"`// number of TX buffer overruns
	RxRateAvg float32 `yaml:"rx_rate_avg"`// transmit rate average (Bytes/s)
	RxMessageCount uint32 `yaml:"rx_message_count"`// count of total messages received
	RxMessageCountSupported uint32 `yaml:"rx_message_count_supported"`// count of total messages received from supported systems and components (for loss statistics)
	RxMessageLostCount uint32 `yaml:"rx_message_lost_count"`
	RxBufferOverruns uint32 `yaml:"rx_buffer_overruns"`// number of RX buffer overruns
	RxParseErrors uint32 `yaml:"rx_parse_errors"`// number of parse errors
	RxPacketDropCount uint32 `yaml:"rx_packet_drop_count"`// number of packet drops
	RxMessageLostRate float32 `yaml:"rx_message_lost_rate"`
	HeartbeatTypeAntennaTracker bool `yaml:"heartbeat_type_antenna_tracker"`// MAV_TYPE_ANTENNA_TRACKER. Heartbeats per type
	HeartbeatTypeGcs bool `yaml:"heartbeat_type_gcs"`// MAV_TYPE_GCS. Heartbeats per type
	HeartbeatTypeOnboardController bool `yaml:"heartbeat_type_onboard_controller"`// MAV_TYPE_ONBOARD_CONTROLLER. Heartbeats per type
	HeartbeatTypeGimbal bool `yaml:"heartbeat_type_gimbal"`// MAV_TYPE_GIMBAL. Heartbeats per type
	HeartbeatTypeAdsb bool `yaml:"heartbeat_type_adsb"`// MAV_TYPE_ADSB. Heartbeats per type
	HeartbeatTypeCamera bool `yaml:"heartbeat_type_camera"`// MAV_TYPE_CAMERA. Heartbeats per type
	HeartbeatComponentTelemetryRadio bool `yaml:"heartbeat_component_telemetry_radio"`// MAV_COMP_ID_TELEMETRY_RADIO. Heartbeats per component
	HeartbeatComponentLog bool `yaml:"heartbeat_component_log"`// MAV_COMP_ID_LOG. Heartbeats per component
	HeartbeatComponentOsd bool `yaml:"heartbeat_component_osd"`// MAV_COMP_ID_OSD. Heartbeats per component
	HeartbeatComponentObstacleAvoidance bool `yaml:"heartbeat_component_obstacle_avoidance"`// MAV_COMP_ID_OBSTACLE_AVOIDANCE. Heartbeats per component
	HeartbeatComponentVio bool `yaml:"heartbeat_component_vio"`// MAV_COMP_ID_VISUAL_INERTIAL_ODOMETRY. Heartbeats per component
	HeartbeatComponentPairingManager bool `yaml:"heartbeat_component_pairing_manager"`// MAV_COMP_ID_PAIRING_MANAGER. Heartbeats per component
	HeartbeatComponentUdpBridge bool `yaml:"heartbeat_component_udp_bridge"`// MAV_COMP_ID_UDP_BRIDGE. Heartbeats per component
	HeartbeatComponentUartBridge bool `yaml:"heartbeat_component_uart_bridge"`// MAV_COMP_ID_UART_BRIDGE. Heartbeats per component
	AvoidanceSystemHealthy bool `yaml:"avoidance_system_healthy"`// Misc component health
}

// NewTelemetryStatus creates a new TelemetryStatus with default values.
func NewTelemetryStatus() *TelemetryStatus {
	self := TelemetryStatus{}
	self.SetDefaults()
	return &self
}

func (t *TelemetryStatus) Clone() *TelemetryStatus {
	c := &TelemetryStatus{}
	c.Timestamp = t.Timestamp
	c.Type = t.Type
	c.Mode = t.Mode
	c.FlowControl = t.FlowControl
	c.Forwarding = t.Forwarding
	c.MavlinkV2 = t.MavlinkV2
	c.Ftp = t.Ftp
	c.Streams = t.Streams
	c.DataRate = t.DataRate
	c.RateMultiplier = t.RateMultiplier
	c.TxRateAvg = t.TxRateAvg
	c.TxErrorRateAvg = t.TxErrorRateAvg
	c.TxMessageCount = t.TxMessageCount
	c.TxBufferOverruns = t.TxBufferOverruns
	c.RxRateAvg = t.RxRateAvg
	c.RxMessageCount = t.RxMessageCount
	c.RxMessageCountSupported = t.RxMessageCountSupported
	c.RxMessageLostCount = t.RxMessageLostCount
	c.RxBufferOverruns = t.RxBufferOverruns
	c.RxParseErrors = t.RxParseErrors
	c.RxPacketDropCount = t.RxPacketDropCount
	c.RxMessageLostRate = t.RxMessageLostRate
	c.HeartbeatTypeAntennaTracker = t.HeartbeatTypeAntennaTracker
	c.HeartbeatTypeGcs = t.HeartbeatTypeGcs
	c.HeartbeatTypeOnboardController = t.HeartbeatTypeOnboardController
	c.HeartbeatTypeGimbal = t.HeartbeatTypeGimbal
	c.HeartbeatTypeAdsb = t.HeartbeatTypeAdsb
	c.HeartbeatTypeCamera = t.HeartbeatTypeCamera
	c.HeartbeatComponentTelemetryRadio = t.HeartbeatComponentTelemetryRadio
	c.HeartbeatComponentLog = t.HeartbeatComponentLog
	c.HeartbeatComponentOsd = t.HeartbeatComponentOsd
	c.HeartbeatComponentObstacleAvoidance = t.HeartbeatComponentObstacleAvoidance
	c.HeartbeatComponentVio = t.HeartbeatComponentVio
	c.HeartbeatComponentPairingManager = t.HeartbeatComponentPairingManager
	c.HeartbeatComponentUdpBridge = t.HeartbeatComponentUdpBridge
	c.HeartbeatComponentUartBridge = t.HeartbeatComponentUartBridge
	c.AvoidanceSystemHealthy = t.AvoidanceSystemHealthy
	return c
}

func (t *TelemetryStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *TelemetryStatus) SetDefaults() {
	t.Timestamp = 0
	t.Type = 0
	t.Mode = 0
	t.FlowControl = false
	t.Forwarding = false
	t.MavlinkV2 = false
	t.Ftp = false
	t.Streams = 0
	t.DataRate = 0
	t.RateMultiplier = 0
	t.TxRateAvg = 0
	t.TxErrorRateAvg = 0
	t.TxMessageCount = 0
	t.TxBufferOverruns = 0
	t.RxRateAvg = 0
	t.RxMessageCount = 0
	t.RxMessageCountSupported = 0
	t.RxMessageLostCount = 0
	t.RxBufferOverruns = 0
	t.RxParseErrors = 0
	t.RxPacketDropCount = 0
	t.RxMessageLostRate = 0
	t.HeartbeatTypeAntennaTracker = false
	t.HeartbeatTypeGcs = false
	t.HeartbeatTypeOnboardController = false
	t.HeartbeatTypeGimbal = false
	t.HeartbeatTypeAdsb = false
	t.HeartbeatTypeCamera = false
	t.HeartbeatComponentTelemetryRadio = false
	t.HeartbeatComponentLog = false
	t.HeartbeatComponentOsd = false
	t.HeartbeatComponentObstacleAvoidance = false
	t.HeartbeatComponentVio = false
	t.HeartbeatComponentPairingManager = false
	t.HeartbeatComponentUdpBridge = false
	t.HeartbeatComponentUartBridge = false
	t.AvoidanceSystemHealthy = false
}

// CloneTelemetryStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTelemetryStatusSlice(dst, src []TelemetryStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TelemetryStatusTypeSupport types.MessageTypeSupport = _TelemetryStatusTypeSupport{}

type _TelemetryStatusTypeSupport struct{}

func (t _TelemetryStatusTypeSupport) New() types.Message {
	return NewTelemetryStatus()
}

func (t _TelemetryStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__TelemetryStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__TelemetryStatus__create())
}

func (t _TelemetryStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__TelemetryStatus__destroy((*C.px4_msgs__msg__TelemetryStatus)(pointer_to_free))
}

func (t _TelemetryStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TelemetryStatus)
	mem := (*C.px4_msgs__msg__TelemetryStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem._type = C.uint8_t(m.Type)
	mem.mode = C.uint8_t(m.Mode)
	mem.flow_control = C.bool(m.FlowControl)
	mem.forwarding = C.bool(m.Forwarding)
	mem.mavlink_v2 = C.bool(m.MavlinkV2)
	mem.ftp = C.bool(m.Ftp)
	mem.streams = C.uint8_t(m.Streams)
	mem.data_rate = C.float(m.DataRate)
	mem.rate_multiplier = C.float(m.RateMultiplier)
	mem.tx_rate_avg = C.float(m.TxRateAvg)
	mem.tx_error_rate_avg = C.float(m.TxErrorRateAvg)
	mem.tx_message_count = C.uint32_t(m.TxMessageCount)
	mem.tx_buffer_overruns = C.uint32_t(m.TxBufferOverruns)
	mem.rx_rate_avg = C.float(m.RxRateAvg)
	mem.rx_message_count = C.uint32_t(m.RxMessageCount)
	mem.rx_message_count_supported = C.uint32_t(m.RxMessageCountSupported)
	mem.rx_message_lost_count = C.uint32_t(m.RxMessageLostCount)
	mem.rx_buffer_overruns = C.uint32_t(m.RxBufferOverruns)
	mem.rx_parse_errors = C.uint32_t(m.RxParseErrors)
	mem.rx_packet_drop_count = C.uint32_t(m.RxPacketDropCount)
	mem.rx_message_lost_rate = C.float(m.RxMessageLostRate)
	mem.heartbeat_type_antenna_tracker = C.bool(m.HeartbeatTypeAntennaTracker)
	mem.heartbeat_type_gcs = C.bool(m.HeartbeatTypeGcs)
	mem.heartbeat_type_onboard_controller = C.bool(m.HeartbeatTypeOnboardController)
	mem.heartbeat_type_gimbal = C.bool(m.HeartbeatTypeGimbal)
	mem.heartbeat_type_adsb = C.bool(m.HeartbeatTypeAdsb)
	mem.heartbeat_type_camera = C.bool(m.HeartbeatTypeCamera)
	mem.heartbeat_component_telemetry_radio = C.bool(m.HeartbeatComponentTelemetryRadio)
	mem.heartbeat_component_log = C.bool(m.HeartbeatComponentLog)
	mem.heartbeat_component_osd = C.bool(m.HeartbeatComponentOsd)
	mem.heartbeat_component_obstacle_avoidance = C.bool(m.HeartbeatComponentObstacleAvoidance)
	mem.heartbeat_component_vio = C.bool(m.HeartbeatComponentVio)
	mem.heartbeat_component_pairing_manager = C.bool(m.HeartbeatComponentPairingManager)
	mem.heartbeat_component_udp_bridge = C.bool(m.HeartbeatComponentUdpBridge)
	mem.heartbeat_component_uart_bridge = C.bool(m.HeartbeatComponentUartBridge)
	mem.avoidance_system_healthy = C.bool(m.AvoidanceSystemHealthy)
}

func (t _TelemetryStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TelemetryStatus)
	mem := (*C.px4_msgs__msg__TelemetryStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.Type = uint8(mem._type)
	m.Mode = uint8(mem.mode)
	m.FlowControl = bool(mem.flow_control)
	m.Forwarding = bool(mem.forwarding)
	m.MavlinkV2 = bool(mem.mavlink_v2)
	m.Ftp = bool(mem.ftp)
	m.Streams = uint8(mem.streams)
	m.DataRate = float32(mem.data_rate)
	m.RateMultiplier = float32(mem.rate_multiplier)
	m.TxRateAvg = float32(mem.tx_rate_avg)
	m.TxErrorRateAvg = float32(mem.tx_error_rate_avg)
	m.TxMessageCount = uint32(mem.tx_message_count)
	m.TxBufferOverruns = uint32(mem.tx_buffer_overruns)
	m.RxRateAvg = float32(mem.rx_rate_avg)
	m.RxMessageCount = uint32(mem.rx_message_count)
	m.RxMessageCountSupported = uint32(mem.rx_message_count_supported)
	m.RxMessageLostCount = uint32(mem.rx_message_lost_count)
	m.RxBufferOverruns = uint32(mem.rx_buffer_overruns)
	m.RxParseErrors = uint32(mem.rx_parse_errors)
	m.RxPacketDropCount = uint32(mem.rx_packet_drop_count)
	m.RxMessageLostRate = float32(mem.rx_message_lost_rate)
	m.HeartbeatTypeAntennaTracker = bool(mem.heartbeat_type_antenna_tracker)
	m.HeartbeatTypeGcs = bool(mem.heartbeat_type_gcs)
	m.HeartbeatTypeOnboardController = bool(mem.heartbeat_type_onboard_controller)
	m.HeartbeatTypeGimbal = bool(mem.heartbeat_type_gimbal)
	m.HeartbeatTypeAdsb = bool(mem.heartbeat_type_adsb)
	m.HeartbeatTypeCamera = bool(mem.heartbeat_type_camera)
	m.HeartbeatComponentTelemetryRadio = bool(mem.heartbeat_component_telemetry_radio)
	m.HeartbeatComponentLog = bool(mem.heartbeat_component_log)
	m.HeartbeatComponentOsd = bool(mem.heartbeat_component_osd)
	m.HeartbeatComponentObstacleAvoidance = bool(mem.heartbeat_component_obstacle_avoidance)
	m.HeartbeatComponentVio = bool(mem.heartbeat_component_vio)
	m.HeartbeatComponentPairingManager = bool(mem.heartbeat_component_pairing_manager)
	m.HeartbeatComponentUdpBridge = bool(mem.heartbeat_component_udp_bridge)
	m.HeartbeatComponentUartBridge = bool(mem.heartbeat_component_uart_bridge)
	m.AvoidanceSystemHealthy = bool(mem.avoidance_system_healthy)
}

func (t _TelemetryStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__TelemetryStatus())
}

type CTelemetryStatus = C.px4_msgs__msg__TelemetryStatus
type CTelemetryStatus__Sequence = C.px4_msgs__msg__TelemetryStatus__Sequence

func TelemetryStatus__Sequence_to_Go(goSlice *[]TelemetryStatus, cSlice CTelemetryStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TelemetryStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__TelemetryStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TelemetryStatus * uintptr(i)),
		))
		TelemetryStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func TelemetryStatus__Sequence_to_C(cSlice *CTelemetryStatus__Sequence, goSlice []TelemetryStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__TelemetryStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__TelemetryStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__TelemetryStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TelemetryStatus * uintptr(i)),
		))
		TelemetryStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func TelemetryStatus__Array_to_Go(goSlice []TelemetryStatus, cSlice []CTelemetryStatus) {
	for i := 0; i < len(cSlice); i++ {
		TelemetryStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TelemetryStatus__Array_to_C(cSlice []CTelemetryStatus, goSlice []TelemetryStatus) {
	for i := 0; i < len(goSlice); i++ {
		TelemetryStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
