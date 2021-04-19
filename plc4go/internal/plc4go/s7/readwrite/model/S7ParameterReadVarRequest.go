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
type S7ParameterReadVarRequest struct {
	Items  []*S7VarRequestParameterItem
	Parent *S7Parameter
}

// The corresponding interface
type IS7ParameterReadVarRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7ParameterReadVarRequest) ParameterType() uint8 {
	return 0x04
}

func (m *S7ParameterReadVarRequest) MessageType() uint8 {
	return 0x01
}

func (m *S7ParameterReadVarRequest) InitializeParent(parent *S7Parameter) {
}

func NewS7ParameterReadVarRequest(items []*S7VarRequestParameterItem) *S7Parameter {
	child := &S7ParameterReadVarRequest{
		Items:  items,
		Parent: NewS7Parameter(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastS7ParameterReadVarRequest(structType interface{}) *S7ParameterReadVarRequest {
	castFunc := func(typ interface{}) *S7ParameterReadVarRequest {
		if casted, ok := typ.(S7ParameterReadVarRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*S7ParameterReadVarRequest); ok {
			return casted
		}
		if casted, ok := typ.(S7Parameter); ok {
			return CastS7ParameterReadVarRequest(casted.Child)
		}
		if casted, ok := typ.(*S7Parameter); ok {
			return CastS7ParameterReadVarRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7ParameterReadVarRequest) GetTypeName() string {
	return "S7ParameterReadVarRequest"
}

func (m *S7ParameterReadVarRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7ParameterReadVarRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Implicit Field (numItems)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *S7ParameterReadVarRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7ParameterReadVarRequestParse(io utils.ReadBuffer) (*S7Parameter, error) {
	io.PullContext("S7ParameterReadVarRequest")

	// Implicit Field (numItems) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	numItems, _numItemsErr := io.ReadUint8("numItems", 8)
	_ = numItems
	if _numItemsErr != nil {
		return nil, errors.Wrap(_numItemsErr, "Error parsing 'numItems' field")
	}
	io.PullContext("items")

	// Array field (items)
	// Count array
	items := make([]*S7VarRequestParameterItem, numItems)
	for curItem := uint16(0); curItem < uint16(numItems); curItem++ {
		_item, _err := S7VarRequestParameterItemParse(io)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'items' field")
		}
		items[curItem] = _item
	}

	io.CloseContext("S7ParameterReadVarRequest")

	// Create a partially initialized instance
	_child := &S7ParameterReadVarRequest{
		Items:  items,
		Parent: &S7Parameter{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *S7ParameterReadVarRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("S7ParameterReadVarRequest")

		// Implicit Field (numItems) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		numItems := uint8(uint8(len(m.Items)))
		_numItemsErr := io.WriteUint8("numItems", 8, (numItems))
		if _numItemsErr != nil {
			return errors.Wrap(_numItemsErr, "Error serializing 'numItems' field")
		}

		// Array Field (items)
		if m.Items != nil {
			io.PushContext("items")
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(io)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
			io.PopContext("items")
		}

		io.PopContext("S7ParameterReadVarRequest")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *S7ParameterReadVarRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "items":
			arrayLoop:
				for {
					token, err = d.Token()
					switch token.(type) {
					case xml.StartElement:
						tok := token.(xml.StartElement)
						var dt *S7VarRequestParameterItem
						if err := d.DecodeElement(&dt, &tok); err != nil {
							return err
						}
						m.Items = append(m.Items, dt)
					default:
						break arrayLoop
					}
				}
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

func (m *S7ParameterReadVarRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
		return err
	}
	for _, arrayElement := range m.Items {
		if err := e.EncodeElement(arrayElement, xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "items"}}); err != nil {
		return err
	}
	return nil
}

func (m S7ParameterReadVarRequest) String() string {
	return string(m.Box("", 120))
}

func (m S7ParameterReadVarRequest) Box(name string, width int) utils.AsciiBox {
	boxName := "S7ParameterReadVarRequest"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Implicit Field (numItems)
		numItems := uint8(uint8(len(m.Items)))
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("NumItems", numItems, -1))
		// Array Field (items)
		if m.Items != nil {
			// Complex array base type
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.Items {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("Items", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
