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

// BACnetApplicationTagUnsignedInteger is the data-structure of this message
type BACnetApplicationTagUnsignedInteger struct {
	*BACnetApplicationTag
	Payload *BACnetTagPayloadUnsignedInteger
}

// IBACnetApplicationTagUnsignedInteger is the corresponding interface of BACnetApplicationTagUnsignedInteger
type IBACnetApplicationTagUnsignedInteger interface {
	IBACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() *BACnetTagPayloadUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint64
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetApplicationTagUnsignedInteger) InitializeParent(parent *BACnetApplicationTag, header *BACnetTagHeader) {
	m.BACnetApplicationTag.Header = header
}

func (m *BACnetApplicationTagUnsignedInteger) GetParent() *BACnetApplicationTag {
	return m.BACnetApplicationTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetApplicationTagUnsignedInteger) GetPayload() *BACnetTagPayloadUnsignedInteger {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetApplicationTagUnsignedInteger) GetActualValue() uint64 {
	return uint64(m.GetPayload().GetActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetApplicationTagUnsignedInteger factory function for BACnetApplicationTagUnsignedInteger
func NewBACnetApplicationTagUnsignedInteger(payload *BACnetTagPayloadUnsignedInteger, header *BACnetTagHeader) *BACnetApplicationTagUnsignedInteger {
	_result := &BACnetApplicationTagUnsignedInteger{
		Payload:              payload,
		BACnetApplicationTag: NewBACnetApplicationTag(header),
	}
	_result.Child = _result
	return _result
}

func CastBACnetApplicationTagUnsignedInteger(structType interface{}) *BACnetApplicationTagUnsignedInteger {
	if casted, ok := structType.(BACnetApplicationTagUnsignedInteger); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetApplicationTagUnsignedInteger); ok {
		return casted
	}
	if casted, ok := structType.(BACnetApplicationTag); ok {
		return CastBACnetApplicationTagUnsignedInteger(casted.Child)
	}
	if casted, ok := structType.(*BACnetApplicationTag); ok {
		return CastBACnetApplicationTagUnsignedInteger(casted.Child)
	}
	return nil
}

func (m *BACnetApplicationTagUnsignedInteger) GetTypeName() string {
	return "BACnetApplicationTagUnsignedInteger"
}

func (m *BACnetApplicationTagUnsignedInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetApplicationTagUnsignedInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetApplicationTagUnsignedInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetApplicationTagUnsignedIntegerParse(readBuffer utils.ReadBuffer, header *BACnetTagHeader) (*BACnetApplicationTagUnsignedInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagUnsignedInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagUnsignedInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadUnsignedIntegerParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := CastBACnetTagPayloadUnsignedInteger(_payload)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_actualValue := payload.GetActualValue()
	actualValue := uint64(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagUnsignedInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagUnsignedInteger")
	}

	// Create a partially initialized instance
	_child := &BACnetApplicationTagUnsignedInteger{
		Payload:              CastBACnetTagPayloadUnsignedInteger(payload),
		BACnetApplicationTag: &BACnetApplicationTag{},
	}
	_child.BACnetApplicationTag.Child = _child
	return _child, nil
}

func (m *BACnetApplicationTagUnsignedInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagUnsignedInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagUnsignedInteger")
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		_payloadErr := writeBuffer.WriteSerializable(m.Payload)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagUnsignedInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagUnsignedInteger")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetApplicationTagUnsignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
