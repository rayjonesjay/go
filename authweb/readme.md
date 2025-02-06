<<<<<<< HEAD
## authweb
auth web aims to teach you on how to authenticate and handle authentication and session tokens
when signing in or login in a website.

It comprises 2 pages login.html and signup.html, which will be written using html , css and js.
Backend will be written in golang.

## 1. Authentication
Authentication is the process of verifying the identity of a user.(ensuring the person logging in is who they
claim to be). This process usually involves providing credentials like username and password.

It is needed since without it anyone can access the restricted parts of your website.

## 2. Tokenization
Tokenization is the process of creating a token (this is a unique string of characters) to represent a
user's session.
Instead of checking the username and password repeatedly just check for the token.
Tokens are temporary.

## 3. session management
Session management tracks a user's activity on a website during their visit. This includes keeping them
logged in for a certain time or until they manually log out.

## project flow
1. a user signs up with username and password
2. they log in with their credentials
3. the server creates a session token and sends it to user
4. the user accesses a private page that displays "hello you succeeded" and the server verifies their token
5. the token expires after 30 seconds, after which they get logged out


## implementation plan
```sh
go get -u github.com/gin-gonic/gin # Web framework
go get -u github.com/dgrijalva/jwt-go # JSON Web Tokens (JWT) library
```
See ![main](main.go) for the backend API
=======
### authentication in web apps

The authentication session of a web app is the heart of its defense against malicious threats.
>>>>>>> 65c7bc714d5b57f3e38a93f2731ef9c1acd071a0
