/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "swap.sale";

export interface MsgSell {
  creator: string;
  amount: string;
  price: string;
}

export interface MsgSellResponse {
  id: number;
}

export interface MsgCancel {
  creator: string;
  id: string;
}

export interface MsgCancelResponse {
}

export interface MsgBuy {
  creator: string;
  id: string;
}

export interface MsgBuyResponse {
}

export interface MsgSellNFT {
  creator: string;
  classId: string;
  nftId: string;
  price: string;
}

export interface MsgSellNFTResponse {
  id: string;
}

export interface MsgBuyNFT {
  creator: string;
  id: string;
}

export interface MsgBuyNFTResponse {
}

export interface MsgCancelNFT {
  creator: string;
  id: string;
}

export interface MsgCancelNFTResponse {
}

function createBaseMsgSell(): MsgSell {
  return { creator: "", amount: "", price: "" };
}

export const MsgSell = {
  encode(message: MsgSell, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    if (message.price !== "") {
      writer.uint32(26).string(message.price);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSell {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSell();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        case 3:
          message.price = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSell {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      price: isSet(object.price) ? String(object.price) : "",
    };
  },

  toJSON(message: MsgSell): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    message.price !== undefined && (obj.price = message.price);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSell>, I>>(object: I): MsgSell {
    const message = createBaseMsgSell();
    message.creator = object.creator ?? "";
    message.amount = object.amount ?? "";
    message.price = object.price ?? "";
    return message;
  },
};

function createBaseMsgSellResponse(): MsgSellResponse {
  return { id: 0 };
}

export const MsgSellResponse = {
  encode(message: MsgSellResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSellResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSellResponse();
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

  fromJSON(object: any): MsgSellResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgSellResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSellResponse>, I>>(object: I): MsgSellResponse {
    const message = createBaseMsgSellResponse();
    message.id = object.id ?? 0;
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

function createBaseMsgBuy(): MsgBuy {
  return { creator: "", id: "" };
}

export const MsgBuy = {
  encode(message: MsgBuy, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBuy {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBuy();
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

  fromJSON(object: any): MsgBuy {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? String(object.id) : "",
    };
  },

  toJSON(message: MsgBuy): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBuy>, I>>(object: I): MsgBuy {
    const message = createBaseMsgBuy();
    message.creator = object.creator ?? "";
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseMsgBuyResponse(): MsgBuyResponse {
  return {};
}

export const MsgBuyResponse = {
  encode(_: MsgBuyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBuyResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBuyResponse();
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

  fromJSON(_: any): MsgBuyResponse {
    return {};
  },

  toJSON(_: MsgBuyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBuyResponse>, I>>(_: I): MsgBuyResponse {
    const message = createBaseMsgBuyResponse();
    return message;
  },
};

function createBaseMsgSellNFT(): MsgSellNFT {
  return { creator: "", classId: "", nftId: "", price: "" };
}

export const MsgSellNFT = {
  encode(message: MsgSellNFT, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.classId !== "") {
      writer.uint32(18).string(message.classId);
    }
    if (message.nftId !== "") {
      writer.uint32(26).string(message.nftId);
    }
    if (message.price !== "") {
      writer.uint32(34).string(message.price);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSellNFT {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSellNFT();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.classId = reader.string();
          break;
        case 3:
          message.nftId = reader.string();
          break;
        case 4:
          message.price = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSellNFT {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      classId: isSet(object.classId) ? String(object.classId) : "",
      nftId: isSet(object.nftId) ? String(object.nftId) : "",
      price: isSet(object.price) ? String(object.price) : "",
    };
  },

  toJSON(message: MsgSellNFT): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.classId !== undefined && (obj.classId = message.classId);
    message.nftId !== undefined && (obj.nftId = message.nftId);
    message.price !== undefined && (obj.price = message.price);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSellNFT>, I>>(object: I): MsgSellNFT {
    const message = createBaseMsgSellNFT();
    message.creator = object.creator ?? "";
    message.classId = object.classId ?? "";
    message.nftId = object.nftId ?? "";
    message.price = object.price ?? "";
    return message;
  },
};

function createBaseMsgSellNFTResponse(): MsgSellNFTResponse {
  return { id: "" };
}

export const MsgSellNFTResponse = {
  encode(message: MsgSellNFTResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSellNFTResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSellNFTResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSellNFTResponse {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: MsgSellNFTResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSellNFTResponse>, I>>(object: I): MsgSellNFTResponse {
    const message = createBaseMsgSellNFTResponse();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseMsgBuyNFT(): MsgBuyNFT {
  return { creator: "", id: "" };
}

export const MsgBuyNFT = {
  encode(message: MsgBuyNFT, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBuyNFT {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBuyNFT();
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

  fromJSON(object: any): MsgBuyNFT {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? String(object.id) : "",
    };
  },

  toJSON(message: MsgBuyNFT): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBuyNFT>, I>>(object: I): MsgBuyNFT {
    const message = createBaseMsgBuyNFT();
    message.creator = object.creator ?? "";
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseMsgBuyNFTResponse(): MsgBuyNFTResponse {
  return {};
}

export const MsgBuyNFTResponse = {
  encode(_: MsgBuyNFTResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBuyNFTResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBuyNFTResponse();
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

  fromJSON(_: any): MsgBuyNFTResponse {
    return {};
  },

  toJSON(_: MsgBuyNFTResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBuyNFTResponse>, I>>(_: I): MsgBuyNFTResponse {
    const message = createBaseMsgBuyNFTResponse();
    return message;
  },
};

function createBaseMsgCancelNFT(): MsgCancelNFT {
  return { creator: "", id: "" };
}

export const MsgCancelNFT = {
  encode(message: MsgCancelNFT, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
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
          message.id = reader.string();
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
      id: isSet(object.id) ? String(object.id) : "",
    };
  },

  toJSON(message: MsgCancelNFT): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancelNFT>, I>>(object: I): MsgCancelNFT {
    const message = createBaseMsgCancelNFT();
    message.creator = object.creator ?? "";
    message.id = object.id ?? "";
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

/** Msg defines the Msg service. */
export interface Msg {
  Sell(request: MsgSell): Promise<MsgSellResponse>;
  Cancel(request: MsgCancel): Promise<MsgCancelResponse>;
  Buy(request: MsgBuy): Promise<MsgBuyResponse>;
  SellNFT(request: MsgSellNFT): Promise<MsgSellNFTResponse>;
  BuyNFT(request: MsgBuyNFT): Promise<MsgBuyNFTResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CancelNFT(request: MsgCancelNFT): Promise<MsgCancelNFTResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Sell = this.Sell.bind(this);
    this.Cancel = this.Cancel.bind(this);
    this.Buy = this.Buy.bind(this);
    this.SellNFT = this.SellNFT.bind(this);
    this.BuyNFT = this.BuyNFT.bind(this);
    this.CancelNFT = this.CancelNFT.bind(this);
  }
  Sell(request: MsgSell): Promise<MsgSellResponse> {
    const data = MsgSell.encode(request).finish();
    const promise = this.rpc.request("swap.sale.Msg", "Sell", data);
    return promise.then((data) => MsgSellResponse.decode(new _m0.Reader(data)));
  }

  Cancel(request: MsgCancel): Promise<MsgCancelResponse> {
    const data = MsgCancel.encode(request).finish();
    const promise = this.rpc.request("swap.sale.Msg", "Cancel", data);
    return promise.then((data) => MsgCancelResponse.decode(new _m0.Reader(data)));
  }

  Buy(request: MsgBuy): Promise<MsgBuyResponse> {
    const data = MsgBuy.encode(request).finish();
    const promise = this.rpc.request("swap.sale.Msg", "Buy", data);
    return promise.then((data) => MsgBuyResponse.decode(new _m0.Reader(data)));
  }

  SellNFT(request: MsgSellNFT): Promise<MsgSellNFTResponse> {
    const data = MsgSellNFT.encode(request).finish();
    const promise = this.rpc.request("swap.sale.Msg", "SellNFT", data);
    return promise.then((data) => MsgSellNFTResponse.decode(new _m0.Reader(data)));
  }

  BuyNFT(request: MsgBuyNFT): Promise<MsgBuyNFTResponse> {
    const data = MsgBuyNFT.encode(request).finish();
    const promise = this.rpc.request("swap.sale.Msg", "BuyNFT", data);
    return promise.then((data) => MsgBuyNFTResponse.decode(new _m0.Reader(data)));
  }

  CancelNFT(request: MsgCancelNFT): Promise<MsgCancelNFTResponse> {
    const data = MsgCancelNFT.encode(request).finish();
    const promise = this.rpc.request("swap.sale.Msg", "CancelNFT", data);
    return promise.then((data) => MsgCancelNFTResponse.decode(new _m0.Reader(data)));
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
