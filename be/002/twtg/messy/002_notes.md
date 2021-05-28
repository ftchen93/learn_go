# Questions

1. What is the difference between int32 and int64?

   One is 32 bits while the other is 64 bits.
   int64 is often a choice when memory isn't an issue.

2. What is the difference between int32 and uint32?

   **int32** is signed integer which include positive and negative range -2147483648 to +2147483647

   **uint32** is unsigned integer which only positive numbers are include range from 0 to 4294967295

3. What happens when we exceed the range of integer types?

   Overflow(during compile) or wraparound(during runtime, after compiled)

   E.g. 1(wraparound):

   var maxUint32 uint32 = 4294967295

   maxUnit32 = maxUnit32 + 1

   E.g. 2(overflow):
   var maxUint32 uint32
   maxUint32 = 4294967296