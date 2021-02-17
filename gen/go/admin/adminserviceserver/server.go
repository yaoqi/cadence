// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by thriftrw-plugin-yarpc
// @generated

package adminserviceserver

import (
	context "context"

	wire "go.uber.org/thriftrw/wire"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"

	admin "github.com/uber/cadence/gen/go/admin"
	replicator "github.com/uber/cadence/gen/go/replicator"
	shared "github.com/uber/cadence/gen/go/shared"
)

// Interface is the server-side interface for the AdminService service.
type Interface interface {
	AddSearchAttribute(
		ctx context.Context,
		Request *admin.AddSearchAttributeRequest,
	) error

	CloseShard(
		ctx context.Context,
		Request *shared.CloseShardRequest,
	) error

	DescribeCluster(
		ctx context.Context,
	) (*admin.DescribeClusterResponse, error)

	DescribeHistoryHost(
		ctx context.Context,
		Request *shared.DescribeHistoryHostRequest,
	) (*shared.DescribeHistoryHostResponse, error)

	DescribeQueue(
		ctx context.Context,
		Request *shared.DescribeQueueRequest,
	) (*shared.DescribeQueueResponse, error)

	DescribeWorkflowExecution(
		ctx context.Context,
		Request *admin.DescribeWorkflowExecutionRequest,
	) (*admin.DescribeWorkflowExecutionResponse, error)

	GetDLQReplicationMessages(
		ctx context.Context,
		Request *replicator.GetDLQReplicationMessagesRequest,
	) (*replicator.GetDLQReplicationMessagesResponse, error)

	GetDomainReplicationMessages(
		ctx context.Context,
		Request *replicator.GetDomainReplicationMessagesRequest,
	) (*replicator.GetDomainReplicationMessagesResponse, error)

	GetReplicationMessages(
		ctx context.Context,
		Request *replicator.GetReplicationMessagesRequest,
	) (*replicator.GetReplicationMessagesResponse, error)

	GetWorkflowExecutionRawHistoryV2(
		ctx context.Context,
		GetRequest *admin.GetWorkflowExecutionRawHistoryV2Request,
	) (*admin.GetWorkflowExecutionRawHistoryV2Response, error)

	MergeDLQMessages(
		ctx context.Context,
		Request *replicator.MergeDLQMessagesRequest,
	) (*replicator.MergeDLQMessagesResponse, error)

	PurgeDLQMessages(
		ctx context.Context,
		Request *replicator.PurgeDLQMessagesRequest,
	) error

	ReadDLQMessages(
		ctx context.Context,
		Request *replicator.ReadDLQMessagesRequest,
	) (*replicator.ReadDLQMessagesResponse, error)

	ReapplyEvents(
		ctx context.Context,
		ReapplyEventsRequest *shared.ReapplyEventsRequest,
	) error

	RefreshWorkflowTasks(
		ctx context.Context,
		Request *shared.RefreshWorkflowTasksRequest,
	) error

	RemoveTask(
		ctx context.Context,
		Request *shared.RemoveTaskRequest,
	) error

	ResendReplicationTasks(
		ctx context.Context,
		Request *admin.ResendReplicationTasksRequest,
	) error

	ResetQueue(
		ctx context.Context,
		Request *shared.ResetQueueRequest,
	) error
}

