#!/bin/bash

echo "Confirming curl is installed..."
# confirm curl is installed
if ! hash curl 2> /dev/null; then
  printf "Error: you do not have 'curl' installed.\n"
  exit 1
fi

echo "Fetching latest version..."
# get latest version
latest_version=$(curl -s https://raw.githubusercontent.com/jonh-a/speedreader/main/latest_version.txt)

# get OS
os_type=$(uname | tr '[:upper:]' '[:lower:]')

# get arch
architecture=$(uname -m)
if [[ "$architecture" == "x86_64" ]]; then
    architecture="amd64"
elif [[ "$architecture" == "aarch64" ]]; then
    architecture="arm64"
fi

echo "Downloading latest release..."
# download the latest release
mkdir ~/.bin
curl -sL "https://github.com/jonh-a/speedreader/releases/download/${latest_version}/speedread_${os_type}_${architecture}" -o ~/.bin/speedreader
chmod +x ~/.bin/speedreader

echo "Adding speedreader to system path..."
# modify system path
if [[ "$SHELL" == *"bash" ]]; then
    echo "export PATH=\$PATH:~/.bin/speedreader" >> ~/.bashrc
elif [[ "$SHELL" == *"zsh" ]]; then
    echo "export PATH=\$PATH:~/.bin/speedreader" >> ~/.zshrc
elif [[ "$SHELL" == *"fish" ]]; then
    echo "set -x PATH ~/.bin/speedreader \$PATH" >> ~/.config/fish/config.fish
fi

# refresh shell
source "$HOME/.${SHELL##*/}rc"

echo "Speedreader is now installed!"