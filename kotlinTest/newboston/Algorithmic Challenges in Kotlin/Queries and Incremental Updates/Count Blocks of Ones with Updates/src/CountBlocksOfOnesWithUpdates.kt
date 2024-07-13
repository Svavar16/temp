class CountBlocksOfOnesWithUpdates(seq: CharSequence) {
    private var char: CharArray = CharArray(seq.length) { seq[it] }
    private var counter: Int = seq.count { it == '1' }

    fun countOnes(): Int {
        return counter
    }

    fun countBlocksOfOnes(): Int {
        var recentOne: Boolean = false
        var blockOfOnes: Int = 0

        for (c in char) {
            var isOne = c == '1'
            if (isOne) {
                if (!recentOne) {
                    blockOfOnes++
                }
            }
        }

        return blockOfOnes
    }

    fun flip(index: Int) {
        if (char[index] == '1') {
            char[index] = '0'
            counter--
        } else if (char[index] == '0') {
            char[index] = '1'
            counter++
        }
    }
}
