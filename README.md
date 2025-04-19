## 🧱 Dependencies

Make sure the following Go tools are installed before proceeding:

### 🔧 `nto-cli`

Used to generate code from internal models.
```bash
go install github.com/opbnq-q/nto-cli@latest
```

### 🛠️ `gorm gen` (`crudgen`)

Used for generating GORM-based CRUD operations.
```bash
go install github.com/kuzgoga/crudgen/cmd/crudgen@latest
```

### 🎨 `wails3`

Used to build modern desktop applications using Go + Web technologies.
```bash
go install -v github.com/wailsapp/wails/v3/cmd/wails3@latest
```

---

## 📦 Code Generation

After installing the dependencies, run the following command to generate code from your model definitions:

```bash
nto-cli ./internal/models
```

This will auto-generate the boilerplate based on the structure and annotations in your `./internal/models` directory.

---

## 📁 Project Structure

```text
.
├── internal
│   └── models         # Your application's data models
├── main.go            # Entry point
└── README.md          # This file
```

---

## 📌 Notes

- Make sure your Go environment is properly configured (`go env`).
- Ensure `$GOPATH/bin` is in your system `PATH` to use the installed tools.
- Regenerate code whenever you change your model definitions.
