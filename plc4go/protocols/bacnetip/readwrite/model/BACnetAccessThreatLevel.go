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

// BACnetAccessThreatLevel is the corresponding interface of BACnetAccessThreatLevel
type BACnetAccessThreatLevel interface {
	utils.LengthAware
	utils.Serializable
	// GetThreatLevel returns ThreatLevel (property field)
	GetThreatLevel() BACnetApplicationTagUnsignedInteger
}

// BACnetAccessThreatLevelExactly can be used when we want exactly this type and not a type which fulfills BACnetAccessThreatLevel.
// This is useful for switch cases.
type BACnetAccessThreatLevelExactly interface {
	BACnetAccessThreatLevel
	isBACnetAccessThreatLevel() bool
}

// _BACnetAccessThreatLevel is the data-structure of this message
type _BACnetAccessThreatLevel struct {
	ThreatLevel BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessThreatLevel) GetThreatLevel() BACnetApplicationTagUnsignedInteger {
	return m.ThreatLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAccessThreatLevel factory function for _BACnetAccessThreatLevel
func NewBACnetAccessThreatLevel(threatLevel BACnetApplicationTagUnsignedInteger) *_BACnetAccessThreatLevel {
	return &_BACnetAccessThreatLevel{ThreatLevel: threatLevel}
}

// Deprecated: use the interface for direct cast
func CastBACnetAccessThreatLevel(structType interface{}) BACnetAccessThreatLevel {
	if casted, ok := structType.(BACnetAccessThreatLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessThreatLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessThreatLevel) GetTypeName() string {
	return "BACnetAccessThreatLevel"
}

func (m *_BACnetAccessThreatLevel) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetAccessThreatLevel) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (threatLevel)
	lengthInBits += m.ThreatLevel.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetAccessThreatLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAccessThreatLevelParse(readBuffer utils.ReadBuffer) (BACnetAccessThreatLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccessThreatLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessThreatLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (threatLevel)
	if pullErr := readBuffer.PullContext("threatLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for threatLevel")
	}
	_threatLevel, _threatLevelErr := BACnetApplicationTagParse(readBuffer)
	if _threatLevelErr != nil {
		return nil, errors.Wrap(_threatLevelErr, "Error parsing 'threatLevel' field of BACnetAccessThreatLevel")
	}
	threatLevel := _threatLevel.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("threatLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for threatLevel")
	}

	if closeErr := readBuffer.CloseContext("BACnetAccessThreatLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessThreatLevel")
	}

	// Create the instance
	return &_BACnetAccessThreatLevel{
		ThreatLevel: threatLevel,
	}, nil
}

func (m *_BACnetAccessThreatLevel) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetAccessThreatLevel"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessThreatLevel")
	}

	// Simple Field (threatLevel)
	if pushErr := writeBuffer.PushContext("threatLevel"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for threatLevel")
	}
	_threatLevelErr := writeBuffer.WriteSerializable(m.GetThreatLevel())
	if popErr := writeBuffer.PopContext("threatLevel"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for threatLevel")
	}
	if _threatLevelErr != nil {
		return errors.Wrap(_threatLevelErr, "Error serializing 'threatLevel' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccessThreatLevel"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessThreatLevel")
	}
	return nil
}

func (m *_BACnetAccessThreatLevel) isBACnetAccessThreatLevel() bool {
	return true
}

func (m *_BACnetAccessThreatLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
