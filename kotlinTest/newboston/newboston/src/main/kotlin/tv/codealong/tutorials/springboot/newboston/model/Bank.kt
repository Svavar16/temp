package tv.codealong.tutorials.springboot.newboston.model

// while we could have created a normal class that we can define everything, data class already does that
// ,so we don't have to implement equals, hashcode, and toString.
data class Bank(val accountNumber: String, val trust: Double, val transactionFee: Int)
