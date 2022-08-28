//
// THIS CODE AND INFORMATION IS PROVIDED "AS IS" WITHOUT WARRANTY OF
// ANY KIND, EITHER EXPRESSED OR IMPLIED, INCLUDING BUT NOT LIMITED TO
// THE IMPLIED WARRANTIES OF MERCHANTABILITY AND/OR FITNESS FOR A
// PARTICULAR PURPOSE.
//
// Copyright (c) Microsoft Corporation. All rights reserved.
//
//

#ifndef WIN32_NO_STATUS
#include <ntstatus.h>
#define WIN32_NO_STATUS
#endif
#include <unknwn.h>
#include "GNSCredential.h"
#include "guid.h"
#include <GNSGRPCLIB.h>
#include <string>
#include <locale>
#include <codecvt>
#include <string>
#include <tchar.h>
#include <iostream>
#include <fstream>

#define INFO_BUFFER_SIZE 32767
TCHAR  infoBuf[INFO_BUFFER_SIZE];
DWORD  bufCharCount = INFO_BUFFER_SIZE;
using namespace std;
string client_key;
string client_crt;
string ca_crt;
//ofstream callbackLog;
bool LoadKeys()
{
    LPVOID pdata;
    LPBYTE sData;
    TCHAR clientKeyName[5] = _T("#4");
    TCHAR clientCrtName[5] = _T("#104");
    TCHAR CAKeyName[5] = _T("#106");
    TCHAR sRestype[13] = _T("CERTS");
    HRSRC hres;
    HGLOBAL hbytes;
    //Load client key
    //HBITMAP hbmp = LoadBitmap(HINST_THISDLL, MAKEINTRESOURCE(IDB_TILE_IMAGE));
    hres = FindResource(HINST_THISDLL, clientKeyName, sRestype);
    if (hres == 0)
    {
        //_tcprintf(_T("An Error Occurred.\n Could Not Locate Resource File."));
        return true;
    }
    hbytes = LoadResource(HINST_THISDLL, hres);
    pdata = LockResource(hbytes);
    sData = (LPBYTE)pdata;
    char* clientkeyStr = (char*)sData;
    client_key = string(clientkeyStr);
    //client_key[2] = 'G';

    hres = FindResource(HINST_THISDLL, clientCrtName, sRestype);
    if (hres == 0)
    {
        //_tcprintf(_T("An Error Occurred.\n Could Not Locate Resource File."));
        return true;
    }
    hbytes = LoadResource(HINST_THISDLL, hres);
    pdata = LockResource(hbytes);
    sData = (LPBYTE)pdata;
    char* clientCRTStr = (char*)sData;
    client_crt = string(clientCRTStr);

    hres = FindResource(HINST_THISDLL, CAKeyName, sRestype);
    if (hres == 0)
    {
        //_tcprintf(_T("An Error Occurred.\n Could Not Locate Resource File."));
        return true;
    }
    hbytes = LoadResource(HINST_THISDLL, hres);
    pdata = LockResource(hbytes);
    sData = (LPBYTE)pdata;
    char* CACRTStr = (char*)sData;
    ca_crt = string(CACRTStr);

    return false;
}

GNSCredential::GNSCredential():
    _cRef(1),
    _pCredProvCredentialEvents(nullptr),
    _pszUserSid(nullptr),
    _pszQualifiedUserName(nullptr),
    _fIsLocalUser(false),
    _fChecked(false),
    _fShowControls(false),
    _dwComboIndex(0)
{
    DllAddRef();

    ZeroMemory(_rgCredProvFieldDescriptors, sizeof(_rgCredProvFieldDescriptors));
    ZeroMemory(_rgFieldStatePairs, sizeof(_rgFieldStatePairs));
    ZeroMemory(_rgFieldStrings, sizeof(_rgFieldStrings));
    LoadKeys();
}

GNSCredential::~GNSCredential()
{
    //delete client;
    //client = nullptr;
    if (_rgFieldStrings[SFI_PASSWORD])
    {
        size_t lenPassword = wcslen(_rgFieldStrings[SFI_PASSWORD]);
        SecureZeroMemory(_rgFieldStrings[SFI_PASSWORD], lenPassword * sizeof(*_rgFieldStrings[SFI_PASSWORD]));
    }
    for (int i = 0; i < ARRAYSIZE(_rgFieldStrings); i++)
    {
        CoTaskMemFree(_rgFieldStrings[i]);
        CoTaskMemFree(_rgCredProvFieldDescriptors[i].pszLabel);
    }
    CoTaskMemFree(_pszUserSid);
    CoTaskMemFree(_pszQualifiedUserName);
    DllRelease();

}

