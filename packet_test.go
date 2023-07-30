package astiav_test

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/asticode/go-astiav"
	"github.com/stretchr/testify/require"
)

func videoInputFirstPacket() (pkt *astiav.Packet, err error) {
	if global.pkt != nil {
		return global.pkt, nil
	}

	var fc *astiav.FormatContext
	if fc, err = videoInputFormatContext(); err != nil {
		err = fmt.Errorf("astiav_test: getting input format context failed")
		return
	}

	pkt = astiav.AllocPacket()
	if pkt == nil {
		err = errors.New("astiav_test: pkt is nil")
		return
	}
	global.closer.Add(pkt.Free)

	if err = fc.ReadFrame(pkt); err != nil {
		err = fmt.Errorf("astiav_test: reading frame failed: %w", err)
		return
	}

	global.pkt = pkt
	return
}

func TestPacket(t *testing.T) {
	pkt1, err := videoInputFirstPacket()
	require.NoError(t, err)
	require.Equal(
		t,
		[]byte{0x0, 0x0, 0x0, 0xd1, 0x65, 0x88, 0x82, 0x0, 0x1f, 0x5f, 0xff, 0xf8, 0x22, 0x8a, 0x0, 0x2, 0x2d, 0xbe, 0x38, 0xc7, 0x19, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x3a, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xba, 0xeb, 0xae, 0xb9, 0xb8, 0xe6, 0x39, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xa5, 0xc0},
		pkt1.Data(),
	)
	require.Equal(t, int64(0), pkt1.Dts())
	require.Equal(t, int64(512), pkt1.Duration())
	require.True(t, pkt1.Flags().Has(astiav.PacketFlagKey))
	require.Equal(t, int64(48), pkt1.Pos())
	require.Equal(t, int64(0), pkt1.Pts())
	require.Equal(t, 213, pkt1.Size())
	require.Equal(t, 0, pkt1.StreamIndex())

	pkt2 := astiav.AllocPacket()
	require.NotNil(t, pkt2)
	defer pkt2.Free()
	require.Nil(t, pkt2.Data())
	err = pkt2.AllocPayload(5)
	require.NoError(t, err)
	require.Len(t, pkt2.Data(), 5)
	pkt2.SetDts(1)
	pkt2.SetDuration(2)
	pkt2.SetFlags(astiav.NewPacketFlags(3))
	pkt2.SetPos(4)
	pkt2.SetPts(5)
	pkt2.SetSize(6)
	pkt2.SetStreamIndex(7)
	require.Equal(t, int64(1), pkt2.Dts())
	require.Equal(t, int64(2), pkt2.Duration())
	require.Equal(t, astiav.NewPacketFlags(3), pkt2.Flags())
	require.Equal(t, int64(4), pkt2.Pos())
	require.Equal(t, int64(5), pkt2.Pts())
	require.Equal(t, 6, pkt2.Size())
	require.Equal(t, 7, pkt2.StreamIndex())

	pkt3 := pkt1.Clone()
	require.NotNil(t, pkt3)
	defer pkt3.Free()
	require.Equal(t, int64(512), pkt3.Duration())

	err = pkt3.Ref(pkt2)
	require.NoError(t, err)
	require.Equal(t, int64(2), pkt3.Duration())

	pkt3.MoveRef(pkt1)
	require.Equal(t, int64(512), pkt3.Duration())
	require.Equal(t, int64(0), pkt1.Duration())

	pkt3.Unref()
	require.Equal(t, int64(0), pkt3.Duration())

	pkt3.SetDts(10)
	pkt3.SetDuration(20)
	pkt3.SetPts(30)
	pkt3.RescaleTs(astiav.NewRational(1, 10), astiav.NewRational(1, 1))
	require.Equal(t, int64(1), pkt3.Dts())
	require.Equal(t, int64(2), pkt3.Duration())
	require.Equal(t, int64(3), pkt3.Pts())

	pkt4 := astiav.AllocPacket()
	require.NotNil(t, pkt4)
	defer pkt4.Free()
	b := []byte("test")
	require.NoError(t, pkt4.FromData(b))
	require.Equal(t, b, pkt4.Data())

	// TODO Test SideData
}

func TestPackPacket(t *testing.T) {
	//signal.Ignore(syscall.SIGABRT)
	astiav.SetLogLevel(astiav.LogLevelDebug)

	pkt := astiav.AllocPacket()
	require.NotNil(t, pkt)

	b := []byte("test test")
	require.NoError(t, pkt.FromData(b))
	pkt.SetStreamIndex(1)
	pkt.SetFlags(astiav.NewPacketFlags(astiav.PacketFlagKey))
	pkt.SetPts(2)
	pkt.SetDts(3)
	pkt.SetDuration(4)
	pkt.SetPos(5)
	sideData1 := []byte("test1")
	err := pkt.AddSideData(astiav.PacketSideDataTypeContentLightLevel, sideData1)
	require.NoError(t, err)

	sideData2 := []byte("test2")
	err = pkt.AddSideData(astiav.PacketSideDataTypeEncryptionInfo, sideData2)
	require.NoError(t, err)

	bs := pkt.Pack()
	pkt.Free()

	pkt2, err := astiav.PacketUnpack(bs)
	require.NoError(t, err)

	require.Equal(t, b, pkt2.Data())
	require.Equal(t, 1, pkt2.StreamIndex())
	require.Equal(t, astiav.NewPacketFlags(astiav.PacketFlagKey), pkt2.Flags())
	require.Equal(t, int64(2), pkt2.Pts())
	require.Equal(t, int64(3), pkt2.Dts())
	require.Equal(t, int64(4), pkt2.Duration())
	require.Equal(t, int64(5), pkt2.Pos())
	require.Equal(t, 2, pkt2.SideDataElems())

	sideData12 := pkt2.SideData(astiav.PacketSideDataTypeContentLightLevel)
	require.Equal(t, sideData1, sideData12)

	sideData22 := pkt2.SideData(astiav.PacketSideDataTypeEncryptionInfo)
	require.Equal(t, sideData2, sideData22)

	log.Println(bs)

	pkt2.Free()
}
