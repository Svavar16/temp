package dev.fly.WeightManagement.WeightManagementSystem.service

import org.springframework.security.oauth2.jwt.JwtDecoder
import org.springframework.security.oauth2.jwt.JwtEncoder


class TokenService(private val jwtDecoder: JwtDecoder, private val jwtEncoder: JwtEncoder) {
}