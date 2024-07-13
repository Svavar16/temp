package com.example.SvavarDemo.service

import com.example.SvavarDemo.datasource.testingdatasource
import com.example.SvavarDemo.datasource.testitemdatasource
import com.example.SvavarDemo.model.testing
import org.springframework.stereotype.Service
import java.util.*

@Service
class testingService(private val testingdatasource: testingdatasource, private val testitemdatasource: testitemdatasource) {
    fun getTestingItem(): List<testing> = testingdatasource.findAll()

    fun getSingleTesting(id: UUID): testing {
        var testingItems = testingdatasource.findById(id)

        if (testingItems.isPresent){
            val testing = testingItems.get()
            testing.testItems = emptyList()
            val testingResults = testitemdatasource.getAllByTestingId(id)
//
            testing.testItems = testingResults

            return testing
        } else {
            val testing = testing(UUID.randomUUID(), "", emptyList())
            return testing
        }
    }
}