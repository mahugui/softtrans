// Code generated by protoc-gen-auxo 0.1, DO NOT EDIT.
// source: tcc_coordinator.proto

package contract

import (
	"context"

	"github.com/cuigh/auxo/net/rpc"
)

var (
	tccService                = &tccServiceClient{rpc.LazyClient{Name: "softtrans.coordinator"}}
	tccResourceManagerService = &tccResourceManagerServiceClient{rpc.LazyClient{Name: "softtrans.coordinator"}}
)

// TCCService comment
type TCCService interface {
	// Begin a transaction
	BeginTrans(context.Context, *BeginTransRequest) (*BeginTransResponse, error)
	// Try a step
	TryStep(context.Context, *TryStepRequest) (*TryStepResponse, error)
	// Confirm a transaction
	ConfirmTrans(context.Context, *ConfirmTransRequest) (*ConfirmTransResponse, error)
	// Mark a transaction confirmed
	ConfirmTransSuccess(context.Context, *ConfirmTransSuccessRequest) (*ConfirmTransSuccessResponse, error)
	// Start rolling back a transaction
	CancelTrans(context.Context, *CancelTransRequest) (*CancelTransResponse, error)
	// Mark a transaction rolled back
	CancelTransSuccess(context.Context, *CancelTransSuccessRequest) (*CancelTransSuccessResponse, error)
	// Get info of the transaction
	GetTrans(context.Context, *GetTransRequest) (*GetTransResponse, error)
	// Get expired transactions
	GetExpiredTransList(context.Context, *GetExpiredTransListRequest) (*GetExpiredTransListResponse, error)
	// Get confirming transactions
	GetConfirmingTransList(context.Context, *GetConfirmingTransListRequest) (*GetConfirmingTransListResponse, error)
	// Get cancelling transactions
	GetCancellingTransList(context.Context, *GetCancellingTransListRequest) (*GetCancellingTransListResponse, error)
}

func GetTCCService() TCCService {
	return tccService
}

type tccServiceClient struct {
	rpc.LazyClient
}

func (s *tccServiceClient) BeginTrans(ctx context.Context, req *BeginTransRequest) (*BeginTransResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(BeginTransResponse)
	err = c.Call(ctx, "TCCService", "BeginTrans", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *tccServiceClient) TryStep(ctx context.Context, req *TryStepRequest) (*TryStepResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(TryStepResponse)
	err = c.Call(ctx, "TCCService", "TryStep", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *tccServiceClient) ConfirmTrans(ctx context.Context, req *ConfirmTransRequest) (*ConfirmTransResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(ConfirmTransResponse)
	err = c.Call(ctx, "TCCService", "ConfirmTrans", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *tccServiceClient) ConfirmTransSuccess(ctx context.Context, req *ConfirmTransSuccessRequest) (*ConfirmTransSuccessResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(ConfirmTransSuccessResponse)
	err = c.Call(ctx, "TCCService", "ConfirmTransSuccess", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *tccServiceClient) CancelTrans(ctx context.Context, req *CancelTransRequest) (*CancelTransResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(CancelTransResponse)
	err = c.Call(ctx, "TCCService", "CancelTrans", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *tccServiceClient) CancelTransSuccess(ctx context.Context, req *CancelTransSuccessRequest) (*CancelTransSuccessResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(CancelTransSuccessResponse)
	err = c.Call(ctx, "TCCService", "CancelTransSuccess", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *tccServiceClient) GetTrans(ctx context.Context, req *GetTransRequest) (*GetTransResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(GetTransResponse)
	err = c.Call(ctx, "TCCService", "GetTrans", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *tccServiceClient) GetExpiredTransList(ctx context.Context, req *GetExpiredTransListRequest) (*GetExpiredTransListResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(GetExpiredTransListResponse)
	err = c.Call(ctx, "TCCService", "GetExpiredTransList", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *tccServiceClient) GetConfirmingTransList(ctx context.Context, req *GetConfirmingTransListRequest) (*GetConfirmingTransListResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(GetConfirmingTransListResponse)
	err = c.Call(ctx, "TCCService", "GetConfirmingTransList", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *tccServiceClient) GetCancellingTransList(ctx context.Context, req *GetCancellingTransListRequest) (*GetCancellingTransListResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(GetCancellingTransListResponse)
	err = c.Call(ctx, "TCCService", "GetCancellingTransList", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type TCCResourceManagerService interface {
	// Confirm (require idempotent)
	Confirm(context.Context, *RMConfirmTransRequest) (*RMConfirmTransResponse, error)
	// Cancel (require idempotent)
	Cancel(context.Context, *RMCancelTransRequest) (*RMCancelTransResponse, error)
}

func GetTCCResourceManagerService() TCCResourceManagerService {
	return tccResourceManagerService
}

type tccResourceManagerServiceClient struct {
	rpc.LazyClient
}

func (s *tccResourceManagerServiceClient) Confirm(ctx context.Context, req *RMConfirmTransRequest) (*RMConfirmTransResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(RMConfirmTransResponse)
	err = c.Call(ctx, "TCCResourceManagerService", "Confirm", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *tccResourceManagerServiceClient) Cancel(ctx context.Context, req *RMCancelTransRequest) (*RMCancelTransResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(RMCancelTransResponse)
	err = c.Call(ctx, "TCCResourceManagerService", "Cancel", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
