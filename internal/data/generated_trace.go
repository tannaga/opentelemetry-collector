// Copyright 2020 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "internal/data_generator/main.go". DO NOT EDIT.
// To regenerate this file run "go run internal/data_generator/main.go".

package data

import (
	otlpcommon "github.com/open-telemetry/opentelemetry-proto/gen/go/common/v1"
	otlpresource "github.com/open-telemetry/opentelemetry-proto/gen/go/resource/v1"
	otlptrace "github.com/open-telemetry/opentelemetry-proto/gen/go/trace/v1"
)

// ResourceSpansSlice logically represents a slice of ResourceSpans.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceSpansSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceSpansSlice struct {
	orig *[]*otlptrace.ResourceSpans
}

// NewResourceSpansSlice creates a ResourceSpansSlice with "len" empty elements.
//
// es := NewResourceSpansSlice(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     // Here should set all the values for e.
// }
func NewResourceSpansSlice(len int) ResourceSpansSlice {
	if len == 0 {
		orig := []*otlptrace.ResourceSpans(nil)
		return ResourceSpansSlice{&orig}
	}
	// Slice for underlying orig.
	origs := make([]otlptrace.ResourceSpans, len)
	// Slice for wrappers.
	wrappers := make([]*otlptrace.ResourceSpans, len)
	for i := range origs {
		wrappers[i] = &origs[i]
	}
	return ResourceSpansSlice{&wrappers}
}

func newResourceSpansSlice(orig *[]*otlptrace.ResourceSpans) ResourceSpansSlice {
	return ResourceSpansSlice{orig}
}

// Len returns the number of elements in the slice.
func (es ResourceSpansSlice) Len() int {
	return len(*es.orig)
}

// Get the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     ... // Do something with the element
// }
func (es ResourceSpansSlice) Get(ix int) ResourceSpans {
	return newResourceSpans((*es.orig)[ix])
}

// Remove the element at the given index from the slice.
// Elements after the removed one are shifted to fill the emptied space.
// The length of the slice is reduced by one.
func (es ResourceSpansSlice) Remove(ix int) {
	(*es.orig)[ix] = (*es.orig)[len(*es.orig)-1]
	(*es.orig)[len(*es.orig)-1] = nil
	*es.orig = (*es.orig)[:len(*es.orig)-1]
}

// Resize the slice. This operation is equivalent with slice[from:to].
func (es ResourceSpansSlice) Resize(from, to int) {
	*es.orig = (*es.orig)[from:to]
}

// InstrumentationLibrarySpans is a collection of spans from a LibraryInstrumentation.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceSpans function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceSpans struct {
	// Wrap OTLP otlptrace.ResourceSpans.
	orig *otlptrace.ResourceSpans
}

// NewResourceSpans creates a new empty ResourceSpans.
func NewResourceSpans() ResourceSpans {
	return ResourceSpans{&otlptrace.ResourceSpans{}}
}

func newResourceSpans(orig *otlptrace.ResourceSpans) ResourceSpans {
	return ResourceSpans{orig}
}

// Resource returns the resource associated with this ResourceSpans.
func (ms ResourceSpans) Resource() Resource {
	if ms.orig.Resource == nil {
		// No Resource available, initialize one to make all operations on Resource available.
		ms.orig.Resource = &otlpresource.Resource{}
	}
	return newResource(ms.orig.Resource)
}

// SetResource replaces the resource associated with this ResourceSpans.
func (ms ResourceSpans) SetResource(v Resource) {
	ms.orig.Resource = v.orig
}

// InstrumentationLibrarySpans returns the InstrumentationLibrarySpans associated with this ResourceSpans.
func (ms ResourceSpans) InstrumentationLibrarySpans() InstrumentationLibrarySpansSlice {
	return newInstrumentationLibrarySpansSlice(&ms.orig.InstrumentationLibrarySpans)
}

// SetInstrumentationLibrarySpans replaces the InstrumentationLibrarySpans associated with this ResourceSpans.
func (ms ResourceSpans) SetInstrumentationLibrarySpans(v InstrumentationLibrarySpansSlice) {
	ms.orig.InstrumentationLibrarySpans = *v.orig
}

