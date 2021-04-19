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
	"reflect"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type BACnetUnconfirmedServiceRequest struct {
	Child IBACnetUnconfirmedServiceRequestChild
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequest interface {
	ServiceChoice() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IBACnetUnconfirmedServiceRequestParent interface {
	SerializeParent(io utils.WriteBuffer, child IBACnetUnconfirmedServiceRequest, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetUnconfirmedServiceRequestChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *BACnetUnconfirmedServiceRequest)
	GetTypeName() string
	IBACnetUnconfirmedServiceRequest
	utils.AsciiBoxer
}

func NewBACnetUnconfirmedServiceRequest() *BACnetUnconfirmedServiceRequest {
	return &BACnetUnconfirmedServiceRequest{}
}

func CastBACnetUnconfirmedServiceRequest(structType interface{}) *BACnetUnconfirmedServiceRequest {
	castFunc := func(typ interface{}) *BACnetUnconfirmedServiceRequest {
		if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequest) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequest"
}

func (m *BACnetUnconfirmedServiceRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequest) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BACnetUnconfirmedServiceRequest) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestParse(io utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequest, error) {
	io.PullContext("BACnetUnconfirmedServiceRequest")

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice, _serviceChoiceErr := io.ReadUint8("serviceChoice", 8)
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetUnconfirmedServiceRequest
	var typeSwitchError error
	switch {
	case serviceChoice == 0x00: // BACnetUnconfirmedServiceRequestIAm
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestIAmParse(io)
	case serviceChoice == 0x01: // BACnetUnconfirmedServiceRequestIHave
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestIHaveParse(io)
	case serviceChoice == 0x02: // BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationParse(io)
	case serviceChoice == 0x03: // BACnetUnconfirmedServiceRequestUnconfirmedEventNotification
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationParse(io)
	case serviceChoice == 0x04: // BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransferParse(io, len)
	case serviceChoice == 0x05: // BACnetUnconfirmedServiceRequestUnconfirmedTextMessage
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedTextMessageParse(io)
	case serviceChoice == 0x06: // BACnetUnconfirmedServiceRequestTimeSynchronization
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestTimeSynchronizationParse(io)
	case serviceChoice == 0x07: // BACnetUnconfirmedServiceRequestWhoHas
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestWhoHasParse(io)
	case serviceChoice == 0x08: // BACnetUnconfirmedServiceRequestWhoIs
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestWhoIsParse(io)
	case serviceChoice == 0x09: // BACnetUnconfirmedServiceRequestUTCTimeSynchronization
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestUTCTimeSynchronizationParse(io)
	case serviceChoice == 0x0A: // BACnetUnconfirmedServiceRequestWriteGroup
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestWriteGroupParse(io)
	case serviceChoice == 0x0B: // BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
		_parent, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleParse(io)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	io.CloseContext("BACnetUnconfirmedServiceRequest")

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *BACnetUnconfirmedServiceRequest) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *BACnetUnconfirmedServiceRequest) SerializeParent(io utils.WriteBuffer, child IBACnetUnconfirmedServiceRequest, serializeChildFunction func() error) error {
	io.PushContext("BACnetUnconfirmedServiceRequest")

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice := uint8(child.ServiceChoice())
	_serviceChoiceErr := io.WriteUint8("serviceChoice", 8, (serviceChoice))

	if _serviceChoiceErr != nil {
		return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	io.PopContext("BACnetUnconfirmedServiceRequest")
	return nil
}

func (m *BACnetUnconfirmedServiceRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	if start.Attr != nil && len(start.Attr) > 0 {
		switch start.Attr[0].Value {
		// BACnetUnconfirmedServiceRequestIHave needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestIHave":
			if m.Child == nil {
				m.Child = &BACnetUnconfirmedServiceRequestIHave{
					Parent: m,
				}
			}
		// BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification":
			if m.Child == nil {
				m.Child = &BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification{
					Parent: m,
				}
			}
		// BACnetUnconfirmedServiceRequestUnconfirmedEventNotification needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUnconfirmedEventNotification":
			if m.Child == nil {
				m.Child = &BACnetUnconfirmedServiceRequestUnconfirmedEventNotification{
					Parent: m,
				}
			}
		// BACnetUnconfirmedServiceRequestUnconfirmedTextMessage needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUnconfirmedTextMessage":
			if m.Child == nil {
				m.Child = &BACnetUnconfirmedServiceRequestUnconfirmedTextMessage{
					Parent: m,
				}
			}
		// BACnetUnconfirmedServiceRequestTimeSynchronization needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestTimeSynchronization":
			if m.Child == nil {
				m.Child = &BACnetUnconfirmedServiceRequestTimeSynchronization{
					Parent: m,
				}
			}
		// BACnetUnconfirmedServiceRequestUTCTimeSynchronization needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUTCTimeSynchronization":
			if m.Child == nil {
				m.Child = &BACnetUnconfirmedServiceRequestUTCTimeSynchronization{
					Parent: m,
				}
			}
		// BACnetUnconfirmedServiceRequestWriteGroup needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestWriteGroup":
			if m.Child == nil {
				m.Child = &BACnetUnconfirmedServiceRequestWriteGroup{
					Parent: m,
				}
			}
		// BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple":
			if m.Child == nil {
				m.Child = &BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple{
					Parent: m,
				}
			}
		}
	}
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			default:
				attr := start.Attr
				if attr == nil || len(attr) <= 0 {
					// TODO: workaround for bug with nested lists
					attr = tok.Attr
				}
				if attr == nil || len(attr) <= 0 {
					panic("Couldn't determine class type for childs of BACnetUnconfirmedServiceRequest")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestIAm":
					var dt *BACnetUnconfirmedServiceRequestIAm
					if m.Child != nil {
						dt = m.Child.(*BACnetUnconfirmedServiceRequestIAm)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestIHave":
					var dt *BACnetUnconfirmedServiceRequestIHave
					if m.Child != nil {
						dt = m.Child.(*BACnetUnconfirmedServiceRequestIHave)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification":
					var dt *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification
					if m.Child != nil {
						dt = m.Child.(*BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUnconfirmedEventNotification":
					var dt *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification
					if m.Child != nil {
						dt = m.Child.(*BACnetUnconfirmedServiceRequestUnconfirmedEventNotification)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer":
					var dt *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
					if m.Child != nil {
						dt = m.Child.(*BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUnconfirmedTextMessage":
					var dt *BACnetUnconfirmedServiceRequestUnconfirmedTextMessage
					if m.Child != nil {
						dt = m.Child.(*BACnetUnconfirmedServiceRequestUnconfirmedTextMessage)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestTimeSynchronization":
					var dt *BACnetUnconfirmedServiceRequestTimeSynchronization
					if m.Child != nil {
						dt = m.Child.(*BACnetUnconfirmedServiceRequestTimeSynchronization)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestWhoHas":
					var dt *BACnetUnconfirmedServiceRequestWhoHas
					if m.Child != nil {
						dt = m.Child.(*BACnetUnconfirmedServiceRequestWhoHas)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestWhoIs":
					var dt *BACnetUnconfirmedServiceRequestWhoIs
					if m.Child != nil {
						dt = m.Child.(*BACnetUnconfirmedServiceRequestWhoIs)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUTCTimeSynchronization":
					var dt *BACnetUnconfirmedServiceRequestUTCTimeSynchronization
					if m.Child != nil {
						dt = m.Child.(*BACnetUnconfirmedServiceRequestUTCTimeSynchronization)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestWriteGroup":
					var dt *BACnetUnconfirmedServiceRequestWriteGroup
					if m.Child != nil {
						dt = m.Child.(*BACnetUnconfirmedServiceRequestWriteGroup)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple":
					var dt *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
					if m.Child != nil {
						dt = m.Child.(*BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				}
			}
		}
	}
}

func (m *BACnetUnconfirmedServiceRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.bacnetip.readwrite." + className[strings.LastIndex(className, ".")+1:]
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	marshaller, ok := m.Child.(xml.Marshaler)
	if !ok {
		return errors.Errorf("child is not castable to Marshaler. Actual type %T", m.Child)
	}
	if err := marshaller.MarshalXML(e, start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m BACnetUnconfirmedServiceRequest) String() string {
	return string(m.Box("", 120))
}

func (m *BACnetUnconfirmedServiceRequest) Box(name string, width int) utils.AsciiBox {
	return m.Child.Box(name, width)
}

func (m *BACnetUnconfirmedServiceRequest) BoxParent(name string, width int, childBoxer func() []utils.AsciiBox) utils.AsciiBox {
	boxName := "BACnetUnconfirmedServiceRequest"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice := uint8(m.Child.ServiceChoice())
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("ServiceChoice", serviceChoice, -1))
	// Switch field (Depending on the discriminator values, passes the boxing to a sub-type)
	boxes = append(boxes, childBoxer()...)
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
