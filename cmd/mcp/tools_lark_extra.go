// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

package mcp

import (
	"encoding/json"
	"fmt"
)

// Additional curated tools promoted from frequent `lark_api` recipes.
// Flag names verified verbatim against the shortcut definitions in
// shortcuts/<domain>/*.go (task +complete, base +record-upsert,
// wiki +node-create, calendar +freebusy). Register order is defined in
// allTools() (tools.go) — these are inserted before the generic passthrough.

// toolTaskComplete exposes `task +complete` — close out a task by guid.
// Pairs with lark_task_create / lark_task_my so the bridge can run the
// full create → list → complete loop without falling back to lark_api.
func toolTaskComplete() tool {
	return tool{
		Name:        "lark_task_complete",
		Description: "Mark a Lark Task (todo) complete by its task id. Use when the user says they finished a task, or to close action items after a meeting. Find the task id first with lark_task_my (match on summary). To re-open or edit a task, use lark_api with the task v2 endpoint.",
		Schema: `{
  "type": "object",
  "properties": {
    "task_id": {"type": "string", "description": "Task id (guid) to complete. Get it from lark_task_my."},
    "as":      {"type": "string", "enum": ["user", "bot"]}
  },
  "required": ["task_id"]
}`,
		Build: func(args map[string]interface{}) ([]string, error) {
			id, ok := argString(args, "task_id")
			if !ok {
				return nil, fmt.Errorf("task_id is required")
			}
			argv := []string{"task", "+complete", "--task-id", id}
			argv = appendIdentity(argv, args)
			return argv, nil
		},
	}
}

// toolBaseRecordUpsert exposes `base +record-upsert` — create-or-update a
// single Bitable record. Lark has no separate CRM/tracker product, so this
// is how plugins write to a Base used as system-of-record (read side is
// lark_base_search). base-token / table-id are injected by the common base
// command wrapper, same as lark_base_search.
func toolBaseRecordUpsert() tool {
	return tool{
		Name:        "lark_base_record_upsert",
		Description: "Create or update one record in a Lark Base (Bitable) table. Use to write to a Base used as a system-of-record — CRM rows, tracker entries, decision logs. Pass `fields` as a flat field→value object (e.g. {\"Name\":\"Alice\",\"Status\":\"Todo\"}); do NOT wrap it in a `fields` key. Reads use lark_base_search; for many rows at once use lark_api with the batch endpoint. To scaffold a brand-new Base/table, use the base-deploy skill instead.",
		Schema: `{
  "type": "object",
  "properties": {
    "base_token": {"type": "string", "description": "Base (Bitable) app token."},
    "table_id":   {"type": "string", "description": "Table id (tblXXX) or display name."},
    "fields":     {"type": "object", "description": "Record field map, e.g. {\"Name\":\"Alice\",\"Status\":\"Todo\"}. Flat object, not wrapped in a fields key."},
    "record_json":{"type": "string", "description": "Advanced: the field-map JSON as a string. Overrides fields when present."},
    "as":         {"type": "string", "enum": ["user", "bot"]}
  },
  "required": ["base_token", "table_id"]
}`,
		Build: func(args map[string]interface{}) ([]string, error) {
			baseToken, ok := argString(args, "base_token")
			if !ok {
				return nil, fmt.Errorf("base_token is required")
			}
			tableID, ok := argString(args, "table_id")
			if !ok {
				return nil, fmt.Errorf("table_id is required")
			}
			var jsonStr string
			if raw, ok := args["fields"]; ok && raw != nil {
				buf, err := json.Marshal(raw)
				if err != nil {
					return nil, fmt.Errorf("fields must be a JSON object: %w", err)
				}
				jsonStr = string(buf)
			} else if s, ok := argString(args, "record_json"); ok && s != "" {
				jsonStr = s
			} else {
				return nil, fmt.Errorf("fields (object) or record_json (string) is required")
			}
			argv := []string{"base", "+record-upsert",
				"--base-token", baseToken,
				"--table-id", tableID,
				"--json", jsonStr,
			}
			argv = appendIdentity(argv, args)
			return argv, nil
		},
	}
}

