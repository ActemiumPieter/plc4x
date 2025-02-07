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

// ChannelSecurityToken is the corresponding interface of ChannelSecurityToken
type ChannelSecurityToken interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetChannelId returns ChannelId (property field)
	GetChannelId() uint32
	// GetTokenId returns TokenId (property field)
	GetTokenId() uint32
	// GetCreatedAt returns CreatedAt (property field)
	GetCreatedAt() int64
	// GetRevisedLifetime returns RevisedLifetime (property field)
	GetRevisedLifetime() uint32
}

// ChannelSecurityTokenExactly can be used when we want exactly this type and not a type which fulfills ChannelSecurityToken.
// This is useful for switch cases.
type ChannelSecurityTokenExactly interface {
	ChannelSecurityToken
	isChannelSecurityToken() bool
}

// _ChannelSecurityToken is the data-structure of this message
type _ChannelSecurityToken struct {
	*_ExtensionObjectDefinition
	ChannelId       uint32
	TokenId         uint32
	CreatedAt       int64
	RevisedLifetime uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ChannelSecurityToken) GetIdentifier() string {
	return "443"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ChannelSecurityToken) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ChannelSecurityToken) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ChannelSecurityToken) GetChannelId() uint32 {
	return m.ChannelId
}

func (m *_ChannelSecurityToken) GetTokenId() uint32 {
	return m.TokenId
}

func (m *_ChannelSecurityToken) GetCreatedAt() int64 {
	return m.CreatedAt
}

func (m *_ChannelSecurityToken) GetRevisedLifetime() uint32 {
	return m.RevisedLifetime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewChannelSecurityToken factory function for _ChannelSecurityToken
func NewChannelSecurityToken(channelId uint32, tokenId uint32, createdAt int64, revisedLifetime uint32) *_ChannelSecurityToken {
	_result := &_ChannelSecurityToken{
		ChannelId:                  channelId,
		TokenId:                    tokenId,
		CreatedAt:                  createdAt,
		RevisedLifetime:            revisedLifetime,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastChannelSecurityToken(structType any) ChannelSecurityToken {
	if casted, ok := structType.(ChannelSecurityToken); ok {
		return casted
	}
	if casted, ok := structType.(*ChannelSecurityToken); ok {
		return *casted
	}
	return nil
}

func (m *_ChannelSecurityToken) GetTypeName() string {
	return "ChannelSecurityToken"
}

func (m *_ChannelSecurityToken) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (channelId)
	lengthInBits += 32

	// Simple field (tokenId)
	lengthInBits += 32

	// Simple field (createdAt)
	lengthInBits += 64

	// Simple field (revisedLifetime)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ChannelSecurityToken) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ChannelSecurityTokenParse(ctx context.Context, theBytes []byte, identifier string) (ChannelSecurityToken, error) {
	return ChannelSecurityTokenParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ChannelSecurityTokenParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ChannelSecurityToken, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ChannelSecurityToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ChannelSecurityToken")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (channelId)
	_channelId, _channelIdErr := readBuffer.ReadUint32("channelId", 32)
	if _channelIdErr != nil {
		return nil, errors.Wrap(_channelIdErr, "Error parsing 'channelId' field of ChannelSecurityToken")
	}
	channelId := _channelId

	// Simple Field (tokenId)
	_tokenId, _tokenIdErr := readBuffer.ReadUint32("tokenId", 32)
	if _tokenIdErr != nil {
		return nil, errors.Wrap(_tokenIdErr, "Error parsing 'tokenId' field of ChannelSecurityToken")
	}
	tokenId := _tokenId

	// Simple Field (createdAt)
	_createdAt, _createdAtErr := readBuffer.ReadInt64("createdAt", 64)
	if _createdAtErr != nil {
		return nil, errors.Wrap(_createdAtErr, "Error parsing 'createdAt' field of ChannelSecurityToken")
	}
	createdAt := _createdAt

	// Simple Field (revisedLifetime)
	_revisedLifetime, _revisedLifetimeErr := readBuffer.ReadUint32("revisedLifetime", 32)
	if _revisedLifetimeErr != nil {
		return nil, errors.Wrap(_revisedLifetimeErr, "Error parsing 'revisedLifetime' field of ChannelSecurityToken")
	}
	revisedLifetime := _revisedLifetime

	if closeErr := readBuffer.CloseContext("ChannelSecurityToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ChannelSecurityToken")
	}

	// Create a partially initialized instance
	_child := &_ChannelSecurityToken{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ChannelId:                  channelId,
		TokenId:                    tokenId,
		CreatedAt:                  createdAt,
		RevisedLifetime:            revisedLifetime,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ChannelSecurityToken) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ChannelSecurityToken) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ChannelSecurityToken"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ChannelSecurityToken")
		}

		// Simple Field (channelId)
		channelId := uint32(m.GetChannelId())
		_channelIdErr := writeBuffer.WriteUint32("channelId", 32, (channelId))
		if _channelIdErr != nil {
			return errors.Wrap(_channelIdErr, "Error serializing 'channelId' field")
		}

		// Simple Field (tokenId)
		tokenId := uint32(m.GetTokenId())
		_tokenIdErr := writeBuffer.WriteUint32("tokenId", 32, (tokenId))
		if _tokenIdErr != nil {
			return errors.Wrap(_tokenIdErr, "Error serializing 'tokenId' field")
		}

		// Simple Field (createdAt)
		createdAt := int64(m.GetCreatedAt())
		_createdAtErr := writeBuffer.WriteInt64("createdAt", 64, (createdAt))
		if _createdAtErr != nil {
			return errors.Wrap(_createdAtErr, "Error serializing 'createdAt' field")
		}

		// Simple Field (revisedLifetime)
		revisedLifetime := uint32(m.GetRevisedLifetime())
		_revisedLifetimeErr := writeBuffer.WriteUint32("revisedLifetime", 32, (revisedLifetime))
		if _revisedLifetimeErr != nil {
			return errors.Wrap(_revisedLifetimeErr, "Error serializing 'revisedLifetime' field")
		}

		if popErr := writeBuffer.PopContext("ChannelSecurityToken"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ChannelSecurityToken")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ChannelSecurityToken) isChannelSecurityToken() bool {
	return true
}

func (m *_ChannelSecurityToken) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
