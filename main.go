package main

import (
	"strings"
)

// Markdown elements important for formatting:
// Document:
//
// # Heading(level: 1, text: Heading)
// Paragraph(text)
//
// ## Heading(level, text)
// Paragraph(text)
//
// ### Heading(level, text)
// (list)
// - ListElement(level: 1, text: ListElement)
// - ListElement
// - ListElement
//
// * ListElement
// * ListElement
// * ListElement
//
// - ListElement
//   - ListElement
//     - ListElement(level: 3, text: ListElement)
//     - ListElement
//
//
// #### Heading
// | Table | Table |
// | ----- | ----- |
// | Table | Table |
//
//

const TabWidth = 4

type NodeType int

const (
	NodeTypeDocument = iota
	NodeTypeHeading
	NodeTypeParagraph
	NodeTypeList
	NodeTypeListElement
)

type Node interface {
	Type() NodeType
	Children() []Node
}

var _ Node = (*Document)(nil)

type Document struct {
	children []Node
}

func (d *Document) Type() NodeType   { return NodeTypeDocument }
func (d *Document) Children() []Node { return d.children }

var _ Node = (*Heading)(nil)

type Heading struct {
	Level int
	Text  string
}

func (h *Heading) Type() NodeType   { return NodeTypeHeading }
func (h *Heading) Children() []Node { return nil }

var _ Node = (*List)(nil)

type List struct {
	elements []Node
}

func (l *List) Type() NodeType   { return NodeTypeList }
func (l *List) Children() []Node { return l.elements }

var _ Node = (*ListElement)(nil)

type ListElement struct {
	Level int
	Text  string
}

func (le *ListElement) Type() NodeType   { return NodeTypeListElement }
func (le *ListElement) Children() []Node { return nil }

var _ Node = (*Paragraph)(nil)

type Paragraph struct {
	Text string
}

func (p *Paragraph) Type() NodeType   { return NodeTypeParagraph }
func (p *Paragraph) Children() []Node { return nil }

func Parse(in string) Node {
	doc := &Document{}

	lines := strings.Split(in, "\n")

	for i := 0; i < len(lines); i++ {
		// skip empty lines
		if strings.TrimSpace(lines[i]) == "" {
			continue
		}

		// heading
		if strings.HasPrefix(lines[i], "#") {
			// get heading lvl
			lvl := 0
			textStart := 0
			for j := range len(lines[i]) {
				if lines[i][j] != '#' {
					textStart = j + 1
					break
				}
				lvl++
			}

			text := lines[i][textStart:]

			doc.children = append(doc.children, &Heading{
				Level: lvl,
				Text:  text,
			})

			continue
		}

		// list
		listStart := i
		// trim all left spaces and check if - is first char
		// - list element at lvl 1
		//   - list element at lvl 2
		//   - list element at lvl 2
		//     - list element at lvl 3
		if strings.HasPrefix(strings.TrimSpace(lines[i]), "-") {
			for i < len(lines) && strings.HasPrefix(strings.TrimSpace(lines[i]), "-") {
				i++
			}

			listElements := make([]Node, 0)
			listLines := lines[listStart:i]

			// replace tabs
			for j := range listLines {
				listLines[j] = strings.ReplaceAll(listLines[j], "\t", strings.Repeat(" ", TabWidth))
			}

			for _, listLine := range listLines {
				lvl := 1
				for j := range listLine {
					if listLine[j] == '-' {
						break
					}
					lvl++
				}
				// TODO: use j as text start, no need for trimming
				text := strings.TrimSpace(listLine) // needed for higher lvl list elements
				text = strings.TrimLeft(text, "-")
				text = strings.TrimSpace(text)

				lvl = lvl/2 + 1

				listElements = append(listElements, &ListElement{
					Level: lvl,
					Text:  text,
				})
			}

			doc.children = append(doc.children, &List{
				elements: listElements,
			})

			continue
		}

		// paragraph
		paragraphStart := i
		for i < len(lines) && strings.TrimSpace(lines[i]) != "" {
			i++
		}

		doc.children = append(doc.children, &Paragraph{
			Text: strings.Join(lines[paragraphStart:i], "\n"),
		})
	}

	return doc
}

func main() {

}
