# lark-mcp-cli — Lark/Feishu × Claude (MCP)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.23-blue.svg)](https://go.dev/)
[![MCP](https://img.shields.io/badge/MCP-stdio%20%2B%20http-7C3AED.svg)](#-lark-mcp--dùng-lark-ngay-trong-claude-desktop--web)
[![Tools](https://img.shields.io/badge/MCP%20tools-21-0EA5E9.svg)](./docs/06-cong-cu-mcp.md)
[![Cowork Skills](https://img.shields.io/badge/Cowork%20skills-marketplace-F59E0B.svg)](https://github.com/larkcowork/lark-cowork-plugins)

[English](./README.md) | **Tiếng Việt** | [中文版](./README.zh.md)

> **Điều khiển Lark/Feishu bằng Claude — đọc mail, tóm họp, dọn task, gửi tin nhắn, tạo doc, duyệt approval — bằng prompt tiếng Việt, ngay trong Claude Desktop (Cowork) hoặc claude.ai (web). Không cần Claude Code, không copy-paste, không cần biết lập trình.**

Đây là `lark-cli` (CLI chính thức của [Lark/Feishu](https://www.larksuite.com/) — 200+ lệnh, 18 nhóm nghiệp vụ) **kèm một MCP server tích hợp sẵn**, để mọi giao diện [Claude](https://claude.ai) thao tác Lark ngay lập tức. Các skill workflow Cowork cấp cao (morning-brief, inbox-zero, base-deploy…) nằm ở một **marketplace plugin riêng, cài tuỳ nhu cầu** → **[lark-cowork-plugins](https://github.com/larkcowork/lark-cowork-plugins)**.

[🚀 Lark MCP](#-lark-mcp--dùng-lark-ngay-trong-claude-desktop--web) · [Cài đặt](#cài-đặt-nhanh) · [21 công cụ MCP](./docs/06-cong-cu-mcp.md) · [Skill Cowork](./docs/05-bo-skill-cowork.md) · [Tài liệu đầy đủ](./docs/README.md) · [Bảo mật](./docs/07-bao-mat-quyen-rieng-tu.md)

---

## 🚀 Lark MCP — dùng Lark ngay trong Claude Desktop & web

> **Điểm mới của bản dựng này:** một MCP server gốc ([Model Context Protocol](https://modelcontextprotocol.io/)) — `lark-cli mcp serve` — biến CLI thành **tay chân của Claude** trong Lark. Mở rộng thêm bằng các skill workflow Cowork từ marketplace [lark-cowork-plugins](https://github.com/larkcowork/lark-cowork-plugins). Hướng tới **business user**, không chỉ lập trình viên.

| | |
| --- | --- |
| 🧰 **21 công cụ MCP** | IM, Mail, Calendar, Docs, Base, Contact, Task, Drive, Sheets, Meetings, OKR + cửa thoát hiểm `lark_api` → [tham chiếu công cụ](./docs/06-cong-cu-mcp.md) |
| 🖥️ **Claude Desktop (Cowork)** | Transport **stdio** local — 1 khối cấu hình, restart, xong → [hướng dẫn desktop](./docs/02-cai-dat-claude-desktop.md) |
| 🌐 **claude.ai (web)** | Transport **HTTP streamable** + Cloudflare Tunnel + **bearer token** tích hợp → [hướng dẫn web](./docs/03-ket-noi-web-claude-ai.md) |
| 🧠 **Skill Cowork (tuỳ chọn)** | `morning-brief`, `inbox-zero`, `meeting-prep`, `task-prioritizer`, `approval-triage`… cài từ marketplace [lark-cowork-plugins](https://github.com/larkcowork/lark-cowork-plugins) → [cách mở rộng](./docs/05-bo-skill-cowork.md) |
| 🔐 **An toàn mặc định** | Dry-run mọi thao tác ghi, mail lưu nháp tới khi xác nhận, audit log, token trong OS keychain, bearer-token so sánh constant-time → [bảo mật](./docs/07-bao-mat-quyen-rieng-tu.md) |

### Cách hoạt động

```
Bạn ──"Sáng nay tôi có gì?"──▶  Claude (Desktop / Web)
                                      │  chọn công cụ + tham số
                                      ▼
                            lark-cli mcp serve   (trên máy bạn)
                                      │  chạy lark-cli <verb> bằng auth CỦA BẠN
                                      ▼
                            open.larksuite.com / open.feishu.cn
```

MCP server spawn subprocess `lark-cli` cho mỗi tool call, tái dùng auth/profile/keychain sẵn có → **data chỉ đi tới Lark**, không qua bên thứ ba (ở chế độ stdio/local).

### Khởi động 60 giây (Claude Desktop)

```bash
# 1. Build & cài binary (thêm lệnh `mcp`)
./scripts/setup-mcp.sh                 # cài vào ~/bin/lark-cli
lark-cli mcp tools                     # phải liệt kê 21 tool

# 2. Đăng nhập Lark một lần (OAuth trình duyệt, token lưu OS keychain)
lark-cli auth login

# 3. Khai báo vào file cấu hình Claude Desktop, rồi restart Claude Desktop:
#    ~/Library/Application Support/Claude/claude_desktop_config.json
```

```json
{
  "mcpServers": {
    "lark-cli": {
      "command": "/đường-dẫn-tuyệt-đối/tới/lark-cli",
      "args": ["mcp", "serve"],
      "env": { "NO_COLOR": "1" }
    }
  }
}
```

Giờ hỏi Cowork: *"Liệt kê lịch hôm nay của tôi"* hoặc *"Tìm liên hệ tên Nguyễn Văn A"*. → **Đầy đủ từng bước:** [docs/02](./docs/02-cai-dat-claude-desktop.md). **Web (claude.ai):** [docs/03](./docs/03-ket-noi-web-claude-ai.md).

### Mở ra web (claude.ai) — an toàn

```bash
export LARK_MCP_BEARER_TOKEN=$(openssl rand -hex 32)        # secret nằm trong biến môi trường
lark-cli mcp serve --transport http --addr 127.0.0.1:3000 --audit-log ~/.lark-mcp-audit.ndjson
cloudflared tunnel --url http://127.0.0.1:3000             # → URL HTTPS công khai
```

Endpoint HTTP bắt buộc `Authorization: Bearer <token>` ở `/` và `/mcp` (so sánh constant-time, `/health` để mở cho liveness). **Tuyệt đối không mở tunnel khi chưa có token.** Thêm URL (`https://…/mcp`) làm Custom Connector trên claude.ai → [hướng dẫn web an toàn](./docs/03-ket-noi-web-claude-ai.md).

### 📚 Tài liệu MCP (cho business user)

| Tài liệu | Nội dung |
| --- | --- |
| [docs/README](./docs/README.md) | Mục lục + bắt đầu 3 bước |
| [01 — Tổng quan & giá trị](./docs/01-tong-quan.md) | Vì sao MCP, ROI |
| [02 — Cài Claude Desktop](./docs/02-cai-dat-claude-desktop.md) | stdio, từng bước |
| [03 — Kết nối web claude.ai](./docs/03-ket-noi-web-claude-ai.md) | HTTP + tunnel + bearer token |
| [04 — Đăng nhập & quyền](./docs/04-dang-nhap-va-quyen.md) | user/bot, scopes |
| [05 — Skill Cowork](./docs/05-bo-skill-cowork.md) | mở rộng qua marketplace |
| [06 — Công cụ MCP](./docs/06-cong-cu-mcp.md) | tham chiếu 21 tool |
| [07 — Bảo mật & riêng tư](./docs/07-bao-mat-quyen-rieng-tu.md) | bearer, audit, data flow |
| [08 — Xử lý sự cố](./docs/08-xu-ly-su-co.md) | gồm sự cố đã biết |
| [09 — Cập nhật & bảo trì](./docs/09-cap-nhat-bao-tri.md) | nâng cấp |
| [MCP_QUICKSTART](./MCP_QUICKSTART.md) | mọi MCP host (Cursor, Zed, Cline…) |

> Muốn xây/mở rộng tool? Xem [`cmd/mcp/README.md`](./cmd/mcp/README.md) cho kiến trúc bridge và quy trình `/mcp-add`.

---

## Vì sao chọn lark-cli?

- **Thiết kế Agent-Native** — bộ [Skill](./skills/) có cấu trúc, tương thích các công cụ AI phổ biến; Agent thao tác Lark không cần cấu hình thêm.
- **Phủ rộng** — 18 nhóm nghiệp vụ, 200+ lệnh, 26 skill lark-* lõi + 21 công cụ MCP; mở rộng thêm bằng marketplace [lark-cowork-plugins](https://github.com/larkcowork/lark-cowork-plugins).
- **Tối ưu cho AI** — mọi lệnh tham số gọn, default thông minh, output có cấu trúc để tăng tỉ lệ gọi thành công.
- **Mã nguồn mở, không rào cản** — giấy phép MIT.
- **Chạy trong vài phút** — tạo app, đăng nhập tương tác, từ cài tới lần gọi API đầu rất nhanh.
- **An toàn, kiểm soát được** — chống injection input, làm sạch output terminal, lưu credential bằng keychain hệ điều hành.
- **Kiến trúc 3 lớp** — Shortcuts (thân thiện người & AI) → API Commands (đồng bộ nền tảng) → Raw API (phủ toàn bộ).

## Năng lực theo nhóm

| Nhóm | Năng lực |
| --- | --- |
| 📅 Calendar | Xem/tạo/cập nhật lịch, mời người, đặt phòng họp, RSVP, tra free/busy |
| 💬 Messenger | Gửi/trả lời tin, tạo & quản group, lịch sử & thread, tìm tin, tải media |
| 📄 Docs | Tạo/đọc/sửa/tìm doc, media & whiteboard |
| 📁 Drive | Upload/download file, tìm doc & wiki, quản lý comment |
| 📊 Base | Bảng, field, record, view, dashboard, workflow, form, role; thống kê & phân tích |
| 📈 Sheets | Tạo/đọc/ghi/append/tìm/export dữ liệu |
| 🖼️ Slides | Tạo & quản presentation, đọc nội dung, thêm/xoá slide |
| ✅ Tasks | Tạo/tra/cập nhật/hoàn tất task; cl.list, subtask, nhắc việc |
| 📚 Wiki | Quản lý không gian tri thức, node, doc |
| 👤 Contact | Tìm người theo tên/email/sđt, lấy hồ sơ |
| 📧 Mail | Đọc/tìm/gửi/trả lời/forward, quản nháp, theo dõi mail mới |
| 🎥 Meetings | Tìm bản ghi họp, truy minutes & recording |
| ✍️ Approval | Tra/duyệt/từ chối/chuyển task, huỷ & CC instance |
| 🎯 OKR | Tra/tạo/cập nhật OKR; objective, key result, tiến độ |

## Cài đặt nhanh

### Yêu cầu

- Node.js (`npm`/`npx`)
- Go `v1.23`+ và Python 3 (chỉ cần khi build từ mã nguồn)

### Build từ mã nguồn (kèm lệnh MCP)

```bash
make build                 # tạo ./lark-cli
./scripts/setup-mcp.sh     # cài vào ~/bin/lark-cli + verify 21 tool
lark-cli config init       # cấu hình app (tương tác, 1 lần)
lark-cli auth login        # đăng nhập
```

> Hướng dẫn cài cho **business user** (không cần dev) — desktop & web: xem [docs/](./docs/README.md).

## Đăng nhập

```bash
lark-cli auth login              # đăng nhập tương tác (TUI chọn domain & mức quyền)
lark-cli auth login --recommend  # tự chọn các scope thường dùng
lark-cli auth status             # kiểm tra user/bot
lark-cli <lệnh> --as user|bot    # chọn danh tính khi chạy lệnh
```

Chi tiết user vs bot, scope, app riêng cho enterprise: [docs/04](./docs/04-dang-nhap-va-quyen.md).

## Hệ thống lệnh 3 lớp

1. **Shortcuts** — 1 lệnh gói nhiều API, thân thiện người & AI (ví dụ `im +messages-send`, `mail +triage`).
2. **API Commands** — 1 lệnh = 1 API, tham số có cấu trúc (`im chats search`…).
3. **Raw API** — tương đương curl, phủ toàn bộ (`lark-cli api POST /open-apis/...`).

AI tự chọn lớp phù hợp dựa vào skill. (Bản tiếng Anh có ví dụ đầy đủ: [README.md](./README.md#three-layer-command-system).)

## Bảo mật & cảnh báo (đọc trước khi dùng)

- Mọi thao tác **ghi** có **dry-run** xem trước; **mail mặc định lưu nháp** (cần `confirm_send=true` mới gửi).
- Token nằm trong **OS keychain**; secret (bearer token) truyền qua **biến môi trường** `LARK_MCP_BEARER_TOKEN`.
- **Mở cổng web = phơi toàn quyền tài khoản Lark** → bắt buộc bật bearer token và/hoặc Cloudflare Access; bật `--audit-log`.
- Chống injection input, làm sạch output terminal.

Đầy đủ: [docs/07 — Bảo mật & quyền riêng tư](./docs/07-bao-mat-quyen-rieng-tu.md).

## Đóng góp

Hoan nghênh đóng góp! Mở [Issue](https://github.com/pluginmd/lark-mcp-cli/issues) hoặc [Pull Request](https://github.com/pluginmd/lark-mcp-cli/pulls). Trước khi mở PR, xem [AGENTS.md](./AGENTS.md) cho quy trình build/test/checklist.

## Giấy phép

Dự án theo giấy phép **MIT**. Là sản phẩm phái sinh từ [larksuite/cli](https://github.com/larksuite/cli) (cũng MIT). Khi chạy, công cụ gọi API Lark/Feishu Open Platform — bạn phải tuân thủ điều khoản & chính sách của Lark/Feishu:

- [Feishu User Terms](https://www.feishu.cn/terms) · [Feishu Privacy](https://www.feishu.cn/privacy)
- [Lark User Terms](https://www.larksuite.com/user-terms-of-service) · [Lark Privacy](https://www.larksuite.com/privacy-policy)
