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

// BACnetConstructedDataAllowGroupDelayInhibit is the data-structure of this message
type BACnetConstructedDataAllowGroupDelayInhibit struct {
	*BACnetConstructedData
	AllowGroupDelayInhibit *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataAllowGroupDelayInhibit is the corresponding interface of BACnetConstructedDataAllowGroupDelayInhibit
type IBACnetConstructedDataAllowGroupDelayInhibit interface {
	IBACnetConstructedData
	// GetAllowGroupDelayInhibit returns AllowGroupDelayInhibit (property field)
	GetAllowGroupDelayInhibit() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataAllowGroupDelayInhibit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAllowGroupDelayInhibit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALLOW_GROUP_DELAY_INHIBIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAllowGroupDelayInhibit) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAllowGroupDelayInhibit) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAllowGroupDelayInhibit) GetAllowGroupDelayInhibit() *BACnetApplicationTagBoolean {
	return m.AllowGroupDelayInhibit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAllowGroupDelayInhibit factory function for BACnetConstructedDataAllowGroupDelayInhibit
func NewBACnetConstructedDataAllowGroupDelayInhibit(allowGroupDelayInhibit *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataAllowGroupDelayInhibit {
	_result := &BACnetConstructedDataAllowGroupDelayInhibit{
		AllowGroupDelayInhibit: allowGroupDelayInhibit,
		BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAllowGroupDelayInhibit(structType interface{}) *BACnetConstructedDataAllowGroupDelayInhibit {
	if casted, ok := structType.(BACnetConstructedDataAllowGroupDelayInhibit); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAllowGroupDelayInhibit); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAllowGroupDelayInhibit(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAllowGroupDelayInhibit(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAllowGroupDelayInhibit) GetTypeName() string {
	return "BACnetConstructedDataAllowGroupDelayInhibit"
}

func (m *BACnetConstructedDataAllowGroupDelayInhibit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAllowGroupDelayInhibit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (allowGroupDelayInhibit)
	lengthInBits += m.AllowGroupDelayInhibit.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataAllowGroupDelayInhibit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAllowGroupDelayInhibitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataAllowGroupDelayInhibit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAllowGroupDelayInhibit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAllowGroupDelayInhibit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (allowGroupDelayInhibit)
	if pullErr := readBuffer.PullContext("allowGroupDelayInhibit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for allowGroupDelayInhibit")
	}
	_allowGroupDelayInhibit, _allowGroupDelayInhibitErr := BACnetApplicationTagParse(readBuffer)
	if _allowGroupDelayInhibitErr != nil {
		return nil, errors.Wrap(_allowGroupDelayInhibitErr, "Error parsing 'allowGroupDelayInhibit' field")
	}
	allowGroupDelayInhibit := CastBACnetApplicationTagBoolean(_allowGroupDelayInhibit)
	if closeErr := readBuffer.CloseContext("allowGroupDelayInhibit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for allowGroupDelayInhibit")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAllowGroupDelayInhibit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAllowGroupDelayInhibit")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAllowGroupDelayInhibit{
		AllowGroupDelayInhibit: CastBACnetApplicationTagBoolean(allowGroupDelayInhibit),
		BACnetConstructedData:  &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAllowGroupDelayInhibit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAllowGroupDelayInhibit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAllowGroupDelayInhibit")
		}

		// Simple Field (allowGroupDelayInhibit)
		if pushErr := writeBuffer.PushContext("allowGroupDelayInhibit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for allowGroupDelayInhibit")
		}
		_allowGroupDelayInhibitErr := writeBuffer.WriteSerializable(m.AllowGroupDelayInhibit)
		if popErr := writeBuffer.PopContext("allowGroupDelayInhibit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for allowGroupDelayInhibit")
		}
		if _allowGroupDelayInhibitErr != nil {
			return errors.Wrap(_allowGroupDelayInhibitErr, "Error serializing 'allowGroupDelayInhibit' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAllowGroupDelayInhibit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAllowGroupDelayInhibit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAllowGroupDelayInhibit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