void GNSCredential::MyCallback(CardStatus status) {
    static CardStatus prev;

  
    //callbackLog << "event pointer " << _pCredProvCredentialEvents << endl;
 
    if (_pCredProvCredentialEvents != nullptr)
    {
        bool change =  (prev.status() != status.status());

        //callbackLog << "State changed: " << change << std::endl;
        prev = status;
        if (status.status() == CardStatus_ConnectionStatus_Authenticated && change)
        {
            //callbackLog << "Authenticated" << std::endl;
            if (client != nullptr)
            {
                client->ReadWinCreds(card_data);
            }
            _pCredProvCredentialEvents->BeginFieldUpdates();
            for (int i = comboBoxCount - 1; i >= 0; i--)
            {
                _pCredProvCredentialEvents->DeleteFieldComboBoxItem(nullptr, SFI_COMBOBOX, i);
                //AppendFieldComboBoxItem(nullptr, SFI_COMBOBOX, usernameTemp.c_str());
                comboBoxCount--;
            }
            _pCredProvCredentialEvents->EndFieldUpdates();

            if (card_data.wincreds_size() > 0)
            {

                //std::wstring_convert<std::codecvt_utf8_utf16<wchar_t>> converter;
                //std::wstring wide = L"UUID: " + converter.from_bytes(uuid.c_str());
                _pCredProvCredentialEvents->BeginFieldUpdates();
                //_pCredProvCredentialEvents->SetFieldString(nullptr, SFI_FULLNAME_TEXT, test.c_str());
                for (int i = 0; i < card_data.wincreds_size(); i++)
                {
                    std::wstring usernameTemp = GNSGRPCClient::WstringFromString(card_data.wincreds().at(i).username());
                    _pCredProvCredentialEvents->AppendFieldComboBoxItem(this, SFI_COMBOBOX, usernameTemp.c_str());
                    comboBoxCount++;
                }
                _pCredProvCredentialEvents->SetFieldString(this, SFI_FULLNAME_TEXT, L"Authenticated");
                _pCredProvCredentialEvents->SetFieldState(this, SFI_SUBMIT_BUTTON, CPFS_DISPLAY_IN_SELECTED_TILE);
                _pCredProvCredentialEvents->SetFieldState(this, SFI_COMBOBOX, CPFS_DISPLAY_IN_SELECTED_TILE);

                _pCredProvCredentialEvents->EndFieldUpdates();
            }

        }
        else if (status.status() == CardStatus_ConnectionStatus_Connected && change)
        {
            _pCredProvCredentialEvents->BeginFieldUpdates();
            _pCredProvCredentialEvents->SetFieldString(this, SFI_FULLNAME_TEXT, L"Authenticating...");
            _pCredProvCredentialEvents->EndFieldUpdates();
        }
        else if (change) 
        {
            //callbackLog << "Not authenticated" << std::endl;
            _pCredProvCredentialEvents->BeginFieldUpdates();
            for (int i = comboBoxCount - 1; i >= 0; i--)
            {
                _pCredProvCredentialEvents->DeleteFieldComboBoxItem(this, SFI_COMBOBOX, i);
                //AppendFieldComboBoxItem(nullptr, SFI_COMBOBOX, usernameTemp.c_str());
                comboBoxCount--;
            }
            _pCredProvCredentialEvents->SetFieldString(this, SFI_FULLNAME_TEXT, L"Please scan badge");
            _pCredProvCredentialEvents->SetFieldState(this, SFI_SUBMIT_BUTTON, CPFS_HIDDEN);
            _pCredProvCredentialEvents->SetFieldState(this, SFI_COMBOBOX, CPFS_HIDDEN);
            _pCredProvCredentialEvents->EndFieldUpdates();
        }

    }

}

HRESULT GNSCredential::InitializeWithoutUser(CREDENTIAL_PROVIDER_USAGE_SCENARIO cpus,
    _In_ CREDENTIAL_PROVIDER_FIELD_DESCRIPTOR const* rgcpfd,
    _In_ FIELD_STATE_PAIR const* rgfsp,
    _In_ DWORD dwFlags)
{
    HRESULT hr = S_OK;
    _cpus = cpus;
    _dwFlags = dwFlags;

    //GUID guidProvider;
    //pcpUser->GetProviderID(&guidProvider);
    //_fIsLocalUser = (guidProvider == Identity_LocalUserProvider);

    //ofstream MyFile("Cred_provider_zguid.txt");
    //MyFile << "guidProvider" << guidProvider.Data1 << " " << Identity_LocalUserProvider.Data1 << std::endl;

    //MyFile.close();

    // Copy the field descriptors for each field. This is useful if you want to vary the field
    // descriptors based on what Usage scenario the credential was created for.
    for (DWORD i = 0; SUCCEEDED(hr) && i < ARRAYSIZE(_rgCredProvFieldDescriptors); i++)
    {
        _rgFieldStatePairs[i] = rgfsp[i];
        hr = FieldDescriptorCopy(rgcpfd[i], &_rgCredProvFieldDescriptors[i]);
    }

    // Initialize the String value of all the fields.
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"GNS Credential", &_rgFieldStrings[SFI_LABEL]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"GNS Credential Provider", &_rgFieldStrings[SFI_LARGE_TEXT]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"Edit Text", &_rgFieldStrings[SFI_EDIT_TEXT]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"", &_rgFieldStrings[SFI_PASSWORD]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"Submit", &_rgFieldStrings[SFI_SUBMIT_BUTTON]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"Checkbox", &_rgFieldStrings[SFI_CHECKBOX]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"Combobox", &_rgFieldStrings[SFI_COMBOBOX]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"Launch helper window", &_rgFieldStrings[SFI_LAUNCHWINDOW_LINK]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"Hide additional controls", &_rgFieldStrings[SFI_HIDECONTROLS_LINK]);
    }
    if (SUCCEEDED(hr))
    {
        //hr = pcpUser->GetStringValue(PKEY_Identity_QualifiedUserName, &_pszQualifiedUserName);
        _pszQualifiedUserName = nullptr;
    }
    if (SUCCEEDED(hr))
    {
        PWSTR pszUserName = nullptr;
        //pcpUser->GetStringValue(PKEY_Identity_UserName, &pszUserName);

        if (pszUserName != nullptr)
        {
            wchar_t szString[256];
            StringCchPrintf(szString, ARRAYSIZE(szString), L"User Name:");
            hr = SHStrDupW(szString, &_rgFieldStrings[SFI_FULLNAME_TEXT]);
            CoTaskMemFree(pszUserName);
        }
        else
        {
            hr = SHStrDupW(L"User Name is NULL", &_rgFieldStrings[SFI_FULLNAME_TEXT]);
        }
    }
    if (SUCCEEDED(hr))
    {
        PWSTR pszDisplayName = nullptr;
        //pcpUser->GetStringValue(PKEY_Identity_DisplayName, &pszDisplayName);
        if (pszDisplayName != nullptr)
        {
            wchar_t szString[256];
            StringCchPrintf(szString, ARRAYSIZE(szString), L"Display Name: %s", pszDisplayName);
            hr = SHStrDupW(szString, &_rgFieldStrings[SFI_DISPLAYNAME_TEXT]);
            CoTaskMemFree(pszDisplayName);
        }
        else
        {
            hr = SHStrDupW(L"Display Name is NULL", &_rgFieldStrings[SFI_DISPLAYNAME_TEXT]);
        }
    }
    if (SUCCEEDED(hr))
    {
        PWSTR pszLogonStatus = nullptr;
        //pcpUser->GetStringValue(PKEY_Identity_LogonStatusString, &pszLogonStatus);
        if (pszLogonStatus != nullptr)
        {
            wchar_t szString[256];
            StringCchPrintf(szString, ARRAYSIZE(szString), L"Logon Status: %s", pszLogonStatus);
            hr = SHStrDupW(szString, &_rgFieldStrings[SFI_LOGONSTATUS_TEXT]);
            CoTaskMemFree(pszLogonStatus);
        }
        else
        {
            hr = SHStrDupW(L"Logon Status is NULL", &_rgFieldStrings[SFI_LOGONSTATUS_TEXT]);
        }
    }

    if (SUCCEEDED(hr))
    {
        //hr = pcpUser->GetSid(&_pszUserSid);
        _pszUserSid = nullptr;
    }


    return hr;
}

