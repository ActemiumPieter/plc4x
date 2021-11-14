/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.spi.codegen.fields;

import org.apache.plc4x.java.spi.codegen.io.DataWriter;
import org.apache.plc4x.java.spi.generation.SerializationException;
import org.apache.plc4x.java.spi.generation.WithWriterArgs;

public class FieldWriterFactory {

    public static <T> void writeConstField(String logicalName, T value, DataWriter<T> dataWriter, WithWriterArgs... writerArgs) throws SerializationException {
        new FieldWriterConst<T>().writeField(logicalName, value, dataWriter, writerArgs);
    }

    public static <T> void writeEnumField(String logicalName, String innerName, T value, DataWriter<T> dataWriter, WithWriterArgs... writerArgs) throws SerializationException {
        new FieldWriterEnum<T>().writeField(logicalName, innerName, value, dataWriter, writerArgs);
    }

    public static <T> void writeDiscriminatorField(String logicalName, T value, DataWriter<T> dataWriter, WithWriterArgs... writerArgs) throws SerializationException {
        new FieldWriterDiscriminator<T>().writeField(logicalName, value, dataWriter, writerArgs);
    }

    public static <T> void writeDiscriminatorEnumField(String logicalName, String innerName, T value, DataWriter<T> dataWriter, WithWriterArgs... writerArgs) throws SerializationException {
        new FieldWriterDiscriminatorEnum<T>().writeField(logicalName, innerName, value, dataWriter, writerArgs);
    }

    public static <T> void writeOptionalField(String logicalName, T value, DataWriter<T> dataWriter, WithWriterArgs... writerArgs) throws SerializationException {
        new FieldWriterOptional<T>().writeField(logicalName, value, dataWriter, writerArgs);
    }

    public static <T> void writeOptionalEnumField(String logicalName, String innerName, T value, DataWriter<T> dataWriter, WithWriterArgs... writerArgs) throws SerializationException {
        new FieldWriterOptionalEnum<T>().writeField(logicalName, innerName, value, dataWriter, writerArgs);
    }

    public static <T> void writePaddingField(String logicalName, int timesPadding, T value, DataWriter<T> dataWriter, WithWriterArgs... writerArgs) throws SerializationException {
        new FieldWriterPadding<T>().writeField(logicalName, timesPadding, value, dataWriter, writerArgs);
    }

    public static <T> void writeReservedField(String logicalName, T value, DataWriter<T> dataWriter, WithWriterArgs... writerArgs) throws SerializationException {
        new FieldWriterReserved<T>().writeField(logicalName, value, dataWriter, writerArgs);
    }

    public static <T> void writeSimpleField(String logicalName, T value, DataWriter<T> dataWriter, WithWriterArgs... writerArgs) throws SerializationException {
        new FieldWriterSimple<T>().writeField(logicalName, value, dataWriter, writerArgs);
    }

    public static <T> void writeSimpleEnumField(String logicalName, String innerName, T value, DataWriter<T> dataWriter, WithWriterArgs... writerArgs) throws SerializationException {
        new FieldWriterSimpleEnum<T>().writeField(logicalName, innerName, value, dataWriter, writerArgs);
    }

}
