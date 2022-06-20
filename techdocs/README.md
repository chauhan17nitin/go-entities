# Tech Docs

## use of dummyOutput inside Present method
In the Present method you will see i am creating a new instance dummyOutput using the reflect type of output. 

The reason behind this if i directly use the value of output then while setting the fields will be unaddressable because of receiving empty struct, that's why i am initializing a dummyOutput.
