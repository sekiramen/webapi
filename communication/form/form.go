// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package form

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/file"
	"github.com/gowebapi/webapi/html"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// file.Blob
// html.HTMLFormElement
// javascript.PromiseFinally

// source idl files:
// promises.idl
// xhr.idl

// transform files:
// promises.go.md
// xhr.go.md

// ReleasableApiResource is used to release underlaying
// allocated resources.
type ReleasableApiResource interface {
	Release()
}

type releasableApiResourceList []ReleasableApiResource

func (a releasableApiResourceList) Release() {
	for _, v := range a {
		v.Release()
	}
}

// workaround for compiler error
func unused(value interface{}) {
	// TODO remove this method
}

type Union struct {
	Value js.Value
}

func (u *Union) JSValue() js.Value {
	return u.Value
}

func UnionFromJS(value js.Value) *Union {
	return &Union{Value: value}
}

// callback: FormDataForEach
type FormDataForEachFunc func(currentValue *Union, currentIndex int, listObj *FormData)

// FormDataForEach is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type FormDataForEach js.Func

func FormDataForEachToJS(callback FormDataForEachFunc) *FormDataForEach {
	if callback == nil {
		return nil
	}
	ret := FormDataForEach(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Union    // javascript: Union currentValue
			_p1 int       // javascript: long currentIndex
			_p2 *FormData // javascript: FormData listObj
		)
		_p0 = UnionFromJS(args[0])
		_p1 = (args[1]).Int()
		_p2 = FormDataFromJS(args[2])
		callback(_p0, _p1, _p2)
		// returning no return value
		return nil
	}))
	return &ret
}

func FormDataForEachFromJS(_value js.Value) FormDataForEachFunc {
	return func(currentValue *Union, currentIndex int, listObj *FormData) {
		var (
			_args [3]interface{}
			_end  int
		)
		_p0 := currentValue.JSValue()
		_args[0] = _p0
		_end++
		_p1 := currentIndex
		_args[1] = _p1
		_end++
		_p2 := listObj.JSValue()
		_args[2] = _p2
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnFulfilled
type PromiseFormDataOnFulfilledFunc func(value *FormData)

// PromiseFormDataOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseFormDataOnFulfilled js.Func

func PromiseFormDataOnFulfilledToJS(callback PromiseFormDataOnFulfilledFunc) *PromiseFormDataOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseFormDataOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *FormData // javascript: FormData value
		)
		_p0 = FormDataFromJS(args[0])
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseFormDataOnFulfilledFromJS(_value js.Value) PromiseFormDataOnFulfilledFunc {
	return func(value *FormData) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := value.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnRejected
type PromiseFormDataOnRejectedFunc func(reason js.Value)

// PromiseFormDataOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseFormDataOnRejected js.Func

func PromiseFormDataOnRejectedToJS(callback PromiseFormDataOnRejectedFunc) *PromiseFormDataOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseFormDataOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 js.Value // javascript: any reason
		)
		_p0 = args[0]
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseFormDataOnRejectedFromJS(_value js.Value) PromiseFormDataOnRejectedFunc {
	return func(reason js.Value) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := reason
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// dictionary: FormDataEntryIteratorValue
type FormDataEntryIteratorValue struct {
	Value []js.Value
	Done  bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *FormDataEntryIteratorValue) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := js.Global().Get("Array").New(len(_this.Value))
	for __idx0, __seq_in0 := range _this.Value {
		__seq_out0 := __seq_in0
		value0.SetIndex(__idx0, __seq_out0)
	}
	out.Set("value", value0)
	value1 := _this.Done
	out.Set("done", value1)
	return out
}

// FormDataEntryIteratorValueFromJS is allocating a new
// FormDataEntryIteratorValue object and copy all values from
// input javascript object
func FormDataEntryIteratorValueFromJS(value js.Wrapper) *FormDataEntryIteratorValue {
	input := value.JSValue()
	var out FormDataEntryIteratorValue
	var (
		value0 []js.Value // javascript: sequence<any> {value Value value}
		value1 bool       // javascript: boolean {done Done done}
	)
	__length0 := input.Get("value").Length()
	__array0 := make([]js.Value, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 js.Value
		__seq_in0 := input.Get("value").Index(__idx0)
		__seq_out0 = __seq_in0
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.Value = value0
	value1 = (input.Get("done")).Bool()
	out.Done = value1
	return &out
}

// dictionary: FormDataKeyIteratorValue
type FormDataKeyIteratorValue struct {
	Value *Union
	Done  bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *FormDataKeyIteratorValue) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Value.JSValue()
	out.Set("value", value0)
	value1 := _this.Done
	out.Set("done", value1)
	return out
}

// FormDataKeyIteratorValueFromJS is allocating a new
// FormDataKeyIteratorValue object and copy all values from
// input javascript object
func FormDataKeyIteratorValueFromJS(value js.Wrapper) *FormDataKeyIteratorValue {
	input := value.JSValue()
	var out FormDataKeyIteratorValue
	var (
		value0 *Union // javascript: Union {value Value value}
		value1 bool   // javascript: boolean {done Done done}
	)
	value0 = UnionFromJS(input.Get("value"))
	out.Value = value0
	value1 = (input.Get("done")).Bool()
	out.Done = value1
	return &out
}

// dictionary: FormDataValueIteratorValue
type FormDataValueIteratorValue struct {
	Value *Union
	Done  bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *FormDataValueIteratorValue) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Value.JSValue()
	out.Set("value", value0)
	value1 := _this.Done
	out.Set("done", value1)
	return out
}

