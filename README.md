Varcaser
========

A library for converting variables between different programming language casing
conventions.

Varcaser is designed around the assumption that the caller is aware of the casing
convention of the input variable names. It is a planned feature to have a casing
detector.

Varcaser is implemented without regular expressions (for now).

Varcaser handles cases such as rendering "AsyncHTTPRequest" in lower_snake_case
as "async_http_request". In {Upper, Lower}CamelCase, that same variable will be
rendered as "AsyncHttpRequest". To preserve the original casing, use
{Upper, Lower}CamelCaseKeepCaps.

**Warning**: Although varcaser.Caser implements the golang.org/x/text/transform
  interface, the Bytes() and Transform() methods have not been tested yet.

Usage Examples
--------------
```
result := Caser{From: LowerCamelCase, To: KebabCase}.String("someInitMethod")
// "some-init-method"

result := Caser{From: LowerCamelCase, To: ScreamingSnakeCase}.String("myConstantVariable")
// MY_CONSTANT_VARIABLE
```

Available Case Conventions
--------------------------

All of the following are exported as CaseConvention structs.

* **LowerSnakeCase**: lower_snake_case
* **ScreamingSnakeCase**: SCREAMING_SNAKE_CASE
* **KebabCase**: kebab-case  (also exported as **SpinalCase**)
* **UpperKebabCase**: Upper-Kebab-Case (also exported as **TrainCase**)
* **ScreamingKebabCase**: SCREAMING-KEBAB-CASE
* **HttpHeaderCase**: HTTP-Header-Case  (NB: Mishandles some conventional acronyms at the moment)
* **UpperCamelCase**: UpperCamelCase  (renders HTTP as Http)
* **LowerCamelCase**: lowerCamelCase  (renders HTTP as Http)
* **UpperCamelCaseKeepCaps**: UpperCamelCaseKeepCaps (renders HTTP as HTTP)
* **LowerCamelCaseKeepCaps**: lowerCamelCaseKeepCaps (renders HTTP as HTTP)
