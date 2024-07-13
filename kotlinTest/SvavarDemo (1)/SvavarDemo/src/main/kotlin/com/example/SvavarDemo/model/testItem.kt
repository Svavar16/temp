package com.example.SvavarDemo.model

import io.swagger.v3.oas.annotations.media.Schema
import jakarta.persistence.Entity
import jakarta.persistence.GeneratedValue
import jakarta.persistence.GenerationType
import jakarta.persistence.Id
import java.util.*

@Entity
@Schema(description = "testItems")
data class testItem(
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    var Id: UUID,
    var Name: String,
    var testingId: UUID,
)
