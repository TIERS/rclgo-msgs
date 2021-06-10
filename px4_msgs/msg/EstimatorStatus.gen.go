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
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/estimator_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/EstimatorStatus", EstimatorStatusTypeSupport)
}
const (
	EstimatorStatus_GPS_CHECK_FAIL_GPS_FIX uint8 = 0// 0 : insufficient fix type (no 3D solution). bits are true when corresponding test has failed
	EstimatorStatus_GPS_CHECK_FAIL_MIN_SAT_COUNT uint8 = 1// 1 : minimum required sat count fail. bits are true when corresponding test has failed
	EstimatorStatus_GPS_CHECK_FAIL_MAX_PDOP uint8 = 2// 2 : maximum allowed PDOP fail. bits are true when corresponding test has failed
	EstimatorStatus_GPS_CHECK_FAIL_MAX_HORZ_ERR uint8 = 3// 3 : maximum allowed horizontal position error fail. bits are true when corresponding test has failed
	EstimatorStatus_GPS_CHECK_FAIL_MAX_VERT_ERR uint8 = 4// 4 : maximum allowed vertical position error fail. bits are true when corresponding test has failed
	EstimatorStatus_GPS_CHECK_FAIL_MAX_SPD_ERR uint8 = 5// 5 : maximum allowed speed error fail. bits are true when corresponding test has failed
	EstimatorStatus_GPS_CHECK_FAIL_MAX_HORZ_DRIFT uint8 = 6// 6 : maximum allowed horizontal position drift fail - requires stationary vehicle. bits are true when corresponding test has failed
	EstimatorStatus_GPS_CHECK_FAIL_MAX_VERT_DRIFT uint8 = 7// 7 : maximum allowed vertical position drift fail - requires stationary vehicle. bits are true when corresponding test has failed
	EstimatorStatus_GPS_CHECK_FAIL_MAX_HORZ_SPD_ERR uint8 = 8// 8 : maximum allowed horizontal speed fail - requires stationary vehicle. bits are true when corresponding test has failed
	EstimatorStatus_GPS_CHECK_FAIL_MAX_VERT_SPD_ERR uint8 = 9// 9 : maximum allowed vertical velocity discrepancy fail. bits are true when corresponding test has failed
	EstimatorStatus_CS_TILT_ALIGN uint8 = 0// 0 - true if the filter tilt alignment is complete
	EstimatorStatus_CS_YAW_ALIGN uint8 = 1// 1 - true if the filter yaw alignment is complete
	EstimatorStatus_CS_GPS uint8 = 2// 2 - true if GPS measurements are being fused
	EstimatorStatus_CS_OPT_FLOW uint8 = 3// 3 - true if optical flow measurements are being fused
	EstimatorStatus_CS_MAG_HDG uint8 = 4// 4 - true if a simple magnetic yaw heading is being fused
	EstimatorStatus_CS_MAG_3D uint8 = 5// 5 - true if 3-axis magnetometer measurement are being fused
	EstimatorStatus_CS_MAG_DEC uint8 = 6// 6 - true if synthetic magnetic declination measurements are being fused
	EstimatorStatus_CS_IN_AIR uint8 = 7// 7 - true when thought to be airborne
	EstimatorStatus_CS_WIND uint8 = 8// 8 - true when wind velocity is being estimated
	EstimatorStatus_CS_BARO_HGT uint8 = 9// 9 - true when baro height is being fused as a primary height reference
	EstimatorStatus_CS_RNG_HGT uint8 = 10// 10 - true when range finder height is being fused as a primary height reference
	EstimatorStatus_CS_GPS_HGT uint8 = 11// 11 - true when GPS height is being fused as a primary height reference
	EstimatorStatus_CS_EV_POS uint8 = 12// 12 - true when local position data from external vision is being fused
	EstimatorStatus_CS_EV_YAW uint8 = 13// 13 - true when yaw data from external vision measurements is being fused
	EstimatorStatus_CS_EV_HGT uint8 = 14// 14 - true when height data from external vision measurements is being fused
	EstimatorStatus_CS_BETA uint8 = 15// 15 - true when synthetic sideslip measurements are being fused
	EstimatorStatus_CS_MAG_FIELD uint8 = 16// 16 - true when only the magnetic field states are updated by the magnetometer
	EstimatorStatus_CS_FIXED_WING uint8 = 17// 17 - true when thought to be operating as a fixed wing vehicle with constrained sideslip
	EstimatorStatus_CS_MAG_FAULT uint8 = 18// 18 - true when the magnetomer has been declared faulty and is no longer being used
	EstimatorStatus_CS_ASPD uint8 = 19// 19 - true when airspeed measurements are being fused
	EstimatorStatus_CS_GND_EFFECT uint8 = 20// 20 - true when when protection from ground effect induced static pressure rise is active
	EstimatorStatus_CS_RNG_STUCK uint8 = 21// 21 - true when a stuck range finder sensor has been detected
	EstimatorStatus_CS_GPS_YAW uint8 = 22// 22 - true when yaw (not ground course) data from a GPS receiver is being fused
	EstimatorStatus_CS_MAG_ALIGNED uint8 = 23// 23 - true when the in-flight mag field alignment has been completed
)

