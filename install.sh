#!/bin/bash

echo ""
echo "Confirming curl is installed..."
# confirm curl is installed
if ! hash curl 2> /dev/null; then
  printf "Error: you do not have 'curl' installed.\n"
  exit 1
fi

echo ""
echo "Fetching latest version..."
# get latest version
latest_version=$(curl -s https://raw.githubusercontent.com/jonh-a/speedreader/main/latest_version.txt)
echo " - Fetched version $latest_version."

# get OS
os_type=$(uname | tr '[:upper:]' '[:lower:]')

# get arch
architecture=$(uname -m)
if [[ "$architecture" == "x86_64" ]]; then
    architecture="amd64"
elif [[ "$architecture" == "aarch64" ]]; then
    architecture="arm64"
fi

echo ""
echo "Downloading latest release..."
# determine the install directory
if [[ "$os_type" == "darwin" ]]; then
    install_dir="$HOME/bin"
    if [ ! -d "$install_dir" ]; then
        echo "Creating directory $install_dir"
        mkdir "$install_dir"
    fi
else
    install_dir="$HOME/.local/bin"
fi

# download and install speedread executable
url="https://github.com/jonh-a/speedreader/releases/download/${latest_version}/speedread_${os_type}_${architecture}"
curl -sL "$url" -o "$install_dir/speedreader"
chmod +x "$install_dir/speedreader"
echo " - Speedreader installed in $install_dir."

echo ""
echo "Adding speedreader to system path..."
# modify system path
if [[ "$SHELL" == *"bash" ]]; then
    echo "export PATH=\$PATH:$install_dir" >> ~/.bashrc
    source ~/.bashrc
    echo " - ~/.bashrc successfully modified."
elif [[ "$SHELL" == *"zsh" ]]; then
    echo "export PATH=\$PATH:$install_dir" >> ~/.zshrc
    source ~/.zshrc
    echo " - ~/.zshrc successfully modified."
elif [[ "$SHELL" == *"fish" ]]; then
    echo "set -x PATH $install_dir \$PATH" >> ~/.config/fish/config.fish
    source ~/.config/fish/config.fish
    echo " - ~/.config/fish/config.fish successfully modified."
fi

echo ""
echo "Speedreader is ready to use!"

echo ""
echo "Try running the following:"
echo "    echo \"Hi there, I don't have too much to say.\" | speedreader -w 250"