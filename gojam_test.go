package gojam

import "testing"

func TestRunThis(t *testing.T) {
    my_message := "lollllzzzz"
    res := RunThis(my_message)
    if res != my_message {
        t.Error("errorrrrrrrrr")
    }
}