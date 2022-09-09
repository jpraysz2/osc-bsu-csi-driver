// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/driver/mount.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	luks "github.com/outscale-dev/osc-bsu-csi-driver/pkg/driver/luks"
	exec "k8s.io/utils/exec"
	mount "k8s.io/utils/mount"
)

// MockMounter is a mock of Mounter interface.
type MockMounter struct {
	ctrl     *gomock.Controller
	recorder *MockMounterMockRecorder
}

// MockMounterMockRecorder is the mock recorder for MockMounter.
type MockMounterMockRecorder struct {
	mock *MockMounter
}

// NewMockMounter creates a new mock instance.
func NewMockMounter(ctrl *gomock.Controller) *MockMounter {
	mock := &MockMounter{ctrl: ctrl}
	mock.recorder = &MockMounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMounter) EXPECT() *MockMounterMockRecorder {
	return m.recorder
}

// CheckLuksPassphrase mocks base method.
func (m *MockMounter) CheckLuksPassphrase(devicePath, passphrase string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckLuksPassphrase", devicePath, passphrase)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckLuksPassphrase indicates an expected call of CheckLuksPassphrase.
func (mr *MockMounterMockRecorder) CheckLuksPassphrase(devicePath, passphrase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckLuksPassphrase", reflect.TypeOf((*MockMounter)(nil).CheckLuksPassphrase), devicePath, passphrase)
}

// Command mocks base method.
func (m *MockMounter) Command(cmd string, args ...string) exec.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Command", varargs...)
	ret0, _ := ret[0].(exec.Cmd)
	return ret0
}

// Command indicates an expected call of Command.
func (mr *MockMounterMockRecorder) Command(cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockMounter)(nil).Command), varargs...)
}

// CommandContext mocks base method.
func (m *MockMounter) CommandContext(ctx context.Context, cmd string, args ...string) exec.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CommandContext", varargs...)
	ret0, _ := ret[0].(exec.Cmd)
	return ret0
}

// CommandContext indicates an expected call of CommandContext.
func (mr *MockMounterMockRecorder) CommandContext(ctx, cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandContext", reflect.TypeOf((*MockMounter)(nil).CommandContext), varargs...)
}

