# Chương 1: Giới thiệu JavaScript

## JavaScript là gì?

JavaScript là ngôn ngữ lập trình động, được sử dụng chủ yếu cho web development. Được tạo bởi Brendan Eich tại Netscape vào năm 1995.

## Tại sao học JavaScript?

### ✅ Ưu điểm

1. **Chạy mọi nơi**: Browser, Server (Node.js), Mobile (React Native)
2. **Dễ bắt đầu**: Chỉ cần browser và text editor
3. **Cộng đồng lớn**: Nhiều thư viện và framework
4. **Nhu cầu cao**: Một trong những ngôn ngữ phổ biến nhất

### 📊 Ứng dụng của JavaScript

- **Frontend**: React, Vue, Angular
- **Backend**: Node.js, Express
- **Mobile**: React Native, Ionic
- **Desktop**: Electron
- **Game Development**: Phaser, Three.js
- **IoT**: Johnny-Five

## JavaScript vs Java

**Không liên quan!** Chỉ giống tên, JavaScript và Java là hai ngôn ngữ hoàn toàn khác nhau.

## ECMAScript

JavaScript tuân theo chuẩn ECMAScript:
- **ES5** (2009): Phiên bản ổn định
- **ES6/ES2015**: Classes, Arrow functions, Modules
- **ES2016-ES2023**: Các tính năng mới liên tục được thêm

## Hello World đầu tiên

### Trong Browser Console
```javascript
console.log("Hello, World!");
```

### Trong HTML
```html
<script>
    console.log("Hello, World!");
    alert("Xin chào JavaScript!");
</script>
```

### Với Node.js
```javascript
// app.js
console.log("Hello, World!");
```

Chạy: `node app.js`

## Các khái niệm cơ bản

### 1. Interpreted Language
JavaScript được thông dịch tại runtime, không cần compile.

### 2. Dynamic Typing
Kiểu dữ liệu được xác định tự động.

### 3. Prototype-based
Sử dụng prototypes thay vì classes (trước ES6).

### 4. First-class Functions
Functions là objects, có thể gán vào biến, truyền như tham số.

## Tóm tắt

- JavaScript là ngôn ngữ web phổ biến nhất
- Chạy được ở nhiều môi trường khác nhau
- Dễ học và có cộng đồng lớn
- Liên tục được cập nhật với các tính năng mới

**Tiếp theo**: [Chương 2: Cài đặt và Môi trường](./02-setup.md)

