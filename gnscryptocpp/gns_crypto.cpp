/************************************************************************
Copyright (C) 2020 - Global Net Solutions Inc.
All rights reserved.
Any use of this code without a written consent from GNS Inc is prohibited.
@Author: Anthony Fanous
*************************************************************************/
#include <fstream>
#include <sstream>
#include <vector>
#include "gns_crypto.h"
#include "conversions.h"
#include <openssl/evp.h>
#include <openssl/pem.h>
#include <openssl/rand.h>
#include <string>
#include <iostream>

void handleErrors(const char* reason) {
  printf(reason);
  printf("\n");
}

void load_from_binary_file(const char* filename, unsigned char arr[], int arr_length)
{

}

void generate_pubkey_from_privkey(const char* privkeyfilename, unsigned char* hex_x, unsigned char* hex_y, bool isDebug)
{
	const char* pwd = "SIANA_DEV";
	//Init the algos
	OpenSSL_add_all_algorithms();

	//Read private key from PEM
	BIO* bio = BIO_new(BIO_s_file());
	BIO_read_filename(bio, privkeyfilename);
	EVP_PKEY* privkey_ = PEM_read_bio_PrivateKey(bio, NULL, NULL, (void*)pwd);
	if (privkey_ == NULL)
          handleErrors("Can't extract private key");

	//convert private key into EC format
	EC_KEY* privkey = EVP_PKEY_get1_EC_KEY(privkey_); //convert the private key in EVP format to EC format

	//convert private key into BIGNUM
	const BIGNUM* bn = EC_KEY_get0_private_key(privkey);

	//generate corresponding pubkey
	BN_CTX* ctx = BN_CTX_new();

	EC_KEY* pubkey = EC_KEY_new_by_curve_name(NID_X9_62_prime256v1); //secp256r1 curve - referred as prime256v1
	if (pubkey == NULL)
          handleErrors("EC_KEY_new_by_curve_name error");
	EC_POINT* ec_point = EC_POINT_new(EC_KEY_get0_group(pubkey));
	if (1 != EC_POINT_mul(EC_KEY_get0_group(pubkey), ec_point, bn, NULL, NULL, ctx))
          handleErrors("EC_POINT_mul");
	if (1 != EC_KEY_set_public_key(pubkey, ec_point))
          handleErrors("EC_KEY_set_public_key");

	BIGNUM* x = BN_new();
	BIGNUM* y = BN_new();
	//Get pubkey in  (x,y) format
	if (1 != EC_POINT_get_affine_coordinates_GFp(EC_KEY_get0_group(pubkey), ec_point, x, y, ctx))
          handleErrors("EC_POINT_get_affine_coordinates_GFp");
	
	char* hex_x_temp = BN_bn2hex(x);
	char* hex_y_temp = BN_bn2hex(y);

	for (int i = 0; i < 64; i++)
	{
		hex_x[i] = hex_x_temp[i];
		hex_y[i] = hex_y_temp[i];
	}

	BIO_free(bio);
	EVP_PKEY_free(privkey_);
	EC_KEY_free(privkey);
	EC_KEY_free(pubkey);
	BN_CTX_free(ctx);
	BN_free(x);
	BN_free(y);
	EC_POINT_free(ec_point);
}