// Do not create instances of this type directly. Always use NewEstimatorStatus
// function instead.
type EstimatorStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	Vibe [3]float32 `yaml:"vibe"`// IMU vibration metrics in the following array locations
	OutputTrackingError [3]float32 `yaml:"output_tracking_error"`// return a vector containing the output predictor angular, velocity and position tracking error magnitudes (rad), (m/s), (m)
	GpsCheckFailFlags uint16 `yaml:"gps_check_fail_flags"`// Bitmask to indicate status of GPS checks - see definition below
	ControlModeFlags uint32 `yaml:"control_mode_flags"`// Bitmask to indicate EKF logic state
	FilterFaultFlags uint32 `yaml:"filter_fault_flags"`// Bitmask to indicate EKF internal faults
	PosHorizAccuracy float32 `yaml:"pos_horiz_accuracy"`// 1-Sigma estimated horizontal position accuracy relative to the estimators origin (m)
	PosVertAccuracy float32 `yaml:"pos_vert_accuracy"`// 1-Sigma estimated vertical position accuracy relative to the estimators origin (m)
	InnovationCheckFlags uint16 `yaml:"innovation_check_flags"`// Bitmask to indicate pass/fail status of innovation consistency checks
	MagTestRatio float32 `yaml:"mag_test_ratio"`// ratio of the largest magnetometer innovation component to the innovation test limit
	VelTestRatio float32 `yaml:"vel_test_ratio"`// ratio of the largest velocity innovation component to the innovation test limit
	PosTestRatio float32 `yaml:"pos_test_ratio"`// ratio of the largest horizontal position innovation component to the innovation test limit
	HgtTestRatio float32 `yaml:"hgt_test_ratio"`// ratio of the vertical position innovation to the innovation test limit
	TasTestRatio float32 `yaml:"tas_test_ratio"`// ratio of the true airspeed innovation to the innovation test limit
	HaglTestRatio float32 `yaml:"hagl_test_ratio"`// ratio of the height above ground innovation to the innovation test limit
	BetaTestRatio float32 `yaml:"beta_test_ratio"`// ratio of the synthetic sideslip innovation to the innovation test limit
	SolutionStatusFlags uint16 `yaml:"solution_status_flags"`// Bitmask indicating which filter kinematic state outputs are valid for flight control use.
	ResetCountVelNe uint8 `yaml:"reset_count_vel_ne"`// number of horizontal position reset events (allow to wrap if count exceeds 255)
	ResetCountVelD uint8 `yaml:"reset_count_vel_d"`// number of vertical velocity reset events (allow to wrap if count exceeds 255)
	ResetCountPosNe uint8 `yaml:"reset_count_pos_ne"`// number of horizontal position reset events (allow to wrap if count exceeds 255)
	ResetCountPodD uint8 `yaml:"reset_count_pod_d"`// number of vertical position reset events (allow to wrap if count exceeds 255)
	ResetCountQuat uint8 `yaml:"reset_count_quat"`// number of quaternion reset events (allow to wrap if count exceeds 255)
	TimeSlip float32 `yaml:"time_slip"`// cumulative amount of time in seconds that the EKF inertial calculation has slipped relative to system time
	PreFltFailInnovHeading bool `yaml:"pre_flt_fail_innov_heading"`
	PreFltFailInnovVelHoriz bool `yaml:"pre_flt_fail_innov_vel_horiz"`
	PreFltFailInnovVelVert bool `yaml:"pre_flt_fail_innov_vel_vert"`
	PreFltFailInnovHeight bool `yaml:"pre_flt_fail_innov_height"`
	PreFltFailMagFieldDisturbed bool `yaml:"pre_flt_fail_mag_field_disturbed"`
	AccelDeviceId uint32 `yaml:"accel_device_id"`
	GyroDeviceId uint32 `yaml:"gyro_device_id"`
	BaroDeviceId uint32 `yaml:"baro_device_id"`
	MagDeviceId uint32 `yaml:"mag_device_id"`
	HealthFlags uint8 `yaml:"health_flags"`// Bitmask to indicate sensor health states (vel, pos, hgt). legacy local position estimator (LPE) flags
	TimeoutFlags uint8 `yaml:"timeout_flags"`// Bitmask to indicate timeout flags (vel, pos, hgt). legacy local position estimator (LPE) flags
}

