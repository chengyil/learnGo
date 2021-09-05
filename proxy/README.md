## Header Set vs Add

``` golang
func handle(w http.ResponseWriter, r *http.Request) {
    w.Header.Add("Header1", "bar")
    w.Header.Set("Header2", "foo")
}
```

What is the difference between Add vs Set?
``` bash
curl -X get http://localhost:3001/hello -H "Header1: one" -H "Header2: two"
```
On the proxied side
Header1:[one bar] 
Header2:[foo]

Add -> Append value to header key
Set -> Replace value to header key

## Accesing Header

``` golang
func handle(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Header.Get("Header1"))
    fmt.Println(r.Header["Header1"])
}
```

What is the difference between Get vs []?
``` bash
curl -X get http://localhost:3001/hello -H "Header1: one"
```
On the proxied side
Get("Header1") one 
["Header1] [one bar]


## Header Key

``` golang
func handle(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Header.Get("content-type"))
    fmt.Println(r.Header.Get(textproto.CanoicalMIMEHeaderKey("content-type")))
}
```
What is the value for each access
``` bash
curl -X get http://localhost:3001/hello -H "conTent-type: application/json"
```
On the proxied side
Get("content-type") "" 
Get(textproto.CanoicalMIMEHeaderKey("content-type")) application/json
