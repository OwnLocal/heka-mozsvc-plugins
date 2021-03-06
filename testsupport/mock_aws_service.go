// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/AdRoll/goamz/aws (interfaces: AWSService)

package testsupport

import (
	gomock "github.com/rafrombrc/gomock/gomock"
	http "net/http"
)

// Mock of AWSService interface
type MockAWSService struct {
	ctrl     *gomock.Controller
	recorder *_MockAWSServiceRecorder
}

// Recorder for MockAWSService (not exported)
type _MockAWSServiceRecorder struct {
	mock *MockAWSService
}

func NewMockAWSService(ctrl *gomock.Controller) *MockAWSService {
	mock := &MockAWSService{ctrl: ctrl}
	mock.recorder = &_MockAWSServiceRecorder{mock}
	return mock
}

func (_m *MockAWSService) EXPECT() *_MockAWSServiceRecorder {
	return _m.recorder
}

func (_m *MockAWSService) BuildError(_param0 *http.Response) error {
	ret := _m.ctrl.Call(_m, "BuildError", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAWSServiceRecorder) BuildError(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BuildError", arg0)
}

func (_m *MockAWSService) Query(_param0 string, _param1 string, _param2 map[string]string) (*http.Response, error) {
	ret := _m.ctrl.Call(_m, "Query", _param0, _param1, _param2)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAWSServiceRecorder) Query(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Query", arg0, arg1, arg2)
}