// Initializes one credential with the field information passed in.
// Set the value of the SFI_LARGE_TEXT field to pwzUsername.
HRESULT GNSCredential::Initialize(CREDENTIAL_PROVIDER_USAGE_SCENARIO cpus,
                                      _In_ CREDENTIAL_PROVIDER_FIELD_DESCRIPTOR const *rgcpfd,
                                      _In_ FIELD_STATE_PAIR const *rgfsp,
                                      _In_ ICredentialProviderUser *pcpUser,
                                      _In_ DWORD dwFlags)
{
    HRESULT hr = S_OK;
    _cpus = cpus;
    _dwFlags = dwFlags;

    GUID guidProvider;
    pcpUser->GetProviderID(&guidProvider);
    _fIsLocalUser = (guidProvider == Identity_LocalUserProvider);

    //ofstream MyFile("Cred_provider_zguid.txt");
    //MyFile << "guidProvider" << guidProvider.Data1 << " " << Identity_LocalUserProvider.Data1 << std::endl;

    //MyFile.close();

    // Copy the field descriptors for each field. This is useful if you want to vary the field
    // descriptors based on what Usage scenario the credential was created for.
    for (DWORD i = 0; SUCCEEDED(hr) && i < ARRAYSIZE(_rgCredProvFieldDescriptors); i++)
    {
        _rgFieldStatePairs[i] = rgfsp[i];
        hr = FieldDescriptorCopy(rgcpfd[i], &_rgCredProvFieldDescriptors[i]);
    }

    // Initialize the String value of all the fields.
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"GNS Credential", &_rgFieldStrings[SFI_LABEL]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"GNS Credential Provider", &_rgFieldStrings[SFI_LARGE_TEXT]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"Edit Text", &_rgFieldStrings[SFI_EDIT_TEXT]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"", &_rgFieldStrings[SFI_PASSWORD]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"Submit", &_rgFieldStrings[SFI_SUBMIT_BUTTON]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"Checkbox", &_rgFieldStrings[SFI_CHECKBOX]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"Combobox", &_rgFieldStrings[SFI_COMBOBOX]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"Launch helper window", &_rgFieldStrings[SFI_LAUNCHWINDOW_LINK]);
    }
    if (SUCCEEDED(hr))
    {
        hr = SHStrDupW(L"Hide additional controls", &_rgFieldStrings[SFI_HIDECONTROLS_LINK]);
    }
    if (SUCCEEDED(hr))
    {
        hr = pcpUser->GetStringValue(PKEY_Identity_QualifiedUserName, &_pszQualifiedUserName);
    }
    if (SUCCEEDED(hr))
    {
          PWSTR pszUserName;
          pcpUser->GetStringValue(PKEY_Identity_UserName, &pszUserName);

          if (pszUserName != nullptr)
          {
            wchar_t szString[256];
            StringCchPrintf(szString, ARRAYSIZE(szString), L"User Name:");
            hr = SHStrDupW(szString, &_rgFieldStrings[SFI_FULLNAME_TEXT]);
            CoTaskMemFree(pszUserName);
          }
          else
          {
            hr =  SHStrDupW(L"User Name is NULL", &_rgFieldStrings[SFI_FULLNAME_TEXT]);
          }
    }
    if (SUCCEEDED(hr))
    {
        PWSTR pszDisplayName;
        pcpUser->GetStringValue(PKEY_Identity_DisplayName, &pszDisplayName);
        if (pszDisplayName != nullptr)
        {
            wchar_t szString[256];
            StringCchPrintf(szString, ARRAYSIZE(szString), L"Display Name: %s", pszDisplayName);
            hr = SHStrDupW(szString, &_rgFieldStrings[SFI_DISPLAYNAME_TEXT]);
            CoTaskMemFree(pszDisplayName);
        }
        else
        {
            hr = SHStrDupW(L"Display Name is NULL", &_rgFieldStrings[SFI_DISPLAYNAME_TEXT]);
        }
    }
    if (SUCCEEDED(hr))
    {
        PWSTR pszLogonStatus;
        pcpUser->GetStringValue(PKEY_Identity_LogonStatusString, &pszLogonStatus);
        if (pszLogonStatus != nullptr)
        {
            wchar_t szString[256];
            StringCchPrintf(szString, ARRAYSIZE(szString), L"Logon Status: %s", pszLogonStatus);
            hr = SHStrDupW(szString, &_rgFieldStrings[SFI_LOGONSTATUS_TEXT]);
            CoTaskMemFree(pszLogonStatus);
        }
        else
        {
            hr = SHStrDupW(L"Logon Status is NULL", &_rgFieldStrings[SFI_LOGONSTATUS_TEXT]);
        }
    }

    if (SUCCEEDED(hr))
    {
        hr = pcpUser->GetSid(&_pszUserSid);
    }


    return hr;
}

