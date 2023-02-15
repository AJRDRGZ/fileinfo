## FileInfo
This package contains two functions for get information about of a file, this has support for unix and windows system.

This package was created for use it in the EDteam's course [Go desde Cero](https://ed.team/cursos/go) 

* `IsHidden`: returns if a file is hidden

```go
    IsHidden(route)
```

* `GetUserAndGroup`: returns the user name and group name of an file, for windows always returns empty. You need pass as argument a fs.FileInfo.Sys() value

```go
    GetUserAndGroup(infoSys)
```

## Disclaimer
This package is only for education use.