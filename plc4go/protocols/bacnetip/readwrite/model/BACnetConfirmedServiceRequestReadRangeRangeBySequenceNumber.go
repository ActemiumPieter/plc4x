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

// BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber is the corresponding interface of BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber
type BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequestReadRangeRange
	// GetReferenceSequenceNumber returns ReferenceSequenceNumber (property field)
	GetReferenceSequenceNumber() BACnetApplicationTagUnsignedInteger
	// GetCount returns Count (property field)
	GetCount() BACnetApplicationTagSignedInteger
}

// BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberExactly interface {
	BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber
	isBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber() bool
}

// _BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber is the data-structure of this message
type _BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber struct {
	*_BACnetConfirmedServiceRequestReadRangeRange
	ReferenceSequenceNumber BACnetApplicationTagUnsignedInteger
	Count                   BACnetApplicationTagSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) InitializeParent(parent BACnetConfirmedServiceRequestReadRangeRange, peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) {
	m.PeekedTagHeader = peekedTagHeader
	m.OpeningTag = openingTag
	m.ClosingTag = closingTag
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetParent() BACnetConfirmedServiceRequestReadRangeRange {
	return m._BACnetConfirmedServiceRequestReadRangeRange
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetReferenceSequenceNumber() BACnetApplicationTagUnsignedInteger {
	return m.ReferenceSequenceNumber
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetCount() BACnetApplicationTagSignedInteger {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber factory function for _BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber
func NewBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber(referenceSequenceNumber BACnetApplicationTagUnsignedInteger, count BACnetApplicationTagSignedInteger, peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber {
	_result := &_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber{
		ReferenceSequenceNumber: referenceSequenceNumber,
		Count:                   count,
		_BACnetConfirmedServiceRequestReadRangeRange: NewBACnetConfirmedServiceRequestReadRangeRange(peekedTagHeader, openingTag, closingTag),
	}
	_result._BACnetConfirmedServiceRequestReadRangeRange._BACnetConfirmedServiceRequestReadRangeRangeChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber(structType interface{}) BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (referenceSequenceNumber)
	lengthInBits += m.ReferenceSequenceNumber.GetLengthInBits()

	// Simple field (count)
	lengthInBits += m.Count.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberParse(readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (referenceSequenceNumber)
	if pullErr := readBuffer.PullContext("referenceSequenceNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referenceSequenceNumber")
	}
	_referenceSequenceNumber, _referenceSequenceNumberErr := BACnetApplicationTagParse(readBuffer)
	if _referenceSequenceNumberErr != nil {
		return nil, errors.Wrap(_referenceSequenceNumberErr, "Error parsing 'referenceSequenceNumber' field of BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber")
	}
	referenceSequenceNumber := _referenceSequenceNumber.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("referenceSequenceNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referenceSequenceNumber")
	}

	// Simple Field (count)
	if pullErr := readBuffer.PullContext("count"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for count")
	}
	_count, _countErr := BACnetApplicationTagParse(readBuffer)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field of BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber")
	}
	count := _count.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("count"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for count")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber{
		_BACnetConfirmedServiceRequestReadRangeRange: &_BACnetConfirmedServiceRequestReadRangeRange{},
		ReferenceSequenceNumber:                      referenceSequenceNumber,
		Count:                                        count,
	}
	_child._BACnetConfirmedServiceRequestReadRangeRange._BACnetConfirmedServiceRequestReadRangeRangeChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber")
		}

		// Simple Field (referenceSequenceNumber)
		if pushErr := writeBuffer.PushContext("referenceSequenceNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for referenceSequenceNumber")
		}
		_referenceSequenceNumberErr := writeBuffer.WriteSerializable(m.GetReferenceSequenceNumber())
		if popErr := writeBuffer.PopContext("referenceSequenceNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for referenceSequenceNumber")
		}
		if _referenceSequenceNumberErr != nil {
			return errors.Wrap(_referenceSequenceNumberErr, "Error serializing 'referenceSequenceNumber' field")
		}

		// Simple Field (count)
		if pushErr := writeBuffer.PushContext("count"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for count")
		}
		_countErr := writeBuffer.WriteSerializable(m.GetCount())
		if popErr := writeBuffer.PopContext("count"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for count")
		}
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) isBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
