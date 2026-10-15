In Go, comparable types are data types whose values can be compared using the equality (==) and inequality (!=) operators, and which can be used as keys in maps.


What types ARE comparable in Go?

• Booleans: bool (true == true)
• Numeric types: All integer (int, uint8, float64, etc.) and floating-point numbers
• Strings: string values
• Pointers: Pointer variables compare equal if they point to the same memory address or are both nil
• Channels: Channel values compare equal if they were created by the same make call or are both nil
• Arrays: An array is comparable if and only if its element type is comparable. Two arrays are equal if their corresponding elements are equal.
• Structs: A struct is comparable if and only if all of its individual fields are comparable. Two structs are equal if all their corresponding fields are equal.
• Interfaces: Interface values are comparable if their dynamic/underlying types are comparable. (Note: Comparing interfaces with non-comparable underlying types at runtime will cause a panic).



What types are NOT comparable in Go?

• Slices: []T cannot be compared with == (except against nil).
• Maps: map[K]V cannot be compared with == (except against nil).
• Functions: func() cannot be compared with == (except against nil).
• Structs or Arrays containing any of the above: If a struct has a slice or map field, the entire struct becomes completely non-comparable and cannot be used as a map key.


Important Interview Distinction: comparable vs cmp.Ordered

• comparable (built-in constraint): Only checks for equality (==, !=) and allows usage as map keys. It includes numbers, strings, booleans, pointers, channels, and valid structs/arrays.
• cmp.Ordered (from the standard cmp package): Checks for relational ordering (<, <=, >, >=). It only supports numbers and strings. Pointers and structs satisfy comparable but do not satisfy cmp.Ordered because they cannot be sorted.