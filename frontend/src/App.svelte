<script>
	import { onMount } from "svelte";
	import MessageDisplay from './components/MessageDisplay.svelte';
	import CustomMessage from './components/CustomMessage.svelte';
	import MessageCount from './components/MessageCount.svelte';
	
	let isDarkMode = false;
	let messages = [
		"Welcome to Svelte + Go!",
		"Testing Air for Fast Development 🚀",
		"Svelte is so reactive!",
		"Dark mode makes everything cooler 😎",
		"Let’s build awesome apps together!"
	];
	let currentMessageIndex = 0;
	let isDarkModeSupported = false; // New variable to check if dark mode is supported

	// Automatically cycle through the messages every 3 seconds
	onMount(() => {
		const interval = setInterval(() => {
			currentMessageIndex = (currentMessageIndex + 1) % messages.length;
		}, 3000);

		// Check for dark mode preference on mount
		if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
			isDarkMode = true; // If the system prefers dark mode, set it to true
			isDarkModeSupported = true; // Dark mode is supported
		}

		return () => clearInterval(interval);
	});
	
	// Toggle dark mode
	const toggleDarkMode = () => {
		isDarkMode = !isDarkMode;
	};
</script>

<main class={isDarkMode ? 'dark' : 'light'}>
	<h1>Svelte + Go</h1>
	<p>This is me trying out Svelte + Go as a potential problem solver for spinning up apps very fast</p>

	<h1>Testing Air 🚀</h1>
	<h2>Cool now that Air works</h2>

	<!-- Animated Message Carousel -->
	<div class="message-carousel">
		<p class="message">{messages[currentMessageIndex]}</p>
	</div>

	<MessageDisplay />
	<CustomMessage />
	<MessageCount />

	<!-- Dark/Light Mode Toggle, Only show if dark mode is supported -->
	{#if isDarkModeSupported}
		<button class="mode-toggle" on:click={toggleDarkMode}>
			{isDarkMode ? 'Switch to Light Mode' : 'Switch to Dark Mode'}
		</button>
	{/if}
</main>

<style>
	/* Global styles */
	main {
	  text-align: center;
	  padding: 2rem;
	  max-width: 600px;
	  margin: auto;
	  transition: all 0.3s ease;
	}
  
	h1 {
	  color: var(--primary-color);
	}
  
	/* Light Mode Styles */
	.light {
	  background-color: var(--background-light);
	  color: var(--text-light);
	}
  
	/* Dark Mode Styles */
	.dark {
	  background-color: var(--background-dark);
	  color: var(--text-dark);
	}
  
	/* Message Carousel */
	.message-carousel {
	  margin-top: 2rem;
	  font-size: 1.5rem;
	  font-weight: bold;
	  transition: opacity 1s ease;
	}
  
	/* Dark/Light Mode Toggle Button */
	.mode-toggle {
	  margin-top: 2rem;
	  padding: 0.5rem 1.5rem;
	  font-size: 1rem;
	  background-color: var(--button-background);
	  color: #fff;
	  border: none;
	  border-radius: 4px;
	  cursor: pointer;
	  transition: background-color 0.3s;
	}
  
	.mode-toggle:hover {
	  background-color: var(--button-hover-background);
	}
</style>
