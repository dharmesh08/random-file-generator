# random-file-generator
random-file-generator

I created this program to create random files with random data. Use it at your own risk.

#### Example

          $ godep go run main.go -dir=C:\Temp -howManyFiles=100 -total_lines=100 -remove-files=true -wait-before-remove=10s


Command Line Options
-dir = Directory where you like to store file\
-how-many-files= default 10
-total-lines = numbers lines in file
-remove-files= remove files default is true
-wait-before-remove= Wait before remove in secods default is 1s

