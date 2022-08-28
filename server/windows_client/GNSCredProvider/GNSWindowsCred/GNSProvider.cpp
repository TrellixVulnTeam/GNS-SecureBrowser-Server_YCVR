//
// THIS CODE AND INFORMATION IS PROVIDED "AS IS" WITHOUT WARRANTY OF
// ANY KIND, EITHER EXPRESSED OR IMPLIED, INCLUDING BUT NOT LIMITED TO
// THE IMPLIED WARRANTIES OF MERCHANTABILITY AND/OR FITNESS FOR A
// PARTICULAR PURPOSE.
//
// Copyright (c) Microsoft Corporation. All rights reserved.
//
// GNSProvider implements ICredentialProvider, which is the main
// interface that logonUI uses to decide which tiles to display.
// In this sample, we will display one tile that uses each of the nine
// available UI controls.

#include <initguid.h>
#include "GNSProvider.h"
#include "GNSCredential.h"
#include "guid.h"
#include <iostream>
#include <fstream>
using namespace std;
/*
    long                                    _cRef;            // Used for reference counting.
    GNSCredential                       *_pCredential;    // SampleV2Credential
    bool                                    _fRecreateEnumeratedCredentials;
    CREDENTIAL_PROVIDER_USAGE_SCENARIO      _cpus;
    ICredentialProviderUserArray            *_pCredProviderUserArray;
    DWORD                                 _dwCredUIFlags;   // The flags representing the Credui Options*/
GNSProvider::GNSProvider():
    _cRef(1),
    _pCredential(nullptr),
    _fRecreateEnumeratedCredentials(true),
    _pCredProviderUserArray(nullptr),
    _dwCredUIFlags(0)
{
    DllAddRef();
}

GNSProvider::~GNSProvider()
{
    if (_pCredential != nullptr)
    {
        _pCredential->Release();
        _pCredential = nullptr;
    }
    if (_pCredProviderUserArray != nullptr)
    {
        _pCredProviderUserArray->Release();
        _pCredProviderUserArray = nullptr;
    }

    DllRelease();
}

// SetUsageScenario is the provider's cue that it's going to be asked for tiles
// in a subsequent call.
HRESULT GNSProvider::SetUsageScenario(
    CREDENTIAL_PROVIDER_USAGE_SCENARIO cpus,
    DWORD dwFlags)
{
        HRESULT hr;

        //ofstream MyFile("Cred_provider_log.txt", std::ios_base::app);
        //MyFile << "SetUsageScenario CPU: " << cpus << endl;
        //MyFile << "SetUsageScenario dwFlags: " << std::hex << dwFlags << endl;
    _cpus = cpus;
    if (cpus == CPUS_CREDUI)
    {
        
        //0x100 = CREDUIWIN_ENUMERATE_ADMINS
        _dwCredUIFlags = dwFlags;  // currently the only flags ever passed in are only valid for the credui scenario
      
    }
    _fRecreateEnumeratedCredentials = true;

    // unlike SampleCredentialProvider, we're not going to enumerate here.  Instead, we'll store off the info
    // and then we'll wait for GetCredentialCount to enumerate.  That way we'll know at enumeration time
    // whether we have a SetSerialization cred to deal with.  That's a bit more important in the credUI case
    // than the logon case (although even in the logon case you could choose to only enumerate the SetSerialization
    // credential if there is one -- that's what the built-in password provider does).
    switch (cpus)
    {
    case CPUS_LOGON:
    case CPUS_UNLOCK_WORKSTATION:
    case CPUS_CREDUI:
        hr = S_OK;
        break;

    case CPUS_CHANGE_PASSWORD:
        hr = E_NOTIMPL;
        break;

    default:
        hr = E_INVALIDARG;
        break;
    }

    return hr;
}

