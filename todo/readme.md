## Implementing JWT Authentication

In this tutorial we are going to explore the fundamentals of JWT authentication, understanding its significance, and then transition into hands on implementation.

Throughout this guide we will cover.

- A brief overview of JWT and its structure.
- Creating a Simple todo application with go
- Examining the Golang-JWT package
- Creating JWT tokens and adding claims using Golang-JWT
- Signing and verifying JWTs

### what is JWT
Json Web Tokens in full, serve as a compact and self-contained data structure for transmitting information securely between parties.

JWTs specify the token type, contain claims about an entity, and ensure integrity through cryptographic signatures.

### STRUCTURE OF A JWT
Jwt contains 3 parts: the header, the payload, and the signature.

1. Header

The header typically consists of two parts: the type of token, which is JWT, and the signing algorithm being used such as HMAC SHA256 or RSA.

An example:

```javascript
{
    "alg":"HS256",
    "typ": "JWT"
}
```

2. Payload

The payload contains claims. Claims are statements about and entity (typicall, the user) and additional data. There are 3 types of claims:
- registered
- public
- private

Registered claims include standard fields like issuer (`iss`), subject (`sub`), audience (`aud`), expiration time (`exp`) and issued at(`iat`).

```js
{
    "sub": "1234567890"
    "name": "Foo Bar"
    "iat": 14345343 // unix time
}
```

3. Signature

The signature is created by combining the encoded header, encoded payload, a secret, and the specified signing algorithm.
It ensures the integrity and authenticity of the token.

example (using HMAC SHA256)

```js
HMACSHA256(
    base64UrlEncode(header) + "." + base64UrlEncode(payload),
    secret
)
```
I think the dot acts as a separator


### Creating a simple todo application

For the purpose of demonstrating JWT authentication, we will create a simple todo list application with [ROLE BASED ACCESS CONTROL](RBAC).

There will be two roles:

- **senior** -> users have elevated privileges and can perform actions such as adding todo items to the list.

- **employee** -> employees have more restricted access and cannot add new todo items.


Roles will be assigned during user authentication and embedded in the JWT claims.

## Requirements
Gin framework
Golang Programming language
Knowledge of Golang
