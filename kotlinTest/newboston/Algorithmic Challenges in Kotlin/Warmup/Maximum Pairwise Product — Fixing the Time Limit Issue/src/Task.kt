fun maximumPairwiseProduct(a: IntArray): Long {
    val len = a.size
    var firstMax = Int.MIN_VALUE
    var secondMax = Int.MIN_VALUE
    for (i in a) {
        if(i > firstMax) {
            secondMax = firstMax
            firstMax = i
        } else if (i > secondMax){
            secondMax = i
        }
    }
    return firstMax.toLong() * secondMax

}
