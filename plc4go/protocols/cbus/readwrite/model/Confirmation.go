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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// Confirmation is the corresponding interface of Confirmation
type Confirmation interface {
	utils.LengthAware
	utils.Serializable
	// GetAlpha returns Alpha (property field)
	GetAlpha() Alpha
	// GetSecondAlpha returns SecondAlpha (property field)
	GetSecondAlpha() Alpha
	// GetConfirmationType returns ConfirmationType (property field)
	GetConfirmationType() ConfirmationType
	// GetIsSuccess returns IsSuccess (virtual field)
	GetIsSuccess() bool
}

// ConfirmationExactly can be used when we want exactly this type and not a type which fulfills Confirmation.
// This is useful for switch cases.
type ConfirmationExactly interface {
	Confirmation
	isConfirmation() bool
}

// _Confirmation is the data-structure of this message
type _Confirmation struct {
	Alpha            Alpha
	SecondAlpha      Alpha
	ConfirmationType ConfirmationType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Confirmation) GetAlpha() Alpha {
	return m.Alpha
}

func (m *_Confirmation) GetSecondAlpha() Alpha {
	return m.SecondAlpha
}

func (m *_Confirmation) GetConfirmationType() ConfirmationType {
	return m.ConfirmationType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_Confirmation) GetIsSuccess() bool {
	secondAlpha := m.SecondAlpha
	_ = secondAlpha
	return bool(bool((m.GetConfirmationType()) == (ConfirmationType_CONFIRMATION_SUCCESSFUL)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConfirmation factory function for _Confirmation
func NewConfirmation(alpha Alpha, secondAlpha Alpha, confirmationType ConfirmationType) *_Confirmation {
	return &_Confirmation{Alpha: alpha, SecondAlpha: secondAlpha, ConfirmationType: confirmationType}
}

// Deprecated: use the interface for direct cast
func CastConfirmation(structType interface{}) Confirmation {
	if casted, ok := structType.(Confirmation); ok {
		return casted
	}
	if casted, ok := structType.(*Confirmation); ok {
		return *casted
	}
	return nil
}

func (m *_Confirmation) GetTypeName() string {
	return "Confirmation"
}

func (m *_Confirmation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_Confirmation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (alpha)
	lengthInBits += m.Alpha.GetLengthInBits()

	// Optional Field (secondAlpha)
	if m.SecondAlpha != nil {
		lengthInBits += m.SecondAlpha.GetLengthInBits()
	}

	// Simple field (confirmationType)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_Confirmation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ConfirmationParse(readBuffer utils.ReadBuffer) (Confirmation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Confirmation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Confirmation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (alpha)
	if pullErr := readBuffer.PullContext("alpha"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alpha")
	}
	_alpha, _alphaErr := AlphaParse(readBuffer)
	if _alphaErr != nil {
		return nil, errors.Wrap(_alphaErr, "Error parsing 'alpha' field of Confirmation")
	}
	alpha := _alpha.(Alpha)
	if closeErr := readBuffer.CloseContext("alpha"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alpha")
	}

	// Optional Field (secondAlpha) (Can be skipped, if a given expression evaluates to false)
	var secondAlpha Alpha = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("secondAlpha"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for secondAlpha")
		}
		_val, _err := AlphaParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'secondAlpha' field of Confirmation")
		default:
			secondAlpha = _val.(Alpha)
			if closeErr := readBuffer.CloseContext("secondAlpha"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for secondAlpha")
			}
		}
	}

	// Simple Field (confirmationType)
	if pullErr := readBuffer.PullContext("confirmationType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for confirmationType")
	}
	_confirmationType, _confirmationTypeErr := ConfirmationTypeParse(readBuffer)
	if _confirmationTypeErr != nil {
		return nil, errors.Wrap(_confirmationTypeErr, "Error parsing 'confirmationType' field of Confirmation")
	}
	confirmationType := _confirmationType
	if closeErr := readBuffer.CloseContext("confirmationType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for confirmationType")
	}

	// Virtual field
	_isSuccess := bool((confirmationType) == (ConfirmationType_CONFIRMATION_SUCCESSFUL))
	isSuccess := bool(_isSuccess)
	_ = isSuccess

	if closeErr := readBuffer.CloseContext("Confirmation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Confirmation")
	}

	// Create the instance
	return &_Confirmation{
		Alpha:            alpha,
		SecondAlpha:      secondAlpha,
		ConfirmationType: confirmationType,
	}, nil
}

func (m *_Confirmation) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("Confirmation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Confirmation")
	}

	// Simple Field (alpha)
	if pushErr := writeBuffer.PushContext("alpha"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for alpha")
	}
	_alphaErr := writeBuffer.WriteSerializable(m.GetAlpha())
	if popErr := writeBuffer.PopContext("alpha"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for alpha")
	}
	if _alphaErr != nil {
		return errors.Wrap(_alphaErr, "Error serializing 'alpha' field")
	}

	// Optional Field (secondAlpha) (Can be skipped, if the value is null)
	var secondAlpha Alpha = nil
	if m.GetSecondAlpha() != nil {
		if pushErr := writeBuffer.PushContext("secondAlpha"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for secondAlpha")
		}
		secondAlpha = m.GetSecondAlpha()
		_secondAlphaErr := writeBuffer.WriteSerializable(secondAlpha)
		if popErr := writeBuffer.PopContext("secondAlpha"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for secondAlpha")
		}
		if _secondAlphaErr != nil {
			return errors.Wrap(_secondAlphaErr, "Error serializing 'secondAlpha' field")
		}
	}

	// Simple Field (confirmationType)
	if pushErr := writeBuffer.PushContext("confirmationType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for confirmationType")
	}
	_confirmationTypeErr := writeBuffer.WriteSerializable(m.GetConfirmationType())
	if popErr := writeBuffer.PopContext("confirmationType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for confirmationType")
	}
	if _confirmationTypeErr != nil {
		return errors.Wrap(_confirmationTypeErr, "Error serializing 'confirmationType' field")
	}
	// Virtual field
	if _isSuccessErr := writeBuffer.WriteVirtual("isSuccess", m.GetIsSuccess()); _isSuccessErr != nil {
		return errors.Wrap(_isSuccessErr, "Error serializing 'isSuccess' field")
	}

	if popErr := writeBuffer.PopContext("Confirmation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Confirmation")
	}
	return nil
}

func (m *_Confirmation) isConfirmation() bool {
	return true
}

func (m *_Confirmation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
