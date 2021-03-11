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
package readwrite

import (
    "errors"
    "github.com/apache/plc4x/plc4go/internal/plc4go/ads/readwrite/model"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

type AdsParserHelper struct {
}

func (m AdsParserHelper) Parse(typeName string, arguments []string, io *utils.ReadBuffer) (spi.Message, error) {
    switch typeName {
    case "AdsMultiRequestItem":
        indexGroup, err := utils.StrToUint32(arguments[0])
        if err != nil {
            return nil, err
        }
        return model.AdsMultiRequestItemParse(io, indexGroup)
    case "AmsTCPPacket":
        return model.AmsTCPPacketParse(io)
    case "State":
        return model.StateParse(io)
    case "AmsPacket":
        return model.AmsPacketParse(io)
    case "AmsSerialFrame":
        return model.AmsSerialFrameParse(io)
    case "AmsSerialAcknowledgeFrame":
        return model.AmsSerialAcknowledgeFrameParse(io)
    case "AdsData":
        var commandId ICommandId
        response, err := utils.StrToBool(arguments[1])
        if err != nil {
            return nil, err
        }
        return model.AdsDataParse(io, commandId, response)
    case "AmsNetId":
        return model.AmsNetIdParse(io)
    case "AdsStampHeader":
        return model.AdsStampHeaderParse(io)
    case "AmsSerialResetFrame":
        return model.AmsSerialResetFrameParse(io)
    case "AdsNotificationSample":
        return model.AdsNotificationSampleParse(io)
    }
    return nil, errors.New("Unsupported type " + typeName)
}
