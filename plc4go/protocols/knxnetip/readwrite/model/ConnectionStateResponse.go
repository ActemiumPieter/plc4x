/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ConnectionStateResponse is the corresponding interface of ConnectionStateResponse
type ConnectionStateResponse interface {
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetStatus returns Status (property field)
	GetStatus() Status
}

// ConnectionStateResponseExactly can be used when we want exactly this type and not a type which fulfills ConnectionStateResponse.
// This is useful for switch cases.
type ConnectionStateResponseExactly interface {
	ConnectionStateResponse
	isConnectionStateResponse() bool
}

// _ConnectionStateResponse is the data-structure of this message
type _ConnectionStateResponse struct {
	*_KnxNetIpMessage
	CommunicationChannelId uint8
	Status                 Status
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionStateResponse) GetMsgType() uint16 {
	return 0x0208
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionStateResponse) InitializeParent(parent KnxNetIpMessage) {}

func (m *_ConnectionStateResponse) GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectionStateResponse) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_ConnectionStateResponse) GetStatus() Status {
	return m.Status
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConnectionStateResponse factory function for _ConnectionStateResponse
func NewConnectionStateResponse(communicationChannelId uint8, status Status) *_ConnectionStateResponse {
	_result := &_ConnectionStateResponse{
		CommunicationChannelId: communicationChannelId,
		Status:                 status,
		_KnxNetIpMessage:       NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastConnectionStateResponse(structType interface{}) ConnectionStateResponse {
	if casted, ok := structType.(ConnectionStateResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionStateResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionStateResponse) GetTypeName() string {
	return "ConnectionStateResponse"
}

func (m *_ConnectionStateResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ConnectionStateResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ConnectionStateResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ConnectionStateResponseParse(readBuffer utils.ReadBuffer) (ConnectionStateResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionStateResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionStateResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (communicationChannelId)
	_communicationChannelId, _communicationChannelIdErr := readBuffer.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field of ConnectionStateResponse")
	}
	communicationChannelId := _communicationChannelId

	// Simple Field (status)
	if pullErr := readBuffer.PullContext("status"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for status")
	}
	_status, _statusErr := StatusParse(readBuffer)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of ConnectionStateResponse")
	}
	status := _status
	if closeErr := readBuffer.CloseContext("status"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for status")
	}

	if closeErr := readBuffer.CloseContext("ConnectionStateResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionStateResponse")
	}

	// Create a partially initialized instance
	_child := &_ConnectionStateResponse{
		_KnxNetIpMessage:       &_KnxNetIpMessage{},
		CommunicationChannelId: communicationChannelId,
		Status:                 status,
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_ConnectionStateResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionStateResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionStateResponse")
		}

		// Simple Field (communicationChannelId)
		communicationChannelId := uint8(m.GetCommunicationChannelId())
		_communicationChannelIdErr := writeBuffer.WriteUint8("communicationChannelId", 8, (communicationChannelId))
		if _communicationChannelIdErr != nil {
			return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
		}

		// Simple Field (status)
		if pushErr := writeBuffer.PushContext("status"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for status")
		}
		_statusErr := writeBuffer.WriteSerializable(m.GetStatus())
		if popErr := writeBuffer.PopContext("status"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for status")
		}
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		if popErr := writeBuffer.PopContext("ConnectionStateResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionStateResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ConnectionStateResponse) isConnectionStateResponse() bool {
	return true
}

func (m *_ConnectionStateResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
