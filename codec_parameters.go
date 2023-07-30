package astiav

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"reflect"
	"unsafe"
)

// https://github.com/FFmpeg/FFmpeg/blob/n5.0/libavcodec/codec_par.h#L52
type CodecParameters struct {
	c *C.struct_AVCodecParameters
}

func AllocCodecParameters() *CodecParameters {
	return newCodecParametersFromC(C.avcodec_parameters_alloc())
}

func newCodecParametersFromC(c *C.struct_AVCodecParameters) *CodecParameters {
	if c == nil {
		return nil
	}
	return &CodecParameters{c: c}
}

func (cp *CodecParameters) Free() {
	C.avcodec_parameters_free(&cp.c)
}

func (cp *CodecParameters) BitRate() int64 {
	return int64(cp.c.bit_rate)
}

func (cp *CodecParameters) SetBitRate(i int64) {
	cp.c.bit_rate = C.long(i)
}

func (cp *CodecParameters) ChannelLayout() ChannelLayout {
	return ChannelLayout(cp.c.channel_layout)
}

func (cp *CodecParameters) SetChannelLayout(l ChannelLayout) {
	cp.c.channel_layout = C.uint64_t(l)
}

func (cp *CodecParameters) Channels() int {
	return int(cp.c.channels)
}

func (cp *CodecParameters) SetChannels(c int) {
	cp.c.channels = C.int(c)
}

func (cp *CodecParameters) CodecID() CodecID {
	return CodecID(cp.c.codec_id)
}

func (cp *CodecParameters) SetCodecID(i CodecID) {
	cp.c.codec_id = uint32(i)
}

func (cp *CodecParameters) CodecTag() CodecTag {
	return CodecTag(cp.c.codec_tag)
}

func (cp *CodecParameters) SetCodecTag(t CodecTag) {
	cp.c.codec_tag = C.uint(t)
}

func (cp *CodecParameters) CodecType() MediaType {
	return MediaType(cp.c.codec_type)
}

func (cp *CodecParameters) SetCodecType(t MediaType) {
	cp.c.codec_type = int32(t)
}

func (cp *CodecParameters) ChromaLocation() ChromaLocation {
	return ChromaLocation(cp.c.chroma_location)
}

func (cp *CodecParameters) SetChromaLocation(l ChromaLocation) {
	cp.c.chroma_location = uint32(l)
}

func (cp *CodecParameters) ColorPrimaries() ColorPrimaries {
	return ColorPrimaries(cp.c.color_primaries)
}

func (cp *CodecParameters) SetColorPrimaries(p ColorPrimaries) {
	cp.c.color_primaries = uint32(p)
}

func (cp *CodecParameters) ColorRange() ColorRange {
	return ColorRange(cp.c.color_range)
}

func (cp *CodecParameters) SetColorRange(r ColorRange) {
	cp.c.color_range = uint32(r)
}

func (cp *CodecParameters) ColorSpace() ColorSpace {
	return ColorSpace(cp.c.color_space)
}

func (cp *CodecParameters) SetColorSpace(s ColorSpace) {
	cp.c.color_space = uint32(s)
}

func (cp *CodecParameters) ColorTransferCharacteristic() ColorTransferCharacteristic {
	return ColorTransferCharacteristic(cp.c.color_trc)
}

func (cp *CodecParameters) SetColorTransferCharacteristic(c ColorTransferCharacteristic) {
	cp.c.color_trc = uint32(c)
}

func (cp *CodecParameters) FrameSize() int {
	return int(cp.c.frame_size)
}

func (cp *CodecParameters) SetFrameSize(i int) {
	cp.c.frame_size = C.int(i)
}

func (cp *CodecParameters) Height() int {
	return int(cp.c.height)
}

func (cp *CodecParameters) SetHeight(h int) {
	cp.c.height = C.int(h)
}

func (cp *CodecParameters) Level() Level {
	return Level(cp.c.level)
}

func (cp *CodecParameters) SetLevel(l Level) {
	cp.c.level = C.int(l)
}

func (cp *CodecParameters) MediaType() MediaType {
	return MediaType(cp.c.codec_type)
}

func (cp *CodecParameters) PixelFormat() PixelFormat {
	return PixelFormat(cp.c.format)
}

