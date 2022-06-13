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

// BACnetConstructedDataMaximumOutput is the data-structure of this message
type BACnetConstructedDataMaximumOutput struct {
	*BACnetConstructedData
	MaximumOutput *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMaximumOutput is the corresponding interface of BACnetConstructedDataMaximumOutput
type IBACnetConstructedDataMaximumOutput interface {
	IBACnetConstructedData
	// GetMaximumOutput returns MaximumOutput (property field)
	GetMaximumOutput() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataMaximumOutput) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMaximumOutput) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAXIMUM_OUTPUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMaximumOutput) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMaximumOutput) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMaximumOutput) GetMaximumOutput() *BACnetApplicationTagReal {
	return m.MaximumOutput
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaximumOutput factory function for BACnetConstructedDataMaximumOutput
func NewBACnetConstructedDataMaximumOutput(maximumOutput *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMaximumOutput {
	_result := &BACnetConstructedDataMaximumOutput{
		MaximumOutput:         maximumOutput,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMaximumOutput(structType interface{}) *BACnetConstructedDataMaximumOutput {
	if casted, ok := structType.(BACnetConstructedDataMaximumOutput); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaximumOutput); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaximumOutput(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaximumOutput(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMaximumOutput) GetTypeName() string {
	return "BACnetConstructedDataMaximumOutput"
}

func (m *BACnetConstructedDataMaximumOutput) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMaximumOutput) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maximumOutput)
	lengthInBits += m.MaximumOutput.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataMaximumOutput) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaximumOutputParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMaximumOutput, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaximumOutput"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaximumOutput")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maximumOutput)
	if pullErr := readBuffer.PullContext("maximumOutput"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maximumOutput")
	}
	_maximumOutput, _maximumOutputErr := BACnetApplicationTagParse(readBuffer)
	if _maximumOutputErr != nil {
		return nil, errors.Wrap(_maximumOutputErr, "Error parsing 'maximumOutput' field")
	}
	maximumOutput := CastBACnetApplicationTagReal(_maximumOutput)
	if closeErr := readBuffer.CloseContext("maximumOutput"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maximumOutput")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaximumOutput"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaximumOutput")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMaximumOutput{
		MaximumOutput:         CastBACnetApplicationTagReal(maximumOutput),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMaximumOutput) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaximumOutput"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaximumOutput")
		}

		// Simple Field (maximumOutput)
		if pushErr := writeBuffer.PushContext("maximumOutput"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maximumOutput")
		}
		_maximumOutputErr := writeBuffer.WriteSerializable(m.MaximumOutput)
		if popErr := writeBuffer.PopContext("maximumOutput"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maximumOutput")
		}
		if _maximumOutputErr != nil {
			return errors.Wrap(_maximumOutputErr, "Error serializing 'maximumOutput' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaximumOutput"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaximumOutput")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMaximumOutput) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
