package com.ilxqx.web_java;

import java.util.List;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface SysUserRepository extends JpaRepository<SysUser, Integer> {
    
    /**
     * Find users by userId between start and end
     *
     * @param start start userId
     * @param end end userId
     * @return list of users
     */
    List<SysUser> findByUserIdBetween(Integer start, Integer end);
}
