/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "teapartycrypto.partychain.party";

export interface TradeOrders {
  index: string;
  tradeAsset: string;
  price: string;
  currency: string;
  amount: string;
  sellerShippingAddr: string;
  sellerNknAddr: string;
  refundAddr: string;
}

function createBaseTradeOrders(): TradeOrders {
  return {
    index: "",
    tradeAsset: "",
    price: "",
    currency: "",
    amount: "",
    sellerShippingAddr: "",
    sellerNknAddr: "",
    refundAddr: "",
  };
}

export const TradeOrders = {
  encode(message: TradeOrders, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.tradeAsset !== "") {
      writer.uint32(18).string(message.tradeAsset);
    }
    if (message.price !== "") {
      writer.uint32(26).string(message.price);
    }
    if (message.currency !== "") {
      writer.uint32(34).string(message.currency);
    }
    if (message.amount !== "") {
      writer.uint32(42).string(message.amount);
    }
    if (message.sellerShippingAddr !== "") {
      writer.uint32(50).string(message.sellerShippingAddr);
    }
    if (message.sellerNknAddr !== "") {
      writer.uint32(58).string(message.sellerNknAddr);
    }
    if (message.refundAddr !== "") {
      writer.uint32(66).string(message.refundAddr);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TradeOrders {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTradeOrders();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.tradeAsset = reader.string();
          break;
        case 3:
          message.price = reader.string();
          break;
        case 4:
          message.currency = reader.string();
          break;
        case 5:
          message.amount = reader.string();
          break;
        case 6:
          message.sellerShippingAddr = reader.string();
          break;
        case 7:
          message.sellerNknAddr = reader.string();
          break;
        case 8:
          message.refundAddr = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TradeOrders {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      tradeAsset: isSet(object.tradeAsset) ? String(object.tradeAsset) : "",
      price: isSet(object.price) ? String(object.price) : "",
      currency: isSet(object.currency) ? String(object.currency) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      sellerShippingAddr: isSet(object.sellerShippingAddr) ? String(object.sellerShippingAddr) : "",
      sellerNknAddr: isSet(object.sellerNknAddr) ? String(object.sellerNknAddr) : "",
      refundAddr: isSet(object.refundAddr) ? String(object.refundAddr) : "",
    };
  },

  toJSON(message: TradeOrders): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.tradeAsset !== undefined && (obj.tradeAsset = message.tradeAsset);
    message.price !== undefined && (obj.price = message.price);
    message.currency !== undefined && (obj.currency = message.currency);
    message.amount !== undefined && (obj.amount = message.amount);
    message.sellerShippingAddr !== undefined && (obj.sellerShippingAddr = message.sellerShippingAddr);
    message.sellerNknAddr !== undefined && (obj.sellerNknAddr = message.sellerNknAddr);
    message.refundAddr !== undefined && (obj.refundAddr = message.refundAddr);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<TradeOrders>, I>>(object: I): TradeOrders {
    const message = createBaseTradeOrders();
    message.index = object.index ?? "";
    message.tradeAsset = object.tradeAsset ?? "";
    message.price = object.price ?? "";
    message.currency = object.currency ?? "";
    message.amount = object.amount ?? "";
    message.sellerShippingAddr = object.sellerShippingAddr ?? "";
    message.sellerNknAddr = object.sellerNknAddr ?? "";
    message.refundAddr = object.refundAddr ?? "";
    return message;
  },
};

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
