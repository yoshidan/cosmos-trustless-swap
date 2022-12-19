/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "swap.swap";

export interface MsgSend {
  creator: string;
  receiver: string;
  amount: string;
  amountToReceive: string;
}

export interface MsgSendResponse {
  id: number;
}

export interface MsgReceive {
  creator: string;
  id: number;
}

export interface MsgReceiveResponse {
}

export interface MsgCancel {
  creator: string;
  id: number;
}

export interface MsgCancelResponse {
}

export interface MsgSendNFT {
  creator: string;
  receiver: string;
  classId: string;
  nftId: string;
  amountToReceive: string;
}

export interface MsgSendNFTResponse {
  id: number;
}

export interface MsgCancelNFT {
  creator: string;
  id: number;
}

export interface MsgCancelNFTResponse {
}

export interface MsgReceiveNFT {
  creator: string;
  id: number;
}

export interface MsgReceiveNFTResponse {
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
  return { id: 0 };
}

export const MsgSendResponse = {
  encode(message: MsgSendResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSendResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSendResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgSendResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSendResponse>, I>>(object: I): MsgSendResponse {
    const message = createBaseMsgSendResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgReceive(): MsgReceive {
  return { creator: "", id: 0 };
}

export const MsgReceive = {
  encode(message: MsgReceive, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
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
          message.id = longToNumber(reader.uint64() as Long);
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
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgReceive): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgReceive>, I>>(object: I): MsgReceive {
    const message = createBaseMsgReceive();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
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
  return { creator: "", id: 0 };
}

export const MsgCancel = {
  encode(message: MsgCancel, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
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
          message.id = longToNumber(reader.uint64() as Long);
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
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgCancel): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancel>, I>>(object: I): MsgCancel {
    const message = createBaseMsgCancel();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
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

function createBaseMsgSendNFT(): MsgSendNFT {
  return { creator: "", receiver: "", classId: "", nftId: "", amountToReceive: "" };
}

export const MsgSendNFT = {
  encode(message: MsgSendNFT, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.classId !== "") {
      writer.uint32(26).string(message.classId);
    }
    if (message.nftId !== "") {
      writer.uint32(34).string(message.nftId);
    }
    if (message.amountToReceive !== "") {
      writer.uint32(42).string(message.amountToReceive);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSendNFT {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSendNFT();
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
          message.classId = reader.string();
          break;
        case 4:
          message.nftId = reader.string();
          break;
        case 5:
          message.amountToReceive = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendNFT {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      classId: isSet(object.classId) ? String(object.classId) : "",
      nftId: isSet(object.nftId) ? String(object.nftId) : "",
      amountToReceive: isSet(object.amountToReceive) ? String(object.amountToReceive) : "",
    };
  },

  toJSON(message: MsgSendNFT): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.classId !== undefined && (obj.classId = message.classId);
    message.nftId !== undefined && (obj.nftId = message.nftId);
    message.amountToReceive !== undefined && (obj.amountToReceive = message.amountToReceive);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSendNFT>, I>>(object: I): MsgSendNFT {
    const message = createBaseMsgSendNFT();
    message.creator = object.creator ?? "";
    message.receiver = object.receiver ?? "";
    message.classId = object.classId ?? "";
    message.nftId = object.nftId ?? "";
    message.amountToReceive = object.amountToReceive ?? "";
    return message;
  },
};

function createBaseMsgSendNFTResponse(): MsgSendNFTResponse {
  return { id: 0 };
}

export const MsgSendNFTResponse = {
  encode(message: MsgSendNFTResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSendNFTResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSendNFTResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendNFTResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgSendNFTResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSendNFTResponse>, I>>(object: I): MsgSendNFTResponse {
    const message = createBaseMsgSendNFTResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgCancelNFT(): MsgCancelNFT {
  return { creator: "", id: 0 };
}

export const MsgCancelNFT = {
  encode(message: MsgCancelNFT, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCancelNFT {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCancelNFT();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCancelNFT {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgCancelNFT): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancelNFT>, I>>(object: I): MsgCancelNFT {
    const message = createBaseMsgCancelNFT();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgCancelNFTResponse(): MsgCancelNFTResponse {
  return {};
}

export const MsgCancelNFTResponse = {
  encode(_: MsgCancelNFTResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCancelNFTResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCancelNFTResponse();
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

  fromJSON(_: any): MsgCancelNFTResponse {
    return {};
  },

  toJSON(_: MsgCancelNFTResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancelNFTResponse>, I>>(_: I): MsgCancelNFTResponse {
    const message = createBaseMsgCancelNFTResponse();
    return message;
  },
};

function createBaseMsgReceiveNFT(): MsgReceiveNFT {
  return { creator: "", id: 0 };
}

export const MsgReceiveNFT = {
  encode(message: MsgReceiveNFT, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgReceiveNFT {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgReceiveNFT();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgReceiveNFT {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgReceiveNFT): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgReceiveNFT>, I>>(object: I): MsgReceiveNFT {
    const message = createBaseMsgReceiveNFT();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgReceiveNFTResponse(): MsgReceiveNFTResponse {
  return {};
}

export const MsgReceiveNFTResponse = {
  encode(_: MsgReceiveNFTResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgReceiveNFTResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgReceiveNFTResponse();
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

  fromJSON(_: any): MsgReceiveNFTResponse {
    return {};
  },

  toJSON(_: MsgReceiveNFTResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgReceiveNFTResponse>, I>>(_: I): MsgReceiveNFTResponse {
    const message = createBaseMsgReceiveNFTResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Send(request: MsgSend): Promise<MsgSendResponse>;
  Receive(request: MsgReceive): Promise<MsgReceiveResponse>;
  Cancel(request: MsgCancel): Promise<MsgCancelResponse>;
  SendNFT(request: MsgSendNFT): Promise<MsgSendNFTResponse>;
  CancelNFT(request: MsgCancelNFT): Promise<MsgCancelNFTResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ReceiveNFT(request: MsgReceiveNFT): Promise<MsgReceiveNFTResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Send = this.Send.bind(this);
    this.Receive = this.Receive.bind(this);
    this.Cancel = this.Cancel.bind(this);
    this.SendNFT = this.SendNFT.bind(this);
    this.CancelNFT = this.CancelNFT.bind(this);
    this.ReceiveNFT = this.ReceiveNFT.bind(this);
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

  SendNFT(request: MsgSendNFT): Promise<MsgSendNFTResponse> {
    const data = MsgSendNFT.encode(request).finish();
    const promise = this.rpc.request("swap.swap.Msg", "SendNFT", data);
    return promise.then((data) => MsgSendNFTResponse.decode(new _m0.Reader(data)));
  }

  CancelNFT(request: MsgCancelNFT): Promise<MsgCancelNFTResponse> {
    const data = MsgCancelNFT.encode(request).finish();
    const promise = this.rpc.request("swap.swap.Msg", "CancelNFT", data);
    return promise.then((data) => MsgCancelNFTResponse.decode(new _m0.Reader(data)));
  }

  ReceiveNFT(request: MsgReceiveNFT): Promise<MsgReceiveNFTResponse> {
    const data = MsgReceiveNFT.encode(request).finish();
    const promise = this.rpc.request("swap.swap.Msg", "ReceiveNFT", data);
    return promise.then((data) => MsgReceiveNFTResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