// NewEstimatorStatus creates a new EstimatorStatus with default values.
func NewEstimatorStatus() *EstimatorStatus {
	self := EstimatorStatus{}
	self.SetDefaults()
	return &self
}

func (t *EstimatorStatus) Clone() *EstimatorStatus {
	c := &EstimatorStatus{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.Vibe = t.Vibe
	c.OutputTrackingError = t.OutputTrackingError
	c.GpsCheckFailFlags = t.GpsCheckFailFlags
	c.ControlModeFlags = t.ControlModeFlags
	c.FilterFaultFlags = t.FilterFaultFlags
	c.PosHorizAccuracy = t.PosHorizAccuracy
	c.PosVertAccuracy = t.PosVertAccuracy
	c.InnovationCheckFlags = t.InnovationCheckFlags
	c.MagTestRatio = t.MagTestRatio
	c.VelTestRatio = t.VelTestRatio
	c.PosTestRatio = t.PosTestRatio
	c.HgtTestRatio = t.HgtTestRatio
	c.TasTestRatio = t.TasTestRatio
	c.HaglTestRatio = t.HaglTestRatio
	c.BetaTestRatio = t.BetaTestRatio
	c.SolutionStatusFlags = t.SolutionStatusFlags
	c.ResetCountVelNe = t.ResetCountVelNe
	c.ResetCountVelD = t.ResetCountVelD
	c.ResetCountPosNe = t.ResetCountPosNe
	c.ResetCountPodD = t.ResetCountPodD
	c.ResetCountQuat = t.ResetCountQuat
	c.TimeSlip = t.TimeSlip
	c.PreFltFailInnovHeading = t.PreFltFailInnovHeading
	c.PreFltFailInnovVelHoriz = t.PreFltFailInnovVelHoriz
	c.PreFltFailInnovVelVert = t.PreFltFailInnovVelVert
	c.PreFltFailInnovHeight = t.PreFltFailInnovHeight
	c.PreFltFailMagFieldDisturbed = t.PreFltFailMagFieldDisturbed
	c.AccelDeviceId = t.AccelDeviceId
	c.GyroDeviceId = t.GyroDeviceId
	c.BaroDeviceId = t.BaroDeviceId
	c.MagDeviceId = t.MagDeviceId
	c.HealthFlags = t.HealthFlags
	c.TimeoutFlags = t.TimeoutFlags
	return c
}

func (t *EstimatorStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *EstimatorStatus) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.Vibe = [3]float32{}
	t.OutputTrackingError = [3]float32{}
	t.GpsCheckFailFlags = 0
	t.ControlModeFlags = 0
	t.FilterFaultFlags = 0
	t.PosHorizAccuracy = 0
	t.PosVertAccuracy = 0
	t.InnovationCheckFlags = 0
	t.MagTestRatio = 0
	t.VelTestRatio = 0
	t.PosTestRatio = 0
	t.HgtTestRatio = 0
	t.TasTestRatio = 0
	t.HaglTestRatio = 0
	t.BetaTestRatio = 0
	t.SolutionStatusFlags = 0
	t.ResetCountVelNe = 0
	t.ResetCountVelD = 0
	t.ResetCountPosNe = 0
	t.ResetCountPodD = 0
	t.ResetCountQuat = 0
	t.TimeSlip = 0
	t.PreFltFailInnovHeading = false
	t.PreFltFailInnovVelHoriz = false
	t.PreFltFailInnovVelVert = false
	t.PreFltFailInnovHeight = false
	t.PreFltFailMagFieldDisturbed = false
	t.AccelDeviceId = 0
	t.GyroDeviceId = 0
	t.BaroDeviceId = 0
	t.MagDeviceId = 0
	t.HealthFlags = 0
	t.TimeoutFlags = 0
}

// CloneEstimatorStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneEstimatorStatusSlice(dst, src []EstimatorStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var EstimatorStatusTypeSupport types.MessageTypeSupport = _EstimatorStatusTypeSupport{}

type _EstimatorStatusTypeSupport struct{}

func (t _EstimatorStatusTypeSupport) New() types.Message {
	return NewEstimatorStatus()
}

func (t _EstimatorStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__EstimatorStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__EstimatorStatus__create())
}

func (t _EstimatorStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__EstimatorStatus__destroy((*C.px4_msgs__msg__EstimatorStatus)(pointer_to_free))
}

func (t _EstimatorStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*EstimatorStatus)
	mem := (*C.px4_msgs__msg__EstimatorStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	cSlice_vibe := mem.vibe[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_vibe)), m.Vibe[:])
	cSlice_output_tracking_error := mem.output_tracking_error[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_output_tracking_error)), m.OutputTrackingError[:])
	mem.gps_check_fail_flags = C.uint16_t(m.GpsCheckFailFlags)
	mem.control_mode_flags = C.uint32_t(m.ControlModeFlags)
	mem.filter_fault_flags = C.uint32_t(m.FilterFaultFlags)
	mem.pos_horiz_accuracy = C.float(m.PosHorizAccuracy)
	mem.pos_vert_accuracy = C.float(m.PosVertAccuracy)
	mem.innovation_check_flags = C.uint16_t(m.InnovationCheckFlags)
	mem.mag_test_ratio = C.float(m.MagTestRatio)
	mem.vel_test_ratio = C.float(m.VelTestRatio)
	mem.pos_test_ratio = C.float(m.PosTestRatio)
	mem.hgt_test_ratio = C.float(m.HgtTestRatio)
	mem.tas_test_ratio = C.float(m.TasTestRatio)
	mem.hagl_test_ratio = C.float(m.HaglTestRatio)
	mem.beta_test_ratio = C.float(m.BetaTestRatio)
	mem.solution_status_flags = C.uint16_t(m.SolutionStatusFlags)
	mem.reset_count_vel_ne = C.uint8_t(m.ResetCountVelNe)
	mem.reset_count_vel_d = C.uint8_t(m.ResetCountVelD)
	mem.reset_count_pos_ne = C.uint8_t(m.ResetCountPosNe)
	mem.reset_count_pod_d = C.uint8_t(m.ResetCountPodD)
	mem.reset_count_quat = C.uint8_t(m.ResetCountQuat)
	mem.time_slip = C.float(m.TimeSlip)
	mem.pre_flt_fail_innov_heading = C.bool(m.PreFltFailInnovHeading)
	mem.pre_flt_fail_innov_vel_horiz = C.bool(m.PreFltFailInnovVelHoriz)
	mem.pre_flt_fail_innov_vel_vert = C.bool(m.PreFltFailInnovVelVert)
	mem.pre_flt_fail_innov_height = C.bool(m.PreFltFailInnovHeight)
	mem.pre_flt_fail_mag_field_disturbed = C.bool(m.PreFltFailMagFieldDisturbed)
	mem.accel_device_id = C.uint32_t(m.AccelDeviceId)
	mem.gyro_device_id = C.uint32_t(m.GyroDeviceId)
	mem.baro_device_id = C.uint32_t(m.BaroDeviceId)
	mem.mag_device_id = C.uint32_t(m.MagDeviceId)
	mem.health_flags = C.uint8_t(m.HealthFlags)
	mem.timeout_flags = C.uint8_t(m.TimeoutFlags)
}

