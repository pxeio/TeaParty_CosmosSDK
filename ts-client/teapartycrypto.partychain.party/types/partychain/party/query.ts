/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { OrdersAwaitingFinalizer } from "./orders_awaiting_finalizer";
import { OrdersUnderWatch } from "./orders_under_watch";
import { Params } from "./params";
import { PendingOrders } from "./pending_orders";
import { TradeOrders } from "./trade_orders";

export const protobufPackage = "teapartycrypto.partychain.party";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetTradeOrdersRequest {
  index: string;
}

export interface QueryGetTradeOrdersResponse {
  tradeOrders: TradeOrders | undefined;
}

export interface QueryAllTradeOrdersRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTradeOrdersResponse {
  tradeOrders: TradeOrders[];
  pagination: PageResponse | undefined;
}

export interface QueryGetPendingOrdersRequest {
  index: string;
}

export interface QueryGetPendingOrdersResponse {
  pendingOrders: PendingOrders | undefined;
}

export interface QueryAllPendingOrdersRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllPendingOrdersResponse {
  pendingOrders: PendingOrders[];
  pagination: PageResponse | undefined;
}

export interface QueryGetOrdersAwaitingFinalizerRequest {
  index: string;
}

export interface QueryGetOrdersAwaitingFinalizerResponse {
  ordersAwaitingFinalizer: OrdersAwaitingFinalizer | undefined;
}

export interface QueryAllOrdersAwaitingFinalizerRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllOrdersAwaitingFinalizerResponse {
  ordersAwaitingFinalizer: OrdersAwaitingFinalizer[];
  pagination: PageResponse | undefined;
}

export interface QueryGetOrdersUnderWatchRequest {
  index: string;
}

export interface QueryGetOrdersUnderWatchResponse {
  ordersUnderWatch: OrdersUnderWatch | undefined;
}

export interface QueryAllOrdersUnderWatchRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllOrdersUnderWatchResponse {
  ordersUnderWatch: OrdersUnderWatch[];
  pagination: PageResponse | undefined;
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

function createBaseQueryGetTradeOrdersRequest(): QueryGetTradeOrdersRequest {
  return { index: "" };
}

export const QueryGetTradeOrdersRequest = {
  encode(message: QueryGetTradeOrdersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetTradeOrdersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetTradeOrdersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTradeOrdersRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetTradeOrdersRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetTradeOrdersRequest>, I>>(object: I): QueryGetTradeOrdersRequest {
    const message = createBaseQueryGetTradeOrdersRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetTradeOrdersResponse(): QueryGetTradeOrdersResponse {
  return { tradeOrders: undefined };
}

export const QueryGetTradeOrdersResponse = {
  encode(message: QueryGetTradeOrdersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.tradeOrders !== undefined) {
      TradeOrders.encode(message.tradeOrders, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetTradeOrdersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetTradeOrdersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tradeOrders = TradeOrders.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTradeOrdersResponse {
    return { tradeOrders: isSet(object.tradeOrders) ? TradeOrders.fromJSON(object.tradeOrders) : undefined };
  },

  toJSON(message: QueryGetTradeOrdersResponse): unknown {
    const obj: any = {};
    message.tradeOrders !== undefined
      && (obj.tradeOrders = message.tradeOrders ? TradeOrders.toJSON(message.tradeOrders) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetTradeOrdersResponse>, I>>(object: I): QueryGetTradeOrdersResponse {
    const message = createBaseQueryGetTradeOrdersResponse();
    message.tradeOrders = (object.tradeOrders !== undefined && object.tradeOrders !== null)
      ? TradeOrders.fromPartial(object.tradeOrders)
      : undefined;
    return message;
  },
};

function createBaseQueryAllTradeOrdersRequest(): QueryAllTradeOrdersRequest {
  return { pagination: undefined };
}

export const QueryAllTradeOrdersRequest = {
  encode(message: QueryAllTradeOrdersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllTradeOrdersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllTradeOrdersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTradeOrdersRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllTradeOrdersRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllTradeOrdersRequest>, I>>(object: I): QueryAllTradeOrdersRequest {
    const message = createBaseQueryAllTradeOrdersRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllTradeOrdersResponse(): QueryAllTradeOrdersResponse {
  return { tradeOrders: [], pagination: undefined };
}

export const QueryAllTradeOrdersResponse = {
  encode(message: QueryAllTradeOrdersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.tradeOrders) {
      TradeOrders.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllTradeOrdersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllTradeOrdersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tradeOrders.push(TradeOrders.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTradeOrdersResponse {
    return {
      tradeOrders: Array.isArray(object?.tradeOrders)
        ? object.tradeOrders.map((e: any) => TradeOrders.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllTradeOrdersResponse): unknown {
    const obj: any = {};
    if (message.tradeOrders) {
      obj.tradeOrders = message.tradeOrders.map((e) => e ? TradeOrders.toJSON(e) : undefined);
    } else {
      obj.tradeOrders = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllTradeOrdersResponse>, I>>(object: I): QueryAllTradeOrdersResponse {
    const message = createBaseQueryAllTradeOrdersResponse();
    message.tradeOrders = object.tradeOrders?.map((e) => TradeOrders.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetPendingOrdersRequest(): QueryGetPendingOrdersRequest {
  return { index: "" };
}

export const QueryGetPendingOrdersRequest = {
  encode(message: QueryGetPendingOrdersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetPendingOrdersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetPendingOrdersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPendingOrdersRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetPendingOrdersRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetPendingOrdersRequest>, I>>(object: I): QueryGetPendingOrdersRequest {
    const message = createBaseQueryGetPendingOrdersRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetPendingOrdersResponse(): QueryGetPendingOrdersResponse {
  return { pendingOrders: undefined };
}

export const QueryGetPendingOrdersResponse = {
  encode(message: QueryGetPendingOrdersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pendingOrders !== undefined) {
      PendingOrders.encode(message.pendingOrders, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetPendingOrdersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetPendingOrdersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pendingOrders = PendingOrders.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPendingOrdersResponse {
    return { pendingOrders: isSet(object.pendingOrders) ? PendingOrders.fromJSON(object.pendingOrders) : undefined };
  },

  toJSON(message: QueryGetPendingOrdersResponse): unknown {
    const obj: any = {};
    message.pendingOrders !== undefined
      && (obj.pendingOrders = message.pendingOrders ? PendingOrders.toJSON(message.pendingOrders) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetPendingOrdersResponse>, I>>(
    object: I,
  ): QueryGetPendingOrdersResponse {
    const message = createBaseQueryGetPendingOrdersResponse();
    message.pendingOrders = (object.pendingOrders !== undefined && object.pendingOrders !== null)
      ? PendingOrders.fromPartial(object.pendingOrders)
      : undefined;
    return message;
  },
};

function createBaseQueryAllPendingOrdersRequest(): QueryAllPendingOrdersRequest {
  return { pagination: undefined };
}

export const QueryAllPendingOrdersRequest = {
  encode(message: QueryAllPendingOrdersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllPendingOrdersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllPendingOrdersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPendingOrdersRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllPendingOrdersRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllPendingOrdersRequest>, I>>(object: I): QueryAllPendingOrdersRequest {
    const message = createBaseQueryAllPendingOrdersRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllPendingOrdersResponse(): QueryAllPendingOrdersResponse {
  return { pendingOrders: [], pagination: undefined };
}

export const QueryAllPendingOrdersResponse = {
  encode(message: QueryAllPendingOrdersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.pendingOrders) {
      PendingOrders.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllPendingOrdersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllPendingOrdersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pendingOrders.push(PendingOrders.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPendingOrdersResponse {
    return {
      pendingOrders: Array.isArray(object?.pendingOrders)
        ? object.pendingOrders.map((e: any) => PendingOrders.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllPendingOrdersResponse): unknown {
    const obj: any = {};
    if (message.pendingOrders) {
      obj.pendingOrders = message.pendingOrders.map((e) => e ? PendingOrders.toJSON(e) : undefined);
    } else {
      obj.pendingOrders = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllPendingOrdersResponse>, I>>(
    object: I,
  ): QueryAllPendingOrdersResponse {
    const message = createBaseQueryAllPendingOrdersResponse();
    message.pendingOrders = object.pendingOrders?.map((e) => PendingOrders.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetOrdersAwaitingFinalizerRequest(): QueryGetOrdersAwaitingFinalizerRequest {
  return { index: "" };
}

export const QueryGetOrdersAwaitingFinalizerRequest = {
  encode(message: QueryGetOrdersAwaitingFinalizerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetOrdersAwaitingFinalizerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetOrdersAwaitingFinalizerRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOrdersAwaitingFinalizerRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetOrdersAwaitingFinalizerRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetOrdersAwaitingFinalizerRequest>, I>>(
    object: I,
  ): QueryGetOrdersAwaitingFinalizerRequest {
    const message = createBaseQueryGetOrdersAwaitingFinalizerRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetOrdersAwaitingFinalizerResponse(): QueryGetOrdersAwaitingFinalizerResponse {
  return { ordersAwaitingFinalizer: undefined };
}

export const QueryGetOrdersAwaitingFinalizerResponse = {
  encode(message: QueryGetOrdersAwaitingFinalizerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ordersAwaitingFinalizer !== undefined) {
      OrdersAwaitingFinalizer.encode(message.ordersAwaitingFinalizer, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetOrdersAwaitingFinalizerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetOrdersAwaitingFinalizerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ordersAwaitingFinalizer = OrdersAwaitingFinalizer.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOrdersAwaitingFinalizerResponse {
    return {
      ordersAwaitingFinalizer: isSet(object.ordersAwaitingFinalizer)
        ? OrdersAwaitingFinalizer.fromJSON(object.ordersAwaitingFinalizer)
        : undefined,
    };
  },

  toJSON(message: QueryGetOrdersAwaitingFinalizerResponse): unknown {
    const obj: any = {};
    message.ordersAwaitingFinalizer !== undefined && (obj.ordersAwaitingFinalizer = message.ordersAwaitingFinalizer
      ? OrdersAwaitingFinalizer.toJSON(message.ordersAwaitingFinalizer)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetOrdersAwaitingFinalizerResponse>, I>>(
    object: I,
  ): QueryGetOrdersAwaitingFinalizerResponse {
    const message = createBaseQueryGetOrdersAwaitingFinalizerResponse();
    message.ordersAwaitingFinalizer =
      (object.ordersAwaitingFinalizer !== undefined && object.ordersAwaitingFinalizer !== null)
        ? OrdersAwaitingFinalizer.fromPartial(object.ordersAwaitingFinalizer)
        : undefined;
    return message;
  },
};

function createBaseQueryAllOrdersAwaitingFinalizerRequest(): QueryAllOrdersAwaitingFinalizerRequest {
  return { pagination: undefined };
}

export const QueryAllOrdersAwaitingFinalizerRequest = {
  encode(message: QueryAllOrdersAwaitingFinalizerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllOrdersAwaitingFinalizerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllOrdersAwaitingFinalizerRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOrdersAwaitingFinalizerRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllOrdersAwaitingFinalizerRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllOrdersAwaitingFinalizerRequest>, I>>(
    object: I,
  ): QueryAllOrdersAwaitingFinalizerRequest {
    const message = createBaseQueryAllOrdersAwaitingFinalizerRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllOrdersAwaitingFinalizerResponse(): QueryAllOrdersAwaitingFinalizerResponse {
  return { ordersAwaitingFinalizer: [], pagination: undefined };
}

export const QueryAllOrdersAwaitingFinalizerResponse = {
  encode(message: QueryAllOrdersAwaitingFinalizerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.ordersAwaitingFinalizer) {
      OrdersAwaitingFinalizer.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllOrdersAwaitingFinalizerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllOrdersAwaitingFinalizerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ordersAwaitingFinalizer.push(OrdersAwaitingFinalizer.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOrdersAwaitingFinalizerResponse {
    return {
      ordersAwaitingFinalizer: Array.isArray(object?.ordersAwaitingFinalizer)
        ? object.ordersAwaitingFinalizer.map((e: any) => OrdersAwaitingFinalizer.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllOrdersAwaitingFinalizerResponse): unknown {
    const obj: any = {};
    if (message.ordersAwaitingFinalizer) {
      obj.ordersAwaitingFinalizer = message.ordersAwaitingFinalizer.map((e) =>
        e ? OrdersAwaitingFinalizer.toJSON(e) : undefined
      );
    } else {
      obj.ordersAwaitingFinalizer = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllOrdersAwaitingFinalizerResponse>, I>>(
    object: I,
  ): QueryAllOrdersAwaitingFinalizerResponse {
    const message = createBaseQueryAllOrdersAwaitingFinalizerResponse();
    message.ordersAwaitingFinalizer = object.ordersAwaitingFinalizer?.map((e) => OrdersAwaitingFinalizer.fromPartial(e))
      || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetOrdersUnderWatchRequest(): QueryGetOrdersUnderWatchRequest {
  return { index: "" };
}

export const QueryGetOrdersUnderWatchRequest = {
  encode(message: QueryGetOrdersUnderWatchRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetOrdersUnderWatchRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetOrdersUnderWatchRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOrdersUnderWatchRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetOrdersUnderWatchRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetOrdersUnderWatchRequest>, I>>(
    object: I,
  ): QueryGetOrdersUnderWatchRequest {
    const message = createBaseQueryGetOrdersUnderWatchRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetOrdersUnderWatchResponse(): QueryGetOrdersUnderWatchResponse {
  return { ordersUnderWatch: undefined };
}

export const QueryGetOrdersUnderWatchResponse = {
  encode(message: QueryGetOrdersUnderWatchResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ordersUnderWatch !== undefined) {
      OrdersUnderWatch.encode(message.ordersUnderWatch, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetOrdersUnderWatchResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetOrdersUnderWatchResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ordersUnderWatch = OrdersUnderWatch.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOrdersUnderWatchResponse {
    return {
      ordersUnderWatch: isSet(object.ordersUnderWatch) ? OrdersUnderWatch.fromJSON(object.ordersUnderWatch) : undefined,
    };
  },

  toJSON(message: QueryGetOrdersUnderWatchResponse): unknown {
    const obj: any = {};
    message.ordersUnderWatch !== undefined && (obj.ordersUnderWatch = message.ordersUnderWatch
      ? OrdersUnderWatch.toJSON(message.ordersUnderWatch)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetOrdersUnderWatchResponse>, I>>(
    object: I,
  ): QueryGetOrdersUnderWatchResponse {
    const message = createBaseQueryGetOrdersUnderWatchResponse();
    message.ordersUnderWatch = (object.ordersUnderWatch !== undefined && object.ordersUnderWatch !== null)
      ? OrdersUnderWatch.fromPartial(object.ordersUnderWatch)
      : undefined;
    return message;
  },
};

function createBaseQueryAllOrdersUnderWatchRequest(): QueryAllOrdersUnderWatchRequest {
  return { pagination: undefined };
}

export const QueryAllOrdersUnderWatchRequest = {
  encode(message: QueryAllOrdersUnderWatchRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllOrdersUnderWatchRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllOrdersUnderWatchRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOrdersUnderWatchRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllOrdersUnderWatchRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllOrdersUnderWatchRequest>, I>>(
    object: I,
  ): QueryAllOrdersUnderWatchRequest {
    const message = createBaseQueryAllOrdersUnderWatchRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllOrdersUnderWatchResponse(): QueryAllOrdersUnderWatchResponse {
  return { ordersUnderWatch: [], pagination: undefined };
}

export const QueryAllOrdersUnderWatchResponse = {
  encode(message: QueryAllOrdersUnderWatchResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.ordersUnderWatch) {
      OrdersUnderWatch.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllOrdersUnderWatchResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllOrdersUnderWatchResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ordersUnderWatch.push(OrdersUnderWatch.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOrdersUnderWatchResponse {
    return {
      ordersUnderWatch: Array.isArray(object?.ordersUnderWatch)
        ? object.ordersUnderWatch.map((e: any) => OrdersUnderWatch.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllOrdersUnderWatchResponse): unknown {
    const obj: any = {};
    if (message.ordersUnderWatch) {
      obj.ordersUnderWatch = message.ordersUnderWatch.map((e) => e ? OrdersUnderWatch.toJSON(e) : undefined);
    } else {
      obj.ordersUnderWatch = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllOrdersUnderWatchResponse>, I>>(
    object: I,
  ): QueryAllOrdersUnderWatchResponse {
    const message = createBaseQueryAllOrdersUnderWatchResponse();
    message.ordersUnderWatch = object.ordersUnderWatch?.map((e) => OrdersUnderWatch.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a TradeOrders by index. */
  TradeOrders(request: QueryGetTradeOrdersRequest): Promise<QueryGetTradeOrdersResponse>;
  /** Queries a list of TradeOrders items. */
  TradeOrdersAll(request: QueryAllTradeOrdersRequest): Promise<QueryAllTradeOrdersResponse>;
  /** Queries a PendingOrders by index. */
  PendingOrders(request: QueryGetPendingOrdersRequest): Promise<QueryGetPendingOrdersResponse>;
  /** Queries a list of PendingOrders items. */
  PendingOrdersAll(request: QueryAllPendingOrdersRequest): Promise<QueryAllPendingOrdersResponse>;
  /** Queries a OrdersAwaitingFinalizer by index. */
  OrdersAwaitingFinalizer(
    request: QueryGetOrdersAwaitingFinalizerRequest,
  ): Promise<QueryGetOrdersAwaitingFinalizerResponse>;
  /** Queries a list of OrdersAwaitingFinalizer items. */
  OrdersAwaitingFinalizerAll(
    request: QueryAllOrdersAwaitingFinalizerRequest,
  ): Promise<QueryAllOrdersAwaitingFinalizerResponse>;
  /** Queries a OrdersUnderWatch by index. */
  OrdersUnderWatch(request: QueryGetOrdersUnderWatchRequest): Promise<QueryGetOrdersUnderWatchResponse>;
  /** Queries a list of OrdersUnderWatch items. */
  OrdersUnderWatchAll(request: QueryAllOrdersUnderWatchRequest): Promise<QueryAllOrdersUnderWatchResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.TradeOrders = this.TradeOrders.bind(this);
    this.TradeOrdersAll = this.TradeOrdersAll.bind(this);
    this.PendingOrders = this.PendingOrders.bind(this);
    this.PendingOrdersAll = this.PendingOrdersAll.bind(this);
    this.OrdersAwaitingFinalizer = this.OrdersAwaitingFinalizer.bind(this);
    this.OrdersAwaitingFinalizerAll = this.OrdersAwaitingFinalizerAll.bind(this);
    this.OrdersUnderWatch = this.OrdersUnderWatch.bind(this);
    this.OrdersUnderWatchAll = this.OrdersUnderWatchAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("teapartycrypto.partychain.party.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  TradeOrders(request: QueryGetTradeOrdersRequest): Promise<QueryGetTradeOrdersResponse> {
    const data = QueryGetTradeOrdersRequest.encode(request).finish();
    const promise = this.rpc.request("teapartycrypto.partychain.party.Query", "TradeOrders", data);
    return promise.then((data) => QueryGetTradeOrdersResponse.decode(new _m0.Reader(data)));
  }

  TradeOrdersAll(request: QueryAllTradeOrdersRequest): Promise<QueryAllTradeOrdersResponse> {
    const data = QueryAllTradeOrdersRequest.encode(request).finish();
    const promise = this.rpc.request("teapartycrypto.partychain.party.Query", "TradeOrdersAll", data);
    return promise.then((data) => QueryAllTradeOrdersResponse.decode(new _m0.Reader(data)));
  }

  PendingOrders(request: QueryGetPendingOrdersRequest): Promise<QueryGetPendingOrdersResponse> {
    const data = QueryGetPendingOrdersRequest.encode(request).finish();
    const promise = this.rpc.request("teapartycrypto.partychain.party.Query", "PendingOrders", data);
    return promise.then((data) => QueryGetPendingOrdersResponse.decode(new _m0.Reader(data)));
  }

  PendingOrdersAll(request: QueryAllPendingOrdersRequest): Promise<QueryAllPendingOrdersResponse> {
    const data = QueryAllPendingOrdersRequest.encode(request).finish();
    const promise = this.rpc.request("teapartycrypto.partychain.party.Query", "PendingOrdersAll", data);
    return promise.then((data) => QueryAllPendingOrdersResponse.decode(new _m0.Reader(data)));
  }

  OrdersAwaitingFinalizer(
    request: QueryGetOrdersAwaitingFinalizerRequest,
  ): Promise<QueryGetOrdersAwaitingFinalizerResponse> {
    const data = QueryGetOrdersAwaitingFinalizerRequest.encode(request).finish();
    const promise = this.rpc.request("teapartycrypto.partychain.party.Query", "OrdersAwaitingFinalizer", data);
    return promise.then((data) => QueryGetOrdersAwaitingFinalizerResponse.decode(new _m0.Reader(data)));
  }

  OrdersAwaitingFinalizerAll(
    request: QueryAllOrdersAwaitingFinalizerRequest,
  ): Promise<QueryAllOrdersAwaitingFinalizerResponse> {
    const data = QueryAllOrdersAwaitingFinalizerRequest.encode(request).finish();
    const promise = this.rpc.request("teapartycrypto.partychain.party.Query", "OrdersAwaitingFinalizerAll", data);
    return promise.then((data) => QueryAllOrdersAwaitingFinalizerResponse.decode(new _m0.Reader(data)));
  }

  OrdersUnderWatch(request: QueryGetOrdersUnderWatchRequest): Promise<QueryGetOrdersUnderWatchResponse> {
    const data = QueryGetOrdersUnderWatchRequest.encode(request).finish();
    const promise = this.rpc.request("teapartycrypto.partychain.party.Query", "OrdersUnderWatch", data);
    return promise.then((data) => QueryGetOrdersUnderWatchResponse.decode(new _m0.Reader(data)));
  }

  OrdersUnderWatchAll(request: QueryAllOrdersUnderWatchRequest): Promise<QueryAllOrdersUnderWatchResponse> {
    const data = QueryAllOrdersUnderWatchRequest.encode(request).finish();
    const promise = this.rpc.request("teapartycrypto.partychain.party.Query", "OrdersUnderWatchAll", data);
    return promise.then((data) => QueryAllOrdersUnderWatchResponse.decode(new _m0.Reader(data)));
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
