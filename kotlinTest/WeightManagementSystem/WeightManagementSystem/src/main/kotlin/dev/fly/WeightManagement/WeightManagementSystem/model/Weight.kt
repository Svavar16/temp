package dev.fly.WeightManagement.WeightManagementSystem.model

import jakarta.persistence.GeneratedValue
import jakarta.persistence.Entity
import jakarta.persistence.GenerationType
import jakarta.persistence.Id
import java.util.Date

@Entity
data class Weight(
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    val id: String? = null,
    val date: Date,
    val weight: Double
)
