Validator (go-playground/validator)

Purpose:
- Validate struct fields using rules written in struct tags.

How it works:
1. Receive any struct (interface{}).
2. Use reflection to inspect every field.
3. Read the `validate` struct tag.
4. Parse validation rules (required, email, min, max, etc.).
5. Apply rules to the field values.
6. Return validation errors if any rule fails.

Important:
- Validator reads only the `validate` tag.
- JSON decoding is done separately using the `json` tag.
- Both libraries use Go's reflection, but each reads different tags.

Mental Model:
Struct
    ↓
Reflection
    ↓
Read tags
    ↓
Execute validation rules
    ↓
Return errors

-----
Reflection standard defination: Reflection is the ability of a program to inspect and manipulate its own types, values, and structure at runtime.

Reflection defination: Go can look inside any variable while the program is running and discover everything about it—even if it didn't know its type when the program was compiled.

----
Reflection

Definition:
Reflection allows a Go program to inspect its own types, values, fields and tags at runtime.

Why it exists:
Libraries don't know the user's structs beforehand. Reflection lets them examine any struct dynamically.

How it works:
Variable
    ↓
reflect.TypeOf()  → type information
reflect.ValueOf() → actual values
    ↓
Read fields
Read tags
Read values
Perform required operation

Used by:
- encoding/json
- validator
- GORM
- YAML/XML packages
- Configuration libraries

Mental Model:
Reflection is an X-ray machine for Go objects. It lets a program look inside any variable while it's running.