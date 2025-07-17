package mocking

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"go.uber.org/mock/gomock"
)

/// ***** Testing using stdlib ***** ///
func TestCountBytes_Reader(t *testing.T) {
	input := []byte("123456789")
	reader := bytes.NewReader(input)

	got, err := CountBytes(reader)
	if err != nil {
		t.Errorf("CountBytes() error = %v, want nil", err)
	}
	if got != len(input) {
		t.Errorf("CountBytes() got = %v, want %v", got, len(input))
	}
}

// There is no way for us to force the reader to return an error. So we can't test our functions error path.

/// ***** Testing using mock ***** ///

type ManualFakeReader struct {
	RemainingData int
	Err           error
}

func (m *ManualFakeReader) Read(p []byte) (n int, err error) {
	if m.Err != nil {
		return 0, m.Err
	}

	if m.RemainingData == 0 {
		return 0, io.EOF
	}

	for i := range p {
		if m.RemainingData == 0 {
			// Reader does not return io.EOF on first end of data
			return i, nil
		}

		p[i] = byte(65 + (i % 26))
		m.RemainingData--
	}
	return len(p), nil
}

func TestCountBytes_ManualMock(t *testing.T) {
	cases := map[string]struct {
		DataSize int
		Err      error
	}{
		"ok":   {DataSize: 64},
		"fail": {Err: errors.New("fail")},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			reader := &ManualFakeReader{
				RemainingData: tt.DataSize,
				Err:           tt.Err,
			}

			got, err := CountBytes(reader)
			if tt.Err != nil && !errors.Is(err, tt.Err) {
				t.Errorf("CountBytes() error = %v, wantErr %v", err, tt.Err)
			}

			if got != tt.DataSize {
				t.Errorf("CountBytes() got = %v, want %v", got, tt.DataSize)
			}
		})
	}

}

/// ***** Testing using generated mock ***** ///

// We generate a mock using go mock
//go:generate mockgen -destination reader.gen.go . Reader
type Reader interface {
	Read(p []byte) (n int, err error)
}

func TestCountBytes_GoMock(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		reader := NewMockReader(ctrl)

		reader.EXPECT().Read(gomock.Any()).Return(1, nil).Times(2)
		reader.EXPECT().Read(gomock.Any()).Return(0, io.EOF).Times(1)
		//reader.EXPECT().Read(gomock.Any()).Return(1, nil) // would cause error since expects go in order, and this will not be called after io.EOF

		got, err := CountBytes(reader)
		if err != nil {
			t.Errorf("CountBytes() error = %v, want nil", err)
		}
		if got != 2 {
			t.Errorf("CountBytes() got = %v, want %v", got, 2)
		}
	})
	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		reader := NewMockReader(ctrl)

		expectedErr := errors.New("fail")
		reader.EXPECT().Read(gomock.Any()).Return(0, expectedErr)

		got, gotErr := CountBytes(reader)
		if !errors.Is(gotErr, expectedErr) {
			t.Errorf("CountBytes() got = %v, want %v", gotErr, expectedErr)
		}
		if got != 0 {
			t.Errorf("CountBytes() got = %v, want %v", got, 0)
		}
	})
}
