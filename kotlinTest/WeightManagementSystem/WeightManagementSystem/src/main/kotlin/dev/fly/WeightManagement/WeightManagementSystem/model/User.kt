package dev.fly.WeightManagement.WeightManagementSystem.model

import jakarta.persistence.Entity
import jakarta.persistence.GeneratedValue
import jakarta.persistence.GenerationType
import jakarta.persistence.Id
import org.springframework.boot.autoconfigure.web.WebProperties.Resources.Chain.Strategy

@Entity
data class User(
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    var id: String? = null,
    var name: String,
    var password: String,
)