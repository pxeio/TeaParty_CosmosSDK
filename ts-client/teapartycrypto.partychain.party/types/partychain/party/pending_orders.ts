/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "teapartycrypto.partychain.party";

export interface PendingOrders {
  index: string;
  buyerEscrowWalletPublicKey: string;
  buyerEscrowWalletPrivateKey: string;
  sellerEscrowWalletPublicKey: string;
  sellerEscrowWalletPrivateKey: string;
  sellerPaymentComplete: boolean;
  sellerPaymentCompleteBlockHeight: number;
  buyerPaymentComplete: boolean;
  buyerPaymentCompleteBlockHeight: number;
  amount: string;
  buyerShippingAddress: string;
  buyerRefundAddress: string;
  buyerNKNAddress: string;
  sellerRefundAddress: string;
  sellerShippingAddress: string;
  sellerNKNAddress: string;
  tradeAsset: string;
  currency: string;
  price: string;
  blockHeight: number;
}

function createBasePendingOrders(): PendingOrders {
  return {
    index: "",
    buyerEscrowWalletPublicKey: "",
    buyerEscrowWalletPrivateKey: "",
    sellerEscrowWalletPublicKey: "",
    sellerEscrowWalletPrivateKey: "",
    sellerPaymentComplete: false,
    sellerPaymentCompleteBlockHeight: 0,
    buyerPaymentComplete: false,
    buyerPaymentCompleteBlockHeight: 0,
    amount: "",
    buyerShippingAddress: "",
    buyerRefundAddress: "",
    buyerNKNAddress: "",
    sellerRefundAddress: "",
    sellerShippingAddress: "",
    sellerNKNAddress: "",
    tradeAsset: "",
    currency: "",
    price: "",
    blockHeight: 0,
  };
}

