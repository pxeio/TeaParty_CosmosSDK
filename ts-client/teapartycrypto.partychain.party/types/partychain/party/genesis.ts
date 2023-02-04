/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { FinalizingOrders } from "./finalizing_orders";
import { OrdersAwaitingFinalizer } from "./orders_awaiting_finalizer";
import { OrdersUnderWatch } from "./orders_under_watch";
import { Params } from "./params";
import { PendingOrders } from "./pending_orders";
import { TradeOrders } from "./trade_orders";

export const protobufPackage = "teapartycrypto.partychain.party";

/** GenesisState defines the party module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  tradeOrdersList: TradeOrders[];
  pendingOrdersList: PendingOrders[];
  ordersAwaitingFinalizerList: OrdersAwaitingFinalizer[];
  ordersUnderWatchList: OrdersUnderWatch[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  finalizingOrdersList: FinalizingOrders[];
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    tradeOrdersList: [],
    pendingOrdersList: [],
    ordersAwaitingFinalizerList: [],
    ordersUnderWatchList: [],
    finalizingOrdersList: [],
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.tradeOrdersList) {
      TradeOrders.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.pendingOrdersList) {
      PendingOrders.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.ordersAwaitingFinalizerList) {
      OrdersAwaitingFinalizer.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.ordersUnderWatchList) {
      OrdersUnderWatch.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.finalizingOrdersList) {
      FinalizingOrders.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.tradeOrdersList.push(TradeOrders.decode(reader, reader.uint32()));
          break;
        case 3:
          message.pendingOrdersList.push(PendingOrders.decode(reader, reader.uint32()));
          break;
        case 4:
          message.ordersAwaitingFinalizerList.push(OrdersAwaitingFinalizer.decode(reader, reader.uint32()));
          break;
        case 5:
          message.ordersUnderWatchList.push(OrdersUnderWatch.decode(reader, reader.uint32()));
          break;
        case 6:
          message.finalizingOrdersList.push(FinalizingOrders.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      tradeOrdersList: Array.isArray(object?.tradeOrdersList)
        ? object.tradeOrdersList.map((e: any) => TradeOrders.fromJSON(e))
        : [],
      pendingOrdersList: Array.isArray(object?.pendingOrdersList)
        ? object.pendingOrdersList.map((e: any) => PendingOrders.fromJSON(e))
        : [],
      ordersAwaitingFinalizerList: Array.isArray(object?.ordersAwaitingFinalizerList)
        ? object.ordersAwaitingFinalizerList.map((e: any) => OrdersAwaitingFinalizer.fromJSON(e))
        : [],
      ordersUnderWatchList: Array.isArray(object?.ordersUnderWatchList)
        ? object.ordersUnderWatchList.map((e: any) => OrdersUnderWatch.fromJSON(e))
        : [],
      finalizingOrdersList: Array.isArray(object?.finalizingOrdersList)
        ? object.finalizingOrdersList.map((e: any) => FinalizingOrders.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.tradeOrdersList) {
      obj.tradeOrdersList = message.tradeOrdersList.map((e) => e ? TradeOrders.toJSON(e) : undefined);
    } else {
      obj.tradeOrdersList = [];
    }
    if (message.pendingOrdersList) {
      obj.pendingOrdersList = message.pendingOrdersList.map((e) => e ? PendingOrders.toJSON(e) : undefined);
    } else {
      obj.pendingOrdersList = [];
    }
    if (message.ordersAwaitingFinalizerList) {
      obj.ordersAwaitingFinalizerList = message.ordersAwaitingFinalizerList.map((e) =>
        e ? OrdersAwaitingFinalizer.toJSON(e) : undefined
      );
    } else {
      obj.ordersAwaitingFinalizerList = [];
    }
    if (message.ordersUnderWatchList) {
      obj.ordersUnderWatchList = message.ordersUnderWatchList.map((e) => e ? OrdersUnderWatch.toJSON(e) : undefined);
    } else {
      obj.ordersUnderWatchList = [];
    }
    if (message.finalizingOrdersList) {
      obj.finalizingOrdersList = message.finalizingOrdersList.map((e) => e ? FinalizingOrders.toJSON(e) : undefined);
    } else {
      obj.finalizingOrdersList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.tradeOrdersList = object.tradeOrdersList?.map((e) => TradeOrders.fromPartial(e)) || [];
    message.pendingOrdersList = object.pendingOrdersList?.map((e) => PendingOrders.fromPartial(e)) || [];
    message.ordersAwaitingFinalizerList =
      object.ordersAwaitingFinalizerList?.map((e) => OrdersAwaitingFinalizer.fromPartial(e)) || [];
    message.ordersUnderWatchList = object.ordersUnderWatchList?.map((e) => OrdersUnderWatch.fromPartial(e)) || [];
    message.finalizingOrdersList = object.finalizingOrdersList?.map((e) => FinalizingOrders.fromPartial(e)) || [];
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
