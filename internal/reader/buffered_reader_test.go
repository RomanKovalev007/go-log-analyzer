package reader

import (
	"os"
	"strings"
	"testing"
)

func TestNewReader(t *testing.T) {
	size := 8192
	reader := NewReader(size)
	if reader.size != size {
		t.Errorf("Expected buffer size %d, got %d", size, reader.size)
	}
}

func TestReadFile_EmptyFile(t *testing.T) {
	// Создаем временный файл
	tmpfile, err := os.CreateTemp("", "empty_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	r := NewReader(4096)
	called := false

	name := tmpfile.Name()

	err = r.ReadFile(&name, func(line string) error {
		called = true
		return nil
	})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if called {
		t.Error("Processor should not be called for empty file")
	}
}

func TestReadFile_NormalCase(t *testing.T) {
	content := "line1\nline2\nline3\n"
	tmpfile, err := os.CreateTemp("", "normal_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	r := NewReader(1024)
	var lines []string

	name := tmpfile.Name()

	err = r.ReadFile(&name, func(line string) error {
		lines = append(lines, line)
		return nil
	})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []string{"line1", "line2", "line3"}
	if len(lines) != len(expected) {
		t.Fatalf("Expected %d lines, got %d", len(expected), len(lines))
	}

	for i := range expected {
		if lines[i] != expected[i] {
			t.Errorf("Line %d mismatch: expected '%s', got '%s'", i, expected[i], lines[i])
		}
	}
}

func TestReadFile_BufferBoundary(t *testing.T) {
	longLine := strings.Repeat("a", 5000) // Строка длиннее буфера
	content := longLine + "\nshort"
	tmpfile, err := os.CreateTemp("", "boundary_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatal(err)
	}

	tmpfile.Close()

	r := NewReader(1024)
	var lines []string

	name := tmpfile.Name()

	err = r.ReadFile(&name, func(line string) error {
		lines = append(lines, line)
		return nil
	})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []string{longLine, "short"}
	if len(lines) != len(expected) {
		t.Fatalf("Expected %d lines, got %d, lines: %v", len(expected), len(lines), lines)
	}
	
	for i := range expected {
        if lines[i] != expected[i] {
            t.Errorf("Line %d mismatch: expected %d chars, got %d", 
                i, len(expected[i]), len(lines[i]))
        }
    }
}