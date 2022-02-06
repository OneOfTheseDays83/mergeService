# Api

| Title            | Merge                                                                                         |
| :--------------- | :-----------                                                                                  |
| Brief            | This API can be used to merge a given number of intervals.  |
| Path             | /merge                                                                                        |
| Method           | POST                                                                                          |
| Body             | JSON with intervals as arrays (c.f. Example request)   |
| Success Response | **Code:** 200 (OK) |
| Error Response   | 400 (Bad Request)<br>500 (Internal Server Error)                                              |

## Constraints
- interval value: 0 < value < 2^31
- count of numbers per interval: 2
- max count of intervals: currently 100.000

## Example request

```shell
curl -v -X POST http://localhost:8000/v1/merge --data-raw '
{
    "intervals": [
        [1, 3],
        [8,15],
        [15,18],
        [2,6],
        [18,22]
     ]
}'
```

## Example response (JSON)

```
[[1,6],[8,15],[16,22]]
```