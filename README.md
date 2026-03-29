# 📧 MailMon

MailMon is a lightweight **email automation tool built in Go** that allows sending emails to multiple recipients using templates and structured data.

It uses a **producer-consumer model** to efficiently process and send emails, making it scalable and easy to extend.

---

## ✨ Features

* 📤 Send emails to multiple users
* 📄 Template-based email generation
* 📊 CSV-based input (names & emails)
* ⚡ Concurrent processing using Go routines
* 🧩 Modular architecture (producer & consumer)

---

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **Concepts Used:**

  * Goroutines & Channels
  * Producer-Consumer Pattern
  * File Handling (CSV)
  * Template Rendering

---

## 📁 Project Structure

```
MailMon/
│── main.go              # Entry point
│── producer.go          # Produces email jobs
│── consumer.go          # Sends emails
│── execute.go           # Execution logic
│── email.tmpl           # Email template
│── names_and_emails.csv # Input data
│── go.mod               # Dependencies
│── info.md              # Additional info
```

---

## ⚙️ How It Works

1. **Read CSV File** → Extract names and email addresses
2. **Generate Email Content** → Using `email.tmpl`
3. **Producer** → Creates email tasks
4. **Consumer** → Sends emails concurrently

---

## ▶️ Getting Started

### 1. Clone the Repository

```
git clone https://github.com/Mksthatsall/MailMon.git
cd MailMon
```

---

### 2. Install Dependencies

```
go mod tidy
```

---

### 3. Run the Project

```
go run main.go
```

---

## 📄 Input Format (CSV)

Example:

```
Name,Email
John Doe,john@example.com
Jane Doe,jane@example.com
```

---

## 📝 Email Template

Modify `email.tmpl` to customize the email content:

```
Hello {{.Name}},

This is a test email.

Regards,
MailMon
```

---

## 🚧 Future Improvements

* ✅ SMTP integration (Gmail / SendGrid)
* 🔐 Authentication support
* 📈 Logging & monitoring
* 🌐 Web UI / Dashboard
* 📦 Docker support

---

## 🤝 Contributing

Feel free to fork the repo and contribute!

---

## 📌 Author

* **Mridul Krishna Sharma**

---

## ⭐ Show Your Support

If you like this project, give it a ⭐ on GitHub!
