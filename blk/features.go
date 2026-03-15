package blk

import "vhost-go/types"

func SupportedFeatures() uint64 {
	var f uint64
	f |= 1 << types.VIRTIO_F_VERSION_1
	f |= 1 << types.VHOST_USER_F_PROTOCOL_FEATURES
	return f
}

func SupportedProtocolFeatures() uint64 {
	var f uint64
	f |= 1 << types.VHOST_USER_PROTOCOL_F_MQ
	f |= 1 << types.VHOST_USER_PROTOCOL_F_CONFIG
	return f
}
