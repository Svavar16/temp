package dev.fly.WeightManagement.WeightManagementSystem.datasource

import dev.fly.WeightManagement.WeightManagementSystem.model.Weight
import org.springframework.data.jpa.repository.JpaRepository

interface WeightDataSource: JpaRepository<Weight, String>