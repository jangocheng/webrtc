package gst

/*
#cgo pkg-config: gstreamer-1.0 gstreamer-app-1.0

#include "gst.h"

*/
import "C"
import (
	"unsafe"
)

// Pipeline is a wrapper for a GStreamer Pipeline
type Pipeline struct {
	Pipeline *C.GstElement
}

// CreatePipeline creates a GStreamer Pipeline
func CreatePipeline() *Pipeline {
	return &Pipeline{Pipeline: C.gst_create_pipeline()}
}

// Start starts the GStreamer Pipeline
func (p *Pipeline) Start() {
	C.gst_start_pipeline(p.Pipeline)
}

// Stop stops the GStreamer Pipeline
func (p *Pipeline) Stop() {
	C.gst_stop_pipeline(p.Pipeline)
}

// Push pushes a buffer on the appsrc of the GStreamer Pipeline
func (p *Pipeline) Push(buffer []byte) {
	b := C.CBytes(buffer)
	defer C.free(unsafe.Pointer(b))
	C.gst_push_buffer(p.Pipeline, b, C.int(len(buffer)))
}
