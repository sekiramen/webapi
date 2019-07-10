// Code generated by webidl-bind. DO NOT EDIT.

package domcore

import "syscall/js"

import (
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// javascript.Promise

// source idl files:
// dom.idl
// dom.webidl.idl
// html.idl
// page-visibility.idl
// service-workers.idl

// transform files:
// dom.go.md
// dom.go.md
// html.go.md
// page-visibility.go.md
// service-workers.go.md

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

// enum: VisibilityState
type VisibilityState int

const (
	HiddenVisibilityState VisibilityState = iota
	VisibleVisibilityState
	PrerenderVisibilityState
)

var visibilityStateToWasmTable = []string{
	"hidden", "visible", "prerender",
}

var visibilityStateFromWasmTable = map[string]VisibilityState{
	"hidden": HiddenVisibilityState, "visible": VisibleVisibilityState, "prerender": PrerenderVisibilityState,
}

// JSValue is converting this enum into a java object
func (this *VisibilityState) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this VisibilityState) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(visibilityStateToWasmTable) {
		return visibilityStateToWasmTable[idx]
	}
	panic("unknown input value")
}

// VisibilityStateFromJS is converting a javascript value into
// a VisibilityState enum value.
func VisibilityStateFromJS(value js.Value) VisibilityState {
	key := value.String()
	conv, ok := visibilityStateFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: EventHandlerNonNull
type EventHandlerFunc func(event *Event) interface{}

// EventHandler is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type EventHandler js.Func

func EventHandlerToJS(callback EventHandlerFunc) *EventHandler {
	if callback == nil {
		return nil
	}
	ret := EventHandler(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Event // javascript: Event event
		)
		_p0 = EventFromJS(args[0])
		_returned := callback(_p0)
		_converted := _returned
		return _converted
	}))
	return &ret
}

func EventHandlerFromJS(_value js.Value) EventHandlerFunc {
	return func(event *Event) (_result interface{}) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := event.JSValue()
		_args[0] = _p0
		_end++
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted js.Value // javascript: any
		)
		_converted = _returned
		_result = _converted
		return
	}
}

// dictionary: AddEventListenerOptions
type AddEventListenerOptions struct {
	Capture bool
	Passive bool
	Once    bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *AddEventListenerOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Capture
	out.Set("capture", value0)
	value1 := _this.Passive
	out.Set("passive", value1)
	value2 := _this.Once
	out.Set("once", value2)
	return out
}

// AddEventListenerOptionsFromJS is allocating a new
// AddEventListenerOptions object and copy all values from
// input javascript object
func AddEventListenerOptionsFromJS(value js.Wrapper) *AddEventListenerOptions {
	input := value.JSValue()
	var out AddEventListenerOptions
	var (
		value0 bool // javascript: boolean {capture Capture capture}
		value1 bool // javascript: boolean {passive Passive passive}
		value2 bool // javascript: boolean {once Once once}
	)
	value0 = (input.Get("capture")).Bool()
	out.Capture = value0
	value1 = (input.Get("passive")).Bool()
	out.Passive = value1
	value2 = (input.Get("once")).Bool()
	out.Once = value2
	return &out
}

// dictionary: CustomEventInit
type CustomEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Detail     js.Value
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CustomEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Detail
	out.Set("detail", value3)
	return out
}

// CustomEventInitFromJS is allocating a new
// CustomEventInit object and copy all values from
// input javascript object
func CustomEventInitFromJS(value js.Wrapper) *CustomEventInit {
	input := value.JSValue()
	var out CustomEventInit
	var (
		value0 bool     // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool     // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool     // javascript: boolean {composed Composed composed}
		value3 js.Value // javascript: any {detail Detail detail}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = input.Get("detail")
	out.Detail = value3
	return &out
}

// dictionary: EventInit
type EventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *EventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	return out
}

