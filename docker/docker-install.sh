#!/bin/bash

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    bash docker-install_linux.sh
elif [[ "$OSTYPE" == "msys" ]] || [[ "$OSTYPE" == "win32" ]]; then
    powershell.exe -ExecutionPolicy Bypass -File docker-install_windows.ps1
else
    echo "Error unsupported operating system: $OSTYPE"
    exit 1
fi
