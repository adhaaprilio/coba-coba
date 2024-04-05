<script>
    import { onMount } from "svelte";
    import { writable } from "svelte/store";

    let input = {
        name: "",
        username: "",
        email: "",
        password: "",
    };

    let data = writable([]);

    const handleInput = (event) => {
        const { name, value } = event.target;
        input = { ...input, [name]: value };
    };

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            const res = await fetch("http://127.0.0.1:8000/v1/register/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(input),
            });
            const responseData = await res.json();
            data.set(responseData);
        } catch (error) {
            console.log(error.message);
        }
    };
</script>

<form on:submit={handleSubmit}>
    <input
        type="text"
        name="name"
        placeholder="Name"
        on:input={handleInput}
        bind:value={input.name}
    />
    <br />
    <input
        type="text"
        name="username"
        placeholder="Username"
        on:input={handleInput}
        bind:value={input.username}
    />
    <br />
    <input
        type="email"
        name="email"
        placeholder="Email"
        on:input={handleInput}
        bind:value={input.email}
    />
    <br />
    <input
        type="password"
        name="password"
        placeholder="Password:"
        on:input={handleInput}
        bind:value={input.password} 
    />
    <br />
    <button type="submit">Register</button>
</form>

<nav>
    <a href="/auth/login">Login</a>
</nav>
