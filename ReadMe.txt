|---------------------|
|  File Name Checker  |
|---------------------|

--------------------------------------------------------------------------------

A small Tool which reads a Text File containing File Name List and ensures
that all the Files in the List do exist. A Path prefix is appended to each
File Name in the List.

Command Line Arguments:
    1-st:	A Folder used as a Prefix for each File in the List.
    2-nd:	A Path to the File containing the List of File Names.

--------------------------------------------------------------------------------

Usage Example.

> go build
> FileNameChecker.exe test test/Names.txt

You should see something similar to:
    "File does not exist: test\ccc.txt"

--------------------------------------------------------------------------------