// InstrumentationLibrarySpansSlice logically represents a slice of InstrumentationLibrarySpans.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewInstrumentationLibrarySpansSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type InstrumentationLibrarySpansSlice struct {
	orig *[]*otlptrace.InstrumentationLibrarySpans
}

// NewInstrumentationLibrarySpansSlice creates a InstrumentationLibrarySpansSlice with "len" empty elements.
//
// es := NewInstrumentationLibrarySpansSlice(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     // Here should set all the values for e.
// }
func NewInstrumentationLibrarySpansSlice(len int) InstrumentationLibrarySpansSlice {
	if len == 0 {
		orig := []*otlptrace.InstrumentationLibrarySpans(nil)
		return InstrumentationLibrarySpansSlice{&orig}
	}
	// Slice for underlying orig.
	origs := make([]otlptrace.InstrumentationLibrarySpans, len)
	// Slice for wrappers.
	wrappers := make([]*otlptrace.InstrumentationLibrarySpans, len)
	for i := range origs {
		wrappers[i] = &origs[i]
	}
	return InstrumentationLibrarySpansSlice{&wrappers}
}

func newInstrumentationLibrarySpansSlice(orig *[]*otlptrace.InstrumentationLibrarySpans) InstrumentationLibrarySpansSlice {
	return InstrumentationLibrarySpansSlice{orig}
}

// Len returns the number of elements in the slice.
func (es InstrumentationLibrarySpansSlice) Len() int {
	return len(*es.orig)
}

// Get the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     ... // Do something with the element
// }
func (es InstrumentationLibrarySpansSlice) Get(ix int) InstrumentationLibrarySpans {
	return newInstrumentationLibrarySpans((*es.orig)[ix])
}

// Remove the element at the given index from the slice.
// Elements after the removed one are shifted to fill the emptied space.
// The length of the slice is reduced by one.
func (es InstrumentationLibrarySpansSlice) Remove(ix int) {
	(*es.orig)[ix] = (*es.orig)[len(*es.orig)-1]
	(*es.orig)[len(*es.orig)-1] = nil
	*es.orig = (*es.orig)[:len(*es.orig)-1]
}

// Resize the slice. This operation is equivalent with slice[from:to].
func (es InstrumentationLibrarySpansSlice) Resize(from, to int) {
	*es.orig = (*es.orig)[from:to]
}

// InstrumentationLibrarySpans is a collection of spans from a LibraryInstrumentation.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewInstrumentationLibrarySpans function to create new instances.
// Important: zero-initialized instance is not valid for use.
type InstrumentationLibrarySpans struct {
	// Wrap OTLP otlptrace.InstrumentationLibrarySpans.
	orig *otlptrace.InstrumentationLibrarySpans
}

// NewInstrumentationLibrarySpans creates a new empty InstrumentationLibrarySpans.
func NewInstrumentationLibrarySpans() InstrumentationLibrarySpans {
	return InstrumentationLibrarySpans{&otlptrace.InstrumentationLibrarySpans{}}
}

func newInstrumentationLibrarySpans(orig *otlptrace.InstrumentationLibrarySpans) InstrumentationLibrarySpans {
	return InstrumentationLibrarySpans{orig}
}

// InstrumentationLibrary returns the instrumentationlibrary associated with this InstrumentationLibrarySpans.
func (ms InstrumentationLibrarySpans) InstrumentationLibrary() InstrumentationLibrary {
	if ms.orig.InstrumentationLibrary == nil {
		// No InstrumentationLibrary available, initialize one to make all operations on InstrumentationLibrary available.
		ms.orig.InstrumentationLibrary = &otlpcommon.InstrumentationLibrary{}
	}
	return newInstrumentationLibrary(ms.orig.InstrumentationLibrary)
}

// SetInstrumentationLibrary replaces the instrumentationlibrary associated with this InstrumentationLibrarySpans.
func (ms InstrumentationLibrarySpans) SetInstrumentationLibrary(v InstrumentationLibrary) {
	ms.orig.InstrumentationLibrary = v.orig
}

