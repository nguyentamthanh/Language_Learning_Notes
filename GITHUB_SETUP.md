# 📤 Hướng dẫn đẩy code lên GitHub

## Bước 1: Tạo repository trên GitHub

1. Đăng nhập vào [GitHub](https://github.com)
2. Click nút "+" ở góc trên bên phải → "New repository"
3. Đặt tên: `Language_Learning_Notes`
4. Mô tả: "Tài liệu học tập đa ngôn ngữ lập trình"
5. Chọn **Public** (hoặc Private nếu muốn)
6. **KHÔNG** tích "Initialize with README" (vì đã có sẵn)
7. Click "Create repository"

## Bước 2: Khởi tạo Git trong thư mục dự án

```bash
# Di chuyển vào thư mục dự án
cd /home/thanh/Language_Learning_Notes

# Khởi tạo Git repository
git init

# Thêm tất cả files
git add .

# Commit lần đầu
git commit -m "Initial commit: Thêm tài liệu học tập đa ngôn ngữ"
```

## Bước 3: Kết nối với GitHub

```bash
# Thêm remote repository (thay YOUR_USERNAME bằng username GitHub của bạn)
git remote add origin https://github.com/YOUR_USERNAME/Language_Learning_Notes.git

# Kiểm tra remote đã được thêm
git remote -v
```

## Bước 4: Push code lên GitHub

```bash
# Push code lên GitHub (lần đầu)
git branch -M main
git push -u origin main
```

Nếu gặp lỗi authentication, bạn có thể:

### Option 1: Sử dụng Personal Access Token
1. GitHub → Settings → Developer settings → Personal access tokens → Tokens (classic)
2. Generate new token
3. Copy token
4. Khi push, nhập username và paste token làm password

### Option 2: Sử dụng SSH
```bash
# Tạo SSH key (nếu chưa có)
ssh-keygen -t ed25519 -C "your_email@example.com"

# Copy public key
cat ~/.ssh/id_ed25519.pub

# Thêm vào GitHub: Settings → SSH and GPG keys → New SSH key
# Sau đó đổi remote URL
git remote set-url origin git@github.com:YOUR_USERNAME/Language_Learning_Notes.git
```

## Bước 5: Kiểm tra trên GitHub

Vào trang repository trên GitHub để xem code đã được đẩy lên thành công!

## Các lệnh Git hữu ích

```bash
# Xem trạng thái
git status

# Xem lịch sử commit
git log

# Xem các thay đổi
git diff

# Thêm file cụ thể
git add filename.md

# Commit với message
git commit -m "Mô tả thay đổi"

# Push lên GitHub
git push

# Pull từ GitHub
git pull

# Tạo branch mới
git checkout -b feature/new-feature

# Chuyển branch
git checkout main

# Merge branch
git merge feature/new-feature
```

## Cập nhật sau này

Khi có thay đổi mới:

```bash
git add .
git commit -m "Mô tả thay đổi"
git push
```

## Troubleshooting

### Lỗi: "remote origin already exists"
```bash
git remote remove origin
git remote add origin https://github.com/YOUR_USERNAME/Language_Learning_Notes.git
```

### Lỗi: "failed to push some refs"
```bash
git pull origin main --rebase
git push
```

### Xem lại remote URL
```bash
git remote get-url origin
```

Chúc bạn thành công! 🚀

