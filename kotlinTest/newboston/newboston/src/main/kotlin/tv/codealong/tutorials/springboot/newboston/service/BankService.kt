package tv.codealong.tutorials.springboot.newboston.service

import org.springframework.stereotype.Service
import tv.codealong.tutorials.springboot.newboston.datasource.BankDataSource
import tv.codealong.tutorials.springboot.newboston.model.Bank
import javax.sql.DataSource

@Service
class BankService(private val dataSource: BankDataSource) {
    fun getBanks(): Collection<Bank> = dataSource.retrieveBanks()
    fun getBank(accountNumber: String): Bank = dataSource.retrieveBank(accountNumber)
    fun addBank(bank: Bank): Bank = dataSource.createBank(bank)
    fun updateBank(bank: Bank): Bank = dataSource.updateBank(bank)
    fun deleteBank(accountNumber: String) = dataSource.deleteBank(accountNumber)
}