// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gnsservice.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers.h>
#else
 #import "GPBProtocolBuffers.h"
#endif

#if GOOGLE_PROTOBUF_OBJC_VERSION < 30004
#error This file was generated by a newer version of protoc which is incompatible with your Protocol Buffer library sources.
#endif
#if 30004 < GOOGLE_PROTOBUF_OBJC_MIN_SUPPORTED_VERSION
#error This file was generated by an older version of protoc which is incompatible with your Protocol Buffer library sources.
#endif

// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

CF_EXTERN_C_BEGIN

@class GNSRPCSiteCred;
@class GNSRPCWinCred;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - Enum GNSRPCCardStatus_ConnectionType

typedef GPB_ENUM(GNSRPCCardStatus_ConnectionType) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  GNSRPCCardStatus_ConnectionType_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  GNSRPCCardStatus_ConnectionType_Usb = 0,
  GNSRPCCardStatus_ConnectionType_Nfc = 1,
};

GPBEnumDescriptor *GNSRPCCardStatus_ConnectionType_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL GNSRPCCardStatus_ConnectionType_IsValidValue(int32_t value);

#pragma mark - Enum GNSRPCCardStatus_ConnectionStatus

typedef GPB_ENUM(GNSRPCCardStatus_ConnectionStatus) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  GNSRPCCardStatus_ConnectionStatus_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  /** Device is not visible to OS */
  GNSRPCCardStatus_ConnectionStatus_Disconnected = 0,

  /** Device is visible but not yet authenticated */
  GNSRPCCardStatus_ConnectionStatus_Connected = 1,

  /** Device is authenticated aka ping-pong is successful */
  GNSRPCCardStatus_ConnectionStatus_Authenticated = 2,
};

GPBEnumDescriptor *GNSRPCCardStatus_ConnectionStatus_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL GNSRPCCardStatus_ConnectionStatus_IsValidValue(int32_t value);

#pragma mark - GNSRPCGnsserviceRoot

/**
 * Exposes the extension registry for this file.
 *
 * The base class provides:
 * @code
 *   + (GPBExtensionRegistry *)extensionRegistry;
 * @endcode
 * which is a @c GPBExtensionRegistry that includes all the extensions defined by
 * this file and all files that it depends on.
 **/
GPB_FINAL @interface GNSRPCGnsserviceRoot : GPBRootObject
@end

#pragma mark - GNSRPCUUID

typedef GPB_ENUM(GNSRPCUUID_FieldNumber) {
  GNSRPCUUID_FieldNumber_Uuid = 1,
  GNSRPCUUID_FieldNumber_Mode = 2,
};

GPB_FINAL @interface GNSRPCUUID : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *uuid;

@property(nonatomic, readwrite) uint32_t mode;

@end

#pragma mark - GNSRPCFreeSites

typedef GPB_ENUM(GNSRPCFreeSites_FieldNumber) {
  GNSRPCFreeSites_FieldNumber_IdxArray = 1,
};

GPB_FINAL @interface GNSRPCFreeSites : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) GPBUInt32Array *idxArray;
/** The number of items in @c idxArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger idxArray_Count;

@end

#pragma mark - GNSRPCFreeWinCreds

typedef GPB_ENUM(GNSRPCFreeWinCreds_FieldNumber) {
  GNSRPCFreeWinCreds_FieldNumber_IdxArray = 1,
};

GPB_FINAL @interface GNSRPCFreeWinCreds : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) GPBUInt32Array *idxArray;
/** The number of items in @c idxArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger idxArray_Count;

@end

#pragma mark - GNSRPCSiteCred

typedef GPB_ENUM(GNSRPCSiteCred_FieldNumber) {
  GNSRPCSiteCred_FieldNumber_Idx = 1,
  GNSRPCSiteCred_FieldNumber_Offset = 2,
  GNSRPCSiteCred_FieldNumber_Code = 3,
  GNSRPCSiteCred_FieldNumber_Username = 4,
  GNSRPCSiteCred_FieldNumber_Password = 5,
  GNSRPCSiteCred_FieldNumber_Misc = 6,
};

GPB_FINAL @interface GNSRPCSiteCred : GPBMessage

/** index 1 to 32 for USB serial;  index 1 to 24 for NFC */
@property(nonatomic, readwrite) uint32_t idx;