// Spans returns the Spans associated with this InstrumentationLibrarySpans.
func (ms InstrumentationLibrarySpans) Spans() SpanSlice {
	return newSpanSlice(&ms.orig.Spans)
}

// SetSpans replaces the Spans associated with this InstrumentationLibrarySpans.
func (ms InstrumentationLibrarySpans) SetSpans(v SpanSlice) {
	ms.orig.Spans = *v.orig
}

// SpanSlice logically represents a slice of Span.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanSlice struct {
	orig *[]*otlptrace.Span
}

// NewSpanSlice creates a SpanSlice with "len" empty elements.
//
// es := NewSpanSlice(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     // Here should set all the values for e.
// }
func NewSpanSlice(len int) SpanSlice {
	if len == 0 {
		orig := []*otlptrace.Span(nil)
		return SpanSlice{&orig}
	}
	// Slice for underlying orig.
	origs := make([]otlptrace.Span, len)
	// Slice for wrappers.
	wrappers := make([]*otlptrace.Span, len)
	for i := range origs {
		wrappers[i] = &origs[i]
	}
	return SpanSlice{&wrappers}
}

func newSpanSlice(orig *[]*otlptrace.Span) SpanSlice {
	return SpanSlice{orig}
}

// Len returns the number of elements in the slice.
func (es SpanSlice) Len() int {
	return len(*es.orig)
}

// Get the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     ... // Do something with the element
// }
func (es SpanSlice) Get(ix int) Span {
	return newSpan((*es.orig)[ix])
}

// Remove the element at the given index from the slice.
// Elements after the removed one are shifted to fill the emptied space.
// The length of the slice is reduced by one.
func (es SpanSlice) Remove(ix int) {
	(*es.orig)[ix] = (*es.orig)[len(*es.orig)-1]
	(*es.orig)[len(*es.orig)-1] = nil
	*es.orig = (*es.orig)[:len(*es.orig)-1]
}

// Resize the slice. This operation is equivalent with slice[from:to].
func (es SpanSlice) Resize(from, to int) {
	*es.orig = (*es.orig)[from:to]
}

// Span represents a single operation within a trace.
// See Span definition in OTLP: https://github.com/open-telemetry/opentelemetry-proto/blob/master/opentelemetry/proto/trace/v1/trace.proto#L37
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpan function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Span struct {
	// Wrap OTLP otlptrace.Span.
	orig *otlptrace.Span
}

// NewSpan creates a new empty Span.
func NewSpan() Span {
	return Span{&otlptrace.Span{}}
}

func newSpan(orig *otlptrace.Span) Span {
	return Span{orig}
}

// TraceID returns the traceid associated with this Span.
func (ms Span) TraceID() TraceID {
	return TraceID(ms.orig.TraceId)
}

// SetTraceID replaces the traceid associated with this Span.
func (ms Span) SetTraceID(v TraceID) {
	ms.orig.TraceId = []byte(v)
}

// SpanID returns the spanid associated with this Span.
func (ms Span) SpanID() SpanID {
	return SpanID(ms.orig.SpanId)
}

// SetSpanID replaces the spanid associated with this Span.
func (ms Span) SetSpanID(v SpanID) {
	ms.orig.SpanId = []byte(v)
}

// TraceState returns the tracestate associated with this Span.
func (ms Span) TraceState() TraceState {
	return TraceState(ms.orig.TraceState)
}

// SetTraceState replaces the tracestate associated with this Span.
func (ms Span) SetTraceState(v TraceState) {
	ms.orig.TraceState = string(v)
}

// ParentSpanID returns the parentspanid associated with this Span.
func (ms Span) ParentSpanID() SpanID {
	return SpanID(ms.orig.ParentSpanId)
}

// SetParentSpanID replaces the parentspanid associated with this Span.
func (ms Span) SetParentSpanID(v SpanID) {
	ms.orig.ParentSpanId = []byte(v)
}

// Name returns the name associated with this Span.
func (ms Span) Name() string {
	return ms.orig.Name
}

// SetName replaces the name associated with this Span.
func (ms Span) SetName(v string) {
	ms.orig.Name = v
}

// Kind returns the kind associated with this Span.
func (ms Span) Kind() SpanKind {
	return SpanKind(ms.orig.Kind)
}

