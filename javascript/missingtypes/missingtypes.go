// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package missingtypes

import js "github.com/gowebapi/webapi/core/js"

// using following types:

// source idl files:
// missingtypes.idl

// transform files:
// missingtypes.go.md

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

// class: Date
type Date struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Date) JSValue() js.Value {
	return _this.Value_JS
}

// DateFromJS is casting a js.Wrapper into Date.
func DateFromJS(value js.Wrapper) *Date {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Date{}
	ret.Value_JS = input
	return ret
}

// class: Dictionary
type Dictionary struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Dictionary) JSValue() js.Value {
	return _this.Value_JS
}

// DictionaryFromJS is casting a js.Wrapper into Dictionary.
func DictionaryFromJS(value js.Wrapper) *Dictionary {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Dictionary{}
	ret.Value_JS = input
	return ret
}

// class: WritableStream
type WritableStream struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *WritableStream) JSValue() js.Value {
	return _this.Value_JS
}

// WritableStreamFromJS is casting a js.Wrapper into WritableStream.
func WritableStreamFromJS(value js.Wrapper) *WritableStream {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &WritableStream{}
	ret.Value_JS = input
	return ret
}