// EventInitFromJS is allocating a new
// EventInit object and copy all values from
// input javascript object
func EventInitFromJS(value js.Wrapper) *EventInit {
	input := value.JSValue()
	var out EventInit
	var (
		value0 bool // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool // javascript: boolean {composed Composed composed}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	return &out
}

// dictionary: EventListenerOptions
type EventListenerOptions struct {
	Capture bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *EventListenerOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Capture
	out.Set("capture", value0)
	return out
}

// EventListenerOptionsFromJS is allocating a new
// EventListenerOptions object and copy all values from
// input javascript object
func EventListenerOptionsFromJS(value js.Wrapper) *EventListenerOptions {
	input := value.JSValue()
	var out EventListenerOptions
	var (
		value0 bool // javascript: boolean {capture Capture capture}
	)
	value0 = (input.Get("capture")).Bool()
	out.Capture = value0
	return &out
}

// dictionary: ExtendableEventInit
type ExtendableEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ExtendableEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	return out
}

// ExtendableEventInitFromJS is allocating a new
// ExtendableEventInit object and copy all values from
// input javascript object
func ExtendableEventInitFromJS(value js.Wrapper) *ExtendableEventInit {
	input := value.JSValue()
	var out ExtendableEventInit
	var (
		value0 bool // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool // javascript: boolean {composed Composed composed}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	return &out
}

// interface: AbortController
type AbortController struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *AbortController) JSValue() js.Value {
	return _this.Value_JS
}

// AbortControllerFromJS is casting a js.Wrapper into AbortController.
func AbortControllerFromJS(value js.Wrapper) *AbortController {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &AbortController{}
	ret.Value_JS = input
	return ret
}

func NewAbortController() (_result *AbortController) {
	_klass := js.Global().Get("AbortController")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *AbortController // javascript: AbortController _what_return_name
	)
	_converted = AbortControllerFromJS(_returned)
	_result = _converted
	return
}

// Signal returning attribute 'signal' with
// type AbortSignal (idl: AbortSignal).
func (_this *AbortController) Signal() *AbortSignal {
	var ret *AbortSignal
	value := _this.Value_JS.Get("signal")
	ret = AbortSignalFromJS(value)
	return ret
}

func (_this *AbortController) Abort() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("abort", _args[0:_end]...)
	return
}

// interface: AbortSignal
type AbortSignal struct {
	EventTarget
}

// AbortSignalFromJS is casting a js.Wrapper into AbortSignal.
func AbortSignalFromJS(value js.Wrapper) *AbortSignal {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &AbortSignal{}
	ret.Value_JS = input
	return ret
}

// Aborted returning attribute 'aborted' with
// type bool (idl: boolean).
func (_this *AbortSignal) Aborted() bool {
	var ret bool
	value := _this.Value_JS.Get("aborted")
	ret = (value).Bool()
	return ret
}

// Onabort returning attribute 'onabort' with
// type EventHandler (idl: EventHandlerNonNull).
func (_this *AbortSignal) Onabort() EventHandlerFunc {
	var ret EventHandlerFunc
	value := _this.Value_JS.Get("onabort")
	if value.Type() != js.TypeNull {
		ret = EventHandlerFromJS(value)
	}
	return ret
}

// SetOnabort setting attribute 'onabort' with
// type EventHandler (idl: EventHandlerNonNull).
func (_this *AbortSignal) SetOnabort(value *EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onabort", input)
}

// interface: CustomEvent
type CustomEvent struct {
	Event
}

// CustomEventFromJS is casting a js.Wrapper into CustomEvent.
func CustomEventFromJS(value js.Wrapper) *CustomEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CustomEvent{}
	ret.Value_JS = input
	return ret
}

func NewCustomEvent(_type string, eventInitDict *CustomEventInit) (_result *CustomEvent) {
	_klass := js.Global().Get("CustomEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *CustomEvent // javascript: CustomEvent _what_return_name
	)
	_converted = CustomEventFromJS(_returned)
	_result = _converted
	return
}

// Detail returning attribute 'detail' with
// type Any (idl: any).
func (_this *CustomEvent) Detail() js.Value {
	var ret js.Value
	value := _this.Value_JS.Get("detail")
	ret = value
	return ret
}

func (_this *CustomEvent) InitCustomEvent(_type string, bubbles *bool, cancelable *bool, detail interface{}) {
	var (
		_args [4]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if bubbles != nil {
		_p1 := bubbles
		_args[1] = _p1
		_end++
	}
	if cancelable != nil {
		_p2 := cancelable
		_args[2] = _p2
		_end++
	}
	if detail != nil {
		_p3 := detail
		_args[3] = _p3
		_end++
	}
	_this.Value_JS.Call("initCustomEvent", _args[0:_end]...)
	return
}

// interface: DOMException
type DOMException struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DOMException) JSValue() js.Value {
	return _this.Value_JS
}

