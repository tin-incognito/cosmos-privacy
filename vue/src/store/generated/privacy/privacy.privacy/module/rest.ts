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

export interface MsgTransferPaymentInfo {
  payment_address?: string;

  /** @format uint64 */
  amount?: string;

  /** @format byte */
  info?: string;
}

export interface PrivacyCommitment {
  index?: string;
  creator?: string;
  is_confidential_asset?: boolean;

  /** @format byte */
  value?: string;
}

export interface PrivacyCommitmentIndex {
  index?: string;
  creator?: string;
}

export type PrivacyMsgAirdropResponse = object;

export type PrivacyMsgCreateCommitmentIndexResponse = object;

export type PrivacyMsgCreateCommitmentResponse = object;

export type PrivacyMsgCreateOnetimeAddressResponse = object;

export type PrivacyMsgCreateOutputCoinResponse = object;

export type PrivacyMsgCreateOutputCoinSerialNumberResponse = object;

export type PrivacyMsgCreateSerialNumberResponse = object;

export type PrivacyMsgCreateTokenResponse = object;

export type PrivacyMsgCreateTxPrivacyDataResponse = object;

export type PrivacyMsgCreateTxResponse = object;

export type PrivacyMsgDeleteCommitmentIndexResponse = object;

export type PrivacyMsgDeleteCommitmentResponse = object;

export type PrivacyMsgDeleteOnetimeAddressResponse = object;

export type PrivacyMsgDeleteOutputCoinResponse = object;

export type PrivacyMsgDeleteOutputCoinSerialNumberResponse = object;

export type PrivacyMsgDeleteSerialNumberResponse = object;

export type PrivacyMsgDeleteTokenResponse = object;

export type PrivacyMsgDeleteTxPrivacyDataResponse = object;

export interface PrivacyMsgTransferResponse {
  msg?: string;
  is_error?: boolean;
}

export type PrivacyMsgUpdateCommitmentIndexResponse = object;

export type PrivacyMsgUpdateCommitmentResponse = object;

export type PrivacyMsgUpdateOnetimeAddressResponse = object;

export type PrivacyMsgUpdateOutputCoinResponse = object;

export type PrivacyMsgUpdateOutputCoinSerialNumberResponse = object;

export type PrivacyMsgUpdateSerialNumberResponse = object;

export type PrivacyMsgUpdateTokenResponse = object;

export type PrivacyMsgUpdateTxPrivacyDataResponse = object;

export interface PrivacyOTACoin {
  index?: string;
  output_coin_index?: string;

  /** @format byte */
  height?: string;
}

export interface PrivacyOnetimeAddress {
  index?: string;
  creator?: string;
  is_confidential_asset?: boolean;

  /** @format byte */
  public_key?: string;

  /** @format byte */
  i?: string;

  /** @format int32 */
  status?: number;
}

export interface PrivacyOutputCoin {
  index?: string;

  /** @format byte */
  serial_number?: string;
  is_confidential_asset?: boolean;

  /** @format byte */
  pub_key?: string;

  /** @format byte */
  value?: string;
}

export interface PrivacyOutputCoinSerialNumber {
  /** @format byte */
  value?: string;
}

/**
 * Params defines the parameters for the module.
 */
export type PrivacyParams = object;