// FormDataValueIteratorValueFromJS is allocating a new
// FormDataValueIteratorValue object and copy all values from
// input javascript object
func FormDataValueIteratorValueFromJS(value js.Wrapper) *FormDataValueIteratorValue {
	input := value.JSValue()
	var out FormDataValueIteratorValue
	var (
		value0 *Union // javascript: Union {value Value value}
		value1 bool   // javascript: boolean {done Done done}
	)
	value0 = UnionFromJS(input.Get("value"))
	out.Value = value0
	value1 = (input.Get("done")).Bool()
	out.Done = value1
	return &out
}

// interface: FormData
type FormData struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FormData) JSValue() js.Value {
	return _this.Value_JS
}

// FormDataFromJS is casting a js.Wrapper into FormData.
func FormDataFromJS(value js.Wrapper) *FormData {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FormData{}
	ret.Value_JS = input
	return ret
}

func NewFormData(form *html.HTMLFormElement) (_result *FormData) {
	_klass := js.Global().Get("FormData")
	var (
		_args [1]interface{}
		_end  int
	)
	if form != nil {
		_p0 := form.JSValue()
		_args[0] = _p0
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *FormData // javascript: FormData _what_return_name
	)
	_converted = FormDataFromJS(_returned)
	_result = _converted
	return
}

func (_this *FormData) Append(name string, value string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_p1 := value
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("append", _args[0:_end]...)
	return
}

func (_this *FormData) Append2(name string, blobValue *file.Blob, filename *string) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_p1 := blobValue.JSValue()
	_args[1] = _p1
	_end++
	if filename != nil {
		_p2 := filename
		_args[2] = _p2
		_end++
	}
	_this.Value_JS.Call("append", _args[0:_end]...)
	return
}

func (_this *FormData) Delete(name string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("delete", _args[0:_end]...)
	return
}

