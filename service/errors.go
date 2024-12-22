package service

import "fmt"

var ErrUniqueConstraintViolation = fmt.Errorf("unique constraint violation")
var ErrForeigKeyViolationt = fmt.Errorf("foreign key violation")
var ErrUnaccountedDBError = fmt.Errorf("unaccounted DB error")
