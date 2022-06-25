// manual-parse/main.go
// go mod init github.com/somedir/manual-parse
// go mod init example/manual-parse
// go build -c application
package main

import (
    "bufio"
    "errors"
    "fmt"
    "io"
    "os"
    "strconv"
)

type config struct {
    numTimes int
    printUsage bool
}

var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]

A greeter application which prints the name you entered <integers> number of times.
`, os.Args[0])

func printUsage(w io.Writer) {
    fmt.Fprintf(w, usageString)
}

func validateArgs(c config) error {
    if !(c.numTimes > 0) {
        return errors.New("Must specify a number greater than 0")
    }
    return nil
}

// TODO - Insert definition of ParseArgs() as earlier
func parseArgs(args []string) (config, error) {
    var numTimes int
    var err error
    c := config{}
    if len(args) != 1 {
        return c, errors.New("Invalid number of arguments")
    }
    if args[0] == "-h" || args[0] == "--help" {
        c.printUsage = true
        return c, nil
    }
    numTimes, err = strconv.Atoi(args[0])
    if err != nil {
        return c, err
    }
    c.numTimes = numTimes
    return c, nil
}

// TODO - Insert definition of getName() () as earlier
func getName(r io.Reader, w io.Writer) (string, error) {
    msg := "Your name pleae? Press the Enter key when done.\n"
    fmt.Fprintf(w, msg)

    scanner := bufio.NewScanner(r)
    scanner.Scan()
    if err := scanner.Err(); err != nil {
        return "", err
    }
    name := scanner.Text()
    if len(name) == 0 {
        return "", errors.New("You didn't enter your name")
    }
    return name, nil
} 

// TODO - Insert definition of greetUser() as earlier
func greetUser(c config, name string, w io.Writer) {
    msg := fmt.Sprintf("Nice to meet you %s\n", name)
    for i := 0; i < c.numTimes; i++ {
        fmt.Fprintf(w, msg)
    }
}

// TODO - Insert definition of runCmd() as earlier
func runCmd(r io.Reader, w io.Writer, c config) error {
    if c.printUsage {
        printUsage(w)
        return nil
    }
    name, err := getName(r, w)
    if err != nil {
        return err
    }
    greetUser(c, name, w)
    return nil
}


func main() {
    c, err := parseArgs(os.Args[1:])
    if err != nil {
        fmt.Fprintln(os.Stdout, err)
        printUsage(os.Stdout)
        os.Exit(1)
    }
    err = validateArgs(c)
    if err != nil {
        fmt.Fprintln(os.Stdout, err)
        printUsage(os.Stdout)
        os.Exit(1)
    }
    err = runCmd(os.Stdin, os.Stdout, c)
    if err != nil {
        fmt.Fprintln(os.Stdout, err)
        os.Exit(1)
    }
}


