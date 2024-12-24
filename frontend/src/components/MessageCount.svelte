<script>
	let messageCount = '';
	let loading = false;
	let error = null;

	async function fetchMessageCount() {
		loading = true;
		error = null;

		try {
			const res = await fetch('/api/message-count');
			const data = await res.json();
			messageCount = data.message;
		} catch (err) {
			error = 'Error: ' + err.message;
		} finally {
			loading = false;
		}
	}
</script>

<main>
	<h1>Message Count</h1>
	<button on:click={fetchMessageCount} disabled={loading}>Fetch Message Count</button>

	{#if loading}
		<p>Loading...</p>
	{/if}

	{#if error}
		<p style="color: red;">{error}</p>
	{/if}

	{#if messageCount}
		<p>{messageCount}</p>
	{/if}
</main>

<style>
	main {
		text-align: center;
		padding: 2rem;
		max-width: 600px;
		margin: auto;
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
