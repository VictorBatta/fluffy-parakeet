// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package extensions

import (
	"sync"
)

// Ensure, that ExtensionMock does implement Extension.
// If this is not the case, regenerate this file with moq.
var _ Extension = &ExtensionMock{}

// ExtensionMock is a mock implementation of Extension.
//
// 	func TestSomethingThatUsesExtension(t *testing.T) {
//
// 		// make and configure a mocked Extension
// 		mockedExtension := &ExtensionMock{
// 			IsLocalFunc: func() bool {
// 				panic("mock out the IsLocal method")
// 			},
// 			NameFunc: func() string {
// 				panic("mock out the Name method")
// 			},
// 			PathFunc: func() string {
// 				panic("mock out the Path method")
// 			},
// 			URLFunc: func() string {
// 				panic("mock out the URL method")
// 			},
// 			UpdateAvailableFunc: func() bool {
// 				panic("mock out the UpdateAvailable method")
// 			},
// 		}
//
// 		// use mockedExtension in code that requires Extension
// 		// and then make assertions.
//
// 	}
type ExtensionMock struct {
	// IsLocalFunc mocks the IsLocal method.
	IsLocalFunc func() bool

	// NameFunc mocks the Name method.
	NameFunc func() string

	// PathFunc mocks the Path method.
	PathFunc func() string

	// URLFunc mocks the URL method.
	URLFunc func() string

	// UpdateAvailableFunc mocks the UpdateAvailable method.
	UpdateAvailableFunc func() bool

	// calls tracks calls to the methods.
	calls struct {
		// IsLocal holds details about calls to the IsLocal method.
		IsLocal []struct {
		}
		// Name holds details about calls to the Name method.
		Name []struct {
		}
		// Path holds details about calls to the Path method.
		Path []struct {
		}
		// URL holds details about calls to the URL method.
		URL []struct {
		}
		// UpdateAvailable holds details about calls to the UpdateAvailable method.
		UpdateAvailable []struct {
		}
	}
	lockIsLocal         sync.RWMutex
	lockName            sync.RWMutex
	lockPath            sync.RWMutex
	lockURL             sync.RWMutex
	lockUpdateAvailable sync.RWMutex
}

// IsLocal calls IsLocalFunc.
func (mock *ExtensionMock) IsLocal() bool {
	if mock.IsLocalFunc == nil {
		panic("ExtensionMock.IsLocalFunc: method is nil but Extension.IsLocal was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsLocal.Lock()
	mock.calls.IsLocal = append(mock.calls.IsLocal, callInfo)
	mock.lockIsLocal.Unlock()
	return mock.IsLocalFunc()
}

// IsLocalCalls gets all the calls that were made to IsLocal.
// Check the length with:
//     len(mockedExtension.IsLocalCalls())
func (mock *ExtensionMock) IsLocalCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsLocal.RLock()
	calls = mock.calls.IsLocal
	mock.lockIsLocal.RUnlock()
	return calls
}

// Name calls NameFunc.
func (mock *ExtensionMock) Name() string {
	if mock.NameFunc == nil {
		panic("ExtensionMock.NameFunc: method is nil but Extension.Name was just called")
	}
	callInfo := struct {
	}{}
	mock.lockName.Lock()
	mock.calls.Name = append(mock.calls.Name, callInfo)
	mock.lockName.Unlock()
	return mock.NameFunc()
}

// NameCalls gets all the calls that were made to Name.
// Check the length with:
//     len(mockedExtension.NameCalls())
func (mock *ExtensionMock) NameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockName.RLock()
	calls = mock.calls.Name
	mock.lockName.RUnlock()
	return calls
}

// Path calls PathFunc.
func (mock *ExtensionMock) Path() string {
	if mock.PathFunc == nil {
		panic("ExtensionMock.PathFunc: method is nil but Extension.Path was just called")
	}
	callInfo := struct {
	}{}
	mock.lockPath.Lock()
	mock.calls.Path = append(mock.calls.Path, callInfo)
	mock.lockPath.Unlock()
	return mock.PathFunc()
}

// PathCalls gets all the calls that were made to Path.
// Check the length with:
//     len(mockedExtension.PathCalls())
func (mock *ExtensionMock) PathCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockPath.RLock()
	calls = mock.calls.Path
	mock.lockPath.RUnlock()
	return calls
}

// URL calls URLFunc.
func (mock *ExtensionMock) URL() string {
	if mock.URLFunc == nil {
		panic("ExtensionMock.URLFunc: method is nil but Extension.URL was just called")
	}
	callInfo := struct {
	}{}
	mock.lockURL.Lock()
	mock.calls.URL = append(mock.calls.URL, callInfo)
	mock.lockURL.Unlock()
	return mock.URLFunc()
}

// URLCalls gets all the calls that were made to URL.
// Check the length with:
//     len(mockedExtension.URLCalls())
func (mock *ExtensionMock) URLCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockURL.RLock()
	calls = mock.calls.URL
	mock.lockURL.RUnlock()
	return calls
}

// UpdateAvailable calls UpdateAvailableFunc.
func (mock *ExtensionMock) UpdateAvailable() bool {
	if mock.UpdateAvailableFunc == nil {
		panic("ExtensionMock.UpdateAvailableFunc: method is nil but Extension.UpdateAvailable was just called")
	}
	callInfo := struct {
	}{}
	mock.lockUpdateAvailable.Lock()
	mock.calls.UpdateAvailable = append(mock.calls.UpdateAvailable, callInfo)
	mock.lockUpdateAvailable.Unlock()
	return mock.UpdateAvailableFunc()
}

// UpdateAvailableCalls gets all the calls that were made to UpdateAvailable.
// Check the length with:
//     len(mockedExtension.UpdateAvailableCalls())
func (mock *ExtensionMock) UpdateAvailableCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockUpdateAvailable.RLock()
	calls = mock.calls.UpdateAvailable
	mock.lockUpdateAvailable.RUnlock()
	return calls
}
