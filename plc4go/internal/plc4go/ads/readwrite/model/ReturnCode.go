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
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

type ReturnCode uint32

type IReturnCode interface {
    Serialize(io utils.WriteBuffer) error
}

const(
    ReturnCode_OK ReturnCode = 0x00
    ReturnCode_INTERNAL_ERROR ReturnCode = 0x01
    ReturnCode_NO_REALTIME ReturnCode = 0x02
    ReturnCode_SAVE_ERROR ReturnCode = 0x03
    ReturnCode_MAILBOX_FULL ReturnCode = 0x04
    ReturnCode_WRONG_HMSG ReturnCode = 0x05
    ReturnCode_TARGET_PORT_NOT_FOUND ReturnCode = 0x06
    ReturnCode_TARGET_HOST_NOT_FOUND ReturnCode = 0x07
    ReturnCode_UNKNOWN_COMMAND_ID ReturnCode = 0x08
    ReturnCode_UNKNOWN_TASK_ID ReturnCode = 0x09
    ReturnCode_NO_IO ReturnCode = 0x0A
    ReturnCode_UNKNOWN_ADS_COMMAND ReturnCode = 0x0B
    ReturnCode_WIN32_ERROR ReturnCode = 0x0C
    ReturnCode_PORT_NOT_CONNECTED ReturnCode = 0x0D
    ReturnCode_INVALID_ADS_LENGTH ReturnCode = 0x0E
    ReturnCode_INVALID_AMS_NET_ID ReturnCode = 0x0F
    ReturnCode_LOW_INSTALLATION_LEVEL ReturnCode = 0x10
    ReturnCode_NO_DEBUGGING_AVAILABLE ReturnCode = 0x11
    ReturnCode_PORT_DEACTIVATED ReturnCode = 0x12
    ReturnCode_PORT_ALREADY_CONNECTED ReturnCode = 0x13
    ReturnCode_ADS_SYNC_WIN32_ERROR ReturnCode = 0x14
    ReturnCode_ADS_SYNC_TIMEOUT ReturnCode = 0x15
    ReturnCode_ADS_SYNC_AMS_ERROR ReturnCode = 0x16
    ReturnCode_NO_INDEX_MAP_FOR_ADS_AVAILABLE ReturnCode = 0x17
    ReturnCode_INVALID_ADS_PORT ReturnCode = 0x18
    ReturnCode_NO_MEMORY ReturnCode = 0x19
    ReturnCode_TCP_SENDING_ERROR ReturnCode = 0x1A
    ReturnCode_HOST_NOT_REACHABLE ReturnCode = 0x1B
    ReturnCode_INVALID_AMS_FRAGMENT ReturnCode = 0x1C
    ReturnCode_ROUTERERR_NOLOCKEDMEMORY ReturnCode = 0x500
    ReturnCode_ROUTERERR_RESIZEMEMORY ReturnCode = 0x501
    ReturnCode_ROUTERERR_MAILBOXFULL ReturnCode = 0x502
    ReturnCode_ROUTERERR_DEBUGBOXFULL ReturnCode = 0x503
    ReturnCode_ROUTERERR_UNKNOWNPORTTYPE ReturnCode = 0x504
    ReturnCode_ROUTERERR_NOTINITIALIZED ReturnCode = 0x505
    ReturnCode_ROUTERERR_PORTALREADYINUSE ReturnCode = 0x506
    ReturnCode_ROUTERERR_NOTREGISTERED ReturnCode = 0x507
    ReturnCode_ROUTERERR_NOMOREQUEUES ReturnCode = 0x508
    ReturnCode_ROUTERERR_INVALIDPORT ReturnCode = 0x509
    ReturnCode_ROUTERERR_NOTACTIVATED ReturnCode = 0x50A
    ReturnCode_ADSERR_DEVICE_ERROR ReturnCode = 0x700
    ReturnCode_ADSERR_DEVICE_SRVNOTSUPP ReturnCode = 0x701
    ReturnCode_ADSERR_DEVICE_INVALIDGRP ReturnCode = 0x702
    ReturnCode_ADSERR_DEVICE_INVALIDOFFSET ReturnCode = 0x703
    ReturnCode_ADSERR_DEVICE_INVALIDACCESS ReturnCode = 0x704
    ReturnCode_ADSERR_DEVICE_INVALIDSIZE ReturnCode = 0x705
    ReturnCode_ADSERR_DEVICE_INVALIDDATA ReturnCode = 0x706
    ReturnCode_ADSERR_DEVICE_NOTREADY ReturnCode = 0x707
    ReturnCode_ADSERR_DEVICE_BUSY ReturnCode = 0x708
    ReturnCode_ADSERR_DEVICE_INVALIDCONTEXT ReturnCode = 0x709
    ReturnCode_ADSERR_DEVICE_NOMEMORY ReturnCode = 0x70A
    ReturnCode_ADSERR_DEVICE_INVALIDPARM ReturnCode = 0x70B
    ReturnCode_ADSERR_DEVICE_NOTFOUND ReturnCode = 0x70C
    ReturnCode_ADSERR_DEVICE_SYNTAX ReturnCode = 0x70D
    ReturnCode_ADSERR_DEVICE_INCOMPATIBLE ReturnCode = 0x70E
    ReturnCode_ADSERR_DEVICE_EXISTS ReturnCode = 0x70F
    ReturnCode_ADSERR_DEVICE_SYMBOLNOTFOUND ReturnCode = 0x710
    ReturnCode_ADSERR_DEVICE_SYMBOLVERSIONINVALID ReturnCode = 0x711
    ReturnCode_ADSERR_DEVICE_INVALIDSTATE ReturnCode = 0x712
    ReturnCode_ADSERR_DEVICE_TRANSMODENOTSUPP ReturnCode = 0x713
    ReturnCode_ADSERR_DEVICE_NOTIFYHNDINVALID ReturnCode = 0x714
    ReturnCode_ADSERR_DEVICE_CLIENTUNKNOWN ReturnCode = 0x715
    ReturnCode_ADSERR_DEVICE_NOMOREHDLS ReturnCode = 0x716
    ReturnCode_ADSERR_DEVICE_INVALIDWATCHSIZE ReturnCode = 0x717
    ReturnCode_ADSERR_DEVICE_NOTINIT ReturnCode = 0x718
    ReturnCode_ADSERR_DEVICE_TIMEOUT ReturnCode = 0x719
    ReturnCode_ADSERR_DEVICE_NOINTERFACE ReturnCode = 0x71A
    ReturnCode_ADSERR_DEVICE_INVALIDINTERFACE ReturnCode = 0x71B
    ReturnCode_ADSERR_DEVICE_INVALIDCLSID ReturnCode = 0x71C
    ReturnCode_ADSERR_DEVICE_INVALIDOBJID ReturnCode = 0x71D
    ReturnCode_ADSERR_DEVICE_PENDING ReturnCode = 0x71E
    ReturnCode_ADSERR_DEVICE_ABORTED ReturnCode = 0x71F
    ReturnCode_ADSERR_DEVICE_WARNING ReturnCode = 0x720
    ReturnCode_ADSERR_DEVICE_INVALIDARRAYIDX ReturnCode = 0x721
    ReturnCode_ADSERR_DEVICE_SYMBOLNOTACTIVE ReturnCode = 0x722
    ReturnCode_ADSERR_DEVICE_ACCESSDENIED ReturnCode = 0x723
    ReturnCode_ADSERR_DEVICE_LICENSENOTFOUND ReturnCode = 0x724
    ReturnCode_ADSERR_DEVICE_LICENSEEXPIRED ReturnCode = 0x725
    ReturnCode_ADSERR_DEVICE_LICENSEEXCEEDED ReturnCode = 0x726
    ReturnCode_ADSERR_DEVICE_LICENSEINVALID ReturnCode = 0x727
    ReturnCode_ADSERR_DEVICE_LICENSESYSTEMID ReturnCode = 0x728
    ReturnCode_ADSERR_DEVICE_LICENSENOTIMELIMIT ReturnCode = 0x729
    ReturnCode_ADSERR_DEVICE_LICENSEFUTUREISSUE ReturnCode = 0x72A
    ReturnCode_ADSERR_DEVICE_LICENSETIMETOLONG ReturnCode = 0x72B
    ReturnCode_ADSERR_DEVICE_EXCEPTION ReturnCode = 0x72c
    ReturnCode_ADSERR_DEVICE_LICENSEDUPLICATED ReturnCode = 0x72D
    ReturnCode_ADSERR_DEVICE_SIGNATUREINVALID ReturnCode = 0x72E
    ReturnCode_ADSERR_DEVICE_CERTIFICATEINVALID ReturnCode = 0x72F
    ReturnCode_ADSERR_CLIENT_ERROR ReturnCode = 0x740
    ReturnCode_ADSERR_CLIENT_INVALIDPARM ReturnCode = 0x741
    ReturnCode_ADSERR_CLIENT_LISTEMPTY ReturnCode = 0x742
    ReturnCode_ADSERR_CLIENT_VARUSED ReturnCode = 0x743
    ReturnCode_ADSERR_CLIENT_DUPLINVOKEID ReturnCode = 0x744
    ReturnCode_ADSERR_CLIENT_SYNCTIMEOUT ReturnCode = 0x745
    ReturnCode_ADSERR_CLIENT_W32ERROR ReturnCode = 0x746
    ReturnCode_ADSERR_CLIENT_TIMEOUTINVALID ReturnCode = 0x747
    ReturnCode_ADSERR_CLIENT_PORTNOTOPEN ReturnCode = 0x748
    ReturnCode_ADSERR_CLIENT_NOAMSADDR ReturnCode = 0x750
    ReturnCode_ADSERR_CLIENT_SYNCINTERNAL ReturnCode = 0x751
    ReturnCode_ADSERR_CLIENT_ADDHASH ReturnCode = 0x752
    ReturnCode_ADSERR_CLIENT_REMOVEHASH ReturnCode = 0x753
    ReturnCode_ADSERR_CLIENT_NOMORESYM ReturnCode = 0x754
    ReturnCode_ADSERR_CLIENT_SYNCRESINVALID ReturnCode = 0x755
    ReturnCode_RTERR_INTERNAL ReturnCode = 0x1000
    ReturnCode_RTERR_BADTIMERPERIODS ReturnCode = 0x1001
    ReturnCode_RTERR_INVALIDTASKPTR ReturnCode = 0x1002
    ReturnCode_RTERR_INVALIDSTACKPTR ReturnCode = 0x1003
    ReturnCode_RTERR_PRIOEXISTS ReturnCode = 0x1004
    ReturnCode_RTERR_NOMORETCB ReturnCode = 0x1005
    ReturnCode_RTERR_NOMORESEMAS ReturnCode = 0x1006
    ReturnCode_RTERR_NOMOREQUEUES ReturnCode = 0x1007
    ReturnCode_RTERR_EXTIRQALREADYDEF ReturnCode = 0x100D
    ReturnCode_RTERR_EXTIRQNOTDEF ReturnCode = 0x100E
    ReturnCode_RTERR_EXTIRQINSTALLFAILED ReturnCode = 0x100F
    ReturnCode_RTERR_IRQLNOTLESSOREQUAL ReturnCode = 0x1010
    ReturnCode_RTERR_VMXNOTSUPPORTED ReturnCode = 0x1017
    ReturnCode_RTERR_VMXDISABLED ReturnCode = 0x1018
    ReturnCode_RTERR_VMXCONTROLSMISSING ReturnCode = 0x1019
    ReturnCode_RTERR_VMXENABLEFAILS ReturnCode = 0x101A
    ReturnCode_WSAETIMEDOUT ReturnCode = 0x274C
    ReturnCode_WSAECONNREFUSED ReturnCode = 0x274D
    ReturnCode_WSAEHOSTUNREACH ReturnCode = 0x2751
)

