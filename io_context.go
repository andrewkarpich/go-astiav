package astiav

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import "unsafe"

// https://github.com/FFmpeg/FFmpeg/blob/n5.0/libavformat/avio.h#L161
type IOContext struct {
	c *C.struct_AVIOContext
}

func NewIOContext() *IOContext {
	return &IOContext{}
}

func NewIOContextFromC(c *C.struct_AVIOContext) *IOContext {
	if c == nil {
		return nil
	}
	return &IOContext{c: c}
}

func AllocIOContext(
	buf unsafe.Pointer,
	bufferSize int,
	flags IOContextFlags,
	opaque unsafe.Pointer,
	readCallBack *[0]byte,
	writeCallBack *[0]byte,
	seekCallBack *[0]byte,
) *IOContext {
	return NewIOContextFromC(
		C.avio_alloc_context(
			(*C.uchar)(buf),
			C.int(bufferSize),
			C.int(flags),
			opaque,
			readCallBack,  // int (*read_packet)(void *opaque, uint8_t *buf, int buf_size)
			writeCallBack, // int (*write_packet)(void *opaque, uint8_t *buf, int buf_size)
			seekCallBack,  // int64_t (*seek)(void *opaque, int64_t offset, int whence)
		),
	)
}

func (ic *IOContext) Closep() error {
	return newError(C.avio_closep(&ic.c))
}

func (ic *IOContext) Open(filename string, flags IOContextFlags) error {
	cfi := C.CString(filename)
	defer C.free(unsafe.Pointer(cfi))
	return newError(C.avio_open(&ic.c, cfi, C.int(flags)))
}

func (ic *IOContext) Write(b []byte) {
	if b == nil {
		return
	}
	C.avio_write(ic.c, (*C.uchar)(unsafe.Pointer(&b[0])), C.int(len(b)))
}

// AvClassName TODO: realize full AvClass
func (ic *IOContext) AvClassName() string {
	return C.GoString(ic.c.av_class.class_name)
}
