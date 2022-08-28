@REM MSYS Dependencies: pacman -S mingw-w64-x86_64-toolchain mingw-w64-x86_64-zlib mingw-w64-x86_64-openssl
@REM pacman -S --needed base-devel

SET PATH=C:\msys64\mingw64\bin;C:\msys64\usr\bin;%PATH%
SET CMAKE_C_COMPILER=gcc
SET CMAKE_CXX_COMPILER=g++

if exist win_build\ (
  rmdir /S /Q win_build
) 
mkdir win_build\gnscrypto\lib

@REM Building GNS crypto 
g++ -O2 -std=c++17 -Wall -c -Ignscryptocpp -march=x86-64 gnscryptocpp\gns_crypto.cpp -o win_build\gnscrypto\gnscrypto.o -static-libgcc -static-libstdc++
ar rc win_build\gnscrypto\lib\libgnscrypto.a win_build\gnscrypto\gnscrypto.o

@REM Building test app
xcopy server\GNSPrivate.pem win_build\gnscrypto\
@REM g++ -O2 -std=c++17 -Wall -march=x86-64 -Ignscryptocpp -o win_build\gnscrypto\gns_crypto_test gnscryptocpp\gns_crypto_test.cpp win_build\gnscrypto\lib\libgnscrypto.a -l:libz.a -l:libcrypto.a -l:libssl.a -lws2_32
g++ -O2 -std=c++17 -Wall -march=x86-64 -Ignscryptocpp -o win_build\gnscrypto\gns_crypto_test gnscryptocpp\gns_crypto_test.cpp win_build\gnscrypto\lib\libgnscrypto.a -l:libz.a -l:libcrypto.a -l:libssl.a -lws2_32 -static-libgcc -static-libstdc++