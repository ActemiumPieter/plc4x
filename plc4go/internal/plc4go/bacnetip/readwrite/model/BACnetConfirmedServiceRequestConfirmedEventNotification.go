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
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type BACnetConfirmedServiceRequestConfirmedEventNotification struct {
	Parent *BACnetConfirmedServiceRequest
}

// The corresponding interface
type IBACnetConfirmedServiceRequestConfirmedEventNotification interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) ServiceChoice() uint8 {
	return 0x02
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func NewBACnetConfirmedServiceRequestConfirmedEventNotification() *BACnetConfirmedServiceRequest {
	child := &BACnetConfirmedServiceRequestConfirmedEventNotification{
		Parent: NewBACnetConfirmedServiceRequest(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetConfirmedServiceRequestConfirmedEventNotification(structType interface{}) *BACnetConfirmedServiceRequestConfirmedEventNotification {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceRequestConfirmedEventNotification {
		if casted, ok := typ.(BACnetConfirmedServiceRequestConfirmedEventNotification); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequestConfirmedEventNotification); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestConfirmedEventNotification(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestConfirmedEventNotification(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedEventNotification"
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestConfirmedEventNotificationParse(io utils.ReadBuffer) (*BACnetConfirmedServiceRequest, error) {
	io.PullContext("BACnetConfirmedServiceRequestConfirmedEventNotification")

	io.CloseContext("BACnetConfirmedServiceRequestConfirmedEventNotification")

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestConfirmedEventNotification{
		Parent: &BACnetConfirmedServiceRequest{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("BACnetConfirmedServiceRequestConfirmedEventNotification")

		io.PopContext("BACnetConfirmedServiceRequestConfirmedEventNotification")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m BACnetConfirmedServiceRequestConfirmedEventNotification) String() string {
	return string(m.Box("", 120))
}

func (m BACnetConfirmedServiceRequestConfirmedEventNotification) Box(name string, width int) utils.AsciiBox {
	boxName := "BACnetConfirmedServiceRequestConfirmedEventNotification"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
