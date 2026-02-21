package attribution

import (
	"bytes"
	"errors"
	"testing"
)

func testBuilderCode() [20]byte {
	var code [20]byte
	copy(code[:], "test-builder-code-00")
	return code
}

func testEncoder() AttributionEncoder {
	return NewEncoder(EncoderConfig{
		BuilderCode: testBuilderCode(),
	})
}

func TestEncode_Success(t *testing.T) {
	enc := testEncoder()
	calldata := []byte{0x01, 0x02, 0x03, 0x04}

	encoded, err := enc.Encode(calldata)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should be original + 4 magic + 20 builder code = 4 + 24 = 28 bytes.
	expected := len(calldata) + len(AttributionMagic) + BuilderCodeLength
	if len(encoded) != expected {
		t.Errorf("expected encoded length %d, got %d", expected, len(encoded))
	}

	// Original calldata should be at the start.
	if !bytes.Equal(encoded[:len(calldata)], calldata) {
		t.Error("original calldata should be preserved at start of encoded data")
	}
}

func TestEncode_EmptyCalldata(t *testing.T) {
	enc := testEncoder()

	encoded, err := enc.Encode([]byte{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should be just the magic + builder code (24 bytes).
	expected := len(AttributionMagic) + BuilderCodeLength
	if len(encoded) != expected {
		t.Errorf("expected %d bytes, got %d", expected, len(encoded))
	}
}

func TestDecode_Success(t *testing.T) {
	enc := testEncoder()
	original := []byte{0xde, 0xad, 0xbe, 0xef}

	encoded, err := enc.Encode(original)
	if err != nil {
		t.Fatalf("encode failed: %v", err)
	}

	attr, err := enc.Decode(encoded)
	if err != nil {
		t.Fatalf("decode failed: %v", err)
	}

	if attr == nil {
		t.Fatal("expected attribution, got nil")
	}

	expected := testBuilderCode()
	if attr.BuilderCode != expected {
		t.Errorf("expected builder code %v, got %v", expected, attr.BuilderCode)
	}
}

func TestRoundtrip(t *testing.T) {
	tests := []struct {
		name     string
		calldata []byte
	}{
		{name: "empty calldata", calldata: []byte{}},
		{name: "small calldata", calldata: []byte{0x01, 0x02}},
		{name: "function selector", calldata: []byte{0xa9, 0x05, 0x9c, 0xbb, 0x00, 0x00, 0x00, 0x01}},
		{name: "large calldata", calldata: make([]byte, 256)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := testEncoder()

			encoded, err := enc.Encode(tt.calldata)
			if err != nil {
				t.Fatalf("encode error: %v", err)
			}

			attr, err := enc.Decode(encoded)
			if err != nil {
				t.Fatalf("decode error: %v", err)
			}

			expected := testBuilderCode()
			if attr.BuilderCode != expected {
				t.Errorf("roundtrip builder code mismatch: expected %v, got %v", expected, attr.BuilderCode)
			}

			// Verify the original calldata is preserved (before the 24-byte suffix).
			originalPart := encoded[:len(encoded)-len(AttributionMagic)-BuilderCodeLength]
			if !bytes.Equal(originalPart, tt.calldata) {
				t.Error("original calldata not preserved after roundtrip")
			}
		})
	}
}

func TestDecode_NoAttribution(t *testing.T) {
	enc := testEncoder()

	// Calldata without attribution magic.
	calldata := make([]byte, 32)
	for i := range calldata {
		calldata[i] = 0xff
	}

	_, err := enc.Decode(calldata)
	if err == nil {
		t.Fatal("expected error for calldata without attribution")
	}
	if !errors.Is(err, ErrNoAttribution) {
		t.Errorf("expected ErrNoAttribution, got %v", err)
	}
}

func TestDecode_TooShort(t *testing.T) {
	enc := testEncoder()

	// Calldata shorter than the minimum attribution suffix (24 bytes).
	calldata := []byte{0x01, 0x02, 0x03}

	_, err := enc.Decode(calldata)
	if err == nil {
		t.Fatal("expected error for too-short calldata")
	}
	if !errors.Is(err, ErrInvalidCalldata) {
		t.Errorf("expected ErrInvalidCalldata, got %v", err)
	}
}

func TestBuilderCodeHex(t *testing.T) {
	var code [20]byte
	for i := range code {
		code[i] = byte(i)
	}

	attr := Attribution{BuilderCode: code}
	hex := attr.BuilderCodeHex()

	if len(hex) != 42 { // "0x" + 40 hex chars
		t.Errorf("expected 42 char hex string, got %d: %s", len(hex), hex)
	}
	if hex[:2] != "0x" {
		t.Errorf("expected 0x prefix, got %s", hex[:2])
	}
}
