
<h1>TODO app</h1>
<h2>Using</h2>
<p>authorization</p>
<p><b>[POST]</b> /signup</p>

```json
{
  "email":"your_email",
  "password": "your_password",
  "first_name": "your_name",
  "last_name":"your_surname"
}
```

<p><b>[POST] - /signin</b></p>

```json
{
  "email": "your_email",
  "password": "your_password"
}
```
<hr>
<p>todo app</p>
<p><b>[GET]</b>  /api/todo - get all tasks </p>
<p><b>[GET]</b>  /api/todo/:id - get task by id </p>
<p><b>[POST]</b>  /api/todo - create task </p>

```json
{
  "title": "title",
  "description": "give description",
  "deadline": "set deadline",
}
```

<p><b>[PUT]</b>  /api/todo/:id - update task by id </p>
<p><b>[POST]</b>  /api/todo/:id/execute - mark execute </p>
<p><b>[DELETE]</b>  /api/todo/:id - delete task by id</p>
<hr>
<h2>Dependencies</h2>
<a href="github.com/valyala/fasthttp">fasthttp</a><br>
<a href="github.com/buaazp/fasthttprouter">fasthttprouter</a><br>
<a href="github.com/urfave/cli/v2">cli</a><br>
<a href="github.com/go-pg/pg">go-pg</a><br>
<a href="github.com/go-pg/pg/orm">orm</a><br>
<a href="github.com/joho/godotenv">godotenv</a><br><hr>
<h2>Installation (linux)</h2>
<h3>Clone the project</h3>
<code>$ git clone https://github.com/gospodinzerkalo/todo_app_golang</code>
<h2>With Docker</h2>
<code>docker-compose build</code><br>
<code>docker-compose up</code>
<h2>with Makefile</h2>
<h3>Create database</h3>
<code>$ psql -c "CREATE DATABASE todo_app" </code> <br> or <br>
<code>$ make createDB</code>

<h3>Build and run</h3>
<code>$ make build</code><br>
<code>$ make run</code>