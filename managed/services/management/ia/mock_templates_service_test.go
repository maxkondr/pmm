// Code generated by mockery v1.0.0. DO NOT EDIT.

package ia

import (
	mock "github.com/stretchr/testify/mock"

	alerting "github.com/percona/pmm/managed/services/management/alerting"
)

// mockTemplatesService is an autogenerated mock type for the templatesService type
type mockTemplatesService struct {
	mock.Mock
}

// GetTemplates provides a mock function with given fields:
func (_m *mockTemplatesService) GetTemplates() map[string]alerting.TemplateInfo {
	ret := _m.Called()

	var r0 map[string]alerting.TemplateInfo
	if rf, ok := ret.Get(0).(func() map[string]alerting.TemplateInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]alerting.TemplateInfo)
		}
	}

	return r0
}
