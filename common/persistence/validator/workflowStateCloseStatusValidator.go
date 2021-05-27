// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package validator

import (
	"fmt"

	"github.com/uber/cadence/common/persistence"

	"github.com/uber/cadence/common/types"
)

var (
	validWorkflowStates = map[int]struct{}{
		persistence.WorkflowStateCreated:   {},
		persistence.WorkflowStateRunning:   {},
		persistence.WorkflowStateCompleted: {},
		persistence.WorkflowStateZombie:    {},
		persistence.WorkflowStateCorrupted: {},
	}

	validWorkflowCloseStatuses = map[int]struct{}{
		persistence.WorkflowCloseStatusNone:           {},
		persistence.WorkflowCloseStatusCompleted:      {},
		persistence.WorkflowCloseStatusFailed:         {},
		persistence.WorkflowCloseStatusCanceled:       {},
		persistence.WorkflowCloseStatusTerminated:     {},
		persistence.WorkflowCloseStatusContinuedAsNew: {},
		persistence.WorkflowCloseStatusTimedOut:       {},
	}
)

// ValidateCreateWorkflowStateCloseStatus validate workflow state and close status
func ValidateCreateWorkflowStateCloseStatus(
	state int,
	closeStatus int,
) error {

	if err := validateWorkflowState(state); err != nil {
		return err
	}
	if err := validateWorkflowCloseStatus(closeStatus); err != nil {
		return err
	}

	// validate workflow state & close status
	if state == persistence.WorkflowStateCompleted || closeStatus != persistence.WorkflowCloseStatusNone {
		return &types.InternalServiceError{
			Message: fmt.Sprintf("Create workflow with invalid state: %v or close status: %v",
				state, closeStatus),
		}
	}
	return nil
}

// ValidateUpdateWorkflowStateCloseStatus validate workflow state and close status
func ValidateUpdateWorkflowStateCloseStatus(
	state int,
	closeStatus int,
) error {

	if err := validateWorkflowState(state); err != nil {
		return err
	}
	if err := validateWorkflowCloseStatus(closeStatus); err != nil {
		return err
	}

	// validate workflow state & close status
	if closeStatus == persistence.WorkflowCloseStatusNone {
		if state == persistence.WorkflowStateCompleted {
			return &types.InternalServiceError{
				Message: fmt.Sprintf("Update workflow with invalid state: %v or close status: %v",
					state, closeStatus),
			}
		}
	} else {
		// WorkflowCloseStatusCompleted
		// WorkflowCloseStatusFailed
		// WorkflowCloseStatusCanceled
		// WorkflowCloseStatusTerminated
		// WorkflowCloseStatusContinuedAsNew
		// WorkflowCloseStatusTimedOut
		if state != persistence.WorkflowStateCompleted {
			return &types.InternalServiceError{
				Message: fmt.Sprintf("Update workflow with invalid state: %v or close status: %v",
					state, closeStatus),
			}
		}
	}
	return nil
}

// validateWorkflowState validate workflow state
func validateWorkflowState(
	state int,
) error {

	if _, ok := validWorkflowStates[state]; !ok {
		return &types.InternalServiceError{
			Message: fmt.Sprintf("Invalid workflow state: %v", state),
		}
	}

	return nil
}

// validateWorkflowCloseStatus validate workflow close status
func validateWorkflowCloseStatus(
	closeStatus int,
) error {

	if _, ok := validWorkflowCloseStatuses[closeStatus]; !ok {
		return &types.InternalServiceError{
			Message: fmt.Sprintf("Invalid workflow close status: %v", closeStatus),
		}
	}

	return nil
}

// ToInternalWorkflowExecutionCloseStatus convert persistence representation of close status to internal representation
func ToInternalWorkflowExecutionCloseStatus(
	closeStatus int,
) *types.WorkflowExecutionCloseStatus {

	switch closeStatus {
	case persistence.WorkflowCloseStatusNone:
		return nil
	case persistence.WorkflowCloseStatusCompleted:
		return types.WorkflowExecutionCloseStatusCompleted.Ptr()
	case persistence.WorkflowCloseStatusFailed:
		return types.WorkflowExecutionCloseStatusFailed.Ptr()
	case persistence.WorkflowCloseStatusCanceled:
		return types.WorkflowExecutionCloseStatusCanceled.Ptr()
	case persistence.WorkflowCloseStatusTerminated:
		return types.WorkflowExecutionCloseStatusTerminated.Ptr()
	case persistence.WorkflowCloseStatusContinuedAsNew:
		return types.WorkflowExecutionCloseStatusContinuedAsNew.Ptr()
	case persistence.WorkflowCloseStatusTimedOut:
		return types.WorkflowExecutionCloseStatusTimedOut.Ptr()
	default:
		panic("Invalid value for enum WorkflowExecutionCloseStatus")
	}
}

// FromInternalWorkflowExecutionCloseStatus convert internal representation of close status to persistence representation
func FromInternalWorkflowExecutionCloseStatus(
	closeStatus *types.WorkflowExecutionCloseStatus,
) int {
	if closeStatus == nil {
		return persistence.WorkflowCloseStatusNone
	}
	switch *closeStatus {
	case types.WorkflowExecutionCloseStatusCompleted:
		return persistence.WorkflowCloseStatusCompleted
	case types.WorkflowExecutionCloseStatusFailed:
		return persistence.WorkflowCloseStatusFailed
	case types.WorkflowExecutionCloseStatusCanceled:
		return persistence.WorkflowCloseStatusCanceled
	case types.WorkflowExecutionCloseStatusTerminated:
		return persistence.WorkflowCloseStatusTerminated
	case types.WorkflowExecutionCloseStatusContinuedAsNew:
		return persistence.WorkflowCloseStatusContinuedAsNew
	case types.WorkflowExecutionCloseStatusTimedOut:
		return persistence.WorkflowCloseStatusTimedOut
	default:
		panic("Invalid value for enum WorkflowExecutionCloseStatus")
	}
}