export interface PrivacyQueryAllCommitmentIndexResponse {
  commitmentIndex?: PrivacyCommitmentIndex[];

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

export interface PrivacyQueryAllCommitmentResponse {
  commitment?: PrivacyCommitment[];

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

export interface PrivacyQueryAllOTACoinResponse {
  oTACoin?: PrivacyOTACoin[];

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

export interface PrivacyQueryAllOnetimeAddressResponse {
  onetimeAddress?: PrivacyOnetimeAddress[];

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

export interface PrivacyQueryAllOutputCoinResponse {
  outputCoin?: PrivacyOutputCoin[];

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

export interface PrivacyQueryAllSerialNumberResponse {
  serialNumber?: PrivacySerialNumber[];

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

export interface PrivacyQueryAllTokenResponse {
  token?: PrivacyToken[];

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

export interface PrivacyQueryAllTxPrivacyDataResponse {
  txPrivacyData?: PrivacyTxPrivacyData[];

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

export interface PrivacyQueryBalanceResponse {
  /** @format uint64 */
  value?: string;
}

export interface PrivacyQueryGetCommitmentIndexResponse {
  commitmentIndex?: PrivacyCommitmentIndex;
}

export interface PrivacyQueryGetCommitmentResponse {
  commitment?: PrivacyCommitment;
}

export interface PrivacyQueryGetOTACoinResponse {
  oTACoin?: PrivacyOTACoin;
}

export interface PrivacyQueryGetOnetimeAddressResponse {
  onetimeAddress?: PrivacyOnetimeAddress;
}

export interface PrivacyQueryGetOutputCoinResponse {
  outputCoin?: PrivacyOutputCoin;
}

export interface PrivacyQueryGetOutputCoinSerialNumberResponse {
  OutputCoinSerialNumber?: PrivacyOutputCoinSerialNumber;
}

export interface PrivacyQueryGetSerialNumberResponse {
  serialNumber?: PrivacySerialNumber;
}

export interface PrivacyQueryGetTokenResponse {
  token?: PrivacyToken;
}

export interface PrivacyQueryGetTxPrivacyDataResponse {
  txPrivacyData?: PrivacyTxPrivacyData;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface PrivacyQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: PrivacyParams;
}

export interface PrivacySerialNumber {
  index?: string;
  creator?: string;
  is_confidential_asset?: boolean;

  /** @format byte */
  value?: string;
}

export interface PrivacyToken {
  index?: string;
  creator?: string;
  name?: string;
  symbol?: string;

  /** @format uint64 */
  amount?: string;

  /** @format byte */
  token_id?: string;
}

export interface PrivacyTxPrivacyData {
  index?: string;
  creator?: string;

  /** @format uint64 */
  lock_time?: string;

  /** @format uint64 */
  fee?: string;

  /** @format byte */
  info?: string;

  /** @format byte */
  sig_pub_key?: string;

  /** @format byte */
  sig?: string;

  /** @format byte */
  proof?: string;

  /** @format byte */
  messages?: string;
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
  /** @format byte */
  next_key?: string;

