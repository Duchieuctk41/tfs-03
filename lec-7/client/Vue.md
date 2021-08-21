# Vuejs
- Open source hỗ trợ build giao diện & SPA
- Viết theo mô hình [MVVM](https://wiki.matbao.net/mvc-mvp-mvvm-la-gi-thong-tin-can-biet-ve-cac-mo-hinh-lap-trinh/)
### Vue instance
- Mọi applications đều cần có 1 vue instance mới
- Sử dụng mô hình MVVM nên tên bắt đầu thường là "vm"
```sh
var vm = new Vue({
  // options
})
```
### Vue lifecycle 
![vue_lifecycle](https://vuejs.org/images/lifecycle.png)
### Template syntax
- text: dùng {{}} để call data, đọc trực tiếp msg > text
```sh
<span>Message: {{ msg }}</span>
```
- rawHTML: dùng v-html để biên dịch >  html
```sh
<p>Using mustaches: {{ rawHtml }}</p>
<p>Using v-html directive: <span v-html="rawHtml"></span></p>
```
- Attributes: dùng v-bind cho thuộc tính attribues
```sh
<div v-bind:id="dynamicId"></div>
```
- Using JavaScript Expressions: hỗ trợ tận răng cho js, viết thẳng code js trong {{}} hay v-bind đều đc (chỉ chứa 1 biểu thức duy nhất)
```sh
{{ number + 1 }}

{{ ok ? 'YES' : 'NO' }}

{{ message.split('').reverse().join('') }}

<div v-bind:id="'list-' + id"></div>
```
- Argument: 1 số directive có thể nhận đối số, viết sau dấu “:”
```sh
<a v-bind:href="url"> ... </a>
<a v-on:click="doSomething"> ... </a>
// truyền biểu thức js vào lm đối số.
<a v-bind:['foo' + bar]="value"> ... </a> 
```
- Modifiers : dấu “.”, cho biết rằng directive nên được ràng buộc theo 1 cách nào đó
```sh
<form v-on:submit.prevent="onSubmit"> ... </form>
```
- shorthand
```sh
<!-- full syntax v-bind -->
<a v-bind:href="url"> ... </a>

<!-- shorthand v-bind -->
<a :href="url"> ... </a>

<!-- full syntax v-on -->
<a v-on:click="doSomething"> ... </a>

<!-- shorthand v-on -->
<a @click="doSomething"> ... </a>
```
### Condition rendering
```sh
<-- v-if -->
<h1 v-if="awesome">Dau cat moi</h1>

<-- v-if-else -->
<h1 v-if="awesome">Hieu thu high</h1>
<h2 v-else>Hieu chu nhat</h2>

<-- show -->
<h1 v-show="ok">Dau cat moi</h1>
```
- v-if và v-show khác nhau: 
### List render
### Computed Properties and Watchers
### Component
### Single file components & Vue template loader