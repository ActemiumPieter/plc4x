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

// BACnetConstructedDataEventEnable is the data-structure of this message
type BACnetConstructedDataEventEnable struct {
	*BACnetConstructedData
	EventEnable *BACnetEventTransitionBitsTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataEventEnable is the corresponding interface of BACnetConstructedDataEventEnable
type IBACnetConstructedDataEventEnable interface {
	IBACnetConstructedData
	// GetEventEnable returns EventEnable (property field)
	GetEventEnable() *BACnetEventTransitionBitsTagged
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

func (m *BACnetConstructedDataEventEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataEventEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataEventEnable) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataEventEnable) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataEventEnable) GetEventEnable() *BACnetEventTransitionBitsTagged {
	return m.EventEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventEnable factory function for BACnetConstructedDataEventEnable
func NewBACnetConstructedDataEventEnable(eventEnable *BACnetEventTransitionBitsTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataEventEnable {
	_result := &BACnetConstructedDataEventEnable{
		EventEnable:           eventEnable,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataEventEnable(structType interface{}) *BACnetConstructedDataEventEnable {
	if casted, ok := structType.(BACnetConstructedDataEventEnable); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventEnable); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataEventEnable(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataEventEnable(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataEventEnable) GetTypeName() string {
	return "BACnetConstructedDataEventEnable"
}

func (m *BACnetConstructedDataEventEnable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataEventEnable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (eventEnable)
	lengthInBits += m.EventEnable.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataEventEnable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventEnableParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataEventEnable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (eventEnable)
	if pullErr := readBuffer.PullContext("eventEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventEnable")
	}
	_eventEnable, _eventEnableErr := BACnetEventTransitionBitsTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _eventEnableErr != nil {
		return nil, errors.Wrap(_eventEnableErr, "Error parsing 'eventEnable' field")
	}
	eventEnable := CastBACnetEventTransitionBitsTagged(_eventEnable)
	if closeErr := readBuffer.CloseContext("eventEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventEnable")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventEnable")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataEventEnable{
		EventEnable:           CastBACnetEventTransitionBitsTagged(eventEnable),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataEventEnable) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventEnable")
		}

		// Simple Field (eventEnable)
		if pushErr := writeBuffer.PushContext("eventEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventEnable")
		}
		_eventEnableErr := writeBuffer.WriteSerializable(m.EventEnable)
		if popErr := writeBuffer.PopContext("eventEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventEnable")
		}
		if _eventEnableErr != nil {
			return errors.Wrap(_eventEnableErr, "Error serializing 'eventEnable' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventEnable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataEventEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
