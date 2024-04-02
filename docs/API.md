## GET
### api/user
```js
{
	name: "...",
	role: 'unauthorized' | 'user' | 'admin',
	...
}
```
### api/search?type=x&text=s
Возвращает не более 20 наиболее релевантных результатов.
```js
// Указывает, где производить поиск - среди учителей, предметов, или и тех и тех
// если поиск по категории не осуществлялся - соответствующий массив вернется пустым
// type = 'techers' | 'subjects' | 'both'

{
	subjects: [
		{
			id: 1,
			name: "...",
			tags: ["1", "2"..],
			score: 9.1
		},
	],
	teachers: [
		{
			id: 1,
			name: "...",
			tags: ["1", "2"..],
			score: 9.1
		},
	]
}
```
### api/teachers?offeset=x&amount=y
```js
[
	{
		name: "...",
		tags: ["1", "2"..],
		score: 9.1
	},
	...
]
```
### api/subjects?offeset=x&amount=y
```js
[
	{
		name: "...",
		tags: ["1", "2"..],
		score: 9.1
	},
]
```
### api/subject/*'id'*
```js
{
	id: 1,
	faculties: ["1", "2"],
	creteria: [
		{
			name: "...",
			rating: 1.0	
		},
	],
	avg_rating: 1.0,
	lecturers: ["1", "2", "3"],
	teachers: ["1", "2", "3"]
}
```
### api/teacher/*'id'*
```js
{
	id: 1,
	avatar: "path/to/image",
	creteria: [
		{
			name: "...",
			rating: 1.0	
		},
	],
	avg_rating: 1.0,
	subjects: ["1", "2", "3"]
}
```
### api/subject/*'id'*/reviews?offset=x&amount=y
```js
[
	{
		id: 1,
		rating: [1.0, 3.4, ],
		text: "...",         // only preview (first x symbols)
		created: "10.02.2023",
		author: "..."
	},
]
```
### api/teacher/*'id'*/reviews?offset=*x*&amount=*y*
```js
[
	{
		avg_rating: 6.6,
		text: "...",         // only preview (first x symbols)
		created: "10.02.2023",
		author: "anon"
	},
]
```
### api/review/*'id'*/
```js
{
	id: 1,
	lecturer: "...",
	teacher: "...",
	subject: "...",
	text: "...",
	author: "anon",
	created: "10.02.2023",
	criteria: [
		{
			name: "...",
			rating: 1.0	
		},
	]
}
```
## OPTIONS
### api/user/verify
?
## POST
### api/review
**request**
```js
{
	lecturer_id: 1,
	teacher_id: 1,
	criteria: [
		{
			name: "...",
			rating: 1.0	
		},
	],
	text: "..." 
}
```
**response**
```
202
```
