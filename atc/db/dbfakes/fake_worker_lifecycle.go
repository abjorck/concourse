// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/db"
)

type FakeWorkerLifecycle struct {
	DeleteFinishedRetiringWorkersStub        func() ([]string, error)
	deleteFinishedRetiringWorkersMutex       sync.RWMutex
	deleteFinishedRetiringWorkersArgsForCall []struct {
	}
	deleteFinishedRetiringWorkersReturns struct {
		result1 []string
		result2 error
	}
	deleteFinishedRetiringWorkersReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	DeleteUnresponsiveEphemeralWorkersStub        func() ([]string, error)
	deleteUnresponsiveEphemeralWorkersMutex       sync.RWMutex
	deleteUnresponsiveEphemeralWorkersArgsForCall []struct {
	}
	deleteUnresponsiveEphemeralWorkersReturns struct {
		result1 []string
		result2 error
	}
	deleteUnresponsiveEphemeralWorkersReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	GetWorkerStateByNameStub        func() (map[string]db.WorkerState, error)
	getWorkerStateByNameMutex       sync.RWMutex
	getWorkerStateByNameArgsForCall []struct {
	}
	getWorkerStateByNameReturns struct {
		result1 map[string]db.WorkerState
		result2 error
	}
	getWorkerStateByNameReturnsOnCall map[int]struct {
		result1 map[string]db.WorkerState
		result2 error
	}
	LandFinishedLandingWorkersStub        func() ([]string, error)
	landFinishedLandingWorkersMutex       sync.RWMutex
	landFinishedLandingWorkersArgsForCall []struct {
	}
	landFinishedLandingWorkersReturns struct {
		result1 []string
		result2 error
	}
	landFinishedLandingWorkersReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	StallUnresponsiveWorkersStub        func() ([]string, error)
	stallUnresponsiveWorkersMutex       sync.RWMutex
	stallUnresponsiveWorkersArgsForCall []struct {
	}
	stallUnresponsiveWorkersReturns struct {
		result1 []string
		result2 error
	}
	stallUnresponsiveWorkersReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorkerLifecycle) DeleteFinishedRetiringWorkers() ([]string, error) {
	fake.deleteFinishedRetiringWorkersMutex.Lock()
	ret, specificReturn := fake.deleteFinishedRetiringWorkersReturnsOnCall[len(fake.deleteFinishedRetiringWorkersArgsForCall)]
	fake.deleteFinishedRetiringWorkersArgsForCall = append(fake.deleteFinishedRetiringWorkersArgsForCall, struct {
	}{})
	stub := fake.DeleteFinishedRetiringWorkersStub
	fakeReturns := fake.deleteFinishedRetiringWorkersReturns
	fake.recordInvocation("DeleteFinishedRetiringWorkers", []interface{}{})
	fake.deleteFinishedRetiringWorkersMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkerLifecycle) DeleteFinishedRetiringWorkersCallCount() int {
	fake.deleteFinishedRetiringWorkersMutex.RLock()
	defer fake.deleteFinishedRetiringWorkersMutex.RUnlock()
	return len(fake.deleteFinishedRetiringWorkersArgsForCall)
}

func (fake *FakeWorkerLifecycle) DeleteFinishedRetiringWorkersCalls(stub func() ([]string, error)) {
	fake.deleteFinishedRetiringWorkersMutex.Lock()
	defer fake.deleteFinishedRetiringWorkersMutex.Unlock()
	fake.DeleteFinishedRetiringWorkersStub = stub
}

