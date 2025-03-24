People who "program" in HTML and call themselves "programmers" can finally be programmers for real.

All code written in the programming language HTML (PL/HTML for short) is syntactically valid HTML code
(validated using [W3C validator][1]), which was one of the main goals during language design.
Also, the names of all statements were handpicked from the list of HTML elements in such a way that name or
description of an element describes the meaning of the statement (`var`, `input`, `output`, etc).

[1]: https://validator.w3.org/#validate-by-input

The following program displays the first `n` numbers in the Fibonacci series.

```html
<!doctype html>
<html lang="en">
    <head>
        <!-- Program to Display Fibonacci Series -->
        <title>`Fibonacci Series`</title>
    </head>
    <body>

        <main>

            <var class="integer">a</var>
            <var class="integer">b</var>
            <var class="integer">c</var>
            <var class="integer">i</var>
            <var class="integer">n</var>

            <data value="0">a</data>
            <data value="1">b</data>
            <data value="1">i</data>

            <output>`n: `</output>
            <input name="n">

            <div data-while="i &leq; n">
                <output>a + ` `</output>
                <data value="a + b">c</data>
                <data value="b">a</data>
                <data value="c">b</data>
                <data value="i + 1">i</data>
            </div>

        </main>

    </body>
</html>

```

**BUG NOTE:** If you get *Unexpected token 'EOF' at 0:0* as an error place an empty line at the end of the file.

Other examples can be found in the *tests/* folder.

# Specification
  * Keywords: `doctype`, `lang`, `html`, `head`, `title`,`body`, `main`, `var`, `class`, `output`, `input`, `name`, `data`, `value`, `div`, `if`, `while`
  * Builtin types: `integer`, `real`, `boolean`, `string`
  * Arithmetical operators: `+`, `-`, `*`, `/`, `%`, `(`, `)`
  * Logical operators: `&and;`, `&or;`, `!`
  * Comparison operators: `&lt;`, `&gt;`, `&leq;`, `&geq;`, `&equals;`, `&ne;`
  * Special characters: `\\`, `\t`, `\n`

# References
Useful reading materials:
  - [CS143 Compilers](https://web.stanford.edu/class/archive/cs/cs143/cs143.1128/)
  - [Implementing Lexers and Parsers](http://www.cse.chalmers.se/edu/year/2015/course/DAT150/lectures/proglang-04.html)
  - [A Simple Recursive Descent Parser](http://math.hws.edu/javanotes/c9/s5.html)
  - [Let’s Build A Simple Interpreter](https://ruslanspivak.com/lsbasi-part1/)

Source code of compilers for some programming languages:
  - [Go](https://github.com/golang/go/blob/master/src/go)
  - [TypeScript](https://github.com/microsoft/TypeScript/tree/master/src/compiler)
