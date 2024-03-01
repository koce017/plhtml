go build -o plhtml.exe
plhtml tests/hello.html tests/interpreter/hello.in.txt
plhtml tests/scopes.html
plhtml tests/factorial.html tests/interpreter/factorial.in.txt
plhtml tests/fibonacci.html tests/interpreter/fibonacci.in.txt
plhtml tests/leap.html tests/interpreter/leap.in.txt
plhtml tests/prime.html tests/interpreter/prime.in.txt