package com.ilxqx.web_java;

import java.util.List;
import java.util.concurrent.ThreadLocalRandom;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class SysUserController {
    private static final Logger LOG = LoggerFactory.getLogger(SysUserController.class);
    private final SysUserRepository sysUserRepository;

    public SysUserController(SysUserRepository sysUserRepository) {
        this.sysUserRepository = sysUserRepository;
    }

    @GetMapping("/")
    public List<SysUser> index() {
        var random = ThreadLocalRandom.current();
        int randomId1 = random.nextInt(0, 1000);
        int randomId2 = random.nextInt(8000, 9000);
        
        LOG.info("Generated random id1: {}, id2: {}", randomId1, randomId2);

        try {
            List<SysUser> users = sysUserRepository.findByUserIdBetween(randomId1, randomId2);
            LOG.info("Found {} users", users.size());
            return users;
        } catch (Exception e) {
            LOG.error("Error: {}", e.getMessage());
            return List.of();
        }
    }
} 