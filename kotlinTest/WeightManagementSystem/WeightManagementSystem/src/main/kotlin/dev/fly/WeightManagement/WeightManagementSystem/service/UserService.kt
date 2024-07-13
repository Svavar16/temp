package dev.fly.WeightManagement.WeightManagementSystem.service

import dev.fly.WeightManagement.WeightManagementSystem.datasource.UsersDataSource
import dev.fly.WeightManagement.WeightManagementSystem.model.User
import org.springframework.beans.factory.annotation.Autowired
import java.util.*

class UserService @Autowired constructor(private val repository: UsersDataSource) {
    fun findById(id: String): Optional<User> = repository.findById(id)
}