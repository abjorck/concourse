// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"io"
	"sync"

	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/worker"
)

type FakeArtifact struct {
	DBArtifactStub        func() db.WorkerArtifact
	dBArtifactMutex       sync.RWMutex
	dBArtifactArgsForCall []struct {
	}
	dBArtifactReturns struct {
		result1 db.WorkerArtifact
	}
	dBArtifactReturnsOnCall map[int]struct {
		result1 db.WorkerArtifact
	}
	InitializedStub        func() bool
	initializedMutex       sync.RWMutex
	initializedArgsForCall []struct {
	}
	initializedReturns struct {
		result1 bool
	}
	initializedReturnsOnCall map[int]struct {
		result1 bool
	}
	RetrieveStub        func(string) (io.ReadCloser, error)
	retrieveMutex       sync.RWMutex
	retrieveArgsForCall []struct {
		arg1 string
	}
	retrieveReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	retrieveReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	StoreStub        func(string, io.Reader) error
	storeMutex       sync.RWMutex
	storeArgsForCall []struct {
		arg1 string
		arg2 io.Reader
	}
	storeReturns struct {
		result1 error
	}
	storeReturnsOnCall map[int]struct {
		result1 error
	}
	VolumeStub        func() worker.Volume
	volumeMutex       sync.RWMutex
	volumeArgsForCall []struct {
	}
	volumeReturns struct {
		result1 worker.Volume
	}
	volumeReturnsOnCall map[int]struct {
		result1 worker.Volume
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeArtifact) DBArtifact() db.WorkerArtifact {
	fake.dBArtifactMutex.Lock()
	ret, specificReturn := fake.dBArtifactReturnsOnCall[len(fake.dBArtifactArgsForCall)]
	fake.dBArtifactArgsForCall = append(fake.dBArtifactArgsForCall, struct {
	}{})
	fake.recordInvocation("DBArtifact", []interface{}{})
	fake.dBArtifactMutex.Unlock()
	if fake.DBArtifactStub != nil {
		return fake.DBArtifactStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.dBArtifactReturns
	return fakeReturns.result1
}

func (fake *FakeArtifact) DBArtifactCallCount() int {
	fake.dBArtifactMutex.RLock()
	defer fake.dBArtifactMutex.RUnlock()
	return len(fake.dBArtifactArgsForCall)
}

func (fake *FakeArtifact) DBArtifactCalls(stub func() db.WorkerArtifact) {
	fake.dBArtifactMutex.Lock()
	defer fake.dBArtifactMutex.Unlock()
	fake.DBArtifactStub = stub
}

func (fake *FakeArtifact) DBArtifactReturns(result1 db.WorkerArtifact) {
	fake.dBArtifactMutex.Lock()
	defer fake.dBArtifactMutex.Unlock()
	fake.DBArtifactStub = nil
	fake.dBArtifactReturns = struct {
		result1 db.WorkerArtifact
	}{result1}
}

func (fake *FakeArtifact) DBArtifactReturnsOnCall(i int, result1 db.WorkerArtifact) {
	fake.dBArtifactMutex.Lock()
	defer fake.dBArtifactMutex.Unlock()
	fake.DBArtifactStub = nil
	if fake.dBArtifactReturnsOnCall == nil {
		fake.dBArtifactReturnsOnCall = make(map[int]struct {
			result1 db.WorkerArtifact
		})
	}
	fake.dBArtifactReturnsOnCall[i] = struct {
		result1 db.WorkerArtifact
	}{result1}
}

func (fake *FakeArtifact) Initialized() bool {
	fake.initializedMutex.Lock()
	ret, specificReturn := fake.initializedReturnsOnCall[len(fake.initializedArgsForCall)]
	fake.initializedArgsForCall = append(fake.initializedArgsForCall, struct {
	}{})
	fake.recordInvocation("Initialized", []interface{}{})
	fake.initializedMutex.Unlock()
	if fake.InitializedStub != nil {
		return fake.InitializedStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.initializedReturns
	return fakeReturns.result1
}

func (fake *FakeArtifact) InitializedCallCount() int {
	fake.initializedMutex.RLock()
	defer fake.initializedMutex.RUnlock()
	return len(fake.initializedArgsForCall)
}

func (fake *FakeArtifact) InitializedCalls(stub func() bool) {
	fake.initializedMutex.Lock()
	defer fake.initializedMutex.Unlock()
	fake.InitializedStub = stub
}

func (fake *FakeArtifact) InitializedReturns(result1 bool) {
	fake.initializedMutex.Lock()
	defer fake.initializedMutex.Unlock()
	fake.InitializedStub = nil
	fake.initializedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeArtifact) InitializedReturnsOnCall(i int, result1 bool) {
	fake.initializedMutex.Lock()
	defer fake.initializedMutex.Unlock()
	fake.InitializedStub = nil
	if fake.initializedReturnsOnCall == nil {
		fake.initializedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.initializedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeArtifact) Retrieve(arg1 string) (io.ReadCloser, error) {
	fake.retrieveMutex.Lock()
	ret, specificReturn := fake.retrieveReturnsOnCall[len(fake.retrieveArgsForCall)]
	fake.retrieveArgsForCall = append(fake.retrieveArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Retrieve", []interface{}{arg1})
	fake.retrieveMutex.Unlock()
	if fake.RetrieveStub != nil {
		return fake.RetrieveStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.retrieveReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeArtifact) RetrieveCallCount() int {
	fake.retrieveMutex.RLock()
	defer fake.retrieveMutex.RUnlock()
	return len(fake.retrieveArgsForCall)
}

func (fake *FakeArtifact) RetrieveCalls(stub func(string) (io.ReadCloser, error)) {
	fake.retrieveMutex.Lock()
	defer fake.retrieveMutex.Unlock()
	fake.RetrieveStub = stub
}

func (fake *FakeArtifact) RetrieveArgsForCall(i int) string {
	fake.retrieveMutex.RLock()
	defer fake.retrieveMutex.RUnlock()
	argsForCall := fake.retrieveArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeArtifact) RetrieveReturns(result1 io.ReadCloser, result2 error) {
	fake.retrieveMutex.Lock()
	defer fake.retrieveMutex.Unlock()
	fake.RetrieveStub = nil
	fake.retrieveReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeArtifact) RetrieveReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.retrieveMutex.Lock()
	defer fake.retrieveMutex.Unlock()
	fake.RetrieveStub = nil
	if fake.retrieveReturnsOnCall == nil {
		fake.retrieveReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.retrieveReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeArtifact) Store(arg1 string, arg2 io.Reader) error {
	fake.storeMutex.Lock()
	ret, specificReturn := fake.storeReturnsOnCall[len(fake.storeArgsForCall)]
	fake.storeArgsForCall = append(fake.storeArgsForCall, struct {
		arg1 string
		arg2 io.Reader
	}{arg1, arg2})
	fake.recordInvocation("Store", []interface{}{arg1, arg2})
	fake.storeMutex.Unlock()
	if fake.StoreStub != nil {
		return fake.StoreStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.storeReturns
	return fakeReturns.result1
}

func (fake *FakeArtifact) StoreCallCount() int {
	fake.storeMutex.RLock()
	defer fake.storeMutex.RUnlock()
	return len(fake.storeArgsForCall)
}

func (fake *FakeArtifact) StoreCalls(stub func(string, io.Reader) error) {
	fake.storeMutex.Lock()
	defer fake.storeMutex.Unlock()
	fake.StoreStub = stub
}

func (fake *FakeArtifact) StoreArgsForCall(i int) (string, io.Reader) {
	fake.storeMutex.RLock()
	defer fake.storeMutex.RUnlock()
	argsForCall := fake.storeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeArtifact) StoreReturns(result1 error) {
	fake.storeMutex.Lock()
	defer fake.storeMutex.Unlock()
	fake.StoreStub = nil
	fake.storeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeArtifact) StoreReturnsOnCall(i int, result1 error) {
	fake.storeMutex.Lock()
	defer fake.storeMutex.Unlock()
	fake.StoreStub = nil
	if fake.storeReturnsOnCall == nil {
		fake.storeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeArtifact) Volume() worker.Volume {
	fake.volumeMutex.Lock()
	ret, specificReturn := fake.volumeReturnsOnCall[len(fake.volumeArgsForCall)]
	fake.volumeArgsForCall = append(fake.volumeArgsForCall, struct {
	}{})
	fake.recordInvocation("Volume", []interface{}{})
	fake.volumeMutex.Unlock()
	if fake.VolumeStub != nil {
		return fake.VolumeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.volumeReturns
	return fakeReturns.result1
}

func (fake *FakeArtifact) VolumeCallCount() int {
	fake.volumeMutex.RLock()
	defer fake.volumeMutex.RUnlock()
	return len(fake.volumeArgsForCall)
}

func (fake *FakeArtifact) VolumeCalls(stub func() worker.Volume) {
	fake.volumeMutex.Lock()
	defer fake.volumeMutex.Unlock()
	fake.VolumeStub = stub
}

func (fake *FakeArtifact) VolumeReturns(result1 worker.Volume) {
	fake.volumeMutex.Lock()
	defer fake.volumeMutex.Unlock()
	fake.VolumeStub = nil
	fake.volumeReturns = struct {
		result1 worker.Volume
	}{result1}
}

func (fake *FakeArtifact) VolumeReturnsOnCall(i int, result1 worker.Volume) {
	fake.volumeMutex.Lock()
	defer fake.volumeMutex.Unlock()
	fake.VolumeStub = nil
	if fake.volumeReturnsOnCall == nil {
		fake.volumeReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
		})
	}
	fake.volumeReturnsOnCall[i] = struct {
		result1 worker.Volume
	}{result1}
}

func (fake *FakeArtifact) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dBArtifactMutex.RLock()
	defer fake.dBArtifactMutex.RUnlock()
	fake.initializedMutex.RLock()
	defer fake.initializedMutex.RUnlock()
	fake.retrieveMutex.RLock()
	defer fake.retrieveMutex.RUnlock()
	fake.storeMutex.RLock()
	defer fake.storeMutex.RUnlock()
	fake.volumeMutex.RLock()
	defer fake.volumeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeArtifact) recordInvocation(key string, args []interface{}) {
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

var _ worker.Artifact = new(FakeArtifact)