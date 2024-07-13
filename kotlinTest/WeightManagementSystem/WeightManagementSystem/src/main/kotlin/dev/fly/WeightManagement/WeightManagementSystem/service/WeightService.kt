package dev.fly.WeightManagement.WeightManagementSystem.service

import dev.fly.WeightManagement.WeightManagementSystem.datasource.WeightDataSource
import dev.fly.WeightManagement.WeightManagementSystem.model.Weight
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class WeightService @Autowired constructor(private val repository: WeightDataSource) {
    fun getAllWeights(): Collection<Weight> = repository.findAll()

}