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

// BACnetConstructedDataEscalatorMode is the data-structure of this message
type BACnetConstructedDataEscalatorMode struct {
	*BACnetConstructedData
	EscalatorMode *BACnetEscalatorModeTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataEscalatorMode is the corresponding interface of BACnetConstructedDataEscalatorMode
type IBACnetConstructedDataEscalatorMode interface {
	IBACnetConstructedData
	// GetEscalatorMode returns EscalatorMode (property field)
	GetEscalatorMode() *BACnetEscalatorModeTagged
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

func (m *BACnetConstructedDataEscalatorMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataEscalatorMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ESCALATOR_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataEscalatorMode) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataEscalatorMode) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataEscalatorMode) GetEscalatorMode() *BACnetEscalatorModeTagged {
	return m.EscalatorMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEscalatorMode factory function for BACnetConstructedDataEscalatorMode
func NewBACnetConstructedDataEscalatorMode(escalatorMode *BACnetEscalatorModeTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataEscalatorMode {
	_result := &BACnetConstructedDataEscalatorMode{
		EscalatorMode:         escalatorMode,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataEscalatorMode(structType interface{}) *BACnetConstructedDataEscalatorMode {
	if casted, ok := structType.(BACnetConstructedDataEscalatorMode); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEscalatorMode); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataEscalatorMode(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataEscalatorMode(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataEscalatorMode) GetTypeName() string {
	return "BACnetConstructedDataEscalatorMode"
}

func (m *BACnetConstructedDataEscalatorMode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataEscalatorMode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (escalatorMode)
	lengthInBits += m.EscalatorMode.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataEscalatorMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEscalatorModeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataEscalatorMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEscalatorMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEscalatorMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (escalatorMode)
	if pullErr := readBuffer.PullContext("escalatorMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for escalatorMode")
	}
	_escalatorMode, _escalatorModeErr := BACnetEscalatorModeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _escalatorModeErr != nil {
		return nil, errors.Wrap(_escalatorModeErr, "Error parsing 'escalatorMode' field")
	}
	escalatorMode := CastBACnetEscalatorModeTagged(_escalatorMode)
	if closeErr := readBuffer.CloseContext("escalatorMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for escalatorMode")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEscalatorMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEscalatorMode")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataEscalatorMode{
		EscalatorMode:         CastBACnetEscalatorModeTagged(escalatorMode),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataEscalatorMode) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEscalatorMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEscalatorMode")
		}

		// Simple Field (escalatorMode)
		if pushErr := writeBuffer.PushContext("escalatorMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for escalatorMode")
		}
		_escalatorModeErr := writeBuffer.WriteSerializable(m.EscalatorMode)
		if popErr := writeBuffer.PopContext("escalatorMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for escalatorMode")
		}
		if _escalatorModeErr != nil {
			return errors.Wrap(_escalatorModeErr, "Error serializing 'escalatorMode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEscalatorMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEscalatorMode")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataEscalatorMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
