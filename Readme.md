# Setenv

**Setenv** Reads the `.env` file in a project root folder and sets it within the running process

### Install

```bash
go get github.com/chukwuemekaajah/setenv
```

### Import

```go
import (
	"github.com/chukwuemekaajah/setenv"
)
```
### Usage

The `setenv` only has one exported function called `SetEnv`. It is to be called at the top level Go file that instantiates a running process. It looks for the `.env` file in the project root folder and sets the corresponding key-value pairs to the `process` environmental variables lookup

```go

import (
	"github.com/chukwuemekaajah/setenv"
)

SetEnv()


```

### Todo
Cater for line comments