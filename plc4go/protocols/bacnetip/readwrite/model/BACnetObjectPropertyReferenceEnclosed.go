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

// BACnetObjectPropertyReferenceEnclosed is the data-structure of this message
type BACnetObjectPropertyReferenceEnclosed struct {
	OpeningTag              *BACnetOpeningTag
	ObjectPropertyReference *BACnetObjectPropertyReference
	ClosingTag              *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetObjectPropertyReferenceEnclosed is the corresponding interface of BACnetObjectPropertyReferenceEnclosed
type IBACnetObjectPropertyReferenceEnclosed interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetObjectPropertyReference returns ObjectPropertyReference (property field)
	GetObjectPropertyReference() *BACnetObjectPropertyReference
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetObjectPropertyReferenceEnclosed) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetObjectPropertyReferenceEnclosed) GetObjectPropertyReference() *BACnetObjectPropertyReference {
	return m.ObjectPropertyReference
}

func (m *BACnetObjectPropertyReferenceEnclosed) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetObjectPropertyReferenceEnclosed factory function for BACnetObjectPropertyReferenceEnclosed
func NewBACnetObjectPropertyReferenceEnclosed(openingTag *BACnetOpeningTag, objectPropertyReference *BACnetObjectPropertyReference, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetObjectPropertyReferenceEnclosed {
	return &BACnetObjectPropertyReferenceEnclosed{OpeningTag: openingTag, ObjectPropertyReference: objectPropertyReference, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetObjectPropertyReferenceEnclosed(structType interface{}) *BACnetObjectPropertyReferenceEnclosed {
	if casted, ok := structType.(BACnetObjectPropertyReferenceEnclosed); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetObjectPropertyReferenceEnclosed); ok {
		return casted
	}
	return nil
}

func (m *BACnetObjectPropertyReferenceEnclosed) GetTypeName() string {
	return "BACnetObjectPropertyReferenceEnclosed"
}

func (m *BACnetObjectPropertyReferenceEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetObjectPropertyReferenceEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (objectPropertyReference)
	lengthInBits += m.ObjectPropertyReference.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetObjectPropertyReferenceEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetObjectPropertyReferenceEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetObjectPropertyReferenceEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetObjectPropertyReferenceEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetObjectPropertyReferenceEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (objectPropertyReference)
	if pullErr := readBuffer.PullContext("objectPropertyReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectPropertyReference")
	}
	_objectPropertyReference, _objectPropertyReferenceErr := BACnetObjectPropertyReferenceParse(readBuffer)
	if _objectPropertyReferenceErr != nil {
		return nil, errors.Wrap(_objectPropertyReferenceErr, "Error parsing 'objectPropertyReference' field")
	}
	objectPropertyReference := CastBACnetObjectPropertyReference(_objectPropertyReference)
	if closeErr := readBuffer.CloseContext("objectPropertyReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectPropertyReference")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetObjectPropertyReferenceEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetObjectPropertyReferenceEnclosed")
	}

	// Create the instance
	return NewBACnetObjectPropertyReferenceEnclosed(openingTag, objectPropertyReference, closingTag, tagNumber), nil
}

func (m *BACnetObjectPropertyReferenceEnclosed) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetObjectPropertyReferenceEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetObjectPropertyReferenceEnclosed")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.OpeningTag)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (objectPropertyReference)
	if pushErr := writeBuffer.PushContext("objectPropertyReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for objectPropertyReference")
	}
	_objectPropertyReferenceErr := writeBuffer.WriteSerializable(m.ObjectPropertyReference)
	if popErr := writeBuffer.PopContext("objectPropertyReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for objectPropertyReference")
	}
	if _objectPropertyReferenceErr != nil {
		return errors.Wrap(_objectPropertyReferenceErr, "Error serializing 'objectPropertyReference' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.ClosingTag)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetObjectPropertyReferenceEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetObjectPropertyReferenceEnclosed")
	}
	return nil
}

func (m *BACnetObjectPropertyReferenceEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
