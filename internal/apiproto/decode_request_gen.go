// Code generated by internal/gen/api/main.go. DO NOT EDIT.

package apiproto

import "encoding/json"

var _ RequestDecoder = (*JSONRequestDecoder)(nil)

// JSONRequestDecoder ...
type JSONRequestDecoder struct{}

// NewJSONRequestDecoder ...
func NewJSONRequestDecoder() *JSONRequestDecoder {
	return &JSONRequestDecoder{}
}

// DecodeBatch ...
func (d *JSONRequestDecoder) DecodeBatch(data []byte) (*BatchRequest, error) {
	var p BatchRequest
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// DecodePublish ...
func (d *JSONRequestDecoder) DecodePublish(data []byte) (*PublishRequest, error) {
	var p PublishRequest
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// DecodeBroadcast ...
func (d *JSONRequestDecoder) DecodeBroadcast(data []byte) (*BroadcastRequest, error) {
	var p BroadcastRequest
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// DecodeSubscribe ...
func (d *JSONRequestDecoder) DecodeSubscribe(data []byte) (*SubscribeRequest, error) {
	var p SubscribeRequest
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// DecodeUnsubscribe ...
func (d *JSONRequestDecoder) DecodeUnsubscribe(data []byte) (*UnsubscribeRequest, error) {
	var p UnsubscribeRequest
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// DecodeDisconnect ...
func (d *JSONRequestDecoder) DecodeDisconnect(data []byte) (*DisconnectRequest, error) {
	var p DisconnectRequest
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// DecodePresence ...
func (d *JSONRequestDecoder) DecodePresence(data []byte) (*PresenceRequest, error) {
	var p PresenceRequest
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// DecodePresenceStats ...
func (d *JSONRequestDecoder) DecodePresenceStats(data []byte) (*PresenceStatsRequest, error) {
	var p PresenceStatsRequest
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// DecodeHistory ...
func (d *JSONRequestDecoder) DecodeHistory(data []byte) (*HistoryRequest, error) {
	var p HistoryRequest
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// DecodeHistoryRemove ...
func (d *JSONRequestDecoder) DecodeHistoryRemove(data []byte) (*HistoryRemoveRequest, error) {
	var p HistoryRemoveRequest
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// DecodeInfo ...
func (d *JSONRequestDecoder) DecodeInfo(data []byte) (*InfoRequest, error) {
	var p InfoRequest
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// DecodeRPC ...
func (d *JSONRequestDecoder) DecodeRPC(data []byte) (*RPCRequest, error) {
	var p RPCRequest
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// DecodeRefresh ...
func (d *JSONRequestDecoder) DecodeRefresh(data []byte) (*RefreshRequest, error) {
	var p RefreshRequest
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// DecodeChannels ...
func (d *JSONRequestDecoder) DecodeChannels(data []byte) (*ChannelsRequest, error) {
	var p ChannelsRequest
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