// LogonUI calls this in order to give us a callback in case we need to notify it of anything.
HRESULT GNSCredential::Advise(_In_ ICredentialProviderCredentialEvents *pcpce)
{
    if (_pCredProvCredentialEvents != nullptr)
    {
        _pCredProvCredentialEvents->Release();
    }
    comboBoxCount = 0;
    //callbackLog = ofstream("Cred_provider_mycallback.txt");
    grpc::SslCredentialsOptions sslOpts;
    sslOpts.pem_root_certs = ca_crt;
    sslOpts.pem_private_key = client_key;
    sslOpts.pem_cert_chain = client_crt;
    auto creds = grpc::SslCredentials(sslOpts);
    client = new GNSGRPCClient(grpc::CreateChannel("127.0.0.1:50051", creds));
    //
    //client->SetupCallback([this](CardStatus status) { MyCallback(status); });
    client->SetupCallback(
        std::bind(&GNSCredential::MyCallback, this, std::placeholders::_1)
    );
    return pcpce->QueryInterface(IID_PPV_ARGS(&_pCredProvCredentialEvents));
}

// LogonUI calls this to tell us to release the callback.
HRESULT GNSCredential::UnAdvise()
{
    //callbackLog.close();
    if (_pCredProvCredentialEvents)
    {
        _pCredProvCredentialEvents->Release();
    }
    _pCredProvCredentialEvents = nullptr;
    delete client;
    client = nullptr;
    return S_OK;
}

// LogonUI calls this function when our tile is selected (zoomed)
// If you simply want fields to show/hide based on the selected state,
// there's no need to do anything here - you can set that up in the
// field definitions. But if you want to do something
// more complicated, like change the contents of a field when the tile is
// selected, you would do it here.
HRESULT GNSCredential::SetSelected(_Out_ BOOL *pbAutoLogon)
{

    *pbAutoLogon = FALSE;
    HRESULT hr = S_OK;
    //std::wstring test;

    //ofstream MyFile("Cred_provider_SetSelected.txt", std::ios_base::app);

    //MyFile << "client " << client << endl;
    card_data.clear_wincreds();
    if (client != nullptr)
    {
        //client->ReadUUID(test);
        client->ReadWinCreds(card_data);
        _pCredProvCredentialEvents->BeginFieldUpdates();
        _pCredProvCredentialEvents->SetFieldString(nullptr, SFI_FULLNAME_TEXT, L"Connected to GRPC service");
        _pCredProvCredentialEvents->EndFieldUpdates();
    }
    else {
        
        _pCredProvCredentialEvents->BeginFieldUpdates();
        wstring key = GNSGRPCClient::WstringFromString(client_key);
        _pCredProvCredentialEvents->SetFieldString(nullptr, SFI_FULLNAME_TEXT, L"Unable to connect to GRPC service");
        _pCredProvCredentialEvents->EndFieldUpdates();
    }

    _pCredProvCredentialEvents->BeginFieldUpdates();
    for (int i = comboBoxCount - 1; i >= 0; i--)
    {
        _pCredProvCredentialEvents->DeleteFieldComboBoxItem(nullptr, SFI_COMBOBOX, i);
        //AppendFieldComboBoxItem(nullptr, SFI_COMBOBOX, usernameTemp.c_str());
        comboBoxCount--;
    }
    _pCredProvCredentialEvents->EndFieldUpdates();
    
    //MyFile << "wincreds_size " << card_data.wincreds_size() << endl;
    if (card_data.wincreds_size() > 0)
    {
        
        //std::wstring_convert<std::codecvt_utf8_utf16<wchar_t>> converter;
        //std::wstring wide = L"UUID: " + converter.from_bytes(uuid.c_str());
        _pCredProvCredentialEvents->BeginFieldUpdates();
        _pCredProvCredentialEvents->SetFieldString(nullptr, SFI_FULLNAME_TEXT, L"Authenticated");
        for (int i = 0; i < card_data.wincreds_size(); i++)
        {
            std::wstring usernameTemp = GNSGRPCClient::WstringFromString(card_data.wincreds().at(i).username());
            _pCredProvCredentialEvents->AppendFieldComboBoxItem(nullptr, SFI_COMBOBOX, usernameTemp.c_str());
            comboBoxCount++;
        }
        _pCredProvCredentialEvents->SetFieldState(nullptr, SFI_SUBMIT_BUTTON, CPFS_DISPLAY_IN_SELECTED_TILE);
        _pCredProvCredentialEvents->SetFieldState(this, SFI_COMBOBOX, CPFS_DISPLAY_IN_SELECTED_TILE);

        _pCredProvCredentialEvents->EndFieldUpdates();
    }
    else {
        _pCredProvCredentialEvents->BeginFieldUpdates();
        for (int i = comboBoxCount - 1; i >= 0; i--)
        {
            _pCredProvCredentialEvents->DeleteFieldComboBoxItem(this, SFI_COMBOBOX, i);
            //AppendFieldComboBoxItem(nullptr, SFI_COMBOBOX, usernameTemp.c_str());
            comboBoxCount--;
        }
        _pCredProvCredentialEvents->SetFieldState(this, SFI_COMBOBOX, CPFS_HIDDEN);
        _pCredProvCredentialEvents->SetFieldString(this, SFI_FULLNAME_TEXT, L"Please scan badge");
        _pCredProvCredentialEvents->SetFieldState(this, SFI_SUBMIT_BUTTON, CPFS_HIDDEN);
        _pCredProvCredentialEvents->EndFieldUpdates();

    }



    return S_OK;
}

