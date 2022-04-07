/**
 * @fileoverview gRPC-Web generated client stub for WikiLinkParser
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.WikiLinkParser = require('./WikiLinkParserService_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.WikiLinkParser.WikiLinkParserClient =
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
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.WikiLinkParser.WikiLinkParserPromiseClient =
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
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.WikiLinkParser.EmptyMsg,
 *   !proto.WikiLinkParser.ConnectionStatus>}
 */
const methodDescriptor_WikiLinkParser_EstablishConnection = new grpc.web.MethodDescriptor(
  '/WikiLinkParser.WikiLinkParser/EstablishConnection',
  grpc.web.MethodType.UNARY,
  proto.WikiLinkParser.EmptyMsg,
  proto.WikiLinkParser.ConnectionStatus,
  /**
   * @param {!proto.WikiLinkParser.EmptyMsg} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.WikiLinkParser.ConnectionStatus.deserializeBinary
);


/**
 * @param {!proto.WikiLinkParser.EmptyMsg} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.WikiLinkParser.ConnectionStatus)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.WikiLinkParser.ConnectionStatus>|undefined}
 *     The XHR Node Readable Stream
 */
proto.WikiLinkParser.WikiLinkParserClient.prototype.establishConnection =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/WikiLinkParser.WikiLinkParser/EstablishConnection',
      request,
      metadata || {},
      methodDescriptor_WikiLinkParser_EstablishConnection,
      callback);
};


/**
 * @param {!proto.WikiLinkParser.EmptyMsg} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.WikiLinkParser.ConnectionStatus>}
 *     Promise that resolves to the response
 */
proto.WikiLinkParser.WikiLinkParserPromiseClient.prototype.establishConnection =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/WikiLinkParser.WikiLinkParser/EstablishConnection',
      request,
      metadata || {},
      methodDescriptor_WikiLinkParser_EstablishConnection);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.WikiLinkParser.EmptyMsg,
 *   !proto.WikiLinkParser.ConnectionStatus>}
 */
const methodDescriptor_WikiLinkParser_AbortConnection = new grpc.web.MethodDescriptor(
  '/WikiLinkParser.WikiLinkParser/AbortConnection',
  grpc.web.MethodType.UNARY,
  proto.WikiLinkParser.EmptyMsg,
  proto.WikiLinkParser.ConnectionStatus,
  /**
   * @param {!proto.WikiLinkParser.EmptyMsg} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.WikiLinkParser.ConnectionStatus.deserializeBinary
);


/**
 * @param {!proto.WikiLinkParser.EmptyMsg} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.WikiLinkParser.ConnectionStatus)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.WikiLinkParser.ConnectionStatus>|undefined}
 *     The XHR Node Readable Stream
 */
proto.WikiLinkParser.WikiLinkParserClient.prototype.abortConnection =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/WikiLinkParser.WikiLinkParser/AbortConnection',
      request,
      metadata || {},
      methodDescriptor_WikiLinkParser_AbortConnection,
      callback);
};


/**
 * @param {!proto.WikiLinkParser.EmptyMsg} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.WikiLinkParser.ConnectionStatus>}
 *     Promise that resolves to the response
 */
proto.WikiLinkParser.WikiLinkParserPromiseClient.prototype.abortConnection =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/WikiLinkParser.WikiLinkParser/AbortConnection',
      request,
      metadata || {},
      methodDescriptor_WikiLinkParser_AbortConnection);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.WikiLinkParser.ParseRequest,
 *   !proto.WikiLinkParser.ParseResponse>}
 */
const methodDescriptor_WikiLinkParser_PublishTask = new grpc.web.MethodDescriptor(
  '/WikiLinkParser.WikiLinkParser/PublishTask',
  grpc.web.MethodType.UNARY,
  proto.WikiLinkParser.ParseRequest,
  proto.WikiLinkParser.ParseResponse,
  /**
   * @param {!proto.WikiLinkParser.ParseRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.WikiLinkParser.ParseResponse.deserializeBinary
);


/**
 * @param {!proto.WikiLinkParser.ParseRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.WikiLinkParser.ParseResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.WikiLinkParser.ParseResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.WikiLinkParser.WikiLinkParserClient.prototype.publishTask =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/WikiLinkParser.WikiLinkParser/PublishTask',
      request,
      metadata || {},
      methodDescriptor_WikiLinkParser_PublishTask,
      callback);
};


/**
 * @param {!proto.WikiLinkParser.ParseRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.WikiLinkParser.ParseResponse>}
 *     Promise that resolves to the response
 */
proto.WikiLinkParser.WikiLinkParserPromiseClient.prototype.publishTask =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/WikiLinkParser.WikiLinkParser/PublishTask',
      request,
      metadata || {},
      methodDescriptor_WikiLinkParser_PublishTask);
};


module.exports = proto.WikiLinkParser;

