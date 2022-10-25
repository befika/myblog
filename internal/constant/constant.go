package constant

import (
	"errors"
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

//DbConnectionString connction string finder from the .env file
func DbConnectionString() (string, error) {
	addr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", "root", "", "localhost", "26257", "blog")
	return addr, nil
}

//StructValidator validates specific struct
func StructValidator(structName interface{}, validate *validator.Validate, trans ut.Translator) []string {
	var errorList []string
	errV := validate.Struct(structName)
	if errV != nil {
		errs := errV.(validator.ValidationErrors)
		for _, e := range errs {
			errorList = append(errorList, e.Translate(trans))
		}
		return errorList
	}
	return nil
}

// wrap field validator with error code
func VerifyInput(structName interface{}, validate *validator.Validate, trans ut.Translator) error {
	errs := StructValidator(structName, validate, trans)
	if errs == nil {
		return nil
	}
	return errors.New("input varification error")

}
