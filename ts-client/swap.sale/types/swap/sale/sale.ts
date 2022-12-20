/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "swap.sale";

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface Sale {
  id: number;
  seller: string;
  amount: string;
  price: string;
}

export interface NFTSale {
  id: number;
  seller: string;
  classId: string;
  nftId: string;
  price: string;
}

function createBaseSale(): Sale {
  return { id: 0, seller: "", amount: "", price: "" };
}

export const Sale = {
  encode(message: Sale, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.seller !== "") {
      writer.uint32(18).string(message.seller);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    if (message.price !== "") {
      writer.uint32(34).string(message.price);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Sale {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSale();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.seller = reader.string();
          break;
        case 3:
          message.amount = reader.string();
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

  fromJSON(object: any): Sale {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      seller: isSet(object.seller) ? String(object.seller) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      price: isSet(object.price) ? String(object.price) : "",
    };
  },

  toJSON(message: Sale): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.seller !== undefined && (obj.seller = message.seller);
    message.amount !== undefined && (obj.amount = message.amount);
    message.price !== undefined && (obj.price = message.price);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Sale>, I>>(object: I): Sale {
    const message = createBaseSale();
    message.id = object.id ?? 0;
    message.seller = object.seller ?? "";
    message.amount = object.amount ?? "";
    message.price = object.price ?? "";
    return message;
  },
};

function createBaseNFTSale(): NFTSale {
  return { id: 0, seller: "", classId: "", nftId: "", price: "" };
}

export const NFTSale = {
  encode(message: NFTSale, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.seller !== "") {
      writer.uint32(18).string(message.seller);
    }
    if (message.classId !== "") {
      writer.uint32(26).string(message.classId);
    }
    if (message.nftId !== "") {
      writer.uint32(34).string(message.nftId);
    }
    if (message.price !== "") {
      writer.uint32(42).string(message.price);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NFTSale {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNFTSale();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.seller = reader.string();
          break;
        case 3:
          message.classId = reader.string();
          break;
        case 4:
          message.nftId = reader.string();
          break;
        case 5:
          message.price = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NFTSale {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      seller: isSet(object.seller) ? String(object.seller) : "",
      classId: isSet(object.classId) ? String(object.classId) : "",
      nftId: isSet(object.nftId) ? String(object.nftId) : "",
      price: isSet(object.price) ? String(object.price) : "",
    };
  },

  toJSON(message: NFTSale): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.seller !== undefined && (obj.seller = message.seller);
    message.classId !== undefined && (obj.classId = message.classId);
    message.nftId !== undefined && (obj.nftId = message.nftId);
    message.price !== undefined && (obj.price = message.price);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NFTSale>, I>>(object: I): NFTSale {
    const message = createBaseNFTSale();
    message.id = object.id ?? 0;
    message.seller = object.seller ?? "";
    message.classId = object.classId ?? "";
    message.nftId = object.nftId ?? "";
    message.price = object.price ?? "";
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
