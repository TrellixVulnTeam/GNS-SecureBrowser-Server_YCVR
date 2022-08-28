// Code generated by gRPC proto compiler.  DO NOT EDIT!
// source: gnsservice.proto

#if !defined(GPB_GRPC_PROTOCOL_ONLY) || !GPB_GRPC_PROTOCOL_ONLY
#import "Gnsservice.pbrpc.h"
#import "Gnsservice.pbobjc.h"
#import <ProtoRPC/ProtoRPCLegacy.h>
#import <RxLibrary/GRXWriter+Immediate.h>


@implementation GNSRPCGNSBadgeData

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wobjc-designated-initializers"

// Designated initializer
- (instancetype)initWithHost:(NSString *)host callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [super initWithHost:host
                 packageName:@""
                 serviceName:@"GNSBadgeData"
                 callOptions:callOptions];
}

- (instancetype)initWithHost:(NSString *)host {
  return [super initWithHost:host
                 packageName:@""
                 serviceName:@"GNSBadgeData"];
}

#pragma clang diagnostic pop

// Override superclass initializer to disallow different package and service names.
- (instancetype)initWithHost:(NSString *)host
                 packageName:(NSString *)packageName
                 serviceName:(NSString *)serviceName {
  return [self initWithHost:host];
}

- (instancetype)initWithHost:(NSString *)host
                 packageName:(NSString *)packageName
                 serviceName:(NSString *)serviceName
                 callOptions:(GRPCCallOptions *)callOptions {
  return [self initWithHost:host callOptions:callOptions];
}

#pragma mark - Class Methods

+ (instancetype)serviceWithHost:(NSString *)host {
  return [[self alloc] initWithHost:host];
}

+ (instancetype)serviceWithHost:(NSString *)host callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [[self alloc] initWithHost:host callOptions:callOptions];
}

#pragma mark - Method Implementations

#pragma mark ReadUUID(GNSBadgeDataParam) returns (UUID)

/**
 * Get UUID of card
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (void)readUUIDWithRequest:(GNSRPCGNSBadgeDataParam *)request handler:(void(^)(GNSRPCUUID *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToReadUUIDWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
/**
 * Get UUID of card
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (GRPCProtoCall *)RPCToReadUUIDWithRequest:(GNSRPCGNSBadgeDataParam *)request handler:(void(^)(GNSRPCUUID *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"ReadUUID"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[GNSRPCUUID class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
/**
 * Get UUID of card
 */
- (GRPCUnaryProtoCall *)readUUIDWithMessage:(GNSRPCGNSBadgeDataParam *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"ReadUUID"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[GNSRPCUUID class]];
}

#pragma mark FormatCard(UUID) returns (GNSBadgeDataParam)

/**
 * Format card 0: sites + wincreds, 1: sites only, 2: wincreds
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (void)formatCardWithRequest:(GNSRPCUUID *)request handler:(void(^)(GNSRPCGNSBadgeDataParam *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToFormatCardWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
/**
 * Format card 0: sites + wincreds, 1: sites only, 2: wincreds
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (GRPCProtoCall *)RPCToFormatCardWithRequest:(GNSRPCUUID *)request handler:(void(^)(GNSRPCGNSBadgeDataParam *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"FormatCard"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[GNSRPCGNSBadgeDataParam class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
/**
 * Format card 0: sites + wincreds, 1: sites only, 2: wincreds
 */
- (GRPCUnaryProtoCall *)formatCardWithMessage:(GNSRPCUUID *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"FormatCard"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[GNSRPCGNSBadgeDataParam class]];
}

#pragma mark GetFreeSites(GNSBadgeDataParam) returns (FreeSites)

/**
 * Get available free indexes for sites (32 max)
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (void)getFreeSitesWithRequest:(GNSRPCGNSBadgeDataParam *)request handler:(void(^)(GNSRPCFreeSites *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToGetFreeSitesWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
/**
 * Get available free indexes for sites (32 max)
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (GRPCProtoCall *)RPCToGetFreeSitesWithRequest:(GNSRPCGNSBadgeDataParam *)request handler:(void(^)(GNSRPCFreeSites *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"GetFreeSites"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[GNSRPCFreeSites class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
/**
 * Get available free indexes for sites (32 max)
 */