// SetKind replaces the kind associated with this Span.
func (ms Span) SetKind(v SpanKind) {
	ms.orig.Kind = otlptrace.Span_SpanKind(v)
}

// StartTime returns the starttime associated with this Span.
func (ms Span) StartTime() TimestampUnixNano {
	return TimestampUnixNano(ms.orig.StartTimeUnixNano)
}

// SetStartTime replaces the starttime associated with this Span.
func (ms Span) SetStartTime(v TimestampUnixNano) {
	ms.orig.StartTimeUnixNano = uint64(v)
}

// EndTime returns the endtime associated with this Span.
func (ms Span) EndTime() TimestampUnixNano {
	return TimestampUnixNano(ms.orig.EndTimeUnixNano)
}

// SetEndTime replaces the endtime associated with this Span.
func (ms Span) SetEndTime(v TimestampUnixNano) {
	ms.orig.EndTimeUnixNano = uint64(v)
}

// Attributes returns the Attributes associated with this Span.
func (ms Span) Attributes() AttributeMap {
	return newAttributeMap(&ms.orig.Attributes)
}

// SetAttributes replaces the Attributes associated with this Span.
func (ms Span) SetAttributes(v AttributeMap) {
	ms.orig.Attributes = *v.orig
}

// DroppedAttributesCount returns the droppedattributescount associated with this Span.
func (ms Span) DroppedAttributesCount() uint32 {
	return ms.orig.DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this Span.
func (ms Span) SetDroppedAttributesCount(v uint32) {
	ms.orig.DroppedAttributesCount = v
}

// Events returns the Events associated with this Span.
func (ms Span) Events() SpanEventSlice {
	return newSpanEventSlice(&ms.orig.Events)
}

// SetEvents replaces the Events associated with this Span.
func (ms Span) SetEvents(v SpanEventSlice) {
	ms.orig.Events = *v.orig
}

// DroppedEventsCount returns the droppedeventscount associated with this Span.
func (ms Span) DroppedEventsCount() uint32 {
	return ms.orig.DroppedEventsCount
}

// SetDroppedEventsCount replaces the droppedeventscount associated with this Span.
func (ms Span) SetDroppedEventsCount(v uint32) {
	ms.orig.DroppedEventsCount = v
}

// Links returns the Links associated with this Span.
func (ms Span) Links() SpanLinkSlice {
	return newSpanLinkSlice(&ms.orig.Links)
}

// SetLinks replaces the Links associated with this Span.
func (ms Span) SetLinks(v SpanLinkSlice) {
	ms.orig.Links = *v.orig
}

// DroppedLinksCount returns the droppedlinkscount associated with this Span.
func (ms Span) DroppedLinksCount() uint32 {
	return ms.orig.DroppedLinksCount
}

// SetDroppedLinksCount replaces the droppedlinkscount associated with this Span.
func (ms Span) SetDroppedLinksCount(v uint32) {
	ms.orig.DroppedLinksCount = v
}

// Status returns the status associated with this Span.
func (ms Span) Status() SpanStatus {
	if ms.orig.Status == nil {
		// No Status available, initialize one to make all operations on SpanStatus available.
		ms.orig.Status = &otlptrace.Status{}
	}
	return newSpanStatus(ms.orig.Status)
}

// SetStatus replaces the status associated with this Span.
func (ms Span) SetStatus(v SpanStatus) {
	ms.orig.Status = v.orig
}

// SpanEventSlice logically represents a slice of SpanEvent.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanEventSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanEventSlice struct {
	orig *[]*otlptrace.Span_Event
}

// NewSpanEventSlice creates a SpanEventSlice with "len" empty elements.
//
// es := NewSpanEventSlice(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     // Here should set all the values for e.
// }
func NewSpanEventSlice(len int) SpanEventSlice {
	if len == 0 {
		orig := []*otlptrace.Span_Event(nil)
		return SpanEventSlice{&orig}
	}
	// Slice for underlying orig.
	origs := make([]otlptrace.Span_Event, len)
	// Slice for wrappers.
	wrappers := make([]*otlptrace.Span_Event, len)
	for i := range origs {
		wrappers[i] = &origs[i]
	}
	return SpanEventSlice{&wrappers}
}