// SetSerialization takes the kind of buffer that you would normally return to LogonUI for
// an authentication attempt.  It's the opposite of ICredentialProviderCredential::GetSerialization.
// GetSerialization is implement by a credential and serializes that credential.  Instead,
// SetSerialization takes the serialization and uses it to create a tile.
//
// SetSerialization is called for two main scenarios.  The first scenario is in the credui case
// where it is prepopulating a tile with credentials that the user chose to store in the OS.
// The second situation is in a remote logon case where the remote client may wish to
// prepopulate a tile with a username, or in some cases, completely populate the tile and
// use it to logon without showing any UI.
//
// If you wish to see an example of SetSerialization, please see either the SampleCredentialProvider
// sample or the SampleCredUICredentialProvider sample.  [The logonUI team says, "The original sample that
// this was built on top of didn't have SetSerialization.  And when we decided SetSerialization was
// important enough to have in the sample, it ended up being a non-trivial amount of work to integrate
// it into the main sample.  We felt it was more important to get these samples out to you quickly than to
// hold them in order to do the work to integrate the SetSerialization changes from SampleCredentialProvider
// into this sample.]
HRESULT GNSProvider::SetSerialization(
    _In_ CREDENTIAL_PROVIDER_CREDENTIAL_SERIALIZATION const * pcpcs)
{
    HRESULT hr = E_INVALIDARG;


    ofstream MyFile("Cred_provider_log.txt", std::ios_base::app);
    MyFile << "SetSerialization is called" << endl;

    if ((CLSID_GNS == pcpcs->clsidCredentialProvider) || (CPUS_CREDUI == _cpus))
    {
        // Get the current AuthenticationPackageID that we are supporting
        ULONG ulNegotiateAuthPackage;
        hr = RetrieveNegotiateAuthPackage(&ulNegotiateAuthPackage);

        MyFile << "CPUS_CREDUI is detected" << endl;

        if (SUCCEEDED(hr))
        {
            if (CPUS_CREDUI == _cpus)
            {
                if (CREDUIWIN_IN_CRED_ONLY & _dwCredUIFlags)
                {
                    // If we are being told to enumerate only the incoming credential, we must not return
                    // success unless we can enumerate it.  We'll set hr to failure here and let it be
                    // overridden if the enumeration logic below succeeds.
                    hr = E_INVALIDARG;
                }
                else if (_dwCredUIFlags & CREDUIWIN_AUTHPACKAGE_ONLY)
                {
                    if (ulNegotiateAuthPackage == pcpcs->ulAuthenticationPackage)
                    {
                        // In the credui case, SetSerialization should only ever return S_OK if it is able to serialize the input cred.
                        // Unfortunately, SetSerialization had to be overloaded to indicate whether or not it will be able to GetSerialization 
                        // for the specific Auth Package that is being requested for CREDUIWIN_AUTHPACKAGE_ONLY to work, so when that flag is 
                        // set, it should return S_FALSE unless it is ALSO able to serialize the input cred, then it can return S_OK.
                        // So in this case, we can set it to be S_FALSE because we support the authpackage, and then if we
                        // can serialize the input cred, it will get overwritten with S_OK.
                        hr = S_FALSE;
                    }
                    else
                    {
                        //we don't support this auth package, so we want to let logonUI know that by failing
                        hr = E_INVALIDARG;
                    }
                }
            }

            if ((ulNegotiateAuthPackage == pcpcs->ulAuthenticationPackage) &&
                (0 < pcpcs->cbSerialization && pcpcs->rgbSerialization))
            {
                KERB_INTERACTIVE_UNLOCK_LOGON* pkil = (KERB_INTERACTIVE_UNLOCK_LOGON*)pcpcs->rgbSerialization;
                if (KerbInteractiveLogon == pkil->Logon.MessageType)
                {
                    // If there isn't a username, we can't serialize or create a tile for this credential.
                    if (0 < pkil->Logon.UserName.Length && pkil->Logon.UserName.Buffer)
                    {
                        if ((CPUS_CREDUI == _cpus) && (CREDUIWIN_PACK_32_WOW & _dwCredUIFlags))
                        {
                            BYTE* rgbNativeSerialization;
                            DWORD cbNativeSerialization;
                            if (SUCCEEDED(KerbInteractiveUnlockLogonRepackNative(pcpcs->rgbSerialization, pcpcs->cbSerialization, &rgbNativeSerialization, &cbNativeSerialization)))
                            {
                                KerbInteractiveUnlockLogonUnpackInPlace((PKERB_INTERACTIVE_UNLOCK_LOGON)rgbNativeSerialization, cbNativeSerialization);

                                _pkiulSetSerialization = (PKERB_INTERACTIVE_UNLOCK_LOGON)rgbNativeSerialization;
                                hr = S_OK;
                            }
                        }
                        else
                        {
                            BYTE* rgbSerialization;
                            rgbSerialization = (BYTE*)HeapAlloc(GetProcessHeap(), 0, pcpcs->cbSerialization);
                            HRESULT hrCreateCred = rgbSerialization ? S_OK : E_OUTOFMEMORY;

                            if (SUCCEEDED(hrCreateCred))
                            {
                                CopyMemory(rgbSerialization, pcpcs->rgbSerialization, pcpcs->cbSerialization);
                                KerbInteractiveUnlockLogonUnpackInPlace((KERB_INTERACTIVE_UNLOCK_LOGON*)rgbSerialization, pcpcs->cbSerialization);

                                if (_pkiulSetSerialization)
                                {
                                    HeapFree(GetProcessHeap(), 0, _pkiulSetSerialization);
                                }
                                _pkiulSetSerialization = (KERB_INTERACTIVE_UNLOCK_LOGON*)rgbSerialization;
                                if (SUCCEEDED(hrCreateCred))
                                {
                                    // we allow success to override the S_FALSE for the CREDUIWIN_AUTHPACKAGE_ONLY, but
                                    // failure to create the cred shouldn't override that we can still handle
                                    // the auth package
                                    hr = hrCreateCred;
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    return hr;
}

// Called by LogonUI to give you a callback.  Providers often use the callback if they
// some event would cause them to need to change the set of tiles that they enumerated.
HRESULT GNSProvider::Advise(
    _In_ ICredentialProviderEvents * pcpe,
    _In_ UINT_PTR upAdviseContext)
{
    UNREFERENCED_PARAMETER(pcpe);
    UNREFERENCED_PARAMETER(upAdviseContext);
    return E_NOTIMPL;
}

// Called by LogonUI when the ICredentialProviderEvents callback is no longer valid.
HRESULT GNSProvider::UnAdvise()
{
    return E_NOTIMPL;
}

// Called by LogonUI to determine the number of fields in your tiles.  This
// does mean that all your tiles must have the same number of fields.
// This number must include both visible and invisible fields. If you want a tile
// to have different fields from the other tiles you enumerate for a given usage
// scenario you must include them all in this count and then hide/show them as desired
// using the field descriptors.
HRESULT GNSProvider::GetFieldDescriptorCount(
    _Out_ DWORD *pdwCount)
{
    *pdwCount = SFI_NUM_FIELDS;
    return S_OK;
}

// Gets the field descriptor for a particular field.
HRESULT GNSProvider::GetFieldDescriptorAt(
    DWORD dwIndex,
    _Outptr_result_nullonfailure_ CREDENTIAL_PROVIDER_FIELD_DESCRIPTOR **ppcpfd)
{
    HRESULT hr;
    *ppcpfd = nullptr;

    // Verify dwIndex is a valid field.
    if ((dwIndex < SFI_NUM_FIELDS) && ppcpfd)
    {
        hr = FieldDescriptorCoAllocCopy(s_rgCredProvFieldDescriptors[dwIndex], ppcpfd);
    }
    else
    {
        hr = E_INVALIDARG;
    }

    return hr;
}

// Sets pdwCount to the number of tiles that we wish to show at this time.
// Sets pdwDefault to the index of the tile which should be used as the default.
// The default tile is the tile which will be shown in the zoomed view by default. If
// more than one provider specifies a default the last used cred prov gets to pick
// the default. If *pbAutoLogonWithDefault is TRUE, LogonUI will immediately call
// GetSerialization on the credential you've specified as the default and will submit
// that credential for authentication without showing any further UI.
HRESULT GNSProvider::GetCredentialCount(
    _Out_ DWORD *pdwCount,
    _Out_ DWORD *pdwDefault,
    _Out_ BOOL *pbAutoLogonWithDefault)
{
    *pdwDefault = CREDENTIAL_PROVIDER_NO_DEFAULT;
    *pbAutoLogonWithDefault = FALSE;

    if (_fRecreateEnumeratedCredentials)
    {
        _fRecreateEnumeratedCredentials = false;
        _ReleaseEnumeratedCredentials();
        _CreateEnumeratedCredentials();
    }

    *pdwCount = 1;
    return S_OK;
}

// Returns the credential at the index specified by dwIndex. This function is called by logonUI to enumerate
// the tiles.
HRESULT GNSProvider::GetCredentialAt(
    DWORD dwIndex,
    _Outptr_result_nullonfailure_ ICredentialProviderCredential **ppcpc)
{
    HRESULT hr = E_INVALIDARG;
    *ppcpc = nullptr;

    if ((dwIndex == 0) && ppcpc && _pCredential != nullptr)
    {
        hr = _pCredential->QueryInterface(IID_PPV_ARGS(ppcpc));
    }
    else {
        hr = E_INVALIDARG;
    }
    return hr;
}

// This function will be called by LogonUI after SetUsageScenario succeeds.
// Sets the User Array with the list of users to be enumerated on the logon screen.
HRESULT GNSProvider::SetUserArray(_In_ ICredentialProviderUserArray *users)
{
    if (_pCredProviderUserArray)
    {
        _pCredProviderUserArray->Release();
    }
    _pCredProviderUserArray = users;
    _pCredProviderUserArray->AddRef();
    _fRecreateEnumeratedCredentials = true;
    return S_OK;
}

void GNSProvider::_CreateEnumeratedCredentials()
{
    switch (_cpus)
    {
    case CPUS_CREDUI:
    case CPUS_LOGON:
    case CPUS_UNLOCK_WORKSTATION:
        {
            _EnumerateCredentials();
            break;
        }
    default:
        break;
    }
}

void GNSProvider::_ReleaseEnumeratedCredentials()
{
    if (_pCredential != nullptr)
    {
        _pCredential->Release();
        _pCredential = nullptr;
    }
}

HRESULT GNSProvider::_EnumerateCredentials()
{
    HRESULT hr = E_UNEXPECTED;
    if (_pCredProviderUserArray != nullptr)
    {
        DWORD dwUserCount;
        _pCredProviderUserArray->GetCount(&dwUserCount);
        if (dwUserCount > 0)
        {
            ICredentialProviderUser *pCredUser;
            hr = _pCredProviderUserArray->GetAt(0, &pCredUser);
            if (SUCCEEDED(hr))
            {
                _pCredential = new(std::nothrow) GNSCredential();
                if (_pCredential != nullptr)
                {
                    hr = _pCredential->Initialize(_cpus, s_rgCredProvFieldDescriptors, s_rgFieldStatePairs, pCredUser,_dwCredUIFlags);
                    if (FAILED(hr))
                    {
                        _pCredential->Release();
                        _pCredential = nullptr;
                    }
                }
                else
                {
                    hr = E_OUTOFMEMORY;
                }
                pCredUser->Release();
            }
        }
    }
    if (_pCredential == nullptr)
    {
        _pCredential = new(std::nothrow) GNSCredential();
        if (_pCredential != nullptr)
        {
            hr = _pCredential->InitializeWithoutUser(_cpus, s_rgCredProvFieldDescriptors, s_rgFieldStatePairs, _dwCredUIFlags);
            if (FAILED(hr))
            {
                _pCredential->Release();
                _pCredential = nullptr;
            }
        }
        else
        {
            hr = E_OUTOFMEMORY;
        }
    }
    return hr;
}

// Boilerplate code to create our provider.
HRESULT GNS_CreateInstance(_In_ REFIID riid, _Outptr_ void **ppv)
{
    HRESULT hr;
    GNSProvider *pProvider = new(std::nothrow) GNSProvider();
    if (pProvider)
    {
        hr = pProvider->QueryInterface(riid, ppv);
        pProvider->Release();
    }
    else
    {
        hr = E_OUTOFMEMORY;
    }
    return hr;
}
