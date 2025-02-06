## Unpacking JSON (DECODING) from an api response

When you get data from a JSON API, you need to *unmarshall* it into a go struct.

Let's say you get data that looks like this from the api.

```json
{
    "name":"John",
    "age":25,
    "email","john@email.com"
}
```

Look at `main.go` for how to fetch and parse such data.


### handling nested json objects
Lets say you get the following JSON from an API.

Example JSON (Nested Object)
```json
{
    "id":1,
    "name","Alice",
    "contact": {
        "email" : "alice@example.com",
        "phone" : "+1234",
    }
}
```

check `nested.go` file on how to handle nested objects from an api