void generate_challenge(unsigned char* raw_challenge, unsigned char* digest_hex, bool isDebug)
{
	if (1 != RAND_bytes(raw_challenge, 256))
		handleErrors("raw_challenge");

	//std::string r = "06ec13d01797a8b28eb4e1e5c0fb15c584ccf3c37ed48fe0e6fb4014ca6c751d23967fdd6c8fda869024f7387cfcd420a735e54a7b052709e24448ef39719431110b684872c9c9942cf66e46c3ced642b6be5cc59b9be66c2b4e9ee419dbc3a835ed66c497452e8bb19d01025675ee1b3ca7847ecdada4bf8d53944660d88ce308534adb35cd44045f49f61f233948712ebbf760ae57a17e8155be81a7a3a3a063484a5e5b14a3ebe07fd889a40a76438709ab3121883496b9f182dd6a78b2f5d3528b7220a3e4368e54fbd58b9e260635a5b3f42b3c617a7c3017fc336060ace764dd4cdf8a37d245b78973132e70f9bc8387f7e39ec91a601dcb9798c5ce1e";
	//HEX::Hex_to_byte(r, raw_challenge);

	unsigned char* digest = (unsigned char*)OPENSSL_malloc(EVP_MD_size(EVP_sha256()));
	unsigned int x = 32; //length of digest in ASCII
	unsigned int* digest_len = &x;

	EVP_MD_CTX* mdctx;
	mdctx = EVP_MD_CTX_create();

	if (1 != EVP_DigestInit_ex(mdctx, EVP_sha256(), NULL))
          handleErrors("EVP_DigestInit_ex");

	if (1 != EVP_DigestUpdate(mdctx, raw_challenge, 256))
          handleErrors("EVP_DigestUpdate");

	if (1 != EVP_DigestFinal_ex(mdctx, digest, digest_len))
          handleErrors("EVP_DigestFinal_ex");
	
	std::string str_out;
	HEX::byte_to_Hex(digest, 32, str_out);
	for (int i = 0; i < 64; i++)
		digest_hex[i] = static_cast<unsigned char>(str_out[i]);

	OPENSSL_free(digest);
	EVP_MD_CTX_destroy(mdctx);
	EVP_cleanup();
}

void ECDSA_dss_verify_signature(int *result,const unsigned char* s_challenge, const unsigned char* raw_challenge, const unsigned char* cert, const int cert_len)
{
	X509* x509 = d2i_X509(NULL, &cert, cert_len);

	//Extract public key
	EVP_PKEY* pubkey = X509_get_pubkey(x509);
	if (pubkey == nullptr) {
		*result = 0;
		return;
	}
	X509_free(x509);

	ECDSA_SIG* ec_sig = ECDSA_SIG_new();


#if OPENSSL_VERSION_NUMBER >= 0x10100000L
	BIGNUM* r = BN_new();
	BIGNUM* s = BN_new();
	if (NULL == BN_bin2bn(s_challenge, 32, r)) {
          handleErrors("BN_bin2bn");
	}

	if (NULL == BN_bin2bn(s_challenge + 32, 32, s)) {
          handleErrors("BN_bin2bn");
	}
	ECDSA_SIG_set0(ec_sig, r, s);
#else
	if (NULL == BN_bin2bn(s_challenge, 32, ec_sig->r)) {
          handleErrors("BN_bin2bn");
	}

	if (NULL == BN_bin2bn(s_challenge + 32, 32, ec_sig->s)) {
          handleErrors("BN_bin2bn");
	}


#endif

	int sig_size = i2d_ECDSA_SIG(ec_sig, NULL);

#if defined(_WIN64) || defined(_WIN32)  
	//std::cout << "sig_size = " << sig_size << std::endl;
#endif
 	unsigned char* sig_bytes = (unsigned char*)malloc(sig_size);
	unsigned char* p;
	memset(sig_bytes, 6, sig_size);
	p = sig_bytes;
	int new_sig_size = i2d_ECDSA_SIG(ec_sig, &p); // The value of p is now sig_bytes + sig_size, and the signature in DER format resides at sig_bytes

	EVP_MD_CTX* ctx = EVP_MD_CTX_create();

	if (1 != EVP_DigestVerifyInit(ctx, NULL, EVP_sha256(), NULL, pubkey))
          handleErrors("BN_bin2bn");
	if (1 != EVP_DigestVerifyUpdate(ctx, raw_challenge, 256))
          handleErrors("BN_bin2bn");
	int AuthStatus = EVP_DigestVerifyFinal(ctx, &sig_bytes[0], sig_size);
	EVP_MD_CTX_destroy(ctx);

	free(sig_bytes);
	ECDSA_SIG_free(ec_sig);
	EVP_PKEY_free(pubkey);
	if(AuthStatus == 1)
	{
		*result = 1;
	}
	else{
		*result = 0;
	}
	//return (AuthStatus == 1);
	;
}

