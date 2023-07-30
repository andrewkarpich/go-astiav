package astiav

import "C"
import (
	"errors"
	"unsafe"
)

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
/*
uint8_t *astiavPacketPack(AVPacket *pkt, int *size) {
	int i, side_data_size = 0;

	side_data_size = sizeof(pkt->side_data_elems);
  	for (i = 0; i < pkt->side_data_elems; i++) {
		side_data_size += sizeof(pkt->side_data[i].type) + sizeof(pkt->side_data[i].size) + pkt->side_data[i].size;
	}

	*size = sizeof(pkt->size) + pkt->size + sizeof(pkt->stream_index) + sizeof(pkt->flags) + sizeof(pkt->pts) + sizeof(pkt->dts) + sizeof(pkt->duration) + sizeof(pkt->pos) + side_data_size;
	uint8_t *buf = av_malloc(*size);

	int s = 0;

	memcpy(buf, &(pkt->size), sizeof(pkt->size));
	s += sizeof(pkt->size);

	memcpy(buf + s, pkt->data, pkt->size);
	s += pkt->size;

	memcpy(buf + s, &(pkt->stream_index), sizeof(pkt->stream_index));
	s += sizeof(pkt->stream_index);

	memcpy(buf + s, &(pkt->flags), sizeof(pkt->flags));
	s += sizeof(pkt->flags);

	memcpy(buf + s, &(pkt->pts), sizeof(pkt->pts));
	s += sizeof(pkt->pts);

	memcpy(buf + s, &(pkt->dts), sizeof(pkt->dts));
	s += sizeof(pkt->dts);

	memcpy(buf + s, &(pkt->duration), sizeof(pkt->duration));
	s += sizeof(pkt->duration);

	memcpy(buf + s, &(pkt->pos), sizeof(pkt->pos));
	s += sizeof(pkt->pos);

	memcpy(buf + s, &(pkt->side_data_elems), sizeof(pkt->side_data_elems));
	s += sizeof(pkt->side_data_elems);

	for (i = 0; i < pkt->side_data_elems; i++) {
		memcpy(buf + s, &(pkt->side_data[i].type), sizeof(pkt->side_data[i].type));
		s += sizeof(pkt->side_data[i].type);

		memcpy(buf + s, &(pkt->side_data[i].size), sizeof(pkt->side_data[i].size));
		s += sizeof(pkt->side_data[i].size);

		memcpy(buf + s, pkt->side_data[i].data, pkt->side_data[i].size);
		s += pkt->side_data[i].size;
	}

	return buf;
}

AVPacket *astiavPacketUnpack(uint8_t *buf) {
	AVPacket *pkt = av_packet_alloc();

	int s = 0;

	memcpy(&(pkt->size), buf, sizeof(pkt->size));
	s += sizeof(pkt->size);

	uint8_t *dbuf = av_malloc(pkt->size);
	memcpy(dbuf, buf + s, pkt->size);
	av_packet_from_data(pkt, dbuf, pkt->size);
	s += pkt->size;

	memcpy(&(pkt->stream_index), buf + s, sizeof(pkt->stream_index));
	s += sizeof(pkt->stream_index);

	memcpy(&(pkt->flags), buf + s, sizeof(pkt->flags));
	s += sizeof(pkt->flags);

	memcpy(&(pkt->pts), buf + s, sizeof(pkt->pts));
	s += sizeof(pkt->pts);

	memcpy(&(pkt->dts), buf + s, sizeof(pkt->dts));
	s += sizeof(pkt->dts);

	memcpy(&(pkt->duration), buf + s, sizeof(pkt->duration));
	s += sizeof(pkt->duration);

	memcpy(&(pkt->pos), buf + s, sizeof(pkt->pos));
	s += sizeof(pkt->pos);

	int side_data_elems = 0;
	memcpy(&side_data_elems, buf + s, sizeof(side_data_elems));
	s += sizeof(side_data_elems);

	int i = 0;
	for (i = 0; i < side_data_elems; i++) {
		enum AVPacketSideDataType side_data_type;
		memcpy(&side_data_type, buf + s, sizeof(side_data_type));
		s += sizeof(side_data_type);

		int side_data_size = 0;
		memcpy(&side_data_size, buf + s, sizeof(side_data_size));
		s += sizeof(side_data_size);

		uint8_t *sdbuf = av_packet_new_side_data(pkt, side_data_type, side_data_size);
		memcpy(sdbuf, buf + s, side_data_size);
		s += side_data_size;
	}

	return pkt;
}
*/
import "C"

// https://github.com/FFmpeg/FFmpeg/blob/n5.0/libavcodec/packet.h#L350
type Packet struct {
	c *C.struct_AVPacket
}

func newPacketFromC(c *C.struct_AVPacket) *Packet {
	if c == nil {
		return nil
	}
	return &Packet{c: c}
}

