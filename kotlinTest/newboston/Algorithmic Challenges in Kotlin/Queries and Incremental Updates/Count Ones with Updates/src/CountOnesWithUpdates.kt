
class CountOnesWithUpdates(seq: CharSequence)  {
// create required fields and methods here
    private var char = CharArray(seq.length) { seq[it] }
    private var count = seq.count { it == '1' }
    fun countOnes() : Int {
        return count
    }

    fun flip(index: Int) {
        if(char[index] === '0'){
            char[index] = '1'
            count++
        }
        else if (char[index] === '1'){
            char[index] = '0'
            count--
        }
    }
}
