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

// BACnetConstructedDataInactiveText is the data-structure of this message
type BACnetConstructedDataInactiveText struct {
	*BACnetConstructedData
	InactiveText *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataInactiveText is the corresponding interface of BACnetConstructedDataInactiveText
type IBACnetConstructedDataInactiveText interface {
	IBACnetConstructedData
	// GetInactiveText returns InactiveText (property field)
	GetInactiveText() *BACnetApplicationTagCharacterString
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

func (m *BACnetConstructedDataInactiveText) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataInactiveText) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INACTIVE_TEXT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataInactiveText) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataInactiveText) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataInactiveText) GetInactiveText() *BACnetApplicationTagCharacterString {
	return m.InactiveText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInactiveText factory function for BACnetConstructedDataInactiveText
func NewBACnetConstructedDataInactiveText(inactiveText *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataInactiveText {
	_result := &BACnetConstructedDataInactiveText{
		InactiveText:          inactiveText,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataInactiveText(structType interface{}) *BACnetConstructedDataInactiveText {
	if casted, ok := structType.(BACnetConstructedDataInactiveText); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInactiveText); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataInactiveText(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataInactiveText(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataInactiveText) GetTypeName() string {
	return "BACnetConstructedDataInactiveText"
}

func (m *BACnetConstructedDataInactiveText) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataInactiveText) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (inactiveText)
	lengthInBits += m.InactiveText.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataInactiveText) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataInactiveTextParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataInactiveText, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInactiveText"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInactiveText")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (inactiveText)
	if pullErr := readBuffer.PullContext("inactiveText"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for inactiveText")
	}
	_inactiveText, _inactiveTextErr := BACnetApplicationTagParse(readBuffer)
	if _inactiveTextErr != nil {
		return nil, errors.Wrap(_inactiveTextErr, "Error parsing 'inactiveText' field")
	}
	inactiveText := CastBACnetApplicationTagCharacterString(_inactiveText)
	if closeErr := readBuffer.CloseContext("inactiveText"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for inactiveText")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInactiveText"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInactiveText")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataInactiveText{
		InactiveText:          CastBACnetApplicationTagCharacterString(inactiveText),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataInactiveText) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInactiveText"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInactiveText")
		}

		// Simple Field (inactiveText)
		if pushErr := writeBuffer.PushContext("inactiveText"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for inactiveText")
		}
		_inactiveTextErr := writeBuffer.WriteSerializable(m.InactiveText)
		if popErr := writeBuffer.PopContext("inactiveText"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for inactiveText")
		}
		if _inactiveTextErr != nil {
			return errors.Wrap(_inactiveTextErr, "Error serializing 'inactiveText' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInactiveText"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInactiveText")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataInactiveText) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