void sign_challenge(const unsigned char* challenge, const size_t challenge_size, const char* privkeyfilename, char* sign_r_hex, char* sign_s_hex, bool isDebug)
{
	const char* pwd = "SIANA_DEV";
	//Init the password.
	OpenSSL_add_all_algorithms();

	BIO* bio = BIO_new(BIO_s_file());
	BIO_read_filename(bio, privkeyfilename);
	EVP_PKEY* privkey = PEM_read_bio_PrivateKey(bio, NULL, NULL, (void*)pwd);
	if (privkey == NULL)
          handleErrors("PEM_read_bio_PrivateKey");

	ECDSA_SIG* ec_sig = ECDSA_SIG_new();
	EVP_MD_CTX* ctx;
	ctx = EVP_MD_CTX_create();

	if (1 != EVP_DigestSignInit(ctx, NULL, EVP_sha256(), NULL, privkey))
          handleErrors("EVP_DigestSignInit");
	if (1 != EVP_DigestSignUpdate(ctx, challenge, challenge_size))
          handleErrors("EVP_DigestSignUpdate");
	size_t req = 0; //required buffer length
	if (1 != EVP_DigestSignFinal(ctx, NULL, &req))
          handleErrors("EVP_DigestSignFinal");
	unsigned char* s_challenge = (unsigned char*)malloc(req);
	if (1 != EVP_DigestSignFinal(ctx, &s_challenge[0], &req))
          handleErrors("EVP_DigestSignFinal");
	const unsigned char* s_challenge2 = s_challenge;
	ECDSA_SIG* ec_sig_decode = d2i_ECDSA_SIG(NULL, &s_challenge2, (long)req);

#if OPENSSL_VERSION_NUMBER >= 0x10100000L

	char* sign_r = BN_bn2hex(ECDSA_SIG_get0_r(ec_sig_decode));
	char* sign_s = BN_bn2hex(ECDSA_SIG_get0_s(ec_sig_decode));
#else
	char* sign_r = BN_bn2hex(ec_sig_decode->r);
	char* sign_s = BN_bn2hex(ec_sig_decode->s);

#endif
	for (int i = 0; i < 64; i++)
	{
		sign_r_hex[i] = sign_r[i];
		sign_s_hex[i] = sign_s[i];
	}

	free(s_challenge);
	BIO_free(bio);
	EVP_PKEY_free(privkey);
	ECDSA_SIG_free(ec_sig);
	ECDSA_SIG_free(ec_sig_decode);
	EVP_cleanup();
	EVP_MD_CTX_destroy(ctx);
	//ECDSA_SIG_free(ec_sig_decode);
}

