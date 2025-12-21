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
