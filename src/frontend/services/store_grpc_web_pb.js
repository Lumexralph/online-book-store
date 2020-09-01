/**
 * @fileoverview gRPC-Web generated client stub for bookstore
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
proto.bookstore = require('./store_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.bookstore.ProductServiceClient =
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
proto.bookstore.ProductServicePromiseClient =
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
 *   !proto.bookstore.AddProductRequest,
 *   !proto.bookstore.AddProductResponse>}
 */
const methodDescriptor_ProductService_AddProduct = new grpc.web.MethodDescriptor(
  '/bookstore.ProductService/AddProduct',
  grpc.web.MethodType.UNARY,
  proto.bookstore.AddProductRequest,
  proto.bookstore.AddProductResponse,
  /**
   * @param {!proto.bookstore.AddProductRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.bookstore.AddProductResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.bookstore.AddProductRequest,
 *   !proto.bookstore.AddProductResponse>}
 */
const methodInfo_ProductService_AddProduct = new grpc.web.AbstractClientBase.MethodInfo(
  proto.bookstore.AddProductResponse,
  /**
   * @param {!proto.bookstore.AddProductRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.bookstore.AddProductResponse.deserializeBinary
);


/**
 * @param {!proto.bookstore.AddProductRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.bookstore.AddProductResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.bookstore.AddProductResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.bookstore.ProductServiceClient.prototype.addProduct =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/bookstore.ProductService/AddProduct',
      request,
      metadata || {},
      methodDescriptor_ProductService_AddProduct,
      callback);
};


/**
 * @param {!proto.bookstore.AddProductRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.bookstore.AddProductResponse>}
 *     A native promise that resolves to the response
 */
proto.bookstore.ProductServicePromiseClient.prototype.addProduct =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/bookstore.ProductService/AddProduct',
      request,
      metadata || {},
      methodDescriptor_ProductService_AddProduct);
};


module.exports = proto.bookstore;

