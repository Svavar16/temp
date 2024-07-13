package tv.codealong.tutorials.springboot.newboston.controller

import com.fasterxml.jackson.databind.ObjectMapper
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Nested
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.TestInstance
import org.junit.jupiter.api.TestInstance.*
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.http.MediaType
import org.springframework.test.web.servlet.*
import tv.codealong.tutorials.springboot.newboston.model.Bank

@SpringBootTest
@AutoConfigureMockMvc
internal class BankControllerTest @Autowired constructor(
    val mockMvc: MockMvc,
    val objectMapper: ObjectMapper
){

    val baseUrl = "/api/banks"

    @Nested
    @DisplayName("A GetBanks()")
    @TestInstance(Lifecycle.PER_CLASS)
    inner class AGetBanks {
        @Test
        fun `Should return all banks`() {
            // when/then
            mockMvc.get(baseUrl)
                .andDo { print() }
                .andExpect {
                    status { isOk() }
                    content { contentType(MediaType.APPLICATION_JSON) }
                    jsonPath("$[0].accountNumber") { value("1234")}
                }
        }
    }


    @Nested
    @DisplayName("B GetBank()")
    @TestInstance(Lifecycle.PER_CLASS)
    inner class BGetBank {
        @Test
        fun `Should return single bank`(){
            // given
            val accountNumber = 1234

            //when
            mockMvc.get("$baseUrl/$accountNumber")

            //then
                .andDo { print() }
                .andExpect {
                    status { isOk() }
                    content { contentType((MediaType.APPLICATION_JSON)) }
                    jsonPath("$.trust") { value("3.14") }
                    jsonPath("$.transactionFee") { value("17") }
                }
        }

        @Test
        fun `Should return NOT FOUND if the account number does not exist`(){
            // given
            val accountNumber = "does_not_exist"

            // when
            mockMvc.get("$baseUrl/$accountNumber")

            // then
                .andDo { print() }
                .andExpect { status { isNotFound() } }
        }
    }
    @Nested
    @DisplayName("C create PostBank()")
    @TestInstance(Lifecycle.PER_CLASS)
    inner class CPostBank{
        @Test
        fun `Should Create New Bank`() {
            // given
            val newBank = Bank("acc123",314.17, 2)
            // when
            val performPost = mockMvc.post(baseUrl) {
                contentType = MediaType.APPLICATION_JSON
                content = objectMapper.writeValueAsString(newBank)
            }

            // then
            performPost
                .andDo { print() }
                .andExpect {
                    status { isCreated() }
                    content {
                        contentType(MediaType.APPLICATION_JSON)
                        json(objectMapper.writeValueAsString(newBank))
                    }
                }
        }
    }

    @Nested
    @DisplayName("D PatchBank()")
    @TestInstance(Lifecycle.PER_CLASS)
    inner class DPatchBank {
        @Test
        fun `Should Patch A Bank`(){
            // given
            val updatedBank = Bank("1234",1.00, 1)
            // when
            var performPatch = mockMvc.patch(baseUrl) {
                contentType = MediaType.APPLICATION_JSON
                content = objectMapper.writeValueAsString(updatedBank)
            }

            // then
            performPatch
                .andDo { print() }
                .andExpect {
                    status { isOk() }
                    content {
                        contentType(MediaType.APPLICATION_JSON)
                        json(objectMapper.writeValueAsString(updatedBank))
                    }
                }

            mockMvc.get("$baseUrl/${updatedBank.accountNumber}")
                .andExpect { content { json(objectMapper.writeValueAsString(updatedBank)) } }
        }

        @Test
        fun `Should return bad request if not bank is found`(){
            // given
            val invalidBank = Bank("15667", 4.0, 1)

            // when
            val performPatch = mockMvc.patch(baseUrl) {
                contentType = MediaType.APPLICATION_JSON
                content = objectMapper.writeValueAsString(invalidBank)
            }

            // then
            performPatch
                .andDo { print() }
                .andExpect { status { isNotFound() } }
        }
    }

    @Nested
    @DisplayName("E DeleteBank()")
    @TestInstance(Lifecycle.PER_CLASS)
    inner class EDeleteBank {
        @Test
        fun `Should delete a bank`(){
            // given
            val accountNumber = 1234
            // when/then
            mockMvc.delete("$baseUrl/$accountNumber")
                .andDo { print() }
                .andExpect { status { isNoContent() } }

            mockMvc.get("$baseUrl/$accountNumber")
                .andExpect { status { isNotFound() } }
        }

//        @Test
//        fun `Should return Not Found when trying to delete an empty bank`(){
//            // given
//            val accountNumber = 1234
//            // when/then
//            mockMvc.delete("$baseUrl/$accountNumber")
//                .andDo { print() }
//                .andExpect { status { isNotFound() } }
//        }>
    }
}