// ExistsPath mocks base method.
func (m *MockMounter) ExistsPath(filename string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsPath", filename)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsPath indicates an expected call of ExistsPath.
func (mr *MockMounterMockRecorder) ExistsPath(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsPath", reflect.TypeOf((*MockMounter)(nil).ExistsPath), filename)
}

// FormatAndMount mocks base method.
func (m *MockMounter) FormatAndMount(source, target, fstype string, options []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatAndMount", source, target, fstype, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// FormatAndMount indicates an expected call of FormatAndMount.
func (mr *MockMounterMockRecorder) FormatAndMount(source, target, fstype, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatAndMount", reflect.TypeOf((*MockMounter)(nil).FormatAndMount), source, target, fstype, options)
}

// GetDeviceName mocks base method.
func (m *MockMounter) GetDeviceName(mountPath string) (string, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceName", mountPath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDeviceName indicates an expected call of GetDeviceName.
func (mr *MockMounterMockRecorder) GetDeviceName(mountPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceName", reflect.TypeOf((*MockMounter)(nil).GetDeviceName), mountPath)
}

// GetDiskFormat mocks base method.
func (m *MockMounter) GetDiskFormat(disk string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskFormat", disk)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskFormat indicates an expected call of GetDiskFormat.
func (mr *MockMounterMockRecorder) GetDiskFormat(disk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskFormat", reflect.TypeOf((*MockMounter)(nil).GetDiskFormat), disk)
}

// GetMountRefs mocks base method.
func (m *MockMounter) GetMountRefs(pathname string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMountRefs", pathname)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMountRefs indicates an expected call of GetMountRefs.
func (mr *MockMounterMockRecorder) GetMountRefs(pathname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMountRefs", reflect.TypeOf((*MockMounter)(nil).GetMountRefs), pathname)
}

// IsCorruptedMnt mocks base method.
func (m *MockMounter) IsCorruptedMnt(arg0 error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCorruptedMnt", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCorruptedMnt indicates an expected call of IsCorruptedMnt.
func (mr *MockMounterMockRecorder) IsCorruptedMnt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCorruptedMnt", reflect.TypeOf((*MockMounter)(nil).IsCorruptedMnt), arg0)
}

// IsLikelyNotMountPoint mocks base method.
func (m *MockMounter) IsLikelyNotMountPoint(file string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLikelyNotMountPoint", file)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLikelyNotMountPoint indicates an expected call of IsLikelyNotMountPoint.
func (mr *MockMounterMockRecorder) IsLikelyNotMountPoint(file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLikelyNotMountPoint", reflect.TypeOf((*MockMounter)(nil).IsLikelyNotMountPoint), file)
}

// IsLuks mocks base method.
func (m *MockMounter) IsLuks(devicePath string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLuks", devicePath)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLuks indicates an expected call of IsLuks.
func (mr *MockMounterMockRecorder) IsLuks(devicePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLuks", reflect.TypeOf((*MockMounter)(nil).IsLuks), devicePath)
}

// IsLuksMapping mocks base method.
func (m *MockMounter) IsLuksMapping(devicePath string) (bool, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLuksMapping", devicePath)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IsLuksMapping indicates an expected call of IsLuksMapping.
func (mr *MockMounterMockRecorder) IsLuksMapping(devicePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLuksMapping", reflect.TypeOf((*MockMounter)(nil).IsLuksMapping), devicePath)
}

// List mocks base method.
func (m *MockMounter) List() ([]mount.MountPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]mount.MountPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMounterMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMounter)(nil).List))
}

// LookPath mocks base method.
func (m *MockMounter) LookPath(file string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookPath", file)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookPath indicates an expected call of LookPath.
func (mr *MockMounterMockRecorder) LookPath(file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookPath", reflect.TypeOf((*MockMounter)(nil).LookPath), file)
}

// LuksClose mocks base method.
func (m *MockMounter) LuksClose(deviceName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LuksClose", deviceName)
	ret0, _ := ret[0].(error)
	return ret0
}

// LuksClose indicates an expected call of LuksClose.
func (mr *MockMounterMockRecorder) LuksClose(deviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LuksClose", reflect.TypeOf((*MockMounter)(nil).LuksClose), deviceName)
}

// LuksFormat mocks base method.
func (m *MockMounter) LuksFormat(devicePath, passphrase string, context luks.LuksContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LuksFormat", devicePath, passphrase, context)
	ret0, _ := ret[0].(error)
	return ret0
}

// LuksFormat indicates an expected call of LuksFormat.
func (mr *MockMounterMockRecorder) LuksFormat(devicePath, passphrase, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LuksFormat", reflect.TypeOf((*MockMounter)(nil).LuksFormat), devicePath, passphrase, context)
}

// LuksOpen mocks base method.
func (m *MockMounter) LuksOpen(devicePath, encryptedDeviceName, passphrase string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LuksOpen", devicePath, encryptedDeviceName, passphrase)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LuksOpen indicates an expected call of LuksOpen.
func (mr *MockMounterMockRecorder) LuksOpen(devicePath, encryptedDeviceName, passphrase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LuksOpen", reflect.TypeOf((*MockMounter)(nil).LuksOpen), devicePath, encryptedDeviceName, passphrase)
}

// LuksResize mocks base method.
func (m *MockMounter) LuksResize(deviceName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LuksResize", deviceName)
	ret0, _ := ret[0].(error)
	return ret0
}

// LuksResize indicates an expected call of LuksResize.
func (mr *MockMounterMockRecorder) LuksResize(deviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LuksResize", reflect.TypeOf((*MockMounter)(nil).LuksResize), deviceName)
}

// MakeDir mocks base method.
func (m *MockMounter) MakeDir(pathname string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeDir", pathname)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeDir indicates an expected call of MakeDir.
func (mr *MockMounterMockRecorder) MakeDir(pathname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeDir", reflect.TypeOf((*MockMounter)(nil).MakeDir), pathname)
}

// MakeFile mocks base method.
func (m *MockMounter) MakeFile(pathname string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeFile", pathname)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeFile indicates an expected call of MakeFile.
func (mr *MockMounterMockRecorder) MakeFile(pathname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeFile", reflect.TypeOf((*MockMounter)(nil).MakeFile), pathname)
}

// Mount mocks base method.
func (m *MockMounter) Mount(source, target, fstype string, options []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mount", source, target, fstype, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mount indicates an expected call of Mount.
func (mr *MockMounterMockRecorder) Mount(source, target, fstype, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mount", reflect.TypeOf((*MockMounter)(nil).Mount), source, target, fstype, options)
}

// MountSensitive mocks base method.
func (m *MockMounter) MountSensitive(source, target, fstype string, options, sensitiveOptions []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MountSensitive", source, target, fstype, options, sensitiveOptions)
	ret0, _ := ret[0].(error)
	return ret0
}

// MountSensitive indicates an expected call of MountSensitive.
func (mr *MockMounterMockRecorder) MountSensitive(source, target, fstype, options, sensitiveOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountSensitive", reflect.TypeOf((*MockMounter)(nil).MountSensitive), source, target, fstype, options, sensitiveOptions)
}

// Unmount mocks base method.
func (m *MockMounter) Unmount(target string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmount", target)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmount indicates an expected call of Unmount.
func (mr *MockMounterMockRecorder) Unmount(target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmount", reflect.TypeOf((*MockMounter)(nil).Unmount), target)
}
