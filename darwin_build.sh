#!/bin/bash
# Code adopted from:
# https://stackoverflow.com/questions/69002453/how-to-build-openssl-for-m1-and-for-intel
arch="$(uname -m)"  # -i is only linux, -m is linux and apple
set -e


if [ -f "darwin_build" ] 
then
    rm -rf darwin_build
fi
mkdir -p darwin_build/openssl
mkdir -p darwin_build/openssl/lib
mkdir -p darwin_build/gnscrypto
mkdir -p darwin_build/gnscrypto/lib
mkdir -p darwin_build/gnscrypto/include



echo '******************************************************'
echo 'OpenSSL:build start'



if [[ "$arch" = x86_64* ]]; then
  export MACOSX_DEPLOYMENT_TARGET=12.0
  cd openssl
  ./Configure darwin64-x86_64-cc shared
  make

elif [[ "$arch" = arm* ]]; then
  export MACOSX_DEPLOYMENT_TARGET=10.15
  cd openssl
  ./Configure enable-rc5 zlib darwin64-arm64-cc no-asm
  make
else
    echo 'Unsupported architecture'
    exit 1
fi


# Copy include and binaries to build folder
cp -r include ../darwin_build/openssl/
cp lib*.* ../darwin_build/openssl/lib/

cd .. 

echo '******************************************************'
echo 'OpenSSL:build done results are in darwin_build/openssl'
echo



echo '******************************************************'
echo 'GNSCRYPTO: starting build'
echo $(pwd)

if [[ "$arch" = x86_64* ]]; then
clang++ -O2 -std=c++17 -Wall -c -Ignscryptocpp -Idarwin_build/openssl/include -arch x86_64 gnscryptocpp/gns_crypto.cpp -o darwin_build/gnscrypto/gnscrypto.o   
elif [[ "$arch" = arm* ]]; then
  #cp -r openssl darwin_build/openssl-arm64 
clang++ -O2 -std=c++17 -Wall -c -Ignscryptocpp -Idarwin_build/openssl/include -arch arm64 gnscryptocpp/gns_crypto.cpp -o darwin_build/gnscrypto/gnscrypto.o    
else
echo 'Unsupported archicture'
exit
fi

# Optimizing 
ar rc darwin_build/gnscrypto/lib/gnscrypto.a darwin_build/gnscrypto/gnscrypto.o

# Combine OpenSSL static lib with GNSCrypto Lib
libtool -static -o darwin_build/gnscrypto/lib/libgnscrypto.a darwin_build/gnscrypto/lib/gnscrypto.a darwin_build/openssl/lib/libssl.a darwin_build/openssl/lib/libcrypto.a

echo
echo '******************************************************'
echo 'GNSCRYPTO: building basic test app'
# Build test app

cp server/GNSPrivate.pem darwin_build/gnscrypto/

if [[ "$arch" = x86_64* ]]; then
clang++ -O2 -std=c++17 -Wall -arch x86_64 -Ignscryptocpp -Idarwin_build/openssl/include -L$(xcrun --sdk macosx --show-sdk-path)/usr/lib -lz\
 -o darwin_build/gnscrypto/gns_crypto_test gnscryptocpp/gns_crypto_test.cpp darwin_build/gnscrypto/lib/libgnscrypto.a
elif [[ "$arch" = arm* ]]; then
  #cp -r openssl darwin_build/openssl-arm64 
clang++ -O2 -std=c++17 -Wall -arch arm64 -Ignscryptocpp -Idarwin_build/openssl/include -L$(xcrun --sdk macosx --show-sdk-path)/usr/lib -lz\
 -o darwin_build/gnscrypto/gns_crypto_test gnscryptocpp/gns_crypto_test.cpp darwin_build/gnscrypto/lib/libgnscrypto.a
else
echo 'Unsupported archicture'
exit
fi


echo
echo '******************************************************'
echo 'GNSCRYPTO: running test app'
# run test app

cd darwin_build/gnscrypto
 ./gns_crypto_test

cd ../..

echo 

echo
echo '******************************************************'
echo 'Building mock GNSDataService gRPC server'
cd server
go mod download
cd mock_server
go build
cd ../..

echo 'The binary mock service is over at server/mock_server/mock_server'