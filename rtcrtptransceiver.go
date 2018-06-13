package webrtc

type RTCRtpTransceiverDirection string

const (
	SendRecv RTCRtpTransceiverDirection = "sendrecv"
	SendOnly RTCRtpTransceiverDirection = "sendonly"
	RecvOnly RTCRtpTransceiverDirection = "recvonly"
	Inactive RTCRtpTransceiverDirection = "inactive"
)

type RTCRTPTransceiver struct {
	currentDirection RTCRtpTransceiverDirection
	mid string
	receiver RTCRTPReceiver
	//sender RTCRTPSender
	stopped bool
}

func (t *RTCRTPTransceiver) CurrentDirection() RTCRtpTransceiverDirection {
	return t.currentDirection
}

func (t *RTCRTPTransceiver) Mid() string {
	return t.mid
}

func (t *RTCRTPTransceiver) Receiver() *RTCRTPReceiver {
	return &t.receiver
}

//func (t *RTCRTPTransceiver) Sender() *RTCRTPSender {
//	return &t.sender
//}

func (t *RTCRTPTransceiver) Stopped() bool {
	return t.stopped
}