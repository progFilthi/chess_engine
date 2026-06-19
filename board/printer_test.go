package board

import (
	"strings"
	"testing"
)

func TestPrinter(t *testing.T) {

	b := NewBoard()

	builder := strings.Builder{}

	b.Print(&builder)

	if !strings.Contains(builder.String(), "*") {
		t.Errorf("expected * for empty square, got %v", builder.String())
	}

	if !strings.Contains(builder.String(), "8") {
		t.Errorf("expected 8 for the 8th rank, got %v", builder.String())
	}

	if !strings.Contains(builder.String(), "1") {
		t.Errorf("expected 8 for the 1st rank, got %v", builder.String())
	}

	if !strings.Contains(builder.String(), "a") {
		t.Errorf("expected a for files, got %v", builder.String())
	}

}