export const PendingOrders = {
  encode(message: PendingOrders, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.buyerEscrowWalletPublicKey !== "") {
      writer.uint32(18).string(message.buyerEscrowWalletPublicKey);
    }
    if (message.buyerEscrowWalletPrivateKey !== "") {
      writer.uint32(26).string(message.buyerEscrowWalletPrivateKey);
    }
    if (message.sellerEscrowWalletPublicKey !== "") {
      writer.uint32(34).string(message.sellerEscrowWalletPublicKey);
    }
    if (message.sellerEscrowWalletPrivateKey !== "") {
      writer.uint32(42).string(message.sellerEscrowWalletPrivateKey);
    }
    if (message.sellerPaymentComplete === true) {
      writer.uint32(48).bool(message.sellerPaymentComplete);
    }
    if (message.sellerPaymentCompleteBlockHeight !== 0) {
      writer.uint32(56).int32(message.sellerPaymentCompleteBlockHeight);
    }
    if (message.buyerPaymentComplete === true) {
      writer.uint32(64).bool(message.buyerPaymentComplete);
    }
    if (message.buyerPaymentCompleteBlockHeight !== 0) {
      writer.uint32(72).int32(message.buyerPaymentCompleteBlockHeight);
    }
    if (message.amount !== "") {
      writer.uint32(82).string(message.amount);
    }
    if (message.buyerShippingAddress !== "") {
      writer.uint32(90).string(message.buyerShippingAddress);
    }
    if (message.buyerRefundAddress !== "") {
      writer.uint32(98).string(message.buyerRefundAddress);
    }
    if (message.buyerNKNAddress !== "") {
      writer.uint32(106).string(message.buyerNKNAddress);
    }
    if (message.sellerRefundAddress !== "") {
      writer.uint32(114).string(message.sellerRefundAddress);
    }
    if (message.sellerShippingAddress !== "") {
      writer.uint32(122).string(message.sellerShippingAddress);
    }
    if (message.sellerNKNAddress !== "") {
      writer.uint32(130).string(message.sellerNKNAddress);
    }
    if (message.tradeAsset !== "") {
      writer.uint32(138).string(message.tradeAsset);
    }
    if (message.currency !== "") {
      writer.uint32(146).string(message.currency);
    }
    if (message.price !== "") {
      writer.uint32(154).string(message.price);
    }
    if (message.blockHeight !== 0) {
      writer.uint32(160).int32(message.blockHeight);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PendingOrders {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePendingOrders();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.buyerEscrowWalletPublicKey = reader.string();
          break;
        case 3:
          message.buyerEscrowWalletPrivateKey = reader.string();
          break;
        case 4:
          message.sellerEscrowWalletPublicKey = reader.string();
          break;
        case 5:
          message.sellerEscrowWalletPrivateKey = reader.string();
          break;
        case 6:
          message.sellerPaymentComplete = reader.bool();
          break;
        case 7:
          message.sellerPaymentCompleteBlockHeight = reader.int32();
          break;
        case 8:
          message.buyerPaymentComplete = reader.bool();
          break;
        case 9:
          message.buyerPaymentCompleteBlockHeight = reader.int32();
          break;
        case 10:
          message.amount = reader.string();
          break;
        case 11:
          message.buyerShippingAddress = reader.string();
          break;
        case 12:
          message.buyerRefundAddress = reader.string();
          break;
        case 13:
          message.buyerNKNAddress = reader.string();
          break;
        case 14:
          message.sellerRefundAddress = reader.string();
          break;
        case 15:
          message.sellerShippingAddress = reader.string();
          break;
        case 16:
          message.sellerNKNAddress = reader.string();
          break;
        case 17:
          message.tradeAsset = reader.string();
          break;
        case 18:
          message.currency = reader.string();
          break;
        case 19:
          message.price = reader.string();
          break;
        case 20:
          message.blockHeight = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PendingOrders {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      buyerEscrowWalletPublicKey: isSet(object.buyerEscrowWalletPublicKey)
        ? String(object.buyerEscrowWalletPublicKey)
        : "",
      buyerEscrowWalletPrivateKey: isSet(object.buyerEscrowWalletPrivateKey)
        ? String(object.buyerEscrowWalletPrivateKey)
        : "",
      sellerEscrowWalletPublicKey: isSet(object.sellerEscrowWalletPublicKey)
        ? String(object.sellerEscrowWalletPublicKey)
        : "",
      sellerEscrowWalletPrivateKey: isSet(object.sellerEscrowWalletPrivateKey)
        ? String(object.sellerEscrowWalletPrivateKey)
        : "",
      sellerPaymentComplete: isSet(object.sellerPaymentComplete) ? Boolean(object.sellerPaymentComplete) : false,
      sellerPaymentCompleteBlockHeight: isSet(object.sellerPaymentCompleteBlockHeight)
        ? Number(object.sellerPaymentCompleteBlockHeight)
        : 0,
      buyerPaymentComplete: isSet(object.buyerPaymentComplete) ? Boolean(object.buyerPaymentComplete) : false,
      buyerPaymentCompleteBlockHeight: isSet(object.buyerPaymentCompleteBlockHeight)
        ? Number(object.buyerPaymentCompleteBlockHeight)
        : 0,
      amount: isSet(object.amount) ? String(object.amount) : "",
      buyerShippingAddress: isSet(object.buyerShippingAddress) ? String(object.buyerShippingAddress) : "",
      buyerRefundAddress: isSet(object.buyerRefundAddress) ? String(object.buyerRefundAddress) : "",
      buyerNKNAddress: isSet(object.buyerNKNAddress) ? String(object.buyerNKNAddress) : "",
      sellerRefundAddress: isSet(object.sellerRefundAddress) ? String(object.sellerRefundAddress) : "",
      sellerShippingAddress: isSet(object.sellerShippingAddress) ? String(object.sellerShippingAddress) : "",
      sellerNKNAddress: isSet(object.sellerNKNAddress) ? String(object.sellerNKNAddress) : "",
      tradeAsset: isSet(object.tradeAsset) ? String(object.tradeAsset) : "",
      currency: isSet(object.currency) ? String(object.currency) : "",
      price: isSet(object.price) ? String(object.price) : "",
      blockHeight: isSet(object.blockHeight) ? Number(object.blockHeight) : 0,
    };
  },

  toJSON(message: PendingOrders): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.buyerEscrowWalletPublicKey !== undefined
      && (obj.buyerEscrowWalletPublicKey = message.buyerEscrowWalletPublicKey);
    message.buyerEscrowWalletPrivateKey !== undefined
      && (obj.buyerEscrowWalletPrivateKey = message.buyerEscrowWalletPrivateKey);
    message.sellerEscrowWalletPublicKey !== undefined
      && (obj.sellerEscrowWalletPublicKey = message.sellerEscrowWalletPublicKey);
    message.sellerEscrowWalletPrivateKey !== undefined
      && (obj.sellerEscrowWalletPrivateKey = message.sellerEscrowWalletPrivateKey);
    message.sellerPaymentComplete !== undefined && (obj.sellerPaymentComplete = message.sellerPaymentComplete);
    message.sellerPaymentCompleteBlockHeight !== undefined
      && (obj.sellerPaymentCompleteBlockHeight = Math.round(message.sellerPaymentCompleteBlockHeight));
    message.buyerPaymentComplete !== undefined && (obj.buyerPaymentComplete = message.buyerPaymentComplete);
    message.buyerPaymentCompleteBlockHeight !== undefined
      && (obj.buyerPaymentCompleteBlockHeight = Math.round(message.buyerPaymentCompleteBlockHeight));
    message.amount !== undefined && (obj.amount = message.amount);
    message.buyerShippingAddress !== undefined && (obj.buyerShippingAddress = message.buyerShippingAddress);
    message.buyerRefundAddress !== undefined && (obj.buyerRefundAddress = message.buyerRefundAddress);
    message.buyerNKNAddress !== undefined && (obj.buyerNKNAddress = message.buyerNKNAddress);
    message.sellerRefundAddress !== undefined && (obj.sellerRefundAddress = message.sellerRefundAddress);
    message.sellerShippingAddress !== undefined && (obj.sellerShippingAddress = message.sellerShippingAddress);
    message.sellerNKNAddress !== undefined && (obj.sellerNKNAddress = message.sellerNKNAddress);
    message.tradeAsset !== undefined && (obj.tradeAsset = message.tradeAsset);
    message.currency !== undefined && (obj.currency = message.currency);
    message.price !== undefined && (obj.price = message.price);
    message.blockHeight !== undefined && (obj.blockHeight = Math.round(message.blockHeight));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PendingOrders>, I>>(object: I): PendingOrders {
    const message = createBasePendingOrders();
    message.index = object.index ?? "";
    message.buyerEscrowWalletPublicKey = object.buyerEscrowWalletPublicKey ?? "";
    message.buyerEscrowWalletPrivateKey = object.buyerEscrowWalletPrivateKey ?? "";
    message.sellerEscrowWalletPublicKey = object.sellerEscrowWalletPublicKey ?? "";
    message.sellerEscrowWalletPrivateKey = object.sellerEscrowWalletPrivateKey ?? "";
    message.sellerPaymentComplete = object.sellerPaymentComplete ?? false;
    message.sellerPaymentCompleteBlockHeight = object.sellerPaymentCompleteBlockHeight ?? 0;
    message.buyerPaymentComplete = object.buyerPaymentComplete ?? false;
    message.buyerPaymentCompleteBlockHeight = object.buyerPaymentCompleteBlockHeight ?? 0;
    message.amount = object.amount ?? "";
    message.buyerShippingAddress = object.buyerShippingAddress ?? "";
    message.buyerRefundAddress = object.buyerRefundAddress ?? "";
    message.buyerNKNAddress = object.buyerNKNAddress ?? "";
    message.sellerRefundAddress = object.sellerRefundAddress ?? "";
    message.sellerShippingAddress = object.sellerShippingAddress ?? "";
    message.sellerNKNAddress = object.sellerNKNAddress ?? "";
    message.tradeAsset = object.tradeAsset ?? "";
    message.currency = object.currency ?? "";
    message.price = object.price ?? "";
    message.blockHeight = object.blockHeight ?? 0;
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