func (_this *FormData) Get(name string) (_result *Union) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("get", _args[0:_end]...)
	var (
		_converted *Union // javascript: Union _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = UnionFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *FormData) GetAll(name string) (_result []*Union) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getAll", _args[0:_end]...)
	var (
		_converted []*Union // javascript: sequence<Union> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*Union, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *Union
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = UnionFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *FormData) Has(name string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("has", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *FormData) Set(name string, value string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_p1 := value
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("set", _args[0:_end]...)
	return
}

func (_this *FormData) Set2(name string, blobValue *file.Blob, filename *string) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_p1 := blobValue.JSValue()
	_args[1] = _p1
	_end++
	if filename != nil {
		_p2 := filename
		_args[2] = _p2
		_end++
	}
	_this.Value_JS.Call("set", _args[0:_end]...)
	return
}

func (_this *FormData) Entries() (_result *FormDataEntryIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("entries", _args[0:_end]...)
	var (
		_converted *FormDataEntryIterator // javascript: FormDataEntryIterator _what_return_name
	)
	_converted = FormDataEntryIteratorFromJS(_returned)
	_result = _converted
	return
}

func (_this *FormData) ForEach(callback *FormDataForEach, optionalThisForCallbackArgument interface{}) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if callback != nil {
		__callback0 = (*callback).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if optionalThisForCallbackArgument != nil {
		_p1 := optionalThisForCallbackArgument
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("forEach", _args[0:_end]...)
	return
}

func (_this *FormData) Keys() (_result *FormDataKeyIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("keys", _args[0:_end]...)
	var (
		_converted *FormDataKeyIterator // javascript: FormDataKeyIterator _what_return_name
	)
	_converted = FormDataKeyIteratorFromJS(_returned)
	_result = _converted
	return
}

func (_this *FormData) Values() (_result *FormDataValueIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("values", _args[0:_end]...)
	var (
		_converted *FormDataValueIterator // javascript: FormDataValueIterator _what_return_name
	)
	_converted = FormDataValueIteratorFromJS(_returned)
	_result = _converted
	return
}

// interface: FormDataEntryIterator
type FormDataEntryIterator struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FormDataEntryIterator) JSValue() js.Value {
	return _this.Value_JS
}

// FormDataEntryIteratorFromJS is casting a js.Wrapper into FormDataEntryIterator.
func FormDataEntryIteratorFromJS(value js.Wrapper) *FormDataEntryIterator {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FormDataEntryIterator{}
	ret.Value_JS = input
	return ret
}

func (_this *FormDataEntryIterator) Next() (_result *FormDataEntryIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("next", _args[0:_end]...)
	var (
		_converted *FormDataEntryIteratorValue // javascript: FormDataEntryIteratorValue _what_return_name
	)
	_converted = FormDataEntryIteratorValueFromJS(_returned)
	_result = _converted
	return
}

// interface: FormDataKeyIterator
type FormDataKeyIterator struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FormDataKeyIterator) JSValue() js.Value {
	return _this.Value_JS
}

// FormDataKeyIteratorFromJS is casting a js.Wrapper into FormDataKeyIterator.
func FormDataKeyIteratorFromJS(value js.Wrapper) *FormDataKeyIterator {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FormDataKeyIterator{}
	ret.Value_JS = input
	return ret
}

func (_this *FormDataKeyIterator) Next() (_result *FormDataKeyIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("next", _args[0:_end]...)
	var (
		_converted *FormDataKeyIteratorValue // javascript: FormDataKeyIteratorValue _what_return_name
	)
	_converted = FormDataKeyIteratorValueFromJS(_returned)
	_result = _converted
	return
}

// interface: FormDataValueIterator
type FormDataValueIterator struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FormDataValueIterator) JSValue() js.Value {
	return _this.Value_JS
}

// FormDataValueIteratorFromJS is casting a js.Wrapper into FormDataValueIterator.
func FormDataValueIteratorFromJS(value js.Wrapper) *FormDataValueIterator {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FormDataValueIterator{}
	ret.Value_JS = input
	return ret
}

func (_this *FormDataValueIterator) Next() (_result *FormDataValueIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("next", _args[0:_end]...)
	var (
		_converted *FormDataValueIteratorValue // javascript: FormDataValueIteratorValue _what_return_name
	)
	_converted = FormDataValueIteratorValueFromJS(_returned)
	_result = _converted
	return
}

// interface: Promise
type PromiseFormData struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseFormData) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseFormDataFromJS is casting a js.Wrapper into PromiseFormData.
func PromiseFormDataFromJS(value js.Wrapper) *PromiseFormData {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PromiseFormData{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseFormData) Then(onFulfilled *PromiseFormDataOnFulfilled, onRejected *PromiseFormDataOnRejected) (_result *PromiseFormData) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFulfilled != nil {
		__callback0 = (*onFulfilled).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if onRejected != nil {

		var __callback1 js.Value
		if onRejected != nil {
			__callback1 = (*onRejected).Value
		} else {
			__callback1 = js.Null()
		}
		_p1 := __callback1
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("then", _args[0:_end]...)
	var (
		_converted *PromiseFormData // javascript: Promise _what_return_name
	)
	_converted = PromiseFormDataFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseFormData) Catch(onRejected *PromiseFormDataOnRejected) (_result *PromiseFormData) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onRejected != nil {
		__callback0 = (*onRejected).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("catch", _args[0:_end]...)
	var (
		_converted *PromiseFormData // javascript: Promise _what_return_name
	)
	_converted = PromiseFormDataFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseFormData) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseFormData) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFinally != nil {
		__callback0 = (*onFinally).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("finally", _args[0:_end]...)
	var (
		_converted *PromiseFormData // javascript: Promise _what_return_name
	)
	_converted = PromiseFormDataFromJS(_returned)
	_result = _converted
	return
}
