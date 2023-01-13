/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "teapartycrypto.partychain.party";

export interface OrdersUnderWatch {
  index: string;
  nknAddress: string;
  walletPrivateKey: string;
  walletPublicKey: string;
  shippingAddress: string;
  refundAddress: string;
  amount: string;
  chain: string;
  paymentComplete: boolean;
}

function createBaseOrdersUnderWatch(): OrdersUnderWatch {
  return {
    index: "",
    nknAddress: "",
    walletPrivateKey: "",
    walletPublicKey: "",
    shippingAddress: "",
    refundAddress: "",
    amount: "",
    chain: "",
    paymentComplete: false,
  };
}

export const OrdersUnderWatch = {
  encode(message: OrdersUnderWatch, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.nknAddress !== "") {
      writer.uint32(18).string(message.nknAddress);
    }
    if (message.walletPrivateKey !== "") {
      writer.uint32(26).string(message.walletPrivateKey);
    }
    if (message.walletPublicKey !== "") {
      writer.uint32(34).string(message.walletPublicKey);
    }
    if (message.shippingAddress !== "") {
      writer.uint32(42).string(message.shippingAddress);
    }
    if (message.refundAddress !== "") {
      writer.uint32(50).string(message.refundAddress);
    }
    if (message.amount !== "") {
      writer.uint32(58).string(message.amount);
    }
    if (message.chain !== "") {
      writer.uint32(66).string(message.chain);
    }
    if (message.paymentComplete === true) {
      writer.uint32(72).bool(message.paymentComplete);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OrdersUnderWatch {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOrdersUnderWatch();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.nknAddress = reader.string();
          break;
        case 3:
          message.walletPrivateKey = reader.string();
          break;
        case 4:
          message.walletPublicKey = reader.string();
          break;
        case 5:
          message.shippingAddress = reader.string();
          break;
        case 6:
          message.refundAddress = reader.string();
          break;
        case 7:
          message.amount = reader.string();
          break;
        case 8:
          message.chain = reader.string();
          break;
        case 9:
          message.paymentComplete = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OrdersUnderWatch {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      nknAddress: isSet(object.nknAddress) ? String(object.nknAddress) : "",
      walletPrivateKey: isSet(object.walletPrivateKey) ? String(object.walletPrivateKey) : "",
      walletPublicKey: isSet(object.walletPublicKey) ? String(object.walletPublicKey) : "",
      shippingAddress: isSet(object.shippingAddress) ? String(object.shippingAddress) : "",
      refundAddress: isSet(object.refundAddress) ? String(object.refundAddress) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      chain: isSet(object.chain) ? String(object.chain) : "",
      paymentComplete: isSet(object.paymentComplete) ? Boolean(object.paymentComplete) : false,
    };
  },

  toJSON(message: OrdersUnderWatch): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.nknAddress !== undefined && (obj.nknAddress = message.nknAddress);
    message.walletPrivateKey !== undefined && (obj.walletPrivateKey = message.walletPrivateKey);
    message.walletPublicKey !== undefined && (obj.walletPublicKey = message.walletPublicKey);
    message.shippingAddress !== undefined && (obj.shippingAddress = message.shippingAddress);
    message.refundAddress !== undefined && (obj.refundAddress = message.refundAddress);
    message.amount !== undefined && (obj.amount = message.amount);
    message.chain !== undefined && (obj.chain = message.chain);
    message.paymentComplete !== undefined && (obj.paymentComplete = message.paymentComplete);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OrdersUnderWatch>, I>>(object: I): OrdersUnderWatch {
    const message = createBaseOrdersUnderWatch();
    message.index = object.index ?? "";
    message.nknAddress = object.nknAddress ?? "";
    message.walletPrivateKey = object.walletPrivateKey ?? "";
    message.walletPublicKey = object.walletPublicKey ?? "";
    message.shippingAddress = object.shippingAddress ?? "";
    message.refundAddress = object.refundAddress ?? "";
    message.amount = object.amount ?? "";
    message.chain = object.chain ?? "";
    message.paymentComplete = object.paymentComplete ?? false;
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
