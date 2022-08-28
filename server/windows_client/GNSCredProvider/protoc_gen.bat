
@REM md .build
@REM cmake .. -G "Visual Studio 17 2022"
@REM cmake --build . --config Release
@REM C:\vcpkg\vcpkg.exe install grpc:x64-windows
@REM c:\vcpkg\vcpkg.exe integrate install
C:\vcpkg\packages\protobuf_x64-windows\tools\protobuf\protoc -I ../../proto gnsservice.proto --cpp_out=grpc/include --cppgrpc_out=grpc/include