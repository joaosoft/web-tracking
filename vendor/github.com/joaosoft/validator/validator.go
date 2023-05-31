package validator

import (
	"github.com/joaosoft/logger"
)

func NewValidator() *Validator {

	v := &Validator{
		tag:       constDefaultValidationTag,
		callbacks: make(map[string]callbackHandler),
		sanitize:  make([]string, 0),
		logger:    logger.NewLogDefault(constDefaultLogTag, logger.InfoLevel),
	}

	v.init()

	return v
}

func (v *Validator) newActiveHandlers() map[string]empty {
	handlers := make(map[string]empty)

	for key, _ := range v.handlersBefore {
		handlers[key] = empty{}
	}

	for key, _ := range v.handlersMiddle {
		handlers[key] = empty{}
	}

	for key, _ := range v.handlersAfter {
		handlers[key] = empty{}
	}

	return handlers
}

func (v *Validator) AddBefore(name string, handler beforeTagHandler) *Validator {
	v.handlersBefore[name] = handler
	v.activeHandlers[name] = empty{}

	return v
}

func (v *Validator) AddMiddle(name string, handler middleTagHandler) *Validator {
	v.handlersMiddle[name] = handler
	v.activeHandlers[name] = empty{}

	return v
}

func (v *Validator) AddAfter(name string, handler afterTagHandler) *Validator {
	v.handlersAfter[name] = handler
	v.activeHandlers[name] = empty{}

	return v
}

func (v *Validator) SetErrorCodeHandler(handler errorCodeHandler) *Validator {
	v.errorCodeHandler = handler

	return v
}

func (v *Validator) SetValidateAll(canValidateAll bool) *Validator {
	v.canValidateAll = canValidateAll

	return v
}

func (v *Validator) SetTag(tag string) *Validator {
	v.tag = tag

	return v
}

func (v *Validator) SetPwdSettings(settings *PwdSettings) *Validator {
	v.pwd.settings = settings

	return v
}

func (v *Validator) SetSanitize(sanitize []string) *Validator {
	v.sanitize = sanitize

	return v
}

func (v *Validator) AddCallback(name string, callback callbackHandler) *Validator {
	v.callbacks[name] = callback

	return v
}

func (v *Validator) Validate(obj interface{}, args ...*argument) []error {
	return NewValidatorHandler(v, args...).handleValidation(obj)
}