func AllocPacket() *Packet {
	return newPacketFromC(C.av_packet_alloc())
}

func PacketUnpack(b []byte) (*Packet, error) {
	//var ptr *C.uint8_t
	//if b != nil {
	//	c := make([]byte, len(b))
	//	copy(c, b)
	//	ptr = (*C.uint8_t)(unsafe.Pointer(&c[0]))
	//}
	//
	//return newError(C.astiavPacketUnpack(p.c, ptr))

	////////

	buf := C.av_malloc(C.size_t(len(b)))
	if buf == nil {
		return nil, errors.New("astiav: allocating buffer failed")
	}
	C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&b[0]), C.size_t(len(b)))
	// From data
	return newPacketFromC(C.astiavPacketUnpack((*C.uint8_t)(unsafe.Pointer(buf)))), nil
}

func (p *Packet) Data() []byte {
	if p.c.data == nil {
		return nil
	}
	return bytesFromC(
		func(size *C.int) *C.uint8_t {
			*size = p.c.size
			return p.c.data
		},
	)
}

func (p *Packet) Dts() int64 {
	return int64(p.c.dts)
}

func (p *Packet) SetDts(v int64) {
	p.c.dts = C.int64_t(v)
}

func (p *Packet) Duration() int64 {
	return int64(p.c.duration)
}

func (p *Packet) SetDuration(d int64) {
	p.c.duration = C.int64_t(d)
}

func (p *Packet) Flags() PacketFlags {
	return PacketFlags(p.c.flags)
}

func (p *Packet) SetFlags(f PacketFlags) {
	p.c.flags = C.int(f)
}

func (p *Packet) Pos() int64 {
	return int64(p.c.pos)
}

func (p *Packet) SetPos(v int64) {
	p.c.pos = C.int64_t(v)
}

func (p *Packet) Pts() int64 {
	return int64(p.c.pts)
}

func (p *Packet) SetPts(v int64) {
	p.c.pts = C.int64_t(v)
}

func (p *Packet) SideData(t PacketSideDataType) []byte {
	return bytesFromC(
		func(size *C.int) *C.uint8_t {
			return C.av_packet_get_side_data(p.c, (C.enum_AVPacketSideDataType)(t), size)
		},
	)
}

func (p *Packet) AddSideData(t PacketSideDataType, b []byte) error {
	//return bytesToC(
	//	b, func(b *C.uint8_t, size C.int) error {
	//		return newError(C.av_packet_add_side_data(p.c, (C.enum_AVPacketSideDataType)(t), b, C.ulong(size)))
	//	},
	//)

	///
	buf := C.av_malloc(C.size_t(len(b)))
	if buf == nil {
		return errors.New("astiav: allocating buffer failed")
	}
	C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&b[0]), C.size_t(len(b)))

	// From data
	return newError(
		C.av_packet_add_side_data(
			p.c,
			(C.enum_AVPacketSideDataType)(t),
			(*C.uint8_t)(unsafe.Pointer(buf)),
			C.size_t(len(b)),
		),
	)
}

func (p *Packet) SideDataElems() int {
	return int(p.c.side_data_elems)
}

func (p *Packet) Size() int {
	return int(p.c.size)
}

func (p *Packet) SetSize(s int) {
	p.c.size = C.int(s)
}

func (p *Packet) StreamIndex() int {
	return int(p.c.stream_index)
}

func (p *Packet) SetStreamIndex(i int) {
	p.c.stream_index = C.int(i)
}

func (p *Packet) Free() {
	C.av_packet_free(&p.c)
}

func (p *Packet) Clone() *Packet {
	return newPacketFromC(C.av_packet_clone(p.c))
}

func (p *Packet) AllocPayload(s int) error {
	return newError(C.av_new_packet(p.c, C.int(s)))
}

func (p *Packet) Ref(src *Packet) error {
	return newError(C.av_packet_ref(p.c, src.c))
}

func (p *Packet) Unref() {
	C.av_packet_unref(p.c)
}

func (p *Packet) MoveRef(src *Packet) {
	C.av_packet_move_ref(p.c, src.c)
}

func (p *Packet) RescaleTs(src, dst Rational) {
	C.av_packet_rescale_ts(p.c, src.c, dst.c)
}

func (p *Packet) FromData(data []byte) error {
	// Create buf
	buf := (*C.uint8_t)(C.av_malloc(C.size_t(len(data))))
	if buf == nil {
		return errors.New("astiav: allocating buffer failed")
	}
	C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&data[0]), C.size_t(len(data)))

	// From data
	return newError(C.av_packet_from_data(p.c, buf, C.int(len(data))))
}

func (p *Packet) Pack() []byte {
	return bytesFromC(
		func(size *C.int) *C.uint8_t {
			return C.astiavPacketPack(p.c, size)
		},
	)
}
