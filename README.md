# 🕒 Task Scheduler CLI

A **lightweight and efficient task scheduler** built with **Go**.  
Allows users to **add, list, and delete scheduled tasks** that run at defined intervals.  

## 🚀 Features
✅ Schedule recurring tasks (e.g., every 5 minutes, every hour)  
✅ Manage tasks via CLI (`add`, `list`, `delete`)  
✅ Persistent task storage using JSON  
✅ Multi-task execution with Goroutines  
✅ Real-time task execution tracking  

## 📥 Installation
```sh
git clone https://github.com/yourusername/task-scheduler.git
cd task-scheduler
go mod tidy
go build -o scheduler
./scheduler
```

## 🛠️ Usage

### ➕ Add a Task
```sh
add "Task Name" X "Task Description"
```
**Example:**  
```sh
add "DB Backup" 60 "Backup PostgreSQL database"
```

### 📋 List All Tasks
```sh
list
```
**Example Output:**
```
📋 Scheduled Tasks:
🔹 ID 1 | Name: DB Backup | Schedule: 60 min | Desc: Backup PostgreSQL database
```

### ❌ Delete a Task
```sh
delete 1
```
✅ If successful:
```
✅ Deleted task 1
```
❌ If the task ID does not exist:
```
❌ Task not found
```

## 🚀 Use Cases
- Automate database backups (`pg_dump` every hour)  
- Run API polling (`GET request` every 5 minutes)  
- Schedule periodic file cleanup (delete temp files every 30 minutes)  
- Trigger IoT actions (e.g., turn on lights at scheduled times)  

## 🔧 Future Improvements
✅ Web UI for task management  
✅ Integration with **Redis or PostgreSQL**  
✅ Support for **cron-like scheduling**  
✅ WebSockets for **real-time updates**  

## 📜 License
This project is **MIT Licensed** – Free to modify and use.

## 🤝 Contributing
Fork the repo & submit pull requests! 🚀  

## 📬 Contact
📧 Email: **iamoluwaseyiolasupo@gmail.com**  
🐙 GitHub: (https://github.com/Oluwaseyi3)  
