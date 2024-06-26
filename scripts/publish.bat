@echo off
setlocal enabledelayedexpansion

REM Get the last version in repository
for /f "delims=" %%a in ('git describe --abbrev=0 --tags') do set "CURRENT_VERSION=%%a"

REM Get the declared version in .version file
for /f "delims=" %%a in (.version) do set "VERSION=%%a"

REM Check if versions are equal
if "!CURRENT_VERSION!" == "!VERSION!" (
    echo Versions are equal. Nothing to publish
    exit 0
)

REM Tidy the repository
go mod tidy

REM Create the new tag version
git tag "!VERSION!"

REM Push the new tag version
git push origin "!VERSION!"

REM Publish the new version
for /f "delims=" %%a in ('go list -m github.com/Diegiwg/tt@"!VERSION!"') do set "PUBLISHED=%%a"

REM Test if the version was published
for /f "tokens=1* delims=:" %%a in ('curl -o NUL -s -w "%{http_code}:%%{url_effective}" https://pkg.go.dev/github.com/Diegiwg/tt@"!VERSION!"') do (
    set "OK=%%a"
)

if "!OK!" == "200" (
    echo Version !VERSION! was published
) else (
    echo Version !VERSION! was not published
    exit 1
)

endlocal