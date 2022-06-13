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

// BACnetNotificationParametersBufferReady is the data-structure of this message
type BACnetNotificationParametersBufferReady struct {
	*BACnetNotificationParameters
	InnerOpeningTag      *BACnetOpeningTag
	BufferProperty       *BACnetDeviceObjectPropertyReferenceEnclosed
	PreviousNotification *BACnetContextTagUnsignedInteger
	CurrentNotification  *BACnetContextTagUnsignedInteger
	InnerClosingTag      *BACnetClosingTag

	// Arguments.
	TagNumber          uint8
	ObjectTypeArgument BACnetObjectType
}

// IBACnetNotificationParametersBufferReady is the corresponding interface of BACnetNotificationParametersBufferReady
type IBACnetNotificationParametersBufferReady interface {
	IBACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() *BACnetOpeningTag
	// GetBufferProperty returns BufferProperty (property field)
	GetBufferProperty() *BACnetDeviceObjectPropertyReferenceEnclosed
	// GetPreviousNotification returns PreviousNotification (property field)
	GetPreviousNotification() *BACnetContextTagUnsignedInteger
	// GetCurrentNotification returns CurrentNotification (property field)
	GetCurrentNotification() *BACnetContextTagUnsignedInteger
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() *BACnetClosingTag
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

func (m *BACnetNotificationParametersBufferReady) InitializeParent(parent *BACnetNotificationParameters, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParameters.OpeningTag = openingTag
	m.BACnetNotificationParameters.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParameters.ClosingTag = closingTag
}

func (m *BACnetNotificationParametersBufferReady) GetParent() *BACnetNotificationParameters {
	return m.BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetNotificationParametersBufferReady) GetInnerOpeningTag() *BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *BACnetNotificationParametersBufferReady) GetBufferProperty() *BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.BufferProperty
}

func (m *BACnetNotificationParametersBufferReady) GetPreviousNotification() *BACnetContextTagUnsignedInteger {
	return m.PreviousNotification
}

func (m *BACnetNotificationParametersBufferReady) GetCurrentNotification() *BACnetContextTagUnsignedInteger {
	return m.CurrentNotification
}

