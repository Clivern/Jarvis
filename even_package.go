// The even package implements a fast function for detecting if an integer
// is even or not.
package even

func Even(x int) bool {
    return x % 2 == 0
}

func Odd(x int) bool {
    return x % 2 == 1
}

/*
cp even_package.go even_package_test.go /Users/Mark/go/src/even/
% go test
% go build
% go install
*/