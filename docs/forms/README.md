# Documentation for form.js

## JavaScript form builder tool

Form-Builder is a small library for creating forms dynamically. It allows many field types, with sorting.

## Installation

### Dependencies

#### Bootstrap CSS

```html

```

#### Sortable.js
https://github.com/SortableJS/Sortable
```html
<script src="https://cdn.jsdelivr.net/npm/sortablejs@latest/Sortable.min.js"></script>
```

#### 

Embed in your HTML file.
```html

<script src="form-builder.js"></script>
```

## Usage
```javascript
let formBuilder = new FormBuilder({
    title: 'Form Builder',
    description: 'This is some description',
    request: '/form-builder.php',
    method: 'POST',
    output: document.getElementById('output'),
    weighted: true,
});
```

### Options
```javascript
let formBuilder = new FormBuilder({
	id: 0, // 
	title: '', // Displays on top of the form
	description: '', // Displays on top of the form
	name: '', // Name of the form
	prefix: '', // Prefix of the form, will be generated on name-change
	weighted: false, // Adds weights to the fields
	output: null, // HTMLElement, output area
	request: '', // Request URL
	method: 'POST', // Request method
	fields: [], // Fields in the form
	regexp: /\W/, // Regexp for making prefix
	deleteRequest: '', // Delete request URL
	deleteMethod: 'DELETE', // Delete request method
});
```

#### `id` option

#### `title` option

#### `description` option

#### `name` option

#### `prefix` option

#### `weighted` option

#### `output` option

#### `request` option

#### `method` option
#### `fields` option
#### `regexp` option
#### `deleteRequest` option
#### `deleteMethod` option

### Method

#### `import(data: string)`