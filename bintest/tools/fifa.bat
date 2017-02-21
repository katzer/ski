@echo off

set "param1=%~1"
set "param2=%~2"

IF "%param1%"=="-t" (
    IF "%param2%"=="app"           ECHO server
    IF "%param2%"=="web"           ECHO web
    IF "%param2%"=="db"            ECHO db
    IF "%param2%"=="unauthorized"  ECHO server
    IF "%param2%"=="offline"       ECHO server
)

IF "%param1%"=="app"           ECHO localhost
IF "%param1%"=="web"           ECHO localhost
IF "%param1%"=="db"            ECHO localhost
IF "%param1%"=="unauthorized"  ECHO user@localhost
IF "%param1%"=="offline"       ECHO user@testhost
