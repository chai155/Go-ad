# Go-ad
A Go tool to generate http urls given the domain strings and generate md5 hash for the generated urls.

## Command
    ```
    Go-ad.exe -parallel < int > < domain(s) >
    ``` 

Optional argument:
-parallel argument represents the number of parallel jobs. By default the value is 10. A positive integer number can be passed to override the default value.
    
Compulsory arguments:
One or more strings separated by space, which represent the domain of the URL has(ve) to be passed as arguments.

Output:
One or more URL(s) and their corresponding MD5 hash values will be printed on the console.

Example(in windows OS):
    ```
    ./Go-ad.exe -parallel 2 google.com fb.com yahoo.com

    Output: http://fb.com fd71094d08a3637e76ccee6637f03713
            http://yahoo.com 873c87c71f8bf1d15a53ce0c0676971f
            http://google.com c7b920f57e553df2bb68272f61570210
    ```