func ReturnCodeByValue(value uint32) ReturnCode {
    switch value {
        case 0x00:
            return ReturnCode_OK
        case 0x01:
            return ReturnCode_INTERNAL_ERROR
        case 0x02:
            return ReturnCode_NO_REALTIME
        case 0x03:
            return ReturnCode_SAVE_ERROR
        case 0x04:
            return ReturnCode_MAILBOX_FULL
        case 0x05:
            return ReturnCode_WRONG_HMSG
        case 0x06:
            return ReturnCode_TARGET_PORT_NOT_FOUND
        case 0x07:
            return ReturnCode_TARGET_HOST_NOT_FOUND
        case 0x08:
            return ReturnCode_UNKNOWN_COMMAND_ID
        case 0x09:
            return ReturnCode_UNKNOWN_TASK_ID
        case 0x0A:
            return ReturnCode_NO_IO
        case 0x0B:
            return ReturnCode_UNKNOWN_ADS_COMMAND
        case 0x0C:
            return ReturnCode_WIN32_ERROR
        case 0x0D:
            return ReturnCode_PORT_NOT_CONNECTED
        case 0x0E:
            return ReturnCode_INVALID_ADS_LENGTH
        case 0x0F:
            return ReturnCode_INVALID_AMS_NET_ID
        case 0x10:
            return ReturnCode_LOW_INSTALLATION_LEVEL
        case 0x1000:
            return ReturnCode_RTERR_INTERNAL
        case 0x1001:
            return ReturnCode_RTERR_BADTIMERPERIODS
        case 0x1002:
            return ReturnCode_RTERR_INVALIDTASKPTR
        case 0x1003:
            return ReturnCode_RTERR_INVALIDSTACKPTR
        case 0x1004:
            return ReturnCode_RTERR_PRIOEXISTS
        case 0x1005:
            return ReturnCode_RTERR_NOMORETCB
        case 0x1006:
            return ReturnCode_RTERR_NOMORESEMAS
        case 0x1007:
            return ReturnCode_RTERR_NOMOREQUEUES
        case 0x100D:
            return ReturnCode_RTERR_EXTIRQALREADYDEF
        case 0x100E:
            return ReturnCode_RTERR_EXTIRQNOTDEF
        case 0x100F:
            return ReturnCode_RTERR_EXTIRQINSTALLFAILED
        case 0x1010:
            return ReturnCode_RTERR_IRQLNOTLESSOREQUAL
        case 0x1017:
            return ReturnCode_RTERR_VMXNOTSUPPORTED
        case 0x1018:
            return ReturnCode_RTERR_VMXDISABLED
        case 0x1019:
            return ReturnCode_RTERR_VMXCONTROLSMISSING
        case 0x101A:
            return ReturnCode_RTERR_VMXENABLEFAILS
        case 0x11:
            return ReturnCode_NO_DEBUGGING_AVAILABLE
        case 0x12:
            return ReturnCode_PORT_DEACTIVATED
        case 0x13:
            return ReturnCode_PORT_ALREADY_CONNECTED
        case 0x14:
            return ReturnCode_ADS_SYNC_WIN32_ERROR
        case 0x15:
            return ReturnCode_ADS_SYNC_TIMEOUT
        case 0x16:
            return ReturnCode_ADS_SYNC_AMS_ERROR
        case 0x17:
            return ReturnCode_NO_INDEX_MAP_FOR_ADS_AVAILABLE
        case 0x18:
            return ReturnCode_INVALID_ADS_PORT
        case 0x19:
            return ReturnCode_NO_MEMORY
        case 0x1A:
            return ReturnCode_TCP_SENDING_ERROR
        case 0x1B:
            return ReturnCode_HOST_NOT_REACHABLE
        case 0x1C:
            return ReturnCode_INVALID_AMS_FRAGMENT
        case 0x274C:
            return ReturnCode_WSAETIMEDOUT
        case 0x274D:
            return ReturnCode_WSAECONNREFUSED
        case 0x2751:
            return ReturnCode_WSAEHOSTUNREACH
        case 0x500:
            return ReturnCode_ROUTERERR_NOLOCKEDMEMORY
        case 0x501:
            return ReturnCode_ROUTERERR_RESIZEMEMORY
        case 0x502:
            return ReturnCode_ROUTERERR_MAILBOXFULL
        case 0x503:
            return ReturnCode_ROUTERERR_DEBUGBOXFULL
        case 0x504:
            return ReturnCode_ROUTERERR_UNKNOWNPORTTYPE
        case 0x505:
            return ReturnCode_ROUTERERR_NOTINITIALIZED
        case 0x506:
            return ReturnCode_ROUTERERR_PORTALREADYINUSE
        case 0x507:
            return ReturnCode_ROUTERERR_NOTREGISTERED
        case 0x508:
            return ReturnCode_ROUTERERR_NOMOREQUEUES
        case 0x509:
            return ReturnCode_ROUTERERR_INVALIDPORT
        case 0x50A:
            return ReturnCode_ROUTERERR_NOTACTIVATED
        case 0x700:
            return ReturnCode_ADSERR_DEVICE_ERROR
        case 0x701:
            return ReturnCode_ADSERR_DEVICE_SRVNOTSUPP
        case 0x702:
            return ReturnCode_ADSERR_DEVICE_INVALIDGRP
        case 0x703:
            return ReturnCode_ADSERR_DEVICE_INVALIDOFFSET
        case 0x704:
            return ReturnCode_ADSERR_DEVICE_INVALIDACCESS
        case 0x705:
            return ReturnCode_ADSERR_DEVICE_INVALIDSIZE
        case 0x706:
            return ReturnCode_ADSERR_DEVICE_INVALIDDATA
        case 0x707:
            return ReturnCode_ADSERR_DEVICE_NOTREADY
        case 0x708:
            return ReturnCode_ADSERR_DEVICE_BUSY
        case 0x709:
            return ReturnCode_ADSERR_DEVICE_INVALIDCONTEXT
        case 0x70A:
            return ReturnCode_ADSERR_DEVICE_NOMEMORY
        case 0x70B:
            return ReturnCode_ADSERR_DEVICE_INVALIDPARM
        case 0x70C:
            return ReturnCode_ADSERR_DEVICE_NOTFOUND
        case 0x70D:
            return ReturnCode_ADSERR_DEVICE_SYNTAX
        case 0x70E:
            return ReturnCode_ADSERR_DEVICE_INCOMPATIBLE
        case 0x70F:
            return ReturnCode_ADSERR_DEVICE_EXISTS
        case 0x710:
            return ReturnCode_ADSERR_DEVICE_SYMBOLNOTFOUND
        case 0x711:
            return ReturnCode_ADSERR_DEVICE_SYMBOLVERSIONINVALID
        case 0x712:
            return ReturnCode_ADSERR_DEVICE_INVALIDSTATE
        case 0x713:
            return ReturnCode_ADSERR_DEVICE_TRANSMODENOTSUPP
        case 0x714:
            return ReturnCode_ADSERR_DEVICE_NOTIFYHNDINVALID
        case 0x715:
            return ReturnCode_ADSERR_DEVICE_CLIENTUNKNOWN
        case 0x716:
            return ReturnCode_ADSERR_DEVICE_NOMOREHDLS
        case 0x717:
            return ReturnCode_ADSERR_DEVICE_INVALIDWATCHSIZE
        case 0x718:
            return ReturnCode_ADSERR_DEVICE_NOTINIT
        case 0x719:
            return ReturnCode_ADSERR_DEVICE_TIMEOUT
        case 0x71A:
            return ReturnCode_ADSERR_DEVICE_NOINTERFACE
        case 0x71B:
            return ReturnCode_ADSERR_DEVICE_INVALIDINTERFACE
        case 0x71C:
            return ReturnCode_ADSERR_DEVICE_INVALIDCLSID
        case 0x71D:
            return ReturnCode_ADSERR_DEVICE_INVALIDOBJID
        case 0x71E:
            return ReturnCode_ADSERR_DEVICE_PENDING
        case 0x71F:
            return ReturnCode_ADSERR_DEVICE_ABORTED
        case 0x720:
            return ReturnCode_ADSERR_DEVICE_WARNING
        case 0x721:
            return ReturnCode_ADSERR_DEVICE_INVALIDARRAYIDX
        case 0x722:
            return ReturnCode_ADSERR_DEVICE_SYMBOLNOTACTIVE
        case 0x723:
            return ReturnCode_ADSERR_DEVICE_ACCESSDENIED
        case 0x724:
            return ReturnCode_ADSERR_DEVICE_LICENSENOTFOUND
        case 0x725:
            return ReturnCode_ADSERR_DEVICE_LICENSEEXPIRED
        case 0x726:
            return ReturnCode_ADSERR_DEVICE_LICENSEEXCEEDED
        case 0x727:
            return ReturnCode_ADSERR_DEVICE_LICENSEINVALID
        case 0x728:
            return ReturnCode_ADSERR_DEVICE_LICENSESYSTEMID
        case 0x729:
            return ReturnCode_ADSERR_DEVICE_LICENSENOTIMELIMIT
        case 0x72A:
            return ReturnCode_ADSERR_DEVICE_LICENSEFUTUREISSUE
        case 0x72B:
            return ReturnCode_ADSERR_DEVICE_LICENSETIMETOLONG
        case 0x72D:
            return ReturnCode_ADSERR_DEVICE_LICENSEDUPLICATED
        case 0x72E:
            return ReturnCode_ADSERR_DEVICE_SIGNATUREINVALID
        case 0x72F:
            return ReturnCode_ADSERR_DEVICE_CERTIFICATEINVALID
        case 0x72c:
            return ReturnCode_ADSERR_DEVICE_EXCEPTION
        case 0x740:
            return ReturnCode_ADSERR_CLIENT_ERROR
        case 0x741:
            return ReturnCode_ADSERR_CLIENT_INVALIDPARM
        case 0x742:
            return ReturnCode_ADSERR_CLIENT_LISTEMPTY
        case 0x743:
            return ReturnCode_ADSERR_CLIENT_VARUSED
        case 0x744:
            return ReturnCode_ADSERR_CLIENT_DUPLINVOKEID
        case 0x745:
            return ReturnCode_ADSERR_CLIENT_SYNCTIMEOUT
        case 0x746:
            return ReturnCode_ADSERR_CLIENT_W32ERROR
        case 0x747:
            return ReturnCode_ADSERR_CLIENT_TIMEOUTINVALID
        case 0x748:
            return ReturnCode_ADSERR_CLIENT_PORTNOTOPEN
        case 0x750:
            return ReturnCode_ADSERR_CLIENT_NOAMSADDR
        case 0x751:
            return ReturnCode_ADSERR_CLIENT_SYNCINTERNAL
        case 0x752:
            return ReturnCode_ADSERR_CLIENT_ADDHASH
        case 0x753:
            return ReturnCode_ADSERR_CLIENT_REMOVEHASH
        case 0x754:
            return ReturnCode_ADSERR_CLIENT_NOMORESYM
        case 0x755:
            return ReturnCode_ADSERR_CLIENT_SYNCRESINVALID
    }
    return 0
}