func (t _EstimatorStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*EstimatorStatus)
	mem := (*C.px4_msgs__msg__EstimatorStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_vibe := mem.vibe[:]
	primitives.Float32__Array_to_Go(m.Vibe[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_vibe)))
	cSlice_output_tracking_error := mem.output_tracking_error[:]
	primitives.Float32__Array_to_Go(m.OutputTrackingError[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_output_tracking_error)))
	m.GpsCheckFailFlags = uint16(mem.gps_check_fail_flags)
	m.ControlModeFlags = uint32(mem.control_mode_flags)
	m.FilterFaultFlags = uint32(mem.filter_fault_flags)
	m.PosHorizAccuracy = float32(mem.pos_horiz_accuracy)
	m.PosVertAccuracy = float32(mem.pos_vert_accuracy)
	m.InnovationCheckFlags = uint16(mem.innovation_check_flags)
	m.MagTestRatio = float32(mem.mag_test_ratio)
	m.VelTestRatio = float32(mem.vel_test_ratio)
	m.PosTestRatio = float32(mem.pos_test_ratio)
	m.HgtTestRatio = float32(mem.hgt_test_ratio)
	m.TasTestRatio = float32(mem.tas_test_ratio)
	m.HaglTestRatio = float32(mem.hagl_test_ratio)
	m.BetaTestRatio = float32(mem.beta_test_ratio)
	m.SolutionStatusFlags = uint16(mem.solution_status_flags)
	m.ResetCountVelNe = uint8(mem.reset_count_vel_ne)
	m.ResetCountVelD = uint8(mem.reset_count_vel_d)
	m.ResetCountPosNe = uint8(mem.reset_count_pos_ne)
	m.ResetCountPodD = uint8(mem.reset_count_pod_d)
	m.ResetCountQuat = uint8(mem.reset_count_quat)
	m.TimeSlip = float32(mem.time_slip)
	m.PreFltFailInnovHeading = bool(mem.pre_flt_fail_innov_heading)
	m.PreFltFailInnovVelHoriz = bool(mem.pre_flt_fail_innov_vel_horiz)
	m.PreFltFailInnovVelVert = bool(mem.pre_flt_fail_innov_vel_vert)
	m.PreFltFailInnovHeight = bool(mem.pre_flt_fail_innov_height)
	m.PreFltFailMagFieldDisturbed = bool(mem.pre_flt_fail_mag_field_disturbed)
	m.AccelDeviceId = uint32(mem.accel_device_id)
	m.GyroDeviceId = uint32(mem.gyro_device_id)
	m.BaroDeviceId = uint32(mem.baro_device_id)
	m.MagDeviceId = uint32(mem.mag_device_id)
	m.HealthFlags = uint8(mem.health_flags)
	m.TimeoutFlags = uint8(mem.timeout_flags)
}

func (t _EstimatorStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__EstimatorStatus())
}

type CEstimatorStatus = C.px4_msgs__msg__EstimatorStatus
type CEstimatorStatus__Sequence = C.px4_msgs__msg__EstimatorStatus__Sequence

func EstimatorStatus__Sequence_to_Go(goSlice *[]EstimatorStatus, cSlice CEstimatorStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]EstimatorStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__EstimatorStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorStatus * uintptr(i)),
		))
		EstimatorStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func EstimatorStatus__Sequence_to_C(cSlice *CEstimatorStatus__Sequence, goSlice []EstimatorStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__EstimatorStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__EstimatorStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__EstimatorStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorStatus * uintptr(i)),
		))
		EstimatorStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func EstimatorStatus__Array_to_Go(goSlice []EstimatorStatus, cSlice []CEstimatorStatus) {
	for i := 0; i < len(cSlice); i++ {
		EstimatorStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func EstimatorStatus__Array_to_C(cSlice []CEstimatorStatus, goSlice []EstimatorStatus) {
	for i := 0; i < len(goSlice); i++ {
		EstimatorStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
