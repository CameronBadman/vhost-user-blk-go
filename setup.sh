#!/bin/bash

set -e

ALPINE_VERSION="3.19.1"
ALPINE_ISO="alpine-standard-${ALPINE_VERSION}-x86_64.iso"
ALPINE_URL="https://dl-cdn.alpinelinux.org/alpine/v3.19/releases/x86_64/${ALPINE_ISO}"
DISK_IMAGE="test-vm.qcow2"
DISK_SIZE="8G"

if [ ! -f "$ALPINE_ISO" ]; then
    echo "Downloading Alpine ${ALPINE_VERSION}..."
    wget -q --show-progress "$ALPINE_URL"
fi

if [ ! -f "$DISK_IMAGE" ]; then
    echo "Creating disk image..."
    qemu-img create -f qcow2 "$DISK_IMAGE" "$DISK_SIZE"
fi

echo "Booting Alpine installer — run 'setup-alpine' then poweroff when done"

sudo qemu-system-x86_64 \
  -enable-kvm \
  -m 512M \
  -cpu host \
  -drive file=test-vm.qcow2,if=virtio \
  -cdrom alpine-standard-3.19.1-x86_64.iso \
  -boot d \
  -nographic \
  -serial mon:stdio \
  -netdev user,id=net0 \
  -device virtio-net-pci,netdev=net0