func (fake *FakeWorkerLifecycle) DeleteFinishedRetiringWorkersReturns(result1 []string, result2 error) {
	fake.deleteFinishedRetiringWorkersMutex.Lock()
	defer fake.deleteFinishedRetiringWorkersMutex.Unlock()
	fake.DeleteFinishedRetiringWorkersStub = nil
	fake.deleteFinishedRetiringWorkersReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerLifecycle) DeleteFinishedRetiringWorkersReturnsOnCall(i int, result1 []string, result2 error) {
	fake.deleteFinishedRetiringWorkersMutex.Lock()
	defer fake.deleteFinishedRetiringWorkersMutex.Unlock()
	fake.DeleteFinishedRetiringWorkersStub = nil
	if fake.deleteFinishedRetiringWorkersReturnsOnCall == nil {
		fake.deleteFinishedRetiringWorkersReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.deleteFinishedRetiringWorkersReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerLifecycle) DeleteUnresponsiveEphemeralWorkers() ([]string, error) {
	fake.deleteUnresponsiveEphemeralWorkersMutex.Lock()
	ret, specificReturn := fake.deleteUnresponsiveEphemeralWorkersReturnsOnCall[len(fake.deleteUnresponsiveEphemeralWorkersArgsForCall)]
	fake.deleteUnresponsiveEphemeralWorkersArgsForCall = append(fake.deleteUnresponsiveEphemeralWorkersArgsForCall, struct {
	}{})
	stub := fake.DeleteUnresponsiveEphemeralWorkersStub
	fakeReturns := fake.deleteUnresponsiveEphemeralWorkersReturns
	fake.recordInvocation("DeleteUnresponsiveEphemeralWorkers", []interface{}{})
	fake.deleteUnresponsiveEphemeralWorkersMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkerLifecycle) DeleteUnresponsiveEphemeralWorkersCallCount() int {
	fake.deleteUnresponsiveEphemeralWorkersMutex.RLock()
	defer fake.deleteUnresponsiveEphemeralWorkersMutex.RUnlock()
	return len(fake.deleteUnresponsiveEphemeralWorkersArgsForCall)
}

func (fake *FakeWorkerLifecycle) DeleteUnresponsiveEphemeralWorkersCalls(stub func() ([]string, error)) {
	fake.deleteUnresponsiveEphemeralWorkersMutex.Lock()
	defer fake.deleteUnresponsiveEphemeralWorkersMutex.Unlock()
	fake.DeleteUnresponsiveEphemeralWorkersStub = stub
}

func (fake *FakeWorkerLifecycle) DeleteUnresponsiveEphemeralWorkersReturns(result1 []string, result2 error) {
	fake.deleteUnresponsiveEphemeralWorkersMutex.Lock()
	defer fake.deleteUnresponsiveEphemeralWorkersMutex.Unlock()
	fake.DeleteUnresponsiveEphemeralWorkersStub = nil
	fake.deleteUnresponsiveEphemeralWorkersReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerLifecycle) DeleteUnresponsiveEphemeralWorkersReturnsOnCall(i int, result1 []string, result2 error) {
	fake.deleteUnresponsiveEphemeralWorkersMutex.Lock()
	defer fake.deleteUnresponsiveEphemeralWorkersMutex.Unlock()
	fake.DeleteUnresponsiveEphemeralWorkersStub = nil
	if fake.deleteUnresponsiveEphemeralWorkersReturnsOnCall == nil {
		fake.deleteUnresponsiveEphemeralWorkersReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.deleteUnresponsiveEphemeralWorkersReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerLifecycle) GetWorkerStateByName() (map[string]db.WorkerState, error) {
	fake.getWorkerStateByNameMutex.Lock()
	ret, specificReturn := fake.getWorkerStateByNameReturnsOnCall[len(fake.getWorkerStateByNameArgsForCall)]
	fake.getWorkerStateByNameArgsForCall = append(fake.getWorkerStateByNameArgsForCall, struct {
	}{})
	stub := fake.GetWorkerStateByNameStub
	fakeReturns := fake.getWorkerStateByNameReturns
	fake.recordInvocation("GetWorkerStateByName", []interface{}{})
	fake.getWorkerStateByNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkerLifecycle) GetWorkerStateByNameCallCount() int {
	fake.getWorkerStateByNameMutex.RLock()
	defer fake.getWorkerStateByNameMutex.RUnlock()
	return len(fake.getWorkerStateByNameArgsForCall)
}

func (fake *FakeWorkerLifecycle) GetWorkerStateByNameCalls(stub func() (map[string]db.WorkerState, error)) {
	fake.getWorkerStateByNameMutex.Lock()
	defer fake.getWorkerStateByNameMutex.Unlock()
	fake.GetWorkerStateByNameStub = stub
}

func (fake *FakeWorkerLifecycle) GetWorkerStateByNameReturns(result1 map[string]db.WorkerState, result2 error) {
	fake.getWorkerStateByNameMutex.Lock()
	defer fake.getWorkerStateByNameMutex.Unlock()
	fake.GetWorkerStateByNameStub = nil
	fake.getWorkerStateByNameReturns = struct {
		result1 map[string]db.WorkerState
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerLifecycle) GetWorkerStateByNameReturnsOnCall(i int, result1 map[string]db.WorkerState, result2 error) {
	fake.getWorkerStateByNameMutex.Lock()
	defer fake.getWorkerStateByNameMutex.Unlock()
	fake.GetWorkerStateByNameStub = nil
	if fake.getWorkerStateByNameReturnsOnCall == nil {
		fake.getWorkerStateByNameReturnsOnCall = make(map[int]struct {
			result1 map[string]db.WorkerState
			result2 error
		})
	}
	fake.getWorkerStateByNameReturnsOnCall[i] = struct {
		result1 map[string]db.WorkerState
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerLifecycle) LandFinishedLandingWorkers() ([]string, error) {
	fake.landFinishedLandingWorkersMutex.Lock()
	ret, specificReturn := fake.landFinishedLandingWorkersReturnsOnCall[len(fake.landFinishedLandingWorkersArgsForCall)]
	fake.landFinishedLandingWorkersArgsForCall = append(fake.landFinishedLandingWorkersArgsForCall, struct {
	}{})
	stub := fake.LandFinishedLandingWorkersStub
	fakeReturns := fake.landFinishedLandingWorkersReturns
	fake.recordInvocation("LandFinishedLandingWorkers", []interface{}{})
	fake.landFinishedLandingWorkersMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkerLifecycle) LandFinishedLandingWorkersCallCount() int {
	fake.landFinishedLandingWorkersMutex.RLock()
	defer fake.landFinishedLandingWorkersMutex.RUnlock()
	return len(fake.landFinishedLandingWorkersArgsForCall)
}

func (fake *FakeWorkerLifecycle) LandFinishedLandingWorkersCalls(stub func() ([]string, error)) {
	fake.landFinishedLandingWorkersMutex.Lock()
	defer fake.landFinishedLandingWorkersMutex.Unlock()
	fake.LandFinishedLandingWorkersStub = stub
}

func (fake *FakeWorkerLifecycle) LandFinishedLandingWorkersReturns(result1 []string, result2 error) {
	fake.landFinishedLandingWorkersMutex.Lock()
	defer fake.landFinishedLandingWorkersMutex.Unlock()
	fake.LandFinishedLandingWorkersStub = nil
	fake.landFinishedLandingWorkersReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerLifecycle) LandFinishedLandingWorkersReturnsOnCall(i int, result1 []string, result2 error) {
	fake.landFinishedLandingWorkersMutex.Lock()
	defer fake.landFinishedLandingWorkersMutex.Unlock()
	fake.LandFinishedLandingWorkersStub = nil
	if fake.landFinishedLandingWorkersReturnsOnCall == nil {
		fake.landFinishedLandingWorkersReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.landFinishedLandingWorkersReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerLifecycle) StallUnresponsiveWorkers() ([]string, error) {
	fake.stallUnresponsiveWorkersMutex.Lock()
	ret, specificReturn := fake.stallUnresponsiveWorkersReturnsOnCall[len(fake.stallUnresponsiveWorkersArgsForCall)]
	fake.stallUnresponsiveWorkersArgsForCall = append(fake.stallUnresponsiveWorkersArgsForCall, struct {
	}{})
	stub := fake.StallUnresponsiveWorkersStub
	fakeReturns := fake.stallUnresponsiveWorkersReturns
	fake.recordInvocation("StallUnresponsiveWorkers", []interface{}{})
	fake.stallUnresponsiveWorkersMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkerLifecycle) StallUnresponsiveWorkersCallCount() int {
	fake.stallUnresponsiveWorkersMutex.RLock()
	defer fake.stallUnresponsiveWorkersMutex.RUnlock()
	return len(fake.stallUnresponsiveWorkersArgsForCall)
}

func (fake *FakeWorkerLifecycle) StallUnresponsiveWorkersCalls(stub func() ([]string, error)) {
	fake.stallUnresponsiveWorkersMutex.Lock()
	defer fake.stallUnresponsiveWorkersMutex.Unlock()
	fake.StallUnresponsiveWorkersStub = stub
}

func (fake *FakeWorkerLifecycle) StallUnresponsiveWorkersReturns(result1 []string, result2 error) {
	fake.stallUnresponsiveWorkersMutex.Lock()
	defer fake.stallUnresponsiveWorkersMutex.Unlock()
	fake.StallUnresponsiveWorkersStub = nil
	fake.stallUnresponsiveWorkersReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerLifecycle) StallUnresponsiveWorkersReturnsOnCall(i int, result1 []string, result2 error) {
	fake.stallUnresponsiveWorkersMutex.Lock()
	defer fake.stallUnresponsiveWorkersMutex.Unlock()
	fake.StallUnresponsiveWorkersStub = nil
	if fake.stallUnresponsiveWorkersReturnsOnCall == nil {
		fake.stallUnresponsiveWorkersReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.stallUnresponsiveWorkersReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerLifecycle) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteFinishedRetiringWorkersMutex.RLock()
	defer fake.deleteFinishedRetiringWorkersMutex.RUnlock()
	fake.deleteUnresponsiveEphemeralWorkersMutex.RLock()
	defer fake.deleteUnresponsiveEphemeralWorkersMutex.RUnlock()
	fake.getWorkerStateByNameMutex.RLock()
	defer fake.getWorkerStateByNameMutex.RUnlock()
	fake.landFinishedLandingWorkersMutex.RLock()
	defer fake.landFinishedLandingWorkersMutex.RUnlock()
	fake.stallUnresponsiveWorkersMutex.RLock()
	defer fake.stallUnresponsiveWorkersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWorkerLifecycle) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.WorkerLifecycle = new(FakeWorkerLifecycle)