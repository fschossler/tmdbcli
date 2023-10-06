#!/bin/bash

# Enable robust error handling and consistent character encoding
set -euf
LC_CTYPE=C

# Specify the GitHub repository owner and name
REPO_OWNER="fschossler"
REPO_NAME="tmdbcli"

# Determine the user's operating system
OS=$(uname -s | tr '[:upper:]' '[:lower:]')

# Define the default version or use the provided argument
DEFAULT_VERSION="latest"
VERSION="${1:-$DEFAULT_VERSION}"

# Construct the URL for the release asset based on the OS and version
if [ "$VERSION" == "latest" ]; then
  URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/latest/download/tmdbcli-$VERSION-$OS-amd64.tar.gz"
else
  URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/download/$VERSION/tmdbcli-$VERSION-$OS-amd64.tar.gz"
fi

# Define the installation directory (where the binary will be placed)
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
    handle_error "Failed to download the file from $url."
  fi
}

# Function to extract a tar.gz file with error handling
extract_tar_gz() {
  local input_file="$1"
  local output_dir="$2"

  # Check if the file is a valid gzip archive
  if ! gzip -t "$input_file" &>/dev/null; then
    handle_error "The downloaded file is not a valid gzip archive."
  fi

  # Extract the file
  if ! tar -C "$output_dir" -xzf "$input_file"; then
    handle