// New prepares an implementation of the AdminService service for
// registration.
//
// 	handler := AdminServiceHandler{}
// 	dispatcher.Register(adminserviceserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "AdminService",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "AddSearchAttribute",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.AddSearchAttribute),
				},
				Signature:    "AddSearchAttribute(Request *admin.AddSearchAttributeRequest)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "CloseShard",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.CloseShard),
				},
				Signature:    "CloseShard(Request *shared.CloseShardRequest)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "DescribeCluster",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.DescribeCluster),
				},
				Signature:    "DescribeCluster() (*admin.DescribeClusterResponse)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "DescribeHistoryHost",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.DescribeHistoryHost),
				},
				Signature:    "DescribeHistoryHost(Request *shared.DescribeHistoryHostRequest) (*shared.DescribeHistoryHostResponse)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "DescribeQueue",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.DescribeQueue),
				},
				Signature:    "DescribeQueue(Request *shared.DescribeQueueRequest) (*shared.DescribeQueueResponse)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "DescribeWorkflowExecution",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.DescribeWorkflowExecution),
				},
				Signature:    "DescribeWorkflowExecution(Request *admin.DescribeWorkflowExecutionRequest) (*admin.DescribeWorkflowExecutionResponse)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "GetDLQReplicationMessages",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.GetDLQReplicationMessages),
				},
				Signature:    "GetDLQReplicationMessages(Request *replicator.GetDLQReplicationMessagesRequest) (*replicator.GetDLQReplicationMessagesResponse)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "GetDomainReplicationMessages",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.GetDomainReplicationMessages),
				},
				Signature:    "GetDomainReplicationMessages(Request *replicator.GetDomainReplicationMessagesRequest) (*replicator.GetDomainReplicationMessagesResponse)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "GetReplicationMessages",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.GetReplicationMessages),
				},
				Signature:    "GetReplicationMessages(Request *replicator.GetReplicationMessagesRequest) (*replicator.GetReplicationMessagesResponse)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "GetWorkflowExecutionRawHistoryV2",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.GetWorkflowExecutionRawHistoryV2),
				},
				Signature:    "GetWorkflowExecutionRawHistoryV2(GetRequest *admin.GetWorkflowExecutionRawHistoryV2Request) (*admin.GetWorkflowExecutionRawHistoryV2Response)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "MergeDLQMessages",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.MergeDLQMessages),
				},
				Signature:    "MergeDLQMessages(Request *replicator.MergeDLQMessagesRequest) (*replicator.MergeDLQMessagesResponse)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "PurgeDLQMessages",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.PurgeDLQMessages),
				},
				Signature:    "PurgeDLQMessages(Request *replicator.PurgeDLQMessagesRequest)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "ReadDLQMessages",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.ReadDLQMessages),
				},
				Signature:    "ReadDLQMessages(Request *replicator.ReadDLQMessagesRequest) (*replicator.ReadDLQMessagesResponse)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "ReapplyEvents",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.ReapplyEvents),
				},
				Signature:    "ReapplyEvents(ReapplyEventsRequest *shared.ReapplyEventsRequest)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "RefreshWorkflowTasks",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RefreshWorkflowTasks),
				},
				Signature:    "RefreshWorkflowTasks(Request *shared.RefreshWorkflowTasksRequest)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "RemoveTask",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RemoveTask),
				},
				Signature:    "RemoveTask(Request *shared.RemoveTaskRequest)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "ResendReplicationTasks",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.ResendReplicationTasks),
				},
				Signature:    "ResendReplicationTasks(Request *admin.ResendReplicationTasksRequest)",
				ThriftModule: admin.ThriftModule,
			},

			thrift.Method{
				Name: "ResetQueue",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.ResetQueue),
				},
				Signature:    "ResetQueue(Request *shared.ResetQueueRequest)",
				ThriftModule: admin.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 18)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

func (h handler) AddSearchAttribute(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_AddSearchAttribute_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.AddSearchAttribute(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_AddSearchAttribute_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) CloseShard(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_CloseShard_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.CloseShard(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_CloseShard_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) DescribeCluster(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_DescribeCluster_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.DescribeCluster(ctx)

	hadError := err != nil
	result, err := admin.AdminService_DescribeCluster_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) DescribeHistoryHost(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_DescribeHistoryHost_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.DescribeHistoryHost(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_DescribeHistoryHost_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) DescribeQueue(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_DescribeQueue_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.DescribeQueue(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_DescribeQueue_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) DescribeWorkflowExecution(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_DescribeWorkflowExecution_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.DescribeWorkflowExecution(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_DescribeWorkflowExecution_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) GetDLQReplicationMessages(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_GetDLQReplicationMessages_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.GetDLQReplicationMessages(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_GetDLQReplicationMessages_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) GetDomainReplicationMessages(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_GetDomainReplicationMessages_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.GetDomainReplicationMessages(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_GetDomainReplicationMessages_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) GetReplicationMessages(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_GetReplicationMessages_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.GetReplicationMessages(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_GetReplicationMessages_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) GetWorkflowExecutionRawHistoryV2(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_GetWorkflowExecutionRawHistoryV2_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.GetWorkflowExecutionRawHistoryV2(ctx, args.GetRequest)

	hadError := err != nil
	result, err := admin.AdminService_GetWorkflowExecutionRawHistoryV2_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) MergeDLQMessages(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_MergeDLQMessages_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.MergeDLQMessages(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_MergeDLQMessages_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) PurgeDLQMessages(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_PurgeDLQMessages_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.PurgeDLQMessages(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_PurgeDLQMessages_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) ReadDLQMessages(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_ReadDLQMessages_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.ReadDLQMessages(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_ReadDLQMessages_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) ReapplyEvents(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_ReapplyEvents_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.ReapplyEvents(ctx, args.ReapplyEventsRequest)

	hadError := err != nil
	result, err := admin.AdminService_ReapplyEvents_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RefreshWorkflowTasks(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_RefreshWorkflowTasks_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.RefreshWorkflowTasks(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_RefreshWorkflowTasks_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RemoveTask(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_RemoveTask_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.RemoveTask(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_RemoveTask_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) ResendReplicationTasks(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_ResendReplicationTasks_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.ResendReplicationTasks(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_ResendReplicationTasks_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) ResetQueue(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args admin.AdminService_ResetQueue_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.ResetQueue(ctx, args.Request)

	hadError := err != nil
	result, err := admin.AdminService_ResetQueue_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}