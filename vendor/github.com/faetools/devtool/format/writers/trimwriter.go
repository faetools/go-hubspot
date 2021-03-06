package writers

import (
	"bytes"
	"io"
	"strings"
)

// NewTrimWriter returns a new writer that trims.
func NewTrimWriter(base io.Writer, cutset string) Writer {
	if cutset == "" {
		return Upgrade(base)
	}

	return NewTrimLeftWriter(NewTrimRightWriter(base, cutset), cutset)
}

// NewTrimLeftWriter returns a new writer that trims from the left.
func NewTrimLeftWriter(base io.Writer, cutset string) Writer {
	if cutset == "" {
		return Upgrade(base)
	}

	tlw := &trimLeftWriter{base: base, cutset: cutset}
	tlw.Writer = NewStateWriter(tlw.ltwStateStart)

	return tlw
}

// NewTrimRightWriter returns a new writer that trims from the right.
func NewTrimRightWriter(base io.Writer, cutset string) Writer {
	if cutset == "" {
		return Upgrade(base)
	}

	return &trimRightWriter{
		buff:   &bytes.Buffer{},
		cutset: cutset,
		w:      Upgrade(base),
	}
}

type trimLeftWriter struct {
	Writer // the underlying state writer

	base   io.Writer // the base writer
	cutset string
}

func (w *trimLeftWriter) ltwStateStart(p []byte) (StateFunc, []byte, int, error) {
	lenNotInCutset := len(bytes.TrimLeft(p, w.cutset))
	if lenNotInCutset == 0 {
		// there still might need to be more trimmed
		return w.ltwStateStart, nil, len(p), nil
	}

	// we're done trimming
	lenDiscarded := len(p) - lenNotInCutset
	return w.stateJustWrite, p[lenDiscarded:], lenDiscarded, nil
}

func (w *trimLeftWriter) stateJustWrite(p []byte) (StateFunc, []byte, int, error) {
	size, err := w.base.Write(p)
	return w.stateJustWrite, p[size:], size, err
}

type trimRightWriter struct {
	buff   *bytes.Buffer
	cutset string
	w      Writer
}

func (w trimRightWriter) Write(p []byte) (int, error) {
	notInCutset := bytes.TrimRight(p, w.cutset)
	lenNotInCutset := len(notInCutset)

	if lenNotInCutset > 0 {
		// we're not at the end, write the buffer content first
		if _, err := io.Copy(w.w, w.buff); err != nil {
			return 0, err
		}
	}

	l := len(p)
	if lenNotInCutset == l {
		// nothing was trimmed, write normally
		return w.w.Write(p)
	}

	// write to underlying writer first
	if size, err := w.w.Write(notInCutset); err != nil {
		return size, err
	}

	// save the rest for later - we may or may not choose to write it
	_, _ = w.buff.Write(p[lenNotInCutset:]) // error always nil
	return l, nil
}

func (w *trimRightWriter) WriteByte(b byte) error {
	r := rune(b)
	for _, c := range w.cutset {
		if c == r {
			// save for later - we may or may not choose to write it
			return w.buff.WriteByte(b)
		}
	}

	// we're not at the end, write the buffer content first
	if _, err := io.Copy(w.w, w.buff); err != nil {
		return err
	}

	// nothing was trimmed, write normally
	return w.w.WriteByte(b)
}

func (w *trimRightWriter) WriteString(s string) (int, error) {
	notInCutset := strings.TrimRight(s, w.cutset)

	lenNotInCutset := len(notInCutset)

	if lenNotInCutset > 0 {
		// we're not at the end, write the buffer content first
		if _, err := io.Copy(w.w, w.buff); err != nil {
			return 0, err
		}
	}

	l := len(s)
	if lenNotInCutset == l {
		// nothing was trimmed, write normally
		return w.w.WriteString(s)
	}

	// write to underlying writer first
	if size, err := w.w.WriteString(notInCutset); err != nil {
		return size, err
	}

	// save the rest for later - we may or may not choose to write it
	_, err := w.buff.WriteString(s[lenNotInCutset:])
	return l, err
}
