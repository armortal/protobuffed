/**
 * @fileoverview gRPC-Web generated client stub for armortal.protobuffed.examples
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.21.12
// source: example.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.armortal = {};
proto.armortal.protobuffed = {};
proto.armortal.protobuffed.examples = require('./example_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.armortal.protobuffed.examples.AuthClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.armortal.protobuffed.examples.AuthPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.armortal.protobuffed.examples.SignInRequest,
 *   !proto.armortal.protobuffed.examples.SignInResponse>}
 */
const methodDescriptor_Auth_SignIn = new grpc.web.MethodDescriptor(
  '/armortal.protobuffed.examples.Auth/SignIn',
  grpc.web.MethodType.UNARY,
  proto.armortal.protobuffed.examples.SignInRequest,
  proto.armortal.protobuffed.examples.SignInResponse,
  /**
   * @param {!proto.armortal.protobuffed.examples.SignInRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.armortal.protobuffed.examples.SignInResponse.deserializeBinary
);


/**
 * @param {!proto.armortal.protobuffed.examples.SignInRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.armortal.protobuffed.examples.SignInResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.armortal.protobuffed.examples.SignInResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.armortal.protobuffed.examples.AuthClient.prototype.signIn =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/armortal.protobuffed.examples.Auth/SignIn',
      request,
      metadata || {},
      methodDescriptor_Auth_SignIn,
      callback);
};


/**
 * @param {!proto.armortal.protobuffed.examples.SignInRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.armortal.protobuffed.examples.SignInResponse>}
 *     Promise that resolves to the response
 */
proto.armortal.protobuffed.examples.AuthPromiseClient.prototype.signIn =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/armortal.protobuffed.examples.Auth/SignIn',
      request,
      metadata || {},
      methodDescriptor_Auth_SignIn);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.armortal.protobuffed.examples.SignUpRequest,
 *   !proto.armortal.protobuffed.examples.SignUpResponse>}
 */
const methodDescriptor_Auth_SignUp = new grpc.web.MethodDescriptor(
  '/armortal.protobuffed.examples.Auth/SignUp',
  grpc.web.MethodType.UNARY,
  proto.armortal.protobuffed.examples.SignUpRequest,
  proto.armortal.protobuffed.examples.SignUpResponse,
  /**
   * @param {!proto.armortal.protobuffed.examples.SignUpRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.armortal.protobuffed.examples.SignUpResponse.deserializeBinary
);


/**
 * @param {!proto.armortal.protobuffed.examples.SignUpRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.armortal.protobuffed.examples.SignUpResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.armortal.protobuffed.examples.SignUpResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.armortal.protobuffed.examples.AuthClient.prototype.signUp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/armortal.protobuffed.examples.Auth/SignUp',
      request,
      metadata || {},
      methodDescriptor_Auth_SignUp,
      callback);
};


/**
 * @param {!proto.armortal.protobuffed.examples.SignUpRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.armortal.protobuffed.examples.SignUpResponse>}
 *     Promise that resolves to the response
 */
proto.armortal.protobuffed.examples.AuthPromiseClient.prototype.signUp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/armortal.protobuffed.examples.Auth/SignUp',
      request,
      metadata || {},
      methodDescriptor_Auth_SignUp);
};


module.exports = proto.armortal.protobuffed.examples;

