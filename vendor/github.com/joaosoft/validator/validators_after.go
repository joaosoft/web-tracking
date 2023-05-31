package validator

import (
	"errors"
	"regexp"
	"strings"
)

func (v *Validator) validate_error(context *ValidatorContext, validationData *ValidationData) []error {
	rtnErrs := make([]error, 0)
	added := make(map[string]bool)
	var errorList []error

	for i, e := range *validationData.Errors {
		if _, ok := validationData.ErrorsReplaced[e]; ok {
			continue
		}

		if v.errorCodeHandler == nil {
			return rtnErrs
		}
		var expected string

		if validationData.Expected != nil {
			expected = validationData.Expected.(string)
		}

		matched, err := regexp.MatchString(constRegexForReplace, expected)
		if err != nil {
			rtnErrs = append(rtnErrs, err)
			return rtnErrs
		}

		if !matched {
			expected, err := v._loadExpectedValue(context, validationData.Expected)
			if err != nil {
				rtnErrs = append(rtnErrs, err)
				return rtnErrs
			}

			strValue := v._convertToString(expected)

			newErr := errors.New(strValue)
			(*validationData.Errors)[i] = newErr
			validationData.ErrorsReplaced[newErr] = true
			errorList = append(errorList, newErr)
		} else {
			replacer := strings.NewReplacer(constTagReplaceStart, "", constTagReplaceEnd, "")
			expected := replacer.Replace(validationData.Expected.(string))

			split := strings.SplitN(expected, ":", 2)
			if len(split) == 0 {
				rtnErrs = append(rtnErrs, ErrorInvalidTagArgument.Format(expected))
				continue
			}

			if _, ok := added[split[0]]; !ok {
				var arguments []interface{}
				if len(split) == 2 {
					splitArgs := strings.Split(split[1], constTagSplitValues)
					for _, arg := range splitArgs {
						arguments = append(arguments, arg)
					}
				}

				validationData.ErrorData = &errorData{
					Code:      split[0],
					Arguments: arguments,
				}

				newErr := v.errorCodeHandler(context, validationData)
				if newErr != nil {
					(*validationData.Errors)[i] = newErr
					validationData.ErrorsReplaced[newErr] = true
					errorList = append(errorList, newErr)
				}

				added[split[0]] = true
			}
		}
	}

	*validationData.Errors = errorList

	return rtnErrs
}
