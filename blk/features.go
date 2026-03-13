package blk

import "vhost-go/types"

func SupportedFeatures() uint64 {
	var f uint64
	f |= 1 << types.VIRTIO_F_VERSION_1
	return f
}

func SupportedProtocolFeatures() uint64 {
	var f uint64
	return f
}
