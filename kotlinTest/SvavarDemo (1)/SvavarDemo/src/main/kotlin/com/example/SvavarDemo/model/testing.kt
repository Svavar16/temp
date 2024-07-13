package com.example.SvavarDemo.model

import io.swagger.v3.oas.annotations.media.Schema
import jakarta.persistence.*
import java.util.*
import kotlin.jvm.Transient

@Entity
@Schema(description = "users model")
data class testing(
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    var Id: UUID,
    var Name: String,

    @Transient
    var testItems: List<testItem> = mutableListOf()
)