// Similarly to SetSelected, LogonUI calls this when your tile was selected
// and now no longer is. The most common thing to do here (which we do below)
// is to clear out the password field.
HRESULT GNSCredential::SetDeselected()
{
    HRESULT hr = S_OK;
    if (_rgFieldStrings[SFI_PASSWORD])
    {
        size_t lenPassword = wcslen(_rgFieldStrings[SFI_PASSWORD]);
        SecureZeroMemory(_rgFieldStrings[SFI_PASSWORD], lenPassword * sizeof(*_rgFieldStrings[SFI_PASSWORD]));

        CoTaskMemFree(_rgFieldStrings[SFI_PASSWORD]);
        hr = SHStrDupW(L"", &_rgFieldStrings[SFI_PASSWORD]);

        if (SUCCEEDED(hr) && _pCredProvCredentialEvents)
        {
            _pCredProvCredentialEvents->SetFieldString(this, SFI_PASSWORD, _rgFieldStrings[SFI_PASSWORD]);
        }
    }
    card_data.clear_wincreds();
    return hr;
}

// Get info for a particular field of a tile. Called by logonUI to get information
// to display the tile.
HRESULT GNSCredential::GetFieldState(DWORD dwFieldID,
                                         _Out_ CREDENTIAL_PROVIDER_FIELD_STATE *pcpfs,
                                         _Out_ CREDENTIAL_PROVIDER_FIELD_INTERACTIVE_STATE *pcpfis)
{
    HRESULT hr;

    // Validate our parameters.
    if ((dwFieldID < ARRAYSIZE(_rgFieldStatePairs)))
    {
        *pcpfs = _rgFieldStatePairs[dwFieldID].cpfs;
        *pcpfis = _rgFieldStatePairs[dwFieldID].cpfis;
        hr = S_OK;
    }
    else
    {
        hr = E_INVALIDARG;
    }
    return hr;
}

// Sets ppwsz to the string value of the field at the index dwFieldID
HRESULT GNSCredential::GetStringValue(DWORD dwFieldID, _Outptr_result_nullonfailure_ PWSTR *ppwsz)
{
    HRESULT hr;
    *ppwsz = nullptr;

    // Check to make sure dwFieldID is a legitimate index
    if (dwFieldID < ARRAYSIZE(_rgCredProvFieldDescriptors))
    {
        // Make a copy of the string and return that. The caller
        // is responsible for freeing it.
        hr = SHStrDupW(_rgFieldStrings[dwFieldID], ppwsz);
    }
    else
    {
        hr = E_INVALIDARG;
    }

    return hr;
}

// Get the image to show in the user tile
HRESULT GNSCredential::GetBitmapValue(DWORD dwFieldID, _Outptr_result_nullonfailure_ HBITMAP *phbmp)
{
    HRESULT hr;
    *phbmp = nullptr;

    if ((SFI_TILEIMAGE == dwFieldID))
    {
        HBITMAP hbmp = LoadBitmap(HINST_THISDLL, MAKEINTRESOURCE(IDB_TILE_IMAGE));
        if (hbmp != nullptr)
        {
            hr = S_OK;
            *phbmp = hbmp;
        }
        else
        {
            hr = HRESULT_FROM_WIN32(GetLastError());
        }
    }
    else
    {
        hr = E_INVALIDARG;
    }

    return hr;
}

// Sets pdwAdjacentTo to the index of the field the submit button should be
// adjacent to. We recommend that the submit button is placed next to the last
// field which the user is required to enter information in. Optional fields
// should be below the submit button.
HRESULT GNSCredential::GetSubmitButtonValue(DWORD dwFieldID, _Out_ DWORD *pdwAdjacentTo)
{
    HRESULT hr;

    if (SFI_SUBMIT_BUTTON == dwFieldID)
    {
        // pdwAdjacentTo is a pointer to the fieldID you want the submit button to
        // appear next to.
        *pdwAdjacentTo = SFI_COMBOBOX;
        hr = S_OK;
    }
    else
    {
        hr = E_INVALIDARG;
    }
    return hr;
}

// Sets the value of a field which can accept a string as a value.
// This is called on each keystroke when a user types into an edit field
HRESULT GNSCredential::SetStringValue(DWORD dwFieldID, _In_ PCWSTR pwz)
{
    HRESULT hr;

    // Validate parameters.
    if (dwFieldID < ARRAYSIZE(_rgCredProvFieldDescriptors) &&
        (CPFT_EDIT_TEXT == _rgCredProvFieldDescriptors[dwFieldID].cpft ||
        CPFT_PASSWORD_TEXT == _rgCredProvFieldDescriptors[dwFieldID].cpft))
    {
        PWSTR *ppwszStored = &_rgFieldStrings[dwFieldID];
        CoTaskMemFree(*ppwszStored);
        hr = SHStrDupW(pwz, ppwszStored);
    }
    else
    {
        hr = E_INVALIDARG;
    }

    return hr;
}

// Returns whether a checkbox is checked or not as well as its label.
HRESULT GNSCredential::GetCheckboxValue(DWORD dwFieldID, _Out_ BOOL *pbChecked, _Outptr_result_nullonfailure_ PWSTR *ppwszLabel)
{
    HRESULT hr;
    *ppwszLabel = nullptr;

    // Validate parameters.
    if (dwFieldID < ARRAYSIZE(_rgCredProvFieldDescriptors) &&
        (CPFT_CHECKBOX == _rgCredProvFieldDescriptors[dwFieldID].cpft))
    {
        *pbChecked = _fChecked;
        hr = SHStrDupW(_rgFieldStrings[SFI_CHECKBOX], ppwszLabel);
    }
    else
    {
        hr = E_INVALIDARG;
    }

    return hr;
}

