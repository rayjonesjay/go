### Tutorial: Developing a RESTful API with GO and GIN

This tutorial introduces basics of writing a RESTful web service API with Go and the Gin Web FrameWork (Gin).

Why Gin? Gin simplifies many coding tasks associated with building web applications, including web services.
You will use gin to route requests, retrieve request details and marshal JSON responses.

In this tutorial, you will build a RESTful API server with two endpoints.

The tutorial includes:
1. design api endpoints
2. create folder for your code
3. create the data
4. write a handler to return all items
5. write a handler to add an item
6. Write a handler to return a specific item

## PART 1: Design API endpoints

This api provides access to a store selling records or vinyl.
So you'll need to provide endpoints through which a client can get and add album for users.

When developing an API, you typically begin by designing endpoints.
Your API's users will have more success if the endpoints are easy to understand.

```
/albums
    - GET - get a list of all albums, returned as json
    - POST - add a new album from request data sent as JSON
```

```
/albums/:id
    GET - get an album by its ID, returning as JSON
```


When client makes a request at GET `/albums` the api will return all the albums as json.