- (GRPCUnaryProtoCall *)getFreeSitesWithMessage:(GNSRPCGNSBadgeDataParam *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"GetFreeSites"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[GNSRPCFreeSites class]];
}

#pragma mark GetFreeWinCreds(GNSBadgeDataParam) returns (FreeWinCreds)

/**
 * Get available windows credential (8 max)
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (void)getFreeWinCredsWithRequest:(GNSRPCGNSBadgeDataParam *)request handler:(void(^)(GNSRPCFreeWinCreds *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToGetFreeWinCredsWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
/**
 * Get available windows credential (8 max)
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (GRPCProtoCall *)RPCToGetFreeWinCredsWithRequest:(GNSRPCGNSBadgeDataParam *)request handler:(void(^)(GNSRPCFreeWinCreds *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"GetFreeWinCreds"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[GNSRPCFreeWinCreds class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
/**
 * Get available windows credential (8 max)
 */
- (GRPCUnaryProtoCall *)getFreeWinCredsWithMessage:(GNSRPCGNSBadgeDataParam *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"GetFreeWinCreds"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[GNSRPCFreeWinCreds class]];
}

#pragma mark ReadSiteCreds(GNSBadgeDataParam) returns (Sites)

/**
 * Read site credentials from badge
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (void)readSiteCredsWithRequest:(GNSRPCGNSBadgeDataParam *)request handler:(void(^)(GNSRPCSites *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToReadSiteCredsWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
/**
 * Read site credentials from badge
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (GRPCProtoCall *)RPCToReadSiteCredsWithRequest:(GNSRPCGNSBadgeDataParam *)request handler:(void(^)(GNSRPCSites *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"ReadSiteCreds"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[GNSRPCSites class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
/**
 * Read site credentials from badge
 */
- (GRPCUnaryProtoCall *)readSiteCredsWithMessage:(GNSRPCGNSBadgeDataParam *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"ReadSiteCreds"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[GNSRPCSites class]];
}

#pragma mark ReadWinCreds(GNSBadgeDataParam) returns (WinCreds)

/**
 * Read windows credentials from badge
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (void)readWinCredsWithRequest:(GNSRPCGNSBadgeDataParam *)request handler:(void(^)(GNSRPCWinCreds *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToReadWinCredsWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
/**
 * Read windows credentials from badge
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (GRPCProtoCall *)RPCToReadWinCredsWithRequest:(GNSRPCGNSBadgeDataParam *)request handler:(void(^)(GNSRPCWinCreds *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"ReadWinCreds"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[GNSRPCWinCreds class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
/**
 * Read windows credentials from badge
 */
- (GRPCUnaryProtoCall *)readWinCredsWithMessage:(GNSRPCGNSBadgeDataParam *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"ReadWinCreds"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[GNSRPCWinCreds class]];
}

#pragma mark DeleteSiteCred(SiteCred) returns (GNSBadgeDataParam)

/**
 * Delete 1 site credential
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (void)deleteSiteCredWithRequest:(GNSRPCSiteCred *)request handler:(void(^)(GNSRPCGNSBadgeDataParam *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToDeleteSiteCredWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
/**
 * Delete 1 site credential
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (GRPCProtoCall *)RPCToDeleteSiteCredWithRequest:(GNSRPCSiteCred *)request handler:(void(^)(GNSRPCGNSBadgeDataParam *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"DeleteSiteCred"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[GNSRPCGNSBadgeDataParam class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
/**
 * Delete 1 site credential
 */
- (GRPCUnaryProtoCall *)deleteSiteCredWithMessage:(GNSRPCSiteCred *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"DeleteSiteCred"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[GNSRPCGNSBadgeDataParam class]];
}

#pragma mark DeleteWinCred(WinCred) returns (GNSBadgeDataParam)

/**
 * Delete 1 Windows credential
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (void)deleteWinCredWithRequest:(GNSRPCWinCred *)request handler:(void(^)(GNSRPCGNSBadgeDataParam *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToDeleteWinCredWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
/**
 * Delete 1 Windows credential
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (GRPCProtoCall *)RPCToDeleteWinCredWithRequest:(GNSRPCWinCred *)request handler:(void(^)(GNSRPCGNSBadgeDataParam *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"DeleteWinCred"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[GNSRPCGNSBadgeDataParam class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
/**
 * Delete 1 Windows credential
 */