func ReturnCodeByName(value string) ReturnCode {
    switch value {
    case "OK":
        return ReturnCode_OK
    case "INTERNAL_ERROR":
        return ReturnCode_INTERNAL_ERROR
    case "NO_REALTIME":
        return ReturnCode_NO_REALTIME
    case "SAVE_ERROR":
        return ReturnCode_SAVE_ERROR
    case "MAILBOX_FULL":
        return ReturnCode_MAILBOX_FULL
    case "WRONG_HMSG":
        return ReturnCode_WRONG_HMSG
    case "TARGET_PORT_NOT_FOUND":
        return ReturnCode_TARGET_PORT_NOT_FOUND
    case "TARGET_HOST_NOT_FOUND":
        return ReturnCode_TARGET_HOST_NOT_FOUND
    case "UNKNOWN_COMMAND_ID":
        return ReturnCode_UNKNOWN_COMMAND_ID
    case "UNKNOWN_TASK_ID":
        return ReturnCode_UNKNOWN_TASK_ID
    case "NO_IO":
        return ReturnCode_NO_IO
    case "UNKNOWN_ADS_COMMAND":
        return ReturnCode_UNKNOWN_ADS_COMMAND
    case "WIN32_ERROR":
        return ReturnCode_WIN32_ERROR
    case "PORT_NOT_CONNECTED":
        return ReturnCode_PORT_NOT_CONNECTED
    case "INVALID_ADS_LENGTH":
        return ReturnCode_INVALID_ADS_LENGTH
    case "INVALID_AMS_NET_ID":
        return ReturnCode_INVALID_AMS_NET_ID
    case "LOW_INSTALLATION_LEVEL":
        return ReturnCode_LOW_INSTALLATION_LEVEL
    case "RTERR_INTERNAL":
        return ReturnCode_RTERR_INTERNAL
    case "RTERR_BADTIMERPERIODS":
        return ReturnCode_RTERR_BADTIMERPERIODS
    case "RTERR_INVALIDTASKPTR":
        return ReturnCode_RTERR_INVALIDTASKPTR
    case "RTERR_INVALIDSTACKPTR":
        return ReturnCode_RTERR_INVALIDSTACKPTR
    case "RTERR_PRIOEXISTS":
        return ReturnCode_RTERR_PRIOEXISTS
    case "RTERR_NOMORETCB":
        return ReturnCode_RTERR_NOMORETCB
    case "RTERR_NOMORESEMAS":
        return ReturnCode_RTERR_NOMORESEMAS
    case "RTERR_NOMOREQUEUES":
        return ReturnCode_RTERR_NOMOREQUEUES
    case "RTERR_EXTIRQALREADYDEF":
        return ReturnCode_RTERR_EXTIRQALREADYDEF
    case "RTERR_EXTIRQNOTDEF":
        return ReturnCode_RTERR_EXTIRQNOTDEF
    case "RTERR_EXTIRQINSTALLFAILED":
        return ReturnCode_RTERR_EXTIRQINSTALLFAILED
    case "RTERR_IRQLNOTLESSOREQUAL":
        return ReturnCode_RTERR_IRQLNOTLESSOREQUAL
    case "RTERR_VMXNOTSUPPORTED":
        return ReturnCode_RTERR_VMXNOTSUPPORTED
    case "RTERR_VMXDISABLED":
        return ReturnCode_RTERR_VMXDISABLED
    case "RTERR_VMXCONTROLSMISSING":
        return ReturnCode_RTERR_VMXCONTROLSMISSING
    case "RTERR_VMXENABLEFAILS":
        return ReturnCode_RTERR_VMXENABLEFAILS
    case "NO_DEBUGGING_AVAILABLE":
        return ReturnCode_NO_DEBUGGING_AVAILABLE
    case "PORT_DEACTIVATED":
        return ReturnCode_PORT_DEACTIVATED
    case "PORT_ALREADY_CONNECTED":
        return ReturnCode_PORT_ALREADY_CONNECTED
    case "ADS_SYNC_WIN32_ERROR":
        return ReturnCode_ADS_SYNC_WIN32_ERROR
    case "ADS_SYNC_TIMEOUT":
        return ReturnCode_ADS_SYNC_TIMEOUT
    case "ADS_SYNC_AMS_ERROR":
        return ReturnCode_ADS_SYNC_AMS_ERROR
    case "NO_INDEX_MAP_FOR_ADS_AVAILABLE":
        return ReturnCode_NO_INDEX_MAP_FOR_ADS_AVAILABLE
    case "INVALID_ADS_PORT":
        return ReturnCode_INVALID_ADS_PORT
    case "NO_MEMORY":
        return ReturnCode_NO_MEMORY
    case "TCP_SENDING_ERROR":
        return ReturnCode_TCP_SENDING_ERROR
    case "HOST_NOT_REACHABLE":
        return ReturnCode_HOST_NOT_REACHABLE
    case "INVALID_AMS_FRAGMENT":
        return ReturnCode_INVALID_AMS_FRAGMENT
    case "WSAETIMEDOUT":
        return ReturnCode_WSAETIMEDOUT
    case "WSAECONNREFUSED":
        return ReturnCode_WSAECONNREFUSED
    case "WSAEHOSTUNREACH":
        return ReturnCode_WSAEHOSTUNREACH
    case "ROUTERERR_NOLOCKEDMEMORY":
        return ReturnCode_ROUTERERR_NOLOCKEDMEMORY
    case "ROUTERERR_RESIZEMEMORY":
        return ReturnCode_ROUTERERR_RESIZEMEMORY
    case "ROUTERERR_MAILBOXFULL":
        return ReturnCode_ROUTERERR_MAILBOXFULL
    case "ROUTERERR_DEBUGBOXFULL":
        return ReturnCode_ROUTERERR_DEBUGBOXFULL
    case "ROUTERERR_UNKNOWNPORTTYPE":
        return ReturnCode_ROUTERERR_UNKNOWNPORTTYPE
    case "ROUTERERR_NOTINITIALIZED":
        return ReturnCode_ROUTERERR_NOTINITIALIZED
    case "ROUTERERR_PORTALREADYINUSE":
        return ReturnCode_ROUTERERR_PORTALREADYINUSE
    case "ROUTERERR_NOTREGISTERED":
        return ReturnCode_ROUTERERR_NOTREGISTERED
    case "ROUTERERR_NOMOREQUEUES":
        return ReturnCode_ROUTERERR_NOMOREQUEUES
    case "ROUTERERR_INVALIDPORT":
        return ReturnCode_ROUTERERR_INVALIDPORT
    case "ROUTERERR_NOTACTIVATED":
        return ReturnCode_ROUTERERR_NOTACTIVATED
    case "ADSERR_DEVICE_ERROR":
        return ReturnCode_ADSERR_DEVICE_ERROR
    case "ADSERR_DEVICE_SRVNOTSUPP":
        return ReturnCode_ADSERR_DEVICE_SRVNOTSUPP
    case "ADSERR_DEVICE_INVALIDGRP":
        return ReturnCode_ADSERR_DEVICE_INVALIDGRP
    case "ADSERR_DEVICE_INVALIDOFFSET":
        return ReturnCode_ADSERR_DEVICE_INVALIDOFFSET
    case "ADSERR_DEVICE_INVALIDACCESS":
        return ReturnCode_ADSERR_DEVICE_INVALIDACCESS
    case "ADSERR_DEVICE_INVALIDSIZE":
        return ReturnCode_ADSERR_DEVICE_INVALIDSIZE
    case "ADSERR_DEVICE_INVALIDDATA":
        return ReturnCode_ADSERR_DEVICE_INVALIDDATA
    case "ADSERR_DEVICE_NOTREADY":
        return ReturnCode_ADSERR_DEVICE_NOTREADY
    case "ADSERR_DEVICE_BUSY":
        return ReturnCode_ADSERR_DEVICE_BUSY
    case "ADSERR_DEVICE_INVALIDCONTEXT":
        return ReturnCode_ADSERR_DEVICE_INVALIDCONTEXT
    case "ADSERR_DEVICE_NOMEMORY":
        return ReturnCode_ADSERR_DEVICE_NOMEMORY
    case "ADSERR_DEVICE_INVALIDPARM":
        return ReturnCode_ADSERR_DEVICE_INVALIDPARM
    case "ADSERR_DEVICE_NOTFOUND":
        return ReturnCode_ADSERR_DEVICE_NOTFOUND
    case "ADSERR_DEVICE_SYNTAX":
        return ReturnCode_ADSERR_DEVICE_SYNTAX
    case "ADSERR_DEVICE_INCOMPATIBLE":
        return ReturnCode_ADSERR_DEVICE_INCOMPATIBLE
    case "ADSERR_DEVICE_EXISTS":
        return ReturnCode_ADSERR_DEVICE_EXISTS
    case "ADSERR_DEVICE_SYMBOLNOTFOUND":
        return ReturnCode_ADSERR_DEVICE_SYMBOLNOTFOUND
    case "ADSERR_DEVICE_SYMBOLVERSIONINVALID":
        return ReturnCode_ADSERR_DEVICE_SYMBOLVERSIONINVALID
    case "ADSERR_DEVICE_INVALIDSTATE":
        return ReturnCode_ADSERR_DEVICE_INVALIDSTATE
    case "ADSERR_DEVICE_TRANSMODENOTSUPP":
        return ReturnCode_ADSERR_DEVICE_TRANSMODENOTSUPP
    case "ADSERR_DEVICE_NOTIFYHNDINVALID":
        return ReturnCode_ADSERR_DEVICE_NOTIFYHNDINVALID
    case "ADSERR_DEVICE_CLIENTUNKNOWN":
        return ReturnCode_ADSERR_DEVICE_CLIENTUNKNOWN
    case "ADSERR_DEVICE_NOMOREHDLS":
        return ReturnCode_ADSERR_DEVICE_NOMOREHDLS
    case "ADSERR_DEVICE_INVALIDWATCHSIZE":
        return ReturnCode_ADSERR_DEVICE_INVALIDWATCHSIZE
    case "ADSERR_DEVICE_NOTINIT":
        return ReturnCode_ADSERR_DEVICE_NOTINIT
    case "ADSERR_DEVICE_TIMEOUT":
        return ReturnCode_ADSERR_DEVICE_TIMEOUT
    case "ADSERR_DEVICE_NOINTERFACE":
        return ReturnCode_ADSERR_DEVICE_NOINTERFACE
    case "ADSERR_DEVICE_INVALIDINTERFACE":
        return ReturnCode_ADSERR_DEVICE_INVALIDINTERFACE
    case "ADSERR_DEVICE_INVALIDCLSID":
        return ReturnCode_ADSERR_DEVICE_INVALIDCLSID
    case "ADSERR_DEVICE_INVALIDOBJID":
        return ReturnCode_ADSERR_DEVICE_INVALIDOBJID
    case "ADSERR_DEVICE_PENDING":
        return ReturnCode_ADSERR_DEVICE_PENDING
    case "ADSERR_DEVICE_ABORTED":
        return ReturnCode_ADSERR_DEVICE_ABORTED
    case "ADSERR_DEVICE_WARNING":
        return ReturnCode_ADSERR_DEVICE_WARNING
    case "ADSERR_DEVICE_INVALIDARRAYIDX":
        return ReturnCode_ADSERR_DEVICE_INVALIDARRAYIDX
    case "ADSERR_DEVICE_SYMBOLNOTACTIVE":
        return ReturnCode_ADSERR_DEVICE_SYMBOLNOTACTIVE
    case "ADSERR_DEVICE_ACCESSDENIED":
        return ReturnCode_ADSERR_DEVICE_ACCESSDENIED
    case "ADSERR_DEVICE_LICENSENOTFOUND":
        return ReturnCode_ADSERR_DEVICE_LICENSENOTFOUND
    case "ADSERR_DEVICE_LICENSEEXPIRED":
        return ReturnCode_ADSERR_DEVICE_LICENSEEXPIRED
    case "ADSERR_DEVICE_LICENSEEXCEEDED":
        return ReturnCode_ADSERR_DEVICE_LICENSEEXCEEDED
    case "ADSERR_DEVICE_LICENSEINVALID":
        return ReturnCode_ADSERR_DEVICE_LICENSEINVALID
    case "ADSERR_DEVICE_LICENSESYSTEMID":
        return ReturnCode_ADSERR_DEVICE_LICENSESYSTEMID
    case "ADSERR_DEVICE_LICENSENOTIMELIMIT":
        return ReturnCode_ADSERR_DEVICE_LICENSENOTIMELIMIT
    case "ADSERR_DEVICE_LICENSEFUTUREISSUE":
        return ReturnCode_ADSERR_DEVICE_LICENSEFUTUREISSUE
    case "ADSERR_DEVICE_LICENSETIMETOLONG":
        return ReturnCode_ADSERR_DEVICE_LICENSETIMETOLONG
    case "ADSERR_DEVICE_LICENSEDUPLICATED":
        return ReturnCode_ADSERR_DEVICE_LICENSEDUPLICATED
    case "ADSERR_DEVICE_SIGNATUREINVALID":
        return ReturnCode_ADSERR_DEVICE_SIGNATUREINVALID
    case "ADSERR_DEVICE_CERTIFICATEINVALID":
        return ReturnCode_ADSERR_DEVICE_CERTIFICATEINVALID
    case "ADSERR_DEVICE_EXCEPTION":
        return ReturnCode_ADSERR_DEVICE_EXCEPTION
    case "ADSERR_CLIENT_ERROR":
        return ReturnCode_ADSERR_CLIENT_ERROR
    case "ADSERR_CLIENT_INVALIDPARM":
        return ReturnCode_ADSERR_CLIENT_INVALIDPARM
    case "ADSERR_CLIENT_LISTEMPTY":
        return ReturnCode_ADSERR_CLIENT_LISTEMPTY
    case "ADSERR_CLIENT_VARUSED":
        return ReturnCode_ADSERR_CLIENT_VARUSED
    case "ADSERR_CLIENT_DUPLINVOKEID":
        return ReturnCode_ADSERR_CLIENT_DUPLINVOKEID
    case "ADSERR_CLIENT_SYNCTIMEOUT":
        return ReturnCode_ADSERR_CLIENT_SYNCTIMEOUT
    case "ADSERR_CLIENT_W32ERROR":
        return ReturnCode_ADSERR_CLIENT_W32ERROR
    case "ADSERR_CLIENT_TIMEOUTINVALID":
        return ReturnCode_ADSERR_CLIENT_TIMEOUTINVALID
    case "ADSERR_CLIENT_PORTNOTOPEN":
        return ReturnCode_ADSERR_CLIENT_PORTNOTOPEN
    case "ADSERR_CLIENT_NOAMSADDR":
        return ReturnCode_ADSERR_CLIENT_NOAMSADDR
    case "ADSERR_CLIENT_SYNCINTERNAL":
        return ReturnCode_ADSERR_CLIENT_SYNCINTERNAL
    case "ADSERR_CLIENT_ADDHASH":
        return ReturnCode_ADSERR_CLIENT_ADDHASH
    case "ADSERR_CLIENT_REMOVEHASH":
        return ReturnCode_ADSERR_CLIENT_REMOVEHASH
    case "ADSERR_CLIENT_NOMORESYM":
        return ReturnCode_ADSERR_CLIENT_NOMORESYM
    case "ADSERR_CLIENT_SYNCRESINVALID":
        return ReturnCode_ADSERR_CLIENT_SYNCRESINVALID
    }
    return 0
}

