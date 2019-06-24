package response

import (
	"fmt"
)

func Example_ResponseError_MarshalJSON() {
	r := ResponseError{ErrUserNotExist, 500}
	result, _ := r.MarshalJSON()
	fmt.Println(string(result))
	//Output:
	//"user not exist"
}
func Example_ResponseError_Error() {
	r := ResponseError{ErrUserNotExist, 500}
	result := r.Error()
	fmt.Println(string(result))
	//Output:
	//
}
