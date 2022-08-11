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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataAckedTransitions is the corresponding interface of BACnetConstructedDataAckedTransitions
type BACnetConstructedDataAckedTransitions interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAckedTransitions returns AckedTransitions (property field)
	GetAckedTransitions() BACnetEventTransitionBitsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEventTransitionBitsTagged
}

// BACnetConstructedDataAckedTransitionsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAckedTransitions.
// This is useful for switch cases.
type BACnetConstructedDataAckedTransitionsExactly interface {
	BACnetConstructedDataAckedTransitions
	isBACnetConstructedDataAckedTransitions() bool
}

// _BACnetConstructedDataAckedTransitions is the data-structure of this message
type _BACnetConstructedDataAckedTransitions struct {
	*_BACnetConstructedData
	AckedTransitions BACnetEventTransitionBitsTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAckedTransitions) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAckedTransitions) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACKED_TRANSITIONS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAckedTransitions) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAckedTransitions) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAckedTransitions) GetAckedTransitions() BACnetEventTransitionBitsTagged {
	return m.AckedTransitions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAckedTransitions) GetActualValue() BACnetEventTransitionBitsTagged {
	return CastBACnetEventTransitionBitsTagged(m.GetAckedTransitions())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAckedTransitions factory function for _BACnetConstructedDataAckedTransitions
func NewBACnetConstructedDataAckedTransitions(ackedTransitions BACnetEventTransitionBitsTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAckedTransitions {
	_result := &_BACnetConstructedDataAckedTransitions{
		AckedTransitions:       ackedTransitions,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAckedTransitions(structType interface{}) BACnetConstructedDataAckedTransitions {
	if casted, ok := structType.(BACnetConstructedDataAckedTransitions); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAckedTransitions); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAckedTransitions) GetTypeName() string {
	return "BACnetConstructedDataAckedTransitions"
}

func (m *_BACnetConstructedDataAckedTransitions) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAckedTransitions) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ackedTransitions)
	lengthInBits += m.AckedTransitions.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAckedTransitions) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAckedTransitionsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAckedTransitions, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAckedTransitions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAckedTransitions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ackedTransitions)
	if pullErr := readBuffer.PullContext("ackedTransitions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ackedTransitions")
	}
	_ackedTransitions, _ackedTransitionsErr := BACnetEventTransitionBitsTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _ackedTransitionsErr != nil {
		return nil, errors.Wrap(_ackedTransitionsErr, "Error parsing 'ackedTransitions' field of BACnetConstructedDataAckedTransitions")
	}
	ackedTransitions := _ackedTransitions.(BACnetEventTransitionBitsTagged)
	if closeErr := readBuffer.CloseContext("ackedTransitions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ackedTransitions")
	}

	// Virtual field
	_actualValue := ackedTransitions
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAckedTransitions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAckedTransitions")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAckedTransitions{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AckedTransitions: ackedTransitions,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAckedTransitions) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAckedTransitions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAckedTransitions")
		}

		// Simple Field (ackedTransitions)
		if pushErr := writeBuffer.PushContext("ackedTransitions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ackedTransitions")
		}
		_ackedTransitionsErr := writeBuffer.WriteSerializable(m.GetAckedTransitions())
		if popErr := writeBuffer.PopContext("ackedTransitions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ackedTransitions")
		}
		if _ackedTransitionsErr != nil {
			return errors.Wrap(_ackedTransitionsErr, "Error serializing 'ackedTransitions' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAckedTransitions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAckedTransitions")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAckedTransitions) isBACnetConstructedDataAckedTransitions() bool {
	return true
}

func (m *_BACnetConstructedDataAckedTransitions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
