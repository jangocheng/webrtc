package webrtc

type RTCDTMFSender interface {}

type RTCRTPReceiver struct {
	dtmf RTCDTMFSender
	transport RTCDTLSTransport
	rtcpTransport RTCDTLSTransport
}