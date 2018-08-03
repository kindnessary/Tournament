// Code generated by mockery v1.0.0. DO NOT EDIT.
package controller

import entity "github.com/dmitriyomelyusik/Tournament/entity"
import mock "github.com/stretchr/testify/mock"

// MockDatabase is an autogenerated mock type for the Database type
type MockDatabase struct {
	mock.Mock
}

// CloseTournament provides a mock function with given fields: _a0
func (_m *MockDatabase) CloseTournament(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePlayer provides a mock function with given fields: _a0, _a1
func (_m *MockDatabase) CreatePlayer(_a0 string, _a1 int) (entity.Player, error) {
	ret := _m.Called(_a0, _a1)

	var r0 entity.Player
	if rf, ok := ret.Get(0).(func(string, int) entity.Player); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(entity.Player)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTournament provides a mock function with given fields: _a0, _a1
func (_m *MockDatabase) CreateTournament(_a0 string, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePlayer provides a mock function with given fields: _a0
func (_m *MockDatabase) DeletePlayer(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTournament provides a mock function with given fields: _a0
func (_m *MockDatabase) DeleteTournament(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetParticipants provides a mock function with given fields: _a0
func (_m *MockDatabase) GetParticipants(_a0 string) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlayer provides a mock function with given fields: _a0
func (_m *MockDatabase) GetPlayer(_a0 string) (entity.Player, error) {
	ret := _m.Called(_a0)

	var r0 entity.Player
	if rf, ok := ret.Get(0).(func(string) entity.Player); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(entity.Player)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTournamentState provides a mock function with given fields: _a0
func (_m *MockDatabase) GetTournamentState(_a0 string) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWinner provides a mock function with given fields: _a0
func (_m *MockDatabase) GetWinner(_a0 string) (entity.Winners, error) {
	ret := _m.Called(_a0)

	var r0 entity.Winners
	if rf, ok := ret.Get(0).(func(string) entity.Winners); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(entity.Winners)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTournamentWinner provides a mock function with given fields: _a0, _a1
func (_m *MockDatabase) SetTournamentWinner(_a0 string, _a1 entity.Winner) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, entity.Winner) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePlayer provides a mock function with given fields: _a0, _a1
func (_m *MockDatabase) UpdatePlayer(_a0 string, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTourAndPlayer provides a mock function with given fields: _a0, _a1
func (_m *MockDatabase) UpdateTourAndPlayer(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}