// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/bbs/db"
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/lager"
)

type FakeVersionDB struct {
	SetVersionStub        func(context.Context, lager.Logger, *models.Version) error
	setVersionMutex       sync.RWMutex
	setVersionArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.Version
	}
	setVersionReturns struct {
		result1 error
	}
	setVersionReturnsOnCall map[int]struct {
		result1 error
	}
	VersionStub        func(context.Context, lager.Logger) (*models.Version, error)
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
	}
	versionReturns struct {
		result1 *models.Version
		result2 error
	}
	versionReturnsOnCall map[int]struct {
		result1 *models.Version
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVersionDB) SetVersion(arg1 context.Context, arg2 lager.Logger, arg3 *models.Version) error {
	fake.setVersionMutex.Lock()
	ret, specificReturn := fake.setVersionReturnsOnCall[len(fake.setVersionArgsForCall)]
	fake.setVersionArgsForCall = append(fake.setVersionArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.Version
	}{arg1, arg2, arg3})
	stub := fake.SetVersionStub
	fakeReturns := fake.setVersionReturns
	fake.recordInvocation("SetVersion", []interface{}{arg1, arg2, arg3})
	fake.setVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVersionDB) SetVersionCallCount() int {
	fake.setVersionMutex.RLock()
	defer fake.setVersionMutex.RUnlock()
	return len(fake.setVersionArgsForCall)
}

func (fake *FakeVersionDB) SetVersionCalls(stub func(context.Context, lager.Logger, *models.Version) error) {
	fake.setVersionMutex.Lock()
	defer fake.setVersionMutex.Unlock()
	fake.SetVersionStub = stub
}

func (fake *FakeVersionDB) SetVersionArgsForCall(i int) (context.Context, lager.Logger, *models.Version) {
	fake.setVersionMutex.RLock()
	defer fake.setVersionMutex.RUnlock()
	argsForCall := fake.setVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeVersionDB) SetVersionReturns(result1 error) {
	fake.setVersionMutex.Lock()
	defer fake.setVersionMutex.Unlock()
	fake.SetVersionStub = nil
	fake.setVersionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVersionDB) SetVersionReturnsOnCall(i int, result1 error) {
	fake.setVersionMutex.Lock()
	defer fake.setVersionMutex.Unlock()
	fake.SetVersionStub = nil
	if fake.setVersionReturnsOnCall == nil {
		fake.setVersionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setVersionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVersionDB) Version(arg1 context.Context, arg2 lager.Logger) (*models.Version, error) {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
	}{arg1, arg2})
	stub := fake.VersionStub
	fakeReturns := fake.versionReturns
	fake.recordInvocation("Version", []interface{}{arg1, arg2})
	fake.versionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVersionDB) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeVersionDB) VersionCalls(stub func(context.Context, lager.Logger) (*models.Version, error)) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = stub
}

func (fake *FakeVersionDB) VersionArgsForCall(i int) (context.Context, lager.Logger) {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	argsForCall := fake.versionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVersionDB) VersionReturns(result1 *models.Version, result2 error) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 *models.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeVersionDB) VersionReturnsOnCall(i int, result1 *models.Version, result2 error) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 *models.Version
			result2 error
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 *models.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeVersionDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setVersionMutex.RLock()
	defer fake.setVersionMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVersionDB) recordInvocation(key string, args []interface{}) {
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

var _ db.VersionDB = new(FakeVersionDB)
