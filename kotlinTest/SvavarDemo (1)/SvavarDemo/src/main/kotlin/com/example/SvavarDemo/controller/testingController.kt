package com.example.SvavarDemo.controller

import com.example.SvavarDemo.model.testing
import com.example.SvavarDemo.service.testingService
import io.swagger.v3.oas.annotations.Operation
import io.swagger.v3.oas.annotations.responses.ApiResponse
import io.swagger.v3.oas.annotations.responses.ApiResponses
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController
import java.util.*

@RestController
@RequestMapping("/api/testing")
class TestingController(private val service: testingService) {
    @GetMapping
    fun getAllTesting(): List<testing> = service.getTestingItem()

    @Operation(summary = "Sets a price for a chosen car", description = "Returns 202 if successful")
    @ApiResponses(
        value = [
            ApiResponse(responseCode = "202", description = "Successful Operation"),
            ApiResponse(responseCode = "404", description = "Such a car does not exist"),
        ]
    )
    @GetMapping("/{testingId}")
    fun getSingleTesting(@PathVariable testingId: UUID): testing = service.getSingleTesting(testingId)
}