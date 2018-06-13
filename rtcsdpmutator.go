package webrtc

import "github.com/pions/webrtc/internal/sdp"

type RTCSDPMutator struct {
	sdp *sdp.SessionDescription
}

type RTCMediaDescriptionMutator struct {
	md *sdp.MediaDescription
}

func (m *RTCMediaDescriptionMutator) IsRTCPMux() bool {
	found, _ := m.md.HasAttribute("rtcp-mux")
	return found;
}