void ECDH(const unsigned char* eph_key, const char* privkeyfilename, unsigned char* encrypt_key, unsigned char* encrypt_iv)
{
	const char* pwd = "SIANA_DEV";
	//Init the algos
	OpenSSL_add_all_algorithms();

	BIO* bio = BIO_new(BIO_s_file());
	BIO_read_filename(bio, privkeyfilename);
	EVP_PKEY* privkey_ = PEM_read_bio_PrivateKey(bio, NULL, NULL, (void*)pwd);
	if (privkey_ == NULL)
          handleErrors("ECDH: PEM_read_bio_PrivateKey");
	//*encrypt_key = (unsigned char)OPENSSL_malloc(16);
	//*encrypt_iv = (unsigned char)OPENSSL_malloc(16);
	
	EC_KEY* pubkey = EC_KEY_new_by_curve_name(NID_X9_62_prime256v1); //secp256r1 curve - referred as prime256v1
	if(pubkey == NULL) handleErrors("ECDH: EC_KEY_new_by_curve_name");
	std::string str_out;
	HEX::byte_to_Hex((void*)eph_key, 65, str_out);
	const char* str_out_hex = str_out.c_str();
	EC_POINT* ec_point = EC_POINT_hex2point(EC_KEY_get0_group(pubkey), str_out_hex, NULL, NULL);
	if (NULL == ec_point)
          handleErrors("ECDH: EC_POINT_hex2point");
	//std::cout << "Signature: " << str_out << std::endl;

	if (1 != EC_KEY_set_public_key(pubkey, ec_point))
          handleErrors("ECDH: EC_KEY_set_public_key");

	/* Calculate the size of the buffer for the shared secret */
	int field_size = EC_GROUP_get_degree(EC_KEY_get0_group(pubkey));
	int secret_len = (field_size + 7) / 8;

	/* Allocate the memory for the shared secret */
	unsigned char* shared_secret = (unsigned char*)OPENSSL_malloc(32);
	EC_KEY* privkey = EVP_PKEY_get1_EC_KEY(privkey_); //convert the private key in EVP format to EC format
	if (NULL == privkey)
          handleErrors("ECDH: EVP_PKEY_get1_EC_KEY");

	/* Derive the shared secret */
	int key_size = ECDH_compute_key(shared_secret, secret_len, ec_point, privkey, NULL);
	if (key_size == -1)
          handleErrors("ECDH: ECDH_compute_key");

	/* Compute digest of the shared secret - called kdf_hash */
	unsigned char* kdf_hash = (unsigned char*)OPENSSL_malloc(EVP_MD_size(EVP_sha256()));
	unsigned int x = EVP_MD_size(EVP_sha256()); //length of digest in ASCII
	unsigned int* kdf_hash_len = &x;

	EVP_MD_CTX* mdctx;
	mdctx = EVP_MD_CTX_create();

	if (1 != EVP_DigestInit_ex(mdctx, EVP_sha256(), NULL))
          handleErrors("");

	if (1 != EVP_DigestUpdate(mdctx, shared_secret, 32))
          handleErrors("");

	if (1 != EVP_DigestFinal_ex(mdctx, kdf_hash, kdf_hash_len))
          handleErrors("");
	EVP_MD_CTX_destroy(mdctx);

	/* Compute digest of the [kdf_hash, shared secret]  - called kdf_hash2 */
	unsigned char* kdf_hash2 = (unsigned char*)OPENSSL_malloc(EVP_MD_size(EVP_sha256()));
	x = EVP_MD_size(EVP_sha256()); //length of digest in ASCII
	unsigned int* kdf_hash2_len = &x;

	unsigned char* shared_secret_app = (unsigned char*)OPENSSL_malloc(64);
	for (int i = 0; i < 32; i++)
		shared_secret_app[i] = kdf_hash[i];
	for(int i=32 ;i < 64; i++)
		shared_secret_app[i] = shared_secret[i-32];
	OPENSSL_free(kdf_hash);

	EVP_MD_CTX* mdctx2;
	mdctx2 = EVP_MD_CTX_create();

	if (1 != EVP_DigestInit_ex(mdctx2, EVP_sha256(), NULL))
          handleErrors("");

	if (1 != EVP_DigestUpdate(mdctx2, shared_secret_app, 64))
          handleErrors("");

	if (1 != EVP_DigestFinal_ex(mdctx2, kdf_hash2, kdf_hash2_len))
          handleErrors("");
	OPENSSL_free(shared_secret_app);
	OPENSSL_free(shared_secret);
	EVP_MD_CTX_destroy(mdctx2);

	for (int i = 0; i < 16; i++) //encrypt_key is the first 16 bytes of kdf_hash2
		encrypt_key[i] = kdf_hash2[i];
	for (int i = 0; i < 16; i++) //encrypt_iv is the last 16 bytes of kdf_hash2
		encrypt_iv[i] = kdf_hash2[i+16];

	OPENSSL_free(kdf_hash2);
	EC_KEY_free(privkey);
	EC_POINT_free(ec_point);
	EC_KEY_free(pubkey);
	EVP_PKEY_free(privkey_);
	BIO_free(bio);
	EVP_cleanup();
}

