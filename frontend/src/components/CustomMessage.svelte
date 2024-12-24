<script>
	let name = '';
	let location = '';
	let userInfo = null;
	let loadingUserInfo = false;
	let errorUserInfo = null;

	// Function to get user info from the API
	async function getUserInfo() {
		if (!name || !location) {
			errorUserInfo = 'Please provide both name and location';
			return;
		}

		loadingUserInfo = true;
		errorUserInfo = null;

		try {
			const res = await fetch(`/api/user-info?name=${name}&location=${location}`);
			const data = await res.json();
			userInfo = data;
		} catch (err) {
			errorUserInfo = 'Error fetching user info: ' + err.message;
			userInfo = null;
		} finally {
			loadingUserInfo = false;
		}
	}
</script>

<main>
	<h1>Custom Message Sender</h1>

	<input type="text" bind:value={name} placeholder="Enter your name" />
	<input type="text" bind:value={location} placeholder="Enter your location" />
	<button on:click={getUserInfo} disabled={loadingUserInfo}>Get User Info</button>

	{#if loadingUserInfo}
		<p>Loading user info...</p>
	{/if}

	{#if errorUserInfo}
		<p style="color: red;">{errorUserInfo}</p>
	{/if}

	{#if userInfo}
		<p><strong>Name:</strong> {userInfo.name}</p>
		<p><strong>Location:</strong> {userInfo.location}</p>
		<p><strong>Age:</strong> {userInfo.age}</p>
	{/if}
</main>

<style>
	main {
		text-align: center;
		padding: 2rem;
		max-width: 600px;
		margin: auto;
	}

	input {
		padding: 0.5rem;
		font-size: 1rem;
		width: 100%;
		max-width: 300px;
		margin-bottom: 1rem;
	}

	button {
		padding: 0.5rem 1rem;
		background-color: #007bff;
		color: white;
		border: none;
		cursor: pointer;
	}

	button:disabled {
		background-color: #cccccc;
	}

	p {
		font-size: 1.2rem;
	}
</style>
