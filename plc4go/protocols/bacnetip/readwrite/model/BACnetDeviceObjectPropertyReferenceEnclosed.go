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

// BACnetDeviceObjectPropertyReferenceEnclosed is the corresponding interface of BACnetDeviceObjectPropertyReferenceEnclosed
type BACnetDeviceObjectPropertyReferenceEnclosed interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetValue returns Value (property field)
	GetValue() BACnetDeviceObjectPropertyReference
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetDeviceObjectPropertyReferenceEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetDeviceObjectPropertyReferenceEnclosed.
// This is useful for switch cases.
type BACnetDeviceObjectPropertyReferenceEnclosedExactly interface {
	BACnetDeviceObjectPropertyReferenceEnclosed
	isBACnetDeviceObjectPropertyReferenceEnclosed() bool
}

// _BACnetDeviceObjectPropertyReferenceEnclosed is the data-structure of this message
type _BACnetDeviceObjectPropertyReferenceEnclosed struct {
	OpeningTag BACnetOpeningTag
	Value      BACnetDeviceObjectPropertyReference
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDeviceObjectPropertyReferenceEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetDeviceObjectPropertyReferenceEnclosed) GetValue() BACnetDeviceObjectPropertyReference {
	return m.Value
}

func (m *_BACnetDeviceObjectPropertyReferenceEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDeviceObjectPropertyReferenceEnclosed factory function for _BACnetDeviceObjectPropertyReferenceEnclosed
func NewBACnetDeviceObjectPropertyReferenceEnclosed(openingTag BACnetOpeningTag, value BACnetDeviceObjectPropertyReference, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetDeviceObjectPropertyReferenceEnclosed {
	return &_BACnetDeviceObjectPropertyReferenceEnclosed{OpeningTag: openingTag, Value: value, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetDeviceObjectPropertyReferenceEnclosed(structType interface{}) BACnetDeviceObjectPropertyReferenceEnclosed {
	if casted, ok := structType.(BACnetDeviceObjectPropertyReferenceEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDeviceObjectPropertyReferenceEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDeviceObjectPropertyReferenceEnclosed) GetTypeName() string {
	return "BACnetDeviceObjectPropertyReferenceEnclosed"
}

func (m *_BACnetDeviceObjectPropertyReferenceEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetDeviceObjectPropertyReferenceEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetDeviceObjectPropertyReferenceEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetDeviceObjectPropertyReferenceEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetDeviceObjectPropertyReferenceEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDeviceObjectPropertyReferenceEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDeviceObjectPropertyReferenceEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetDeviceObjectPropertyReferenceEnclosed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	_value, _valueErr := BACnetDeviceObjectPropertyReferenceParse(readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetDeviceObjectPropertyReferenceEnclosed")
	}
	value := _value.(BACnetDeviceObjectPropertyReference)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetDeviceObjectPropertyReferenceEnclosed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetDeviceObjectPropertyReferenceEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDeviceObjectPropertyReferenceEnclosed")
	}

	// Create the instance
	return &_BACnetDeviceObjectPropertyReferenceEnclosed{
		TagNumber:  tagNumber,
		OpeningTag: openingTag,
		Value:      value,
		ClosingTag: closingTag,
	}, nil
}

func (m *_BACnetDeviceObjectPropertyReferenceEnclosed) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetDeviceObjectPropertyReferenceEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDeviceObjectPropertyReferenceEnclosed")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (value)
	if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for value")
	}
	_valueErr := writeBuffer.WriteSerializable(m.GetValue())
	if popErr := writeBuffer.PopContext("value"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for value")
	}
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDeviceObjectPropertyReferenceEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDeviceObjectPropertyReferenceEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetDeviceObjectPropertyReferenceEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetDeviceObjectPropertyReferenceEnclosed) isBACnetDeviceObjectPropertyReferenceEnclosed() bool {
	return true
}

func (m *_BACnetDeviceObjectPropertyReferenceEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
