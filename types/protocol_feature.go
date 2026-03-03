package types

type ProtocolFeature uint64

const (
	ProtocolFeatureMQ                   ProtocolFeature = 1 << 0
	ProtocolFeatureLogShmfd             ProtocolFeature = 1 << 1
	ProtocolFeatureRarp                 ProtocolFeature = 1 << 2
	ProtocolFeatureReplyAck             ProtocolFeature = 1 << 3
	ProtocolFeatureMTU                  ProtocolFeature = 1 << 4
	ProtocolFeatureBackendReq           ProtocolFeature = 1 << 5
	ProtocolFeatureCrossEndian          ProtocolFeature = 1 << 6
	ProtocolFeatureCryptoSession        ProtocolFeature = 1 << 7
	ProtocolFeaturePagefault            ProtocolFeature = 1 << 8
	ProtocolFeatureConfig               ProtocolFeature = 1 << 9
	ProtocolFeatureBackendSendFd        ProtocolFeature = 1 << 10
	ProtocolFeatureHostNotifier         ProtocolFeature = 1 << 11
	ProtocolFeatureInflightShmfd        ProtocolFeature = 1 << 12
	ProtocolFeatureResetDevice          ProtocolFeature = 1 << 13
	ProtocolFeatureInbandNotifications  ProtocolFeature = 1 << 14
	ProtocolFeatureConfigureMemSlots    ProtocolFeature = 1 << 15
	ProtocolFeatureStatus               ProtocolFeature = 1 << 16
	ProtocolFeatureXenMmap              ProtocolFeature = 1 << 17
	ProtocolFeatureSharedObject         ProtocolFeature = 1 << 18
	ProtocolFeatureDeviceState          ProtocolFeature = 1 << 19
	ProtocolFeatureGetVringBaseInflight ProtocolFeature = 1 << 20
)