// DOMExceptionFromJS is casting a js.Wrapper into DOMException.
func DOMExceptionFromJS(value js.Wrapper) *DOMException {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DOMException{}
	ret.Value_JS = input
	return ret
}

const (
	INDEX_SIZE_ERR              int = 1
	DOMSTRING_SIZE_ERR          int = 2
	HIERARCHY_REQUEST_ERR       int = 3
	WRONG_DOCUMENT_ERR          int = 4
	INVALID_CHARACTER_ERR       int = 5
	NO_DATA_ALLOWED_ERR         int = 6
	NO_MODIFICATION_ALLOWED_ERR int = 7
	NOT_FOUND_ERR               int = 8
	NOT_SUPPORTED_ERR           int = 9
	INUSE_ATTRIBUTE_ERR         int = 10
	INVALID_STATE_ERR           int = 11
	SYNTAX_ERR                  int = 12
	INVALID_MODIFICATION_ERR    int = 13
	NAMESPACE_ERR               int = 14
	INVALID_ACCESS_ERR          int = 15
	VALIDATION_ERR              int = 16
	TYPE_MISMATCH_ERR           int = 17
	SECURITY_ERR                int = 18
	NETWORK_ERR                 int = 19
	ABORT_ERR                   int = 20
	URL_MISMATCH_ERR            int = 21
	QUOTA_EXCEEDED_ERR          int = 22
	TIMEOUT_ERR                 int = 23
	INVALID_NODE_TYPE_ERR       int = 24
	DATA_CLONE_ERR              int = 25
)

func NewDOMException(message *string, name *string) (_result *DOMException) {
	_klass := js.Global().Get("DOMException")
	var (
		_args [2]interface{}
		_end  int
	)
	if message != nil {
		_p0 := message
		_args[0] = _p0
		_end++
	}
	if name != nil {
		_p1 := name
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *DOMException // javascript: DOMException _what_return_name
	)
	_converted = DOMExceptionFromJS(_returned)
	_result = _converted
	return
}

// Name returning attribute 'name' with
// type string (idl: DOMString).
func (_this *DOMException) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// Message returning attribute 'message' with
// type string (idl: DOMString).
func (_this *DOMException) Message() string {
	var ret string
	value := _this.Value_JS.Get("message")
	ret = (value).String()
	return ret
}

// Code returning attribute 'code' with
// type int (idl: unsigned short).
func (_this *DOMException) Code() int {
	var ret int
	value := _this.Value_JS.Get("code")
	ret = (value).Int()
	return ret
}

// interface: DOMStringList
type DOMStringList struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DOMStringList) JSValue() js.Value {
	return _this.Value_JS
}

// DOMStringListFromJS is casting a js.Wrapper into DOMStringList.
func DOMStringListFromJS(value js.Wrapper) *DOMStringList {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DOMStringList{}
	ret.Value_JS = input
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *DOMStringList) Length() uint {
	var ret uint
	value := _this.Value_JS.Get("length")
	ret = (uint)((value).Int())
	return ret
}

func (_this *DOMStringList) Item(index uint) (_result *string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("item", _args[0:_end]...)
	var (
		_converted *string // javascript: DOMString _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		__tmp := (_returned).String()
		_converted = &__tmp
	}
	_result = _converted
	return
}

func (_this *DOMStringList) Contains(string string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := string
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("contains", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

// interface: DOMStringMap
type DOMStringMap struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DOMStringMap) JSValue() js.Value {
	return _this.Value_JS
}

// DOMStringMapFromJS is casting a js.Wrapper into DOMStringMap.
func DOMStringMapFromJS(value js.Wrapper) *DOMStringMap {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DOMStringMap{}
	ret.Value_JS = input
	return ret
}

// interface: DOMTokenList
type DOMTokenList struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DOMTokenList) JSValue() js.Value {
	return _this.Value_JS
}

// DOMTokenListFromJS is casting a js.Wrapper into DOMTokenList.
func DOMTokenListFromJS(value js.Wrapper) *DOMTokenList {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DOMTokenList{}
	ret.Value_JS = input
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *DOMTokenList) Length() uint {
	var ret uint
	value := _this.Value_JS.Get("length")
	ret = (uint)((value).Int())
	return ret
}

