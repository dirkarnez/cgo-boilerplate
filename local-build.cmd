REM run as Administrator
@echo off
cd /d %~dp0
set DOWNLOAD_DIR=%USERPROFILE%\Downloads
set DOWNLOAD_DIR_LINUX=%DOWNLOAD_DIR:\=/%
SET PATH=%DOWNLOAD_DIR%\PortableGit\bin;%DOWNLOAD_DIR%\go\bin;%DOWNLOAD_DIR%\mingw64-win32;%DOWNLOAD_DIR%\mingw64-win32\bin;%DOWNLOAD_DIR%\cmake-3.22.0-rc1-windows-x86_64\bin;

cd pointcpp
cmake.exe -G"MinGW Makefiles" ^
-B../build
cd ../build

:rebuild_and_startapp
mingw32-make.exe

cd ..
go build -ldflags="-extldflags '-static -lstdc++'"

pause
REM GOTO rebuild_and_startapp