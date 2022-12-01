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

### Format Rupiah
```go
import "github.com/erajayatech/go-helper"

func main() {
	amountRupiah := helper.FormatRupiah(3000)

	fmt.Println(amountRupiah)
}
```

```bash
"Rp Rp 3.000"
```

### Format Gender
```go
import "github.com/erajayatech/go-helper"

func main() {
	gender := helper.FormatGender(1)

	fmt.Println(gender)
}
```

```bash
"M"
```

### Must Get Env
```go
import "github.com/erajayatech/go-helper"

func main() {
	env := helper.MustGetEnv("MODE")

	fmt.Println(env)
}
```

```bash
"local"
```

### Format Info Text
```go
import "github.com/erajayatech/go-helper"

func main() {
	env := helper.FormatInfoText("update data","11000034","processing","system")

	fmt.Println(env)
}
```

```bash
"update data #11000034 processing - updated_by: system"
```

### Expected Int
```go
import "github.com/erajayatech/go-helper"

func main() {
	value := helper.ExpectedInt(6.5)

	fmt.Println(value)
}
```

```bash
6
```

### Expected Int64
```go
import "github.com/erajayatech/go-helper"

func main() {
	var amount int64 = 67

	value := helper.ExpectedInt64(amount)

	fmt.Println(value)
}
```

```bash
67
```

### Expected String
```go
import "github.com/erajayatech/go-helper"

func main() {
	value := helper.ExpectedString(67)

	fmt.Println(value)
}
```

```bash
"67"
```

### Float To String
```go
import "github.com/erajayatech/go-helper"

func main() {
	value := helper.FloatToString(0.06)

	fmt.Println(value)
}
```

```bash
"0.060000"
```

### Validate Date Format
```go
import "github.com/erajayatech/go-helper"

func main() {
	date, err := helper.ValidateDateFormat("01-09-2010")

	fmt.Println(date)
}
```

```bash
"01-09-2010"
```


## License
[MIT](https://choosealicense.com/licenses/mit/)