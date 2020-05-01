 package com.bernardolobato.user.microsservice.security;
 

public final class SecurityConstants {

    public static final String AUTH_LOGIN_URL = "/api/autentica/";
    public static final String JWT_SECRET = "gObQ9HOQhrUEdlvuWkawk_JtX021F5g3rEzgNynXZRiXENZyfyIkoln1krJ7gotY_qhXetIcsNz3eXJRNdftEw";

    public static final String TOKEN_HEADER = "Authorization";
    public static final String TOKEN_PREFIX = "Bearer ";
    public static final String TOKEN_TYPE = "JWT";
    public static final String TOKEN_ISSUER = "secure-api";
    public static final String TOKEN_AUDIENCE = "secure-app";

    private SecurityConstants() {
        throw new IllegalStateException("Cannot create instance of static util class");
    }
}