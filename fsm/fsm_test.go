package fsm

import (
	"bytes"
	"context"
	"fmt"
	"testing"
)

func TestFSMStart(t *testing.T) {
	ctx := context.TODO()
	buf := bytes.NewBuffer(nil)
	pfb := func(_ context.Context, state string, _ map[string]interface{}) {
		fmt.Fprintf(buf, "before state %s\n", state)
	}
	pfa := func(_ context.Context, state string, _ map[string]interface{}) {
		fmt.Fprintf(buf, "after state %s\n", state)
	}
	f := New(StateInitial("1"), StateHookBefore(pfb), StateHookAfter(pfa))
	f1 := func(_ context.Context, args map[string]interface{}) (string, map[string]interface{}, error) {
		if v, ok := args["request"].(string); !ok || v == "" {
			return "", nil, fmt.Errorf("empty request")
		}
		return "2", map[string]interface{}{"response": "test2"}, nil
	}
	f2 := func(_ context.Context, args map[string]interface{}) (string, map[string]interface{}, error) {
		if v, ok := args["response"].(string); !ok || v == "" {
			return "", nil, fmt.Errorf("empty response")
		}
		return "", map[string]interface{}{"response": "test"}, nil
	}
	f.State("1", f1)
	f.State("2", f2)
	args, err := f.Start(ctx, map[string]interface{}{"request": "test1"})
	if err != nil {
		t.Fatal(err)
	} else if v, ok := args["response"].(string); !ok || v == "" {
		t.Fatalf("nil rsp: %#+v", args)
	} else if v != "test" {
		t.Fatalf("invalid rsp %#+v", args)
	}

	if !bytes.Contains(buf.Bytes(), []byte(`before state 1`)) ||
		!bytes.Contains(buf.Bytes(), []byte(`before state 2`)) ||
		!bytes.Contains(buf.Bytes(), []byte(`after state 1`)) ||
		!bytes.Contains(buf.Bytes(), []byte(`after state 2`)) {
		t.Fatalf("fsm not works properly or hooks error, buf: %s", buf.Bytes())
	}
}
