# go-iscygpty

Golang port of iscygpty which is a part of [ptycheck](https://github.com/k-takata/ptycheck).
Check if a program running on mintty (or some other terminal) on Cygwin/MSYS.


## API

Only three functions are provided.

* `GetPipeName()`
* `IsCygwinPty()`
* `IsCygwinPtyUsed()`

See `_exapmle/example.go` for usage.


## License

[The MIT License](LICENSE)
