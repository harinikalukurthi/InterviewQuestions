1. what is struct?



The issue you're encountering stems from the fact that iterating over a map in Go doesn't guarantee any specific order. The order in which keys are retrieved from the map is random, which is why your output is not matching the expected order.

To ensure the output is in a consistent order, you should sort the keys before returning them. Here's an updated version of your Colors function: