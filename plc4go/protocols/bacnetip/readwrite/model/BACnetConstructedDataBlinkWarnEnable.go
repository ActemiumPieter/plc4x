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

// BACnetConstructedDataBlinkWarnEnable is the data-structure of this message
type BACnetConstructedDataBlinkWarnEnable struct {
	*BACnetConstructedData
	BlinkWarnEnable *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataBlinkWarnEnable is the corresponding interface of BACnetConstructedDataBlinkWarnEnable
type IBACnetConstructedDataBlinkWarnEnable interface {
	IBACnetConstructedData
	// GetBlinkWarnEnable returns BlinkWarnEnable (property field)
	GetBlinkWarnEnable() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataBlinkWarnEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataBlinkWarnEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BLINK_WARN_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBlinkWarnEnable) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBlinkWarnEnable) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBlinkWarnEnable) GetBlinkWarnEnable() *BACnetApplicationTagBoolean {
	return m.BlinkWarnEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBlinkWarnEnable factory function for BACnetConstructedDataBlinkWarnEnable
func NewBACnetConstructedDataBlinkWarnEnable(blinkWarnEnable *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataBlinkWarnEnable {
	_result := &BACnetConstructedDataBlinkWarnEnable{
		BlinkWarnEnable:       blinkWarnEnable,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBlinkWarnEnable(structType interface{}) *BACnetConstructedDataBlinkWarnEnable {
	if casted, ok := structType.(BACnetConstructedDataBlinkWarnEnable); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBlinkWarnEnable); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBlinkWarnEnable(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBlinkWarnEnable(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBlinkWarnEnable) GetTypeName() string {
	return "BACnetConstructedDataBlinkWarnEnable"
}

func (m *BACnetConstructedDataBlinkWarnEnable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBlinkWarnEnable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (blinkWarnEnable)
	lengthInBits += m.BlinkWarnEnable.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataBlinkWarnEnable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBlinkWarnEnableParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataBlinkWarnEnable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBlinkWarnEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBlinkWarnEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (blinkWarnEnable)
	if pullErr := readBuffer.PullContext("blinkWarnEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for blinkWarnEnable")
	}
	_blinkWarnEnable, _blinkWarnEnableErr := BACnetApplicationTagParse(readBuffer)
	if _blinkWarnEnableErr != nil {
		return nil, errors.Wrap(_blinkWarnEnableErr, "Error parsing 'blinkWarnEnable' field")
	}
	blinkWarnEnable := CastBACnetApplicationTagBoolean(_blinkWarnEnable)
	if closeErr := readBuffer.CloseContext("blinkWarnEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for blinkWarnEnable")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBlinkWarnEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBlinkWarnEnable")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBlinkWarnEnable{
		BlinkWarnEnable:       CastBACnetApplicationTagBoolean(blinkWarnEnable),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBlinkWarnEnable) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBlinkWarnEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBlinkWarnEnable")
		}

		// Simple Field (blinkWarnEnable)
		if pushErr := writeBuffer.PushContext("blinkWarnEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for blinkWarnEnable")
		}
		_blinkWarnEnableErr := writeBuffer.WriteSerializable(m.BlinkWarnEnable)
		if popErr := writeBuffer.PopContext("blinkWarnEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for blinkWarnEnable")
		}
		if _blinkWarnEnableErr != nil {
			return errors.Wrap(_blinkWarnEnableErr, "Error serializing 'blinkWarnEnable' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBlinkWarnEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBlinkWarnEnable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBlinkWarnEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
