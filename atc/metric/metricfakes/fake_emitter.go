// Code generated by counterfeiter. DO NOT EDIT.
package metricfakes

import (
	"sync"

	"code.cloudfoundry.org/lager/v3"
	"github.com/concourse/concourse/atc/metric"
)

type FakeEmitter struct {
	EmitStub        func(lager.Logger, metric.Event)
	emitMutex       sync.RWMutex
	emitArgsForCall []struct {
		arg1 lager.Logger
		arg2 metric.Event
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEmitter) Emit(arg1 lager.Logger, arg2 metric.Event) {
	fake.emitMutex.Lock()
	fake.emitArgsForCall = append(fake.emitArgsForCall, struct {
		arg1 lager.Logger
		arg2 metric.Event
	}{arg1, arg2})
	stub := fake.EmitStub
	fake.recordInvocation("Emit", []interface{}{arg1, arg2})
	fake.emitMutex.Unlock()
	if stub != nil {
		fake.EmitStub(arg1, arg2)
	}
}

func (fake *FakeEmitter) EmitCallCount() int {
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	return len(fake.emitArgsForCall)
}

func (fake *FakeEmitter) EmitCalls(stub func(lager.Logger, metric.Event)) {
	fake.emitMutex.Lock()
	defer fake.emitMutex.Unlock()
	fake.EmitStub = stub
}

func (fake *FakeEmitter) EmitArgsForCall(i int) (lager.Logger, metric.Event) {
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	argsForCall := fake.emitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEmitter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEmitter) recordInvocation(key string, args []interface{}) {
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

var _ metric.Emitter = new(FakeEmitter)
