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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestRequestKey is the data-structure of this message
type BACnetConfirmedServiceRequestRequestKey struct {
	*BACnetConfirmedServiceRequest
	BytesOfRemovedService []byte

	// Arguments.
	ServiceRequestLength        uint16
	ServiceRequestPayloadLength uint16
}

// IBACnetConfirmedServiceRequestRequestKey is the corresponding interface of BACnetConfirmedServiceRequestRequestKey
type IBACnetConfirmedServiceRequestRequestKey interface {
	IBACnetConfirmedServiceRequest
	// GetBytesOfRemovedService returns BytesOfRemovedService (property field)
	GetBytesOfRemovedService() []byte
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

func (m *BACnetConfirmedServiceRequestRequestKey) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_REQUEST_KEY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestRequestKey) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestRequestKey) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestRequestKey) GetBytesOfRemovedService() []byte {
	return m.BytesOfRemovedService
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestRequestKey factory function for BACnetConfirmedServiceRequestRequestKey
func NewBACnetConfirmedServiceRequestRequestKey(bytesOfRemovedService []byte, serviceRequestLength uint16, serviceRequestPayloadLength uint16) *BACnetConfirmedServiceRequestRequestKey {
	_result := &BACnetConfirmedServiceRequestRequestKey{
		BytesOfRemovedService:         bytesOfRemovedService,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestRequestKey(structType interface{}) *BACnetConfirmedServiceRequestRequestKey {
	if casted, ok := structType.(BACnetConfirmedServiceRequestRequestKey); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestRequestKey); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestRequestKey(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestRequestKey(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestRequestKey) GetTypeName() string {
	return "BACnetConfirmedServiceRequestRequestKey"
}

func (m *BACnetConfirmedServiceRequestRequestKey) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestRequestKey) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.BytesOfRemovedService) > 0 {
		lengthInBits += 8 * uint16(len(m.BytesOfRemovedService))
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestRequestKey) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestRequestKeyParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16, serviceRequestPayloadLength uint16) (*BACnetConfirmedServiceRequestRequestKey, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestRequestKey"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestRequestKey")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (bytesOfRemovedService)
	numberOfBytesbytesOfRemovedService := int(serviceRequestPayloadLength)
	bytesOfRemovedService, _readArrayErr := readBuffer.ReadByteArray("bytesOfRemovedService", numberOfBytesbytesOfRemovedService)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'bytesOfRemovedService' field")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestRequestKey"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestRequestKey")
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestRequestKey{
		BytesOfRemovedService:         bytesOfRemovedService,
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestRequestKey) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestRequestKey"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestRequestKey")
		}

		// Array Field (bytesOfRemovedService)
		if m.BytesOfRemovedService != nil {
			// Byte Array field (bytesOfRemovedService)
			_writeArrayErr := writeBuffer.WriteByteArray("bytesOfRemovedService", m.BytesOfRemovedService)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'bytesOfRemovedService' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestRequestKey"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestRequestKey")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestRequestKey) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
