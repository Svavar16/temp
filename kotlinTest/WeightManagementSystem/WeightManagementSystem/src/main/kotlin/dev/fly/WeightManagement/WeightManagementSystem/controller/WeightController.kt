package dev.fly.WeightManagement.WeightManagementSystem.controller

import dev.fly.WeightManagement.WeightManagementSystem.service.WeightService
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/weight")
class WeightController(private val weightService: WeightService) {
}