package org.apache.plc4x;

import org.apache.nifi.annotation.documentation.CapabilityDescription;
import org.apache.nifi.annotation.documentation.Tags;
import org.apache.nifi.controller.ControllerService;
import org.apache.plc4x.java.api.PlcConnection;

import java.util.concurrent.ExecutionException;

@Tags({"PLC4X"})
@CapabilityDescription("PLC4X service API")
public interface PLCConnection extends ControllerService {

        /**
         * Get the connection string for the PLC.
         *
         * @return The connection string.
         */
        String getConnectionString();
    }

