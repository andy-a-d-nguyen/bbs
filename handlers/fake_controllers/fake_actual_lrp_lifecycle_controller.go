// Code generated by counterfeiter. DO NOT EDIT.
package fake_controllers

import (
	"context"
	"sync"

	"code.cloudfoundry.org/bbs/handlers"
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/lager"
)

type FakeActualLRPLifecycleController struct {
	ClaimActualLRPStub        func(context.Context, lager.Logger, string, int32, *models.ActualLRPInstanceKey) error
	claimActualLRPMutex       sync.RWMutex
	claimActualLRPArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 int32
		arg5 *models.ActualLRPInstanceKey
	}
	claimActualLRPReturns struct {
		result1 error
	}
	claimActualLRPReturnsOnCall map[int]struct {
		result1 error
	}
	CrashActualLRPStub        func(context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, string) error
	crashActualLRPMutex       sync.RWMutex
	crashActualLRPArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 *models.ActualLRPInstanceKey
		arg5 string
	}
	crashActualLRPReturns struct {
		result1 error
	}
	crashActualLRPReturnsOnCall map[int]struct {
		result1 error
	}
	FailActualLRPStub        func(context.Context, lager.Logger, *models.ActualLRPKey, string) error
	failActualLRPMutex       sync.RWMutex
	failActualLRPArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 string
	}
	failActualLRPReturns struct {
		result1 error
	}
	failActualLRPReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveActualLRPStub        func(context.Context, lager.Logger, string, int32, *models.ActualLRPInstanceKey) error
	removeActualLRPMutex       sync.RWMutex
	removeActualLRPArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 int32
		arg5 *models.ActualLRPInstanceKey
	}
	removeActualLRPReturns struct {
		result1 error
	}
	removeActualLRPReturnsOnCall map[int]struct {
		result1 error
	}
	RetireActualLRPStub        func(context.Context, lager.Logger, *models.ActualLRPKey) error
	retireActualLRPMutex       sync.RWMutex
	retireActualLRPArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
	}
	retireActualLRPReturns struct {
		result1 error
	}
	retireActualLRPReturnsOnCall map[int]struct {
		result1 error
	}
	StartActualLRPStub        func(context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, *models.ActualLRPNetInfo) error
	startActualLRPMutex       sync.RWMutex
	startActualLRPArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 *models.ActualLRPInstanceKey
		arg5 *models.ActualLRPNetInfo
	}
	startActualLRPReturns struct {
		result1 error
	}
	startActualLRPReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeActualLRPLifecycleController) ClaimActualLRP(arg1 context.Context, arg2 lager.Logger, arg3 string, arg4 int32, arg5 *models.ActualLRPInstanceKey) error {
	fake.claimActualLRPMutex.Lock()
	ret, specificReturn := fake.claimActualLRPReturnsOnCall[len(fake.claimActualLRPArgsForCall)]
	fake.claimActualLRPArgsForCall = append(fake.claimActualLRPArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 int32
		arg5 *models.ActualLRPInstanceKey
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.ClaimActualLRPStub
	fakeReturns := fake.claimActualLRPReturns
	fake.recordInvocation("ClaimActualLRP", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.claimActualLRPMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeActualLRPLifecycleController) ClaimActualLRPCallCount() int {
	fake.claimActualLRPMutex.RLock()
	defer fake.claimActualLRPMutex.RUnlock()
	return len(fake.claimActualLRPArgsForCall)
}

func (fake *FakeActualLRPLifecycleController) ClaimActualLRPCalls(stub func(context.Context, lager.Logger, string, int32, *models.ActualLRPInstanceKey) error) {
	fake.claimActualLRPMutex.Lock()
	defer fake.claimActualLRPMutex.Unlock()
	fake.ClaimActualLRPStub = stub
}

func (fake *FakeActualLRPLifecycleController) ClaimActualLRPArgsForCall(i int) (context.Context, lager.Logger, string, int32, *models.ActualLRPInstanceKey) {
	fake.claimActualLRPMutex.RLock()
	defer fake.claimActualLRPMutex.RUnlock()
	argsForCall := fake.claimActualLRPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeActualLRPLifecycleController) ClaimActualLRPReturns(result1 error) {
	fake.claimActualLRPMutex.Lock()
	defer fake.claimActualLRPMutex.Unlock()
	fake.ClaimActualLRPStub = nil
	fake.claimActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeActualLRPLifecycleController) ClaimActualLRPReturnsOnCall(i int, result1 error) {
	fake.claimActualLRPMutex.Lock()
	defer fake.claimActualLRPMutex.Unlock()
	fake.ClaimActualLRPStub = nil
	if fake.claimActualLRPReturnsOnCall == nil {
		fake.claimActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.claimActualLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeActualLRPLifecycleController) CrashActualLRP(arg1 context.Context, arg2 lager.Logger, arg3 *models.ActualLRPKey, arg4 *models.ActualLRPInstanceKey, arg5 string) error {
	fake.crashActualLRPMutex.Lock()
	ret, specificReturn := fake.crashActualLRPReturnsOnCall[len(fake.crashActualLRPArgsForCall)]
	fake.crashActualLRPArgsForCall = append(fake.crashActualLRPArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 *models.ActualLRPInstanceKey
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.CrashActualLRPStub
	fakeReturns := fake.crashActualLRPReturns
	fake.recordInvocation("CrashActualLRP", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.crashActualLRPMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeActualLRPLifecycleController) CrashActualLRPCallCount() int {
	fake.crashActualLRPMutex.RLock()
	defer fake.crashActualLRPMutex.RUnlock()
	return len(fake.crashActualLRPArgsForCall)
}

func (fake *FakeActualLRPLifecycleController) CrashActualLRPCalls(stub func(context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, string) error) {
	fake.crashActualLRPMutex.Lock()
	defer fake.crashActualLRPMutex.Unlock()
	fake.CrashActualLRPStub = stub
}

func (fake *FakeActualLRPLifecycleController) CrashActualLRPArgsForCall(i int) (context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, string) {
	fake.crashActualLRPMutex.RLock()
	defer fake.crashActualLRPMutex.RUnlock()
	argsForCall := fake.crashActualLRPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeActualLRPLifecycleController) CrashActualLRPReturns(result1 error) {
	fake.crashActualLRPMutex.Lock()
	defer fake.crashActualLRPMutex.Unlock()
	fake.CrashActualLRPStub = nil
	fake.crashActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeActualLRPLifecycleController) CrashActualLRPReturnsOnCall(i int, result1 error) {
	fake.crashActualLRPMutex.Lock()
	defer fake.crashActualLRPMutex.Unlock()
	fake.CrashActualLRPStub = nil
	if fake.crashActualLRPReturnsOnCall == nil {
		fake.crashActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.crashActualLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeActualLRPLifecycleController) FailActualLRP(arg1 context.Context, arg2 lager.Logger, arg3 *models.ActualLRPKey, arg4 string) error {
	fake.failActualLRPMutex.Lock()
	ret, specificReturn := fake.failActualLRPReturnsOnCall[len(fake.failActualLRPArgsForCall)]
	fake.failActualLRPArgsForCall = append(fake.failActualLRPArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.FailActualLRPStub
	fakeReturns := fake.failActualLRPReturns
	fake.recordInvocation("FailActualLRP", []interface{}{arg1, arg2, arg3, arg4})
	fake.failActualLRPMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeActualLRPLifecycleController) FailActualLRPCallCount() int {
	fake.failActualLRPMutex.RLock()
	defer fake.failActualLRPMutex.RUnlock()
	return len(fake.failActualLRPArgsForCall)
}

func (fake *FakeActualLRPLifecycleController) FailActualLRPCalls(stub func(context.Context, lager.Logger, *models.ActualLRPKey, string) error) {
	fake.failActualLRPMutex.Lock()
	defer fake.failActualLRPMutex.Unlock()
	fake.FailActualLRPStub = stub
}

func (fake *FakeActualLRPLifecycleController) FailActualLRPArgsForCall(i int) (context.Context, lager.Logger, *models.ActualLRPKey, string) {
	fake.failActualLRPMutex.RLock()
	defer fake.failActualLRPMutex.RUnlock()
	argsForCall := fake.failActualLRPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeActualLRPLifecycleController) FailActualLRPReturns(result1 error) {
	fake.failActualLRPMutex.Lock()
	defer fake.failActualLRPMutex.Unlock()
	fake.FailActualLRPStub = nil
	fake.failActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeActualLRPLifecycleController) FailActualLRPReturnsOnCall(i int, result1 error) {
	fake.failActualLRPMutex.Lock()
	defer fake.failActualLRPMutex.Unlock()
	fake.FailActualLRPStub = nil
	if fake.failActualLRPReturnsOnCall == nil {
		fake.failActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.failActualLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeActualLRPLifecycleController) RemoveActualLRP(arg1 context.Context, arg2 lager.Logger, arg3 string, arg4 int32, arg5 *models.ActualLRPInstanceKey) error {
	fake.removeActualLRPMutex.Lock()
	ret, specificReturn := fake.removeActualLRPReturnsOnCall[len(fake.removeActualLRPArgsForCall)]
	fake.removeActualLRPArgsForCall = append(fake.removeActualLRPArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 int32
		arg5 *models.ActualLRPInstanceKey
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.RemoveActualLRPStub
	fakeReturns := fake.removeActualLRPReturns
	fake.recordInvocation("RemoveActualLRP", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.removeActualLRPMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeActualLRPLifecycleController) RemoveActualLRPCallCount() int {
	fake.removeActualLRPMutex.RLock()
	defer fake.removeActualLRPMutex.RUnlock()
	return len(fake.removeActualLRPArgsForCall)
}

func (fake *FakeActualLRPLifecycleController) RemoveActualLRPCalls(stub func(context.Context, lager.Logger, string, int32, *models.ActualLRPInstanceKey) error) {
	fake.removeActualLRPMutex.Lock()
	defer fake.removeActualLRPMutex.Unlock()
	fake.RemoveActualLRPStub = stub
}

func (fake *FakeActualLRPLifecycleController) RemoveActualLRPArgsForCall(i int) (context.Context, lager.Logger, string, int32, *models.ActualLRPInstanceKey) {
	fake.removeActualLRPMutex.RLock()
	defer fake.removeActualLRPMutex.RUnlock()
	argsForCall := fake.removeActualLRPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeActualLRPLifecycleController) RemoveActualLRPReturns(result1 error) {
	fake.removeActualLRPMutex.Lock()
	defer fake.removeActualLRPMutex.Unlock()
	fake.RemoveActualLRPStub = nil
	fake.removeActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeActualLRPLifecycleController) RemoveActualLRPReturnsOnCall(i int, result1 error) {
	fake.removeActualLRPMutex.Lock()
	defer fake.removeActualLRPMutex.Unlock()
	fake.RemoveActualLRPStub = nil
	if fake.removeActualLRPReturnsOnCall == nil {
		fake.removeActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeActualLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeActualLRPLifecycleController) RetireActualLRP(arg1 context.Context, arg2 lager.Logger, arg3 *models.ActualLRPKey) error {
	fake.retireActualLRPMutex.Lock()
	ret, specificReturn := fake.retireActualLRPReturnsOnCall[len(fake.retireActualLRPArgsForCall)]
	fake.retireActualLRPArgsForCall = append(fake.retireActualLRPArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
	}{arg1, arg2, arg3})
	stub := fake.RetireActualLRPStub
	fakeReturns := fake.retireActualLRPReturns
	fake.recordInvocation("RetireActualLRP", []interface{}{arg1, arg2, arg3})
	fake.retireActualLRPMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeActualLRPLifecycleController) RetireActualLRPCallCount() int {
	fake.retireActualLRPMutex.RLock()
	defer fake.retireActualLRPMutex.RUnlock()
	return len(fake.retireActualLRPArgsForCall)
}

func (fake *FakeActualLRPLifecycleController) RetireActualLRPCalls(stub func(context.Context, lager.Logger, *models.ActualLRPKey) error) {
	fake.retireActualLRPMutex.Lock()
	defer fake.retireActualLRPMutex.Unlock()
	fake.RetireActualLRPStub = stub
}

func (fake *FakeActualLRPLifecycleController) RetireActualLRPArgsForCall(i int) (context.Context, lager.Logger, *models.ActualLRPKey) {
	fake.retireActualLRPMutex.RLock()
	defer fake.retireActualLRPMutex.RUnlock()
	argsForCall := fake.retireActualLRPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeActualLRPLifecycleController) RetireActualLRPReturns(result1 error) {
	fake.retireActualLRPMutex.Lock()
	defer fake.retireActualLRPMutex.Unlock()
	fake.RetireActualLRPStub = nil
	fake.retireActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeActualLRPLifecycleController) RetireActualLRPReturnsOnCall(i int, result1 error) {
	fake.retireActualLRPMutex.Lock()
	defer fake.retireActualLRPMutex.Unlock()
	fake.RetireActualLRPStub = nil
	if fake.retireActualLRPReturnsOnCall == nil {
		fake.retireActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.retireActualLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeActualLRPLifecycleController) StartActualLRP(arg1 context.Context, arg2 lager.Logger, arg3 *models.ActualLRPKey, arg4 *models.ActualLRPInstanceKey, arg5 *models.ActualLRPNetInfo) error {
	fake.startActualLRPMutex.Lock()
	ret, specificReturn := fake.startActualLRPReturnsOnCall[len(fake.startActualLRPArgsForCall)]
	fake.startActualLRPArgsForCall = append(fake.startActualLRPArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 *models.ActualLRPInstanceKey
		arg5 *models.ActualLRPNetInfo
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.StartActualLRPStub
	fakeReturns := fake.startActualLRPReturns
	fake.recordInvocation("StartActualLRP", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.startActualLRPMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeActualLRPLifecycleController) StartActualLRPCallCount() int {
	fake.startActualLRPMutex.RLock()
	defer fake.startActualLRPMutex.RUnlock()
	return len(fake.startActualLRPArgsForCall)
}

func (fake *FakeActualLRPLifecycleController) StartActualLRPCalls(stub func(context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, *models.ActualLRPNetInfo) error) {
	fake.startActualLRPMutex.Lock()
	defer fake.startActualLRPMutex.Unlock()
	fake.StartActualLRPStub = stub
}

func (fake *FakeActualLRPLifecycleController) StartActualLRPArgsForCall(i int) (context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, *models.ActualLRPNetInfo) {
	fake.startActualLRPMutex.RLock()
	defer fake.startActualLRPMutex.RUnlock()
	argsForCall := fake.startActualLRPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeActualLRPLifecycleController) StartActualLRPReturns(result1 error) {
	fake.startActualLRPMutex.Lock()
	defer fake.startActualLRPMutex.Unlock()
	fake.StartActualLRPStub = nil
	fake.startActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeActualLRPLifecycleController) StartActualLRPReturnsOnCall(i int, result1 error) {
	fake.startActualLRPMutex.Lock()
	defer fake.startActualLRPMutex.Unlock()
	fake.StartActualLRPStub = nil
	if fake.startActualLRPReturnsOnCall == nil {
		fake.startActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startActualLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeActualLRPLifecycleController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.claimActualLRPMutex.RLock()
	defer fake.claimActualLRPMutex.RUnlock()
	fake.crashActualLRPMutex.RLock()
	defer fake.crashActualLRPMutex.RUnlock()
	fake.failActualLRPMutex.RLock()
	defer fake.failActualLRPMutex.RUnlock()
	fake.removeActualLRPMutex.RLock()
	defer fake.removeActualLRPMutex.RUnlock()
	fake.retireActualLRPMutex.RLock()
	defer fake.retireActualLRPMutex.RUnlock()
	fake.startActualLRPMutex.RLock()
	defer fake.startActualLRPMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeActualLRPLifecycleController) recordInvocation(key string, args []interface{}) {
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

var _ handlers.ActualLRPLifecycleController = new(FakeActualLRPLifecycleController)
