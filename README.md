Varcaser
========

Varcaser is a library for converting variables between different programming
language casing conventions.

Varcaser is now stable. Although bugs that are discovered will be fixed, there
is no plan to add more features to Varcaser.

The varcaser transformation function `Caser.String(string)` is designed around
the assumption that the caller is aware of the casing convention of the input
variable names. If this is not the case, the caller can use the
`Detect([]string)` function on the input variable strings to retrieve a
`Splitter` object, if that is possible. This `Splitter` object takes care of
decomposing an input variable name into its component parts.

The case transformation component of Varcaser is implemented without regular
expressions.

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

In addition, it is easy to build a custom CaseConvention your own use, if you
need one that isn't provided here.

Updates
-------

**2015-11-07**

Changed how the detector module works. Instead of returning a CaseConvention
object, it returns a Splitter, which simplifies the logic significantly, with
only a small sacrifice in functionality.


**2015-11-05**

Added the detector module, which provides the ability to detect the variable
casing conventions.


**2015-06-24**

Removing SpinalCase and TrainCase because the former makes me feel queasy and
they're both unnecessary.