// Code generated by counterfeiter. DO NOT EDIT.
package routingfakes

import (
	"context"
	"sync"

	"github.com/livekit/livekit-server/pkg/routing"
	"github.com/livekit/protocol/livekit"
)

type FakeSignalClient struct {
	ActiveCountStub        func() int
	activeCountMutex       sync.RWMutex
	activeCountArgsForCall []struct {
	}
	activeCountReturns struct {
		result1 int
	}
	activeCountReturnsOnCall map[int]struct {
		result1 int
	}
	StartParticipantSignalStub        func(context.Context, livekit.RoomName, routing.ParticipantInit, livekit.NodeID) (livekit.ConnectionID, routing.MessageSink, routing.MessageSource, error)
	startParticipantSignalMutex       sync.RWMutex
	startParticipantSignalArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 routing.ParticipantInit
		arg4 livekit.NodeID
	}
	startParticipantSignalReturns struct {
		result1 livekit.ConnectionID
		result2 routing.MessageSink
		result3 routing.MessageSource
		result4 error
	}
	startParticipantSignalReturnsOnCall map[int]struct {
		result1 livekit.ConnectionID
		result2 routing.MessageSink
		result3 routing.MessageSource
		result4 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSignalClient) ActiveCount() int {
	fake.activeCountMutex.Lock()
	ret, specificReturn := fake.activeCountReturnsOnCall[len(fake.activeCountArgsForCall)]
	fake.activeCountArgsForCall = append(fake.activeCountArgsForCall, struct {
	}{})
	stub := fake.ActiveCountStub
	fakeReturns := fake.activeCountReturns
	fake.recordInvocation("ActiveCount", []interface{}{})
	fake.activeCountMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSignalClient) ActiveCountCallCount() int {
	fake.activeCountMutex.RLock()
	defer fake.activeCountMutex.RUnlock()
	return len(fake.activeCountArgsForCall)
}

func (fake *FakeSignalClient) ActiveCountCalls(stub func() int) {
	fake.activeCountMutex.Lock()
	defer fake.activeCountMutex.Unlock()
	fake.ActiveCountStub = stub
}

func (fake *FakeSignalClient) ActiveCountReturns(result1 int) {
	fake.activeCountMutex.Lock()
	defer fake.activeCountMutex.Unlock()
	fake.ActiveCountStub = nil
	fake.activeCountReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeSignalClient) ActiveCountReturnsOnCall(i int, result1 int) {
	fake.activeCountMutex.Lock()
	defer fake.activeCountMutex.Unlock()
	fake.ActiveCountStub = nil
	if fake.activeCountReturnsOnCall == nil {
		fake.activeCountReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.activeCountReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeSignalClient) StartParticipantSignal(arg1 context.Context, arg2 livekit.RoomName, arg3 routing.ParticipantInit, arg4 livekit.NodeID) (livekit.ConnectionID, routing.MessageSink, routing.MessageSource, error) {
	fake.startParticipantSignalMutex.Lock()
	ret, specificReturn := fake.startParticipantSignalReturnsOnCall[len(fake.startParticipantSignalArgsForCall)]
	fake.startParticipantSignalArgsForCall = append(fake.startParticipantSignalArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 routing.ParticipantInit
		arg4 livekit.NodeID
	}{arg1, arg2, arg3, arg4})
	stub := fake.StartParticipantSignalStub
	fakeReturns := fake.startParticipantSignalReturns
	fake.recordInvocation("StartParticipantSignal", []interface{}{arg1, arg2, arg3, arg4})
	fake.startParticipantSignalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakeSignalClient) StartParticipantSignalCallCount() int {
	fake.startParticipantSignalMutex.RLock()
	defer fake.startParticipantSignalMutex.RUnlock()
	return len(fake.startParticipantSignalArgsForCall)
}

func (fake *FakeSignalClient) StartParticipantSignalCalls(stub func(context.Context, livekit.RoomName, routing.ParticipantInit, livekit.NodeID) (livekit.ConnectionID, routing.MessageSink, routing.MessageSource, error)) {
	fake.startParticipantSignalMutex.Lock()
	defer fake.startParticipantSignalMutex.Unlock()
	fake.StartParticipantSignalStub = stub
}

func (fake *FakeSignalClient) StartParticipantSignalArgsForCall(i int) (context.Context, livekit.RoomName, routing.ParticipantInit, livekit.NodeID) {
	fake.startParticipantSignalMutex.RLock()
	defer fake.startParticipantSignalMutex.RUnlock()
	argsForCall := fake.startParticipantSignalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSignalClient) StartParticipantSignalReturns(result1 livekit.ConnectionID, result2 routing.MessageSink, result3 routing.MessageSource, result4 error) {
	fake.startParticipantSignalMutex.Lock()
	defer fake.startParticipantSignalMutex.Unlock()
	fake.StartParticipantSignalStub = nil
	fake.startParticipantSignalReturns = struct {
		result1 livekit.ConnectionID
		result2 routing.MessageSink
		result3 routing.MessageSource
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeSignalClient) StartParticipantSignalReturnsOnCall(i int, result1 livekit.ConnectionID, result2 routing.MessageSink, result3 routing.MessageSource, result4 error) {
	fake.startParticipantSignalMutex.Lock()
	defer fake.startParticipantSignalMutex.Unlock()
	fake.StartParticipantSignalStub = nil
	if fake.startParticipantSignalReturnsOnCall == nil {
		fake.startParticipantSignalReturnsOnCall = make(map[int]struct {
			result1 livekit.ConnectionID
			result2 routing.MessageSink
			result3 routing.MessageSource
			result4 error
		})
	}
	fake.startParticipantSignalReturnsOnCall[i] = struct {
		result1 livekit.ConnectionID
		result2 routing.MessageSink
		result3 routing.MessageSource
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeSignalClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.activeCountMutex.RLock()
	defer fake.activeCountMutex.RUnlock()
	fake.startParticipantSignalMutex.RLock()
	defer fake.startParticipantSignalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSignalClient) recordInvocation(key string, args []interface{}) {
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

var _ routing.SignalClient = new(FakeSignalClient)