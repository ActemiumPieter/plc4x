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

// BACnetActionList is the data-structure of this message
type BACnetActionList struct {
	InnerOpeningTag *BACnetOpeningTag
	Action          []*BACnetActionCommand
	InnerClosingTag *BACnetClosingTag
}

// IBACnetActionList is the corresponding interface of BACnetActionList
type IBACnetActionList interface {
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() *BACnetOpeningTag
	// GetAction returns Action (property field)
	GetAction() []*BACnetActionCommand
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
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetActionList) GetInnerOpeningTag() *BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *BACnetActionList) GetAction() []*BACnetActionCommand {
	return m.Action
}

func (m *BACnetActionList) GetInnerClosingTag() *BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetActionList factory function for BACnetActionList
func NewBACnetActionList(innerOpeningTag *BACnetOpeningTag, action []*BACnetActionCommand, innerClosingTag *BACnetClosingTag) *BACnetActionList {
	return &BACnetActionList{InnerOpeningTag: innerOpeningTag, Action: action, InnerClosingTag: innerClosingTag}
}

func CastBACnetActionList(structType interface{}) *BACnetActionList {
	if casted, ok := structType.(BACnetActionList); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetActionList); ok {
		return casted
	}
	return nil
}

func (m *BACnetActionList) GetTypeName() string {
	return "BACnetActionList"
}

func (m *BACnetActionList) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetActionList) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Array field
	if len(m.Action) > 0 {
		for _, element := range m.Action {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetActionList) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetActionListParse(readBuffer utils.ReadBuffer) (*BACnetActionList, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetActionList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetActionList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(0)))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field")
	}
	innerOpeningTag := CastBACnetOpeningTag(_innerOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Array field (action)
	if pullErr := readBuffer.PullContext("action", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for action")
	}
	// Terminated array
	action := make([]*BACnetActionCommand, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, 0)) {
			_item, _err := BACnetActionCommandParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'action' field")
			}
			action = append(action, CastBACnetActionCommand(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("action", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for action")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(0)))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field")
	}
	innerClosingTag := CastBACnetClosingTag(_innerClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetActionList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetActionList")
	}

	// Create the instance
	return NewBACnetActionList(innerOpeningTag, action, innerClosingTag), nil
}

func (m *BACnetActionList) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetActionList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetActionList")
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

	// Array Field (action)
	if m.Action != nil {
		if pushErr := writeBuffer.PushContext("action", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for action")
		}
		for _, _element := range m.Action {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'action' field")
			}
		}
		if popErr := writeBuffer.PopContext("action", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for action")
		}
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

	if popErr := writeBuffer.PopContext("BACnetActionList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetActionList")
	}
	return nil
}

func (m *BACnetActionList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
