### ONE TIME PAD

In cryptography, the one-time pad OTP, is an encryption technique that is impossible to crack.

It requires the use of a single-use pre-shared key that is larger than or equal to the size of the message being sent.

The pre-shared key cannot be less than the length of the message.

In this technique, a plaintext is paired with a random key (also referred to as the One Time Pad).

The each bit or character of the plaintext is encrypted by combining it with the corresponding bit or character from the pad using modular addition.

#### Modular addition

Modular addition is a fundamental concept in number theory and compute science, widely used in
cryptography, coding theory, and digital signal processing.

Modular addition involves  adding two numbers and then taking the remainder when the sum is divided by a modulus.

For integers a and b, a positive integer N(the modulus) the modular addition of a and b is given by.

```text
(a+b) mod n 
```
Example
```js
let a = 7;
let b = 5;
let n = 6;
let modResult = (a + b) % n
console.log(modResult) // 0
```


***

The resulting ciphertext will be impossible to decrypt or break if the following conditions are met:

1. the key must be at least as long as the plaintext
2. the key must be truly random
3. the key must never be reused in whole or in part
4. the key must be kept secret by communicating parties

