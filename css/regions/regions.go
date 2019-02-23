// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package regions

import js "github.com/gowebapi/webapi/core/failjs"

import (
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/dom/domcore"
)

// using following types:
// dom.Node
// dom.Range
// domcore.EventTarget

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

// interface: NamedFlow
type NamedFlow struct {
	domcore.EventTarget
}

// NamedFlowFromJS is casting a js.Wrapper into NamedFlow.
func NamedFlowFromJS(value js.Wrapper) *NamedFlow {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &NamedFlow{}
	ret.Value_JS = input
	return ret
}

// Name returning attribute 'name' with
// type string (idl: DOMString).
func (_this *NamedFlow) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// Overset returning attribute 'overset' with
// type bool (idl: boolean).
func (_this *NamedFlow) Overset() bool {
	var ret bool
	value := _this.Value_JS.Get("overset")
	ret = (value).Bool()
	return ret
}

// FirstEmptyRegionIndex returning attribute 'firstEmptyRegionIndex' with
// type int (idl: short).
func (_this *NamedFlow) FirstEmptyRegionIndex() int {
	var ret int
	value := _this.Value_JS.Get("firstEmptyRegionIndex")
	ret = (value).Int()
	return ret
}

func (_this *NamedFlow) GetRegions() (_result []*Region) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getRegions", _args[0:_end]...)
	var (
		_converted []*Region // javascript: sequence<Region> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*Region, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *Region
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = RegionFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *NamedFlow) GetContent() (_result []*dom.Node) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getContent", _args[0:_end]...)
	var (
		_converted []*dom.Node // javascript: sequence<Node> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*dom.Node, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *dom.Node
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = dom.NodeFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *NamedFlow) GetRegionsByContent(node *dom.Node) (_result []*Region) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := node.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getRegionsByContent", _args[0:_end]...)
	var (
		_converted []*Region // javascript: sequence<Region> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*Region, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *Region
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = RegionFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

// interface: NamedFlowMap
type NamedFlowMap struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *NamedFlowMap) JSValue() js.Value {
	return _this.Value_JS
}

// NamedFlowMapFromJS is casting a js.Wrapper into NamedFlowMap.
func NamedFlowMapFromJS(value js.Wrapper) *NamedFlowMap {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &NamedFlowMap{}
	ret.Value_JS = input
	return ret
}

func (_this *NamedFlowMap) Get(flowName string) (_result *NamedFlow) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := flowName
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("get", _args[0:_end]...)
	var (
		_converted *NamedFlow // javascript: NamedFlow _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		_converted = NamedFlowFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *NamedFlowMap) Has(flowName string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := flowName
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

func (_this *NamedFlowMap) Set(flowName string, flowValue *NamedFlow) (_result *NamedFlowMap) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := flowName
	_args[0] = _p0
	_end++
	_p1 := flowValue.JSValue()
	_args[1] = _p1
	_end++
	_returned := _this.Value_JS.Call("set", _args[0:_end]...)
	var (
		_converted *NamedFlowMap // javascript: NamedFlowMap _what_return_name
	)
	_converted = NamedFlowMapFromJS(_returned)
	_result = _converted
	return
}

func (_this *NamedFlowMap) Delete(flowName string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := flowName
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("delete", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

// interface: Region
type Region struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Region) JSValue() js.Value {
	return _this.Value_JS
}

// RegionFromJS is casting a js.Wrapper into Region.
func RegionFromJS(value js.Wrapper) *Region {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Region{}
	ret.Value_JS = input
	return ret
}

// RegionOverset returning attribute 'regionOverset' with
// type string (idl: DOMString).
func (_this *Region) RegionOverset() string {
	var ret string
	value := _this.Value_JS.Get("regionOverset")
	ret = (value).String()
	return ret
}

func (_this *Region) GetRegionFlowRanges() (_result []*dom.Range) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getRegionFlowRanges", _args[0:_end]...)
	var (
		_converted []*dom.Range // javascript: sequence<Range> _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		__length0 := _returned.Length()
		__array0 := make([]*dom.Range, __length0, __length0)
		for __idx0 := 0; __idx0 < __length0; __idx0++ {
			var __seq_out0 *dom.Range
			__seq_in0 := _returned.Index(__idx0)
			__seq_out0 = dom.RangeFromJS(__seq_in0)
			__array0[__idx0] = __seq_out0
		}
		_converted = __array0
	}
	_result = _converted
	return
}
