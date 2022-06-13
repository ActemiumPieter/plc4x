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

// BACnetHostNPortEnclosed is the data-structure of this message
type BACnetHostNPortEnclosed struct {
	OpeningTag      *BACnetOpeningTag
	BacnetHostNPort *BACnetHostNPort
	ClosingTag      *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetHostNPortEnclosed is the corresponding interface of BACnetHostNPortEnclosed
type IBACnetHostNPortEnclosed interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetBacnetHostNPort returns BacnetHostNPort (property field)
	GetBacnetHostNPort() *BACnetHostNPort
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

func (m *BACnetHostNPortEnclosed) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetHostNPortEnclosed) GetBacnetHostNPort() *BACnetHostNPort {
	return m.BacnetHostNPort
}

func (m *BACnetHostNPortEnclosed) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetHostNPortEnclosed factory function for BACnetHostNPortEnclosed
func NewBACnetHostNPortEnclosed(openingTag *BACnetOpeningTag, bacnetHostNPort *BACnetHostNPort, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetHostNPortEnclosed {
	return &BACnetHostNPortEnclosed{OpeningTag: openingTag, BacnetHostNPort: bacnetHostNPort, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetHostNPortEnclosed(structType interface{}) *BACnetHostNPortEnclosed {
	if casted, ok := structType.(BACnetHostNPortEnclosed); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetHostNPortEnclosed); ok {
		return casted
	}
	return nil
}

func (m *BACnetHostNPortEnclosed) GetTypeName() string {
	return "BACnetHostNPortEnclosed"
}

func (m *BACnetHostNPortEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetHostNPortEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (bacnetHostNPort)
	lengthInBits += m.BacnetHostNPort.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetHostNPortEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetHostNPortEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetHostNPortEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostNPortEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostNPortEnclosed")
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

	// Simple Field (bacnetHostNPort)
	if pullErr := readBuffer.PullContext("bacnetHostNPort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bacnetHostNPort")
	}
	_bacnetHostNPort, _bacnetHostNPortErr := BACnetHostNPortParse(readBuffer)
	if _bacnetHostNPortErr != nil {
		return nil, errors.Wrap(_bacnetHostNPortErr, "Error parsing 'bacnetHostNPort' field")
	}
	bacnetHostNPort := CastBACnetHostNPort(_bacnetHostNPort)
	if closeErr := readBuffer.CloseContext("bacnetHostNPort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bacnetHostNPort")
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

	if closeErr := readBuffer.CloseContext("BACnetHostNPortEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostNPortEnclosed")
	}

	// Create the instance
	return NewBACnetHostNPortEnclosed(openingTag, bacnetHostNPort, closingTag, tagNumber), nil
}

func (m *BACnetHostNPortEnclosed) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetHostNPortEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetHostNPortEnclosed")
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

	// Simple Field (bacnetHostNPort)
	if pushErr := writeBuffer.PushContext("bacnetHostNPort"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for bacnetHostNPort")
	}
	_bacnetHostNPortErr := writeBuffer.WriteSerializable(m.BacnetHostNPort)
	if popErr := writeBuffer.PopContext("bacnetHostNPort"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for bacnetHostNPort")
	}
	if _bacnetHostNPortErr != nil {
		return errors.Wrap(_bacnetHostNPortErr, "Error serializing 'bacnetHostNPort' field")
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

	if popErr := writeBuffer.PopContext("BACnetHostNPortEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetHostNPortEnclosed")
	}
	return nil
}

func (m *BACnetHostNPortEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
