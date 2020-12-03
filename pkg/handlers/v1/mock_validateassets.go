// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/assetValidator.go

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	domain "github.com/asecurityteam/nexpose-asset-producer/pkg/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAssetValidator is a mock of AssetValidator interface
type MockAssetValidator struct {
	ctrl     *gomock.Controller
	recorder *MockAssetValidatorMockRecorder
}

// MockAssetValidatorMockRecorder is the mock recorder for MockAssetValidator
type MockAssetValidatorMockRecorder struct {
	mock *MockAssetValidator
}

// NewMockAssetValidator creates a new mock instance
func NewMockAssetValidator(ctrl *gomock.Controller) *MockAssetValidator {
	mock := &MockAssetValidator{ctrl: ctrl}
	mock.recorder = &MockAssetValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAssetValidator) EXPECT() *MockAssetValidatorMockRecorder {
	return m.recorder
}

// ValidateAssets mocks base method
func (m *MockAssetValidator) ValidateAssets(ctx context.Context, assets []domain.Asset, scanID string) ([]domain.AssetEvent, []error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAssets", ctx, assets, scanID)
	ret0, _ := ret[0].([]domain.AssetEvent)
	ret1, _ := ret[1].([]error)
	return ret0, ret1
}

// ValidateAssets indicates an expected call of ValidateAssets
func (mr *MockAssetValidatorMockRecorder) ValidateAssets(ctx, assets, scanID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAssets", reflect.TypeOf((*MockAssetValidator)(nil).ValidateAssets), ctx, assets, scanID)
}
