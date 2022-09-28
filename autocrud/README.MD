# Autocrud Module Scaffolding

To build the executable autocrud helper, you need to open main.go, then replace this :

```go
var moduleLocation = "../modules/"
```
with
```go
var moduleLocation = "./modules/"
```

then you can run the command with *go build*, and move the executable path to backendbillingdashboard root directory.