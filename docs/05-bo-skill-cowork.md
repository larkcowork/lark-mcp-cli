# 05 — Mở rộng: bộ skill Cowork (lark-cowork-plugins)

> `lark-cli` lõi chỉ gồm **MCP server + 26 skill lark-\* lõi** (lark-base, lark-im, lark-mail…). Các "công thức" workflow cấp cao bên dưới (morning-brief, inbox-zero, base-deploy…) **không còn đóng gói trong repo này** — chúng đã tách thành một **marketplace plugin riêng** để cập nhật độc lập và cài tuỳ nhu cầu.
>
> 📦 **Marketplace:** [github.com/larkcowork/lark-cowork-plugins](https://github.com/larkcowork/lark-cowork-plugins)

## Cài thêm để mở rộng năng lực

Bộ Cowork đóng gói thành các **plugin theo nhóm** (daily-assistant, crm-sales, governance, knowledge-docs, delivery-eng, lark-base-deploy…). Bạn cài plugin nào cần plugin đó:

```bash
# Thêm marketplace
/plugin marketplace add larkcowork/lark-cowork-plugins

# Cài nhóm bạn cần (ví dụ)
/plugin install daily-assistant@lark-cowork
/plugin install lark-base-deploy@lark-cowork
```

> Tuỳ phiên bản Claude Code/Cowork, có thể quản lý qua lệnh `/plugin` hoặc UI Connectors/Plugins. Xem README của marketplace để biết cách cài cập nhật nhất.

Sau khi cài, chỉ cần **nói tự nhiên** — mỗi skill có "câu kích hoạt" gợi ý; nói gần đúng là được.

> Mẹo: bắt đầu ngày bằng **"morning"** để có bản tóm tắt đầu ngày.

---

## Có gì trong marketplace (tham khảo nhanh)

### Nhóm "Đầu ngày / Tổng hợp" — plugin `daily-assistant`

| Skill | Hỏi gì | Làm gì |
|---|---|---|
| **morning-brief** | "morning", "sáng nay có gì" | Bản tóm đầu ngày: gộp mail + chat + approval + task ưu tiên (≤15 dòng). Có 5 biến thể: default/ic/exec/pm/sales |
| **daily-digest** | "tổng kết hôm nay", "wrap up" | Tóm cuối ngày: họp đã dự + task đã xong + điểm nổi bật trong inbox |
| **weekly-review** | "weekly review", "báo cáo tuần" | Tổng hợp lịch + task + OKR thành một mạch tường thuật |
| **overwhelm-triage** | "tôi quá tải" | Hỏi 1 câu để biết bạn ngợp ở đâu (mail/chat/task/họp) rồi điều hướng đúng skill |

### Nhóm "Mail & Chat / Họp & Lịch" — plugin `daily-assistant`

| Skill | Hỏi gì | Làm gì |
|---|---|---|
| **inbox-zero** | "clear inbox", "xử lý hết mail tồn" | Phân loại mail urgent/important/fyi/noise, hướng tới inbox trống |
| **im-digest** | "các group có gì", "tóm tắt chat" | Phân loại N tin mới mỗi group thành cần-action / cần-biết / bỏ qua |
| **meeting-prep** | "chuẩn bị họp X", "action items từ meeting" | Trước họp: gom context. Sau họp: trích action item thành task |
| **calendar-optimizer** | "tôi họp quá nhiều" | Soi 30 ngày họp, gợi ý decline/gộp/chuyển-async |
| **focus-mode** | "focus 2 tiếng", "DND", "deep work" | Chặn lịch + bật DND + báo team |
| **one-on-one-prep** | "1:1 prep với <người>" | Brief 1:1: OKR, task gần đây, ghi chú cũ, câu hỏi gợi ý |
| **contact-360** | "tôi sắp gặp <tên>" | Hồ sơ 360°: IM + mail + họp + doc + task của người đó |
| **task-prioritizer** | "việc nào quan trọng", "top 5 today" | Xếp hạng task theo deadline × rủi ro × OKR × người giao; nêu lý do |

### Nhóm "CRM / Sales" — plugin `crm-sales`

| Skill | Hỏi gì | Làm gì |
|---|---|---|
| **client-followup** | "khách im lặng", "follow-up khách" | Phát hiện liên hệ CRM lâu chưa động (>21 ngày), **soạn nháp** mail tái kết nối (chỉ nháp) |
| **deal-update** | "cập nhật deal sau gọi" | Lấy minutes → trích pain/budget/timeline → cập nhật Base → soạn nháp follow-up |
| **pipeline-review** | "pipeline review", "tổng quan deal" | Quét pipeline theo stage, deal kẹt, sắp chốt, xu hướng win-rate |

### Nhóm "Quản trị / Approval" — plugin `governance`

| Skill | Hỏi gì | Làm gì |
|---|---|---|
| **approval-triage** | "có gì cần duyệt", "queue duyệt" | Đọc approval pending, gợi ý APPROVE/CHECK/REJECT kèm trích policy |
| **approval-flow-sla** | "nghẽn duyệt ở đâu", "SLA duyệt" | Phân tích luồng approval: nghẽn ở node/người nào, chấm SLA |
| **decision-logger** | "log lại decision", "chốt cái này" | Phát hiện quyết định trong IM/minutes → ghi vào bảng Base |
| **permission-audit** | "permission audit", "quét quyền", "PII" | Quét Drive/Doc/Wiki/Base tìm quyền rủi ro (public/ngoài tổ chức/PII) — chỉ đọc |

### Nhóm "Tài liệu & Tri thức" — plugin `knowledge-docs`

| Skill | Hỏi gì | Làm gì |
|---|---|---|
| **doc-from-template** | "tạo doc theo template", "weekly report doc" | Soạn doc/wiki/sheet từ template có tên |
| **doc-restructure** | "wiki bị bừa", "wiki cleanup" | Soi Wiki: trang cũ/mồ côi/trùng → đề xuất archive/gộp/đổi cha |

### Nhóm "Kỹ thuật / Vận hành" — plugin `delivery-eng`

| Skill | Hỏi gì | Làm gì |
|---|---|---|
| **incident-retro** | "postmortem cho SEV-X" | Dựng postmortem blameless từ timeline IM on-call |
| **sprint-retro** | "sprint retro", "/retro" | Bản nháp retro cuối sprint: ticket đã đóng, velocity, blocker |

### Triển khai Lark Base end-to-end — plugin `lark-base-deploy`

| Skill | Hỏi gì | Làm gì |
|---|---|---|
| **base-deploy** | "triển khai base cho team", "/base-deploy" | Orchestrator 8 phase (Discovery→Handover), dựng Base hoàn chỉnh + dashboard + automation |

### Phát triển bridge MCP (cho người mở rộng CLI) — plugin `lark-cli-dev`

| Skill / lệnh | Dùng khi |
|---|---|
| **lark-cli-mcp** | Thêm/sửa/gỡ công cụ MCP, vá lỗi Claude Desktop mất kết nối, đổi schema tool |
| `/mcp-tools` `/mcp-test` `/mcp-call` `/mcp-doctor` `/mcp-add` `/mcp-rebuild` | Bộ lệnh nhanh thao tác trên bridge MCP của repo này |

---

> An toàn: các skill có thao tác **gửi** (mail, tin nhắn) luôn **soạn nháp / xem trước** trước, không tự gửi. Bạn duyệt rồi mới thật. Xem [07](07-bao-mat-quyen-rieng-tu.md).
