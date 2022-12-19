/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "swap.swap";

export enum SwapStatus {
  Active = 0,
  Closed = 1,
  Cancelled = 2,
  UNRECOGNIZED = -1,
}

export function swapStatusFromJSON(object: any): SwapStatus {
  switch (object) {
    case 0:
    case "Active":
      return SwapStatus.Active;
    case 1:
    case "Closed":
      return SwapStatus.Closed;
    case 2:
    case "Cancelled":
      return SwapStatus.Cancelled;
    case -1:
    case "UNRECOGNIZED":
    default:
      return SwapStatus.UNRECOGNIZED;
  }
}

export function swapStatusToJSON(object: SwapStatus): string {
  switch (object) {
    case SwapStatus.Active:
      return "Active";
    case SwapStatus.Closed:
      return "Closed";
    case SwapStatus.Cancelled:
      return "Cancelled";
    case SwapStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface Swap {
  id: number;
  sender: string;
  receiver: string;
  amount: string;
  amountToReceive: string;
  status: SwapStatus;
}

function createBaseSwap(): Swap {
  return { id: 0, sender: "", receiver: "", amount: "", amountToReceive: "", status: 0 };
}

export const Swap = {
  encode(message: Swap, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.sender !== "") {
      writer.uint32(18).string(message.sender);
    }
    if (message.receiver !== "") {
      writer.uint32(26).string(message.receiver);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    if (message.amountToReceive !== "") {
      writer.uint32(42).string(message.amountToReceive);
    }
    if (message.status !== 0) {
      writer.uint32(48).int32(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Swap {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSwap();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.sender = reader.string();
          break;
        case 3:
          message.receiver = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        case 5:
          message.amountToReceive = reader.string();
          break;
        case 6:
          message.status = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Swap {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      sender: isSet(object.sender) ? String(object.sender) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      amountToReceive: isSet(object.amountToReceive) ? String(object.amountToReceive) : "",
      status: isSet(object.status) ? swapStatusFromJSON(object.status) : 0,
    };
  },

  toJSON(message: Swap): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.sender !== undefined && (obj.sender = message.sender);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.amount !== undefined && (obj.amount = message.amount);
    message.amountToReceive !== undefined && (obj.amountToReceive = message.amountToReceive);
    message.status !== undefined && (obj.status = swapStatusToJSON(message.status));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Swap>, I>>(object: I): Swap {
    const message = createBaseSwap();
    message.id = object.id ?? 0;
    message.sender = object.sender ?? "";
    message.receiver = object.receiver ?? "";
    message.amount = object.amount ?? "";
    message.amountToReceive = object.amountToReceive ?? "";
    message.status = object.status ?? 0;
    return message;
  },
};

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
