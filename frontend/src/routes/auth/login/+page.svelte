<script>
    import { onMount } from "svelte";
    import { writable } from "svelte/store";

    let input = {
        username: "",
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
            const res = await fetch("http://127.0.0.1:8000/v1/user/login", {
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
        name="username"
        placeholder="Username"
        on:input={handleInput}
        bind:value={input.username}
    />
    <br />
    <input
        type="password"
        name="password"
        placeholder="Password"
        on:input={handleInput}
        bind:value={input.password}
    />
    <br />
    <button type="submit">Login</button>
</form>

<nav>
    <a href="/auth/register">Register</a>
</nav>
