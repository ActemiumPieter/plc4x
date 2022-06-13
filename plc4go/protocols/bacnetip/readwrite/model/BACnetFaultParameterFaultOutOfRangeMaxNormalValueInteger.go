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

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger is the data-structure of this message
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger struct {
	*BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	IntegerValue *BACnetApplicationTagSignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger
type IBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger interface {
	IBACnetFaultParameterFaultOutOfRangeMaxNormalValue
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() *BACnetApplicationTagSignedInteger
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

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) InitializeParent(parent *BACnetFaultParameterFaultOutOfRangeMaxNormalValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetFaultParameterFaultOutOfRangeMaxNormalValue.OpeningTag = openingTag
	m.BACnetFaultParameterFaultOutOfRangeMaxNormalValue.PeekedTagHeader = peekedTagHeader
	m.BACnetFaultParameterFaultOutOfRangeMaxNormalValue.ClosingTag = closingTag
}

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetParent() *BACnetFaultParameterFaultOutOfRangeMaxNormalValue {
	return m.BACnetFaultParameterFaultOutOfRangeMaxNormalValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetIntegerValue() *BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger factory function for BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger(integerValue *BACnetApplicationTagSignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger {
	_result := &BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger{
		IntegerValue: integerValue,
		BACnetFaultParameterFaultOutOfRangeMaxNormalValue: NewBACnetFaultParameterFaultOutOfRangeMaxNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger(structType interface{}) *BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger); ok {
		return casted
	}
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValue); ok {
		return CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger(casted.Child)
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValue); ok {
		return CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger(casted.Child)
	}
	return nil
}

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"
}

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultOutOfRangeMaxNormalValueIntegerParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (integerValue)
	if pullErr := readBuffer.PullContext("integerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for integerValue")
	}
	_integerValue, _integerValueErr := BACnetApplicationTagParse(readBuffer)
	if _integerValueErr != nil {
		return nil, errors.Wrap(_integerValueErr, "Error parsing 'integerValue' field")
	}
	integerValue := CastBACnetApplicationTagSignedInteger(_integerValue)
	if closeErr := readBuffer.CloseContext("integerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for integerValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger")
	}

	// Create a partially initialized instance
	_child := &BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger{
		IntegerValue: CastBACnetApplicationTagSignedInteger(integerValue),
		BACnetFaultParameterFaultOutOfRangeMaxNormalValue: &BACnetFaultParameterFaultOutOfRangeMaxNormalValue{},
	}
	_child.BACnetFaultParameterFaultOutOfRangeMaxNormalValue.Child = _child
	return _child, nil
}

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger")
		}

		// Simple Field (integerValue)
		if pushErr := writeBuffer.PushContext("integerValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for integerValue")
		}
		_integerValueErr := writeBuffer.WriteSerializable(m.IntegerValue)
		if popErr := writeBuffer.PopContext("integerValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for integerValue")
		}
		if _integerValueErr != nil {
			return errors.Wrap(_integerValueErr, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
