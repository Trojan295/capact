// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"fmt"
	"io"
	"strconv"
)

type CreateTypeInstanceInput struct {
	TypeRef string               `json:"typeRef"`
	Tags    []*TagReferenceInput `json:"tags"`
	Value   interface{}          `json:"value"`
}

type TagFilterInput struct {
	Path string      `json:"path"`
	Rule *FilterRule `json:"rule"`
	// If not provided, latest revision for a given Tag is used
	Revision *string `json:"revision"`
}

type TagReference struct {
	Path     string `json:"path"`
	Revision string `json:"revision"`
}

type TagReferenceInput struct {
	Path string `json:"path"`
	// If not provided, latest revision for a given Tag is used
	Revision *string `json:"revision"`
}

type TypeInstance struct {
	Metadata        *TypeInstanceMetadata `json:"metadata"`
	ResourceVersion int                   `json:"resourceVersion"`
	Spec            *TypeInstanceSpec     `json:"spec"`
}

type TypeInstanceFilter struct {
	Tag     []*TagFilterInput   `json:"tag"`
	TypeRef *TypeRefFilterInput `json:"typeRef"`
}

type TypeInstanceInstrumentation struct {
	Metrics *TypeInstanceInstrumentationMetrics `json:"metrics"`
	Health  *TypeInstanceInstrumentationHealth  `json:"health"`
}

type TypeInstanceInstrumentationHealth struct {
	URL    *string                                  `json:"url"`
	Method *HTTPRequestMethod                       `json:"method"`
	Status *TypeInstanceInstrumentationHealthStatus `json:"status"`
}

type TypeInstanceInstrumentationMetrics struct {
	Endpoint   *string                                        `json:"endpoint"`
	Regex      *string                                        `json:"regex"`
	Dashboards []*TypeInstanceInstrumentationMetricsDashboard `json:"dashboards"`
}

type TypeInstanceInstrumentationMetricsDashboard struct {
	URL string `json:"url"`
}

type TypeInstanceMetadata struct {
	ID   string          `json:"id"`
	Tags []*TagReference `json:"tags"`
}

type TypeInstanceSpec struct {
	TypeRef         *TypeReference               `json:"typeRef"`
	Value           interface{}                  `json:"value"`
	Instrumentation *TypeInstanceInstrumentation `json:"instrumentation"`
}

type TypeRefFilterInput struct {
	Path string `json:"path"`
	// If not provided, latest revision for a given Type is used
	Revision *string `json:"revision"`
}

type TypeReference struct {
	Path     string `json:"path"`
	Revision string `json:"revision"`
}

type UpdateTypeInstanceInput struct {
	TypeRef         string               `json:"typeRef"`
	Tags            []*TagReferenceInput `json:"tags"`
	Value           interface{}          `json:"value"`
	ResourceVersion int                  `json:"resourceVersion"`
}

type FilterRule string

const (
	FilterRuleInclude FilterRule = "INCLUDE"
	FilterRuleExclude FilterRule = "EXCLUDE"
)

var AllFilterRule = []FilterRule{
	FilterRuleInclude,
	FilterRuleExclude,
}

func (e FilterRule) IsValid() bool {
	switch e {
	case FilterRuleInclude, FilterRuleExclude:
		return true
	}
	return false
}

func (e FilterRule) String() string {
	return string(e)
}

func (e *FilterRule) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FilterRule(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FilterRule", str)
	}
	return nil
}

func (e FilterRule) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type HTTPRequestMethod string

const (
	HTTPRequestMethodGet  HTTPRequestMethod = "GET"
	HTTPRequestMethodPost HTTPRequestMethod = "POST"
)

var AllHTTPRequestMethod = []HTTPRequestMethod{
	HTTPRequestMethodGet,
	HTTPRequestMethodPost,
}

func (e HTTPRequestMethod) IsValid() bool {
	switch e {
	case HTTPRequestMethodGet, HTTPRequestMethodPost:
		return true
	}
	return false
}

func (e HTTPRequestMethod) String() string {
	return string(e)
}

func (e *HTTPRequestMethod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HTTPRequestMethod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HTTPRequestMethod", str)
	}
	return nil
}

func (e HTTPRequestMethod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TypeInstanceInstrumentationHealthStatus string

const (
	TypeInstanceInstrumentationHealthStatusUnknown TypeInstanceInstrumentationHealthStatus = "UNKNOWN"
	TypeInstanceInstrumentationHealthStatusReady   TypeInstanceInstrumentationHealthStatus = "READY"
	TypeInstanceInstrumentationHealthStatusFailing TypeInstanceInstrumentationHealthStatus = "FAILING"
)

var AllTypeInstanceInstrumentationHealthStatus = []TypeInstanceInstrumentationHealthStatus{
	TypeInstanceInstrumentationHealthStatusUnknown,
	TypeInstanceInstrumentationHealthStatusReady,
	TypeInstanceInstrumentationHealthStatusFailing,
}

func (e TypeInstanceInstrumentationHealthStatus) IsValid() bool {
	switch e {
	case TypeInstanceInstrumentationHealthStatusUnknown, TypeInstanceInstrumentationHealthStatusReady, TypeInstanceInstrumentationHealthStatusFailing:
		return true
	}
	return false
}

func (e TypeInstanceInstrumentationHealthStatus) String() string {
	return string(e)
}

func (e *TypeInstanceInstrumentationHealthStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TypeInstanceInstrumentationHealthStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TypeInstanceInstrumentationHealthStatus", str)
	}
	return nil
}

func (e TypeInstanceInstrumentationHealthStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}