// Value returning attribute 'value' with
// type string (idl: DOMString).
func (_this *DOMTokenList) Value() string {
	var ret string
	value := _this.Value_JS.Get("value")
	ret = (value).String()
	return ret
}

// ToString is an alias for Value.
func (_this *DOMTokenList) ToString() string {
	return _this.Value()
}

// SetValue setting attribute 'value' with
// type string (idl: DOMString).
func (_this *DOMTokenList) SetValue(value string) {
	input := value
	_this.Value_JS.Set("value", input)
}

func (_this *DOMTokenList) Item(index uint) (_result *string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("item", _args[0:_end]...)
	var (
		_converted *string // javascript: DOMString _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		__tmp := (_returned).String()
		_converted = &__tmp
	}
	_result = _converted
	return
}

func (_this *DOMTokenList) Contains(token string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := token
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("contains", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *DOMTokenList) Add(tokens ...string) {
	var (
		_args []interface{} = make([]interface{}, 0+len(tokens))
		_end  int
	)
	for _, __in := range tokens {
		__out := __in
		_args[_end] = __out
		_end++
	}
	_this.Value_JS.Call("add", _args[0:_end]...)
	return
}

func (_this *DOMTokenList) Remove(tokens ...string) {
	var (
		_args []interface{} = make([]interface{}, 0+len(tokens))
		_end  int
	)
	for _, __in := range tokens {
		__out := __in
		_args[_end] = __out
		_end++
	}
	_this.Value_JS.Call("remove", _args[0:_end]...)
	return
}

func (_this *DOMTokenList) Toggle(token string, force *bool) (_result bool) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := token
	_args[0] = _p0
	_end++
	if force != nil {
		_p1 := force
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("toggle", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *DOMTokenList) Replace(token string, newToken string) (_result bool) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := token
	_args[0] = _p0
	_end++
	_p1 := newToken
	_args[1] = _p1
	_end++
	_returned := _this.Value_JS.Call("replace", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *DOMTokenList) Supports(token string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := token
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("supports", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

// interface: Event
type Event struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Event) JSValue() js.Value {
	return _this.Value_JS
}

// EventFromJS is casting a js.Wrapper into Event.
func EventFromJS(value js.Wrapper) *Event {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Event{}
	ret.Value_JS = input
	return ret
}

const (
	NONE            int = 0
	CAPTURING_PHASE int = 1
	AT_TARGET       int = 2
	BUBBLING_PHASE  int = 3
)

func NewEvent(_type string, eventInitDict *EventInit) (_result *Event) {
	_klass := js.Global().Get("Event")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Event // javascript: Event _what_return_name
	)
	_converted = EventFromJS(_returned)
	_result = _converted
	return
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
func (_this *Event) Type() string {
	var ret string
	value := _this.Value_JS.Get("type")
	ret = (value).String()
	return ret
}

// Target returning attribute 'target' with
// type EventTarget (idl: EventTarget).
func (_this *Event) Target() *EventTarget {
	var ret *EventTarget
	value := _this.Value_JS.Get("target")
	if value.Type() != js.TypeNull {
		ret = EventTargetFromJS(value)
	}
	return ret
}

// SrcElement returning attribute 'srcElement' with
// type EventTarget (idl: EventTarget).
func (_this *Event) SrcElement() *EventTarget {
	var ret *EventTarget
	value := _this.Value_JS.Get("srcElement")
	if value.Type() != js.TypeNull {
		ret = EventTargetFromJS(value)
	}
	return ret
}

// CurrentTarget returning attribute 'currentTarget' with
// type EventTarget (idl: EventTarget).
func (_this *Event) CurrentTarget() *EventTarget {
	var ret *EventTarget
	value := _this.Value_JS.Get("currentTarget")
	if value.Type() != js.TypeNull {
		ret = EventTargetFromJS(value)
	}
	return ret
}

// EventPhase returning attribute 'eventPhase' with
// type int (idl: unsigned short).
func (_this *Event) EventPhase() int {
	var ret int
	value := _this.Value_JS.Get("eventPhase")
	ret = (value).Int()
	return ret
}

// CancelBubble returning attribute 'cancelBubble' with
// type bool (idl: boolean).
func (_this *Event) CancelBubble() bool {
	var ret bool
	value := _this.Value_JS.Get("cancelBubble")
	ret = (value).Bool()
	return ret
}

// SetCancelBubble setting attribute 'cancelBubble' with
// type bool (idl: boolean).
func (_this *Event) SetCancelBubble(value bool) {
	input := value
	_this.Value_JS.Set("cancelBubble", input)
}

// Bubbles returning attribute 'bubbles' with
// type bool (idl: boolean).
func (_this *Event) Bubbles() bool {
	var ret bool
	value := _this.Value_JS.Get("bubbles")
	ret = (value).Bool()
	return ret
}

// Cancelable returning attribute 'cancelable' with
// type bool (idl: boolean).
func (_this *Event) Cancelable() bool {
	var ret bool
	value := _this.Value_JS.Get("cancelable")
	ret = (value).Bool()
	return ret
}

// ReturnValue returning attribute 'returnValue' with
// type bool (idl: boolean).
func (_this *Event) ReturnValue() bool {
	var ret bool
	value := _this.Value_JS.Get("returnValue")
	ret = (value).Bool()
	return ret
}

// SetReturnValue setting attribute 'returnValue' with
// type bool (idl: boolean).
func (_this *Event) SetReturnValue(value bool) {
	input := value
	_this.Value_JS.Set("returnValue", input)
}

// DefaultPrevented returning attribute 'defaultPrevented' with
// type bool (idl: boolean).
func (_this *Event) DefaultPrevented() bool {
	var ret bool
	value := _this.Value_JS.Get("defaultPrevented")
	ret = (value).Bool()
	return ret
}

// Composed returning attribute 'composed' with
// type bool (idl: boolean).
func (_this *Event) Composed() bool {
	var ret bool
	value := _this.Value_JS.Get("composed")
	ret = (value).Bool()
	return ret
}

// IsTrusted returning attribute 'isTrusted' with
// type bool (idl: boolean).
func (_this *Event) IsTrusted() bool {
	var ret bool
	value := _this.Value_JS.Get("isTrusted")
	ret = (value).Bool()
	return ret
}

// TimeStamp returning attribute 'timeStamp' with
// type float64 (idl: double).
func (_this *Event) TimeStamp() float64 {
	var ret float64
	value := _this.Value_JS.Get("timeStamp")
	ret = (value).Float()
	return ret
}

func (_this *Event) ComposedPath() (_result []*EventTarget) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("composedPath", _args[0:_end]...)
	var (
		_converted []*EventTarget // javascript: sequence<EventTarget> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*EventTarget, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *EventTarget
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = EventTargetFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *Event) StopPropagation() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("stopPropagation", _args[0:_end]...)
	return
}

func (_this *Event) StopImmediatePropagation() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("stopImmediatePropagation", _args[0:_end]...)
	return
}

func (_this *Event) PreventDefault() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("preventDefault", _args[0:_end]...)
	return
}

