/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export type PartyMsgAccountWatchFailureResponse = object;

export type PartyMsgAccountWatchOutcomeResponse = object;

export type PartyMsgBuyResponse = object;

export type PartyMsgCancelResponse = object;

export type PartyMsgSubmitSellResponse = object;

export interface PartyOrdersAwaitingFinalizer {
  index?: string;
  nknAddress?: string;
  walletPrivateKey?: string;
  walletPublicKey?: string;
  shippingAddress?: string;
  refundAddress?: string;
  amount?: string;
  chain?: string;
}

/**
 * Params defines the parameters for the module.
 */
export type PartyParams = object;

export interface PartyPendingOrders {
  index?: string;
  buyerEscrowWalletPublicKey?: string;
  buyerEscrowWalletPrivateKey?: string;
  sellerEscrowWalletPublicKey?: string;
  sellerEscrowWalletPrivateKey?: string;
  sellerPaymentComplete?: boolean;

  /** @format int32 */
  sellerPaymentCompleteBlockHeight?: number;
  buyerPaymentComplete?: boolean;

  /** @format int32 */
  buyerPaymentCompleteBlockHeight?: number;
  amount?: string;
  buyerShippingAddress?: string;
  buyerRefundAddress?: string;
  buyerNKNAddress?: string;
  sellerRefundAddress?: string;
  sellerShippingAddress?: string;
  sellerNKNAddress?: string;
  tradeAsset?: string;
  currency?: string;
  price?: string;

  /** @format int32 */
  blockHeight?: number;
}

export interface PartyQueryAllOrdersAwaitingFinalizerResponse {
  ordersAwaitingFinalizer?: PartyOrdersAwaitingFinalizer[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface PartyQueryAllPendingOrdersResponse {
  pendingOrders?: PartyPendingOrders[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface PartyQueryAllTradeOrdersResponse {
  tradeOrders?: PartyTradeOrders[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface PartyQueryGetOrdersAwaitingFinalizerResponse {
  ordersAwaitingFinalizer?: PartyOrdersAwaitingFinalizer;
}

export interface PartyQueryGetPendingOrdersResponse {
  pendingOrders?: PartyPendingOrders;
}

export interface PartyQueryGetTradeOrdersResponse {
  tradeOrders?: PartyTradeOrders;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface PartyQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: PartyParams;
}

export interface PartyTradeOrders {
  index?: string;
  tradeAsset?: string;
  price?: string;
  currency?: string;
  amount?: string;
  sellerShippingAddr?: string;
  sellerNknAddr?: string;
  refundAddr?: string;
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /**
   * next_key is the key to be passed to PageRequest.key to
   * query the next page most efficiently. It will be empty if
   * there are no more results.
   * @format byte
   */
  next_key?: string;

  /**
   * total is total number of results available if PageRequest.count_total
   * was set, its value is undefined otherwise
   * @format uint64
   */
  total?: string;
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title partychain/party/genesis.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryOrdersAwaitingFinalizerAll
   * @summary Queries a list of OrdersAwaitingFinalizer items.
   * @request GET:/TeaPartyCrypto/partychain/party/orders_awaiting_finalizer
   */
  queryOrdersAwaitingFinalizerAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PartyQueryAllOrdersAwaitingFinalizerResponse, RpcStatus>({
      path: `/TeaPartyCrypto/partychain/party/orders_awaiting_finalizer`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOrdersAwaitingFinalizer
   * @summary Queries a OrdersAwaitingFinalizer by index.
   * @request GET:/TeaPartyCrypto/partychain/party/orders_awaiting_finalizer/{index}
   */
  queryOrdersAwaitingFinalizer = (index: string, params: RequestParams = {}) =>
    this.request<PartyQueryGetOrdersAwaitingFinalizerResponse, RpcStatus>({
      path: `/TeaPartyCrypto/partychain/party/orders_awaiting_finalizer/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/TeaPartyCrypto/partychain/party/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<PartyQueryParamsResponse, RpcStatus>({
      path: `/TeaPartyCrypto/partychain/party/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPendingOrdersAll
   * @summary Queries a list of PendingOrders items.
   * @request GET:/TeaPartyCrypto/partychain/party/pending_orders
   */
  queryPendingOrdersAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PartyQueryAllPendingOrdersResponse, RpcStatus>({
      path: `/TeaPartyCrypto/partychain/party/pending_orders`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPendingOrders
   * @summary Queries a PendingOrders by index.
   * @request GET:/TeaPartyCrypto/partychain/party/pending_orders/{index}
   */
  queryPendingOrders = (index: string, params: RequestParams = {}) =>
    this.request<PartyQueryGetPendingOrdersResponse, RpcStatus>({
      path: `/TeaPartyCrypto/partychain/party/pending_orders/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTradeOrdersAll
   * @summary Queries a list of TradeOrders items.
   * @request GET:/TeaPartyCrypto/partychain/party/trade_orders
   */
  queryTradeOrdersAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PartyQueryAllTradeOrdersResponse, RpcStatus>({
      path: `/TeaPartyCrypto/partychain/party/trade_orders`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTradeOrders
   * @summary Queries a TradeOrders by index.
   * @request GET:/TeaPartyCrypto/partychain/party/trade_orders/{index}
   */
  queryTradeOrders = (index: string, params: RequestParams = {}) =>
    this.request<PartyQueryGetTradeOrdersResponse, RpcStatus>({
      path: `/TeaPartyCrypto/partychain/party/trade_orders/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
