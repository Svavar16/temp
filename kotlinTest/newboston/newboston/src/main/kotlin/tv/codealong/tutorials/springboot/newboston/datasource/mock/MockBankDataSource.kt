package tv.codealong.tutorials.springboot.newboston.datasource.mock

import org.springframework.stereotype.Repository
import tv.codealong.tutorials.springboot.newboston.datasource.BankDataSource
import tv.codealong.tutorials.springboot.newboston.model.Bank
import java.lang.IllegalArgumentException

@Repository
class MockBankDataSource : BankDataSource {
    val banks = mutableListOf<Bank>(
        Bank("1234", 3.14, 17),
        Bank("1214", 17.0, 17),
        Bank("1537", 0.0, 17),
    )

    override fun retrieveBanks(): Collection<Bank>  =  banks
    override fun retrieveBank(accountNumber: String): Bank = banks.firstOrNull() { it.accountNumber == accountNumber}
        ?: throw NoSuchElementException("Counld not found a bank with the accountNumber: $accountNumber")

    override fun createBank(bank: Bank): Bank {
        if (banks.any{ it.accountNumber == bank.accountNumber}){
            throw IllegalArgumentException("Bank already exists")
        }
        banks.add(bank)

        return bank
    }

    override fun updateBank(bank: Bank): Bank {
        var currentBank = banks.firstOrNull { it.accountNumber == bank.accountNumber }
            ?: throw NoSuchElementException("Could not find the bank!")

        banks.remove(currentBank)
        banks.add(bank)

        return bank
    }

    override fun deleteBank(accountNumber: String) {
        var currentBank = banks.firstOrNull { it.accountNumber == accountNumber }
            ?: throw NoSuchElementException("Could not find the bank")

        banks.remove(currentBank)
    }
}