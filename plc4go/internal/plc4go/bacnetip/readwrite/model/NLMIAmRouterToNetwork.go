//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type NLMIAmRouterToNetwork struct {
	DestinationNetworkAddress []uint16
	Parent                    *NLM
}

// The corresponding interface
type INLMIAmRouterToNetwork interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *NLMIAmRouterToNetwork) MessageType() uint8 {
	return 0x1
}

func (m *NLMIAmRouterToNetwork) InitializeParent(parent *NLM, vendorId *uint16) {
	m.Parent.VendorId = vendorId
}

func NewNLMIAmRouterToNetwork(destinationNetworkAddress []uint16, vendorId *uint16) *NLM {
	child := &NLMIAmRouterToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		Parent:                    NewNLM(vendorId),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastNLMIAmRouterToNetwork(structType interface{}) *NLMIAmRouterToNetwork {
	castFunc := func(typ interface{}) *NLMIAmRouterToNetwork {
		if casted, ok := typ.(NLMIAmRouterToNetwork); ok {
			return &casted
		}
		if casted, ok := typ.(*NLMIAmRouterToNetwork); ok {
			return casted
		}
		if casted, ok := typ.(NLM); ok {
			return CastNLMIAmRouterToNetwork(casted.Child)
		}
		if casted, ok := typ.(*NLM); ok {
			return CastNLMIAmRouterToNetwork(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *NLMIAmRouterToNetwork) GetTypeName() string {
	return "NLMIAmRouterToNetwork"
}

func (m *NLMIAmRouterToNetwork) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *NLMIAmRouterToNetwork) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Array field
	if len(m.DestinationNetworkAddress) > 0 {
		lengthInBits += 16 * uint16(len(m.DestinationNetworkAddress))
	}

	return lengthInBits
}

func (m *NLMIAmRouterToNetwork) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func NLMIAmRouterToNetworkParse(io utils.ReadBuffer, apduLength uint16, messageType uint8) (*NLM, error) {
	io.PullContext("NLMIAmRouterToNetwork")
	io.PullContext("destinationNetworkAddress")

	// Array field (destinationNetworkAddress)
	// Length array
	destinationNetworkAddress := make([]uint16, 0)
	_destinationNetworkAddressLength := uint16(apduLength) - uint16(uint16(utils.InlineIf(bool(bool(bool(bool((messageType) >= (128)))) && bool(bool(bool((messageType) <= (255))))), func() uint16 { return uint16(uint16(3)) }, func() uint16 { return uint16(uint16(1)) })))
	_destinationNetworkAddressEndPos := io.GetPos() + uint16(_destinationNetworkAddressLength)
	for io.GetPos() < _destinationNetworkAddressEndPos {
		_item, _err := io.ReadUint16("", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'destinationNetworkAddress' field")
		}
		destinationNetworkAddress = append(destinationNetworkAddress, _item)
	}

	io.CloseContext("NLMIAmRouterToNetwork")

	// Create a partially initialized instance
	_child := &NLMIAmRouterToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		Parent:                    &NLM{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *NLMIAmRouterToNetwork) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("NLMIAmRouterToNetwork")

		// Array Field (destinationNetworkAddress)
		if m.DestinationNetworkAddress != nil {
			io.PushContext("destinationNetworkAddress")
			for _, _element := range m.DestinationNetworkAddress {
				_elementErr := io.WriteUint16("", 16, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'destinationNetworkAddress' field")
				}
			}
			io.PopContext("destinationNetworkAddress")
		}

		io.PopContext("NLMIAmRouterToNetwork")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *NLMIAmRouterToNetwork) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "destinationNetworkAddress":
				var data []uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DestinationNetworkAddress = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *NLMIAmRouterToNetwork) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.DestinationNetworkAddress, xml.StartElement{Name: xml.Name{Local: "destinationNetworkAddress"}}); err != nil {
		return err
	}
	return nil
}

func (m NLMIAmRouterToNetwork) String() string {
	return string(m.Box("", 120))
}

func (m NLMIAmRouterToNetwork) Box(name string, width int) utils.AsciiBox {
	boxName := "NLMIAmRouterToNetwork"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Array Field (destinationNetworkAddress)
		if m.DestinationNetworkAddress != nil {
			// Simple array base type uint16 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.DestinationNetworkAddress {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("DestinationNetworkAddress", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
