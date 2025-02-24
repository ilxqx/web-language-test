package com.ilxqx.web_java;

import java.time.LocalDate;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

@Entity
@Table(name = "sys_users")
public class SysUser {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "user_id")
    private Integer userId;

    @Column(name = "login_name", nullable = false)
    private String loginName;

    @Column(name = "user_pwd")
    private String userPwd;

    @Column(name = "pwd_expired_days")
    private Integer pwdExpiredDays;

    @Column(name = "created_time", nullable = false)
    private LocalDate createdTime;

    // Getters and Setters
    public Integer getUserId() {
        return userId;
    }

    public void setUserId(Integer userId) {
        this.userId = userId;
    }

    public String getLoginName() {
        return loginName;
    }

    public void setLoginName(String loginName) {
        this.loginName = loginName;
    }

    public String getUserPwd() {
        return userPwd;
    }

    public void setUserPwd(String userPwd) {
        this.userPwd = userPwd;
    }

    public Integer getPwdExpiredDays() {
        return pwdExpiredDays;
    }

    public void setPwdExpiredDays(Integer pwdExpiredDays) {
        this.pwdExpiredDays = pwdExpiredDays;
    }

    public LocalDate getCreatedTime() {
        return createdTime;
    }

    public void setCreatedTime(LocalDate createdTime) {
        this.createdTime = createdTime;
    }
}
