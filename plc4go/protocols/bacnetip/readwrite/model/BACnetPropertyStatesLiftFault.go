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

// BACnetPropertyStatesLiftFault is the corresponding interface of BACnetPropertyStatesLiftFault
type BACnetPropertyStatesLiftFault interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLiftFault returns LiftFault (property field)
	GetLiftFault() BACnetLiftFaultTagged
}

// BACnetPropertyStatesLiftFaultExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesLiftFault.
// This is useful for switch cases.
type BACnetPropertyStatesLiftFaultExactly interface {
	BACnetPropertyStatesLiftFault
	isBACnetPropertyStatesLiftFault() bool
}

// _BACnetPropertyStatesLiftFault is the data-structure of this message
type _BACnetPropertyStatesLiftFault struct {
	*_BACnetPropertyStates
	LiftFault BACnetLiftFaultTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLiftFault) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesLiftFault) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLiftFault) GetLiftFault() BACnetLiftFaultTagged {
	return m.LiftFault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLiftFault factory function for _BACnetPropertyStatesLiftFault
func NewBACnetPropertyStatesLiftFault(liftFault BACnetLiftFaultTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLiftFault {
	_result := &_BACnetPropertyStatesLiftFault{
		LiftFault:             liftFault,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLiftFault(structType interface{}) BACnetPropertyStatesLiftFault {
	if casted, ok := structType.(BACnetPropertyStatesLiftFault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLiftFault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLiftFault) GetTypeName() string {
	return "BACnetPropertyStatesLiftFault"
}

func (m *_BACnetPropertyStatesLiftFault) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesLiftFault) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (liftFault)
	lengthInBits += m.LiftFault.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesLiftFault) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesLiftFaultParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesLiftFault, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLiftFault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLiftFault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (liftFault)
	if pullErr := readBuffer.PullContext("liftFault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for liftFault")
	}
	_liftFault, _liftFaultErr := BACnetLiftFaultTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _liftFaultErr != nil {
		return nil, errors.Wrap(_liftFaultErr, "Error parsing 'liftFault' field of BACnetPropertyStatesLiftFault")
	}
	liftFault := _liftFault.(BACnetLiftFaultTagged)
	if closeErr := readBuffer.CloseContext("liftFault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for liftFault")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLiftFault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLiftFault")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesLiftFault{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		LiftFault:             liftFault,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesLiftFault) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLiftFault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLiftFault")
		}

		// Simple Field (liftFault)
		if pushErr := writeBuffer.PushContext("liftFault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for liftFault")
		}
		_liftFaultErr := writeBuffer.WriteSerializable(m.GetLiftFault())
		if popErr := writeBuffer.PopContext("liftFault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for liftFault")
		}
		if _liftFaultErr != nil {
			return errors.Wrap(_liftFaultErr, "Error serializing 'liftFault' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLiftFault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLiftFault")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLiftFault) isBACnetPropertyStatesLiftFault() bool {
	return true
}

func (m *_BACnetPropertyStatesLiftFault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
