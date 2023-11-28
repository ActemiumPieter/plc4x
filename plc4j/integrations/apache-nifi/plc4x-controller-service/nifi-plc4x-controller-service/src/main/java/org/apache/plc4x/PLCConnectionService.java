package org.apache.plc4x;

import org.apache.nifi.annotation.documentation.CapabilityDescription;
import org.apache.nifi.annotation.documentation.Tags;
import org.apache.nifi.components.PropertyDescriptor;
import org.apache.nifi.controller.AbstractControllerService;
import org.apache.nifi.controller.ConfigurationContext;
import org.apache.nifi.controller.ControllerServiceInitializationContext;
import org.apache.nifi.processor.exception.ProcessException;
import org.apache.nifi.reporting.InitializationException;

import java.util.List;

@Tags({"plc4x"})
@CapabilityDescription("Custom NiFi controller service for managing the connection to a PLC")
public class PLCConnectionService extends AbstractControllerService  {
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




}

