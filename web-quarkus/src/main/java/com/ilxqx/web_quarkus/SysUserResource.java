package com.ilxqx.web_quarkus;

import java.util.List;
import java.util.concurrent.ThreadLocalRandom;

import io.quarkus.logging.Log;
import io.smallrye.common.annotation.RunOnVirtualThread;
import jakarta.inject.Inject;
import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;

@Path("/")
public class SysUserResource {
    @Inject
    SysUserRepository sysUserRepository;

    @RunOnVirtualThread
    @GET
    @Produces(MediaType.APPLICATION_JSON)
    public List<SysUser> index() {
        var random = ThreadLocalRandom.current();
        int randomId1 = random.nextInt(1000);
        int randomId2 = random.nextInt(8000, 9000);
        
        Log.infof("Generated random id1: %d, id2: %d", randomId1, randomId2);

        try {
            List<SysUser> users = sysUserRepository.findByUserIdBetween(randomId1, randomId2);
            Log.infof("Found %d users", users.size());
            return users;
        } catch (Exception e) {
            Log.errorf("Error: %s", e.getMessage());
            return List.of();
        }
    }
} 