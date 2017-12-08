// generated by "charlatan -dir=testdata/funcer -output=testdata/funcer/funcer.go Funcer".  DO NOT EDIT.

package main

import (
	"reflect"
	"testing"
)

// FuncParameterInvocation represents a single call of FakeFuncer.FuncParameter
type FuncParameterInvocation struct {
	Parameters struct {
		Ident1 func(string) string
	}
}

// FuncReturnInvocation represents a single call of FakeFuncer.FuncReturn
type FuncReturnInvocation struct {
	Results struct {
		Ident1 func(string) string
	}
}

/*
FakeFuncer is a mock implementation of Funcer for testing.
Use it in your tests as in this example:

	package example

	func TestWithFuncer(t *testing.T) {
		f := &main.FakeFuncer{
			FuncParameterHook: func(ident1 func(string) string) () {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeFuncParameter ...
		f.AssertFuncParameterCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeFuncParameter.
*/
type FakeFuncer struct {
	FuncParameterHook func(func(string) string)
	FuncReturnHook    func() func(string) string

	FuncParameterCalls []*FuncParameterInvocation
	FuncReturnCalls    []*FuncReturnInvocation
}

// NewFakeFuncerDefaultPanic returns an instance of FakeFuncer with all hooks configured to panic
func NewFakeFuncerDefaultPanic() *FakeFuncer {
	return &FakeFuncer{
		FuncParameterHook: func(func(string) string) {
			panic("Unexpected call to Funcer.FuncParameter")
		},
		FuncReturnHook: func() (ident1 func(string) string) {
			panic("Unexpected call to Funcer.FuncReturn")
		},
	}
}

// NewFakeFuncerDefaultFatal returns an instance of FakeFuncer with all hooks configured to call t.Fatal
func NewFakeFuncerDefaultFatal(t *testing.T) *FakeFuncer {
	return &FakeFuncer{
		FuncParameterHook: func(func(string) string) {
			t.Fatal("Unexpected call to Funcer.FuncParameter")
			return
		},
		FuncReturnHook: func() (ident1 func(string) string) {
			t.Fatal("Unexpected call to Funcer.FuncReturn")
			return
		},
	}
}

// NewFakeFuncerDefaultError returns an instance of FakeFuncer with all hooks configured to call t.Error
func NewFakeFuncerDefaultError(t *testing.T) *FakeFuncer {
	return &FakeFuncer{
		FuncParameterHook: func(func(string) string) {
			t.Error("Unexpected call to Funcer.FuncParameter")
			return
		},
		FuncReturnHook: func() (ident1 func(string) string) {
			t.Error("Unexpected call to Funcer.FuncReturn")
			return
		},
	}
}

func (f *FakeFuncer) Reset() {
	f.FuncParameterCalls = []*FuncParameterInvocation{}
	f.FuncReturnCalls = []*FuncReturnInvocation{}
}

func (_f1 *FakeFuncer) FuncParameter(ident1 func(string) string) {
	invocation := new(FuncParameterInvocation)

	invocation.Parameters.Ident1 = ident1

	_f1.FuncParameterHook(ident1)

	_f1.FuncParameterCalls = append(_f1.FuncParameterCalls, invocation)

	return
}

// FuncParameterCalled returns true if FakeFuncer.FuncParameter was called
func (f *FakeFuncer) FuncParameterCalled() bool {
	return len(f.FuncParameterCalls) != 0
}

// AssertFuncParameterCalled calls t.Error if FakeFuncer.FuncParameter was not called
func (f *FakeFuncer) AssertFuncParameterCalled(t *testing.T) {
	t.Helper()
	if len(f.FuncParameterCalls) == 0 {
		t.Error("FakeFuncer.FuncParameter not called, expected at least one")
	}
}

// FuncParameterNotCalled returns true if FakeFuncer.FuncParameter was not called
func (f *FakeFuncer) FuncParameterNotCalled() bool {
	return len(f.FuncParameterCalls) == 0
}

// AssertFuncParameterNotCalled calls t.Error if FakeFuncer.FuncParameter was called
func (f *FakeFuncer) AssertFuncParameterNotCalled(t *testing.T) {
	t.Helper()
	if len(f.FuncParameterCalls) != 0 {
		t.Error("FakeFuncer.FuncParameter called, expected none")
	}
}

// FuncParameterCalledOnce returns true if FakeFuncer.FuncParameter was called exactly once
func (f *FakeFuncer) FuncParameterCalledOnce() bool {
	return len(f.FuncParameterCalls) == 1
}

