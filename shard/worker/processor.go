package worker;

import(
	Net    "net"
	Bytes  "bytes"
	Binary "encoding/binary"
	Adler  "hash/adler32"
	// api
	ProcV1 "github.com/PxnPub/pxnMetrics/worker/proc_v1"
);



func (worker *Worker) Process(data []byte, src *Net.Addr) ([]byte, err) {
	reader := Bytes.NewReader(data);
	var first       byte;
	var size        uint16;
	var checksum    uint16;
	var index_crypt byte;
	var index_proto byte;
	var payload     []byte;
	// check first byte
	if err := Binary.Read(reader, Binary.BigEndian, &first); err != nil { return nil, err; }
	// header version
	switch first {
	case 0x07: // 7 byte header v1
		if err := Binary.Read(reader, Binary.BigEndian, &size       ); err != nil { return nil, err; }
		if err := Binary.Read(reader, Binary.BigEndian, &checksum   ); err != nil { return nil, err; }
		if err := Binary.Read(reader, Binary.BigEndian, &index_crypt); err != nil { return nil, err; }
		if err := Binary.Read(reader, Binary.BigEndian, &index_proto); err != nil { return nil, err; }
		var buffer Bytes.Buffer;
		if _, err := buffer.ReadFrom(reader); err != nil { return nil, err; }
		payload = buffer.Bytes();
	// invalid packet (first byte)
	default: return nil, Errors.New("Invalid first byte: %X (%d)", first, first);
	}
	// packet size
	if len(payload) != int(size) { return nil, Errors.New("Invalid packet length"      ); }
	if size <    10              { return nil, Errors.New("Invalid short packet length"); }
	if size > 10000              { return nil, Errors.New("Invalid long packet length" ); }
	// checksum
	hash32 := Adler.Checksum(payload);
	hash16 := uint16(((hash32 >> 16) & 0xFFFF) ^ (hash32 & 0xFFFF)) ^ worker.ChecksumSeed;
	if checksum != hash16 { return nil, Errors.New("Invalid packet checksum"); }
	// encryption
	switch index_crypt {
	// plain text
//TODO: disable this by default
	case 0x00: break;
	// AES-GCM 128
	case 0x01: panic(Errors.New("UNFINISHED"));
	// AES-GCM 192
	case 0x02: panic(Errors.New("UNFINISHED"));
	// AES-GCM 256
	case 0x03: panic(Errors.New("UNFINISHED"));
	default: return nil, Fmt.Errorf("Invalid encryption value: %X (%d)", index_crypt, index_crypt);
	}
	// protocol version
	switch index_proto {
	// Submit V1
	case 0x01: {
		reply, err := ProcV1.Process(payload);
		if err != nil { return nil, err; }
		// send reply
		hash32 := Adler.Checksum(reply);
		hash16 := uint16(((hash32 >> 16) & 0xFFFF) ^ (hash32 & 0xFFFF)) ^ worker.ChecksumSeed;
		size := len(reply);
		out := make([]byte, size+7);
		out[0] = 0x07;                          // header size
		out[1] = byte((size   >> 16) & 0xFFFF); // payload size high
		out[2] = byte( size          & 0xFFFF); // payload size low
		out[3] = byte((hash16 >> 16) & 0xFFFF); // checksum high
		out[4] = byte( hash16        & 0xFFFF); // checksum low
		out[5] = index_crypt;                   // encryption
		out[6] = index_proto;                   // protocol version
		copy(out[7:], reply);
		return out, nil;
	}
	default: break;
	}
	return nil, Fmt.Errorf("Invalid protocol version: %X (%d)", index_proto, index_proto);
}
