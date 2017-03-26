# Variables, Expressions, and Templates

Pretty much any configuration key should accept variables. This will mean that every configuration key will probably have `string` as it's type. This should be fine for pretty much everywhere, as we can use a template engine of some kind (Go's text templates?) for subsituting variables at runtime.

The time when this is not as clear is when we're looking at `when`, because of it's expressions. Unless we can find a template language that is expressive enough to return a simple "true/false" string from some expression in an easy enough to read way, then we may have some problems here.

If we can't find something to do that, we may have to compile the expression from a template, and then run the expression. For example:

Template: `{{.resource.console-src.branch}} == 'master'`.
Expression: `'master' == 'master'`
Result: `true`

This makes sense, and should be reasonably fast, but it's not very elegant. This could be improved over time though, and isn't a high priority.

EDIT:
* Templates: https://github.com/flosch/pongo2
* Expressions: https://github.com/flosch/pongo2/blob/master/template_tests/expressions.tpl
