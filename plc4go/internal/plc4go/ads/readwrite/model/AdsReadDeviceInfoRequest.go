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

// The data-structure of this message
type AdsReadDeviceInfoRequest struct {
    Parent *AdsData
    IAdsReadDeviceInfoRequest
}

// The corresponding interface
type IAdsReadDeviceInfoRequest interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsReadDeviceInfoRequest) CommandId() CommandId {
    return CommandId_ADS_READ_DEVICE_INFO
}

func (m *AdsReadDeviceInfoRequest) Response() bool {
    return false
}


func (m *AdsReadDeviceInfoRequest) InitializeParent(parent *AdsData) {
}

func NewAdsReadDeviceInfoRequest() *AdsData {
    child := &AdsReadDeviceInfoRequest{
        Parent: NewAdsData(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastAdsReadDeviceInfoRequest(structType interface{}) *AdsReadDeviceInfoRequest {
    castFunc := func(typ interface{}) *AdsReadDeviceInfoRequest {
        if casted, ok := typ.(AdsReadDeviceInfoRequest); ok {
            return &casted
        }
        if casted, ok := typ.(*AdsReadDeviceInfoRequest); ok {
            return casted
        }
        if casted, ok := typ.(AdsData); ok {
            return CastAdsReadDeviceInfoRequest(casted.Child)
        }
        if casted, ok := typ.(*AdsData); ok {
            return CastAdsReadDeviceInfoRequest(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *AdsReadDeviceInfoRequest) GetTypeName() string {
    return "AdsReadDeviceInfoRequest"
}

func (m *AdsReadDeviceInfoRequest) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *AdsReadDeviceInfoRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func AdsReadDeviceInfoRequestParse(io *utils.ReadBuffer) (*AdsData, error) {

    // Create a partially initialized instance
    _child := &AdsReadDeviceInfoRequest{
        Parent: &AdsData{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *AdsReadDeviceInfoRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsReadDeviceInfoRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            }
        }
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
    }
}

func (m *AdsReadDeviceInfoRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

