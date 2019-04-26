// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/worker"
)

type FakeClient struct {
	CreateArtifactStub        func(lager.Logger, string) (atc.WorkerArtifact, error)
	createArtifactMutex       sync.RWMutex
	createArtifactArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	createArtifactReturns struct {
		result1 atc.WorkerArtifact
		result2 error
	}
	createArtifactReturnsOnCall map[int]struct {
		result1 atc.WorkerArtifact
		result2 error
	}
	FindContainerStub        func(lager.Logger, int, string) (worker.Container, bool, error)
	findContainerMutex       sync.RWMutex
	findContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
	}
	findContainerReturns struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	findContainerReturnsOnCall map[int]struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	FindVolumeStub        func(lager.Logger, int, string) (worker.Volume, bool, error)
	findVolumeMutex       sync.RWMutex
	findVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
	}
	findVolumeReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	findVolumeReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	StoreStub        func(lager.Logger, int, atc.WorkerArtifact, string, io.ReadCloser) error
	storeMutex       sync.RWMutex
	storeArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
		arg3 atc.WorkerArtifact
		arg4 string
		arg5 io.ReadCloser
	}
	storeReturns struct {
		result1 error
	}
	storeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) CreateArtifact(arg1 lager.Logger, arg2 string) (atc.WorkerArtifact, error) {
	fake.createArtifactMutex.Lock()
	ret, specificReturn := fake.createArtifactReturnsOnCall[len(fake.createArtifactArgsForCall)]
	fake.createArtifactArgsForCall = append(fake.createArtifactArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateArtifact", []interface{}{arg1, arg2})
	fake.createArtifactMutex.Unlock()
	if fake.CreateArtifactStub != nil {
		return fake.CreateArtifactStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createArtifactReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) CreateArtifactCallCount() int {
	fake.createArtifactMutex.RLock()
	defer fake.createArtifactMutex.RUnlock()
	return len(fake.createArtifactArgsForCall)
}

func (fake *FakeClient) CreateArtifactCalls(stub func(lager.Logger, string) (atc.WorkerArtifact, error)) {
	fake.createArtifactMutex.Lock()
	defer fake.createArtifactMutex.Unlock()
	fake.CreateArtifactStub = stub
}

func (fake *FakeClient) CreateArtifactArgsForCall(i int) (lager.Logger, string) {
	fake.createArtifactMutex.RLock()
	defer fake.createArtifactMutex.RUnlock()
	argsForCall := fake.createArtifactArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) CreateArtifactReturns(result1 atc.WorkerArtifact, result2 error) {
	fake.createArtifactMutex.Lock()
	defer fake.createArtifactMutex.Unlock()
	fake.CreateArtifactStub = nil
	fake.createArtifactReturns = struct {
		result1 atc.WorkerArtifact
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CreateArtifactReturnsOnCall(i int, result1 atc.WorkerArtifact, result2 error) {
	fake.createArtifactMutex.Lock()
	defer fake.createArtifactMutex.Unlock()
	fake.CreateArtifactStub = nil
	if fake.createArtifactReturnsOnCall == nil {
		fake.createArtifactReturnsOnCall = make(map[int]struct {
			result1 atc.WorkerArtifact
			result2 error
		})
	}
	fake.createArtifactReturnsOnCall[i] = struct {
		result1 atc.WorkerArtifact
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindContainer(arg1 lager.Logger, arg2 int, arg3 string) (worker.Container, bool, error) {
	fake.findContainerMutex.Lock()
	ret, specificReturn := fake.findContainerReturnsOnCall[len(fake.findContainerArgsForCall)]
	fake.findContainerArgsForCall = append(fake.findContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("FindContainer", []interface{}{arg1, arg2, arg3})
	fake.findContainerMutex.Unlock()
	if fake.FindContainerStub != nil {
		return fake.FindContainerStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.findContainerReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) FindContainerCallCount() int {
	fake.findContainerMutex.RLock()
	defer fake.findContainerMutex.RUnlock()
	return len(fake.findContainerArgsForCall)
}

func (fake *FakeClient) FindContainerCalls(stub func(lager.Logger, int, string) (worker.Container, bool, error)) {
	fake.findContainerMutex.Lock()
	defer fake.findContainerMutex.Unlock()
	fake.FindContainerStub = stub
}

func (fake *FakeClient) FindContainerArgsForCall(i int) (lager.Logger, int, string) {
	fake.findContainerMutex.RLock()
	defer fake.findContainerMutex.RUnlock()
	argsForCall := fake.findContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) FindContainerReturns(result1 worker.Container, result2 bool, result3 error) {
	fake.findContainerMutex.Lock()
	defer fake.findContainerMutex.Unlock()
	fake.FindContainerStub = nil
	fake.findContainerReturns = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) FindContainerReturnsOnCall(i int, result1 worker.Container, result2 bool, result3 error) {
	fake.findContainerMutex.Lock()
	defer fake.findContainerMutex.Unlock()
	fake.FindContainerStub = nil
	if fake.findContainerReturnsOnCall == nil {
		fake.findContainerReturnsOnCall = make(map[int]struct {
			result1 worker.Container
			result2 bool
			result3 error
		})
	}
	fake.findContainerReturnsOnCall[i] = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) FindVolume(arg1 lager.Logger, arg2 int, arg3 string) (worker.Volume, bool, error) {
	fake.findVolumeMutex.Lock()
	ret, specificReturn := fake.findVolumeReturnsOnCall[len(fake.findVolumeArgsForCall)]
	fake.findVolumeArgsForCall = append(fake.findVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("FindVolume", []interface{}{arg1, arg2, arg3})
	fake.findVolumeMutex.Unlock()
	if fake.FindVolumeStub != nil {
		return fake.FindVolumeStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.findVolumeReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) FindVolumeCallCount() int {
	fake.findVolumeMutex.RLock()
	defer fake.findVolumeMutex.RUnlock()
	return len(fake.findVolumeArgsForCall)
}

func (fake *FakeClient) FindVolumeCalls(stub func(lager.Logger, int, string) (worker.Volume, bool, error)) {
	fake.findVolumeMutex.Lock()
	defer fake.findVolumeMutex.Unlock()
	fake.FindVolumeStub = stub
}

func (fake *FakeClient) FindVolumeArgsForCall(i int) (lager.Logger, int, string) {
	fake.findVolumeMutex.RLock()
	defer fake.findVolumeMutex.RUnlock()
	argsForCall := fake.findVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) FindVolumeReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.findVolumeMutex.Lock()
	defer fake.findVolumeMutex.Unlock()
	fake.FindVolumeStub = nil
	fake.findVolumeReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) FindVolumeReturnsOnCall(i int, result1 worker.Volume, result2 bool, result3 error) {
	fake.findVolumeMutex.Lock()
	defer fake.findVolumeMutex.Unlock()
	fake.FindVolumeStub = nil
	if fake.findVolumeReturnsOnCall == nil {
		fake.findVolumeReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 bool
			result3 error
		})
	}
	fake.findVolumeReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) Store(arg1 lager.Logger, arg2 int, arg3 atc.WorkerArtifact, arg4 string, arg5 io.ReadCloser) error {
	fake.storeMutex.Lock()
	ret, specificReturn := fake.storeReturnsOnCall[len(fake.storeArgsForCall)]
	fake.storeArgsForCall = append(fake.storeArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
		arg3 atc.WorkerArtifact
		arg4 string
		arg5 io.ReadCloser
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("Store", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.storeMutex.Unlock()
	if fake.StoreStub != nil {
		return fake.StoreStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.storeReturns
	return fakeReturns.result1
}

func (fake *FakeClient) StoreCallCount() int {
	fake.storeMutex.RLock()
	defer fake.storeMutex.RUnlock()
	return len(fake.storeArgsForCall)
}

func (fake *FakeClient) StoreCalls(stub func(lager.Logger, int, atc.WorkerArtifact, string, io.ReadCloser) error) {
	fake.storeMutex.Lock()
	defer fake.storeMutex.Unlock()
	fake.StoreStub = stub
}

func (fake *FakeClient) StoreArgsForCall(i int) (lager.Logger, int, atc.WorkerArtifact, string, io.ReadCloser) {
	fake.storeMutex.RLock()
	defer fake.storeMutex.RUnlock()
	argsForCall := fake.storeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeClient) StoreReturns(result1 error) {
	fake.storeMutex.Lock()
	defer fake.storeMutex.Unlock()
	fake.StoreStub = nil
	fake.storeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) StoreReturnsOnCall(i int, result1 error) {
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

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createArtifactMutex.RLock()
	defer fake.createArtifactMutex.RUnlock()
	fake.findContainerMutex.RLock()
	defer fake.findContainerMutex.RUnlock()
	fake.findVolumeMutex.RLock()
	defer fake.findVolumeMutex.RUnlock()
	fake.storeMutex.RLock()
	defer fake.storeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ worker.Client = new(FakeClient)
