A confusion in the exercise's audit section:
    The audit asks to pass the flag after a string like this:
        "HeY GuYs" --output=orange.txt
    This is actually not possible. Flags can only come before any other non-flag arguments. This is from the official go documentation (https://pkg.go.dev/flag):
        "  Flag parsing stops just before the first non-flag argument ("-" is a non-flag argument) or after the terminator "--".  "
    Therefore this code asks to pass the flag before the string that is to be printed.


To be able to run the code, first all the files from repo need to be downloaded.

The code acccepts a banner (output.txt by default) and two string as arguments. The first string is the string to be written into a file, the second one determines the style of the string (standard, shadow, or thinkertoy). To run the code, type this in terminal:
go run main.go --output=hello.txt "some string" shadow


The string has certain constraints: 
* Only the characters that are present in the "standard.txt" file can be printed.
* To print special characters backslash (\) needs to be placed before the character:
                                                                                    \" \\
* To add a new line either '\n' or '\\n' can be used.

All the characters that are available are:
 !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~
(the first character is white space)