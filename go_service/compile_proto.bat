@echo off
echo 🔧 Compiling protobuf files...

cd /d C:\Users\IGMA\Desktop\GitRepo\music_fed364847\go_service

:: Создаем папку gen если её нет
if not exist "gen" mkdir gen
if not exist "gen\serverGRPC" mkdir gen\serverGRPC

:: Компилируем proto файлы
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos\*.proto

:: Перемещаем сгенерированные файлы в gen\serverGRPC
if exist "protos\*.pb.go" (
    move protos\*.pb.go gen\serverGRPC\
    echo ✅ Files moved to gen\serverGRPC\
)

echo ✅ Proto compilation completed!
pause