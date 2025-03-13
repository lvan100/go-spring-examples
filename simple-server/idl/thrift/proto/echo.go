// Code generated by Thrift Compiler (0.16.0). DO NOT EDIT.

package proto

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

// Attributes:
//   - Message
type EchoRequest struct {
	Message string `thrift:"message,1,required" db:"message" json:"message"`
}

func NewEchoRequest() *EchoRequest {
	return &EchoRequest{}
}

func (p *EchoRequest) GetMessage() string {
	return p.Message
}
func (p *EchoRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetMessage bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
				issetMessage = true
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetMessage {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Message is not set"))
	}
	return nil
}

func (p *EchoRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = v
	}
	return nil
}

func (p *EchoRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "EchoRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EchoRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "message", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Message)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
	}
	return err
}

func (p *EchoRequest) Equals(other *EchoRequest) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.Message != other.Message {
		return false
	}
	return true
}

func (p *EchoRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoRequest(%+v)", *p)
}

// Attributes:
//   - Message
type EchoResponse struct {
	Message string `thrift:"message,1,required" db:"message" json:"message"`
}

func NewEchoResponse() *EchoResponse {
	return &EchoResponse{}
}

func (p *EchoResponse) GetMessage() string {
	return p.Message
}
func (p *EchoResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetMessage bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
				issetMessage = true
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetMessage {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Message is not set"))
	}
	return nil
}

func (p *EchoResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = v
	}
	return nil
}

func (p *EchoResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "EchoResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EchoResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "message", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Message)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
	}
	return err
}

func (p *EchoResponse) Equals(other *EchoResponse) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.Message != other.Message {
		return false
	}
	return true
}

func (p *EchoResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoResponse(%+v)", *p)
}

type EchoService interface {
	// Parameters:
	//  - Req
	Echo(ctx context.Context, req *EchoRequest) (_r *EchoResponse, _err error)
}

type EchoServiceClient struct {
	c    thrift.TClient
	meta thrift.ResponseMeta
}

func NewEchoServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *EchoServiceClient {
	return &EchoServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewEchoServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *EchoServiceClient {
	return &EchoServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewEchoServiceClient(c thrift.TClient) *EchoServiceClient {
	return &EchoServiceClient{
		c: c,
	}
}

func (p *EchoServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *EchoServiceClient) LastResponseMeta_() thrift.ResponseMeta {
	return p.meta
}

func (p *EchoServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
	p.meta = meta
}

// Parameters:
//   - Req
func (p *EchoServiceClient) Echo(ctx context.Context, req *EchoRequest) (_r *EchoResponse, _err error) {
	var _args0 EchoServiceEchoArgs
	_args0.Req = req
	var _result2 EchoServiceEchoResult
	var _meta1 thrift.ResponseMeta
	_meta1, _err = p.Client_().Call(ctx, "echo", &_args0, &_result2)
	p.SetLastResponseMeta_(_meta1)
	if _err != nil {
		return
	}
	if _ret3 := _result2.GetSuccess(); _ret3 != nil {
		return _ret3, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "echo failed: unknown result")
}

type EchoServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      EchoService
}

func (p *EchoServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *EchoServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *EchoServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewEchoServiceProcessor(handler EchoService) *EchoServiceProcessor {

	self4 := &EchoServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self4.processorMap["echo"] = &echoServiceProcessorEcho{handler: handler}
	return self4
}

func (p *EchoServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
	if err2 != nil {
		return false, thrift.WrapTException(err2)
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(ctx, thrift.STRUCT)
	iprot.ReadMessageEnd(ctx)
	x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
	x5.Write(ctx, oprot)
	oprot.WriteMessageEnd(ctx)
	oprot.Flush(ctx)
	return false, x5

}

type echoServiceProcessorEcho struct {
	handler EchoService
}

func (p *echoServiceProcessorEcho) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := EchoServiceEchoArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "echo", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := EchoServiceEchoResult{}
	var retval *EchoResponse
	if retval, err2 = p.handler.Echo(ctx, args.Req); err2 != nil {
		tickerCancel()
		if err2 == thrift.ErrAbandonRequest {
			return false, thrift.WrapTException(err2)
		}
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing echo: "+err2.Error())
		oprot.WriteMessageBegin(ctx, "echo", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return true, thrift.WrapTException(err2)
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 = oprot.WriteMessageBegin(ctx, "echo", thrift.REPLY, seqId); err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//   - Req
type EchoServiceEchoArgs struct {
	Req *EchoRequest `thrift:"req,1" db:"req" json:"req"`
}

func NewEchoServiceEchoArgs() *EchoServiceEchoArgs {
	return &EchoServiceEchoArgs{}
}

var EchoServiceEchoArgs_Req_DEFAULT *EchoRequest

func (p *EchoServiceEchoArgs) GetReq() *EchoRequest {
	if !p.IsSetReq() {
		return EchoServiceEchoArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *EchoServiceEchoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *EchoServiceEchoArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *EchoServiceEchoArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Req = &EchoRequest{}
	if err := p.Req.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
	}
	return nil
}

func (p *EchoServiceEchoArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "echo_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EchoServiceEchoArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "req", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err)
	}
	if err := p.Req.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err)
	}
	return err
}

func (p *EchoServiceEchoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoArgs(%+v)", *p)
}

// Attributes:
//   - Success
type EchoServiceEchoResult struct {
	Success *EchoResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewEchoServiceEchoResult() *EchoServiceEchoResult {
	return &EchoServiceEchoResult{}
}

var EchoServiceEchoResult_Success_DEFAULT *EchoResponse

func (p *EchoServiceEchoResult) GetSuccess() *EchoResponse {
	if !p.IsSetSuccess() {
		return EchoServiceEchoResult_Success_DEFAULT
	}
	return p.Success
}
func (p *EchoServiceEchoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EchoServiceEchoResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *EchoServiceEchoResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &EchoResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *EchoServiceEchoResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "echo_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EchoServiceEchoResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *EchoServiceEchoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoResult(%+v)", *p)
}
