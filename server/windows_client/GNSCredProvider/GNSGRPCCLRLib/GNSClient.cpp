#include "pch.h"
#include "GNSClient.h"
#include <tchar.h>
using namespace System;
using namespace System::Runtime::InteropServices;
using namespace GNSGRPCNET;
void GNSClient::Init() {
    string client_key;
    string client_crt;
    string ca_crt;
    LPVOID pdata;
    LPBYTE sData;
    TCHAR sRestype[13] = _T("CERTS");
    TCHAR CAKeyName[5] = _T("#102");
    TCHAR clientCrtName[5] = _T("#103");
    TCHAR clientKeyName[5] = _T("#104");
    HRSRC hres;
    HGLOBAL hbytes;
    //Load client key
    hres = FindResource(NULL, clientKeyName, sRestype);
    if (hres == 0)
    {
        _tcprintf(_T("An Error Occurred.\n Could Not Locate Resource File."));
    }
    hbytes = LoadResource(NULL, hres);
    pdata = LockResource(hbytes);
    sData = (LPBYTE)pdata;
    char* clientkeyStr = (char*)sData;
    client_key = string(clientkeyStr);
    //client_key[2] = 'G';

    hres = FindResource(NULL, clientCrtName, sRestype);
    if (hres == 0)
    {
        _tcprintf(_T("An Error Occurred.\n Could Not Locate Resource File."));
    }
    hbytes = LoadResource(NULL, hres);
    pdata = LockResource(hbytes);
    sData = (LPBYTE)pdata;
    char* clientCRTStr = (char*)sData;
    client_crt = string(clientCRTStr);

    hres = FindResource(NULL, CAKeyName, sRestype);
    if (hres == 0)
    {
        _tcprintf(_T("An Error Occurred.\n Could Not Locate Resource File."));
    }
    hbytes = LoadResource(NULL, hres);
    pdata = LockResource(hbytes);
    sData = (LPBYTE)pdata;
    char* CACRTStr = (char*)sData;
    ca_crt = string(CACRTStr);

    grpc::SslCredentialsOptions sslOpts;
    sslOpts.pem_root_certs = ca_crt;
    sslOpts.pem_private_key = client_key;
    sslOpts.pem_cert_chain = client_crt;
    auto creds = grpc::SslCredentials(sslOpts);
    //sslOpts.pem_private_key = string(clientkeyStr);

    _nativePtr = new GNSGRPCClient(
        grpc::CreateChannel("127.0.0.1:50051",
            creds));

    //https://codereview.stackexchange.com/questions/137897/interop-between-c-and-c-via-c-cli-with-callbacks
    ////CliCpp::OnConnectDelegate^ managed_on_connect = gcnew CliCpp::OnConnectDelegate(this, &WrapperClass::OnConnect);
    //ManagedCallbackHandler^ managed_callback = gcnew ManagedCallbackHandler(this, &GNSClient::Callback);
    IntPtr stub_ptr = Marshal::GetFunctionPointerForDelegate(cb);
    typedef void(__stdcall* CPPNativeCB)(GNSRPC::CardStatus);
    auto callbackFunctionCpp = static_cast<CPPNativeCB>(stub_ptr.ToPointer());
    _nativePtr->SetupCallback(callbackFunctionCpp);

}

void GNSClient::Callback(GNSRPC::CardStatus status)
{
    static GNSRPC::CardStatus prev;
    bool change = (prev.status() != status.status());

    if (status.status() == CardStatus::Authenticated && change)
    {
        cb(ManagedCardStatus::Ready);
    }
    else if (status.status() == CardStatus::Connected && change)
    {
        cb(ManagedCardStatus::Connected);
    }
    else if(change)
    {
        cb(ManagedCardStatus::NotReady);
    }
}

void GNSClient::SetupCallback(ManagedCallbackHandler^ handler)
{
    cb = handler;


}