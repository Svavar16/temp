package dev.fly.WeightManagement.WeightManagementSystem.config

import org.springframework.context.annotation.Configuration
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity

@Configuration
@EnableWebSecurity
public class SecurityConfiguration : WebSecurityConfigurationAdapter {

}