package com.example.SvavarDemo.datasource

import com.example.SvavarDemo.model.testing
import org.springframework.data.jpa.repository.JpaRepository
import java.util.*

interface testingdatasource: JpaRepository<testing, UUID>  {
}
