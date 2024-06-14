#!/bin/bash

command_exists () {
    command -v "$1" &> /dev/null ;
}

if ! command_exists go ; then
    echo "Go could not be found. Installing Go..."
    sudo apt update
    sudo apt install -y golang
fi

if ! command_exists subfinder ; then
    echo "subfinder could not be found. Installing subfinder..."
    sudo apt update
    sudo apt install -y subfinder
fi

if ! command_exists theHarvester ; then
    echo "theHarvester could not be found. Installing theHarvester..."
    sudo apt update
    sudo apt install -y theharvester
fi

if ! command_exists dnsrecon ; then
    echo "dnsrecon could not be found. Installing dnsrecon..."
    sudo apt update
    sudo apt install -y dnsrecon
fi

echo "Building the Forager tool..."
if go build -o forage forage.go; then
    # Check if /usr/local/bin/forage exists and is a directory
    if [ -d "/usr/local/bin/forage" ]; then
        echo "/usr/local/bin/forage exists and is a directory. Removing it..."
        sudo rm -rf /usr/local/bin/forage
    elif [ -f "/usr/local/bin/forage" ]; then
        echo "/usr/local/bin/forage exists and is a file. Removing it..."
        sudo rm /usr/local/bin/forage
    fi

    # Move the binary to /usr/local/bin
    sudo mv forage /usr/local/bin/
    sudo chmod +x /usr/local/bin/forage
    echo "Installation complete. You can now run 'forage example.com'"
else
    echo "Build failed. forage binary not found."
fi
