package v1

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

/* A packet header is composed of 8 bytes in this order:
 *
 * Byte Offset  Description      Bits  Values
 * 0            Packet Type      8     1 - action frame, 2 - payload frame.
 * 1            Payload Format   8     1 - JSON object, 2 - UTF8-encoded string, 3 - Node Buffer.
 * 2            Deflated         8     0 - uncompressed, 1 - compressed / deflated (zlib-based compression).
 * 3            Unknown          8     Always 0. Possibly reserved for future use by Ubiquiti?
 * 4-7          Payload Size:    32    Size of payload in network-byte order (big endian).
 */

const (
	HeaderSize                    = 8
	HeaderPacketPosition          = 0
	HeaderPayloadFormatPosition   = 1
	HeaderPayloadDeflatedPosition = 2
	HeaderPayloadSizePosition     = 4
)

type frameType int

const (
	ActionFrameType frameType = iota + 1
	PayloadFrameType
)

type format int

const (
	JSONFormat format = iota + 1
	StringFormat
	Buffer
)

type WsFrame struct {
	data   []byte
	format format
}

type WsAction struct {
	Action      string `json:"action"`
	NewUpdateID string `json:"newUpdateId"`
	ModelKey    string `json:"modelKey"`
	ID          string `json:"id"`
}

type Message struct {
	Action  WsFrame
	Payload WsFrame
}

func DecodeWsMessage(data []byte) (*Message, error) {
	firstPayloadSize := binary.BigEndian.Uint32(data[HeaderPayloadSizePosition:])
	secondPayloadSize := binary.BigEndian.Uint32(data[firstPayloadSize+HeaderSize+HeaderPayloadSizePosition:])

	if len(data) != int(firstPayloadSize+secondPayloadSize+(2*HeaderSize)) { //nolint:gomnd // by spec
		return nil, errors.New("invalid packet size")
	}

	actionFrame, err := decodeFrame(data[0:firstPayloadSize+HeaderSize], ActionFrameType)

	if err != nil {
		return nil, err
	}

	payloadFrame, err := decodeFrame(data[firstPayloadSize+HeaderSize:], PayloadFrameType)

	if err != nil {
		return nil, err
	}

	return &Message{
		Action:  *actionFrame,
		Payload: *payloadFrame,
	}, nil
}

func decodeFrame(data []byte, frameType frameType) (*WsFrame, error) {
	dataFrameType := data[HeaderPacketPosition]
	if dataFrameType != byte(frameType) {
		return nil, errors.New("invalid frame type")
	}

	payloadDeflated := data[HeaderPayloadDeflatedPosition] == 1

	payload := data[HeaderSize:]

	if payloadDeflated {
		r, err := zlib.NewReader(bytes.NewBuffer(payload))

		if err != nil {
			return nil, err
		}
		defer func() {
			_ = r.Close()
		}()

		newBuffer := bytes.Buffer{}

		_, err = io.Copy(&newBuffer, r) //nolint:gosec // not exposed
		if err != nil {
			return nil, err
		}

		payload = newBuffer.Bytes()
	}

	return &WsFrame{data: payload, format: format(data[HeaderPayloadFormatPosition])}, nil
}

func (f WsFrame) GetJSON(out interface{}) error {
	if f.format != JSONFormat {
		return errors.New("invalid format for action")
	}
	err := json.Unmarshal(f.data, out)
	if err != nil {
		return err
	}
	return nil
}

func (f WsFrame) String() string {
	return string(f.data)
}

func (f WsFrame) GetRAWData() []byte {
	tmp := make([]byte, len(f.data))
	copy(tmp, f.data)
	return tmp
}

func (m Message) GetAction() (*WsAction, error) {
	action := &WsAction{}

	if err := m.Action.GetJSON(action); err != nil {
		return nil, err
	}

	return action, nil
}

func (m Message) String() string {
	action, _ := m.GetAction()
	return fmt.Sprintf("%s/%s: %s", action.Action, action.ModelKey, m.Payload)
}
