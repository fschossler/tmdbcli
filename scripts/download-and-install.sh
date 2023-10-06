#!/bin/bash

# Improve error handling
set -euf
LC_CTYPE=C

REPO_OWNER="fschossler"
REPO_NAME="tmdbcli"
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
DEFAULT_VERSION="latest"
VERSION="${1:-$DEFAULT_VERSION}"
URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/download/$VERSION/tmdbcli_v$VERSION-$OS-amd64.tar.gz"
INSTALL_DIR="/usr/local/bin"

# Function to handle errors and exit
handle_error() {
  local error_message="$1"
  echo "❌ Error: $error_message"
  exit 1
}

# Function to check if the user has sudo privileges
check_sudo() {
  if [ "$(id -u)" -ne 0 ]; then
    echo "🔒 This installation script may require root privileges to move the binary."
    sudo echo "Prompting for sudo password..."
  fi
}

# Function to download a file with error handling
download_file() {
  local url="$1"
  local output_file="$2"

  if ! curl -L -o "$output_file" "$url"; then
    handle_error "❌ Failed to download the file from $url."
  fi
}

# Function to extract a tar.gz file with error handling
extract_tar_gz() {
  local input_file="$1"
  local output_dir="$2"

  if ! tar -C "$output_dir" -xzf "$input_file"; then
    handle_error "❌ Failed to extract the tar.gz file."
  fi
}

# Check if root privileges are required
check_sudo

# Download the release asset
echo "📥 Downloading tmdbcli..."
download_file "$URL" "/tmp/tmdbcli.tar.gz"

# Extract the binary
echo "📦 Extracting tmdbcli..."
extract_tar_gz "/tmp/tmdbcli.tar.gz" "/tmp"

# Move the binary to the installation directory
echo "🚚 Moving tmdbcli to $INSTALL_DIR (requires sudo)..."
if ! sudo mv "/tmp/tmdbcli" "$INSTALL_DIR"; then
  handle_error "Failed to move the binary to $INSTALL_DIR."
fi

# Cleanup the downloaded tar.gz file
echo "🧹 Cleaning up..."
rm -f "/tmp/tmdbcli.tar.gz"

# Check if the binary is now available in the user's PATH
if command -v tmdbcli &>/dev/null; then
  echo "✅ Installation complete. You can now use 'tmdbcli' from the command line."
else
  handle_error "Installation failed. Please make sure to add $INSTALL_DIR to your PATH."
fi

echo "==========================================================="
echo "👀 Please don't forget to follow the Requirements in the README for everything works perfectly."