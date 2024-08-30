package org.project_kessel.inventory.client;


import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.project_kessel.clients.authn.AuthenticationConfig;

import java.util.HashMap;
import java.util.Hashtable;
import java.util.Optional;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import static io.smallrye.common.constraint.Assert.assertNotNull;
import static org.junit.jupiter.api.Assertions.*;

public class InventoryGrpcClientsManagerTest {

    public static Config.AuthenticationConfig dummyAuthConfigWithGoodOIDCClientCredentials() {
        return new Config.AuthenticationConfig() {
            @Override
            public AuthenticationConfig.AuthMode mode() {
                return AuthenticationConfig.AuthMode.OIDC_CLIENT_CREDENTIALS; // any non-disabled value
            }

            @Override
            public Optional<Config.OIDCClientCredentialsConfig> clientCredentialsConfig() {
                return Optional.of(new Config.OIDCClientCredentialsConfig() {
                    @Override
                    public String issuer() {
                        return "http://localhost:8090";
                    }

                    @Override
                    public String clientId() {
                        return "test";
                    }

                    @Override
                    public String clientSecret() {
                        return "test";
                    }

                    @Override
                    public Optional<String[]> scope() {
                        return Optional.empty();
                    }

                    @Override
                    public Optional<String> oidcClientCredentialsMinterImplementation() {
                        return Optional.empty();
                    }
                });
            }
        };
    }
}


