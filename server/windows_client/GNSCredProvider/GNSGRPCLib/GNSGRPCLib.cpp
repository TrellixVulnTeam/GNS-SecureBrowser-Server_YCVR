// GNSGRPCLib.cpp : Defines the functions for the static library.
//

#pragma comment(lib, "ws2_32.lib")


#include "GNSGRPCLib.h"
#include <locale>
#include <codecvt>
#include <string>
#include <thread>
//std::wstring_convert<std::codecvt_utf8_utf16<wchar_t>> converter;
//std::string narrow = converter.to_bytes(wide_utf16_source_string);
//std::wstring wide = converter.from_bytes(narrow_utf8_source_string)

using grpc::Channel;
using grpc::ClientContext;
using grpc::ClientReader;
using grpc::ClientReaderWriter;
using grpc::ClientWriter;
using grpc::Status;
using namespace GNSRPC;


std::wstring GNSGRPCClient::WstringFromString(std::string input)
{
    std::wstring_convert<std::codecvt_utf8_utf16<wchar_t>> converter;
    return converter.from_bytes(input);
}


bool GNSGRPCClient::ReadUUID(std::wstring& result) {
        UUID data;
        ClientContext context;
        GNSBadgeDataParam empty;
        std::wstring wide;
        Status status = stub_->ReadUUID(&context, empty, &data);
        if (!status.ok()) {
            
            wide = WstringFromString(status.error_message());
            result = wide;
            //std::cout << "ReadUUID rpc failed." << std::endl;
            return true;
        }
        result = WstringFromString(data.uuid());
        return false;
 }

bool GNSGRPCClient::ReadWinCreds(WinCreds & result){
    ClientContext context;
    GNSBadgeDataParam empty;

    Status status = stub_->ReadWinCreds(&context, empty, &result);
    if (!status.ok()) {

        return true;
    }
    return false;
}

bool GNSGRPCClient::ReadSites(Sites& result) {
    ClientContext context;
    GNSBadgeDataParam empty;

    Status status = stub_->ReadSiteCreds(&context, empty, &result);
    if (!status.ok()) {

        return true;
    }
    return false;
}

bool GNSGRPCClient::WriteWinCred(const WinCred& input) {
    ClientContext context;
    GNSBadgeDataParam empty;

    Status status = stub_->WriteWinCred(&context, input, &empty);
    if (!status.ok()) {

        return true;
    }
    return false;
}

bool GNSGRPCClient::WriteSite(const SiteCred& input) {
    ClientContext context;
    GNSBadgeDataParam empty;

    Status status = stub_->WriteSiteCred(&context, input, &empty);
    if (!status.ok()) {

        return true;
    }
    return false;
}

bool GNSGRPCClient::DeleteWinCred(const WinCred& input) {
    ClientContext context;
    GNSBadgeDataParam empty;

    Status status = stub_->DeleteWinCred(&context, input, &empty);
    if (!status.ok()) {

        return true;
    }
    return false;
}

bool GNSGRPCClient::DeleteSite(const SiteCred& input) {
    ClientContext context;
    GNSBadgeDataParam empty;

    Status status = stub_->DeleteSiteCred(&context, input, &empty);
    if (!status.ok()) {

        return true;
    }
    return false;
}


void GNSGRPCClient::SetupCallback(std::function<void(CardStatus)> callback) {
    GNSRPC::GNSBadgeData::Stub* stub = stub_.get();

    backgroundThread = std::make_unique<std::thread>(
        [=] {


        while (true && !exiting_)
        {
            try {
                //auto stream = 
                ClientContext context;
                GNSBadgeDataParam empty;
                std::unique_ptr<ClientReader<GNSRPC::CardStatus> > reader(stub->StreamCardStatus(&context, empty));

                GNSRPC::CardStatus status;
                while (reader->Read(&status) && !exiting_) {
                    //std::cout << "current status: " << status.status();
                    callback(status);
                }
                reader.release();
                Sleep(1500);
            }
            catch (...)
            {
                
            }
            CardStatus disconnected;
            disconnected.set_status(CardStatus_ConnectionStatus_Disconnected);
            callback(disconnected);
        }
       });
}

GNSGRPCClient::~GNSGRPCClient()
{
    if (backgroundThread) {
        exiting_ = true;
        backgroundThread->join();
    }
}
