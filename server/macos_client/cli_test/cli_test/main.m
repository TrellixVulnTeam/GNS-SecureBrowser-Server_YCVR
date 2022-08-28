//
//  main.m
//  cli_test
//
//  Created by Peter Pham on 2/25/22.
//

#import <Foundation/Foundation.h>
#import "Gnsservice.pbrpc.h"
#import <GRPCClient/GRPCTransport.h>

static NSString * const kHostAddress = @"localhost:50051";

@interface GNSRPCSitesHandler: NSObject<GRPCProtoResponseHandler>
@end
// A response handler object dispatching messages to main queue
@implementation GNSRPCSitesHandler
- (dispatch_queue_t)dispatchQueue {
  return dispatch_get_main_queue();
}
- (void)didReceiveProtoMessage:(GPBMessage *)message {
    if ( [message isKindOfClass:[GNSRPCSites class]]) {
        NSLog(@"We got GNSRPCSItes");
        GNSRPCSites *sites = (GNSRPCSites *)message;
        for( int i=0; i < sites.sitesArray.count;i++)
        {
            NSLog(@"Site[%d] idx: %d username: %@ password: %@\n",
                  i,
                  sites.sitesArray[i].idx,
                  sites.sitesArray[i].username,
                  sites.sitesArray[i].password);
        }

    }
}
@end

int main(int argc, const char * argv[]) {
    dispatch_group_t serviceGroup = dispatch_group_create();
    GRPCMutableCallOptions *options = [[GRPCMutableCallOptions alloc] init];
    options.transport = GRPCDefaultTransportImplList.core_insecure;
    GNSRPCGNSBadgeData *client = [[GNSRPCGNSBadgeData alloc] initWithHost: kHostAddress];
    GNSRPCGNSBadgeDataParam *n = [[GNSRPCGNSBadgeDataParam alloc] init];
    dispatch_group_enter(serviceGroup);
    
    for(int i=0; i< 10 ;i++)
    {
    
    [[client readSiteCredsWithMessage:n
            responseHandler: [[GNSRPCSitesHandler alloc] init]
                          callOptions: options] start ];
    }
    dispatch_group_notify(serviceGroup, dispatch_get_main_queue(), ^{
        exit(EXIT_SUCCESS);
    });
    dispatch_main();
}