func newSpanEventSlice(orig *[]*otlptrace.Span_Event) SpanEventSlice {
	return SpanEventSlice{orig}
}

// Len returns the number of elements in the slice.
func (es SpanEventSlice) Len() int {
	return len(*es.orig)
}

// Get the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     ... // Do something with the element
// }
func (es SpanEventSlice) Get(ix int) SpanEvent {
	return newSpanEvent((*es.orig)[ix])
}

// Remove the element at the given index from the slice.
// Elements after the removed one are shifted to fill the emptied space.
// The length of the slice is reduced by one.
func (es SpanEventSlice) Remove(ix int) {
	(*es.orig)[ix] = (*es.orig)[len(*es.orig)-1]
	(*es.orig)[len(*es.orig)-1] = nil
	*es.orig = (*es.orig)[:len(*es.orig)-1]
}

// Resize the slice. This operation is equivalent with slice[from:to].
func (es SpanEventSlice) Resize(from, to int) {
	*es.orig = (*es.orig)[from:to]
}

// SpanEvent is a time-stamped annotation of the span, consisting of user-supplied
// text description and key-value pairs. See OTLP for event definition.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanEvent function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanEvent struct {
	// Wrap OTLP otlptrace.Span_Event.
	orig *otlptrace.Span_Event
}

// NewSpanEvent creates a new empty SpanEvent.
func NewSpanEvent() SpanEvent {
	return SpanEvent{&otlptrace.Span_Event{}}
}

func newSpanEvent(orig *otlptrace.Span_Event) SpanEvent {
	return SpanEvent{orig}
}

// Timestamp returns the timestamp associated with this SpanEvent.
func (ms SpanEvent) Timestamp() TimestampUnixNano {
	return TimestampUnixNano(ms.orig.TimeUnixNano)
}

// SetTimestamp replaces the timestamp associated with this SpanEvent.
func (ms SpanEvent) SetTimestamp(v TimestampUnixNano) {
	ms.orig.TimeUnixNano = uint64(v)
}

// Name returns the name associated with this SpanEvent.
func (ms SpanEvent) Name() string {
	return ms.orig.Name
}

// SetName replaces the name associated with this SpanEvent.
func (ms SpanEvent) SetName(v string) {
	ms.orig.Name = v
}

// Attributes returns the Attributes associated with this SpanEvent.
func (ms SpanEvent) Attributes() AttributeMap {
	return newAttributeMap(&ms.orig.Attributes)
}

// SetAttributes replaces the Attributes associated with this SpanEvent.
func (ms SpanEvent) SetAttributes(v AttributeMap) {
	ms.orig.Attributes = *v.orig
}

// DroppedAttributesCount returns the droppedattributescount associated with this SpanEvent.
func (ms SpanEvent) DroppedAttributesCount() uint32 {
	return ms.orig.DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this SpanEvent.
func (ms SpanEvent) SetDroppedAttributesCount(v uint32) {
	ms.orig.DroppedAttributesCount = v
}

// SpanLinkSlice logically represents a slice of SpanLink.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanLinkSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanLinkSlice struct {
	orig *[]*otlptrace.Span_Link
}

// NewSpanLinkSlice creates a SpanLinkSlice with "len" empty elements.
//
// es := NewSpanLinkSlice(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     // Here should set all the values for e.
// }
func NewSpanLinkSlice(len int) SpanLinkSlice {
	if len == 0 {
		orig := []*otlptrace.Span_Link(nil)
		return SpanLinkSlice{&orig}
	}
	// Slice for underlying orig.
	origs := make([]otlptrace.Span_Link, len)
	// Slice for wrappers.
	wrappers := make([]*otlptrace.Span_Link, len)
	for i := range origs {
		wrappers[i] = &origs[i]
	}
	return SpanLinkSlice{&wrappers}
}

func newSpanLinkSlice(orig *[]*otlptrace.Span_Link) SpanLinkSlice {
	return SpanLinkSlice{orig}
}

// Len returns the number of elements in the slice.
func (es SpanLinkSlice) Len() int {
	return len(*es.orig)
}

