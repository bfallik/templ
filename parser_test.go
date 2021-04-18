package templ

import (
	"testing"

	"github.com/a-h/lexical/input"
	"github.com/a-h/lexical/parse"
	"github.com/google/go-cmp/cmp"
)

func TestPackageParserErrors(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected parseError
	}{
		{
			name:  "unterminated package",
			input: "{% package",
			expected: newParseError(
				"package literal not terminated",
				Position{
					Index: 0,
					Line:  1,
					Col:   0,
				},
				Position{
					Index: 10,
					Line:  1,
					Col:   10,
				},
			),
		},
		{
			name:  "unterminated package, new line",
			input: "{% package \n%}",
			expected: newParseError(
				"package literal not terminated",
				Position{
					Index: 0,
					Line:  1,
					Col:   0,
				},
				Position{
					Index: 10,
					Line:  1,
					Col:   10,
				},
			),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			pi := input.NewFromString(tt.input)
			actual := packageParser{}.Parse(pi)
			if actual.Success {
				t.Errorf("expected parsing to fail, but it succeeded")
			}
			if diff := cmp.Diff(tt.expected, actual.Error); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestParsers(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		parser   parse.Function
		expected interface{}
	}{
		{
			name:   "package: standard",
			input:  `{% package templ %}`,
			parser: packageParser{}.Parse,
			expected: Package{
				Expression: "templ",
			},
		},
		{
			name:   "import: named",
			input:  `{% import name "github.com/a-h/something" %}`,
			parser: importParser,
			expected: Import{
				Expression: `name "github.com/a-h/something"`,
			},
		},
		{
			name:   "import: default",
			input:  `{% import "github.com/a-h/something" %}`,
			parser: importParser,
			expected: Import{
				Expression: `"github.com/a-h/something"`,
			},
		},
		{
			name:   "templateFileWhitespace: various spaces",
			input:  "  \t ",
			parser: whitespaceParser,
			expected: Whitespace{
				Value: "  \t ",
			},
		},
		{
			name:   "whitespace: spaces and newline",
			input:  " \n",
			parser: whitespaceParser,
			expected: Whitespace{
				Value: " \n",
			},
		},
		{
			name:   "whitespace: newline",
			input:  "\n",
			parser: whitespaceParser,
			expected: Whitespace{
				Value: "\n",
			},
		},
		{
			name: "template: no parameters",
			input: `{% templ Name() %}
{% endtmpl %}`,
			parser: templateParser,
			expected: Template{
				Name:                "Name",
				ParameterExpression: "",
				Children:            []Node{},
			},
		},
		{
			name: "template: single parameter",
			input: `{% templ Name(p Parameter) %}
{% endtmpl %}`,
			parser: templateParser,
			expected: Template{
				Name:                "Name",
				ParameterExpression: "p Parameter",
				Children:            []Node{},
			},
		},
		{
			name:   "constant attribute",
			input:  ` href="test"`,
			parser: constAttrParser,
			expected: ConstantAttribute{
				Name:  "href",
				Value: "test",
			},
		},
		{
			name:   "element: self-closing with single constant attribute",
			input:  `<a href="test"/>`,
			parser: elementSelfClosingParser,
			expected: elementSelfClosing{
				Name: "a",
				Attributes: []Attribute{
					ConstantAttribute{
						Name:  "href",
						Value: "test",
					},
				},
			},
		},
		{
			name:   "element: self-closing with single expression attribute",
			input:  `<a href={%= "test" %}/>`,
			parser: elementSelfClosingParser,
			expected: elementSelfClosing{
				Name: "a",
				Attributes: []Attribute{
					ExpressionAttribute{
						Name: "href",
						Value: StringExpression{
							Expression: `"test"`,
						},
					},
				},
			},
		},
		{
			name:   "element: self-closing with multiple constant attributes",
			input:  `<a href="test" style="text-underline: auto"/>`,
			parser: elementSelfClosingParser,
			expected: elementSelfClosing{
				Name: "a",
				Attributes: []Attribute{
					ConstantAttribute{
						Name:  "href",
						Value: "test",
					},
					ConstantAttribute{
						Name:  "style",
						Value: "text-underline: auto",
					},
				},
			},
		},
		{
			name:   "element: self-closing with multiple constant and expr attributes",
			input:  `<a href="test" title={%= localisation.Get("a_title") %} style="text-underline: auto"/>`,
			parser: elementSelfClosingParser,
			expected: elementSelfClosing{
				Name: "a",
				Attributes: []Attribute{
					ConstantAttribute{
						Name:  "href",
						Value: "test",
					},
					ExpressionAttribute{
						Name: "title",
						Value: StringExpression{
							`localisation.Get("a_title")`,
						},
					},
					ConstantAttribute{
						Name:  "style",
						Value: "text-underline: auto",
					},
				},
			},
		},
		{
			name:   "element: self closing with no attributes",
			input:  `<hr/>`,
			parser: elementSelfClosingParser,
			expected: elementSelfClosing{
				Name:       "hr",
				Attributes: []Attribute{},
			},
		},
		{
			name:   "element: self closing with attribute",
			input:  `<hr style="padding: 10px" />`,
			parser: elementSelfClosingParser,
			expected: elementSelfClosing{
				Name: "hr",
				Attributes: []Attribute{
					ConstantAttribute{
						Name:  "style",
						Value: "padding: 10px",
					},
				},
			},
		},
		{
			name:   "element: open",
			input:  `<a>`,
			parser: elementOpenTagParser,
			expected: elementOpenTag{
				Name:       "a",
				Attributes: []Attribute{},
			},
		},
		{
			name:   "element: open with attributes",
			input:  `<div id="123" style="padding: 10px">`,
			parser: elementOpenTagParser,
			expected: elementOpenTag{
				Name: "div",
				Attributes: []Attribute{
					ConstantAttribute{
						Name:  "id",
						Value: "123",
					},
					ConstantAttribute{
						Name:  "style",
						Value: "padding: 10px",
					},
				},
			},
		},
		{
			name:   "element: open and close",
			input:  `<a></a>`,
			parser: elementParser{}.Parse,
			expected: Element{
				Name:       "a",
				Attributes: []Attribute{},
			},
		},
		{
			name:   "element: with self-closing child element",
			input:  `<a><b/></a>`,
			parser: elementParser{}.Parse,
			expected: Element{
				Name:       "a",
				Attributes: []Attribute{},
				Children: []Node{
					Element{
						Name:       "b",
						Attributes: []Attribute{},
					},
				},
			},
		},
		{
			name:   "element: with non-self-closing child element",
			input:  `<a><b></b></a>`,
			parser: elementParser{}.Parse,
			expected: Element{
				Name:       "a",
				Attributes: []Attribute{},
				Children: []Node{
					Element{
						Name:       "b",
						Attributes: []Attribute{},
					},
				},
			},
		},
		{
			name:   "element: containing space",
			input:  `<a> <b> </b> </a>`,
			parser: elementParser{}.Parse,
			expected: Element{
				Name:       "a",
				Attributes: []Attribute{},
				Children: []Node{
					Element{
						Name:       "b",
						Attributes: []Attribute{},
					},
				},
			},
		},
		{
			name:   "element: with multiple child elements",
			input:  `<a><b></b><c><d/></c></a>`,
			parser: elementParser{}.Parse,
			expected: Element{
				Name:       "a",
				Attributes: []Attribute{},
				Children: []Node{
					Element{
						Name:       "b",
						Attributes: []Attribute{},
					},
					Element{
						Name:       "c",
						Attributes: []Attribute{},
						Children: []Node{
							Element{
								Name:       "d",
								Attributes: []Attribute{},
							},
						},
					},
				},
			},
		},
		{
			name:   "nodestringexpression: constant",
			input:  `{%= "test" %}`,
			parser: stringExpressionParser,
			expected: StringExpression{
				Expression: `"test"`,
			},
		},
		{
			name:   "element: containing string expression",
			input:  `<div>{%= "test" %}</div>`,
			parser: elementParser{}.Parse,
			expected: Element{
				Name:       "div",
				Attributes: []Attribute{},
				Children: []Node{
					StringExpression{
						Expression: `"test"`,
					},
				},
			},
		},
		{
			name: "template: containing element",
			input: `{% templ Name(p Parameter) %}
<span>{%= "span content" %}</span>
{% endtmpl %}`,
			parser: templateParser,
			expected: Template{
				Name:                "Name",
				ParameterExpression: "p Parameter",
				Children: []Node{
					Element{
						Name:       "span",
						Attributes: []Attribute{},
						Children: []Node{
							StringExpression{
								Expression: `"span content"`,
							},
						},
					},
					Whitespace{
						Value: "\n",
					},
				},
			},
		},
		{
			name: "template: containing nested elements",
			input: `{% templ Name(p Parameter) %}
<div>
  {%= "div content" %}
  <span>
	{%= "span content" %}
  </span>
</div>
{% endtmpl %}`,
			parser: templateParser,
			expected: Template{
				Name:                "Name",
				ParameterExpression: "p Parameter",
				Children: []Node{
					Element{
						Name:       "div",
						Attributes: []Attribute{},
						Children: []Node{
							StringExpression{Expression: `"div content"`},
							Element{
								Name:       "span",
								Attributes: []Attribute{},
								Children: []Node{
									StringExpression{Expression: `"span content"`},
								},
							},
						},
					},
					Whitespace{
						Value: "\n",
					},
				},
			},
		},
		{
			name: "if: simple expression",
			input: `{% if p.Test %}
<span>
  {%= "span content" %}
</span>
{% endif %}
`,
			parser: ifExpressionParser{}.Parse,
			expected: IfExpression{
				Expression: `p.Test`,
				Then: []Node{
					Element{
						Name:       "span",
						Attributes: []Attribute{},
						Children: []Node{
							StringExpression{Expression: `"span content"`},
						},
					},
					Whitespace{
						Value: "\n",
					},
				},
				Else: []Node{},
			},
		},
		{
			name: "if: else",
			input: `{% if p.A %}
	{%= "A" %}
{% else %}
	{%= "B" %}
{% endif %}`,
			parser: ifExpressionParser{}.Parse,
			expected: IfExpression{
				Expression: `p.A`,
				Then: []Node{
					Whitespace{Value: "\t"},
					StringExpression{Expression: `"A"`},
					Whitespace{Value: "\n"},
				},
				Else: []Node{
					Whitespace{Value: "\n\t"},
					StringExpression{Expression: `"B"`},
					Whitespace{Value: "\n"},
				},
			},
		},
		{
			name: "if: nested",
			input: `{% if p.A %}
					{% if p.B %}
						<div>{%= "B" %}</div>
					{% endif %}
				{% endif %}`,
			parser: ifExpressionParser{}.Parse,
			expected: IfExpression{
				Expression: `p.A`,
				Then: []Node{
					Whitespace{Value: "\t\t\t\t\t"},
					IfExpression{
						Expression: `p.B`,
						Then: []Node{
							Whitespace{Value: "\t\t\t\t\t\t"},
							Element{
								Name:       "div",
								Attributes: []Attribute{},
								Children: []Node{
									StringExpression{Expression: `"B"`},
								},
							},
							Whitespace{Value: "\n\t\t\t\t\t"},
						},
						Else: []Node{},
					},
					Whitespace{Value: "\n\t\t\t\t"},
				},
				Else: []Node{},
			},
		},
		{
			name: "template: containing if element",
			input: `{% templ Name(p Parameter) %}
	{% if p.Test %}
		<span>
			{%= "span content" %}
		</span>
	{% endif %}
{% endtmpl %}`,
			parser: templateParser,
			expected: Template{
				Name:                "Name",
				ParameterExpression: "p Parameter",
				Children: []Node{
					Whitespace{Value: "\t"},
					IfExpression{
						Expression: `p.Test`,
						Then: []Node{
							Whitespace{Value: "\t\t"},
							Element{
								Name:       "span",
								Attributes: []Attribute{},
								Children: []Node{
									StringExpression{Expression: `"span content"`},
								},
							},
							Whitespace{
								Value: "\n\t",
							},
						},
						Else: []Node{},
					},
					Whitespace{
						Value: "\n",
					},
				},
			},
		},
		{
			name: "for: simple",
			input: `{% for _, item := range p.Items %}
					<div>{%= item %}</div>
				{% endfor %}`,
			parser: forExpressionParser{}.Parse,
			expected: ForExpression{
				Expression: `_, item := range p.Items`,
				Children: []Node{
					Whitespace{Value: "\t\t\t\t\t"},
					Element{
						Name:       "div",
						Attributes: []Attribute{},
						Children: []Node{
							StringExpression{Expression: `item`},
						},
					},
					Whitespace{Value: "\n\t\t\t\t"},
				},
			},
		},
		{
			name:   "call: simple",
			input:  `{% call Other(p.Test) %}`,
			parser: callTemplateExpressionParser{}.Parse,
			expected: CallTemplateExpression{
				Name:               "Other",
				ArgumentExpression: `p.Test`,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			input := input.NewFromString(tt.input)
			result := tt.parser(input)
			if result.Error != nil {
				t.Fatalf("paser error: %v", result.Error)
			}
			if !result.Success {
				t.Fatalf("failed to parse at %d", input.Index())
			}
			if diff := cmp.Diff(tt.expected, result.Item); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}