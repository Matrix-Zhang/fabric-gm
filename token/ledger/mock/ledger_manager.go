// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/Matrix-Zhang/fabric-gm/token/ledger"
)

type LedgerManager struct {
	GetLedgerReaderStub        func(channel string) (ledger.LedgerReader, error)
	getLedgerReaderMutex       sync.RWMutex
	getLedgerReaderArgsForCall []struct {
		channel string
	}
	getLedgerReaderReturns struct {
		result1 ledger.LedgerReader
		result2 error
	}
	getLedgerReaderReturnsOnCall map[int]struct {
		result1 ledger.LedgerReader
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *LedgerManager) GetLedgerReader(channel string) (ledger.LedgerReader, error) {
	fake.getLedgerReaderMutex.Lock()
	ret, specificReturn := fake.getLedgerReaderReturnsOnCall[len(fake.getLedgerReaderArgsForCall)]
	fake.getLedgerReaderArgsForCall = append(fake.getLedgerReaderArgsForCall, struct {
		channel string
	}{channel})
	fake.recordInvocation("GetLedgerReader", []interface{}{channel})
	fake.getLedgerReaderMutex.Unlock()
	if fake.GetLedgerReaderStub != nil {
		return fake.GetLedgerReaderStub(channel)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getLedgerReaderReturns.result1, fake.getLedgerReaderReturns.result2
}

func (fake *LedgerManager) GetLedgerReaderCallCount() int {
	fake.getLedgerReaderMutex.RLock()
	defer fake.getLedgerReaderMutex.RUnlock()
	return len(fake.getLedgerReaderArgsForCall)
}

func (fake *LedgerManager) GetLedgerReaderArgsForCall(i int) string {
	fake.getLedgerReaderMutex.RLock()
	defer fake.getLedgerReaderMutex.RUnlock()
	return fake.getLedgerReaderArgsForCall[i].channel
}

func (fake *LedgerManager) GetLedgerReaderReturns(result1 ledger.LedgerReader, result2 error) {
	fake.GetLedgerReaderStub = nil
	fake.getLedgerReaderReturns = struct {
		result1 ledger.LedgerReader
		result2 error
	}{result1, result2}
}

func (fake *LedgerManager) GetLedgerReaderReturnsOnCall(i int, result1 ledger.LedgerReader, result2 error) {
	fake.GetLedgerReaderStub = nil
	if fake.getLedgerReaderReturnsOnCall == nil {
		fake.getLedgerReaderReturnsOnCall = make(map[int]struct {
			result1 ledger.LedgerReader
			result2 error
		})
	}
	fake.getLedgerReaderReturnsOnCall[i] = struct {
		result1 ledger.LedgerReader
		result2 error
	}{result1, result2}
}

func (fake *LedgerManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getLedgerReaderMutex.RLock()
	defer fake.getLedgerReaderMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *LedgerManager) recordInvocation(key string, args []interface{}) {
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

var _ ledger.LedgerManager = new(LedgerManager)
