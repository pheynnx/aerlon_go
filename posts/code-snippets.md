---
Title: Code Block Preview
Date: September 8, 2023
Slug: code-blocks
Series: coding
Categories:
  - code
  - snippets
  - go
  - rust
  - js
  - ts
  - nim
  - zig
Published: true
Featured: false
PostSnippet: Testing coding syntax and html code blocks.
---

```TypeScript
const a: number = 50;
```

```Python
print("Hello World!")
```

```JavaScript
const main = () => {console.log("Hello World")}
```

```nim
echo "Hello nimlang!"
```

```Rust
pub async fn add(x: int, y: int) -> int {
    x + y
};

add(2, 7).await;
```

```go
func main() {
    go func() {
     fmt.Println("go channels")
  }()
}
```

```zig
const std = @import("std");
const parseInt = std.fmt.parseInt;
```

```html
<html>
  <head> </head>
  <body></body>
</html>
```

```scss
.test {
  border: 1px solid #567fbc;

  &:hover {
    background-color: #567fbc;
  }
}
```
