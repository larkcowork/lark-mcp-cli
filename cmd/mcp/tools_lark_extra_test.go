// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

package mcp

import (
	"strings"
	"testing"
)

// argvStr joins an argv for easy substring assertions.
func argvStr(t *testing.T, tl tool, args map[string]interface{}) string {
	t.Helper()
	argv, err := tl.Build(args)
	if err != nil {
		t.Fatalf("%s Build error: %v", tl.Name, err)
	}
	return strings.Join(argv, " ")
}

func TestTaskCompleteBuild(t *testing.T) {
	got := argvStr(t, toolTaskComplete(), map[string]interface{}{"task_id": "t-123", "as": "user"})
	for _, want := range []string{"task +complete", "--task-id t-123", "--as user"} {
		if !strings.Contains(got, want) {
			t.Errorf("argv %q missing %q", got, want)
		}
	}
	if _, err := toolTaskComplete().Build(map[string]interface{}{}); err == nil {
		t.Error("expected error when task_id missing")
	}
}

func TestBaseRecordUpsertBuild(t *testing.T) {
	got := argvStr(t, toolBaseRecordUpsert(), map[string]interface{}{
		"base_token": "bascn1", "table_id": "tbl1",
		"fields": map[string]interface{}{"Name": "Alice", "Status": "Todo"},
	})
	for _, want := range []string{"base +record-upsert", "--base-token bascn1", "--table-id tbl1", "--json"} {
		if !strings.Contains(got, want) {
			t.Errorf("argv %q missing %q", got, want)
		}
	}
	if !strings.Contains(got, `"Name":"Alice"`) {
		t.Errorf("argv %q missing marshalled fields", got)
	}
	// record_json fallback
	got2 := argvStr(t, toolBaseRecordUpsert(), map[string]interface{}{
		"base_token": "b", "table_id": "t", "record_json": `{"X":1}`,
	})
	if !strings.Contains(got2, `--json {"X":1}`) {
		t.Errorf("record_json not used: %q", got2)
	}
	// errors
	if _, err := toolBaseRecordUpsert().Build(map[string]interface{}{"table_id": "t", "fields": map[string]interface{}{"a": 1}}); err == nil {
		t.Error("expected error when base_token missing")
	}
	if _, err := toolBaseRecordUpsert().Build(map[string]interface{}{"base_token": "b", "table_id": "t"}); err == nil {
		t.Error("expected error when neither fields nor record_json provided")
	}
}

func TestWikiNodeCreateBuild(t *testing.T) {
	got := argvStr(t, toolWikiNodeCreate(), map[string]interface{}{
		"title": "Runbook", "space_id": "my_library", "node_type": "origin",
	})
	for _, want := range []string{"wiki +node-create", "--title Runbook", "--space-id my_library", "--node-type origin"} {
		if !strings.Contains(got, want) {
			t.Errorf("argv %q missing %q", got, want)
		}
	}
	if _, err := toolWikiNodeCreate().Build(map[string]interface{}{}); err == nil {
		t.Error("expected error when title missing")
	}
}

func TestCalendarFreebusyBuild(t *testing.T) {
	got := argvStr(t, toolCalendarFreebusy(), map[string]interface{}{"user_id": "ou_x", "start": "2026-06-07"})
	for _, want := range []string{"calendar +freebusy", "--user-id ou_x", "--start 2026-06-07"} {
		if !strings.Contains(got, want) {
			t.Errorf("argv %q missing %q", got, want)
		}
	}
	// no args → bare verb (defaults to caller/today in the shortcut)
	got2 := argvStr(t, toolCalendarFreebusy(), map[string]interface{}{})
	if strings.TrimSpace(got2) != "calendar +freebusy" {
		t.Errorf("expected bare verb, got %q", got2)
	}
}
