package inliner

import (
	"fmt"
	"testing"
)

func TestQualifiedRule(t *testing.T) {
	input := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />

    <style type="text/css">
      body {
        background-color: #f2f2f2t;
      }

      p {
        font-family: 'Helvetica Neue', Verdana, sans-serif;
      }
    </style>
  </head>
  <body>
    <p>
      Inline me please!
    </p>
  </body>
</html>`

	expectedOutput := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
  </head>
  <body style="background-color: #f2f2f2t;">
    <p style="font-family: 'Helvetica Neue', Verdana, sans-serif;">
      Inline me please!
    </p>
  </body>
</html>`

	output, err := Inline(input)
	if err != nil {
		t.Fatal("Failed to inline html:", err)
	}

	if output != expectedOutput {
		t.Fatal(fmt.Sprintf("CSS inliner error\n   Expected:\n\"%s\"\n   Got:\n\"%s\"", expectedOutput, output))
	}
}