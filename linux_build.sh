mkdir -p linux_build/gnscrypto/lib
g++ -O2 -std=c++17 -Wall -c -Ignscryptocpp gnscryptocpp/gns_crypto.cpp -o linux_build/gnscrypto/gnscrypto.o  
ar rc linux_build/gnscrypto/lib/libgnscrypto.a linux_build/gnscrypto/gnscrypto.o