func (m *BACnetNotificationParametersBufferReady) GetInnerClosingTag() *BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersBufferReady factory function for BACnetNotificationParametersBufferReady
func NewBACnetNotificationParametersBufferReady(innerOpeningTag *BACnetOpeningTag, bufferProperty *BACnetDeviceObjectPropertyReferenceEnclosed, previousNotification *BACnetContextTagUnsignedInteger, currentNotification *BACnetContextTagUnsignedInteger, innerClosingTag *BACnetClosingTag, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *BACnetNotificationParametersBufferReady {
	_result := &BACnetNotificationParametersBufferReady{
		InnerOpeningTag:              innerOpeningTag,
		BufferProperty:               bufferProperty,
		PreviousNotification:         previousNotification,
		CurrentNotification:          currentNotification,
		InnerClosingTag:              innerClosingTag,
		BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetNotificationParametersBufferReady(structType interface{}) *BACnetNotificationParametersBufferReady {
	if casted, ok := structType.(BACnetNotificationParametersBufferReady); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersBufferReady); ok {
		return casted
	}
	if casted, ok := structType.(BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersBufferReady(casted.Child)
	}
	if casted, ok := structType.(*BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersBufferReady(casted.Child)
	}
	return nil
}

func (m *BACnetNotificationParametersBufferReady) GetTypeName() string {
	return "BACnetNotificationParametersBufferReady"
}

func (m *BACnetNotificationParametersBufferReady) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersBufferReady) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (bufferProperty)
	lengthInBits += m.BufferProperty.GetLengthInBits()

	// Simple field (previousNotification)
	lengthInBits += m.PreviousNotification.GetLengthInBits()

	// Simple field (currentNotification)
	lengthInBits += m.CurrentNotification.GetLengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersBufferReady) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersBufferReadyParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, peekedTagNumber uint8) (*BACnetNotificationParametersBufferReady, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersBufferReady"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersBufferReady")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field")
	}
	innerOpeningTag := CastBACnetOpeningTag(_innerOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Simple Field (bufferProperty)
	if pullErr := readBuffer.PullContext("bufferProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bufferProperty")
	}
	_bufferProperty, _bufferPropertyErr := BACnetDeviceObjectPropertyReferenceEnclosedParse(readBuffer, uint8(uint8(0)))
	if _bufferPropertyErr != nil {
		return nil, errors.Wrap(_bufferPropertyErr, "Error parsing 'bufferProperty' field")
	}
	bufferProperty := CastBACnetDeviceObjectPropertyReferenceEnclosed(_bufferProperty)
	if closeErr := readBuffer.CloseContext("bufferProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bufferProperty")
	}

	// Simple Field (previousNotification)
	if pullErr := readBuffer.PullContext("previousNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for previousNotification")
	}
	_previousNotification, _previousNotificationErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _previousNotificationErr != nil {
		return nil, errors.Wrap(_previousNotificationErr, "Error parsing 'previousNotification' field")
	}
	previousNotification := CastBACnetContextTagUnsignedInteger(_previousNotification)
	if closeErr := readBuffer.CloseContext("previousNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for previousNotification")
	}

	// Simple Field (currentNotification)
	if pullErr := readBuffer.PullContext("currentNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for currentNotification")
	}
	_currentNotification, _currentNotificationErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _currentNotificationErr != nil {
		return nil, errors.Wrap(_currentNotificationErr, "Error parsing 'currentNotification' field")
	}
	currentNotification := CastBACnetContextTagUnsignedInteger(_currentNotification)
	if closeErr := readBuffer.CloseContext("currentNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for currentNotification")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field")
	}
	innerClosingTag := CastBACnetClosingTag(_innerClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersBufferReady"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersBufferReady")
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersBufferReady{
		InnerOpeningTag:              CastBACnetOpeningTag(innerOpeningTag),
		BufferProperty:               CastBACnetDeviceObjectPropertyReferenceEnclosed(bufferProperty),
		PreviousNotification:         CastBACnetContextTagUnsignedInteger(previousNotification),
		CurrentNotification:          CastBACnetContextTagUnsignedInteger(currentNotification),
		InnerClosingTag:              CastBACnetClosingTag(innerClosingTag),
		BACnetNotificationParameters: &BACnetNotificationParameters{},
	}
	_child.BACnetNotificationParameters.Child = _child
	return _child, nil
}

func (m *BACnetNotificationParametersBufferReady) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersBufferReady"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersBufferReady")
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerOpeningTag")
		}
		_innerOpeningTagErr := writeBuffer.WriteSerializable(m.InnerOpeningTag)
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerOpeningTag")
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (bufferProperty)
		if pushErr := writeBuffer.PushContext("bufferProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bufferProperty")
		}
		_bufferPropertyErr := writeBuffer.WriteSerializable(m.BufferProperty)
		if popErr := writeBuffer.PopContext("bufferProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bufferProperty")
		}
		if _bufferPropertyErr != nil {
			return errors.Wrap(_bufferPropertyErr, "Error serializing 'bufferProperty' field")
		}

		// Simple Field (previousNotification)
		if pushErr := writeBuffer.PushContext("previousNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for previousNotification")
		}
		_previousNotificationErr := writeBuffer.WriteSerializable(m.PreviousNotification)
		if popErr := writeBuffer.PopContext("previousNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for previousNotification")
		}
		if _previousNotificationErr != nil {
			return errors.Wrap(_previousNotificationErr, "Error serializing 'previousNotification' field")
		}

		// Simple Field (currentNotification)
		if pushErr := writeBuffer.PushContext("currentNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for currentNotification")
		}
		_currentNotificationErr := writeBuffer.WriteSerializable(m.CurrentNotification)
		if popErr := writeBuffer.PopContext("currentNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for currentNotification")
		}
		if _currentNotificationErr != nil {
			return errors.Wrap(_currentNotificationErr, "Error serializing 'currentNotification' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerClosingTag")
		}
		_innerClosingTagErr := writeBuffer.WriteSerializable(m.InnerClosingTag)
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerClosingTag")
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersBufferReady"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersBufferReady")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersBufferReady) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
