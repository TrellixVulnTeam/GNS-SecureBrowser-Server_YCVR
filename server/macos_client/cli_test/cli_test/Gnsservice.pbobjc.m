// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gnsservice.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers_RuntimeSupport.h>
#else
 #import "GPBProtocolBuffers_RuntimeSupport.h"
#endif

#import <stdatomic.h>

#import "Gnsservice.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"
#pragma clang diagnostic ignored "-Wdollar-in-identifier-extension"

#pragma mark - Objective C Class declarations
// Forward declarations of Objective C classes that we can use as
// static values in struct initializers.
// We don't use [Foo class] because it is not a static value.
GPBObjCClassDeclaration(GNSRPCSiteCred);
GPBObjCClassDeclaration(GNSRPCWinCred);

#pragma mark - GNSRPCGnsserviceRoot

@implementation GNSRPCGnsserviceRoot

// No extensions in the file and no imports, so no need to generate
// +extensionRegistry.

@end

#pragma mark - GNSRPCGnsserviceRoot_FileDescriptor

static GPBFileDescriptor *GNSRPCGnsserviceRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@""
                                                 objcPrefix:@"GNSRPC"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - GNSRPCUUID

@implementation GNSRPCUUID

@dynamic uuid;
@dynamic mode;

typedef struct GNSRPCUUID__storage_ {
  uint32_t _has_storage_[1];
  uint32_t mode;
  NSString *uuid;
} GNSRPCUUID__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "uuid",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCUUID_FieldNumber_Uuid,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(GNSRPCUUID__storage_, uuid),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "mode",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCUUID_FieldNumber_Mode,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(GNSRPCUUID__storage_, mode),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeUInt32,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GNSRPCUUID class]
                                     rootClass:[GNSRPCGnsserviceRoot class]
                                          file:GNSRPCGnsserviceRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GNSRPCUUID__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - GNSRPCFreeSites

@implementation GNSRPCFreeSites

@dynamic idxArray, idxArray_Count;

typedef struct GNSRPCFreeSites__storage_ {
  uint32_t _has_storage_[1];
  GPBUInt32Array *idxArray;
} GNSRPCFreeSites__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "idxArray",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCFreeSites_FieldNumber_IdxArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(GNSRPCFreeSites__storage_, idxArray),
        .flags = (GPBFieldFlags)(GPBFieldRepeated | GPBFieldPacked),
        .dataType = GPBDataTypeUInt32,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GNSRPCFreeSites class]
                                     rootClass:[GNSRPCGnsserviceRoot class]
                                          file:GNSRPCGnsserviceRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GNSRPCFreeSites__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - GNSRPCFreeWinCreds

@implementation GNSRPCFreeWinCreds

@dynamic idxArray, idxArray_Count;

typedef struct GNSRPCFreeWinCreds__storage_ {
  uint32_t _has_storage_[1];
  GPBUInt32Array *idxArray;
} GNSRPCFreeWinCreds__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "idxArray",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCFreeWinCreds_FieldNumber_IdxArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(GNSRPCFreeWinCreds__storage_, idxArray),
        .flags = (GPBFieldFlags)(GPBFieldRepeated | GPBFieldPacked),
        .dataType = GPBDataTypeUInt32,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GNSRPCFreeWinCreds class]
                                     rootClass:[GNSRPCGnsserviceRoot class]
                                          file:GNSRPCGnsserviceRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GNSRPCFreeWinCreds__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - GNSRPCSiteCred

@implementation GNSRPCSiteCred

@dynamic idx;
@dynamic offset;
@dynamic code;
@dynamic username;
@dynamic password;
@dynamic misc;

typedef struct GNSRPCSiteCred__storage_ {
  uint32_t _has_storage_[1];
  uint32_t idx;
  uint32_t offset;
  NSString *code;
  NSString *username;
  NSString *password;
  NSString *misc;
} GNSRPCSiteCred__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "idx",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCSiteCred_FieldNumber_Idx,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(GNSRPCSiteCred__storage_, idx),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeUInt32,
      },
      {
        .name = "offset",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCSiteCred_FieldNumber_Offset,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(GNSRPCSiteCred__storage_, offset),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeUInt32,
      },
      {
        .name = "code",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCSiteCred_FieldNumber_Code,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(GNSRPCSiteCred__storage_, code),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "username",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCSiteCred_FieldNumber_Username,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(GNSRPCSiteCred__storage_, username),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "password",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCSiteCred_FieldNumber_Password,
        .hasIndex = 4,
        .offset = (uint32_t)offsetof(GNSRPCSiteCred__storage_, password),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "misc",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCSiteCred_FieldNumber_Misc,
        .hasIndex = 5,
        .offset = (uint32_t)offsetof(GNSRPCSiteCred__storage_, misc),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GNSRPCSiteCred class]
                                     rootClass:[GNSRPCGnsserviceRoot class]
                                          file:GNSRPCGnsserviceRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GNSRPCSiteCred__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - GNSRPCSites

