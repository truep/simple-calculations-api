# simple-calculations-api
anylapus first test

Base Path `/api`

Implements four ariphmetical methods:
 * addition(сложение)  `/add`
 * subtraction(вычитание) `/sub`
 * multiplication(умножение) `/mul`
 * division(деление) `/div`

Waiting parameters from query, like `?a=1&b=2`
### build & start server
```sh
docker-compose up -d --build
```

### query example
```
http://localhost:8080/api/add?a=4&b=5

```

### Response examples
Good request
```sh
15:29 $ curl -X GET "http://localhost:8080/api/add?a=4&b=5"
{"Success":true,"ErrCode":"200","Value":9}
```
Bad request
```sh
15:30 $ curl -X GET "http://localhost:8080/api/div?a=4&b=0"
{"Success":false,"ErrCode":"400"}
```


### tests
simple test

```
make test
```

simple covarage
```
make covtest
```