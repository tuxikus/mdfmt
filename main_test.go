package main

import (
	"fmt"
	"reflect"
	"testing"
)

func dumpForTest(t *testing.T, want, got Node) {
	t.Error("Parse result does not match expected output")
	fmt.Println("=== want ===")
	dump(want.Children())
	fmt.Println("=== got ===")
	dump(got.Children())
}

func printFmtForTest(t *testing.T, want, got string, parsed Node) {
	t.Error("Want != got")
	fmt.Println("=== want ===")
	fmt.Println(want)
	fmt.Println("=== got ===")
	fmt.Println(got)
	dump(parsed.Children())
}

func TestParseEmptyDocument(t *testing.T) {
	input := ""
	want := &Document{}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseLevelOneHeading(t *testing.T) {
	input := "# Foo"
	want := &Document{
		children: []Node{
			&Heading{
				Level: 1,
				Text:  "Foo",
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseLevelTwoHeading(t *testing.T) {
	input := "## Foo Bar"
	want := &Document{
		children: []Node{
			&Heading{
				Level: 2,
				Text:  "Foo Bar",
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseLevel9Heading(t *testing.T) {
	input := "######### Foo Bar Baz"
	want := &Document{
		children: []Node{
			&Heading{
				Level: 9,
				Text:  "Foo Bar Baz",
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseSingleLineParagraph(t *testing.T) {
	input := "Foo"
	want := &Document{
		children: []Node{
			&Paragraph{
				Text: "Foo",
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseMultiLineParagraph(t *testing.T) {
	input := `Foo Faz
Bar Baz`
	want := &Document{
		children: []Node{
			&Paragraph{
				Text: "Foo Faz\nBar Baz",
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseHeadingParagraph(t *testing.T) {
	input := `# Heading
Foo Faz
Bar Baz`
	want := &Document{
		children: []Node{
			&Heading{
				Level: 1,
				Text:  "Heading",
			},
			&Paragraph{
				Text: "Foo Faz\nBar Baz",
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseHeadingParagraphTailingNewLines(t *testing.T) {
	input := `# Heading
Foo Faz
Bar Baz


`
	want := &Document{
		children: []Node{
			&Heading{
				Level: 1,
				Text:  "Heading",
			},
			&Paragraph{
				Text: "Foo Faz\nBar Baz",
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestParseHeadingTwoParagraphs(t *testing.T) {
	input := `# Heading
Foo Faz
Bar Baz

Second paragraph`
	want := &Document{
		children: []Node{
			&Heading{
				Level: 1,
				Text:  "Heading",
			},
			&Paragraph{
				Text: "Foo Faz\nBar Baz",
			},
			&Paragraph{
				Text: "Second paragraph",
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseList(t *testing.T) {
	input := "- Foo"
	want := &Document{
		children: []Node{
			&List{
				elements: []Node{
					&ListElement{
						Level: 1,
						Text:  "Foo",
					},
				},
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseListMoreHyphens(t *testing.T) {
	input := "- Foo-Bar-Baz"
	want := &Document{
		children: []Node{
			&List{
				elements: []Node{
					&ListElement{
						Level: 1,
						Text:  "Foo-Bar-Baz",
					},
				},
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseLongList(t *testing.T) {
	input := `- one
- two
- three
- four
- five
- six
- seven
- foo
- bar
- hello
- world`
	want := &Document{
		children: []Node{
			&List{
				elements: []Node{
					&ListElement{
						Level: 1,
						Text:  "one",
					},
					&ListElement{
						Level: 1,
						Text:  "two",
					},
					&ListElement{
						Level: 1,
						Text:  "three",
					},
					&ListElement{
						Level: 1,
						Text:  "four",
					},
					&ListElement{
						Level: 1,
						Text:  "five",
					},
					&ListElement{
						Level: 1,
						Text:  "six",
					},
					&ListElement{
						Level: 1,
						Text:  "seven",
					},
					&ListElement{
						Level: 1,
						Text:  "foo",
					},
					&ListElement{
						Level: 1,
						Text:  "bar",
					},
					&ListElement{
						Level: 1,
						Text:  "hello",
					},
					&ListElement{
						Level: 1,
						Text:  "world",
					},
				},
			},
		},
	}

	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseMultiLevelList(t *testing.T) {
	input := `- Foo
  - Bar`
	want := &Document{
		children: []Node{
			&List{
				elements: []Node{
					&ListElement{
						Level: 1,
						Text:  "Foo",
					},
					&ListElement{
						Level: 2,
						Text:  "Bar",
					},
				},
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseMultiLevelListLong(t *testing.T) {
	input := `- Foo
  - Bar
    - Baz
      - hello
        - world
          - test
		    - long
		      - list
		        - here`
	want := &Document{
		children: []Node{
			&List{
				elements: []Node{
					&ListElement{
						Level: 1,
						Text:  "Foo",
					},
					&ListElement{
						Level: 2,
						Text:  "Bar",
					},
					&ListElement{
						Level: 3,
						Text:  "Baz",
					},
					&ListElement{
						Level: 4,
						Text:  "hello",
					},
					&ListElement{
						Level: 5,
						Text:  "world",
					},
					&ListElement{
						Level: 6,
						Text:  "test",
					},
					&ListElement{
						Level: 7,
						Text:  "long",
					},
					&ListElement{
						Level: 8,
						Text:  "list",
					},
					&ListElement{
						Level: 9,
						Text:  "here",
					},
				},
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseMultiLevelListLongComplex(t *testing.T) {
	input := `- Foo
  - Bar
  - Bar2
  - longer element
    - Baz
      - hello
        - world
        - foo
        - foo
        - foo
        - foo
          - test
		    - long
		      - list
		        - here
		          - even
		            - more
		              - elements
		                - apple
		                  - banana
		                - coconut
		              - test
		              - abc
		              - i
		              - dont
		              - know`
	want := &Document{
		children: []Node{
			&List{
				elements: []Node{
					&ListElement{
						Level: 1,
						Text:  "Foo",
					},
					&ListElement{
						Level: 2,
						Text:  "Bar",
					},
					&ListElement{
						Level: 2,
						Text:  "Bar2",
					},
					&ListElement{
						Level: 2,
						Text:  "longer element",
					},
					&ListElement{
						Level: 3,
						Text:  "Baz",
					},
					&ListElement{
						Level: 4,
						Text:  "hello",
					},
					&ListElement{
						Level: 5,
						Text:  "world",
					},
					&ListElement{
						Level: 5,
						Text:  "foo",
					},
					&ListElement{
						Level: 5,
						Text:  "foo",
					},
					&ListElement{
						Level: 5,
						Text:  "foo",
					},
					&ListElement{
						Level: 5,
						Text:  "foo",
					},
					&ListElement{
						Level: 6,
						Text:  "test",
					},
					&ListElement{
						Level: 7,
						Text:  "long",
					},
					&ListElement{
						Level: 8,
						Text:  "list",
					},
					&ListElement{
						Level: 9,
						Text:  "here",
					},
					&ListElement{
						Level: 10,
						Text:  "even",
					},
					&ListElement{
						Level: 11,
						Text:  "more",
					},
					&ListElement{
						Level: 12,
						Text:  "elements",
					},
					&ListElement{
						Level: 13,
						Text:  "apple",
					},
					&ListElement{
						Level: 14,
						Text:  "banana",
					},
					&ListElement{
						Level: 13,
						Text:  "coconut",
					},
					&ListElement{
						Level: 12,
						Text:  "test",
					},
					&ListElement{
						Level: 12,
						Text:  "abc",
					},
					&ListElement{
						Level: 12,
						Text:  "i",
					},
					&ListElement{
						Level: 12,
						Text:  "dont",
					},
					&ListElement{
						Level: 12,
						Text:  "know",
					},
				},
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseHeadingWithList(t *testing.T) {
	input := `# Heading

- element 1
- element 2
- element 3
`

	want := &Document{
		children: []Node{
			&Heading{
				Level: 1,
				Text:  "Heading",
			},
			&List{
				elements: []Node{
					&ListElement{
						Level: 1,
						Text:  "element 1",
					},
					&ListElement{
						Level: 1,
						Text:  "element 2",
					},
					&ListElement{
						Level: 1,
						Text:  "element 3",
					},
				},
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseHeadingWithListNoNL(t *testing.T) {
	input := `# Heading
- element 1
- element 2
- element 3
`

	want := &Document{
		children: []Node{
			&Heading{
				Level: 1,
				Text:  "Heading",
			},
			&List{
				elements: []Node{
					&ListElement{
						Level: 1,
						Text:  "element 1",
					},
					&ListElement{
						Level: 1,
						Text:  "element 2",
					},
					&ListElement{
						Level: 1,
						Text:  "element 3",
					},
				},
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseHeadingWithListComplex(t *testing.T) {
	input := `# Heading

- element 1
- element 2
- element 3

## Next heading
- element 1
### The lvl 3 heading

- foo
  - bar
    - baz
`

	want := &Document{
		children: []Node{
			&Heading{
				Level: 1,
				Text:  "Heading",
			},
			&List{
				elements: []Node{
					&ListElement{
						Level: 1,
						Text:  "element 1",
					},
					&ListElement{
						Level: 1,
						Text:  "element 2",
					},
					&ListElement{
						Level: 1,
						Text:  "element 3",
					},
				},
			},
			&Heading{
				Level: 2,
				Text:  "Next heading",
			},
			&List{
				elements: []Node{
					&ListElement{
						Level: 1,
						Text:  "element 1",
					},
				},
			},
			&Heading{
				Level: 3,
				Text:  "The lvl 3 heading",
			},
			&List{
				elements: []Node{
					&ListElement{
						Level: 1,
						Text:  "foo",
					},
					&ListElement{
						Level: 2,
						Text:  "bar",
					},
					&ListElement{
						Level: 3,
						Text:  "baz",
					},
				},
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseTable(t *testing.T) {
	input := "| Table |"
	want := &Document{
		children: []Node{
			&Table{
				rows: []Node{
					&TableRow{
						elements: []Node{
							&TableElement{
								Text: "Table",
							},
						},
					},
				},
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseTableTwoByTwo(t *testing.T) {
	input := `| one | two |
| three | four |`
	want := &Document{
		children: []Node{
			&Table{
				rows: []Node{
					&TableRow{
						elements: []Node{
							&TableElement{
								Text: "one",
							},
							&TableElement{
								Text: "two",
							},
						},
					},
					&TableRow{
						elements: []Node{
							&TableElement{
								Text: "three",
							},
							&TableElement{
								Text: "four",
							},
						},
					},
				},
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseTableTwoByTwoMissingValue(t *testing.T) {
	input := `| one | |
| three | four |`
	want := &Document{
		children: []Node{
			&Table{
				rows: []Node{
					&TableRow{
						elements: []Node{
							&TableElement{
								Text: "one",
							},
							&TableElement{
								Text: "",
							},
						},
					},
					&TableRow{
						elements: []Node{
							&TableElement{
								Text: "three",
							},
							&TableElement{
								Text: "four",
							},
						},
					},
				},
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseTableTwoByTwoMissingValueMoreWhitespace(t *testing.T) {
	input := `| one |       |
| three | four |`
	want := &Document{
		children: []Node{
			&Table{
				rows: []Node{
					&TableRow{
						elements: []Node{
							&TableElement{
								Text: "one",
							},
							&TableElement{
								Text: "",
							},
						},
					},
					&TableRow{
						elements: []Node{
							&TableElement{
								Text: "three",
							},
							&TableElement{
								Text: "four",
							},
						},
					},
				},
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestParseTableWithHeading(t *testing.T) {
	input := `# Header
| table header a | table header b |
| ----- | ----- |
| element a | element b |`
	want := &Document{
		children: []Node{
			&Heading{
				Level: 1,
				Text:  "Header",
			},
			&Table{
				rows: []Node{
					&TableRow{
						elements: []Node{
							&TableElement{
								Text: "table header a",
							},
							&TableElement{
								Text: "table header b",
							},
						},
					},
					&TableRow{
						elements: []Node{
							&TableElement{
								Text: "-----",
							},
							&TableElement{
								Text: "-----",
							},
						},
					},
					&TableRow{
						elements: []Node{
							&TableElement{
								Text: "element a",
							},
							&TableElement{
								Text: "element b",
							},
						},
					},
				},
			},
		},
	}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		dumpForTest(t, want, got)
	}
}

func TestFmtHeadingWithParagraph(t *testing.T) {
	input := `# header
some text`
	want := `# header

some text`

	parsed := Parse(input)
	got := Fmt(Parse(input))

	if want != got {
		printFmtForTest(t, want, got, parsed)
	}
}

func TestFmtHeadingWithTwoParagraphs(t *testing.T) {
	input := `# header
some text




even more text`
	want := `# header

some text

even more text`

	parsed := Parse(input)
	got := Fmt(Parse(input))

	if want != got {
		printFmtForTest(t, want, got, parsed)
	}
}

func TestFmtHeadingWithLongParagraph(t *testing.T) {
	input := `# Heading

Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.`
	want := `# Heading

Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.`

	parsed := Parse(input)
	got := Fmt(Parse(input))

	if want != got {
		printFmtForTest(t, want, got, parsed)
	}
}

func TestFmtHeadingWithMoreLongParagraphs(t *testing.T) {
	input := `# Heading
Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.







Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.



Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.`
	want := `# Heading

Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.

Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.

Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.`

	parsed := Parse(input)
	got := Fmt(Parse(input))

	if want != got {
		printFmtForTest(t, want, got, parsed)
	}
}

func TestFmtHeadingWithWrappedParagraph(t *testing.T) {
	input := `# Heading
Lorem ipsum dolor sit amet consectetur adipiscing elit.
Quisque faucibus ex sapien vitae pellentesque sem
placerat. In id cursus mi pretium tellus duis convallis.
Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus
fringilla lacus nec metus bibendum egestas. Iaculis massa
nisl malesuada lacinia integer nunc posuere. Ut hendrerit
semper vel class aptent taciti sociosqu. Ad litora torquent
per conubia nostra inceptos himenaeos.`
	want := `# Heading

Lorem ipsum dolor sit amet consectetur adipiscing elit.
Quisque faucibus ex sapien vitae pellentesque sem
placerat. In id cursus mi pretium tellus duis convallis.
Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus
fringilla lacus nec metus bibendum egestas. Iaculis massa
nisl malesuada lacinia integer nunc posuere. Ut hendrerit
semper vel class aptent taciti sociosqu. Ad litora torquent
per conubia nostra inceptos himenaeos.`

	parsed := Parse(input)
	got := Fmt(Parse(input))

	if want != got {
		printFmtForTest(t, want, got, parsed)
	}
}

func TestFmtHeadingsWithParagraphs(t *testing.T) {
	input := `# header
some text

more text

## next heading
with a paragraph`
	want := `# header

some text

more text

## next heading

with a paragraph`

	parsed := Parse(input)
	got := Fmt(Parse(input))

	if want != got {
		printFmtForTest(t, want, got, parsed)
	}
}
