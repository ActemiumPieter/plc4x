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

// BACnetPropertyStatesAction is the corresponding interface of BACnetPropertyStatesAction
type BACnetPropertyStatesAction interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetAction returns Action (property field)
	GetAction() BACnetActionTagged
}

// BACnetPropertyStatesActionExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesAction.
// This is useful for switch cases.
type BACnetPropertyStatesActionExactly interface {
	BACnetPropertyStatesAction
	isBACnetPropertyStatesAction() bool
}

// _BACnetPropertyStatesAction is the data-structure of this message
type _BACnetPropertyStatesAction struct {
	*_BACnetPropertyStates
	Action BACnetActionTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesAction) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesAction) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesAction) GetAction() BACnetActionTagged {
	return m.Action
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesAction factory function for _BACnetPropertyStatesAction
func NewBACnetPropertyStatesAction(action BACnetActionTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesAction {
	_result := &_BACnetPropertyStatesAction{
		Action:                action,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesAction(structType interface{}) BACnetPropertyStatesAction {
	if casted, ok := structType.(BACnetPropertyStatesAction); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesAction); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesAction) GetTypeName() string {
	return "BACnetPropertyStatesAction"
}

func (m *_BACnetPropertyStatesAction) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesAction) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (action)
	lengthInBits += m.Action.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesAction) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesActionParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesAction, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesAction"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesAction")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (action)
	if pullErr := readBuffer.PullContext("action"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for action")
	}
	_action, _actionErr := BACnetActionTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _actionErr != nil {
		return nil, errors.Wrap(_actionErr, "Error parsing 'action' field of BACnetPropertyStatesAction")
	}
	action := _action.(BACnetActionTagged)
	if closeErr := readBuffer.CloseContext("action"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for action")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesAction"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesAction")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesAction{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		Action:                action,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesAction) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesAction"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesAction")
		}

		// Simple Field (action)
		if pushErr := writeBuffer.PushContext("action"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for action")
		}
		_actionErr := writeBuffer.WriteSerializable(m.GetAction())
		if popErr := writeBuffer.PopContext("action"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for action")
		}
		if _actionErr != nil {
			return errors.Wrap(_actionErr, "Error serializing 'action' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesAction"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesAction")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesAction) isBACnetPropertyStatesAction() bool {
	return true
}

func (m *_BACnetPropertyStatesAction) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
