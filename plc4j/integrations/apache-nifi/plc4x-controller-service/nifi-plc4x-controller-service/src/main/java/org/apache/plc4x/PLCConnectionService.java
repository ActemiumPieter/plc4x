package org.apache.plc4x;

import org.apache.nifi.annotation.documentation.CapabilityDescription;
import org.apache.nifi.annotation.documentation.Tags;
import org.apache.nifi.annotation.lifecycle.OnEnabled;
import org.apache.nifi.components.PropertyDescriptor;
import org.apache.nifi.controller.AbstractControllerService;
import org.apache.nifi.controller.ConfigurationContext;
import org.apache.nifi.controller.ControllerServiceInitializationContext;
import org.apache.nifi.processor.util.StandardValidators;
import org.apache.nifi.reporting.InitializationException;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.utils.cache.CachedPlcConnectionManager;

import java.time.Duration;
import java.util.List;


@Tags({"plc4x"})
@CapabilityDescription("Custom NiFi controller service for managing the connection to a PLC")
public class PLCConnectionService extends AbstractControllerService implements MyService{
    public static final PropertyDescriptor PLC_CONNECTION_STRING = new PropertyDescriptor.Builder()
        .name("PLC Connection string")
        .description("The connection string for the PLC.")
        .addValidator(new Plc4xConnectionStringValidator())
        .required(true)
        .build();

    public static final PropertyDescriptor MAX_LEASE_TIME = new PropertyDescriptor.Builder()
        .name("Connection MaxLeaseTime")
        .description("The connection maximum lease time")
        .addValidator(StandardValidators.POSITIVE_INTEGER_VALIDATOR)
        .defaultValue("1000")
        .required(false)
        .build();


    public static final PropertyDescriptor MAX_WAIT_TIME = new PropertyDescriptor.Builder()
        .name("MaxWaitTime")
        .description("The connection maximum wait time")
        .addValidator(StandardValidators.POSITIVE_INTEGER_VALIDATOR)
        .defaultValue("500")
        .required(false)
        .build();

    String connectionString;

    protected CachedPlcConnectionManager getConnectionManager() {
        return connectionManager;
    }
    public CachedPlcConnectionManager connectionManager;
    private static final List<PropertyDescriptor> PROPERTY_DESCRIPTORS;

    static {
        PROPERTY_DESCRIPTORS = List.of(PLC_CONNECTION_STRING, MAX_LEASE_TIME, MAX_WAIT_TIME);
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

    public void refreshConnectionManager() {
        connectionManager = CachedPlcConnectionManager.getBuilder()
            .withMaxLeaseTime(Duration.ofSeconds(maxLeaseTime))
            .withMaxWaitTime(Duration.ofSeconds(maxWaitTime))
            .build();
    }

    public void closeConnection(){
        getConnectionManager().removeCachedConnection(PLC_CONNECTION_STRING.toString());
        refreshConnectionManager();
    }

    long maxLeaseTime;
    long maxWaitTime;
    @OnEnabled
    public void onConfigured(final ConfigurationContext context) throws InitializationException{
        connectionString = context.getProperty(PLC_CONNECTION_STRING).getValue();
        if (connectionString == null || connectionString.isEmpty()){
            throw new InitializationException("PLC_CONNECTION_STRING is null or empty");
        }
        maxLeaseTime = Long.parseLong(context.getProperty(MAX_LEASE_TIME).getValue());
        maxWaitTime = Long.parseLong(context.getProperty(MAX_WAIT_TIME).getValue());
        refreshConnectionManager();
    }

    @Override
    public PlcConnection getConnection() throws PlcConnectionException  {
        try {
            return getConnectionManager().getConnection(getPlc4xConnecionString());
        } catch (PlcConnectionException e) {
            throw new PlcConnectionException(e);
        }
    }

    @Override
    public String getPlc4xConnecionString() {
        return connectionString;
    }
}

