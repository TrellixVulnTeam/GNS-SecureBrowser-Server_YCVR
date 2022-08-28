//
//  DataService.swift
//  GNSInternalTools
//
//  Created by Peter Pham on 2/23/22.
//

import Foundation
import GRPC
import NIOCore
import NIOPosix
import NIOSSL

let hostName: String = "127.0.0.1"
let port:Int = 50051
class api {
    func Connect() throws -> GNSRPC_GNSBadgeDataClient {
        let group = MultiThreadedEventLoopGroup(numberOfThreads: 1)
        let builder: ClientConnection.Builder
        let private_key_path  = Bundle.main.path(forResource: "Client_unencrypted", ofType: "key")
        let private_crt_path  = Bundle.main.path(forResource: "Client_unencrypted", ofType: "crt")
        let root_ca_path  = Bundle.main.path(forResource: "global-net-solutions.com", ofType: "crt")
        let private_key = try NIOSSLPrivateKey(file: private_key_path!, format: .pem)
        let private_chain = try NIOSSLCertificate(file: private_crt_path!, format: .pem)
        
        var configuration = TLSConfiguration.makeClientConfiguration()
        configuration.trustRoots = .file(root_ca_path!)
        let sslContext = try NIOSSLContext(configuration: configuration)
        let handler = try NIOSSLClientHandler(context: sslContext, serverHostname: hostName + "\(port)")
        builder = ClientConnection.secure(group: group).withTLS(privateKey: private_key)
                                                    .withTLS(certificateChain: [private_chain])
                                                    .withTLS(trustRoots: configuration.trustRoots!)
                                                    
                                   
        let connection = builder.connect(host: hostName, port: port)
        return GNSRPC_GNSBadgeDataClient(channel: connection)
    }
    
    func readUUID(client: GNSRPC_GNSBadgeDataClient) -> EventLoopFuture<GNSRPC_UUID> {
        print("read UUID")
        let n = GNSRPC_GNSBadgeDataParam.init()
        let sitesReq = client.readUUID(n)
        return sitesReq.response
        
    }
    
    
    func getSites(client: GNSRPC_GNSBadgeDataClient) -> EventLoopFuture<GNSRPC_FreeSites> {
        print("Calling get site")
        let n = GNSRPC_GNSBadgeDataParam.init()
        let sitesReq = client.getFreeSites(n)
        return sitesReq.response
        
    }
    
    func readSites(client: GNSRPC_GNSBadgeDataClient) -> EventLoopFuture<GNSRPC_Sites> {
        print("Calling read site")
        let sitesReq = client.readSiteCreds(GNSRPC_GNSBadgeDataParam.init())
        return sitesReq.response
        
    }
    func writeSite(client: GNSRPC_GNSBadgeDataClient, data: GNSRPC_SiteCred) -> EventLoopFuture<GNSRPC_GNSBadgeDataParam> {
        print("Calling read site")
        let resp = client.writeSiteCred(data)
        return resp.response
        
    }
    
    func formatCard(client: GNSRPC_GNSBadgeDataClient, mode: GNSRPC_UUID) -> EventLoopFuture<GNSRPC_GNSBadgeDataParam> {
        let resp = client.formatCard(mode)
        return resp.response
        
    }
}


