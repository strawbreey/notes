---
title: "Bootstrap Utilities"
date: 2020-12-09T13:13:47+08:00
draft: false
---

### 目录

1. border

```html
<!-- Additive -->
<span class="border"></span>
<span class="border-top"></span>
<span class="border-right"></span>
<span class="border-bottom"></span>
<span class="border-left"></span>

<!-- Subtractive -->
<span class="border-0"></span>
<span class="border-top-0"></span>
<span class="border-right-0"></span>
<span class="border-bottom-0"></span>
<span class="border-left-0"></span>

<!-- Border color -->
<span class="border border-primary"></span>
<span class="border border-secondary"></span>
<span class="border border-success"></span>
<span class="border border-danger"></span>
<span class="border border-warning"></span>
<span class="border border-info"></span>
<span class="border border-light"></span>
<span class="border border-dark"></span>
<span class="border border-white"></span>

<!-- Border-radius -->
<img src="..." class="rounded" alt="...">
<img src="..." class="rounded-top" alt="...">
<img src="..." class="rounded-right" alt="...">
<img src="..." class="rounded-bottom" alt="...">
<img src="..." class="rounded-left" alt="...">
<img src="..." class="rounded-circle" alt="...">
<img src="..." class="rounded-pill" alt="...">
<img src="..." class="rounded-0" alt="...">

 <!-- border-radius Sizes -->
<img src="..." class="rounded-sm" alt="...">
<img src="..." class="rounded-lg" alt="...">

```

2. Clearfix

```html
<div class="clearfix">...</div>

<style>
// Mixin itself
@mixin clearfix() {
  &::after {
    display: block;
    content: "";
    clear: both;
  }
}

// Usage as a mixin
.element {
  @include clearfix;
}
</style>
```

3. Close icon

```html
<button type="button" class="close" aria-label="Close">
  <span aria-hidden="true">&times;</span>
</button>
```

4. Colors

```html
<p class="text-primary">.text-primary</p>
<p class="text-secondary">.text-secondary</p>
<p class="text-success">.text-success</p>
<p class="text-danger">.text-danger</p>
<p class="text-warning">.text-warning</p>
<p class="text-info">.text-info</p>
<p class="text-light bg-dark">.text-light</p>
<p class="text-dark">.text-dark</p>
<p class="text-body">.text-body</p>
<p class="text-muted">.text-muted</p>
<p class="text-white bg-dark">.text-white</p>
<p class="text-black-50">.text-black-50</p>
<p class="text-white-50 bg-dark">.text-white-50</p>

<div class="p-3 mb-2 bg-primary text-white">.bg-primary</div>
<div class="p-3 mb-2 bg-secondary text-white">.bg-secondary</div>
<div class="p-3 mb-2 bg-success text-white">.bg-success</div>
<div class="p-3 mb-2 bg-danger text-white">.bg-danger</div>
<div class="p-3 mb-2 bg-warning text-dark">.bg-warning</div>
<div class="p-3 mb-2 bg-info text-white">.bg-info</div>
<div class="p-3 mb-2 bg-light text-dark">.bg-light</div>
<div class="p-3 mb-2 bg-dark text-white">.bg-dark</div>
<div class="p-3 mb-2 bg-white text-dark">.bg-white</div>
<div class="p-3 mb-2 bg-transparent text-dark">.bg-transparent</div>

.bg-gradient-primary
.bg-gradient-secondary
.bg-gradient-success
.bg-gradient-danger
.bg-gradient-warning
.bg-gradient-info
.bg-gradient-light
.bg-gradient-dark
```

5. Display

none
inline
inline-block
block
table
table-cell
table-row
flex
inline-flex

Hiding elements

Screen Size	Class
Hidden on all	.d-none
Hidden only on xs	.d-none .d-sm-block
Hidden only on sm	.d-sm-none .d-md-block
Hidden only on md	.d-md-none .d-lg-block
Hidden only on lg	.d-lg-none .d-xl-block
Hidden only on xl	.d-xl-none
Visible on all	.d-block
Visible only on xs	.d-block .d-sm-none
Visible only on sm	.d-none .d-sm-block .d-md-none
Visible only on md	.d-none .d-md-block .d-lg-none
Visible only on lg	.d-none .d-lg-block .d-xl-none
Visible only on xl	.d-none .d-xl-block



### 参考链接

- https://getbootstrap.com/docs/5.0/utilities/api/