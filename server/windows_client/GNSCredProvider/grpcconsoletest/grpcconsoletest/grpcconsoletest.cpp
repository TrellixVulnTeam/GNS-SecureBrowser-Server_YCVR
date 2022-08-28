// grpcconsoletest.cpp : This file contains the 'main' function. Program execution begins and ends there.
//
#include <iostream>
#include <windows.h>
#include <GNSGRPCLIB.h>
#include <tchar.h>

using namespace std;

string client_key;
string client_crt;
string ca_crt;
bool LoadKeys()
{
    LPVOID pdata;
    LPBYTE sData;
    TCHAR clientKeyName[5] = _T("#105");
    TCHAR clientCrtName[5] = _T("#106");
    TCHAR CAKeyName[5] = _T("#107");
    TCHAR sRestype[13] = _T("CERTS");
    HRSRC hres;
    HGLOBAL hbytes;
    //Load client key
    hres = FindResource(NULL, clientKeyName, sRestype);
    if (hres == 0)
    {
        _tcprintf(_T("An Error Occurred.\n Could Not Locate Resource File."));
        return true;
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
        return true;
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
        return true;
    }
    hbytes = LoadResource(NULL, hres);
    pdata = LockResource(hbytes);
    sData = (LPBYTE)pdata;
    char* CACRTStr = (char*)sData;
    ca_crt = string(CACRTStr);

    return false;
}

void MyCallback(CardStatus input)
{
    cout << "Received card status: " << input.status() << endl;
}

int main()
{
    LoadKeys();
    

    grpc::SslCredentialsOptions sslOpts;
    sslOpts.pem_root_certs = ca_crt;
    sslOpts.pem_private_key = client_key;
    sslOpts.pem_cert_chain = client_crt;
    auto creds = grpc::SslCredentials(sslOpts);

    wstring uuid;

    //sslOpts.pem_private_key = string(clientkeyStr);

    GNSGRPCClient client(
        grpc::CreateChannel("127.0.0.1:50051",
            creds));

    wcout << L"Hello World!\n";
    client.ReadUUID(uuid);
    wcout << L"UUID: " << uuid;

    WinCreds result;
    bool error = client.ReadWinCreds(result);
    cout << "Found: " << result.wincreds_size() << " creds " << endl;
    for (int i = 0; i < result.wincreds_size(); ++i)
    {
        cout << "Idx: " << result.wincreds().at(i).idx() << " Username : " << result.wincreds().at(i).username() << endl;
    }

    client.SetupCallback(&MyCallback);

    std::cin.clear(); std::cin.ignore(INT_MAX, '\n');
    std::cin.get();
    return 0;
}

// Run program: Ctrl + F5 or Debug > Start Without Debugging menu
// Debug program: F5 or Debug > Start Debugging menu

// Tips for Getting Started: 
//   1. Use the Solution Explorer window to add/manage files
//   2. Use the Team Explorer window to connect to source control
//   3. Use the Output window to see build output and other messages
//   4. Use the Error List window to view errors
//   5. Go to Project > Add New Item to create new code files, or Project > Add Existing Item to add existing code files to the project
//   6. In the future, to open this project again, go to File > Open > Project and select the .sln file
