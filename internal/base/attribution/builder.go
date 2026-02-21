package attribution

import (
	"bytes"
	"fmt"
	"time"
)

// AttributionEncoder defines operations for embedding and extracting ERC-8021
// builder attribution codes in transaction calldata.
type AttributionEncoder interface {
	// Encode appends the configured builder code to the provided calldata.
	// The builder code occupies the last 20 bytes of the resulting calldata,
	// preceded by the 4-byte AttributionMagic marker.
	Encode(calldata []byte) ([]byte, error)

	// Decode extracts the ERC-8021 builder code from the last 24 bytes of
	// calldata (4 magic bytes + 20 builder code bytes).
	// Returns ErrNoAttribution if the magic marker is not present.
	// Returns ErrInvalidCalldata if calldata is too short.
	Decode(calldata []byte) (*Attribution, error)
}

// EncoderConfig holds configuration for the ERC-8021 attribution encoder.
type EncoderConfig struct {
	// BuilderCode is the 20-byte Ethereum address identifying the builder.
	// This is appended to all outgoing transaction calldata.
	BuilderCode [20]byte
}

// encoder implements AttributionEncoder for ERC-8021 calldata manipulation.
type encoder struct {
	cfg EncoderConfig
}

// NewEncoder creates an AttributionEncoder with the given builder configuration.
func NewEncoder(cfg EncoderConfig) AttributionEncoder {
	return &encoder{cfg: cfg}
}

// Encode appends the ERC-8021 builder attribution to the provided calldata.
// The output format is: [original calldata] [4-byte magic] [20-byte builder code]
// This adds 24 bytes to the calldata length.
func (e *encoder) Encode(calldata []byte) ([]byte, error) {
	if len(e.cfg.BuilderCode) != BuilderCodeLength {
		return nil, fmt.Errorf("attribution: %w", ErrInvalidBuilderCode)
	}

	result := make([]byte, 0, len(calldata)+len(AttributionMagic)+BuilderCodeLength)
	result = append(result, calldata...)
	result = append(result, []byte(AttributionMagic)...)
	result = append(result, e.cfg.BuilderCode[:]...)

	return result, nil
}

// Decode extracts the ERC-8021 builder code from transaction calldata.
// It checks the last 24 bytes for the magic marker followed by the builder code.
// Returns ErrInvalidCalldata if calldata is too short.
// Returns ErrNoAttribution if the magic marker is absent.
func (e *encoder) Decode(calldata []byte) (*Attribution, error) {
	const attributionSuffixLen = len(AttributionMagic) + BuilderCodeLength // 24 bytes

	if len(calldata) < attributionSuffixLen {
		return nil, fmt.Errorf("attribution: calldata length %d: %w", len(calldata), ErrInvalidCalldata)
	}

	// Extract the last 24 bytes.
	suffix := calldata[len(calldata)-attributionSuffixLen:]
	magic := suffix[:len(AttributionMagic)]
	code := suffix[len(AttributionMagic):]

	if !bytes.Equal(magic, []byte(AttributionMagic)) {
		return nil, fmt.Errorf("attribution: magic not found: %w", ErrNoAttribution)
	}

	var builderCode [20]byte
	copy(builderCode[:], code)

	return &Attribution{
		BuilderCode: builderCode,
		Timestamp:   time.Now(),
	}, nil
}
