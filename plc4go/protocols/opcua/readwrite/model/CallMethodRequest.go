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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// CallMethodRequest is the corresponding interface of CallMethodRequest
type CallMethodRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetObjectId returns ObjectId (property field)
	GetObjectId() NodeId
	// GetMethodId returns MethodId (property field)
	GetMethodId() NodeId
	// GetNoOfInputArguments returns NoOfInputArguments (property field)
	GetNoOfInputArguments() int32
	// GetInputArguments returns InputArguments (property field)
	GetInputArguments() []Variant
}

// CallMethodRequestExactly can be used when we want exactly this type and not a type which fulfills CallMethodRequest.
// This is useful for switch cases.
type CallMethodRequestExactly interface {
	CallMethodRequest
	isCallMethodRequest() bool
}

// _CallMethodRequest is the data-structure of this message
type _CallMethodRequest struct {
	*_ExtensionObjectDefinition
	ObjectId           NodeId
	MethodId           NodeId
	NoOfInputArguments int32
	InputArguments     []Variant
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CallMethodRequest) GetIdentifier() string {
	return "706"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CallMethodRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_CallMethodRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CallMethodRequest) GetObjectId() NodeId {
	return m.ObjectId
}

func (m *_CallMethodRequest) GetMethodId() NodeId {
	return m.MethodId
}

func (m *_CallMethodRequest) GetNoOfInputArguments() int32 {
	return m.NoOfInputArguments
}

func (m *_CallMethodRequest) GetInputArguments() []Variant {
	return m.InputArguments
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCallMethodRequest factory function for _CallMethodRequest
func NewCallMethodRequest(objectId NodeId, methodId NodeId, noOfInputArguments int32, inputArguments []Variant) *_CallMethodRequest {
	_result := &_CallMethodRequest{
		ObjectId:                   objectId,
		MethodId:                   methodId,
		NoOfInputArguments:         noOfInputArguments,
		InputArguments:             inputArguments,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCallMethodRequest(structType any) CallMethodRequest {
	if casted, ok := structType.(CallMethodRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CallMethodRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CallMethodRequest) GetTypeName() string {
	return "CallMethodRequest"
}

func (m *_CallMethodRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (objectId)
	lengthInBits += m.ObjectId.GetLengthInBits(ctx)

	// Simple field (methodId)
	lengthInBits += m.MethodId.GetLengthInBits(ctx)

	// Simple field (noOfInputArguments)
	lengthInBits += 32

	// Array field
	if len(m.InputArguments) > 0 {
		for _curItem, element := range m.InputArguments {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.InputArguments), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CallMethodRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CallMethodRequestParse(ctx context.Context, theBytes []byte, identifier string) (CallMethodRequest, error) {
	return CallMethodRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func CallMethodRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (CallMethodRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CallMethodRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CallMethodRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectId)
	if pullErr := readBuffer.PullContext("objectId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectId")
	}
	_objectId, _objectIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _objectIdErr != nil {
		return nil, errors.Wrap(_objectIdErr, "Error parsing 'objectId' field of CallMethodRequest")
	}
	objectId := _objectId.(NodeId)
	if closeErr := readBuffer.CloseContext("objectId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectId")
	}

	// Simple Field (methodId)
	if pullErr := readBuffer.PullContext("methodId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for methodId")
	}
	_methodId, _methodIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _methodIdErr != nil {
		return nil, errors.Wrap(_methodIdErr, "Error parsing 'methodId' field of CallMethodRequest")
	}
	methodId := _methodId.(NodeId)
	if closeErr := readBuffer.CloseContext("methodId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for methodId")
	}

	// Simple Field (noOfInputArguments)
	_noOfInputArguments, _noOfInputArgumentsErr := readBuffer.ReadInt32("noOfInputArguments", 32)
	if _noOfInputArgumentsErr != nil {
		return nil, errors.Wrap(_noOfInputArgumentsErr, "Error parsing 'noOfInputArguments' field of CallMethodRequest")
	}
	noOfInputArguments := _noOfInputArguments

	// Array field (inputArguments)
	if pullErr := readBuffer.PullContext("inputArguments", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for inputArguments")
	}
	// Count array
	inputArguments := make([]Variant, utils.Max(noOfInputArguments, 0))
	// This happens when the size is set conditional to 0
	if len(inputArguments) == 0 {
		inputArguments = nil
	}
	{
		_numItems := uint16(utils.Max(noOfInputArguments, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := VariantParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'inputArguments' field of CallMethodRequest")
			}
			inputArguments[_curItem] = _item.(Variant)
		}
	}
	if closeErr := readBuffer.CloseContext("inputArguments", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for inputArguments")
	}

	if closeErr := readBuffer.CloseContext("CallMethodRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CallMethodRequest")
	}

	// Create a partially initialized instance
	_child := &_CallMethodRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ObjectId:                   objectId,
		MethodId:                   methodId,
		NoOfInputArguments:         noOfInputArguments,
		InputArguments:             inputArguments,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_CallMethodRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CallMethodRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CallMethodRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CallMethodRequest")
		}

		// Simple Field (objectId)
		if pushErr := writeBuffer.PushContext("objectId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectId")
		}
		_objectIdErr := writeBuffer.WriteSerializable(ctx, m.GetObjectId())
		if popErr := writeBuffer.PopContext("objectId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectId")
		}
		if _objectIdErr != nil {
			return errors.Wrap(_objectIdErr, "Error serializing 'objectId' field")
		}

		// Simple Field (methodId)
		if pushErr := writeBuffer.PushContext("methodId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for methodId")
		}
		_methodIdErr := writeBuffer.WriteSerializable(ctx, m.GetMethodId())
		if popErr := writeBuffer.PopContext("methodId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for methodId")
		}
		if _methodIdErr != nil {
			return errors.Wrap(_methodIdErr, "Error serializing 'methodId' field")
		}

		// Simple Field (noOfInputArguments)
		noOfInputArguments := int32(m.GetNoOfInputArguments())
		_noOfInputArgumentsErr := writeBuffer.WriteInt32("noOfInputArguments", 32, (noOfInputArguments))
		if _noOfInputArgumentsErr != nil {
			return errors.Wrap(_noOfInputArgumentsErr, "Error serializing 'noOfInputArguments' field")
		}

		// Array Field (inputArguments)
		if pushErr := writeBuffer.PushContext("inputArguments", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for inputArguments")
		}
		for _curItem, _element := range m.GetInputArguments() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetInputArguments()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'inputArguments' field")
			}
		}
		if popErr := writeBuffer.PopContext("inputArguments", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for inputArguments")
		}

		if popErr := writeBuffer.PopContext("CallMethodRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CallMethodRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CallMethodRequest) isCallMethodRequest() bool {
	return true
}

func (m *_CallMethodRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
