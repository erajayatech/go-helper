# go-helper
helper for go app

## Installation

```bash
go get github.com/erajayatech/go-helper
```

## Usage

### Validation Struct
```go
import "github.com/erajayatech/go-helper"

// struct
type Request struct {
	Email  string `json:"email" example:"andrietrilaksono@gmail.com" validate:"required" msg:"error_invalid_email"`
	Name string `json:"name" example:"andrie" validate:"required" msg:"error_invalid_name"`
}

var RequestErrorMessage = map[string]string{
	"error_invalid_email":  "email is required",
	"error_invalid_name":  "name is required",
}

func main() {
    var (
        request Request
    )

    errorMessage := helper.ValidateStruct(request, RequestErrorMessage)

    a, _ := json.MarshalIndent(errorMessage, "", "\t")
	fmt.Print(string(a))
}
```

```bash
{
	"email": "email is required",
	"name": "name is required"
}
```

## License
[MIT](https://choosealicense.com/licenses/mit/)