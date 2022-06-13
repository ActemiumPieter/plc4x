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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// ConnectionRequestInformationTunnelConnection is the data-structure of this message
type ConnectionRequestInformationTunnelConnection struct {
	*ConnectionRequestInformation
	KnxLayer KnxLayer
}

// IConnectionRequestInformationTunnelConnection is the corresponding interface of ConnectionRequestInformationTunnelConnection
type IConnectionRequestInformationTunnelConnection interface {
	IConnectionRequestInformation
	// GetKnxLayer returns KnxLayer (property field)
	GetKnxLayer() KnxLayer
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *ConnectionRequestInformationTunnelConnection) GetConnectionType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ConnectionRequestInformationTunnelConnection) InitializeParent(parent *ConnectionRequestInformation) {
}

func (m *ConnectionRequestInformationTunnelConnection) GetParent() *ConnectionRequestInformation {
	return m.ConnectionRequestInformation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *ConnectionRequestInformationTunnelConnection) GetKnxLayer() KnxLayer {
	return m.KnxLayer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConnectionRequestInformationTunnelConnection factory function for ConnectionRequestInformationTunnelConnection
func NewConnectionRequestInformationTunnelConnection(knxLayer KnxLayer) *ConnectionRequestInformationTunnelConnection {
	_result := &ConnectionRequestInformationTunnelConnection{
		KnxLayer:                     knxLayer,
		ConnectionRequestInformation: NewConnectionRequestInformation(),
	}
	_result.Child = _result
	return _result
}

func CastConnectionRequestInformationTunnelConnection(structType interface{}) *ConnectionRequestInformationTunnelConnection {
	if casted, ok := structType.(ConnectionRequestInformationTunnelConnection); ok {
		return &casted
	}
	if casted, ok := structType.(*ConnectionRequestInformationTunnelConnection); ok {
		return casted
	}
	if casted, ok := structType.(ConnectionRequestInformation); ok {
		return CastConnectionRequestInformationTunnelConnection(casted.Child)
	}
	if casted, ok := structType.(*ConnectionRequestInformation); ok {
		return CastConnectionRequestInformationTunnelConnection(casted.Child)
	}
	return nil
}

func (m *ConnectionRequestInformationTunnelConnection) GetTypeName() string {
	return "ConnectionRequestInformationTunnelConnection"
}

func (m *ConnectionRequestInformationTunnelConnection) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ConnectionRequestInformationTunnelConnection) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (knxLayer)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *ConnectionRequestInformationTunnelConnection) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ConnectionRequestInformationTunnelConnectionParse(readBuffer utils.ReadBuffer) (*ConnectionRequestInformationTunnelConnection, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionRequestInformationTunnelConnection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionRequestInformationTunnelConnection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (knxLayer)
	if pullErr := readBuffer.PullContext("knxLayer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for knxLayer")
	}
	_knxLayer, _knxLayerErr := KnxLayerParse(readBuffer)
	if _knxLayerErr != nil {
		return nil, errors.Wrap(_knxLayerErr, "Error parsing 'knxLayer' field")
	}
	knxLayer := _knxLayer
	if closeErr := readBuffer.CloseContext("knxLayer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for knxLayer")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	if closeErr := readBuffer.CloseContext("ConnectionRequestInformationTunnelConnection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionRequestInformationTunnelConnection")
	}

	// Create a partially initialized instance
	_child := &ConnectionRequestInformationTunnelConnection{
		KnxLayer:                     knxLayer,
		ConnectionRequestInformation: &ConnectionRequestInformation{},
	}
	_child.ConnectionRequestInformation.Child = _child
	return _child, nil
}

func (m *ConnectionRequestInformationTunnelConnection) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionRequestInformationTunnelConnection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionRequestInformationTunnelConnection")
		}

		// Simple Field (knxLayer)
		if pushErr := writeBuffer.PushContext("knxLayer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for knxLayer")
		}
		_knxLayerErr := writeBuffer.WriteSerializable(m.KnxLayer)
		if popErr := writeBuffer.PopContext("knxLayer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for knxLayer")
		}
		if _knxLayerErr != nil {
			return errors.Wrap(_knxLayerErr, "Error serializing 'knxLayer' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("ConnectionRequestInformationTunnelConnection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionRequestInformationTunnelConnection")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ConnectionRequestInformationTunnelConnection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
