#!/bin/bash

set -e

if [[ "$EUID" -ne 0 ]]; then
    echo "This script must be run with sudo or as root."
    exit 1
fi

echo "STAGE: LOAD Variables"
source .env
source ./version_setup_variables.sh
source /etc/os-release
WORK_DIR="/opt"
UBUNTU_CODENAME_ID="ubuntu${VERSION_ID//./}"

echo "System Version: $UBUNTU_CODENAME_ID"

echo "STAGE: INSTALL Dependencies"
apt-get update && \
apt-get install -y \
    curl \
    wget \
    unzip \
    git \
    cmake \
    g++ \
    gcc \
    build-essential \
    libboost-all-dev \
    libprotobuf-dev \
    protobuf-compiler \
    libopenblas-dev \
    libhdf5-dev \
    libyaml-cpp-dev \
    libtcmalloc-minimal4 \
    libgoogle-glog-dev \
    libgflags-dev \
    liblzma-dev \
    libzip-dev \
    zlib1g-dev && \
rm -rf /var/lib/apt/lists/*

cd $WORK_DIR

echo "STAGE: INSTALL CUDA"
CUDA_INSTALLER=cuda_${CUDA_RELEASE_VERSION}_${CUDA_DRIVER_VERSION}_linux.run

wget -O ${CUDA_INSTALLER} https://developer.download.nvidia.com/compute/cuda/${CUDA_RELEASE_VERSION}/local_installers/${CUDA_INSTALLER}
sh ./${CUDA_INSTALLER} --silent --toolkit --override --installpath=${CUDA_INSTALL_PATH}
rm ${CUDA_INSTALLER}

CUDA_SYSTEM_PIN=cuda-${UBUNTU_CODENAME_ID}.pin
CUDA_REPO_INSTALLER=cuda-repo-${UBUNTU_CODENAME_ID}-${CUDA_VERSION}-local_${CUDA_RELEASE_VERSION}-${CUDA_DRIVER_VERSION_PATH}_amd64.deb

wget https://developer.download.nvidia.com/compute/cuda/repos/${UBUNTU_CODENAME_ID}/x86_64/${CUDA_SYSTEM_PIN}
mv ${CUDA_SYSTEM_PIN} /etc/apt/preferences.d/cuda-repository-pin-600
wget https://developer.download.nvidia.com/compute/cuda/$CUDA_RELEASE_VERSION/local_installers/${CUDA_REPO_INSTALLER}
dpkg -i $CUDA_REPO_INSTALLER
cp /var/cuda-repo-${UBUNTU_CODENAME_ID}-${CUDA_VERSION}-local/cuda-*-keyring.gpg /usr/share/keyrings/
apt-get update
apt-get -y install cuda-toolkit-${CUDA_VERSION}
rm -rf ${CUDA_REPO_INSTALLER}

echo "STAGE: INSTALL MARIAN"
MARIAN_PROJECT_DIR=$WORK_DIR/marian

git clone https://github.com/marian-nmt/marian.git $MARIAN_PROJECT_DIR
cd $MARIAN_PROJECT_DIR
git checkout ${MARIAN_VERSION}
mkdir -p build && cd build
cmake -DCMAKE_BUILD_TYPE=Release \
    -DCOMPILE_CUDA=ON \
    -DUSE_NCC=OFF \
    -DCUDA_TOOLKIT_ROOT_DIR="/usr/local/cuda" \
    -DCMAKE_LIBRARY_PATH="/usr/local/cuda/lib64" \
    -DCMAKE_INCLUDE_PATH="/usr/local/cuda/include" \
    ..
make -j8
mkdir -p ${MARIAN_INSTALL_PATH}
cp marian-decoder ${MARIAN_INSTALL_PATH}
chown -R $SUDO_USER:$SUDO_USER ${MARIAN_INSTALL_PATH}
cd $WORK_DIR
rm -rf $MARIAN_PROJECT_DIR

echo "STAGE: INSTALL Vocabulary"
OPUS_BT_ARCHIVE=opus+bt-${OPUS_BT_VERSION}.zip
OPUS_ARCHIVE=opus-${OPUS_VERSION}.zip

wget https://object.pouta.csc.fi/Tatoeba-MT-models/${OPUS_LANG}/${OPUS_BT_ARCHIVE}
mkdir -p ${VOCAB_BT_PATH}
unzip ${OPUS_BT_ARCHIVE} -d ${VOCAB_BT_PATH}
chown -R $SUDO_USER:$SUDO_USER ${VOCAB_BT_PATH}
rm ${OPUS_BT_ARCHIVE}

wget https://object.pouta.csc.fi/Tatoeba-MT-models/${OPUS_LANG}/${OPUS_ARCHIVE}
mkdir -p ${VOCAB_PATH}
unzip ${OPUS_ARCHIVE} -d ${VOCAB_PATH}
chown -R $SUDO_USER:$SUDO_USER ${VOCAB_PATH}
rm ${OPUS_ARCHIVE}
