#pragma once
#include "..\GNSGRPCLib\GNSGRPCLib.h"
#include <msclr\marshal_cppstd.h>
using namespace System;
using namespace std;
namespace GNSGRPCNET
{
    public enum class ManagedCardStatus {
        /// <summary>
        /// Badge is not connected or we haven't scan fingerprint
        /// </summary>
        NotReady,
        /// <summary>
        /// Card is connected and we are going to attempt to authenticate
        /// </summary>
        Connected,
        /// <summary>
        /// Completed handshake and card is ready for use
        /// </summary>
        Ready
    };

    delegate void ManagedCallbackHandler(ManagedCardStatus status);

	public ref class GNSClient
	{
    private:
        GNSGRPCClient* _nativePtr;
        ManagedCallbackHandler^ cb;
        void Callback(GNSRPC::CardStatus status);
    public:

        void Init();
        void SetupCallback(ManagedCallbackHandler^ handler);
        GNSClient() {

        }
        !GNSClient() {
            if (_nativePtr != nullptr) {
                delete _nativePtr;
                _nativePtr = nullptr;
            }
        };

	};

}

