ARCHIVE_EXTENSION="tar.gz"

# Get OS and arch

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
WIN_OS="windowsnt mingw64_nt-10.0-22000"
if echo "$WIN_OS" | grep -q "$OS"; then
    OS="windows"
    ARCHIVE_EXTENSION="zip"
fi
SUPP_OS="darwin linux windows"
if ! echo "$SUPP_OS" | grep -q "$OS"; then
    echo "Unsupported OS: $OS"
    exit 1
fi
ARCH=$(uname -m | sed 's/x86_64/amd64/')

RELEASE_TYPE=$OS-$ARCH

# Get Release url

GH_RESPONSE=$(curl -fsSL https://api.github.com/repos/Oriun/doctor-compose/releases/latest)

substr=${GH_RESPONSE%%$RELEASE_TYPE*}
index=${#substr}
RELEASE_URL=${GH_RESPONSE:$index}
substr=${RELEASE_URL%%browser_download_url*}
index=${#substr}
RELEASE_URL=${RELEASE_URL:$index:200}
substr=${RELEASE_URL##*.$ARCHIVE_EXTENSION}
index=${#substr}
RELEASE_URL=${RELEASE_URL:24:(${#RELEASE_URL} - $index - 24)}

# Download and extract

echo "Downloading latest release for $RELEASE_TYPE from:"
echo $RELEASE_URL

curl -fSL $RELEASE_URL -o doctor-compose$ARCHIVE_EXTENSION
tar -xvf doctor-compose$ARCHIVE_EXTENSION
rm doctor-compose$ARCHIVE_EXTENSION
sudo mv doctor-compose /usr/local/bin/

echo "doctor-compose installed successfully"
