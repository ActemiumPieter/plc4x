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

// BACnetFaultParameterFaultCharacterString is the data-structure of this message
type BACnetFaultParameterFaultCharacterString struct {
	*BACnetFaultParameter
	OpeningTag        *BACnetOpeningTag
	ListOfFaultValues *BACnetFaultParameterFaultCharacterStringListOfFaultValues
	ClosingTag        *BACnetClosingTag
}

// IBACnetFaultParameterFaultCharacterString is the corresponding interface of BACnetFaultParameterFaultCharacterString
type IBACnetFaultParameterFaultCharacterString interface {
	IBACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetListOfFaultValues returns ListOfFaultValues (property field)
	GetListOfFaultValues() *BACnetFaultParameterFaultCharacterStringListOfFaultValues
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
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetFaultParameterFaultCharacterString) InitializeParent(parent *BACnetFaultParameter, peekedTagHeader *BACnetTagHeader) {
	m.BACnetFaultParameter.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetFaultParameterFaultCharacterString) GetParent() *BACnetFaultParameter {
	return m.BACnetFaultParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetFaultParameterFaultCharacterString) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetFaultParameterFaultCharacterString) GetListOfFaultValues() *BACnetFaultParameterFaultCharacterStringListOfFaultValues {
	return m.ListOfFaultValues
}

func (m *BACnetFaultParameterFaultCharacterString) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultCharacterString factory function for BACnetFaultParameterFaultCharacterString
func NewBACnetFaultParameterFaultCharacterString(openingTag *BACnetOpeningTag, listOfFaultValues *BACnetFaultParameterFaultCharacterStringListOfFaultValues, closingTag *BACnetClosingTag, peekedTagHeader *BACnetTagHeader) *BACnetFaultParameterFaultCharacterString {
	_result := &BACnetFaultParameterFaultCharacterString{
		OpeningTag:           openingTag,
		ListOfFaultValues:    listOfFaultValues,
		ClosingTag:           closingTag,
		BACnetFaultParameter: NewBACnetFaultParameter(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetFaultParameterFaultCharacterString(structType interface{}) *BACnetFaultParameterFaultCharacterString {
	if casted, ok := structType.(BACnetFaultParameterFaultCharacterString); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(BACnetFaultParameter); ok {
		return CastBACnetFaultParameterFaultCharacterString(casted.Child)
	}
	if casted, ok := structType.(*BACnetFaultParameter); ok {
		return CastBACnetFaultParameterFaultCharacterString(casted.Child)
	}
	return nil
}

func (m *BACnetFaultParameterFaultCharacterString) GetTypeName() string {
	return "BACnetFaultParameterFaultCharacterString"
}

func (m *BACnetFaultParameterFaultCharacterString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetFaultParameterFaultCharacterString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (listOfFaultValues)
	lengthInBits += m.ListOfFaultValues.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetFaultParameterFaultCharacterString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultCharacterStringParse(readBuffer utils.ReadBuffer) (*BACnetFaultParameterFaultCharacterString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(1)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (listOfFaultValues)
	if pullErr := readBuffer.PullContext("listOfFaultValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfFaultValues")
	}
	_listOfFaultValues, _listOfFaultValuesErr := BACnetFaultParameterFaultCharacterStringListOfFaultValuesParse(readBuffer, uint8(uint8(0)))
	if _listOfFaultValuesErr != nil {
		return nil, errors.Wrap(_listOfFaultValuesErr, "Error parsing 'listOfFaultValues' field")
	}
	listOfFaultValues := CastBACnetFaultParameterFaultCharacterStringListOfFaultValues(_listOfFaultValues)
	if closeErr := readBuffer.CloseContext("listOfFaultValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfFaultValues")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(1)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultCharacterString")
	}

	// Create a partially initialized instance
	_child := &BACnetFaultParameterFaultCharacterString{
		OpeningTag:           CastBACnetOpeningTag(openingTag),
		ListOfFaultValues:    CastBACnetFaultParameterFaultCharacterStringListOfFaultValues(listOfFaultValues),
		ClosingTag:           CastBACnetClosingTag(closingTag),
		BACnetFaultParameter: &BACnetFaultParameter{},
	}
	_child.BACnetFaultParameter.Child = _child
	return _child, nil
}

func (m *BACnetFaultParameterFaultCharacterString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultCharacterString")
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

		// Simple Field (listOfFaultValues)
		if pushErr := writeBuffer.PushContext("listOfFaultValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfFaultValues")
		}
		_listOfFaultValuesErr := writeBuffer.WriteSerializable(m.ListOfFaultValues)
		if popErr := writeBuffer.PopContext("listOfFaultValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfFaultValues")
		}
		if _listOfFaultValuesErr != nil {
			return errors.Wrap(_listOfFaultValuesErr, "Error serializing 'listOfFaultValues' field")
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

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultCharacterString")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetFaultParameterFaultCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
