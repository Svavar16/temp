package com.example.SvavarDemo.datasource

import com.example.SvavarDemo.model.testItem
import org.springframework.data.jpa.repository.JpaRepository
import java.util.UUID

interface testitemdatasource: JpaRepository<testItem, UUID> {
    fun getAllByTestingId(testingId: UUID): List<testItem>
}