// toolWikiNodeCreate exposes `wiki +node-create` — create a page (node) in a
// Wiki space. Knowledge-base plugins land KB articles, runbooks, and specs in
// Wiki rather than leaving them as chat messages.
func toolWikiNodeCreate() tool {
	return tool{
		Name:        "lark_wiki_node_create",
		Description: "Create a node (page) in a Lark Wiki space. Use when landing durable knowledge — KB articles, runbooks, postmortems, specs — into the org wiki. The new node defaults to a docx document you can then fill (lark_doc_create with wiki_node/wiki_space, or lark_api docx blocks). Use space_id='my_library' for the personal library. For a plain Drive doc instead of a wiki page, use lark_doc_create.",
		Schema: `{
  "type": "object",
  "properties": {
    "space_id":          {"type": "string", "description": "Target wiki space id; use 'my_library' for the personal document library."},
    "parent_node_token": {"type": "string", "description": "Parent wiki node token; if set, the new node is created under that parent."},
    "title":             {"type": "string", "description": "Node title."},
    "node_type":         {"type": "string", "enum": ["origin", "shortcut"], "description": "Node type (default origin)."},
    "obj_type":          {"type": "string", "description": "Target object type (default docx)."},
    "origin_node_token": {"type": "string", "description": "Source node token when node_type=shortcut."},
    "as":                {"type": "string", "enum": ["user", "bot"]}
  },
  "required": ["title"]
}`,
		Build: func(args map[string]interface{}) ([]string, error) {
			if _, ok := argString(args, "title"); !ok {
				return nil, fmt.Errorf("title is required")
			}
			argv := []string{"wiki", "+node-create"}
			argv = appendFlag(argv, "space-id", mustString(args, "space_id"))
			argv = appendFlag(argv, "parent-node-token", mustString(args, "parent_node_token"))
			argv = appendFlag(argv, "title", mustString(args, "title"))
			argv = appendFlag(argv, "node-type", mustString(args, "node_type"))
			argv = appendFlag(argv, "obj-type", mustString(args, "obj_type"))
			argv = appendFlag(argv, "origin-node-token", mustString(args, "origin_node_token"))
			argv = appendIdentity(argv, args)
			return argv, nil
		},
	}
}

// toolCalendarFreebusy exposes `calendar +freebusy` — check availability of
// the caller or another user before scheduling. Pairs with
// lark_calendar_create (resolve a free slot, then create the event).
func toolCalendarFreebusy() tool {
	return tool{
		Name:        "lark_calendar_freebusy",
		Description: "Check free/busy availability for the caller or another user over a time window. Use when you need to find an open meeting slot before calling lark_calendar_create, answer 'is X free at 3pm?', or compare several people's availability before proposing a time. Resolve each person's open_id via lark_contact_search first — availability is keyed by open_id, which cannot be guessed. Returns busy intervals only; project with the jq arg to keep just the start/end times you need and avoid payload bloat. For listing the caller's own scheduled events (not availability), use lark_calendar_agenda instead; for already-concluded meetings use lark_vc_search.",
		Schema: `{
  "type": "object",
  "properties": {
    "start":   {"type": "string", "description": "Start time (ISO 8601, default today)."},
    "end":     {"type": "string", "description": "End time (ISO 8601, default end of start day)."},
    "user_id": {"type": "string", "description": "Target user open_id (ou_ prefix). Default: current user."},
    "jq":      {"type": "string", "description": "Optional jq projection (e.g. '.data.freebusy[] | {start_time, end_time}')."},
    "as":      {"type": "string", "enum": ["user", "bot"]}
  }
}`,
		Build: func(args map[string]interface{}) ([]string, error) {
			argv := []string{"calendar", "+freebusy"}
			argv = appendFlag(argv, "start", mustString(args, "start"))
			argv = appendFlag(argv, "end", mustString(args, "end"))
			argv = appendFlag(argv, "user-id", mustString(args, "user_id"))
			argv = appendJq(argv, args)
			argv = appendIdentity(argv, args)
			return argv, nil
		},
	}
}