func (cp *CodecParameters) SetPixelFormat(f PixelFormat) {
	cp.c.format = C.int(f)
}

func (cp *CodecParameters) Profile() Profile {
	return Profile(cp.c.profile)
}

func (cp *CodecParameters) SetProfile(p Profile) {
	cp.c.profile = C.int(p)
}

func (cp *CodecParameters) SampleAspectRatio() Rational {
	return newRationalFromC(cp.c.sample_aspect_ratio)
}

func (cp *CodecParameters) SetSampleAspectRatio(r Rational) {
	cp.c.sample_aspect_ratio = r.c
}

func (cp *CodecParameters) SampleFormat() SampleFormat {
	return SampleFormat(cp.c.format)
}

func (cp *CodecParameters) SetSampleFormat(f SampleFormat) {
	cp.c.format = C.int(f)
}

func (cp *CodecParameters) SampleRate() int {
	return int(cp.c.sample_rate)
}

func (cp *CodecParameters) SetSampleRate(r int) {
	cp.c.sample_rate = C.int(r)
}

func (cp *CodecParameters) Width() int {
	return int(cp.c.width)
}

func (cp *CodecParameters) SetWidth(w int) {
	cp.c.width = C.int(w)
}

func (cp *CodecParameters) BitsPerCodedSample() int {
	return int(cp.c.bits_per_coded_sample)
}

func (cp *CodecParameters) SetBitsPerCodedSample(b int) {
	cp.c.bits_per_coded_sample = C.int(b)
}

func (cp *CodecParameters) BitsRerRawSample() int {
	return int(cp.c.bits_per_raw_sample)
}

func (cp *CodecParameters) SetBitsRerRawSample(b int) {
	cp.c.bits_per_raw_sample = C.int(b)
}

func (cp *CodecParameters) VideoDelay() int {
	return int(cp.c.video_delay)
}

func (cp *CodecParameters) SetVideoDelay(b int) {
	cp.c.video_delay = C.int(b)
}

func (cp *CodecParameters) BlockAlign() int {
	return int(cp.c.block_align)
}

func (cp *CodecParameters) SetBlockAlign(b int) {
	cp.c.block_align = C.int(b)
}

func (cp *CodecParameters) InitialPadding() int {
	return int(cp.c.initial_padding)
}

func (cp *CodecParameters) SetInitialPadding(b int) {
	cp.c.initial_padding = C.int(b)
}

func (cp *CodecParameters) TrailingPadding() int {
	return int(cp.c.trailing_padding)
}

func (cp *CodecParameters) SetTrailingPadding(b int) {
	cp.c.trailing_padding = C.int(b)
}

func (cp *CodecParameters) SeekPreroll() int {
	return int(cp.c.seek_preroll)
}

func (cp *CodecParameters) SetSeekPreroll(b int) {
	cp.c.seek_preroll = C.int(b)
}

// TODO: to enum
func (cp *CodecParameters) FieldOrder() uint32 {
	return uint32(cp.c.field_order)
}

// TODO: to enum
func (cp *CodecParameters) SetFieldOrder(b uint32) {
	cp.c.field_order = b
}

// TODO: to normal view
func (cp *CodecParameters) GetExtraData() []byte {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cp.c.extradata)),
		Len:  int(cp.c.extradata_size),
		Cap:  int(cp.c.extradata_size),
	}

	return *((*[]byte)(unsafe.Pointer(&header)))
}

// TODO: to normal view
func (cp *CodecParameters) SetExtraData(data []byte) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&data))

	cp.c.extradata = (*C.uint8_t)(unsafe.Pointer(header.Data))
	cp.c.extradata_size = C.int(header.Len)
}

func (cp *CodecParameters) FromCodecContext(cc *CodecContext) error {
	return newError(C.avcodec_parameters_from_context(cp.c, cc.c))
}

func (cp *CodecParameters) ToCodecContext(cc *CodecContext) error {
	return newError(C.avcodec_parameters_to_context(cc.c, cp.c))
}

func (cp *CodecParameters) Copy(dst *CodecParameters) error {
	return newError(C.avcodec_parameters_copy(dst.c, cp.c))
}