func (_this *Event) InitEvent(_type string, bubbles *bool, cancelable *bool) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if bubbles != nil {
		_p1 := bubbles
		_args[1] = _p1
		_end++
	}
	if cancelable != nil {
		_p2 := cancelable
		_args[2] = _p2
		_end++
	}
	_this.Value_JS.Call("initEvent", _args[0:_end]...)
	return
}

// EventListener is a callback interface.
type EventListener interface {
	HandleEvent(event *Event)
}

// EventListenerValue is javascript reference value for callback interface EventListener.
// This is holding the underlaying javascript object.
type EventListenerValue struct {
	// Value is the underlying javascript object or function.
	Value js.Value
	// Functions is the underlying function objects that is allocated for the interface callback
	Functions [1]js.Func
	// Go interface to invoke
	impl      EventListener
	function  func(event *Event)
	useInvoke bool
}

// JSValue is returning the javascript object that implements this callback interface
func (t *EventListenerValue) JSValue() js.Value {
	return t.Value
}

// Release is releasing all resources that is allocated.
func (t *EventListenerValue) Release() {
	for i := range t.Functions {
		if t.Functions[i].Type() != js.TypeUndefined {
			t.Functions[i].Release()
		}
	}
}

// NewEventListener is allocating a new javascript object that
// implements EventListener.
func NewEventListener(callback EventListener) *EventListenerValue {
	ret := &EventListenerValue{impl: callback}
	ret.Value = js.Global().Get("Object").New()
	ret.Functions[0] = ret.allocateHandleEvent()
	ret.Value.Set("handleEvent", ret.Functions[0])
	return ret
}