func CastReturnCode(structType interface{}) ReturnCode {
    castFunc := func(typ interface{}) ReturnCode {
        if sReturnCode, ok := typ.(ReturnCode); ok {
            return sReturnCode
        }
        return 0
    }
    return castFunc(structType)
}

func (m ReturnCode) LengthInBits() uint16 {
    return 32
}

func (m ReturnCode) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ReturnCodeParse(io *utils.ReadBuffer) (ReturnCode, error) {
    val, err := io.ReadUint32(32)
    if err != nil {
        return 0, nil
    }
    return ReturnCodeByValue(val), nil
}

func (e ReturnCode) Serialize(io utils.WriteBuffer) error {
    err := io.WriteUint32(32, uint32(e))
    return err
}

func (e ReturnCode) String() string {
    switch e {
    case ReturnCode_OK:
        return "OK"
    case ReturnCode_INTERNAL_ERROR:
        return "INTERNAL_ERROR"
    case ReturnCode_NO_REALTIME:
        return "NO_REALTIME"
    case ReturnCode_SAVE_ERROR:
        return "SAVE_ERROR"
    case ReturnCode_MAILBOX_FULL:
        return "MAILBOX_FULL"
    case ReturnCode_WRONG_HMSG:
        return "WRONG_HMSG"
    case ReturnCode_TARGET_PORT_NOT_FOUND:
        return "TARGET_PORT_NOT_FOUND"
    case ReturnCode_TARGET_HOST_NOT_FOUND:
        return "TARGET_HOST_NOT_FOUND"
    case ReturnCode_UNKNOWN_COMMAND_ID:
        return "UNKNOWN_COMMAND_ID"
    case ReturnCode_UNKNOWN_TASK_ID:
        return "UNKNOWN_TASK_ID"
    case ReturnCode_NO_IO:
        return "NO_IO"
    case ReturnCode_UNKNOWN_ADS_COMMAND:
        return "UNKNOWN_ADS_COMMAND"
    case ReturnCode_WIN32_ERROR:
        return "WIN32_ERROR"
    case ReturnCode_PORT_NOT_CONNECTED:
        return "PORT_NOT_CONNECTED"
    case ReturnCode_INVALID_ADS_LENGTH:
        return "INVALID_ADS_LENGTH"
    case ReturnCode_INVALID_AMS_NET_ID:
        return "INVALID_AMS_NET_ID"
    case ReturnCode_LOW_INSTALLATION_LEVEL:
        return "LOW_INSTALLATION_LEVEL"
    case ReturnCode_RTERR_INTERNAL:
        return "RTERR_INTERNAL"
    case ReturnCode_RTERR_BADTIMERPERIODS:
        return "RTERR_BADTIMERPERIODS"
    case ReturnCode_RTERR_INVALIDTASKPTR:
        return "RTERR_INVALIDTASKPTR"
    case ReturnCode_RTERR_INVALIDSTACKPTR:
        return "RTERR_INVALIDSTACKPTR"
    case ReturnCode_RTERR_PRIOEXISTS:
        return "RTERR_PRIOEXISTS"
    case ReturnCode_RTERR_NOMORETCB:
        return "RTERR_NOMORETCB"
    case ReturnCode_RTERR_NOMORESEMAS:
        return "RTERR_NOMORESEMAS"
    case ReturnCode_RTERR_NOMOREQUEUES:
        return "RTERR_NOMOREQUEUES"
    case ReturnCode_RTERR_EXTIRQALREADYDEF:
        return "RTERR_EXTIRQALREADYDEF"
    case ReturnCode_RTERR_EXTIRQNOTDEF:
        return "RTERR_EXTIRQNOTDEF"
    case ReturnCode_RTERR_EXTIRQINSTALLFAILED:
        return "RTERR_EXTIRQINSTALLFAILED"
    case ReturnCode_RTERR_IRQLNOTLESSOREQUAL:
        return "RTERR_IRQLNOTLESSOREQUAL"
    case ReturnCode_RTERR_VMXNOTSUPPORTED:
        return "RTERR_VMXNOTSUPPORTED"
    case ReturnCode_RTERR_VMXDISABLED:
        return "RTERR_VMXDISABLED"
    case ReturnCode_RTERR_VMXCONTROLSMISSING:
        return "RTERR_VMXCONTROLSMISSING"
    case ReturnCode_RTERR_VMXENABLEFAILS:
        return "RTERR_VMXENABLEFAILS"
    case ReturnCode_NO_DEBUGGING_AVAILABLE:
        return "NO_DEBUGGING_AVAILABLE"
    case ReturnCode_PORT_DEACTIVATED:
        return "PORT_DEACTIVATED"
    case ReturnCode_PORT_ALREADY_CONNECTED:
        return "PORT_ALREADY_CONNECTED"
    case ReturnCode_ADS_SYNC_WIN32_ERROR:
        return "ADS_SYNC_WIN32_ERROR"
    case ReturnCode_ADS_SYNC_TIMEOUT:
        return "ADS_SYNC_TIMEOUT"
    case ReturnCode_ADS_SYNC_AMS_ERROR:
        return "ADS_SYNC_AMS_ERROR"
    case ReturnCode_NO_INDEX_MAP_FOR_ADS_AVAILABLE:
        return "NO_INDEX_MAP_FOR_ADS_AVAILABLE"
    case ReturnCode_INVALID_ADS_PORT:
        return "INVALID_ADS_PORT"
    case ReturnCode_NO_MEMORY:
        return "NO_MEMORY"
    case ReturnCode_TCP_SENDING_ERROR:
        return "TCP_SENDING_ERROR"
    case ReturnCode_HOST_NOT_REACHABLE:
        return "HOST_NOT_REACHABLE"
    case ReturnCode_INVALID_AMS_FRAGMENT:
        return "INVALID_AMS_FRAGMENT"
    case ReturnCode_WSAETIMEDOUT:
        return "WSAETIMEDOUT"
    case ReturnCode_WSAECONNREFUSED:
        return "WSAECONNREFUSED"
    case ReturnCode_WSAEHOSTUNREACH:
        return "WSAEHOSTUNREACH"
    case ReturnCode_ROUTERERR_NOLOCKEDMEMORY:
        return "ROUTERERR_NOLOCKEDMEMORY"
    case ReturnCode_ROUTERERR_RESIZEMEMORY:
        return "ROUTERERR_RESIZEMEMORY"
    case ReturnCode_ROUTERERR_MAILBOXFULL:
        return "ROUTERERR_MAILBOXFULL"
    case ReturnCode_ROUTERERR_DEBUGBOXFULL:
        return "ROUTERERR_DEBUGBOXFULL"
    case ReturnCode_ROUTERERR_UNKNOWNPORTTYPE:
        return "ROUTERERR_UNKNOWNPORTTYPE"
    case ReturnCode_ROUTERERR_NOTINITIALIZED:
        return "ROUTERERR_NOTINITIALIZED"
    case ReturnCode_ROUTERERR_PORTALREADYINUSE:
        return "ROUTERERR_PORTALREADYINUSE"
    case ReturnCode_ROUTERERR_NOTREGISTERED:
        return "ROUTERERR_NOTREGISTERED"
    case ReturnCode_ROUTERERR_NOMOREQUEUES:
        return "ROUTERERR_NOMOREQUEUES"
    case ReturnCode_ROUTERERR_INVALIDPORT:
        return "ROUTERERR_INVALIDPORT"
    case ReturnCode_ROUTERERR_NOTACTIVATED:
        return "ROUTERERR_NOTACTIVATED"
    case ReturnCode_ADSERR_DEVICE_ERROR:
        return "ADSERR_DEVICE_ERROR"
    case ReturnCode_ADSERR_DEVICE_SRVNOTSUPP:
        return "ADSERR_DEVICE_SRVNOTSUPP"
    case ReturnCode_ADSERR_DEVICE_INVALIDGRP:
        return "ADSERR_DEVICE_INVALIDGRP"
    case ReturnCode_ADSERR_DEVICE_INVALIDOFFSET:
        return "ADSERR_DEVICE_INVALIDOFFSET"
    case ReturnCode_ADSERR_DEVICE_INVALIDACCESS:
        return "ADSERR_DEVICE_INVALIDACCESS"
    case ReturnCode_ADSERR_DEVICE_INVALIDSIZE:
        return "ADSERR_DEVICE_INVALIDSIZE"
    case ReturnCode_ADSERR_DEVICE_INVALIDDATA:
        return "ADSERR_DEVICE_INVALIDDATA"
    case ReturnCode_ADSERR_DEVICE_NOTREADY:
        return "ADSERR_DEVICE_NOTREADY"
    case ReturnCode_ADSERR_DEVICE_BUSY:
        return "ADSERR_DEVICE_BUSY"
    case ReturnCode_ADSERR_DEVICE_INVALIDCONTEXT:
        return "ADSERR_DEVICE_INVALIDCONTEXT"
    case ReturnCode_ADSERR_DEVICE_NOMEMORY:
        return "ADSERR_DEVICE_NOMEMORY"
    case ReturnCode_ADSERR_DEVICE_INVALIDPARM:
        return "ADSERR_DEVICE_INVALIDPARM"
    case ReturnCode_ADSERR_DEVICE_NOTFOUND:
        return "ADSERR_DEVICE_NOTFOUND"
    case ReturnCode_ADSERR_DEVICE_SYNTAX:
        return "ADSERR_DEVICE_SYNTAX"
    case ReturnCode_ADSERR_DEVICE_INCOMPATIBLE:
        return "ADSERR_DEVICE_INCOMPATIBLE"
    case ReturnCode_ADSERR_DEVICE_EXISTS:
        return "ADSERR_DEVICE_EXISTS"
    case ReturnCode_ADSERR_DEVICE_SYMBOLNOTFOUND:
        return "ADSERR_DEVICE_SYMBOLNOTFOUND"
    case ReturnCode_ADSERR_DEVICE_SYMBOLVERSIONINVALID:
        return "ADSERR_DEVICE_SYMBOLVERSIONINVALID"
    case ReturnCode_ADSERR_DEVICE_INVALIDSTATE:
        return "ADSERR_DEVICE_INVALIDSTATE"
    case ReturnCode_ADSERR_DEVICE_TRANSMODENOTSUPP:
        return "ADSERR_DEVICE_TRANSMODENOTSUPP"
    case ReturnCode_ADSERR_DEVICE_NOTIFYHNDINVALID:
        return "ADSERR_DEVICE_NOTIFYHNDINVALID"
    case ReturnCode_ADSERR_DEVICE_CLIENTUNKNOWN:
        return "ADSERR_DEVICE_CLIENTUNKNOWN"
    case ReturnCode_ADSERR_DEVICE_NOMOREHDLS:
        return "ADSERR_DEVICE_NOMOREHDLS"
    case ReturnCode_ADSERR_DEVICE_INVALIDWATCHSIZE:
        return "ADSERR_DEVICE_INVALIDWATCHSIZE"
    case ReturnCode_ADSERR_DEVICE_NOTINIT:
        return "ADSERR_DEVICE_NOTINIT"
    case ReturnCode_ADSERR_DEVICE_TIMEOUT:
        return "ADSERR_DEVICE_TIMEOUT"
    case ReturnCode_ADSERR_DEVICE_NOINTERFACE:
        return "ADSERR_DEVICE_NOINTERFACE"
    case ReturnCode_ADSERR_DEVICE_INVALIDINTERFACE:
        return "ADSERR_DEVICE_INVALIDINTERFACE"
    case ReturnCode_ADSERR_DEVICE_INVALIDCLSID:
        return "ADSERR_DEVICE_INVALIDCLSID"
    case ReturnCode_ADSERR_DEVICE_INVALIDOBJID:
        return "ADSERR_DEVICE_INVALIDOBJID"
    case ReturnCode_ADSERR_DEVICE_PENDING:
        return "ADSERR_DEVICE_PENDING"
    case ReturnCode_ADSERR_DEVICE_ABORTED:
        return "ADSERR_DEVICE_ABORTED"
    case ReturnCode_ADSERR_DEVICE_WARNING:
        return "ADSERR_DEVICE_WARNING"
    case ReturnCode_ADSERR_DEVICE_INVALIDARRAYIDX:
        return "ADSERR_DEVICE_INVALIDARRAYIDX"
    case ReturnCode_ADSERR_DEVICE_SYMBOLNOTACTIVE:
        return "ADSERR_DEVICE_SYMBOLNOTACTIVE"
    case ReturnCode_ADSERR_DEVICE_ACCESSDENIED:
        return "ADSERR_DEVICE_ACCESSDENIED"
    case ReturnCode_ADSERR_DEVICE_LICENSENOTFOUND:
        return "ADSERR_DEVICE_LICENSENOTFOUND"
    case ReturnCode_ADSERR_DEVICE_LICENSEEXPIRED:
        return "ADSERR_DEVICE_LICENSEEXPIRED"
    case ReturnCode_ADSERR_DEVICE_LICENSEEXCEEDED:
        return "ADSERR_DEVICE_LICENSEEXCEEDED"
    case ReturnCode_ADSERR_DEVICE_LICENSEINVALID:
        return "ADSERR_DEVICE_LICENSEINVALID"
    case ReturnCode_ADSERR_DEVICE_LICENSESYSTEMID:
        return "ADSERR_DEVICE_LICENSESYSTEMID"
    case ReturnCode_ADSERR_DEVICE_LICENSENOTIMELIMIT:
        return "ADSERR_DEVICE_LICENSENOTIMELIMIT"
    case ReturnCode_ADSERR_DEVICE_LICENSEFUTUREISSUE:
        return "ADSERR_DEVICE_LICENSEFUTUREISSUE"
    case ReturnCode_ADSERR_DEVICE_LICENSETIMETOLONG:
        return "ADSERR_DEVICE_LICENSETIMETOLONG"
    case ReturnCode_ADSERR_DEVICE_LICENSEDUPLICATED:
        return "ADSERR_DEVICE_LICENSEDUPLICATED"
    case ReturnCode_ADSERR_DEVICE_SIGNATUREINVALID:
        return "ADSERR_DEVICE_SIGNATUREINVALID"
    case ReturnCode_ADSERR_DEVICE_CERTIFICATEINVALID:
        return "ADSERR_DEVICE_CERTIFICATEINVALID"
    case ReturnCode_ADSERR_DEVICE_EXCEPTION:
        return "ADSERR_DEVICE_EXCEPTION"
    case ReturnCode_ADSERR_CLIENT_ERROR:
        return "ADSERR_CLIENT_ERROR"
    case ReturnCode_ADSERR_CLIENT_INVALIDPARM:
        return "ADSERR_CLIENT_INVALIDPARM"
    case ReturnCode_ADSERR_CLIENT_LISTEMPTY:
        return "ADSERR_CLIENT_LISTEMPTY"
    case ReturnCode_ADSERR_CLIENT_VARUSED:
        return "ADSERR_CLIENT_VARUSED"
    case ReturnCode_ADSERR_CLIENT_DUPLINVOKEID:
        return "ADSERR_CLIENT_DUPLINVOKEID"
    case ReturnCode_ADSERR_CLIENT_SYNCTIMEOUT:
        return "ADSERR_CLIENT_SYNCTIMEOUT"
    case ReturnCode_ADSERR_CLIENT_W32ERROR:
        return "ADSERR_CLIENT_W32ERROR"
    case ReturnCode_ADSERR_CLIENT_TIMEOUTINVALID:
        return "ADSERR_CLIENT_TIMEOUTINVALID"
    case ReturnCode_ADSERR_CLIENT_PORTNOTOPEN:
        return "ADSERR_CLIENT_PORTNOTOPEN"
    case ReturnCode_ADSERR_CLIENT_NOAMSADDR:
        return "ADSERR_CLIENT_NOAMSADDR"
    case ReturnCode_ADSERR_CLIENT_SYNCINTERNAL:
        return "ADSERR_CLIENT_SYNCINTERNAL"
    case ReturnCode_ADSERR_CLIENT_ADDHASH:
        return "ADSERR_CLIENT_ADDHASH"
    case ReturnCode_ADSERR_CLIENT_REMOVEHASH:
        return "ADSERR_CLIENT_REMOVEHASH"
    case ReturnCode_ADSERR_CLIENT_NOMORESYM:
        return "ADSERR_CLIENT_NOMORESYM"
    case ReturnCode_ADSERR_CLIENT_SYNCRESINVALID:
        return "ADSERR_CLIENT_SYNCRESINVALID"
    }
    return ""
}
