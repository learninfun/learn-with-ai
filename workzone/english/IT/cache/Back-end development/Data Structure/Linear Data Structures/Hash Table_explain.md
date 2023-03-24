

A hash table is a data structure that stores key-value pairs, in which an array of fixed size is used to store the data. When a value is stored in the hash table, a hash function is used to generate an index number, which is used to find the appropriate location in the array to store the data.

For example, let's say we want to create a hash table to store the ages of a group of people. We could use their names as keys, and their ages as values.

To store the data, we would first define an array of a certain size, such as 10. We would then use a hash function to generate an index number for each name. The hash function might simply add up the ASCII values of the characters in the name, and then take the remainder of that sum when divided by 10 (since our array is of size 10).

So if we wanted to store the age of someone named "John", the hash function might generate an index of 6. We would then store John's age in the 6th location of the array.

To retrieve John's age later on, we would use the same process - generate an index using the hash function (which would again be 6 for John's name), and then retrieve the age stored at that location in the array.