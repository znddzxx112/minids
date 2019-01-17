// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package tpingservice

import (
	"bytes"
	"reflect"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type TpingService interface {
  // Parameters:
  //  - Action
  //  - Content
  Ping(ctx context.Context, action string, content string) (r string, err error)
}

type TpingServiceClient struct {
  c thrift.TClient
}

// Deprecated: Use NewTpingService instead
func NewTpingServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TpingServiceClient {
  return &TpingServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

// Deprecated: Use NewTpingService instead
func NewTpingServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TpingServiceClient {
  return &TpingServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewTpingServiceClient(c thrift.TClient) *TpingServiceClient {
  return &TpingServiceClient{
    c: c,
  }
}

// Parameters:
//  - Action
//  - Content
func (p *TpingServiceClient) Ping(ctx context.Context, action string, content string) (r string, err error) {
  var _args0 TpingServicePingArgs
  _args0.Action = action
  _args0.Content = content
  var _result1 TpingServicePingResult
  if err = p.c.Call(ctx, "ping", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

type TpingServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler TpingService
}

func (p *TpingServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *TpingServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *TpingServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewTpingServiceProcessor(handler TpingService) *TpingServiceProcessor {

  self2 := &TpingServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["ping"] = &tpingServiceProcessorPing{handler:handler}
return self2
}

func (p *TpingServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x3

}

type tpingServiceProcessorPing struct {
  handler TpingService
}

func (p *tpingServiceProcessorPing) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := TpingServicePingArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("ping", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := TpingServicePingResult{}
var retval string
  var err2 error
  if retval, err2 = p.handler.Ping(ctx, args.Action, args.Content); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ping: " + err2.Error())
    oprot.WriteMessageBegin("ping", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("ping", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Action
//  - Content
type TpingServicePingArgs struct {
  Action string `thrift:"action,1" db:"action" json:"action"`
  Content string `thrift:"content,2" db:"content" json:"content"`
}

func NewTpingServicePingArgs() *TpingServicePingArgs {
  return &TpingServicePingArgs{}
}


func (p *TpingServicePingArgs) GetAction() string {
  return p.Action
}

func (p *TpingServicePingArgs) GetContent() string {
  return p.Content
}
func (p *TpingServicePingArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TpingServicePingArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Action = v
}
  return nil
}

func (p *TpingServicePingArgs)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Content = v
}
  return nil
}

func (p *TpingServicePingArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("ping_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TpingServicePingArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("action", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:action: ", p), err) }
  if err := oprot.WriteString(string(p.Action)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.action (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:action: ", p), err) }
  return err
}

func (p *TpingServicePingArgs) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("content", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:content: ", p), err) }
  if err := oprot.WriteString(string(p.Content)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.content (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:content: ", p), err) }
  return err
}

func (p *TpingServicePingArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TpingServicePingArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TpingServicePingResult struct {
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewTpingServicePingResult() *TpingServicePingResult {
  return &TpingServicePingResult{}
}

var TpingServicePingResult_Success_DEFAULT string
func (p *TpingServicePingResult) GetSuccess() string {
  if !p.IsSetSuccess() {
    return TpingServicePingResult_Success_DEFAULT
  }
return *p.Success
}
func (p *TpingServicePingResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *TpingServicePingResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TpingServicePingResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *TpingServicePingResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("ping_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TpingServicePingResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *TpingServicePingResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TpingServicePingResult(%+v)", *p)
}


