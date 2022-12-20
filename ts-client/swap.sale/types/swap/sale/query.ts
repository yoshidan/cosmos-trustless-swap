/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Params } from "./params";
import { NFTSale, Sale } from "./sale";

export const protobufPackage = "swap.sale";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryShowRequest {
  seller: string;
  id: number;
}

export interface QueryShowResponse {
  sale: Sale | undefined;
}

export interface QueryShowNFTRequest {
  seller: string;
  id: number;
}

export interface QueryShowNFTResponse {
  sale: NFTSale | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryShowRequest(): QueryShowRequest {
  return { seller: "", id: 0 };
}

export const QueryShowRequest = {
  encode(message: QueryShowRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.seller !== "") {
      writer.uint32(10).string(message.seller);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryShowRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryShowRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.seller = reader.string();
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

  fromJSON(object: any): QueryShowRequest {
    return { seller: isSet(object.seller) ? String(object.seller) : "", id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryShowRequest): unknown {
    const obj: any = {};
    message.seller !== undefined && (obj.seller = message.seller);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryShowRequest>, I>>(object: I): QueryShowRequest {
    const message = createBaseQueryShowRequest();
    message.seller = object.seller ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryShowResponse(): QueryShowResponse {
  return { sale: undefined };
}

export const QueryShowResponse = {
  encode(message: QueryShowResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sale !== undefined) {
      Sale.encode(message.sale, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryShowResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryShowResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sale = Sale.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryShowResponse {
    return { sale: isSet(object.sale) ? Sale.fromJSON(object.sale) : undefined };
  },

  toJSON(message: QueryShowResponse): unknown {
    const obj: any = {};
    message.sale !== undefined && (obj.sale = message.sale ? Sale.toJSON(message.sale) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryShowResponse>, I>>(object: I): QueryShowResponse {
    const message = createBaseQueryShowResponse();
    message.sale = (object.sale !== undefined && object.sale !== null) ? Sale.fromPartial(object.sale) : undefined;
    return message;
  },
};

function createBaseQueryShowNFTRequest(): QueryShowNFTRequest {
  return { seller: "", id: 0 };
}

export const QueryShowNFTRequest = {
  encode(message: QueryShowNFTRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.seller !== "") {
      writer.uint32(10).string(message.seller);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryShowNFTRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryShowNFTRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.seller = reader.string();
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

  fromJSON(object: any): QueryShowNFTRequest {
    return { seller: isSet(object.seller) ? String(object.seller) : "", id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryShowNFTRequest): unknown {
    const obj: any = {};
    message.seller !== undefined && (obj.seller = message.seller);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryShowNFTRequest>, I>>(object: I): QueryShowNFTRequest {
    const message = createBaseQueryShowNFTRequest();
    message.seller = object.seller ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryShowNFTResponse(): QueryShowNFTResponse {
  return { sale: undefined };
}

export const QueryShowNFTResponse = {
  encode(message: QueryShowNFTResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sale !== undefined) {
      NFTSale.encode(message.sale, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryShowNFTResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryShowNFTResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sale = NFTSale.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryShowNFTResponse {
    return { sale: isSet(object.sale) ? NFTSale.fromJSON(object.sale) : undefined };
  },

  toJSON(message: QueryShowNFTResponse): unknown {
    const obj: any = {};
    message.sale !== undefined && (obj.sale = message.sale ? NFTSale.toJSON(message.sale) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryShowNFTResponse>, I>>(object: I): QueryShowNFTResponse {
    const message = createBaseQueryShowNFTResponse();
    message.sale = (object.sale !== undefined && object.sale !== null) ? NFTSale.fromPartial(object.sale) : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Show items. */
  Show(request: QueryShowRequest): Promise<QueryShowResponse>;
  /** Queries a list of ShowNFT items. */
  ShowNFT(request: QueryShowNFTRequest): Promise<QueryShowNFTResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Show = this.Show.bind(this);
    this.ShowNFT = this.ShowNFT.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("swap.sale.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Show(request: QueryShowRequest): Promise<QueryShowResponse> {
    const data = QueryShowRequest.encode(request).finish();
    const promise = this.rpc.request("swap.sale.Query", "Show", data);
    return promise.then((data) => QueryShowResponse.decode(new _m0.Reader(data)));
  }

  ShowNFT(request: QueryShowNFTRequest): Promise<QueryShowNFTResponse> {
    const data = QueryShowNFTRequest.encode(request).finish();
    const promise = this.rpc.request("swap.sale.Query", "ShowNFT", data);
    return promise.then((data) => QueryShowNFTResponse.decode(new _m0.Reader(data)));
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
