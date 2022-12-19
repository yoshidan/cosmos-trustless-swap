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

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

export type SwapMsgCancelNFTResponse = object;

export type SwapMsgCancelResponse = object;

export type SwapMsgReceiveNFTResponse = object;

export type SwapMsgReceiveResponse = object;

export interface SwapMsgSendNFTResponse {
  /** @format uint64 */
  id?: string;
}

export interface SwapMsgSendResponse {
  /** @format uint64 */
  id?: string;
}

export interface SwapNFTSwap {
  /** @format uint64 */
  id?: string;
  sender?: string;
  receiver?: string;
  classId?: string;
  nftId?: string;
  amountToReceive?: string;
}

/**
 * Params defines the parameters for the module.
 */
export type SwapParams = object;

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface SwapQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: SwapParams;
}

export interface SwapQueryShowNFTResponse {
  swap?: SwapNFTSwap;
}

export interface SwapQueryShowResponse {
  /** QueryParamsResponse is response type for the Query/Params RPC method. */
  swap?: SwapSwap;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface SwapSwap {
  /** @format uint64 */
  id?: string;
  sender?: string;
  receiver?: string;
  amount?: string;
  amountToReceive?: string;
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
 * @title swap/swap/genesis.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/swap/swap/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<SwapQueryParamsResponse, RpcStatus>({
      path: `/swap/swap/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryShow
   * @summary Queries a list of Show items.
   * @request GET:/swap/swap/show/{id}
   */
  queryShow = (id: string, params: RequestParams = {}) =>
    this.request<SwapQueryShowResponse, RpcStatus>({
      path: `/swap/swap/show/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryShowNft
   * @summary Queries a list of ShowNFT items.
   * @request GET:/swap/swap/show_nft/{id}
   */
  queryShowNft = (id: string, params: RequestParams = {}) =>
    this.request<SwapQueryShowNFTResponse, RpcStatus>({
      path: `/swap/swap/show_nft/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
