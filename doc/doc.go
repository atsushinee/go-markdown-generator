package doc

import (
	"io/ioutil"
	"os"
	"strings"
)

const LevelTitle = 3
const LevelNormal = 5
const LevelWord = 6

type MarkDown struct {
	builder *strings.Builder
}

func NewMarkDown() *MarkDown {
	m := new(MarkDown)
	m.builder = new(strings.Builder)
	return m
}

func (m *MarkDown) PutLevel1Title(content string) *MarkDown {
	m.PutTitle(content, 1)
	return m
}

func (m *MarkDown) write(content string) {
	m.builder.WriteString(content)
}
func (m *MarkDown) PutTitle(content string, level int) *MarkDown {
	m.write(strings.Repeat("#", level) + " " + content)
	m.PutLine()
	return m
}

func (m *MarkDown) PutWordLine(content string) *MarkDown {
	m.PutWord(content)
	m.PutLine()
	return m
}

func (m *MarkDown) PutWord(content string) *MarkDown {
	m.write(content)
	return m
}
func (m *MarkDown) PutLine() *MarkDown {
	m.write("\n")
	return m
}

func (m *MarkDown) PutJson(content string) *MarkDown {
	m.write("``` json\n" + content + "\n```\n")
	return m
}

func (m *MarkDown) PutCodeLine(content string) *MarkDown {
	m.PutCode(content)
	m.PutLine()
	return m
}
func (m *MarkDown) PutCode(content string) *MarkDown {
	m.write("`" + content + "`")
	return m
}

func (m *MarkDown) PutTable(t *Table) *MarkDown {
	m.builder.WriteString(t.String())
	return m
}

func (m *MarkDown) Export(filename string) error {
	return ioutil.WriteFile(filename, []byte(m.builder.String()), os.ModePerm)
}

func (m *MarkDown) String() string {
	return m.builder.String()
}

type Table struct {
	body [][]string
}

func (t *Table) SetTitle(col int, content string) *Table {
	t.body[0][col] = content
	return t
}
func (t *Table) SetContent(row, col int, content string) *Table {
	row = row + 2
	t.body[row][col] = content
	return t
}

func (t *Table) String() string {
	var buffer strings.Builder
	for _, row := range t.body {
		buffer.WriteString("|")
		for _, col := range row {
			buffer.WriteString(col)
			buffer.WriteString("|")
		}
		buffer.WriteString("\n")

	}
	return buffer.String()
}

func NewTable(row, col int) *Table {
	t := new(Table)
	row = row + 2
	t.body = make([][]string, row)
	for i := 0; i < row; i++ {
		t.body[i] = make([]string, col)
		if i == 1 {
			for j := 0; j < col; j++ {
				t.body[i][j] = "----"
			}
		}
	}
	return t
}
