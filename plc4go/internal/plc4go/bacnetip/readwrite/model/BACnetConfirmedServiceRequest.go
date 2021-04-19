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
type BACnetConfirmedServiceRequest struct {
	Child IBACnetConfirmedServiceRequestChild
}

// The corresponding interface
type IBACnetConfirmedServiceRequest interface {
	ServiceChoice() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IBACnetConfirmedServiceRequestParent interface {
	SerializeParent(io utils.WriteBuffer, child IBACnetConfirmedServiceRequest, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetConfirmedServiceRequestChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *BACnetConfirmedServiceRequest)
	GetTypeName() string
	IBACnetConfirmedServiceRequest
	utils.AsciiBoxer
}

func NewBACnetConfirmedServiceRequest() *BACnetConfirmedServiceRequest {
	return &BACnetConfirmedServiceRequest{}
}

func CastBACnetConfirmedServiceRequest(structType interface{}) *BACnetConfirmedServiceRequest {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceRequest {
		if casted, ok := typ.(BACnetConfirmedServiceRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequest); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceRequest) GetTypeName() string {
	return "BACnetConfirmedServiceRequest"
}

func (m *BACnetConfirmedServiceRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequest) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BACnetConfirmedServiceRequest) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestParse(io utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {
	io.PullContext("BACnetConfirmedServiceRequest")

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice, _serviceChoiceErr := io.ReadUint8("serviceChoice", 8)
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetConfirmedServiceRequest
	var typeSwitchError error
	switch {
	case serviceChoice == 0x00: // BACnetConfirmedServiceRequestAcknowledgeAlarm
		_parent, typeSwitchError = BACnetConfirmedServiceRequestAcknowledgeAlarmParse(io)
	case serviceChoice == 0x01: // BACnetConfirmedServiceRequestConfirmedCOVNotification
		_parent, typeSwitchError = BACnetConfirmedServiceRequestConfirmedCOVNotificationParse(io, len)
	case serviceChoice == 0x02: // BACnetConfirmedServiceRequestConfirmedEventNotification
		_parent, typeSwitchError = BACnetConfirmedServiceRequestConfirmedEventNotificationParse(io)
	case serviceChoice == 0x04: // BACnetConfirmedServiceRequestGetEnrollmentSummary
		_parent, typeSwitchError = BACnetConfirmedServiceRequestGetEnrollmentSummaryParse(io)
	case serviceChoice == 0x05: // BACnetConfirmedServiceRequestSubscribeCOV
		_parent, typeSwitchError = BACnetConfirmedServiceRequestSubscribeCOVParse(io)
	case serviceChoice == 0x06: // BACnetConfirmedServiceRequestAtomicReadFile
		_parent, typeSwitchError = BACnetConfirmedServiceRequestAtomicReadFileParse(io)
	case serviceChoice == 0x07: // BACnetConfirmedServiceRequestAtomicWriteFile
		_parent, typeSwitchError = BACnetConfirmedServiceRequestAtomicWriteFileParse(io)
	case serviceChoice == 0x08: // BACnetConfirmedServiceRequestAddListElement
		_parent, typeSwitchError = BACnetConfirmedServiceRequestAddListElementParse(io)
	case serviceChoice == 0x09: // BACnetConfirmedServiceRequestRemoveListElement
		_parent, typeSwitchError = BACnetConfirmedServiceRequestRemoveListElementParse(io)
	case serviceChoice == 0x0A: // BACnetConfirmedServiceRequestCreateObject
		_parent, typeSwitchError = BACnetConfirmedServiceRequestCreateObjectParse(io)
	case serviceChoice == 0x0B: // BACnetConfirmedServiceRequestDeleteObject
		_parent, typeSwitchError = BACnetConfirmedServiceRequestDeleteObjectParse(io)
	case serviceChoice == 0x0C: // BACnetConfirmedServiceRequestReadProperty
		_parent, typeSwitchError = BACnetConfirmedServiceRequestReadPropertyParse(io)
	case serviceChoice == 0x0E: // BACnetConfirmedServiceRequestReadPropertyMultiple
		_parent, typeSwitchError = BACnetConfirmedServiceRequestReadPropertyMultipleParse(io)
	case serviceChoice == 0x0F: // BACnetConfirmedServiceRequestWriteProperty
		_parent, typeSwitchError = BACnetConfirmedServiceRequestWritePropertyParse(io, len)
	case serviceChoice == 0x10: // BACnetConfirmedServiceRequestWritePropertyMultiple
		_parent, typeSwitchError = BACnetConfirmedServiceRequestWritePropertyMultipleParse(io)
	case serviceChoice == 0x11: // BACnetConfirmedServiceRequestDeviceCommunicationControl
		_parent, typeSwitchError = BACnetConfirmedServiceRequestDeviceCommunicationControlParse(io)
	case serviceChoice == 0x12: // BACnetConfirmedServiceRequestConfirmedPrivateTransfer
		_parent, typeSwitchError = BACnetConfirmedServiceRequestConfirmedPrivateTransferParse(io)
	case serviceChoice == 0x13: // BACnetConfirmedServiceRequestConfirmedTextMessage
		_parent, typeSwitchError = BACnetConfirmedServiceRequestConfirmedTextMessageParse(io)
	case serviceChoice == 0x14: // BACnetConfirmedServiceRequestReinitializeDevice
		_parent, typeSwitchError = BACnetConfirmedServiceRequestReinitializeDeviceParse(io)
	case serviceChoice == 0x15: // BACnetConfirmedServiceRequestVTOpen
		_parent, typeSwitchError = BACnetConfirmedServiceRequestVTOpenParse(io)
	case serviceChoice == 0x16: // BACnetConfirmedServiceRequestVTClose
		_parent, typeSwitchError = BACnetConfirmedServiceRequestVTCloseParse(io)
	case serviceChoice == 0x17: // BACnetConfirmedServiceRequestVTData
		_parent, typeSwitchError = BACnetConfirmedServiceRequestVTDataParse(io)
	case serviceChoice == 0x18: // BACnetConfirmedServiceRequestRemovedAuthenticate
		_parent, typeSwitchError = BACnetConfirmedServiceRequestRemovedAuthenticateParse(io)
	case serviceChoice == 0x19: // BACnetConfirmedServiceRequestRemovedRequestKey
		_parent, typeSwitchError = BACnetConfirmedServiceRequestRemovedRequestKeyParse(io)
	case serviceChoice == 0x0D: // BACnetConfirmedServiceRequestRemovedReadPropertyConditional
		_parent, typeSwitchError = BACnetConfirmedServiceRequestRemovedReadPropertyConditionalParse(io)
	case serviceChoice == 0x1A: // BACnetConfirmedServiceRequestReadRange
		_parent, typeSwitchError = BACnetConfirmedServiceRequestReadRangeParse(io)
	case serviceChoice == 0x1B: // BACnetConfirmedServiceRequestLifeSafetyOperation
		_parent, typeSwitchError = BACnetConfirmedServiceRequestLifeSafetyOperationParse(io)
	case serviceChoice == 0x1C: // BACnetConfirmedServiceRequestSubscribeCOVProperty
		_parent, typeSwitchError = BACnetConfirmedServiceRequestSubscribeCOVPropertyParse(io)
	case serviceChoice == 0x1D: // BACnetConfirmedServiceRequestGetEventInformation
		_parent, typeSwitchError = BACnetConfirmedServiceRequestGetEventInformationParse(io)
	case serviceChoice == 0x1E: // BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
		_parent, typeSwitchError = BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleParse(io)
	case serviceChoice == 0x1F: // BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
		_parent, typeSwitchError = BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleParse(io)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	io.CloseContext("BACnetConfirmedServiceRequest")

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *BACnetConfirmedServiceRequest) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *BACnetConfirmedServiceRequest) SerializeParent(io utils.WriteBuffer, child IBACnetConfirmedServiceRequest, serializeChildFunction func() error) error {
	io.PushContext("BACnetConfirmedServiceRequest")

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

	io.PopContext("BACnetConfirmedServiceRequest")
	return nil
}

func (m *BACnetConfirmedServiceRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	if start.Attr != nil && len(start.Attr) > 0 {
		switch start.Attr[0].Value {
		// BACnetConfirmedServiceRequestAcknowledgeAlarm needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestAcknowledgeAlarm":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestAcknowledgeAlarm{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestConfirmedEventNotification needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestConfirmedEventNotification":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestConfirmedEventNotification{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestGetEnrollmentSummary needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestGetEnrollmentSummary":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestGetEnrollmentSummary{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestAtomicReadFile needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestAtomicReadFile":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestAtomicReadFile{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestAtomicWriteFile needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestAtomicWriteFile":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestAtomicWriteFile{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestAddListElement needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestAddListElement":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestAddListElement{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestRemoveListElement needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestRemoveListElement":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestRemoveListElement{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestCreateObject needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestCreateObject":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestCreateObject{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestDeleteObject needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestDeleteObject":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestDeleteObject{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestReadPropertyMultiple needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestReadPropertyMultiple":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestReadPropertyMultiple{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestWritePropertyMultiple needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestWritePropertyMultiple":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestWritePropertyMultiple{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestDeviceCommunicationControl needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestDeviceCommunicationControl":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestDeviceCommunicationControl{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestConfirmedPrivateTransfer needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestConfirmedPrivateTransfer":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestConfirmedPrivateTransfer{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestConfirmedTextMessage needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestConfirmedTextMessage":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestConfirmedTextMessage{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestReinitializeDevice needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestReinitializeDevice":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestReinitializeDevice{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestVTOpen needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestVTOpen":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestVTOpen{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestVTClose needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestVTClose":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestVTClose{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestVTData needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestVTData":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestVTData{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestRemovedAuthenticate needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestRemovedAuthenticate":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestRemovedAuthenticate{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestRemovedRequestKey needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestRemovedRequestKey":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestRemovedRequestKey{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestRemovedReadPropertyConditional needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestRemovedReadPropertyConditional":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestRemovedReadPropertyConditional{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestReadRange needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestReadRange":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestReadRange{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestLifeSafetyOperation needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestLifeSafetyOperation":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestLifeSafetyOperation{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestSubscribeCOVProperty needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestSubscribeCOVProperty":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestSubscribeCOVProperty{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestGetEventInformation needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestGetEventInformation":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestGetEventInformation{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple{
					Parent: m,
				}
			}
		// BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple":
			if m.Child == nil {
				m.Child = &BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple{
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
					panic("Couldn't determine class type for childs of BACnetConfirmedServiceRequest")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestAcknowledgeAlarm":
					var dt *BACnetConfirmedServiceRequestAcknowledgeAlarm
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestAcknowledgeAlarm)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestConfirmedCOVNotification":
					var dt *BACnetConfirmedServiceRequestConfirmedCOVNotification
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestConfirmedCOVNotification)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestConfirmedEventNotification":
					var dt *BACnetConfirmedServiceRequestConfirmedEventNotification
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestConfirmedEventNotification)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestGetEnrollmentSummary":
					var dt *BACnetConfirmedServiceRequestGetEnrollmentSummary
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestGetEnrollmentSummary)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestSubscribeCOV":
					var dt *BACnetConfirmedServiceRequestSubscribeCOV
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestSubscribeCOV)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestAtomicReadFile":
					var dt *BACnetConfirmedServiceRequestAtomicReadFile
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestAtomicReadFile)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestAtomicWriteFile":
					var dt *BACnetConfirmedServiceRequestAtomicWriteFile
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestAtomicWriteFile)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestAddListElement":
					var dt *BACnetConfirmedServiceRequestAddListElement
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestAddListElement)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestRemoveListElement":
					var dt *BACnetConfirmedServiceRequestRemoveListElement
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestRemoveListElement)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestCreateObject":
					var dt *BACnetConfirmedServiceRequestCreateObject
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestCreateObject)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestDeleteObject":
					var dt *BACnetConfirmedServiceRequestDeleteObject
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestDeleteObject)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestReadProperty":
					var dt *BACnetConfirmedServiceRequestReadProperty
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestReadProperty)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestReadPropertyMultiple":
					var dt *BACnetConfirmedServiceRequestReadPropertyMultiple
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestReadPropertyMultiple)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestWriteProperty":
					var dt *BACnetConfirmedServiceRequestWriteProperty
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestWriteProperty)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestWritePropertyMultiple":
					var dt *BACnetConfirmedServiceRequestWritePropertyMultiple
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestWritePropertyMultiple)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestDeviceCommunicationControl":
					var dt *BACnetConfirmedServiceRequestDeviceCommunicationControl
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestDeviceCommunicationControl)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestConfirmedPrivateTransfer":
					var dt *BACnetConfirmedServiceRequestConfirmedPrivateTransfer
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestConfirmedPrivateTransfer)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestConfirmedTextMessage":
					var dt *BACnetConfirmedServiceRequestConfirmedTextMessage
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestConfirmedTextMessage)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestReinitializeDevice":
					var dt *BACnetConfirmedServiceRequestReinitializeDevice
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestReinitializeDevice)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestVTOpen":
					var dt *BACnetConfirmedServiceRequestVTOpen
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestVTOpen)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestVTClose":
					var dt *BACnetConfirmedServiceRequestVTClose
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestVTClose)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestVTData":
					var dt *BACnetConfirmedServiceRequestVTData
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestVTData)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestRemovedAuthenticate":
					var dt *BACnetConfirmedServiceRequestRemovedAuthenticate
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestRemovedAuthenticate)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestRemovedRequestKey":
					var dt *BACnetConfirmedServiceRequestRemovedRequestKey
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestRemovedRequestKey)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestRemovedReadPropertyConditional":
					var dt *BACnetConfirmedServiceRequestRemovedReadPropertyConditional
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestRemovedReadPropertyConditional)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestReadRange":
					var dt *BACnetConfirmedServiceRequestReadRange
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestReadRange)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestLifeSafetyOperation":
					var dt *BACnetConfirmedServiceRequestLifeSafetyOperation
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestLifeSafetyOperation)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestSubscribeCOVProperty":
					var dt *BACnetConfirmedServiceRequestSubscribeCOVProperty
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestSubscribeCOVProperty)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestGetEventInformation":
					var dt *BACnetConfirmedServiceRequestGetEventInformation
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestGetEventInformation)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple":
					var dt *BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple":
					var dt *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
					if m.Child != nil {
						dt = m.Child.(*BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple)
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

func (m *BACnetConfirmedServiceRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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

func (m BACnetConfirmedServiceRequest) String() string {
	return string(m.Box("", 120))
}

func (m *BACnetConfirmedServiceRequest) Box(name string, width int) utils.AsciiBox {
	return m.Child.Box(name, width)
}

func (m *BACnetConfirmedServiceRequest) BoxParent(name string, width int, childBoxer func() []utils.AsciiBox) utils.AsciiBox {
	boxName := "BACnetConfirmedServiceRequest"
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
