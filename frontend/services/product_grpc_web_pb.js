/**
 * @fileoverview gRPC-Web generated client stub for product.internal.proto_files.domain
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.product = {};
proto.product.internal = {};
proto.product.internal.proto_files = {};
proto.product.internal.proto_files.domain = require('./product_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.product.internal.proto_files.domain.ProductServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.product.internal.proto_files.domain.ProductServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.product.internal.proto_files.domain.AddProductRequest,
 *   !proto.product.internal.proto_files.domain.AddProductResponse>}
 */
const methodDescriptor_ProductService_AddProduct = new grpc.web.MethodDescriptor(
  '/product.internal.proto_files.domain.ProductService/AddProduct',
  grpc.web.MethodType.UNARY,
  proto.product.internal.proto_files.domain.AddProductRequest,
  proto.product.internal.proto_files.domain.AddProductResponse,
  /**
   * @param {!proto.product.internal.proto_files.domain.AddProductRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.product.internal.proto_files.domain.AddProductResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.product.internal.proto_files.domain.AddProductRequest,
 *   !proto.product.internal.proto_files.domain.AddProductResponse>}
 */
const methodInfo_ProductService_AddProduct = new grpc.web.AbstractClientBase.MethodInfo(
  proto.product.internal.proto_files.domain.AddProductResponse,
  /**
   * @param {!proto.product.internal.proto_files.domain.AddProductRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.product.internal.proto_files.domain.AddProductResponse.deserializeBinary
);


/**
 * @param {!proto.product.internal.proto_files.domain.AddProductRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.product.internal.proto_files.domain.AddProductResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.product.internal.proto_files.domain.AddProductResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.product.internal.proto_files.domain.ProductServiceClient.prototype.addProduct =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/product.internal.proto_files.domain.ProductService/AddProduct',
      request,
      metadata || {},
      methodDescriptor_ProductService_AddProduct,
      callback);
};


/**
 * @param {!proto.product.internal.proto_files.domain.AddProductRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.product.internal.proto_files.domain.AddProductResponse>}
 *     A native promise that resolves to the response
 */
proto.product.internal.proto_files.domain.ProductServicePromiseClient.prototype.addProduct =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/product.internal.proto_files.domain.ProductService/AddProduct',
      request,
      metadata || {},
      methodDescriptor_ProductService_AddProduct);
};


module.exports = proto.product.internal.proto_files.domain;

