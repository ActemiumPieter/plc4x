package org.apache.plc4x;

import org.apache.nifi.annotation.documentation.CapabilityDescription;
import org.apache.nifi.annotation.documentation.Tags;
import org.apache.nifi.annotation.lifecycle.OnEnabled;
import org.apache.nifi.annotation.lifecycle.OnShutdown;
import org.apache.nifi.components.PropertyDescriptor;
import org.apache.nifi.controller.AbstractControllerService;
import org.apache.nifi.controller.ConfigurationContext;
import org.apache.nifi.controller.ControllerServiceInitializationContext;
import org.apache.nifi.reporting.InitializationException;

import java.util.List;
import java.util.concurrent.ExecutionException;

@Tags({"plc4x"})
@CapabilityDescription("Custom NiFi controller service for managing the connection to a PLC")
public class PLCConnectionService extends AbstractControllerService implements MyService{
    public static final PropertyDescriptor PLC_CONNECTION_STRING = new PropertyDescriptor.Builder()
        .name("PLC Connection string")
        .description("The connection string for the PLC.")
        .required(true)
        .build();

    String connectionString;
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
        try {
            super.init(config);
        }catch (Exception e){
            throw new InitializationException(e);
        }
    }



    @Override
    public void execute() throws ExecutionException {

    }

    @OnEnabled
    public void onConfigured(final ConfigurationContext context) throws InitializationException{
        connectionString = context.getProperty(PLC_CONNECTION_STRING).getValue();
    }

    @Override
    public String getProperty(){
        return connectionString;
    }
}

