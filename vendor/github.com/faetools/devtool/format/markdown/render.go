package markdown

import (
	"bytes"
	"fmt"

	"github.com/faetools/devtool/format/writers"
	"github.com/yuin/goldmark"

	// meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/ast"
	extast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

func md() goldmark.Markdown {
	return goldmark.New(
		// later
		// goldmark.WithExtensions(extension.GFM), // meta.Meta,
		goldmark.WithRendererOptions(defaultRenderOpt()))
}

func defaultRenderOpt() renderer.Option {
	return renderer.WithNodeRenderers(util.Prioritized(NewRenderer(), defaultPriority))
}

func defaultParser() parser.Parser { return md().Parser() }

// Format formats Markdown.
func Format(src []byte, options ...renderer.Option) ([]byte, error) {
	return Render(src, defaultParser().Parse(text.NewReader(src)), options...)
}

// Render renders a given parsed document.
func Render(src []byte, doc ast.Node, options ...renderer.Option) ([]byte, error) {
	b := &bytes.Buffer{}
	w := writers.NewTrimWriter(b, sNewLine)

	opts := append([]renderer.Option{defaultRenderOpt()}, options...)

	err := renderer.NewRenderer(opts...).Render(w, src, doc)
	if err != nil {
		return nil, err
	}

	b.Write(newLine)

	// trim all new lines but add one at the end
	return b.Bytes(), nil
}

// A Renderer struct is an implementation of renderer.NodeRenderer that renders
// nodes as markdown.
type Renderer struct{ Config }

// NewRenderer returns a new Renderer with given options.
func NewRenderer(opts ...Option) renderer.NodeRenderer {
	r := &Renderer{Config: NewConfig()}

	for _, opt := range opts {
		opt.SetMarkdownOption(&r.Config)
	}

	return r
}

// RegisterFuncs implements NodeRenderer.RegisterFuncs .
func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	// blocks

	reg.Register(ast.KindHeading, r.renderHeading)
	reg.Register(ast.KindBlockquote, r.renderBlockquote)
	reg.Register(ast.KindCodeBlock, r.renderCodeBlock)
	reg.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
	reg.Register(ast.KindHTMLBlock, r.renderHTMLBlock)
	reg.Register(ast.KindList, r.renderList)
	reg.Register(ast.KindListItem, r.renderListItem)
	reg.Register(ast.KindParagraph, r.renderParagraph)
	reg.Register(ast.KindTextBlock, r.renderTextBlock)
	reg.Register(ast.KindThematicBreak, r.renderThematicBreak)
	reg.Register(extast.KindStrikethrough, r.renderStrikethrough)

	// inlines

	reg.Register(ast.KindAutoLink, r.renderAutoLink)
	reg.Register(ast.KindCodeSpan, r.renderCodeSpan)
	reg.Register(ast.KindEmphasis, r.renderEmphasis)
	reg.Register(ast.KindImage, r.renderImage)
	reg.Register(ast.KindLink, r.renderLink)
	reg.Register(ast.KindRawHTML, r.renderRawHTML)
	reg.Register(ast.KindText, r.renderText)
}

func (r *Renderer) renderHeading(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		_, _ = w.Write(twoNewLines)
		return ast.WalkContinue, nil
	}

	repeat(w, bHash, node.(*ast.Heading).Level)
	_ = w.WriteByte(bSpace)

	return ast.WalkContinue, nil
}

func (r *Renderer) renderBlockquote(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		_, _ = w.Write(twoNewLines)
		return ast.WalkContinue, nil
	}

	indent := getIndentOfListItem(n)
	if indent > 0 {
		// separate from the list with a new line
		_ = w.WriteByte(bNewLine)
	}

	// write all indents
	iw := writers.NewIndentWriter(w, indent)

	// write as blockquote
	bw := newBlockquoteWriter(iw)

	err := r.renderChildren(bw, source, n)
	if err != nil {
		return 0, err
	}

	return ast.WalkSkipChildren, nil
}

func (r *Renderer) renderCodeBlock(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	indent := getIndentOfListItem(n)
	if indent == 0 {
		// no indent, so we can transform into a fenced code block
		return r.renderFencedCodeBlock(w, source, n, entering)
	}

	if !entering {
		return ast.WalkContinue, nil
	}

	// separate from the rest
	_ = w.WriteByte(bNewLine)

	l := n.Lines().Len()
	lines := make([][]byte, l)
	for i := 0; i < l; i++ {
		line := n.Lines().At(i)
		lines[i] = line.Value(source)
	}

	// shift all lines to the left
	for hasSpacePrefix(lines) {
		lines = trimSpacePrefix(lines)
	}

	// write lines with indent
	iw := writers.NewIndentWriter(w, indent+1)
	for _, line := range lines {
		_, _ = iw.Write(line)
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_, _ = w.Write(codeBlockFence)

	if !entering {
		_, _ = w.Write(twoNewLines)
		return ast.WalkContinue, nil
	}

	lang := getLanguage(source, node)

	// e.g. "```go\n"
	_, _ = w.Write(lang)
	_ = w.WriteByte(bNewLine)

	writeFormattedLines(w, source, node, string(lang))

	return ast.WalkContinue, nil
}

func (r *Renderer) renderHTMLBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		// write the closure line, if any
		if cl := node.(*ast.HTMLBlock).ClosureLine; cl.Start > -1 {
			_, _ = w.Write(cl.Value(source))
		}

		_ = w.WriteByte(bNewLine)
		return ast.WalkContinue, nil
	}

	writeFormattedLines(w, source, node, langHTML)

	return ast.WalkContinue, nil
}