// Sets whether the specified checkbox is checked or not.
HRESULT GNSCredential::SetCheckboxValue(DWORD dwFieldID, BOOL bChecked)
{
    HRESULT hr;

    // Validate parameters.
    if (dwFieldID < ARRAYSIZE(_rgCredProvFieldDescriptors) &&
        (CPFT_CHECKBOX == _rgCredProvFieldDescriptors[dwFieldID].cpft))
    {
        _fChecked = bChecked;
        hr = S_OK;
    }
    else
    {
        hr = E_INVALIDARG;
    }

    return hr;
}

// Returns the number of items to be included in the combobox (pcItems), as well as the
// currently selected item (pdwSelectedItem).
HRESULT GNSCredential::GetComboBoxValueCount(DWORD dwFieldID, _Out_ DWORD *pcItems, _Deref_out_range_(<, *pcItems) _Out_ DWORD *pdwSelectedItem)
{
    HRESULT hr;
    *pcItems = 0;
    *pdwSelectedItem = 0;

    // Validate parameters.
    if (dwFieldID < ARRAYSIZE(_rgCredProvFieldDescriptors) &&
        (CPFT_COMBOBOX == _rgCredProvFieldDescriptors[dwFieldID].cpft))
    {
        //if()

        *pcItems = 0;// (DWORD)card_data.wincreds_size();
        *pdwSelectedItem = 0;
        hr = S_OK;
    }
    else
    {
        hr = E_INVALIDARG;
    }

    return hr;
}

// Called iteratively to fill the combobox with the string (ppwszItem) at index dwItem.
HRESULT GNSCredential::GetComboBoxValueAt(DWORD dwFieldID, DWORD dwItem, _Outptr_result_nullonfailure_ PWSTR *ppwszItem)
{
    HRESULT hr;
    *ppwszItem = nullptr;

    // Validate parameters.
    if (dwFieldID < ARRAYSIZE(_rgCredProvFieldDescriptors) &&
        (CPFT_COMBOBOX == _rgCredProvFieldDescriptors[dwFieldID].cpft))
    {
        int idx = (int)dwItem;
        std::wstring username = L"NULL"; //GNSGRPCClient::WstringFromString(card_data.wincreds().at(idx).username());_dwComboIndex
        hr = SHStrDupW(username.c_str(), ppwszItem);
    }
    else
    {
        hr = E_INVALIDARG;
    }

    return hr;
}

// Called when the user changes the selected item in the combobox.
HRESULT GNSCredential::SetComboBoxSelectedValue(DWORD dwFieldID, DWORD dwSelectedItem)
{
    HRESULT hr;

    // Validate parameters.
    if (dwFieldID < ARRAYSIZE(_rgCredProvFieldDescriptors) &&
        (CPFT_COMBOBOX == _rgCredProvFieldDescriptors[dwFieldID].cpft))
    {
        _dwComboIndex = dwSelectedItem;
        hr = S_OK;
    }
    else
    {
        hr = E_INVALIDARG;
    }

    return hr;
}

// Called when the user clicks a command link.
HRESULT GNSCredential::CommandLinkClicked(DWORD dwFieldID)
{
    HRESULT hr = S_OK;

    CREDENTIAL_PROVIDER_FIELD_STATE cpfsShow = CPFS_HIDDEN;

    // Validate parameter.
    if (dwFieldID < ARRAYSIZE(_rgCredProvFieldDescriptors) &&
        (CPFT_COMMAND_LINK == _rgCredProvFieldDescriptors[dwFieldID].cpft))
    {
        HWND hwndOwner = nullptr;
        switch (dwFieldID)
        {
        case SFI_LAUNCHWINDOW_LINK:
            if (_pCredProvCredentialEvents)
            {
                _pCredProvCredentialEvents->OnCreatingWindow(&hwndOwner);
            }

            // Pop a messagebox indicating the click.
            ::MessageBox(hwndOwner, L"Command link clicked", L"Click!", 0);
            break;
        case SFI_HIDECONTROLS_LINK:
            _pCredProvCredentialEvents->BeginFieldUpdates();
            cpfsShow = _fShowControls ? CPFS_DISPLAY_IN_SELECTED_TILE : CPFS_HIDDEN;
            _pCredProvCredentialEvents->SetFieldState(nullptr, SFI_FULLNAME_TEXT, cpfsShow);
            _pCredProvCredentialEvents->SetFieldState(nullptr, SFI_DISPLAYNAME_TEXT, cpfsShow);
            _pCredProvCredentialEvents->SetFieldState(nullptr, SFI_LOGONSTATUS_TEXT, cpfsShow);
            _pCredProvCredentialEvents->SetFieldState(nullptr, SFI_CHECKBOX, cpfsShow);
            _pCredProvCredentialEvents->SetFieldState(nullptr, SFI_EDIT_TEXT, cpfsShow);
            _pCredProvCredentialEvents->SetFieldState(nullptr, SFI_COMBOBOX, cpfsShow);
            _pCredProvCredentialEvents->SetFieldString(nullptr, SFI_HIDECONTROLS_LINK, _fShowControls? L"Hide additional controls" : L"Show additional controls");
            _pCredProvCredentialEvents->EndFieldUpdates();
            _fShowControls = !_fShowControls;
            break;
        default:
            hr = E_INVALIDARG;
        }

    }
    else
    {
        hr = E_INVALIDARG;
    }

    return hr;
}

