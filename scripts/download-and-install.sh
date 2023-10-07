#!/bin/bash

set -euf
LC_CTYPE=C

REPO_OWNER="fschossler"
REPO_NAME="tmdbcli"
VERSION=$(curl -s https://api.github.com/repos/$REPO_OWNER/$REPO_NAME/releases/latest | grep -o '"tag_name": ".*"' | cut -d'"' -f4)
INSTALL_DIR="/usr/local/bin"
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/latest/download/tmdbcli-$VERSION-$OS-amd64.tar.gz"

# Handle errors and exit
handle_error() {
  local error_message="$1"
  echo "❌ Error: $error_message"
  exit 1
}

# Check if the user has sudo privileges
check_sudo() {
  if [ "$(id -u)" -ne 0 ]; then
    echo "🔒 This installation script may require root privileges to move the binary."
    sudo echo "Prompting for sudo password..."
  fi
}

# Download a file with error handling
download_file() {
  local url="$1"
  local output_file="$2"

  if ! curl -s -L -o "$output_file" "$url"; then
    handle_error "Failed to download the file from $url."
  fi
}

# Extract a tar.gz file with error handling
extract_tar_gz() {
  local input_file="$1"
  local output_dir="$2"

  if ! tar -C "$output_dir" -xzf "$input_file"; then
    handle_error "Failed to extract the tar.gz file."
  fi
}

check_sudo

echo "📥 Downloading tmdbcli..."
download_file "$URL" "/tmp/tmdbcli.tar.gz"

echo "📦 Extracting tmdbcli..."
extract_tar_gz "/tmp/tmdbcli.tar.gz" "/tmp"

echo "🚚 Moving tmdbcli to $INSTALL_DIR (requires sudo)..."
if ! sudo mv "/tmp/tmdbcli" "$INSTALL_DIR"; then
  handle_error "Failed to move the binary to $INSTALL_DIR."
fi

echo "🧹 Cleaning up 'tar.gz' ..."
rm -f "/tmp/tmdbcli.tar.gz"

# Check if the binary is now available in the user's PATH
if command -v tmdbcli &>/dev/null; then
  echo "✅ Installation complete. You can now use 'tmdbcli' from the command line."
else
  handle_error "Installation failed. Please make sure to add $INSTALL_DIR to your PATH."
fi

echo "👀 Please don't forget to follow the Requirements in the README for everything works perfectly."