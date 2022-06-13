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

// BACnetConstructedDataControlledVariableReference is the data-structure of this message
type BACnetConstructedDataControlledVariableReference struct {
	*BACnetConstructedData
	ControlledVariableReference *BACnetObjectPropertyReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataControlledVariableReference is the corresponding interface of BACnetConstructedDataControlledVariableReference
type IBACnetConstructedDataControlledVariableReference interface {
	IBACnetConstructedData
	// GetControlledVariableReference returns ControlledVariableReference (property field)
	GetControlledVariableReference() *BACnetObjectPropertyReference
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

func (m *BACnetConstructedDataControlledVariableReference) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataControlledVariableReference) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CONTROLLED_VARIABLE_REFERENCE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataControlledVariableReference) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataControlledVariableReference) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataControlledVariableReference) GetControlledVariableReference() *BACnetObjectPropertyReference {
	return m.ControlledVariableReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataControlledVariableReference factory function for BACnetConstructedDataControlledVariableReference
func NewBACnetConstructedDataControlledVariableReference(controlledVariableReference *BACnetObjectPropertyReference, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataControlledVariableReference {
	_result := &BACnetConstructedDataControlledVariableReference{
		ControlledVariableReference: controlledVariableReference,
		BACnetConstructedData:       NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataControlledVariableReference(structType interface{}) *BACnetConstructedDataControlledVariableReference {
	if casted, ok := structType.(BACnetConstructedDataControlledVariableReference); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataControlledVariableReference); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataControlledVariableReference(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataControlledVariableReference(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataControlledVariableReference) GetTypeName() string {
	return "BACnetConstructedDataControlledVariableReference"
}

func (m *BACnetConstructedDataControlledVariableReference) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataControlledVariableReference) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (controlledVariableReference)
	lengthInBits += m.ControlledVariableReference.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataControlledVariableReference) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataControlledVariableReferenceParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataControlledVariableReference, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataControlledVariableReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataControlledVariableReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (controlledVariableReference)
	if pullErr := readBuffer.PullContext("controlledVariableReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for controlledVariableReference")
	}
	_controlledVariableReference, _controlledVariableReferenceErr := BACnetObjectPropertyReferenceParse(readBuffer)
	if _controlledVariableReferenceErr != nil {
		return nil, errors.Wrap(_controlledVariableReferenceErr, "Error parsing 'controlledVariableReference' field")
	}
	controlledVariableReference := CastBACnetObjectPropertyReference(_controlledVariableReference)
	if closeErr := readBuffer.CloseContext("controlledVariableReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for controlledVariableReference")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataControlledVariableReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataControlledVariableReference")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataControlledVariableReference{
		ControlledVariableReference: CastBACnetObjectPropertyReference(controlledVariableReference),
		BACnetConstructedData:       &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataControlledVariableReference) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataControlledVariableReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataControlledVariableReference")
		}

		// Simple Field (controlledVariableReference)
		if pushErr := writeBuffer.PushContext("controlledVariableReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for controlledVariableReference")
		}
		_controlledVariableReferenceErr := writeBuffer.WriteSerializable(m.ControlledVariableReference)
		if popErr := writeBuffer.PopContext("controlledVariableReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for controlledVariableReference")
		}
		if _controlledVariableReferenceErr != nil {
			return errors.Wrap(_controlledVariableReferenceErr, "Error serializing 'controlledVariableReference' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataControlledVariableReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataControlledVariableReference")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataControlledVariableReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
