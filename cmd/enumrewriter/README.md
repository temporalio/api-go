### Enum Rewriter

This tool exists to rewrite the `String()` methods of all proto-generated enums to be compatible with temporal's current CLI, UI, etc APIs. It walks the AST of generated files and replaces the contents of all String() methods define on enumerated types with a switch statement returning a "shorthand" value: given an enum like `ENCODING_TYPE_PROTO3` this will rewrite the string method to return `Proto3`.