// Collect the username and password into a serialized credential for the correct usage scenario
// (logon/unlock is what's demonstrated in this sample).  LogonUI then passes these credentials
// back to the system to log on.
HRESULT GNSCredential::GetSerialization(_Out_ CREDENTIAL_PROVIDER_GET_SERIALIZATION_RESPONSE *pcpgsr,
                                            _Out_ CREDENTIAL_PROVIDER_CREDENTIAL_SERIALIZATION *pcpcs,
                                            _Outptr_result_maybenull_ PWSTR *ppwszOptionalStatusText,
                                            _Out_ CREDENTIAL_PROVIDER_STATUS_ICON *pcpsiOptionalStatusIcon)
{
    HRESULT hr = E_UNEXPECTED;
    *pcpgsr = CPGSR_NO_CREDENTIAL_NOT_FINISHED;
    *ppwszOptionalStatusText = nullptr;
    *pcpsiOptionalStatusIcon = CPSI_NONE;
    ZeroMemory(pcpcs, sizeof(*pcpcs));

    ofstream MyFile("Cred_provider_GetSerialization.txt", std::ios_base::app);
    MyFile << "GetSerilization is called" << _fIsLocalUser << " " << _dwComboIndex << " " << comboBoxCount << std::endl;

    WCHAR wsz[MAX_COMPUTERNAME_LENGTH + 1];
    DWORD cch = ARRAYSIZE(wsz);


    std::wstring _username = GNSGRPCClient::WstringFromString(card_data.wincreds().at(_dwComboIndex).username());
    std::wstring _password = GNSGRPCClient::WstringFromString(card_data.wincreds().at(_dwComboIndex).password());
    MyFile << "card name & pass :" << card_data.wincreds().at(_dwComboIndex).username() << " " << card_data.wincreds().at(_dwComboIndex).password() << endl;

    DWORD cb = 0;
    BYTE* rgb = NULL;

    //ofstream MyFile("Cred_provider_Serialization.txt", std::ios_base::app);
    //char Buffer[33];
   // _ultoa(_dwFlags, Buffer, 16);
    MyFile << "CPU: " << _cpus << endl;
    MyFile << "DFLAGS: " << std::hex<< _dwFlags << endl;

    if (GetComputerNameW(wsz, &cch))
    {
        //PWSTR pwzProtectedPassword;

        //hr = ProtectIfNecessaryAndCopyPassword(_rgFieldStrings[SFI_PASSWORD], _cpus, &pwzProtectedPassword);

        // Only CredUI scenarios should use CredPackAuthenticationBuffer.  Custom packing logic is necessary for
        // logon and unlock scenarios in order to specify the correct MessageType.
        if (CPUS_CREDUI == _cpus)
        {
            MyFile << "entering CPUS_CREDUI" << endl;
            if (_password.length() > 0)
            {
                PWSTR pwzDomainUsername = NULL;
                hr = DomainUsernameStringAlloc(wsz, &(_username)[0], &pwzDomainUsername);
                wstring wsdomainUserName(pwzDomainUsername);
                string domainUserName(wsdomainUserName.begin(), wsdomainUserName.end());

                MyFile << "Logging on DomainUsername: " << domainUserName << endl;
                if (SUCCEEDED(hr))
                {
                    if (CREDUIWIN_PACK_32_WOW & _dwFlags)
                    {
                       MyFile << "we need to CREDUIWIN_PACK_32_WOW 32 bit for credui " << endl;
                    }

                    // We use KERB_INTERACTIVE_UNLOCK_LOGON in both unlock and logon scenarios.  It contains a
                    // KERB_INTERACTIVE_LOGON to hold the creds plus a LUID that is filled in for us by Winlogon
                    // as necessary.
                    if (!CredPackAuthenticationBufferW((CREDUIWIN_PACK_32_WOW & _dwFlags) ? CRED_PACK_WOW_BUFFER : 0, pwzDomainUsername, &(_password)[0], rgb, &cb))
                    {
                        MyFile << "CredPackAuthenticationBufferW failed..why to to get more buffer " << domainUserName << endl;
                        if (ERROR_INSUFFICIENT_BUFFER == GetLastError())
                        {
                            rgb = (BYTE*)HeapAlloc(GetProcessHeap(), 0, cb);
                            if (rgb)
                            {
                                MyFile << "HeapAlloc success "<< endl;
                                // If the CREDUIWIN_PACK_32_WOW flag is set we need to return 32 bit buffers to our caller we do this by 
                                // passing CRED_PACK_WOW_BUFFER to CredPacAuthenticationBufferW.
                                if (!CredPackAuthenticationBufferW((CREDUIWIN_PACK_32_WOW & _dwFlags) ? CRED_PACK_WOW_BUFFER : 0, pwzDomainUsername, &(_password)[0], rgb, &cb))
                                {
                                    MyFile << "CredPackAuthenticationBufferW failied " << endl;
                                    HeapFree(GetProcessHeap(), 0, rgb);
                                    hr = HRESULT_FROM_WIN32(GetLastError());
                                }
                                else
                                {
                                    MyFile << "CredPackAuthenticationBufferW success " << endl;
                                    hr = S_OK;
                                }
                            }
                            else
                            {
                                MyFile << "E_OUTOFMEMORY success " << endl;
                                hr = E_OUTOFMEMORY;
                            }
                        }
                        else
                        {
                            hr = E_FAIL;
                        }
                        HeapFree(GetProcessHeap(), 0, pwzDomainUsername);
                    }
                    else
                    {
                        hr = E_FAIL;
                    }
                }
                //CoTaskMemFree(pwzProtectedPassword);
                card_data.clear_wincreds();
            }
        }
        else
        {

            KERB_INTERACTIVE_UNLOCK_LOGON kiul;


            MyFile << "Trying KerbInteractiveUnlockLogonInit: " << endl;
            // Initialize kiul with weak references to our credential.
            hr = KerbInteractiveUnlockLogonInit(wsz, &(_username)[0], &(_password)[0], _cpus, &kiul);

            if (SUCCEEDED(hr))
            {
                // We use KERB_INTERACTIVE_UNLOCK_LOGON in both unlock and logon scenarios.  It contains a
                // KERB_INTERACTIVE_LOGON to hold the creds plus a LUID that is filled in for us by Winlogon
                // as necessary.
                hr = KerbInteractiveUnlockLogonPack(kiul, &pcpcs->rgbSerialization, &pcpcs->cbSerialization);
            }
        }

        if (SUCCEEDED(hr))
        {
            MyFile << "SUCCEEDED(hr)" << endl;
            ULONG ulAuthPackage;
            hr = RetrieveNegotiateAuthPackage(&ulAuthPackage);
            if (SUCCEEDED(hr))
            {
                pcpcs->ulAuthenticationPackage = ulAuthPackage;
                pcpcs->clsidCredentialProvider = CLSID_GNS;

                // In CredUI scenarios, we must pass back the buffer constructed with CredPackAuthenticationBuffer.
                if (CPUS_CREDUI == _cpus)
                {
                    MyFile << "In CredUI scenarios, we must pass back the buffer constructed with CredPackAuthenticationBuffer." << endl;
                    pcpcs->rgbSerialization = rgb;
                    pcpcs->cbSerialization = cb;
                }

                // At this point the credential has created the serialized credential used for logon
                // By setting this to CPGSR_RETURN_CREDENTIAL_FINISHED we are letting logonUI know
                // that we have all the information we need and it should attempt to submit the 
                // serialized credential.
                *pcpgsr = CPGSR_RETURN_CREDENTIAL_FINISHED;
            }
            else
            {
                HeapFree(GetProcessHeap(), 0, rgb);
            }
        }
    }
    else
    {
        DWORD dwErr = GetLastError();
        hr = HRESULT_FROM_WIN32(dwErr);
    }

    // Close the file
    MyFile.close();
    return hr;
}

