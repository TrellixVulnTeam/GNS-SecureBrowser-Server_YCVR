#include "gns_crypto.h"
#include <string>
#include <fstream>
#include <sstream>
#include <iostream>

int main(){
    unsigned char x[65];
    unsigned char y[65];

    std::cout << "testing loading pubkey from GNSPrivate.pem\n";
    generate_pubkey_from_privkey("GNSPrivate.pem",x,y,true);
    x[64]=0;
    y[64]=0;


    std::cout << x;
    std::cout << std::endl;
    std::cout << y << std::endl;


}