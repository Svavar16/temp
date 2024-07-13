fun countBlocksOfOnes(seq: CharSequence): Int {
    var oneRecently: Boolean = false
    var counter: Int = 0

    for (c in seq) {
        var isOne = c == '1'
        if (isOne) {
            if (!oneRecently) {
                counter++
            }

        }

        oneRecently = isOne
    }

    return counter
}
