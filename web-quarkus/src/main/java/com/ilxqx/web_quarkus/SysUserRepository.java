package com.ilxqx.web_quarkus;

import java.util.List;

import io.quarkus.hibernate.orm.panache.PanacheRepository;
import jakarta.enterprise.context.ApplicationScoped;

@ApplicationScoped
public class SysUserRepository implements PanacheRepository<SysUser> {

    /**
     * Find users by userId between start and end
     *
     * @param start start userId
     * @param end end userId
     * @return list of users
     */
    public List<SysUser> findByUserIdBetween(Integer start, Integer end) {
        return list("userId between ?1 and ?2", start, end);
    }
}