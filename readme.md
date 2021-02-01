# Prefix/Infix Calculator

This is a prefix/infix notation calculator.

I've deployed a simple frontend on my personal website http://majolo.uk/kheiron

## Backend

* Written in Golang as an http server
* Unit testing suite
* Two different implementations (mainly to try out each approach)
    1. prefix notation was solved by converting the input to tokens and then recursively reducing starting from the right hand side
    2. infix notation was solved using a regex based solution, and replacing the string in place
* The solutions could be implemented to be algorithmically faster (prefix notation could be O(n) if implemented using a linked list for example) but the tradeoff for added code complexity didn't seem worth it based on the use case

### Rest API Service

```
POST /api/kheiron/prefix json
POST /api/kheiron/infix json
```

## Frontend

* Written in ReactJS, with minimal css
* The relevant component is extracted and inside `/frontend`

## Infrastructure

* The domain majolo.uk is registered and DNS is managed by AWS Route53.
* The backend is deployed as a Golang http server, wrapped by [algnhsa](https://github.com/akrylysov/algnhsa) to deploy it on AWS Lambda. This problem suits Lambda well as it's stateless and we expect low API usage, so we pay per request rather than constant server uptime.
* The Lambda is accessible via a proxy resource on API Gateway.
* The frontend is deployed as a static site hosted on S3.
* If this was project was a production codebase I'd provision these resources using Terraform, but these were configured manually.