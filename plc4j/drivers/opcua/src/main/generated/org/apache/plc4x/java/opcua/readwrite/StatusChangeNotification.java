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
package org.apache.plc4x.java.opcua.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class StatusChangeNotification extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "820";
  }

  // Properties.
  protected final StatusCode status;
  protected final DiagnosticInfo diagnosticInfo;

  public StatusChangeNotification(StatusCode status, DiagnosticInfo diagnosticInfo) {
    super();
    this.status = status;
    this.diagnosticInfo = diagnosticInfo;
  }

  public StatusCode getStatus() {
    return status;
  }

  public DiagnosticInfo getDiagnosticInfo() {
    return diagnosticInfo;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("StatusChangeNotification");

    // Implicit Field (notificationLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int notificationLength = (int) (getLengthInBytes());
    writeImplicitField("notificationLength", notificationLength, writeSignedInt(writeBuffer, 32));

    // Simple Field (status)
    writeSimpleField("status", status, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (diagnosticInfo)
    writeSimpleField("diagnosticInfo", diagnosticInfo, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("StatusChangeNotification");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    StatusChangeNotification _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (notificationLength)
    lengthInBits += 32;

    // Simple field (status)
    lengthInBits += status.getLengthInBits();

    // Simple field (diagnosticInfo)
    lengthInBits += diagnosticInfo.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("StatusChangeNotification");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int notificationLength = readImplicitField("notificationLength", readSignedInt(readBuffer, 32));

    StatusCode status =
        readSimpleField(
            "status",
            new DataReaderComplexDefault<>(() -> StatusCode.staticParse(readBuffer), readBuffer));

    DiagnosticInfo diagnosticInfo =
        readSimpleField(
            "diagnosticInfo",
            new DataReaderComplexDefault<>(
                () -> DiagnosticInfo.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("StatusChangeNotification");
    // Create the instance
    return new StatusChangeNotificationBuilderImpl(status, diagnosticInfo);
  }

  public static class StatusChangeNotificationBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final StatusCode status;
    private final DiagnosticInfo diagnosticInfo;

    public StatusChangeNotificationBuilderImpl(StatusCode status, DiagnosticInfo diagnosticInfo) {
      this.status = status;
      this.diagnosticInfo = diagnosticInfo;
    }

    public StatusChangeNotification build() {
      StatusChangeNotification statusChangeNotification =
          new StatusChangeNotification(status, diagnosticInfo);
      return statusChangeNotification;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof StatusChangeNotification)) {
      return false;
    }
    StatusChangeNotification that = (StatusChangeNotification) o;
    return (getStatus() == that.getStatus())
        && (getDiagnosticInfo() == that.getDiagnosticInfo())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getStatus(), getDiagnosticInfo());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