@implementation GNSRPCSites

@dynamic sitesArray, sitesArray_Count;

typedef struct GNSRPCSites__storage_ {
  uint32_t _has_storage_[1];
  NSMutableArray *sitesArray;
} GNSRPCSites__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "sitesArray",
        .dataTypeSpecific.clazz = GPBObjCClass(GNSRPCSiteCred),
        .number = GNSRPCSites_FieldNumber_SitesArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(GNSRPCSites__storage_, sitesArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GNSRPCSites class]
                                     rootClass:[GNSRPCGnsserviceRoot class]
                                          file:GNSRPCGnsserviceRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GNSRPCSites__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - GNSRPCWinCred

@implementation GNSRPCWinCred

@dynamic idx;
@dynamic domain;
@dynamic username;
@dynamic password;

typedef struct GNSRPCWinCred__storage_ {
  uint32_t _has_storage_[1];
  uint32_t idx;
  NSString *domain;
  NSString *username;
  NSString *password;
} GNSRPCWinCred__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "idx",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCWinCred_FieldNumber_Idx,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(GNSRPCWinCred__storage_, idx),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeUInt32,
      },
      {
        .name = "domain",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCWinCred_FieldNumber_Domain,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(GNSRPCWinCred__storage_, domain),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "username",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCWinCred_FieldNumber_Username,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(GNSRPCWinCred__storage_, username),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "password",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCWinCred_FieldNumber_Password,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(GNSRPCWinCred__storage_, password),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GNSRPCWinCred class]
                                     rootClass:[GNSRPCGnsserviceRoot class]
                                          file:GNSRPCGnsserviceRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GNSRPCWinCred__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - GNSRPCWinCreds

@implementation GNSRPCWinCreds

@dynamic wincredsArray, wincredsArray_Count;

typedef struct GNSRPCWinCreds__storage_ {
  uint32_t _has_storage_[1];
  NSMutableArray *wincredsArray;
} GNSRPCWinCreds__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "wincredsArray",
        .dataTypeSpecific.clazz = GPBObjCClass(GNSRPCWinCred),
        .number = GNSRPCWinCreds_FieldNumber_WincredsArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(GNSRPCWinCreds__storage_, wincredsArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GNSRPCWinCreds class]
                                     rootClass:[GNSRPCGnsserviceRoot class]
                                          file:GNSRPCGnsserviceRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GNSRPCWinCreds__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - GNSRPCCardStatus

@implementation GNSRPCCardStatus

@dynamic type;
@dynamic status;

typedef struct GNSRPCCardStatus__storage_ {
  uint32_t _has_storage_[1];
  GNSRPCCardStatus_ConnectionType type;
  GNSRPCCardStatus_ConnectionStatus status;
} GNSRPCCardStatus__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "type",
        .dataTypeSpecific.enumDescFunc = GNSRPCCardStatus_ConnectionType_EnumDescriptor,
        .number = GNSRPCCardStatus_FieldNumber_Type,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(GNSRPCCardStatus__storage_, type),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldHasEnumDescriptor | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeEnum,
      },
      {
        .name = "status",
        .dataTypeSpecific.enumDescFunc = GNSRPCCardStatus_ConnectionStatus_EnumDescriptor,
        .number = GNSRPCCardStatus_FieldNumber_Status,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(GNSRPCCardStatus__storage_, status),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldHasEnumDescriptor | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeEnum,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GNSRPCCardStatus class]
                                     rootClass:[GNSRPCGnsserviceRoot class]
                                          file:GNSRPCGnsserviceRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GNSRPCCardStatus__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

int32_t GNSRPCCardStatus_Type_RawValue(GNSRPCCardStatus *message) {
  GPBDescriptor *descriptor = [GNSRPCCardStatus descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:GNSRPCCardStatus_FieldNumber_Type];
  return GPBGetMessageRawEnumField(message, field);
}

