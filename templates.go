package main

const createQuestionPage string = `
<h1>Create Question</h1>
<form method="post" action="/addQuestion">
	<label for="body">Question</label>
	<input type="text" id="body" name="body">
	<label for="reponse">Reponse</label>
	<input type="text" id="reponse" name="reponse">
	<button type="submit">Create</button>
</form>
`

const indexPage string = `
<h1>Login</h1>
<form method="post" action="/login">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
<form method="post" action="/register">
	<button type="submit">Register</button>
</form>
`

const registerPage string = `
<h1>Login</h1>
<form method="post" action="/addUser">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Register</button>
</form>
`
