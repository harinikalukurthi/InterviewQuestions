Context will carry deadlines




Creation of context:
context.TODO()
context.Background()


we can send signals to stop program using the following:

context.WithValue()
context.WithCancel()
context.WithTimeout()
context.WithDeadline()
