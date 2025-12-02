export const codeTemplates = {
    c: `#include <stdio.h>

int main() {
    printf("Hello, World!\\n");
    return 0;
}`,

    cpp: `#include <iostream>
using namespace std;

int main() {
    cout << "Hello, World!" << endl;
    return 0;
}`,

    python: `def main():
    print("Hello, World!")

if __name__ == "__main__":
    main()`,

    java: `public class Main {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
    }
}`,

    javascript: `function main() {
    console.log("Hello, World!");
}

main();`,

    go: `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`,

    rust: `fn main() {
    println!("Hello, World!");
}`,

    php: `<?php
echo "Hello, World!\\n";
?>`
};

export const getTemplate = (language) => {
    return codeTemplates[language] || codeTemplates.python;
};
