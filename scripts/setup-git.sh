#!/bin/bash

# Script để setup Git và push lên GitHub

echo "🚀 Thiết lập Git cho Language Learning Notes"
echo "=============================================="
echo ""

# Kiểm tra Git đã cài đặt chưa
if ! command -v git &> /dev/null; then
    echo "❌ Git chưa được cài đặt. Vui lòng cài đặt Git trước."
    exit 1
fi

echo "✅ Git đã được cài đặt"
echo ""

# Khởi tạo Git nếu chưa có
if [ ! -d ".git" ]; then
    echo "📦 Khởi tạo Git repository..."
    git init
    echo "✅ Đã khởi tạo Git repository"
else
    echo "✅ Git repository đã tồn tại"
fi

echo ""
echo "📝 Nhập thông tin GitHub của bạn:"
read -p "GitHub username: " GITHUB_USERNAME
read -p "Repository name (mặc định: Language_Learning_Notes): " REPO_NAME
REPO_NAME=${REPO_NAME:-Language_Learning_Notes}

echo ""
echo "🔗 Thêm remote repository..."
git remote remove origin 2>/dev/null
git remote add origin "https://github.com/${GITHUB_USERNAME}/${REPO_NAME}.git"

echo ""
echo "📋 Thêm tất cả files..."
git add .

echo ""
echo "💾 Commit lần đầu..."
git commit -m "Initial commit: Thêm tài liệu học tập đa ngôn ngữ"

echo ""
echo "🌿 Đặt branch chính là main..."
git branch -M main

echo ""
echo "📤 Push lên GitHub..."
echo "⚠️  Bạn sẽ cần nhập username và password (hoặc Personal Access Token)"
git push -u origin main

echo ""
echo "✅ Hoàn thành! Kiểm tra repository tại:"
echo "   https://github.com/${GITHUB_USERNAME}/${REPO_NAME}"

