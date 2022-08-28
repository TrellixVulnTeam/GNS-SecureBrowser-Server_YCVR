Instructions to regenerate these keys in the future:
https://github.com/square/certstrap

 ```
 ./certstrap init --common-name global-net-solutions.com

./certstrap request-cert -ip 127.0.0.1 --common-name "GNSGRPCServer"
./certstrap sign GNSGRPCServer --CA global-net-solutions.com

./certstrap request-cert  --common-name "Client"
./certstrap sign Client --CA global-net-solutions.com

 ```