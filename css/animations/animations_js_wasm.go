// Code generated by webidl-bind. DO NOT EDIT.

package animations

import "syscall/js"

import (
	"github.com/gowebapi/webapi/css/ccsom"
	"github.com/gowebapi/webapi/dom/domcore"
)

// using following types:
// ccsom.CSSRule
// ccsom.CSSRuleList
// ccsom.CSSStyleDeclaration
// domcore.Event

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

// dictionary: AnimationEventInit
type AnimationEventInit struct {
	Bubbles       bool
	Cancelable    bool
	Composed      bool
	AnimationName string
	ElapsedTime   float64
	PseudoElement string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *AnimationEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.AnimationName
	out.Set("animationName", value3)
	value4 := _this.ElapsedTime
	out.Set("elapsedTime", value4)
	value5 := _this.PseudoElement
	out.Set("pseudoElement", value5)
	return out
}

// AnimationEventInitFromJS is allocating a new
// AnimationEventInit object and copy all values from
// input javascript object
func AnimationEventInitFromJS(value js.Wrapper) *AnimationEventInit {
	input := value.JSValue()
	var out AnimationEventInit
	var (
		value0 bool    // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool    // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool    // javascript: boolean {composed Composed composed}
		value3 string  // javascript: DOMString {animationName AnimationName animationName}
		value4 float64 // javascript: double {elapsedTime ElapsedTime elapsedTime}
		value5 string  // javascript: DOMString {pseudoElement PseudoElement pseudoElement}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = (input.Get("animationName")).String()
	out.AnimationName = value3
	value4 = (input.Get("elapsedTime")).Float()
	out.ElapsedTime = value4
	value5 = (input.Get("pseudoElement")).String()
	out.PseudoElement = value5
	return &out
}

// interface: AnimationEvent
type AnimationEvent struct {
	domcore.Event
}

// AnimationEventFromJS is casting a js.Wrapper into AnimationEvent.
func AnimationEventFromJS(value js.Wrapper) *AnimationEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &AnimationEvent{}
	ret.Value_JS = input
	return ret
}

func NewAnimationEvent(_type string, animationEventInitDict *AnimationEventInit) (_result *AnimationEvent) {
	_klass := js.Global().Get("AnimationEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if animationEventInitDict != nil {
		_p1 := animationEventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *AnimationEvent // javascript: AnimationEvent _what_return_name
	)
	_converted = AnimationEventFromJS(_returned)
	_result = _converted
	return
}

// AnimationName returning attribute 'animationName' with
// type string (idl: DOMString).
func (_this *AnimationEvent) AnimationName() string {
	var ret string
	value := _this.Value_JS.Get("animationName")
	ret = (value).String()
	return ret
}

// ElapsedTime returning attribute 'elapsedTime' with
// type float64 (idl: double).
func (_this *AnimationEvent) ElapsedTime() float64 {
	var ret float64
	value := _this.Value_JS.Get("elapsedTime")
	ret = (value).Float()
	return ret
}

// PseudoElement returning attribute 'pseudoElement' with
// type string (idl: DOMString).
func (_this *AnimationEvent) PseudoElement() string {
	var ret string
	value := _this.Value_JS.Get("pseudoElement")
	ret = (value).String()
	return ret
}

// interface: CSSKeyframeRule
type CSSKeyframeRule struct {
	ccsom.CSSRule
}

// CSSKeyframeRuleFromJS is casting a js.Wrapper into CSSKeyframeRule.
func CSSKeyframeRuleFromJS(value js.Wrapper) *CSSKeyframeRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSKeyframeRule{}
	ret.Value_JS = input
	return ret
}

// KeyText returning attribute 'keyText' with
// type string (idl: DOMString).
func (_this *CSSKeyframeRule) KeyText() string {
	var ret string
	value := _this.Value_JS.Get("keyText")
	ret = (value).String()
	return ret
}

// SetKeyText setting attribute 'keyText' with
// type string (idl: DOMString).
func (_this *CSSKeyframeRule) SetKeyText(value string) {
	input := value
	_this.Value_JS.Set("keyText", input)
}

// Style returning attribute 'style' with
// type ccsom.CSSStyleDeclaration (idl: CSSStyleDeclaration).
func (_this *CSSKeyframeRule) Style() *ccsom.CSSStyleDeclaration {
	var ret *ccsom.CSSStyleDeclaration
	value := _this.Value_JS.Get("style")
	ret = ccsom.CSSStyleDeclarationFromJS(value)
	return ret
}

// interface: CSSKeyframesRule
type CSSKeyframesRule struct {
	ccsom.CSSRule
}

// CSSKeyframesRuleFromJS is casting a js.Wrapper into CSSKeyframesRule.
func CSSKeyframesRuleFromJS(value js.Wrapper) *CSSKeyframesRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSKeyframesRule{}
	ret.Value_JS = input
	return ret
}

// Name returning attribute 'name' with
// type string (idl: DOMString).
func (_this *CSSKeyframesRule) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// SetName setting attribute 'name' with
// type string (idl: DOMString).
func (_this *CSSKeyframesRule) SetName(value string) {
	input := value
	_this.Value_JS.Set("name", input)
}

// CssRules returning attribute 'cssRules' with
// type ccsom.CSSRuleList (idl: CSSRuleList).
func (_this *CSSKeyframesRule) CssRules() *ccsom.CSSRuleList {
	var ret *ccsom.CSSRuleList
	value := _this.Value_JS.Get("cssRules")
	ret = ccsom.CSSRuleListFromJS(value)
	return ret
}

func (_this *CSSKeyframesRule) AppendRule(rule string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := rule
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("appendRule", _args[0:_end]...)
	return
}

func (_this *CSSKeyframesRule) DeleteRule(_select string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := _select
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("deleteRule", _args[0:_end]...)
	return
}

func (_this *CSSKeyframesRule) FindRule(_select string) (_result *CSSKeyframeRule) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := _select
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("findRule", _args[0:_end]...)
	var (
		_converted *CSSKeyframeRule // javascript: CSSKeyframeRule _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		_converted = CSSKeyframeRuleFromJS(_returned)
	}
	_result = _converted
	return
}
