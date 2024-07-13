
fun countOnes(seq: CharSequence): Int {
    var counter: Int = 0
    for (i in seq){
        if (i === '1'){
            counter++
        }
    }
    return counter
}