void SetGNSRPCCardStatus_Type_RawValue(GNSRPCCardStatus *message, int32_t value) {
  GPBDescriptor *descriptor = [GNSRPCCardStatus descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:GNSRPCCardStatus_FieldNumber_Type];
  GPBSetMessageRawEnumField(message, field, value);
}

int32_t GNSRPCCardStatus_Status_RawValue(GNSRPCCardStatus *message) {
  GPBDescriptor *descriptor = [GNSRPCCardStatus descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:GNSRPCCardStatus_FieldNumber_Status];
  return GPBGetMessageRawEnumField(message, field);
}

void SetGNSRPCCardStatus_Status_RawValue(GNSRPCCardStatus *message, int32_t value) {
  GPBDescriptor *descriptor = [GNSRPCCardStatus descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:GNSRPCCardStatus_FieldNumber_Status];
  GPBSetMessageRawEnumField(message, field, value);
}

#pragma mark - Enum GNSRPCCardStatus_ConnectionType

GPBEnumDescriptor *GNSRPCCardStatus_ConnectionType_EnumDescriptor(void) {
  static _Atomic(GPBEnumDescriptor*) descriptor = nil;
  if (!descriptor) {
    static const char *valueNames =
        "Usb\000Nfc\000";
    static const int32_t values[] = {
        GNSRPCCardStatus_ConnectionType_Usb,
        GNSRPCCardStatus_ConnectionType_Nfc,
    };
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(GNSRPCCardStatus_ConnectionType)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:GNSRPCCardStatus_ConnectionType_IsValidValue];
    GPBEnumDescriptor *expected = nil;
    if (!atomic_compare_exchange_strong(&descriptor, &expected, worker)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL GNSRPCCardStatus_ConnectionType_IsValidValue(int32_t value__) {
  switch (value__) {
    case GNSRPCCardStatus_ConnectionType_Usb:
    case GNSRPCCardStatus_ConnectionType_Nfc:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - Enum GNSRPCCardStatus_ConnectionStatus

GPBEnumDescriptor *GNSRPCCardStatus_ConnectionStatus_EnumDescriptor(void) {
  static _Atomic(GPBEnumDescriptor*) descriptor = nil;
  if (!descriptor) {
    static const char *valueNames =
        "Disconnected\000Connected\000Authenticated\000";
    static const int32_t values[] = {
        GNSRPCCardStatus_ConnectionStatus_Disconnected,
        GNSRPCCardStatus_ConnectionStatus_Connected,
        GNSRPCCardStatus_ConnectionStatus_Authenticated,
    };
    static const char *extraTextFormatInfo = "\003\000\014\000\001\t\000\002\r\000";
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(GNSRPCCardStatus_ConnectionStatus)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:GNSRPCCardStatus_ConnectionStatus_IsValidValue
                              extraTextFormatInfo:extraTextFormatInfo];
    GPBEnumDescriptor *expected = nil;
    if (!atomic_compare_exchange_strong(&descriptor, &expected, worker)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL GNSRPCCardStatus_ConnectionStatus_IsValidValue(int32_t value__) {
  switch (value__) {
    case GNSRPCCardStatus_ConnectionStatus_Disconnected:
    case GNSRPCCardStatus_ConnectionStatus_Connected:
    case GNSRPCCardStatus_ConnectionStatus_Authenticated:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - GNSRPCCommand

@implementation GNSRPCCommand

@dynamic cmd;

typedef struct GNSRPCCommand__storage_ {
  uint32_t _has_storage_[1];
  NSString *cmd;
} GNSRPCCommand__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "cmd",
        .dataTypeSpecific.clazz = Nil,
        .number = GNSRPCCommand_FieldNumber_Cmd,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(GNSRPCCommand__storage_, cmd),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GNSRPCCommand class]
                                     rootClass:[GNSRPCGnsserviceRoot class]
                                          file:GNSRPCGnsserviceRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GNSRPCCommand__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - GNSRPCGNSBadgeDataParam

@implementation GNSRPCGNSBadgeDataParam


typedef struct GNSRPCGNSBadgeDataParam__storage_ {
  uint32_t _has_storage_[1];
} GNSRPCGNSBadgeDataParam__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GNSRPCGNSBadgeDataParam class]
                                     rootClass:[GNSRPCGnsserviceRoot class]
                                          file:GNSRPCGnsserviceRoot_FileDescriptor()
                                        fields:NULL
                                    fieldCount:0
                                   storageSize:sizeof(GNSRPCGNSBadgeDataParam__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
