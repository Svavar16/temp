package dev.fly.WeightManagement.WeightManagementSystem.datasource

import dev.fly.WeightManagement.WeightManagementSystem.model.User
import org.springframework.data.jpa.repository.JpaRepository

interface UsersDataSource: JpaRepository<User, String> {
}