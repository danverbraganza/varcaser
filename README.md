Varcaser
========

A library for converting variables between different programming language casing
conventions.

The varcaser transformation function `Caser.String(string)` is designed around
the assumption that the caller is aware of the casing convention of the input
variable names. If this is not the case, the caller can use the
`Detect([]string)` function on the input variable strings to retrieve the source
`CaseConvention` object, if that is possible. At present, this detection is
pretty naive, and will fail if all of the variables passed in are not all
formatted according to the same `CasingConvention`.

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

result := Caser{From: LowerCamelCase,
       To: ScreamingSnakeCase}.String("myConstantVariable")
// "MY_CONSTANT_VARIABLE"

fromCase := c, err := Detect([]string{"Abcd", "MyVar", "ThisIsMyVar"})
// CaseConvention{
//        JoinStyle: camelJoinStyle,
//        SubsequentCase: strings.Title,
//        InitialCase: strings.Title
// }

```

Available Case Conventions
--------------------------

All of the following are exported as CaseConvention structs.

* `LowerSnakeCase`: `lower_snake_case`
* `ScreamingSnakeCase`: `SCREAMING_SNAKE_CASE`
* `KebabCase`: `kebab-case`
* `ScreamingKebabCase`: `SCREAMING-KEBAB-CASE`
* `HttpHeaderCase`: `HTTP-Header-Case`  (NB: Mishandles some conventional acronyms at the moment)
* `UpperCamelCase`: `UpperCamelCase`  (renders HTTP as Http)
* `LowerCamelCase`: `lowerCamelCase`  (renders HTTP as Http)
* `UpperCamelCaseKeepCaps`: `UpperCamelCaseKeepCaps` (renders HTTP as HTTP)
* `LowerCamelCaseKeepCaps`: `lowerCamelCaseKeepCaps` (renders HTTP as HTTP)


Updates
-------

**2015-11-05**

Added the detector module, which provides the ability to detect the variable
casing conventions.


**2015-06-24**

Removing SpinalCase and TrainCase because the former makes me feel queasy and
they're both unnecessary.