#!/bin/bash
arch="$(uname -m)"  # -i is only linux, -m is linux and apple
set -e

if [[ "$arch" = x86_64* ]]; then
clang++ -std=c++17 -Wall -c -Ignscryptocpp -Idarwin_build/openssl/include -arch x86_64 gnscryptocpp/gns_crypto.cpp -o darwin_build/gnscrypto/gnscrypto.o
ar rc darwin_build/gnscrypto/lib/gnscrypto.a darwin_build/gnscrypto/gnscrypto.o
libtool -static -o darwin_build/gnscrypto/lib/libgnscrypto.a darwin_build/gnscrypto/lib/gnscrypto.a darwin_build/openssl/lib/libssl.a darwin_build/openssl/lib/libcrypto.a

elif [[ "$arch" = arm* ]]; then
clang++ -std=c++17 -Wall -c -Ignscryptocpp -Idarwin_build/openssl/include -arch arm64 gnscryptocpp/gns_crypto.cpp -o darwin_build/gnscrypto/gnscrypto.o
ar rc darwin_build/gnscrypto/lib/gnscrypto.a darwin_build/gnscrypto/gnscrypto.o
libtool -static -o darwin_build/gnscrypto/lib/libgnscrypto.a darwin_build/gnscrypto/lib/gnscrypto.a darwin_build/openssl/lib/libssl.a darwin_build/openssl/lib/libcrypto.a
else
    echo 'Unsupported architecture'
    exit 1
fi




rm server/gnsdeviceserver
cd server
go build
cd ..