- (GRPCUnaryProtoCall *)deleteWinCredWithMessage:(GNSRPCWinCred *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"DeleteWinCred"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[GNSRPCGNSBadgeDataParam class]];
}

#pragma mark WriteSiteCred(SiteCred) returns (GNSBadgeDataParam)

/**
 * Write a site credential at location of offset
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (void)writeSiteCredWithRequest:(GNSRPCSiteCred *)request handler:(void(^)(GNSRPCGNSBadgeDataParam *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToWriteSiteCredWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
/**
 * Write a site credential at location of offset
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (GRPCProtoCall *)RPCToWriteSiteCredWithRequest:(GNSRPCSiteCred *)request handler:(void(^)(GNSRPCGNSBadgeDataParam *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"WriteSiteCred"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[GNSRPCGNSBadgeDataParam class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
/**
 * Write a site credential at location of offset
 */
- (GRPCUnaryProtoCall *)writeSiteCredWithMessage:(GNSRPCSiteCred *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"WriteSiteCred"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[GNSRPCGNSBadgeDataParam class]];
}

#pragma mark WriteWinCred(WinCred) returns (GNSBadgeDataParam)

/**
 * Write a windows credential at location of idx
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (void)writeWinCredWithRequest:(GNSRPCWinCred *)request handler:(void(^)(GNSRPCGNSBadgeDataParam *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToWriteWinCredWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
/**
 * Write a windows credential at location of idx
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (GRPCProtoCall *)RPCToWriteWinCredWithRequest:(GNSRPCWinCred *)request handler:(void(^)(GNSRPCGNSBadgeDataParam *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"WriteWinCred"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[GNSRPCGNSBadgeDataParam class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
/**
 * Write a windows credential at location of idx
 */
- (GRPCUnaryProtoCall *)writeWinCredWithMessage:(GNSRPCWinCred *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"WriteWinCred"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[GNSRPCGNSBadgeDataParam class]];
}

#pragma mark StreamCardStatus(GNSBadgeDataParam) returns (stream CardStatus)

/**
 * Ping-pong like in-out data stream to report CardStatus changes to RPC
 * client
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (void)streamCardStatusWithRequest:(GNSRPCGNSBadgeDataParam *)request eventHandler:(void(^)(BOOL done, GNSRPCCardStatus *_Nullable response, NSError *_Nullable error))eventHandler{
  [[self RPCToStreamCardStatusWithRequest:request eventHandler:eventHandler] start];
}
// Returns a not-yet-started RPC object.
/**
 * Ping-pong like in-out data stream to report CardStatus changes to RPC
 * client
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (GRPCProtoCall *)RPCToStreamCardStatusWithRequest:(GNSRPCGNSBadgeDataParam *)request eventHandler:(void(^)(BOOL done, GNSRPCCardStatus *_Nullable response, NSError *_Nullable error))eventHandler{
  return [self RPCToMethod:@"StreamCardStatus"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[GNSRPCCardStatus class]
        responsesWriteable:[GRXWriteable writeableWithEventHandler:eventHandler]];
}
/**
 * Ping-pong like in-out data stream to report CardStatus changes to RPC
 * client
 */
- (GRPCUnaryProtoCall *)streamCardStatusWithMessage:(GNSRPCGNSBadgeDataParam *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"StreamCardStatus"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[GNSRPCCardStatus class]];
}

#pragma mark Execute(Command) returns (GNSBadgeDataParam)

/**
 * arbitrary client to server command to implement commands like switching
 * between USB vs NFC
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (void)executeWithRequest:(GNSRPCCommand *)request handler:(void(^)(GNSRPCGNSBadgeDataParam *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToExecuteWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
/**
 * arbitrary client to server command to implement commands like switching
 * between USB vs NFC
 *
 * This method belongs to a set of APIs that have been deprecated. Using the v2 API is recommended.
 */
- (GRPCProtoCall *)RPCToExecuteWithRequest:(GNSRPCCommand *)request handler:(void(^)(GNSRPCGNSBadgeDataParam *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"Execute"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[GNSRPCGNSBadgeDataParam class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
/**
 * arbitrary client to server command to implement commands like switching
 * between USB vs NFC
 */
- (GRPCUnaryProtoCall *)executeWithMessage:(GNSRPCCommand *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"Execute"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[GNSRPCGNSBadgeDataParam class]];
}

@end
#endif
