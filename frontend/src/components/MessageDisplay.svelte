<script>
    let name = '';
    let responseMessage = '';
    let isLoading = false;
  
    // Function to handle the form submission
    async function sendMessage() {
      if (!name) {
        responseMessage = 'Please enter your name.';
        return;
      }
  
      isLoading = true;
  
      try {
        const res = await fetch('/api/custom-message', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ name: name }),
        });
  
        if (!res.ok) {
          throw new Error('Failed to send message');
        }
  
        const data = await res.json();
        responseMessage = data.message;
      } catch (error) {
        responseMessage = `Error: ${error.message}`;
      } finally {
        isLoading = false;
      }
    }
  </script>
  
  <main>
    <h2>Send Your Name to Go API</h2>
  
    <form on:submit|preventDefault={sendMessage}>
      <label for="name">Enter your name:</label>
      <input type="text" id="name" bind:value={name} placeholder="Your name" />
      <button type="submit" disabled={isLoading}>
        {isLoading ? 'Sending...' : 'Send'}
      </button>
    </form>
  
    <p>{responseMessage}</p>
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
      margin: 0.5rem;
      width: 70%;
    }
  
    button {
      padding: 0.5rem 1rem;
      background-color: #4CAF50;
      color: white;
      border: none;
      cursor: pointer;
    }
  
    button:disabled {
      background-color: #ccc;
      cursor: not-allowed;
    }
  
    p {
      margin-top: 1rem;
      font-size: 1.2rem;
    }
  </style>
  