// NewEventListenerFunc is allocating a new javascript
// function is implements
// EventListener interface.
func NewEventListenerFunc(f func(event *Event)) *EventListenerValue {
	// single function will result in javascript function type, not an object
	ret := &EventListenerValue{function: f}
	ret.Functions[0] = ret.allocateHandleEvent()
	ret.Value = ret.Functions[0].Value
	return ret
}

// EventListenerFromJS is taking an javascript object that reference to a
// callback interface and return a corresponding interface that can be used
// to invoke on that element.
func EventListenerFromJS(value js.Wrapper) *EventListenerValue {
	input := value.JSValue()
	if input.Type() == js.TypeObject {
		return &EventListenerValue{Value: input}
	}
	if input.Type() == js.TypeFunction {
		return &EventListenerValue{Value: input, useInvoke: true}
	}
	panic("unsupported type")
}

func (t *EventListenerValue) allocateHandleEvent() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Event // javascript: Event event
		)
		_p0 = EventFromJS(args[0])
		if t.function != nil {
			t.function(_p0)
		} else {
			t.impl.HandleEvent(_p0)
		}
		// returning no return value
		return nil
	})
}

func (_this *EventListenerValue) HandleEvent(event *Event) {
	if _this.function != nil {
		_this.function(event)
	}
	if _this.impl != nil {
		_this.impl.HandleEvent(event)
	}
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := event.JSValue()
	_args[0] = _p0
	_end++
	if _this.useInvoke {
		// invoke a javascript function
		_this.Value.Invoke(_args[0:_end]...)
	} else {
		_this.Value.Call("handleEvent", _args[0:_end]...)
	}
	return
}

// interface: EventTarget
type EventTarget struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *EventTarget) JSValue() js.Value {
	return _this.Value_JS
}

// EventTargetFromJS is casting a js.Wrapper into EventTarget.
func EventTargetFromJS(value js.Wrapper) *EventTarget {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &EventTarget{}
	ret.Value_JS = input
	return ret
}

func NewEventTarget() (_result *EventTarget) {
	_klass := js.Global().Get("EventTarget")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *EventTarget // javascript: EventTarget _what_return_name
	)
	_converted = EventTargetFromJS(_returned)
	_result = _converted
	return
}

func (_this *EventTarget) AddEventListener(_type string, callback *EventListenerValue, options *Union) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	_p1 := callback.JSValue()
	_args[1] = _p1
	_end++
	if options != nil {
		_p2 := options.JSValue()
		_args[2] = _p2
		_end++
	}
	_this.Value_JS.Call("addEventListener", _args[0:_end]...)
	return
}

func (_this *EventTarget) RemoveEventListener(_type string, callback *EventListenerValue, options *Union) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	_p1 := callback.JSValue()
	_args[1] = _p1
	_end++
	if options != nil {
		_p2 := options.JSValue()
		_args[2] = _p2
		_end++
	}
	_this.Value_JS.Call("removeEventListener", _args[0:_end]...)
	return
}

func (_this *EventTarget) DispatchEvent(event *Event) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := event.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("dispatchEvent", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

// interface: ExtendableEvent
type ExtendableEvent struct {
	Event
}

// ExtendableEventFromJS is casting a js.Wrapper into ExtendableEvent.
func ExtendableEventFromJS(value js.Wrapper) *ExtendableEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ExtendableEvent{}
	ret.Value_JS = input
	return ret
}

func NewExtendableEvent(_type string, eventInitDict *ExtendableEventInit) (_result *ExtendableEvent) {
	_klass := js.Global().Get("ExtendableEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *ExtendableEvent // javascript: ExtendableEvent _what_return_name
	)
	_converted = ExtendableEventFromJS(_returned)
	_result = _converted
	return
}

func (_this *ExtendableEvent) WaitUntil(f *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := f.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("waitUntil", _args[0:_end]...)
	return
}
