package org.apache.plc4x;

import org.apache.nifi.annotation.documentation.CapabilityDescription;
import org.apache.nifi.annotation.documentation.Tags;
import org.apache.nifi.components.PropertyDescriptor;
import org.apache.nifi.controller.AbstractControllerService;
import org.apache.nifi.controller.ConfigurationContext;
import org.apache.nifi.controller.ControllerServiceInitializationContext;
import org.apache.nifi.reporting.InitializationException;
import org.apache.plc4x.java.api.PlcConnection;

import java.util.List;
import java.util.concurrent.ExecutionException;

@Tags({"plc4x"})
@CapabilityDescription("Custom NiFi controller service for managing the connection to a PLC")
public class PLCConnectionService extends AbstractControllerService implements PLCConnection {
    public static final PropertyDescriptor PLC_CONNECTION_STRING = new PropertyDescriptor.Builder()
        .name("PLC Connection string")
        .description("The connection string for the PLC.")
        .required(true)
        .build();

    private static final List<PropertyDescriptor> PROPERTY_DESCRIPTORS;

    static {
        PROPERTY_DESCRIPTORS = List.of(PLC_CONNECTION_STRING);
    }

    @Override
    protected List<PropertyDescriptor> getSupportedPropertyDescriptors() {
        return PROPERTY_DESCRIPTORS;
    }

    @Override
    protected void init(final ControllerServiceInitializationContext config) throws InitializationException {
        super.init(config);
    }


    public String getConnectionString() {
        ConfigurationContext context = getConfigurationContext();
        return context.getProperty(PLC_CONNECTION_STRING).getValue();
    }
}

