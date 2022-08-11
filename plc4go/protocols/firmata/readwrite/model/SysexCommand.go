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

// SysexCommand is the corresponding interface of SysexCommand
type SysexCommand interface {
	utils.LengthAware
	utils.Serializable
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() uint8
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
}

// SysexCommandExactly can be used when we want exactly this type and not a type which fulfills SysexCommand.
// This is useful for switch cases.
type SysexCommandExactly interface {
	SysexCommand
	isSysexCommand() bool
}

// _SysexCommand is the data-structure of this message
type _SysexCommand struct {
	_SysexCommandChildRequirements
}

type _SysexCommandChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetCommandType() uint8
	GetResponse() bool
}

type SysexCommandParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child SysexCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type SysexCommandChild interface {
	utils.Serializable
	InitializeParent(parent SysexCommand)
	GetParent() *SysexCommand

	GetTypeName() string
	SysexCommand
}

// NewSysexCommand factory function for _SysexCommand
func NewSysexCommand() *_SysexCommand {
	return &_SysexCommand{}
}

// Deprecated: use the interface for direct cast
func CastSysexCommand(structType interface{}) SysexCommand {
	if casted, ok := structType.(SysexCommand); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommand); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommand) GetTypeName() string {
	return "SysexCommand"
}

func (m *_SysexCommand) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (commandType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SysexCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandParse(readBuffer utils.ReadBuffer, response bool) (SysexCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (commandType) (Used as input to a switch field)
	commandType, _commandTypeErr := readBuffer.ReadUint8("commandType", 8)
	if _commandTypeErr != nil {
		return nil, errors.Wrap(_commandTypeErr, "Error parsing 'commandType' field of SysexCommand")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type SysexCommandChildSerializeRequirement interface {
		SysexCommand
		InitializeParent(SysexCommand)
		GetParent() SysexCommand
	}
	var _childTemp interface{}
	var _child SysexCommandChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandType == 0x00: // SysexCommandExtendedId
		_childTemp, typeSwitchError = SysexCommandExtendedIdParse(readBuffer, response)
	case commandType == 0x69 && response == bool(false): // SysexCommandAnalogMappingQueryRequest
		_childTemp, typeSwitchError = SysexCommandAnalogMappingQueryRequestParse(readBuffer, response)
	case commandType == 0x69 && response == bool(true): // SysexCommandAnalogMappingQueryResponse
		_childTemp, typeSwitchError = SysexCommandAnalogMappingQueryResponseParse(readBuffer, response)
	case commandType == 0x6A: // SysexCommandAnalogMappingResponse
		_childTemp, typeSwitchError = SysexCommandAnalogMappingResponseParse(readBuffer, response)
	case commandType == 0x6B: // SysexCommandCapabilityQuery
		_childTemp, typeSwitchError = SysexCommandCapabilityQueryParse(readBuffer, response)
	case commandType == 0x6C: // SysexCommandCapabilityResponse
		_childTemp, typeSwitchError = SysexCommandCapabilityResponseParse(readBuffer, response)
	case commandType == 0x6D: // SysexCommandPinStateQuery
		_childTemp, typeSwitchError = SysexCommandPinStateQueryParse(readBuffer, response)
	case commandType == 0x6E: // SysexCommandPinStateResponse
		_childTemp, typeSwitchError = SysexCommandPinStateResponseParse(readBuffer, response)
	case commandType == 0x6F: // SysexCommandExtendedAnalog
		_childTemp, typeSwitchError = SysexCommandExtendedAnalogParse(readBuffer, response)
	case commandType == 0x71: // SysexCommandStringData
		_childTemp, typeSwitchError = SysexCommandStringDataParse(readBuffer, response)
	case commandType == 0x79 && response == bool(false): // SysexCommandReportFirmwareRequest
		_childTemp, typeSwitchError = SysexCommandReportFirmwareRequestParse(readBuffer, response)
	case commandType == 0x79 && response == bool(true): // SysexCommandReportFirmwareResponse
		_childTemp, typeSwitchError = SysexCommandReportFirmwareResponseParse(readBuffer, response)
	case commandType == 0x7A: // SysexCommandSamplingInterval
		_childTemp, typeSwitchError = SysexCommandSamplingIntervalParse(readBuffer, response)
	case commandType == 0x7E: // SysexCommandSysexNonRealtime
		_childTemp, typeSwitchError = SysexCommandSysexNonRealtimeParse(readBuffer, response)
	case commandType == 0x7F: // SysexCommandSysexRealtime
		_childTemp, typeSwitchError = SysexCommandSysexRealtimeParse(readBuffer, response)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v, response=%v]", commandType, response)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of SysexCommand")
	}
	_child = _childTemp.(SysexCommandChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("SysexCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommand")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_SysexCommand) SerializeParent(writeBuffer utils.WriteBuffer, child SysexCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("SysexCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SysexCommand")
	}

	// Discriminator Field (commandType) (Used as input to a switch field)
	commandType := uint8(child.GetCommandType())
	_commandTypeErr := writeBuffer.WriteUint8("commandType", 8, (commandType))

	if _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("SysexCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SysexCommand")
	}
	return nil
}

func (m *_SysexCommand) isSysexCommand() bool {
	return true
}

func (m *_SysexCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
