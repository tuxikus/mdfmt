package main

import (
	"reflect"
	"testing"
)

func TestParseEmptyDocument(t *testing.T) {
	input := ""
	want := &Document{}
	got := Parse(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %s, got %s", want, got)
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
		t.Errorf("want %s, got %s", want, got)
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
		t.Errorf("want %s, got %s", want, got)
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
		t.Errorf("want %s, got %s", want, got)
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
		t.Errorf("want %s, got %s", want, got)
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
		t.Errorf("want %s, got %s", want, got)
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
		t.Errorf("want %s, got %s", want, got)
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
		t.Errorf("want %s, got %s", want, got)
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
		t.Errorf("want %s, got %s", want, got)
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
		t.Errorf("want %s, got %s", want, got)
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
		t.Errorf("want %s, got %s", want, got)
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
		t.Errorf("want %s, got %s", want, got)
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
		t.Errorf("want %s, got %s", want, got)
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
		t.Errorf("want %s, got %s", want, got)
	}
}
