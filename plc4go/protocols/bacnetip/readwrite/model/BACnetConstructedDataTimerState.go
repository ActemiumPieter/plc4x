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

// BACnetConstructedDataTimerState is the corresponding interface of BACnetConstructedDataTimerState
type BACnetConstructedDataTimerState interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTimerState returns TimerState (property field)
	GetTimerState() BACnetTimerStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetTimerStateTagged
}

// BACnetConstructedDataTimerStateExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTimerState.
// This is useful for switch cases.
type BACnetConstructedDataTimerStateExactly interface {
	BACnetConstructedDataTimerState
	isBACnetConstructedDataTimerState() bool
}

// _BACnetConstructedDataTimerState is the data-structure of this message
type _BACnetConstructedDataTimerState struct {
	*_BACnetConstructedData
	TimerState BACnetTimerStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimerState) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimerState) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIMER_STATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimerState) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTimerState) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimerState) GetTimerState() BACnetTimerStateTagged {
	return m.TimerState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimerState) GetActualValue() BACnetTimerStateTagged {
	return CastBACnetTimerStateTagged(m.GetTimerState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimerState factory function for _BACnetConstructedDataTimerState
func NewBACnetConstructedDataTimerState(timerState BACnetTimerStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimerState {
	_result := &_BACnetConstructedDataTimerState{
		TimerState:             timerState,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimerState(structType interface{}) BACnetConstructedDataTimerState {
	if casted, ok := structType.(BACnetConstructedDataTimerState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimerState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimerState) GetTypeName() string {
	return "BACnetConstructedDataTimerState"
}

func (m *_BACnetConstructedDataTimerState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataTimerState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timerState)
	lengthInBits += m.TimerState.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimerState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimerStateParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimerState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimerState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimerState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timerState)
	if pullErr := readBuffer.PullContext("timerState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timerState")
	}
	_timerState, _timerStateErr := BACnetTimerStateTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _timerStateErr != nil {
		return nil, errors.Wrap(_timerStateErr, "Error parsing 'timerState' field of BACnetConstructedDataTimerState")
	}
	timerState := _timerState.(BACnetTimerStateTagged)
	if closeErr := readBuffer.CloseContext("timerState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timerState")
	}

	// Virtual field
	_actualValue := timerState
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimerState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimerState")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTimerState{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		TimerState: timerState,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTimerState) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimerState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimerState")
		}

		// Simple Field (timerState)
		if pushErr := writeBuffer.PushContext("timerState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timerState")
		}
		_timerStateErr := writeBuffer.WriteSerializable(m.GetTimerState())
		if popErr := writeBuffer.PopContext("timerState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timerState")
		}
		if _timerStateErr != nil {
			return errors.Wrap(_timerStateErr, "Error serializing 'timerState' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimerState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimerState")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimerState) isBACnetConstructedDataTimerState() bool {
	return true
}

func (m *_BACnetConstructedDataTimerState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