  /** @format uint64 */
  total?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title privacy/commitment.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryBalance
   * @summary Queries a list of Balance items.
   * @request GET:/privacy/privacy/balance/{privateKey}
   */
  queryBalance = (privateKey: string, params: RequestParams = {}) =>
    this.request<PrivacyQueryBalanceResponse, RpcStatus>({
      path: `/privacy/privacy/balance/${privateKey}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCommitmentAll
   * @summary Queries a list of Commitment items.
   * @request GET:/privacy/privacy/commitment
   */
  queryCommitmentAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PrivacyQueryAllCommitmentResponse, RpcStatus>({
      path: `/privacy/privacy/commitment`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCommitment
   * @summary Queries a Commitment by index.
   * @request GET:/privacy/privacy/commitment/{index}
   */
  queryCommitment = (index: string, params: RequestParams = {}) =>
    this.request<PrivacyQueryGetCommitmentResponse, RpcStatus>({
      path: `/privacy/privacy/commitment/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCommitmentIndexAll
   * @summary Queries a list of CommitmentIndex items.
   * @request GET:/privacy/privacy/commitment_index
   */
  queryCommitmentIndexAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PrivacyQueryAllCommitmentIndexResponse, RpcStatus>({
      path: `/privacy/privacy/commitment_index`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCommitmentIndex
   * @summary Queries a CommitmentIndex by index.
   * @request GET:/privacy/privacy/commitment_index/{index}
   */
  queryCommitmentIndex = (index: string, params: RequestParams = {}) =>
    this.request<PrivacyQueryGetCommitmentIndexResponse, RpcStatus>({
      path: `/privacy/privacy/commitment_index/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOnetimeAddressAll
   * @summary Queries a list of OnetimeAddress items.
   * @request GET:/privacy/privacy/onetime_address
   */
  queryOnetimeAddressAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PrivacyQueryAllOnetimeAddressResponse, RpcStatus>({
      path: `/privacy/privacy/onetime_address`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOnetimeAddress
   * @summary Queries a OnetimeAddress by index.
   * @request GET:/privacy/privacy/onetime_address/{index}
   */
  queryOnetimeAddress = (index: string, params: RequestParams = {}) =>
    this.request<PrivacyQueryGetOnetimeAddressResponse, RpcStatus>({
      path: `/privacy/privacy/onetime_address/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOtaCoinAll
   * @summary Queries a list of OTACoin items.
   * @request GET:/privacy/privacy/ota_coin
   */
  queryOtaCoinAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PrivacyQueryAllOTACoinResponse, RpcStatus>({
      path: `/privacy/privacy/ota_coin`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOtaCoin
   * @summary Queries a OTACoin by index.
   * @request GET:/privacy/privacy/ota_coin/{index}
   */
  queryOtaCoin = (index: string, params: RequestParams = {}) =>
    this.request<PrivacyQueryGetOTACoinResponse, RpcStatus>({
      path: `/privacy/privacy/ota_coin/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOutputCoinAll
   * @summary Queries a list of OutputCoin items.
   * @request GET:/privacy/privacy/output_coin
   */
  queryOutputCoinAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PrivacyQueryAllOutputCoinResponse, RpcStatus>({
      path: `/privacy/privacy/output_coin`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOutputCoin
   * @summary Queries a OutputCoin by index.
   * @request GET:/privacy/privacy/output_coin/{index}
   */
  queryOutputCoin = (index: string, params: RequestParams = {}) =>
    this.request<PrivacyQueryGetOutputCoinResponse, RpcStatus>({
      path: `/privacy/privacy/output_coin/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOutputCoinSerialNumber
   * @summary Queries a OutputCoinSerialNumber by index.
   * @request GET:/privacy/privacy/output_coin_serial_number
   */
  queryOutputCoinSerialNumber = (params: RequestParams = {}) =>
    this.request<PrivacyQueryGetOutputCoinSerialNumberResponse, RpcStatus>({
      path: `/privacy/privacy/output_coin_serial_number`,
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
   * @request GET:/privacy/privacy/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<PrivacyQueryParamsResponse, RpcStatus>({
      path: `/privacy/privacy/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySerialNumberAll
   * @summary Queries a list of SerialNumber items.
   * @request GET:/privacy/privacy/serial_number
   */
  querySerialNumberAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PrivacyQueryAllSerialNumberResponse, RpcStatus>({
      path: `/privacy/privacy/serial_number`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySerialNumber
   * @summary Queries a SerialNumber by index.
   * @request GET:/privacy/privacy/serial_number/{index}
   */
  querySerialNumber = (index: string, params: RequestParams = {}) =>
    this.request<PrivacyQueryGetSerialNumberResponse, RpcStatus>({
      path: `/privacy/privacy/serial_number/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTokenAll
   * @summary Queries a list of Token items.
   * @request GET:/privacy/privacy/token
   */
  queryTokenAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PrivacyQueryAllTokenResponse, RpcStatus>({
      path: `/privacy/privacy/token`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryToken
   * @summary Queries a Token by index.
   * @request GET:/privacy/privacy/token/{index}
   */
  queryToken = (index: string, params: RequestParams = {}) =>
    this.request<PrivacyQueryGetTokenResponse, RpcStatus>({
      path: `/privacy/privacy/token/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTxPrivacyDataAll
   * @summary Queries a list of TxPrivacyData items.
   * @request GET:/privacy/privacy/tx_privacy_data
   */
  queryTxPrivacyDataAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PrivacyQueryAllTxPrivacyDataResponse, RpcStatus>({
      path: `/privacy/privacy/tx_privacy_data`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTxPrivacyData
   * @summary Queries a TxPrivacyData by index.
   * @request GET:/privacy/privacy/tx_privacy_data/{index}
   */
  queryTxPrivacyData = (index: string, params: RequestParams = {}) =>
    this.request<PrivacyQueryGetTxPrivacyDataResponse, RpcStatus>({
      path: `/privacy/privacy/tx_privacy_data/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