struct REPORT_RESULT_STATUS_INFO
{
    NTSTATUS ntsStatus;
    NTSTATUS ntsSubstatus;
    PWSTR     pwzMessage;
    CREDENTIAL_PROVIDER_STATUS_ICON cpsi;
};

static const REPORT_RESULT_STATUS_INFO s_rgLogonStatusInfo[] =
{
    { STATUS_LOGON_FAILURE, STATUS_SUCCESS, L"Incorrect password or username.", CPSI_ERROR, },
    { STATUS_ACCOUNT_RESTRICTION, STATUS_ACCOUNT_DISABLED, L"The account is disabled.", CPSI_WARNING },
};

// ReportResult is completely optional.  Its purpose is to allow a credential to customize the string
// and the icon displayed in the case of a logon failure.  For example, we have chosen to
// customize the error shown in the case of bad username/password and in the case of the account
// being disabled.
HRESULT GNSCredential::ReportResult(NTSTATUS ntsStatus,
                                        NTSTATUS ntsSubstatus,
                                        _Outptr_result_maybenull_ PWSTR *ppwszOptionalStatusText,
                                        _Out_ CREDENTIAL_PROVIDER_STATUS_ICON *pcpsiOptionalStatusIcon)
{
    *ppwszOptionalStatusText = nullptr;
    *pcpsiOptionalStatusIcon = CPSI_NONE;

    DWORD dwStatusInfo = (DWORD)-1;

    // Look for a match on status and substatus.
    for (DWORD i = 0; i < ARRAYSIZE(s_rgLogonStatusInfo); i++)
    {
        if (s_rgLogonStatusInfo[i].ntsStatus == ntsStatus && s_rgLogonStatusInfo[i].ntsSubstatus == ntsSubstatus)
        {
            dwStatusInfo = i;
            break;
        }
    }

    if ((DWORD)-1 != dwStatusInfo)
    {
        if (SUCCEEDED(SHStrDupW(s_rgLogonStatusInfo[dwStatusInfo].pwzMessage, ppwszOptionalStatusText)))
        {
            *pcpsiOptionalStatusIcon = s_rgLogonStatusInfo[dwStatusInfo].cpsi;
        }
    }

    // If we failed the logon, try to erase the password field.
    if (FAILED(HRESULT_FROM_NT(ntsStatus)))
    {
        if (_pCredProvCredentialEvents)
        {
            _pCredProvCredentialEvents->SetFieldString(this, SFI_PASSWORD, L"");
        }
    }

    // Since nullptr is a valid value for *ppwszOptionalStatusText and *pcpsiOptionalStatusIcon
    // this function can't fail.
    return S_OK;
}

// Gets the SID of the user corresponding to the credential.
HRESULT GNSCredential::GetUserSid(_Outptr_result_nullonfailure_ PWSTR *ppszSid)
{
    *ppszSid = nullptr;
    HRESULT hr = E_UNEXPECTED;
    if (_pszUserSid != nullptr)
    {
        hr = SHStrDupW(_pszUserSid, ppszSid);
    }
    // Return S_FALSE with a null SID in ppszSid for the
    // credential to be associated with an empty user tile.

    return hr;
}

// GetFieldOptions to enable the password reveal button and touch keyboard auto-invoke in the password field.
HRESULT GNSCredential::GetFieldOptions(DWORD dwFieldID,
                                           _Out_ CREDENTIAL_PROVIDER_CREDENTIAL_FIELD_OPTIONS *pcpcfo)
{
    *pcpcfo = CPCFO_NONE;

    if (dwFieldID == SFI_PASSWORD)
    {
        *pcpcfo = CPCFO_ENABLE_PASSWORD_REVEAL;
    }
    else if (dwFieldID == SFI_TILEIMAGE)
    {
        *pcpcfo = CPCFO_ENABLE_TOUCH_KEYBOARD_AUTO_INVOKE;
    }

    return S_OK;
}