// Get the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     ... // Do something with the element
// }
func (es SpanLinkSlice) Get(ix int) SpanLink {
	return newSpanLink((*es.orig)[ix])
}

// Remove the element at the given index from the slice.
// Elements after the removed one are shifted to fill the emptied space.
// The length of the slice is reduced by one.
func (es SpanLinkSlice) Remove(ix int) {
	(*es.orig)[ix] = (*es.orig)[len(*es.orig)-1]
	(*es.orig)[len(*es.orig)-1] = nil
	*es.orig = (*es.orig)[:len(*es.orig)-1]
}

// Resize the slice. This operation is equivalent with slice[from:to].
func (es SpanLinkSlice) Resize(from, to int) {
	*es.orig = (*es.orig)[from:to]
}

// SpanLink is a pointer from the current span to another span in the same trace or in a
// different trace. See OTLP for link definition.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanLink function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanLink struct {
	// Wrap OTLP otlptrace.Span_Link.
	orig *otlptrace.Span_Link
}

// NewSpanLink creates a new empty SpanLink.
func NewSpanLink() SpanLink {
	return SpanLink{&otlptrace.Span_Link{}}
}

func newSpanLink(orig *otlptrace.Span_Link) SpanLink {
	return SpanLink{orig}
}

// TraceID returns the traceid associated with this SpanLink.
func (ms SpanLink) TraceID() TraceID {
	return TraceID(ms.orig.TraceId)
}

// SetTraceID replaces the traceid associated with this SpanLink.
func (ms SpanLink) SetTraceID(v TraceID) {
	ms.orig.TraceId = []byte(v)
}

// SpanID returns the spanid associated with this SpanLink.
func (ms SpanLink) SpanID() SpanID {
	return SpanID(ms.orig.SpanId)
}

// SetSpanID replaces the spanid associated with this SpanLink.
func (ms SpanLink) SetSpanID(v SpanID) {
	ms.orig.SpanId = []byte(v)
}

// TraceState returns the tracestate associated with this SpanLink.
func (ms SpanLink) TraceState() TraceState {
	return TraceState(ms.orig.TraceState)
}

// SetTraceState replaces the tracestate associated with this SpanLink.
func (ms SpanLink) SetTraceState(v TraceState) {
	ms.orig.TraceState = string(v)
}

// Attributes returns the Attributes associated with this SpanLink.
func (ms SpanLink) Attributes() AttributeMap {
	return newAttributeMap(&ms.orig.Attributes)
}

// SetAttributes replaces the Attributes associated with this SpanLink.
func (ms SpanLink) SetAttributes(v AttributeMap) {
	ms.orig.Attributes = *v.orig
}

// DroppedAttributesCount returns the droppedattributescount associated with this SpanLink.
func (ms SpanLink) DroppedAttributesCount() uint32 {
	return ms.orig.DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this SpanLink.
func (ms SpanLink) SetDroppedAttributesCount(v uint32) {
	ms.orig.DroppedAttributesCount = v
}

// SpanStatus is an optional final status for this span. Semantically when Status wasn't set
// it is means span ended without errors and assume Status.Ok (code = 0).
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanStatus function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanStatus struct {
	// Wrap OTLP otlptrace.Status.
	orig *otlptrace.Status
}

// NewSpanStatus creates a new empty SpanStatus.
func NewSpanStatus() SpanStatus {
	return SpanStatus{&otlptrace.Status{}}
}

func newSpanStatus(orig *otlptrace.Status) SpanStatus {
	return SpanStatus{orig}
}

// Code returns the code associated with this SpanStatus.
func (ms SpanStatus) Code() StatusCode {
	return StatusCode(ms.orig.Code)
}

// SetCode replaces the code associated with this SpanStatus.
func (ms SpanStatus) SetCode(v StatusCode) {
	ms.orig.Code = otlptrace.Status_StatusCode(v)
}

// Message returns the message associated with this SpanStatus.
func (ms SpanStatus) Message() string {
	return ms.orig.Message
}

// SetMessage replaces the message associated with this SpanStatus.
func (ms SpanStatus) SetMessage(v string) {
	ms.orig.Message = v
}