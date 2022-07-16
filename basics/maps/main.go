package main

import "fmt"

func main() {
    var langs = make(map[string]string)

    langs["Go"] = "Go is a compiled programming language"
    langs["Python"] = "Python is a high-level, interpreted, object-oriented programming language"
    langs["Ruby"] = "Ruby is a dynamic, reflective, object-oriented, general-purpose programming language"
    langs["Java"] = "Java is a programming language and computing platform first released by Sun Microsystems"
    langs["C"] = "C is a general-purpose, imperative computer programming language"
    langs["C++"] = "C++ is a general-purpose, high-level, statically-linked, compiled, object-oriented programming language"
    langs["C#"] = "C# is a general-purpose, object-oriented, component-based, high-level programming language"
    langs["PHP"] = "PHP is a server-side scripting language designed for web development"
    langs["Perl"] = "Perl is a high-level, dynamically interpreted, multi-paradigm, compiled programming language"
    langs["JavaScript"] = "JavaScript is a high-level, dynamic, untyped, object-oriented programming language"
    langs["Scala"] = "Scala is a general-purpose programming language that combines the power of Java and C++"

    fmt.Println(langs)
    fmt.Println(langs["Python"])

    delete(langs, "Python")

    fmt.Println(langs)

    for key, value := range langs {
        fmt.Println(key, value)
    }
}
