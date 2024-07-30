/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v2beta1

import (
	v2beta1 "github.com/kyverno/kyverno/api/kyverno/v2beta1"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// ConditionApplyConfiguration represents an declarative configuration of the Condition type for use
// with apply.
type ConditionApplyConfiguration struct {
	RawKey   *v1.JSON                   `json:"key,omitempty"`
	Operator *v2beta1.ConditionOperator `json:"operator,omitempty"`
	RawValue *v1.JSON                   `json:"value,omitempty"`
	Message  *string                    `json:"message,omitempty"`
}

// ConditionApplyConfiguration constructs an declarative configuration of the Condition type for use with
// apply.
func Condition() *ConditionApplyConfiguration {
	return &ConditionApplyConfiguration{}
}

// WithRawKey sets the RawKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RawKey field is set to the value of the last call.
func (b *ConditionApplyConfiguration) WithRawKey(value v1.JSON) *ConditionApplyConfiguration {
	b.RawKey = &value
	return b
}

// WithOperator sets the Operator field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Operator field is set to the value of the last call.
func (b *ConditionApplyConfiguration) WithOperator(value v2beta1.ConditionOperator) *ConditionApplyConfiguration {
	b.Operator = &value
	return b
}

// WithRawValue sets the RawValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RawValue field is set to the value of the last call.
func (b *ConditionApplyConfiguration) WithRawValue(value v1.JSON) *ConditionApplyConfiguration {
	b.RawValue = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *ConditionApplyConfiguration) WithMessage(value string) *ConditionApplyConfiguration {
	b.Message = &value
	return b
}
