val age: Int = 24
var greeting: String? = "Hello"
val nullValue: String? = null //Must use a '?' if null value is possible

fun main() {
    println("Hello World")

    val name: String = "Alan Lyne" // Value cannot be changed
    println(name)
    var otherName: String = "Alex" // Value can be changed
    otherName = "Venus"
    println(otherName)

    greeting = "Hi"
    println(greeting)
    println(age)

    greeting = null
    println(greeting)

    val animal = "Cat" // Type inferance
    println(animal)
}