func (r *Renderer) renderList(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering && node.Parent().Kind() != ast.KindListItem {
		// two new lines at the end if not inside another list
		_, _ = w.Write(twoNewLines)
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderListItem(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	indent := getIndentOfListItem(node)
	if !entering {
		if // there is no next list item
		node.NextSibling() != nil &&
			// no content
			(node.LastChild() == nil ||
				// we have not rendered a blockquote
				node.LastChild().Kind() != ast.KindBlockquote) {
			_ = w.WriteByte(bNewLine)
		}

		return ast.WalkSkipChildren, nil
	}

	iw := writers.NewIndentWriter(w, indent)

	p := node.Parent()
	if !p.(*ast.List).IsOrdered() {
		_, _ = iw.Write(preUListItem)
	} else {
		_, _ = iw.WriteString(fmt.Sprintf(preOListItemFormat, numElement(node)))
	}

	if err := r.renderChildren(iw, source, node); err != nil {
		return 0, err
	}

	return ast.WalkSkipChildren, nil
}

func (r *Renderer) renderParagraph(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	indent := getIndentOfListItem(n)
	partOfList := indent > 0

	if !entering {
		if partOfList {
			if n.NextSibling() != nil {
				_ = w.WriteByte(bNewLine)
			}
		} else {
			_, _ = w.Write(twoNewLines)
		}

		return ast.WalkContinue, nil
	}

	if partOfList && n.PreviousSibling() != nil {
		_ = w.WriteByte(bNewLine)
		iw := writers.NewIndentWriter(w, indent)
		if err := r.renderChildren(iw, source, n); err != nil {
			return 0, err
		}
		_ = w.WriteByte(bNewLine)
	} else {
		if err := r.renderChildren(w, source, n); err != nil {
			return 0, err
		}
	}

	return ast.WalkSkipChildren, nil
}

func (r *Renderer) renderTextBlock(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering && n.NextSibling() != nil {
		_ = w.WriteByte(bNewLine)
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderThematicBreak(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		_, _ = w.Write(twoNewLines)
		return ast.WalkContinue, nil
	}

	_, _ = w.Write(thematicBreak)

	return ast.WalkContinue, nil
}

func (r *Renderer) renderAutoLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		_ = w.WriteByte(bGreaterThan)
		return ast.WalkContinue, nil
	}

	_ = w.WriteByte(bLessThan)

	n := node.(*ast.AutoLink)
	_, _ = w.Write(n.URL(source))

	return ast.WalkContinue, nil
}

func (r *Renderer) renderCodeSpan(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	// check if backticks need to be escaped
	if c := node.FirstChild(); c != nil &&
		bytes.Contains(c.(*ast.Text).Segment.Value(source), []byte{bGraveAccent}) {
		_ = w.WriteByte(bGraveAccent)
	}

	if !entering {
		_ = w.WriteByte(bGraveAccent)
		return ast.WalkContinue, nil
	}

	_ = w.WriteByte(bGraveAccent)

	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		segment := c.(*ast.Text).Segment
		_, _ = w.Write(segment.Value(source))
	}

	return ast.WalkSkipChildren, nil
}

func (r *Renderer) renderEmphasis(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Emphasis)

	if n.Attributes() != nil {
		panic(n.Attributes())
	}

	if r.Config.Terminal {
		if entering {
			switch n.Level {
			case 1:
				_, _ = w.Write(tItalic)
			case 2:
				_, _ = w.Write(tBold)
			}
		} else {
			_, _ = w.Write(tReset)
		}

		return ast.WalkContinue, nil
	}

	switch n.Level {
	case 1:
		_ = w.WriteByte(bAsterisk)
	case 2:
		_, _ = w.Write(twoAsterisks)
	default:
		panic(n.Level)
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Link)
	if entering {
		_ = w.WriteByte(bLeftSquareBracket)
		return ast.WalkContinue, nil
	}

	_, _ = w.Write(linkTransition)
	_, _ = w.Write(n.Destination) // link

	if len(n.Title) != 0 {
		_, _ = w.Write(linkTitleStart)
		_, _ = w.Write(n.Title)
		_ = w.WriteByte(bQuotationMark)
	}

	_ = w.WriteByte(bRightParenthesis)

	return ast.WalkContinue, nil
}

func (r *Renderer) renderImage(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.Image)
	_, _ = w.Write(imageStart)
	_, _ = w.Write(n.Text(source)) // alt
	_, _ = w.Write(linkTransition)
	_, _ = w.Write(n.Destination) // link

	if len(n.Title) != 0 {
		_, _ = w.Write(linkTitleStart)
		_, _ = w.Write(n.Title)
		_ = w.WriteByte(bQuotationMark)
	}

	_ = w.WriteByte(bRightParenthesis)
	return ast.WalkSkipChildren, nil
}

func (r *Renderer) renderRawHTML(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkSkipChildren, nil
	}

	_, _ = w.Write(node.Text(source))

	segs := node.(*ast.RawHTML).Segments
	for i, l := 0, segs.Len(); i < l; i++ {
		seg := segs.At(i)
		_, _ = w.Write(seg.Value(source))
	}

	return ast.WalkSkipChildren, nil
}

func (r *Renderer) renderText(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.Text)
	_, _ = w.Write(n.Segment.Value(source))

	if n.IsRaw() {
		return ast.WalkContinue, nil
	}

	if n.SoftLineBreak() {
		switch {
		case r.HardWraps:
			_, _ = w.Write(twoNewLines)
		default:
			_ = w.WriteByte(bNewLine)
		}
		return ast.WalkContinue, nil
	}

	if n.HardLineBreak() {
		_, _ = w.Write(twoNewLines)
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderStrikethrough(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	return ast.WalkContinue, nil
}
