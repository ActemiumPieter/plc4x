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

// BACnetConstructedDataIntervalOffset is the corresponding interface of BACnetConstructedDataIntervalOffset
type BACnetConstructedDataIntervalOffset interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIntervalOffset returns IntervalOffset (property field)
	GetIntervalOffset() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataIntervalOffsetExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIntervalOffset.
// This is useful for switch cases.
type BACnetConstructedDataIntervalOffsetExactly interface {
	BACnetConstructedDataIntervalOffset
	isBACnetConstructedDataIntervalOffset() bool
}

// _BACnetConstructedDataIntervalOffset is the data-structure of this message
type _BACnetConstructedDataIntervalOffset struct {
	*_BACnetConstructedData
	IntervalOffset BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntervalOffset) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIntervalOffset) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTERVAL_OFFSET
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntervalOffset) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIntervalOffset) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntervalOffset) GetIntervalOffset() BACnetApplicationTagUnsignedInteger {
	return m.IntervalOffset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntervalOffset) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetIntervalOffset())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIntervalOffset factory function for _BACnetConstructedDataIntervalOffset
func NewBACnetConstructedDataIntervalOffset(intervalOffset BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntervalOffset {
	_result := &_BACnetConstructedDataIntervalOffset{
		IntervalOffset:         intervalOffset,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntervalOffset(structType interface{}) BACnetConstructedDataIntervalOffset {
	if casted, ok := structType.(BACnetConstructedDataIntervalOffset); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntervalOffset); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntervalOffset) GetTypeName() string {
	return "BACnetConstructedDataIntervalOffset"
}

func (m *_BACnetConstructedDataIntervalOffset) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataIntervalOffset) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (intervalOffset)
	lengthInBits += m.IntervalOffset.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntervalOffset) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIntervalOffsetParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIntervalOffset, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntervalOffset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntervalOffset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (intervalOffset)
	if pullErr := readBuffer.PullContext("intervalOffset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for intervalOffset")
	}
	_intervalOffset, _intervalOffsetErr := BACnetApplicationTagParse(readBuffer)
	if _intervalOffsetErr != nil {
		return nil, errors.Wrap(_intervalOffsetErr, "Error parsing 'intervalOffset' field of BACnetConstructedDataIntervalOffset")
	}
	intervalOffset := _intervalOffset.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("intervalOffset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for intervalOffset")
	}

	// Virtual field
	_actualValue := intervalOffset
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntervalOffset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntervalOffset")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIntervalOffset{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		IntervalOffset: intervalOffset,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIntervalOffset) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntervalOffset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntervalOffset")
		}

		// Simple Field (intervalOffset)
		if pushErr := writeBuffer.PushContext("intervalOffset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for intervalOffset")
		}
		_intervalOffsetErr := writeBuffer.WriteSerializable(m.GetIntervalOffset())
		if popErr := writeBuffer.PopContext("intervalOffset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for intervalOffset")
		}
		if _intervalOffsetErr != nil {
			return errors.Wrap(_intervalOffsetErr, "Error serializing 'intervalOffset' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntervalOffset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntervalOffset")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntervalOffset) isBACnetConstructedDataIntervalOffset() bool {
	return true
}

func (m *_BACnetConstructedDataIntervalOffset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
