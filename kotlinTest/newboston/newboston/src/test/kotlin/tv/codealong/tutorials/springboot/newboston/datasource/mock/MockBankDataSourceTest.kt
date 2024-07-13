package tv.codealong.tutorials.springboot.newboston.datasource.mock

import org.assertj.core.api.Assertions
import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class MockBankDataSourceTest {

    private val mockBankDataSource = MockBankDataSource()

    @Test
    fun `should provide a collection of banks`(){
        // when
        val banks = mockBankDataSource.retrieveBanks()

        // then
        assertThat(banks).isNotEmpty
        assertThat(banks.size).isGreaterThanOrEqualTo(3)
    }

    @Test
    fun `Should provide some mock data`(){
        //when
        val banks = mockBankDataSource.retrieveBanks()

        // then
        assertThat(banks).allMatch { it.accountNumber.isNotBlank() }
        assertThat(banks).anyMatch { it.trust != 0.0 }
        assertThat(banks).anyMatch { it.transactionFee != 0 }

    }
}