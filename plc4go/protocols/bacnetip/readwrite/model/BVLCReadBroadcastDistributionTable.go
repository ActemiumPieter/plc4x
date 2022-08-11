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

// BVLCReadBroadcastDistributionTable is the corresponding interface of BVLCReadBroadcastDistributionTable
type BVLCReadBroadcastDistributionTable interface {
	utils.LengthAware
	utils.Serializable
	BVLC
}

// BVLCReadBroadcastDistributionTableExactly can be used when we want exactly this type and not a type which fulfills BVLCReadBroadcastDistributionTable.
// This is useful for switch cases.
type BVLCReadBroadcastDistributionTableExactly interface {
	BVLCReadBroadcastDistributionTable
	isBVLCReadBroadcastDistributionTable() bool
}

// _BVLCReadBroadcastDistributionTable is the data-structure of this message
type _BVLCReadBroadcastDistributionTable struct {
	*_BVLC
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCReadBroadcastDistributionTable) GetBvlcFunction() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCReadBroadcastDistributionTable) InitializeParent(parent BVLC) {}

func (m *_BVLCReadBroadcastDistributionTable) GetParent() BVLC {
	return m._BVLC
}

// NewBVLCReadBroadcastDistributionTable factory function for _BVLCReadBroadcastDistributionTable
func NewBVLCReadBroadcastDistributionTable() *_BVLCReadBroadcastDistributionTable {
	_result := &_BVLCReadBroadcastDistributionTable{
		_BVLC: NewBVLC(),
	}
	_result._BVLC._BVLCChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBVLCReadBroadcastDistributionTable(structType interface{}) BVLCReadBroadcastDistributionTable {
	if casted, ok := structType.(BVLCReadBroadcastDistributionTable); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCReadBroadcastDistributionTable); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCReadBroadcastDistributionTable) GetTypeName() string {
	return "BVLCReadBroadcastDistributionTable"
}

func (m *_BVLCReadBroadcastDistributionTable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BVLCReadBroadcastDistributionTable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_BVLCReadBroadcastDistributionTable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCReadBroadcastDistributionTableParse(readBuffer utils.ReadBuffer) (BVLCReadBroadcastDistributionTable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCReadBroadcastDistributionTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCReadBroadcastDistributionTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BVLCReadBroadcastDistributionTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCReadBroadcastDistributionTable")
	}

	// Create a partially initialized instance
	_child := &_BVLCReadBroadcastDistributionTable{
		_BVLC: &_BVLC{},
	}
	_child._BVLC._BVLCChildRequirements = _child
	return _child, nil
}

func (m *_BVLCReadBroadcastDistributionTable) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCReadBroadcastDistributionTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCReadBroadcastDistributionTable")
		}

		if popErr := writeBuffer.PopContext("BVLCReadBroadcastDistributionTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCReadBroadcastDistributionTable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BVLCReadBroadcastDistributionTable) isBVLCReadBroadcastDistributionTable() bool {
	return true
}

func (m *_BVLCReadBroadcastDistributionTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
