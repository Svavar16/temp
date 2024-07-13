package tv.codealong.tutorials.springboot.newboston.service

import io.mockk.every
import io.mockk.mockk
import io.mockk.verify
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test
import tv.codealong.tutorials.springboot.newboston.datasource.BankDataSource
import tv.codealong.tutorials.springboot.newboston.model.Bank

internal class BankServiceTest {

    private val dataSource: BankDataSource = mockk(relaxed = true)
    private val bankService = BankService(dataSource)

    @Test
    fun `Should call its data source to retrieve banks`(){
        // given

        // when
        bankService.getBanks()

        // then
        verify(exactly = 1) { dataSource.retrieveBanks() }
    }
}