/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Params } from "./params";

export const protobufPackage = "swap.swap";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryShowRequest {
  id: string;
}

export interface QueryShowResponse {
  sender: string;
  receiver: string;
  amountToSend: string;
  amountToReceive: string;
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
  return { id: "" };
}

export const QueryShowRequest = {
  encode(message: QueryShowRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
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
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryShowRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: QueryShowRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryShowRequest>, I>>(object: I): QueryShowRequest {
    const message = createBaseQueryShowRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseQueryShowResponse(): QueryShowResponse {
  return { sender: "", receiver: "", amountToSend: "", amountToReceive: "" };
}

export const QueryShowResponse = {
  encode(message: QueryShowResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.amountToSend !== "") {
      writer.uint32(26).string(message.amountToSend);
    }
    if (message.amountToReceive !== "") {
      writer.uint32(34).string(message.amountToReceive);
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
          message.sender = reader.string();
          break;
        case 2:
          message.receiver = reader.string();
          break;
        case 3:
          message.amountToSend = reader.string();
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

  fromJSON(object: any): QueryShowResponse {
    return {
      sender: isSet(object.sender) ? String(object.sender) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      amountToSend: isSet(object.amountToSend) ? String(object.amountToSend) : "",
      amountToReceive: isSet(object.amountToReceive) ? String(object.amountToReceive) : "",
    };
  },

  toJSON(message: QueryShowResponse): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.amountToSend !== undefined && (obj.amountToSend = message.amountToSend);
    message.amountToReceive !== undefined && (obj.amountToReceive = message.amountToReceive);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryShowResponse>, I>>(object: I): QueryShowResponse {
    const message = createBaseQueryShowResponse();
    message.sender = object.sender ?? "";
    message.receiver = object.receiver ?? "";
    message.amountToSend = object.amountToSend ?? "";
    message.amountToReceive = object.amountToReceive ?? "";
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Show items. */
  Show(request: QueryShowRequest): Promise<QueryShowResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Show = this.Show.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("swap.swap.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Show(request: QueryShowRequest): Promise<QueryShowResponse> {
    const data = QueryShowRequest.encode(request).finish();
    const promise = this.rpc.request("swap.swap.Query", "Show", data);
    return promise.then((data) => QueryShowResponse.decode(new _m0.Reader(data)));
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
