/************************************************************************
Copyright (C) 2020 - Global Net Solutions Inc.
All rights reserved.
Any use of this code without a written consent from GNS Inc is prohibited.
@Author: Anthony Fanous
*************************************************************************/

#include <string.h>
#include <openssl/evp.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" 
{
#endif

/* ==================================================================================
Function: handleErrors
Description: Function that asserts and exits when errors are encountered
================================================================================== */
void handleErrors(const char* reason);

/* ==================================================================================
Function: load_from_binary_file
Description: Function to populate an ASCII array from a binary file. Only used for 
demonstration and should not be used on target
================================================================================== */
//void load_from_binary_file(const char* filename, unsigned char arr[], int arr_length);

/* ==================================================================================
Function: generate_pubkey_from_privkey
Description: Function to generate the public key in (x,y) format from private key
================================================================================== */
void generate_pubkey_from_privkey(const char* privkeyfilename, unsigned char* hex_x, unsigned char* hex_y, bool isDebug);

/* ==================================================================================
Function: generate_challenge
Description: generate a random challenge (raw challenge) and compute its digest in 
hex to be sent to STSAFE on VCOM
================================================================================== */
void generate_challenge(unsigned char* raw_challenge, unsigned char* digest_hex, bool isDebug);

/* ==================================================================================
Function: ECDSA_dss_verify_signature
Description: Verify the signed challenge received from STSAFE against the raw_
challenge sent given the public key extracted from the certificate
================================================================================== */
void ECDSA_dss_verify_signature(int *result,const unsigned char* s_challenge, const unsigned char* raw_challenge, const unsigned char* cert, const int cert_len);

/* ==================================================================================
Function: sign_challenge
Description: sign the challenge received from STSAFE, and compute the (r,s) of the 
signature digest
================================================================================== */
void sign_challenge(const unsigned char* challenge, const size_t challenge_size, const char* privkeyfilename, char* sign_r_hex, char* sign_s_hex, bool isDebug);

/* ==================================================================================
Function: ECDH
Description: Elliptic Key Diffie Helman Key establishmentComputes the shared secret 
from the ephemeral key received from STSAFE. Then computes the AES key and IV to be 
used for encoding and decoding
===================================================================================== */
void ECDH(const unsigned char* eph_key, const char* privkeyfilename, unsigned char* encrypt_key, unsigned char* encrypt_iv);

/* ==================================================================================
Function: AES128_Encrypt
Description: AES_128 Encryption of data given key and IV
===================================================================================== */
void AES128_Encrypt(const unsigned char* data, const int data_len, const unsigned char* encrypt_key, const unsigned char* encrypt_iv, unsigned char* encrypted_data_hex, 
	int* encrypted_data_hex_len, bool isDebug);

/* ==================================================================================
Function: AES128_Decrypt
Description: AES_128 Decryption of data given key and IV
===================================================================================== */
void AES128_Decrypt(const unsigned char* encrypted_data, const int encrypted_data_len, const unsigned char* encrypt_key, const unsigned char* encrypt_iv, 
	unsigned char* decrypted_data, int* decrypted_data_len, bool isDebug);


// End extern "C"
#ifdef __cplusplus
}
#endif
