/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "teapartycrypto.partychain.party";

export interface MsgSubmitSell {
  creator: string;
  tradeAsset: string;
  price: string;
  currency: string;
  amount: string;
  sellerShippingAddr: string;
  sellerNknAddr: string;
  refundAddr: string;
}

export interface MsgSubmitSellResponse {
}

export interface MsgBuy {
  creator: string;
  txID: string;
  buyerShippingAddress: string;
  buyerNKNAddress: string;
  refundAddress: string;
}

export interface MsgBuyResponse {
}

export interface MsgCancel {
  creator: string;
}

export interface MsgCancelResponse {
}

export interface MsgAccountWatchOutcome {
  creator: string;
  txID: string;
  buyer: boolean;
  paymentOutcome: string;
}

export interface MsgAccountWatchOutcomeResponse {
}

export interface MsgAccountWatchFailure {
  creator: string;
  txID: string;
}

export interface MsgAccountWatchFailureResponse {
}

function createBaseMsgSubmitSell(): MsgSubmitSell {
  return {
    creator: "",
    tradeAsset: "",
    price: "",
    currency: "",
    amount: "",
    sellerShippingAddr: "",
    sellerNknAddr: "",
    refundAddr: "",
  };
}

export const MsgSubmitSell = {
  encode(message: MsgSubmitSell, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitSell {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitSell();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
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

  fromJSON(object: any): MsgSubmitSell {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      tradeAsset: isSet(object.tradeAsset) ? String(object.tradeAsset) : "",
      price: isSet(object.price) ? String(object.price) : "",
      currency: isSet(object.currency) ? String(object.currency) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      sellerShippingAddr: isSet(object.sellerShippingAddr) ? String(object.sellerShippingAddr) : "",
      sellerNknAddr: isSet(object.sellerNknAddr) ? String(object.sellerNknAddr) : "",
      refundAddr: isSet(object.refundAddr) ? String(object.refundAddr) : "",
    };
  },

  toJSON(message: MsgSubmitSell): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tradeAsset !== undefined && (obj.tradeAsset = message.tradeAsset);
    message.price !== undefined && (obj.price = message.price);
    message.currency !== undefined && (obj.currency = message.currency);
    message.amount !== undefined && (obj.amount = message.amount);
    message.sellerShippingAddr !== undefined && (obj.sellerShippingAddr = message.sellerShippingAddr);
    message.sellerNknAddr !== undefined && (obj.sellerNknAddr = message.sellerNknAddr);
    message.refundAddr !== undefined && (obj.refundAddr = message.refundAddr);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitSell>, I>>(object: I): MsgSubmitSell {
    const message = createBaseMsgSubmitSell();
    message.creator = object.creator ?? "";
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

function createBaseMsgSubmitSellResponse(): MsgSubmitSellResponse {
  return {};
}

export const MsgSubmitSellResponse = {
  encode(_: MsgSubmitSellResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitSellResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitSellResponse();
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

  fromJSON(_: any): MsgSubmitSellResponse {
    return {};
  },

  toJSON(_: MsgSubmitSellResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitSellResponse>, I>>(_: I): MsgSubmitSellResponse {
    const message = createBaseMsgSubmitSellResponse();
    return message;
  },
};

function createBaseMsgBuy(): MsgBuy {
  return { creator: "", txID: "", buyerShippingAddress: "", buyerNKNAddress: "", refundAddress: "" };
}

export const MsgBuy = {
  encode(message: MsgBuy, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.txID !== "") {
      writer.uint32(18).string(message.txID);
    }
    if (message.buyerShippingAddress !== "") {
      writer.uint32(26).string(message.buyerShippingAddress);
    }
    if (message.buyerNKNAddress !== "") {
      writer.uint32(34).string(message.buyerNKNAddress);
    }
    if (message.refundAddress !== "") {
      writer.uint32(42).string(message.refundAddress);
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
          message.txID = reader.string();
          break;
        case 3:
          message.buyerShippingAddress = reader.string();
          break;
        case 4:
          message.buyerNKNAddress = reader.string();
          break;
        case 5:
          message.refundAddress = reader.string();
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
      txID: isSet(object.txID) ? String(object.txID) : "",
      buyerShippingAddress: isSet(object.buyerShippingAddress) ? String(object.buyerShippingAddress) : "",
      buyerNKNAddress: isSet(object.buyerNKNAddress) ? String(object.buyerNKNAddress) : "",
      refundAddress: isSet(object.refundAddress) ? String(object.refundAddress) : "",
    };
  },

  toJSON(message: MsgBuy): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.txID !== undefined && (obj.txID = message.txID);
    message.buyerShippingAddress !== undefined && (obj.buyerShippingAddress = message.buyerShippingAddress);
    message.buyerNKNAddress !== undefined && (obj.buyerNKNAddress = message.buyerNKNAddress);
    message.refundAddress !== undefined && (obj.refundAddress = message.refundAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBuy>, I>>(object: I): MsgBuy {
    const message = createBaseMsgBuy();
    message.creator = object.creator ?? "";
    message.txID = object.txID ?? "";
    message.buyerShippingAddress = object.buyerShippingAddress ?? "";
    message.buyerNKNAddress = object.buyerNKNAddress ?? "";
    message.refundAddress = object.refundAddress ?? "";
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

function createBaseMsgCancel(): MsgCancel {
  return { creator: "" };
}

export const MsgCancel = {
  encode(message: MsgCancel, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
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
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCancel {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgCancel): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancel>, I>>(object: I): MsgCancel {
    const message = createBaseMsgCancel();
    message.creator = object.creator ?? "";
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

function createBaseMsgAccountWatchOutcome(): MsgAccountWatchOutcome {
  return { creator: "", txID: "", buyer: false, paymentOutcome: "" };
}

export const MsgAccountWatchOutcome = {
  encode(message: MsgAccountWatchOutcome, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.txID !== "") {
      writer.uint32(18).string(message.txID);
    }
    if (message.buyer === true) {
      writer.uint32(24).bool(message.buyer);
    }
    if (message.paymentOutcome !== "") {
      writer.uint32(34).string(message.paymentOutcome);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAccountWatchOutcome {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAccountWatchOutcome();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.txID = reader.string();
          break;
        case 3:
          message.buyer = reader.bool();
          break;
        case 4:
          message.paymentOutcome = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAccountWatchOutcome {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      txID: isSet(object.txID) ? String(object.txID) : "",
      buyer: isSet(object.buyer) ? Boolean(object.buyer) : false,
      paymentOutcome: isSet(object.paymentOutcome) ? String(object.paymentOutcome) : "",
    };
  },

  toJSON(message: MsgAccountWatchOutcome): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.txID !== undefined && (obj.txID = message.txID);
    message.buyer !== undefined && (obj.buyer = message.buyer);
    message.paymentOutcome !== undefined && (obj.paymentOutcome = message.paymentOutcome);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAccountWatchOutcome>, I>>(object: I): MsgAccountWatchOutcome {
    const message = createBaseMsgAccountWatchOutcome();
    message.creator = object.creator ?? "";
    message.txID = object.txID ?? "";
    message.buyer = object.buyer ?? false;
    message.paymentOutcome = object.paymentOutcome ?? "";
    return message;
  },
};

function createBaseMsgAccountWatchOutcomeResponse(): MsgAccountWatchOutcomeResponse {
  return {};
}

export const MsgAccountWatchOutcomeResponse = {
  encode(_: MsgAccountWatchOutcomeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAccountWatchOutcomeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAccountWatchOutcomeResponse();
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

  fromJSON(_: any): MsgAccountWatchOutcomeResponse {
    return {};
  },

  toJSON(_: MsgAccountWatchOutcomeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAccountWatchOutcomeResponse>, I>>(_: I): MsgAccountWatchOutcomeResponse {
    const message = createBaseMsgAccountWatchOutcomeResponse();
    return message;
  },
};

function createBaseMsgAccountWatchFailure(): MsgAccountWatchFailure {
  return { creator: "", txID: "" };
}

export const MsgAccountWatchFailure = {
  encode(message: MsgAccountWatchFailure, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.txID !== "") {
      writer.uint32(18).string(message.txID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAccountWatchFailure {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAccountWatchFailure();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.txID = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAccountWatchFailure {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      txID: isSet(object.txID) ? String(object.txID) : "",
    };
  },

  toJSON(message: MsgAccountWatchFailure): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.txID !== undefined && (obj.txID = message.txID);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAccountWatchFailure>, I>>(object: I): MsgAccountWatchFailure {
    const message = createBaseMsgAccountWatchFailure();
    message.creator = object.creator ?? "";
    message.txID = object.txID ?? "";
    return message;
  },
};

function createBaseMsgAccountWatchFailureResponse(): MsgAccountWatchFailureResponse {
  return {};
}

export const MsgAccountWatchFailureResponse = {
  encode(_: MsgAccountWatchFailureResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAccountWatchFailureResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAccountWatchFailureResponse();
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

  fromJSON(_: any): MsgAccountWatchFailureResponse {
    return {};
  },

  toJSON(_: MsgAccountWatchFailureResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAccountWatchFailureResponse>, I>>(_: I): MsgAccountWatchFailureResponse {
    const message = createBaseMsgAccountWatchFailureResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  SubmitSell(request: MsgSubmitSell): Promise<MsgSubmitSellResponse>;
  Buy(request: MsgBuy): Promise<MsgBuyResponse>;
  Cancel(request: MsgCancel): Promise<MsgCancelResponse>;
  AccountWatchOutcome(request: MsgAccountWatchOutcome): Promise<MsgAccountWatchOutcomeResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AccountWatchFailure(request: MsgAccountWatchFailure): Promise<MsgAccountWatchFailureResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.SubmitSell = this.SubmitSell.bind(this);
    this.Buy = this.Buy.bind(this);
    this.Cancel = this.Cancel.bind(this);
    this.AccountWatchOutcome = this.AccountWatchOutcome.bind(this);
    this.AccountWatchFailure = this.AccountWatchFailure.bind(this);
  }
  SubmitSell(request: MsgSubmitSell): Promise<MsgSubmitSellResponse> {
    const data = MsgSubmitSell.encode(request).finish();
    const promise = this.rpc.request("teapartycrypto.partychain.party.Msg", "SubmitSell", data);
    return promise.then((data) => MsgSubmitSellResponse.decode(new _m0.Reader(data)));
  }

  Buy(request: MsgBuy): Promise<MsgBuyResponse> {
    const data = MsgBuy.encode(request).finish();
    const promise = this.rpc.request("teapartycrypto.partychain.party.Msg", "Buy", data);
    return promise.then((data) => MsgBuyResponse.decode(new _m0.Reader(data)));
  }

  Cancel(request: MsgCancel): Promise<MsgCancelResponse> {
    const data = MsgCancel.encode(request).finish();
    const promise = this.rpc.request("teapartycrypto.partychain.party.Msg", "Cancel", data);
    return promise.then((data) => MsgCancelResponse.decode(new _m0.Reader(data)));
  }

  AccountWatchOutcome(request: MsgAccountWatchOutcome): Promise<MsgAccountWatchOutcomeResponse> {
    const data = MsgAccountWatchOutcome.encode(request).finish();
    const promise = this.rpc.request("teapartycrypto.partychain.party.Msg", "AccountWatchOutcome", data);
    return promise.then((data) => MsgAccountWatchOutcomeResponse.decode(new _m0.Reader(data)));
  }

  AccountWatchFailure(request: MsgAccountWatchFailure): Promise<MsgAccountWatchFailureResponse> {
    const data = MsgAccountWatchFailure.encode(request).finish();
    const promise = this.rpc.request("teapartycrypto.partychain.party.Msg", "AccountWatchFailure", data);
    return promise.then((data) => MsgAccountWatchFailureResponse.decode(new _m0.Reader(data)));
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
