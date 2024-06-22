## Description

Ascii-art is a program that receives a `string` as an argument and displays the `string` in a graphic representation using ASCII.

## Authors

- [x] najwang
- [x] ramuiruri
- [x] dochiel

## Usage: how to use the program

```bash
go run main.go "message as string"
```

## Implementation details: algorithm

1. Download the three banner files (`shadow`, `standard`, and `thinkertoy`) to a subdirectory, `banners`, in the repository root.

   After this task, we expect the following file structure:
    ```text
    ascii-art
        banners
            shadow.txt
            standard.txt
            thinkertoy.txt
    ```

2. Create a `main.go` program, that parses commandline arguments in the following format:

    - `go run . "Hello"`: The first argument is the text whose graphics is to be displayed.
    - `go run . "Hello" standard`: The first argument is the text whose graphics is to be displayed, while the second argument selects any of the given graphics files to use (must be one of `shadow`, `standard`, and `thinkertoy`).
    - `go run . "Hello" standard "World" shadow "This is working" thinkertoy`: The first argument is the text whose graphics is to be displayed, while the second argument selects any of the given graphics files to use (must be one of `shadow`, `standard`, and `thinkertoy`). Consequently, the third argument is the text whose graphics is to be displayed, while the fourth argument selects any of the given graphics files to use (must be one of `shadow`, `standard`, and `thinkertoy`). Consequently, the fifth argument is the text whose graphics is to be displayed, while the sixth argument selects any of the given graphics files to use (must be one of `shadow`, `standard`, and `thinkertoy`). We follow this format for an infinite number of such formatted command-line arguments.

3. Once done parsing the commandline args, the `main.go` file should also split the input string (the string to be displayed) into lines, then feed each line separately to a line graphics drawing function, `drawln`, that takes the line of string to be drawn, alongside the banner format, either of (`shadow`, `standard`, and `thinkertoy`). This should be done in a separate `draw` function, defined in the `main.go` file.

   > See task 5 below for the implementation details of the function, `drawln`.

4. Create and implement a module, `graphics`, that given a specific banner file, either of (`shadow`, `standard`, and `thinkertoy`), reads the banner file following the banner file format to generate a map of the ASCII characters to their graphical string representation. This map should be such that, the keys are the ASCII characters, while the values are the graphic representation of the ASCII characters, that is:

    <table>
    <tr>
    <td>key</td>
    <td>value</td>
    </tr>

    <tr>
    <td>H</td>
    <td>
    <pre>
     _    _  
    | |  | |
    | |__| |
    |  __  |
    | |  | |
    |_|  |_|
    </pre>
    </td>
    </tr>

    <tr>
    <td>T</td>
    <td>
    <pre>
     _______ 
    |__   __|
       | |   
       | |   
       | |   
       |_|   
    </pre>
    </td>

    </tr>
    </table>

   Given the Banner Format
    - Each character has a height of 8 lines.
    - Characters are separated by a new line \n.

   We come up with the following algorithm to read the file:

    ```text
    Open file
    map <- new map
    Start with the space character as the first ASCII character
    current_ascii_character <- ' ' 
    while not at end of file:
        Read the character separating newline
        graphics [8]string <- Read the next 8 lines containing the graphics
        map[current_ascii_character] <- graphics
        current_ascii_character++
    Close the file
    ```

5. Now to draw the graphics for a line of text, the `drawln` function follows the following steps:

    ```text
   for each graphics line from n=0 to n=7 do;
        for each character in line do;
            print the nth graphics line, i.e. graphics[n]
   print a newline to end the current line
    ```
   


