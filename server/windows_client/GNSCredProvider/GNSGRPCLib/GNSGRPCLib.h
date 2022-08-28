#pragma once
#include <iostream>
#include <grpc/grpc.h>
#include <grpcpp/channel.h>
#include <grpcpp/client_context.h>
#include <grpcpp/create_channel.h>
#include <grpcpp/security/credentials.h>
#include <gnsservice.pb.h>
#include <gnsservice.grpc.pb.h>
#include <atomic>

using grpc::Channel;
using grpc::ClientContext;
using grpc::ClientReader;
using grpc::ClientReaderWriter;
using grpc::ClientWriter;
using grpc::Status;

using namespace GNSRPC;
class GNSGRPCClient {
   
public:
    GNSGRPCClient(std::shared_ptr<Channel> channel)
        : stub_(GNSRPC::GNSBadgeData::NewStub(channel)) {};
    virtual ~GNSGRPCClient();
    
    bool ReadUUID(std::wstring& result);

    bool ReadWinCreds(WinCreds& result);

    bool ReadSites(Sites& result);

    bool WriteWinCred(const WinCred& input);

    bool WriteSite(const SiteCred& input);

    bool DeleteWinCred(const WinCred& input);

    bool DeleteSite(const SiteCred& input);

    static std::wstring WstringFromString(std::string in);

    void SetupCallback(std::function<void(CardStatus)>);
private:
    std::unique_ptr<GNSRPC::GNSBadgeData::Stub> stub_;
    std::unique_ptr<std::thread> backgroundThread;
    std::atomic_bool exiting_ = false;

};