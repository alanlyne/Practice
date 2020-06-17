// Credit: Derek Banas
//import java.util.random;

fun main(args: Array<String>) {
    println("Hello, world!")

    // ----- VARAIBLES -----

    val name = "Alan" // Read only variable

    var myAge = 24 // Changable variable

    // Kotlin uses type inference, but can alo define the type

    var bigInt: Int = Int.MAX_VALUE
    var smallInt: Int = Int.MIN_VALUE

    println("Biggest Int: " + bigInt)
    println("Smallest Int: " + smallInt)

    var bigLong: Long = Long.MAX_VALUE
    var smallLong: Long = Long.MIN_VALUE

    println("Biggest Long: " + bigLong)
    println("Smallest Long: " + smallLong)

    var bigDouble: Double = Double.MAX_VALUE
    var smallDouble: Double = Double.MIN_VALUE

    println("Biggest Double: " + bigDouble)
    println("Smallest Double: " + smallDouble)

    var bigFloat: Float = Float.MAX_VALUE
    var smallFloat: Float = Float.MIN_VALUE

    println("Biggest Float: " + bigFloat)
    println("Smallest Float: " + smallFloat)

    // Double are normally precise to 15 digits
    var dblNum1: Double = 1.111111111111111
    var dblNum2: Double = 1.111111111111111

    println("Sum: " + (dblNum1 + dblNum2))

    /* There is also
    Short 16 Bytes
    Byte 8 Bytes
    */

    // Booleans are either true or false
    if(true is Boolean){
        print("true is boolean\n")
    }

    // Characters are single quoted characters
    var letterGrade: Char = 'A'

    println("A is a Char: " + (letterGrade is Char))

    // ----- CASTING -----
    // Can cast from type to another
    // toShort, toInt, toLong, toFloat, toDouble, toChar, toString

    println("3.14 to Int: " + (3.14.toInt()))
    println("A to Int: " + (letterGrade.toInt())) // Ascii of A is 65
    println("65 to Char: " + (65.toChar()))

    // ----- STRINGS -----
    // Strings are double quoted series of characters

    val myName = "Alan Lyne"

    val longStr = """ This is a
    long string """

    var fName = "Alan"
    var lName = "Lyne"

    // Change values
    fName = "Venus"

    // Combine strings
    var fullName = fName + " " + lName

    // String interpolation
    println("Name: $fullName")

    // Can perform other operations with {}
    println("1 + 2 = ${1 + 2}")

    // Get length
    println("String length: ${longStr.length}")

    var str1 = "A random string"
    var str2 = "a random string"

    // Compare strings
    println("Strings Equal: ${str1.equals(str2)}")

    // Compare strings
    // 0: Equal, Neagative if less, Positive if greater
    println("Compare A to B: ${"A".compareTo("B")}") // A is 1 less than B, therefore -1

    // Get character at index
    println("2nd index: ${str1.get(2)}")

    // Get a substring from start up to, but not including, end
    println("Index 2-7: ${str1.subSequence(2,10)}")

    // Check if a string contains another
    println("Contains random: ${str1.contains("random")}")

    // ----- ARRAYS -----

}
