# jsend

Implements the JSend specification for gin. See [JSend specification](https://github.com/omniti-labs/jsend) for details.

## Usage

```go
func (h Health) Status(c *gin.Context) {
  data := gin.H{"application": gin.H{"version": "20190808"}}
  resp := jsend.R{CTX: c, StatusCode: http.StatusOK, Data: data}
  resp.JSON()
}
```

The above will output the following JSON encoded message:

```json
{
  "status": "success",
  "data": {
    "application": {
      "version": "20190808"
    }
  }
}
```

## Thanks

* https://github.com/joaodlf/jsend
* https://github.com/gamegos/jsend