// AssertFuncParameterCalledOnce calls t.Error if FakeFuncer.FuncParameter was not called exactly once
func (f *FakeFuncer) AssertFuncParameterCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.FuncParameterCalls) != 1 {
		t.Errorf("FakeFuncer.FuncParameter called %d times, expected 1", len(f.FuncParameterCalls))
	}
}

// FuncParameterCalledN returns true if FakeFuncer.FuncParameter was called at least n times
func (f *FakeFuncer) FuncParameterCalledN(n int) bool {
	return len(f.FuncParameterCalls) >= n
}

// AssertFuncParameterCalledN calls t.Error if FakeFuncer.FuncParameter was called less than n times
func (f *FakeFuncer) AssertFuncParameterCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.FuncParameterCalls) < n {
		t.Errorf("FakeFuncer.FuncParameter called %d times, expected >= %d", len(f.FuncParameterCalls), n)
	}
}

// FuncParameterCalledWith returns true if FakeFuncer.FuncParameter was called with the given values
func (_f2 *FakeFuncer) FuncParameterCalledWith(ident1 func(string) string) (found bool) {
	for _, call := range _f2.FuncParameterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	return
}

// AssertFuncParameterCalledWith calls t.Error if FakeFuncer.FuncParameter was not called with the given values
func (_f3 *FakeFuncer) AssertFuncParameterCalledWith(t *testing.T, ident1 func(string) string) {
	t.Helper()
	var found bool
	for _, call := range _f3.FuncParameterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeFuncer.FuncParameter not called with expected parameters")
	}
}

// FuncParameterCalledOnceWith returns true if FakeFuncer.FuncParameter was called exactly once with the given values
func (_f4 *FakeFuncer) FuncParameterCalledOnceWith(ident1 func(string) string) bool {
	var count int
	for _, call := range _f4.FuncParameterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	return count == 1
}

// AssertFuncParameterCalledOnceWith calls t.Error if FakeFuncer.FuncParameter was not called exactly once with the given values
func (_f5 *FakeFuncer) AssertFuncParameterCalledOnceWith(t *testing.T, ident1 func(string) string) {
	t.Helper()
	var count int
	for _, call := range _f5.FuncParameterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeFuncer.FuncParameter called %d times with expected parameters, expected one", count)
	}
}

func (_f6 *FakeFuncer) FuncReturn() (ident1 func(string) string) {
	invocation := new(FuncReturnInvocation)

	ident1 = _f6.FuncReturnHook()

	invocation.Results.Ident1 = ident1

	_f6.FuncReturnCalls = append(_f6.FuncReturnCalls, invocation)

	return
}

// FuncReturnCalled returns true if FakeFuncer.FuncReturn was called
func (f *FakeFuncer) FuncReturnCalled() bool {
	return len(f.FuncReturnCalls) != 0
}

// AssertFuncReturnCalled calls t.Error if FakeFuncer.FuncReturn was not called
func (f *FakeFuncer) AssertFuncReturnCalled(t *testing.T) {
	t.Helper()
	if len(f.FuncReturnCalls) == 0 {
		t.Error("FakeFuncer.FuncReturn not called, expected at least one")
	}
}

// FuncReturnNotCalled returns true if FakeFuncer.FuncReturn was not called
func (f *FakeFuncer) FuncReturnNotCalled() bool {
	return len(f.FuncReturnCalls) == 0
}

// AssertFuncReturnNotCalled calls t.Error if FakeFuncer.FuncReturn was called
func (f *FakeFuncer) AssertFuncReturnNotCalled(t *testing.T) {
	t.Helper()
	if len(f.FuncReturnCalls) != 0 {
		t.Error("FakeFuncer.FuncReturn called, expected none")
	}
}

// FuncReturnCalledOnce returns true if FakeFuncer.FuncReturn was called exactly once
func (f *FakeFuncer) FuncReturnCalledOnce() bool {
	return len(f.FuncReturnCalls) == 1
}

// AssertFuncReturnCalledOnce calls t.Error if FakeFuncer.FuncReturn was not called exactly once
func (f *FakeFuncer) AssertFuncReturnCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.FuncReturnCalls) != 1 {
		t.Errorf("FakeFuncer.FuncReturn called %d times, expected 1", len(f.FuncReturnCalls))
	}
}

// FuncReturnCalledN returns true if FakeFuncer.FuncReturn was called at least n times
func (f *FakeFuncer) FuncReturnCalledN(n int) bool {
	return len(f.FuncReturnCalls) >= n
}

// AssertFuncReturnCalledN calls t.Error if FakeFuncer.FuncReturn was called less than n times
func (f *FakeFuncer) AssertFuncReturnCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.FuncReturnCalls) < n {
		t.Errorf("FakeFuncer.FuncReturn called %d times, expected >= %d", len(f.FuncReturnCalls), n)
	}
}
