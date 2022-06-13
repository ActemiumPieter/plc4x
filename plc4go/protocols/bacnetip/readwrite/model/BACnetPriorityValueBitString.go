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

// BACnetPriorityValueBitString is the data-structure of this message
type BACnetPriorityValueBitString struct {
	*BACnetPriorityValue
	BitStringValue *BACnetApplicationTagBitString

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

// IBACnetPriorityValueBitString is the corresponding interface of BACnetPriorityValueBitString
type IBACnetPriorityValueBitString interface {
	IBACnetPriorityValue
	// GetBitStringValue returns BitStringValue (property field)
	GetBitStringValue() *BACnetApplicationTagBitString
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

func (m *BACnetPriorityValueBitString) InitializeParent(parent *BACnetPriorityValue, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPriorityValue.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPriorityValueBitString) GetParent() *BACnetPriorityValue {
	return m.BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPriorityValueBitString) GetBitStringValue() *BACnetApplicationTagBitString {
	return m.BitStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueBitString factory function for BACnetPriorityValueBitString
func NewBACnetPriorityValueBitString(bitStringValue *BACnetApplicationTagBitString, peekedTagHeader *BACnetTagHeader, objectTypeArgument BACnetObjectType) *BACnetPriorityValueBitString {
	_result := &BACnetPriorityValueBitString{
		BitStringValue:      bitStringValue,
		BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPriorityValueBitString(structType interface{}) *BACnetPriorityValueBitString {
	if casted, ok := structType.(BACnetPriorityValueBitString); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPriorityValueBitString); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPriorityValue); ok {
		return CastBACnetPriorityValueBitString(casted.Child)
	}
	if casted, ok := structType.(*BACnetPriorityValue); ok {
		return CastBACnetPriorityValueBitString(casted.Child)
	}
	return nil
}

func (m *BACnetPriorityValueBitString) GetTypeName() string {
	return "BACnetPriorityValueBitString"
}

func (m *BACnetPriorityValueBitString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPriorityValueBitString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (bitStringValue)
	lengthInBits += m.BitStringValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPriorityValueBitString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityValueBitStringParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (*BACnetPriorityValueBitString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueBitString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueBitString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bitStringValue)
	if pullErr := readBuffer.PullContext("bitStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bitStringValue")
	}
	_bitStringValue, _bitStringValueErr := BACnetApplicationTagParse(readBuffer)
	if _bitStringValueErr != nil {
		return nil, errors.Wrap(_bitStringValueErr, "Error parsing 'bitStringValue' field")
	}
	bitStringValue := CastBACnetApplicationTagBitString(_bitStringValue)
	if closeErr := readBuffer.CloseContext("bitStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bitStringValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueBitString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueBitString")
	}

	// Create a partially initialized instance
	_child := &BACnetPriorityValueBitString{
		BitStringValue:      CastBACnetApplicationTagBitString(bitStringValue),
		BACnetPriorityValue: &BACnetPriorityValue{},
	}
	_child.BACnetPriorityValue.Child = _child
	return _child, nil
}

func (m *BACnetPriorityValueBitString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueBitString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueBitString")
		}

		// Simple Field (bitStringValue)
		if pushErr := writeBuffer.PushContext("bitStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bitStringValue")
		}
		_bitStringValueErr := writeBuffer.WriteSerializable(m.BitStringValue)
		if popErr := writeBuffer.PopContext("bitStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bitStringValue")
		}
		if _bitStringValueErr != nil {
			return errors.Wrap(_bitStringValueErr, "Error serializing 'bitStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueBitString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueBitString")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPriorityValueBitString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
