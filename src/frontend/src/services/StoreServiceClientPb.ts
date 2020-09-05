/**
 * @fileoverview gRPC-Web generated client stub for bookstore
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';

import {
  AddProductRequest,
  AddProductResponse} from './store_pb';

export class ProductServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoAddProduct = new grpcWeb.AbstractClientBase.MethodInfo(
    AddProductResponse,
    (request: AddProductRequest) => {
      return request.serializeBinary();
    },
    AddProductResponse.deserializeBinary
  );

  addProduct(
    request: AddProductRequest,
    metadata: grpcWeb.Metadata | null): Promise<AddProductResponse>;

  addProduct(
    request: AddProductRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: AddProductResponse) => void): grpcWeb.ClientReadableStream<AddProductResponse>;

  addProduct(
    request: AddProductRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: AddProductResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        new URL('/bookstore.ProductService/AddProduct', this.hostname_).toString(),
        request,
        metadata || {},
        this.methodInfoAddProduct,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/bookstore.ProductService/AddProduct',
    request,
    metadata || {},
    this.methodInfoAddProduct);
  }

}

