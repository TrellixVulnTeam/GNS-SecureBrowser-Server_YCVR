# Build
## Dependencies
- Golang installed: https://go.dev/dl/


## Build script

### macOS
- run darwin_build.sh
   - This will build openssl library
   - mock_service api

### Windows extra steps for credential manager
- Visual Studio
- VCPKG (C++ package manager to manage GRP C++ libraries)
   - run integration script to install VCPKG to use with visual studio IDE
   - vcpkg install grpc[codegen]:x64-windows-static
   - add <VcpkgTriplet Condition="'$(Platform)'=='x64'">x64-windows-static</VcpkgTriplet> to setup static link

### Linux
- Must install libssl-dev (sudo apt install libssl-dev)
- run linux_build.sh

## Run mock server
- The binary from build script should be produced in server/mock_server/mock_server
   - The mock service should listen on 127.0.0.1:50051
# GNS-SecureBrowser-Server
