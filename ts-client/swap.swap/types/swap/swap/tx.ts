/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "swap.swap";

export interface MsgSend {
  creator: string;
  receiver: string;
  amount: string;
  amountToReceive: string;
}

export interface MsgSendResponse {
}

export interface MsgReceive {
  creator: string;
  id: string;
}

export interface MsgReceiveResponse {
}

export interface MsgCancel {
  creator: string;
  id: string;
}

export interface MsgCancelResponse {
}

function createBaseMsgSend(): MsgSend {
  return { creator: "", receiver: "", amount: "", amountToReceive: "" };
}

export const MsgSend = {
  encode(message: MsgSend, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    if (message.amountToReceive !== "") {
      writer.uint32(34).string(message.amountToReceive);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSend {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSend();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.receiver = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        case 4:
          message.amountToReceive = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSend {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      amountToReceive: isSet(object.amountToReceive) ? String(object.amountToReceive) : "",
    };
  },

  toJSON(message: MsgSend): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.amount !== undefined && (obj.amount = message.amount);
    message.amountToReceive !== undefined && (obj.amountToReceive = message.amountToReceive);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSend>, I>>(object: I): MsgSend {
    const message = createBaseMsgSend();
    message.creator = object.creator ?? "";
    message.receiver = object.receiver ?? "";
    message.amount = object.amount ?? "";
    message.amountToReceive = object.amountToReceive ?? "";
    return message;
  },
};

function createBaseMsgSendResponse(): MsgSendResponse {
  return {};
}

export const MsgSendResponse = {
  encode(_: MsgSendResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSendResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSendResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSendResponse {
    return {};
  },

  toJSON(_: MsgSendResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSendResponse>, I>>(_: I): MsgSendResponse {
    const message = createBaseMsgSendResponse();
    return message;
  },
};

function createBaseMsgReceive(): MsgReceive {
  return { creator: "", id: "" };
}

export const MsgReceive = {
  encode(message: MsgReceive, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgReceive {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgReceive();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgReceive {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? String(object.id) : "",
    };
  },

  toJSON(message: MsgReceive): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgReceive>, I>>(object: I): MsgReceive {
    const message = createBaseMsgReceive();
    message.creator = object.creator ?? "";
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseMsgReceiveResponse(): MsgReceiveResponse {
  return {};
}

export const MsgReceiveResponse = {
  encode(_: MsgReceiveResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgReceiveResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgReceiveResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgReceiveResponse {
    return {};
  },

  toJSON(_: MsgReceiveResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgReceiveResponse>, I>>(_: I): MsgReceiveResponse {
    const message = createBaseMsgReceiveResponse();
    return message;
  },
};

function createBaseMsgCancel(): MsgCancel {
  return { creator: "", id: "" };
}

export const MsgCancel = {
  encode(message: MsgCancel, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCancel {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCancel();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCancel {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? String(object.id) : "",
    };
  },

  toJSON(message: MsgCancel): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancel>, I>>(object: I): MsgCancel {
    const message = createBaseMsgCancel();
    message.creator = object.creator ?? "";
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseMsgCancelResponse(): MsgCancelResponse {
  return {};
}

export const MsgCancelResponse = {
  encode(_: MsgCancelResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCancelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCancelResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCancelResponse {
    return {};
  },

  toJSON(_: MsgCancelResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancelResponse>, I>>(_: I): MsgCancelResponse {
    const message = createBaseMsgCancelResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Send(request: MsgSend): Promise<MsgSendResponse>;
  Receive(request: MsgReceive): Promise<MsgReceiveResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Cancel(request: MsgCancel): Promise<MsgCancelResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Send = this.Send.bind(this);
    this.Receive = this.Receive.bind(this);
    this.Cancel = this.Cancel.bind(this);
  }
  Send(request: MsgSend): Promise<MsgSendResponse> {
    const data = MsgSend.encode(request).finish();
    const promise = this.rpc.request("swap.swap.Msg", "Send", data);
    return promise.then((data) => MsgSendResponse.decode(new _m0.Reader(data)));
  }

  Receive(request: MsgReceive): Promise<MsgReceiveResponse> {
    const data = MsgReceive.encode(request).finish();
    const promise = this.rpc.request("swap.swap.Msg", "Receive", data);
    return promise.then((data) => MsgReceiveResponse.decode(new _m0.Reader(data)));
  }

  Cancel(request: MsgCancel): Promise<MsgCancelResponse> {
    const data = MsgCancel.encode(request).finish();
    const promise = this.rpc.request("swap.swap.Msg", "Cancel", data);
    return promise.then((data) => MsgCancelResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