/** absolute byte offset of where the data is it's idx*blocklen */
@property(nonatomic, readwrite) uint32_t offset;

@property(nonatomic, readwrite, copy, null_resettable) NSString *code;

@property(nonatomic, readwrite, copy, null_resettable) NSString *username;

@property(nonatomic, readwrite, copy, null_resettable) NSString *password;

@property(nonatomic, readwrite, copy, null_resettable) NSString *misc;

@end

#pragma mark - GNSRPCSites

typedef GPB_ENUM(GNSRPCSites_FieldNumber) {
  GNSRPCSites_FieldNumber_SitesArray = 1,
};

GPB_FINAL @interface GNSRPCSites : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<GNSRPCSiteCred*> *sitesArray;
/** The number of items in @c sitesArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger sitesArray_Count;

@end

#pragma mark - GNSRPCWinCred

typedef GPB_ENUM(GNSRPCWinCred_FieldNumber) {
  GNSRPCWinCred_FieldNumber_Idx = 1,
  GNSRPCWinCred_FieldNumber_Domain = 2,
  GNSRPCWinCred_FieldNumber_Username = 3,
  GNSRPCWinCred_FieldNumber_Password = 4,
};

GPB_FINAL @interface GNSRPCWinCred : GPBMessage

/** index 1 through 8 for USB serial; no NFC windows cred yet */
@property(nonatomic, readwrite) uint32_t idx;

@property(nonatomic, readwrite, copy, null_resettable) NSString *domain;

@property(nonatomic, readwrite, copy, null_resettable) NSString *username;

@property(nonatomic, readwrite, copy, null_resettable) NSString *password;

@end

#pragma mark - GNSRPCWinCreds

typedef GPB_ENUM(GNSRPCWinCreds_FieldNumber) {
  GNSRPCWinCreds_FieldNumber_WincredsArray = 1,
};

GPB_FINAL @interface GNSRPCWinCreds : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<GNSRPCWinCred*> *wincredsArray;
/** The number of items in @c wincredsArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger wincredsArray_Count;

@end

#pragma mark - GNSRPCCardStatus

typedef GPB_ENUM(GNSRPCCardStatus_FieldNumber) {
  GNSRPCCardStatus_FieldNumber_Type = 1,
  GNSRPCCardStatus_FieldNumber_Status = 2,
};

GPB_FINAL @interface GNSRPCCardStatus : GPBMessage

@property(nonatomic, readwrite) GNSRPCCardStatus_ConnectionType type;

@property(nonatomic, readwrite) GNSRPCCardStatus_ConnectionStatus status;

@end

/**
 * Fetches the raw value of a @c GNSRPCCardStatus's @c type property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t GNSRPCCardStatus_Type_RawValue(GNSRPCCardStatus *message);
/**
 * Sets the raw value of an @c GNSRPCCardStatus's @c type property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetGNSRPCCardStatus_Type_RawValue(GNSRPCCardStatus *message, int32_t value);

/**
 * Fetches the raw value of a @c GNSRPCCardStatus's @c status property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t GNSRPCCardStatus_Status_RawValue(GNSRPCCardStatus *message);
/**
 * Sets the raw value of an @c GNSRPCCardStatus's @c status property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetGNSRPCCardStatus_Status_RawValue(GNSRPCCardStatus *message, int32_t value);

#pragma mark - GNSRPCCommand

typedef GPB_ENUM(GNSRPCCommand_FieldNumber) {
  GNSRPCCommand_FieldNumber_Cmd = 1,
};

GPB_FINAL @interface GNSRPCCommand : GPBMessage

/** command string to send to RPC server for example to switch */
@property(nonatomic, readwrite, copy, null_resettable) NSString *cmd;

@end

#pragma mark - GNSRPCGNSBadgeDataParam

/**
 * Null param place holder
 **/
GPB_FINAL @interface GNSRPCGNSBadgeDataParam : GPBMessage

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