void AES128_Encrypt(const unsigned char* data, const int data_len, const unsigned char* encrypt_key, const unsigned char* encrypt_iv, unsigned char* encrypted_data_hex, 
	int* encrypted_data_hex_len, bool isDebug)
{
	EVP_CIPHER_CTX* ctx = EVP_CIPHER_CTX_new();
	EVP_CIPHER_CTX_init(ctx);
	if (1 != EVP_EncryptInit_ex(ctx, EVP_aes_128_cbc(), NULL, encrypt_key, encrypt_iv))
          handleErrors("AES128_Encrypt");
	OPENSSL_assert(EVP_CIPHER_CTX_key_length(ctx) == 16);
	OPENSSL_assert(EVP_CIPHER_CTX_iv_length(ctx) == 16);

	//unsigned char* outbuf = new unsigned char[2048];
	unsigned char* outbuf = (unsigned char*)OPENSSL_malloc(2048);
	int outlen=0;

	if (1 != EVP_EncryptUpdate(ctx, outbuf, &outlen, data, data_len))
          handleErrors("AES128_Encrypt");
	int tmplen = 0;

	if (1 != EVP_EncryptFinal_ex(ctx, outbuf + outlen, &tmplen))
          handleErrors("AES128_Encrypt");
	outlen += tmplen;

	/*std::ofstream fout("writezoneencrypted3.txt", std::ios::binary);
	for (int i = 0; i < outlen; i++)
		fout << outbuf[i];
	fout.close();*/

	*encrypted_data_hex_len = 2*outlen; //2 for Hex
	std::string temp;
	HEX::byte_to_Hex((void*)outbuf, outlen, temp);

	for (int i = 0; i < *encrypted_data_hex_len; i++)
		encrypted_data_hex[i] = temp[i];

	OPENSSL_free(outbuf);
	EVP_CIPHER_CTX_cleanup(ctx);
	EVP_cleanup();
}
/** NEED **/
void AES128_Decrypt(const unsigned char* encrypted_data, const int encrypted_data_len, const unsigned char* encrypt_key, const unsigned char* encrypt_iv, 
	unsigned char* decrypted_data, int* decrypted_data_len, bool isDebug)
{

	//printf("C++ debug input length: %d\n",encrypted_data_len);
	EVP_CIPHER_CTX* ctx = EVP_CIPHER_CTX_new();
	EVP_CIPHER_CTX_init(ctx);
	if (1 != EVP_DecryptInit_ex(ctx, EVP_aes_128_cbc(), NULL, encrypt_key, encrypt_iv))
	{
          handleErrors("AES128_Decrypt");
		  printf("Error in EVP_DecryptInit_ex\n");

	}

	OPENSSL_assert(EVP_CIPHER_CTX_key_length(ctx) == 16);
	OPENSSL_assert(EVP_CIPHER_CTX_iv_length(ctx) == 16);

	//unsigned char* outbuf = new unsigned char[2048];
	unsigned char* outbuf = (unsigned char*)OPENSSL_malloc(2048);
	int outlen=0;

	if (1 != EVP_DecryptUpdate(ctx, outbuf, &outlen, encrypted_data, encrypted_data_len))
	{
          handleErrors("AES128_Decrypt");
		  printf("Error in EVP_DecryptUpdate\n");
	}

	int tmplen = 0;

	if (1 != EVP_DecryptFinal_ex(ctx, outbuf + outlen, &tmplen))
	{
          handleErrors("AES128_Decrypt");
		  printf("Error in EVP_DecryptFinal_EX\n");
	}
	outlen += tmplen;

	*decrypted_data_len = outlen;

	for (int i = 0; i < *decrypted_data_len; i++)
		decrypted_data[i] = outbuf[i];

	OPENSSL_free(outbuf);
	EVP_CIPHER_CTX_cleanup(ctx);
	EVP_cleanup();
}