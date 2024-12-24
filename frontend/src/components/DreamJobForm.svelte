<!-- <script>
    import { onMount } from 'svelte';
    
    // State variables for the form
    let name = '';
    let age = '';
    let dreamJob = '';
    let responseMessage = '';
    
    // Function to handle form submission
    async function submitForm() {
      const formData = {
        name,
        age: parseInt(age), // Ensure age is an integer
        dream_job: dreamJob,
      };
  
      try {
        const res = await fetch('http://localhost:8080/api/submit-dream-job', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(formData),
        });
  
        if (!res.ok) {
          throw new Error('Failed to submit form');
        }
  
        const data = await res.json();
        responseMessage = data.message || 'Form submitted successfully!';
      } catch (error) {
        responseMessage = `Error: ${error.message}`;
      }
    }
  </script>
  
  <h2>What's your dream job?</h2>
  
  <form on:submit|preventDefault={submitForm}>
    <div>
      <label for="name">Name:</label>
      <input id="name" type="text" bind:value={name} required />
    </div>
    <div>
      <label for="age">Age:</label>
      <input id="age" type="number" bind:value={age} required />
    </div>
    <div>
      <label for="dreamJob">Dream Job:</label>
      <input id="dreamJob" type="text" bind:value={dreamJob} required />
    </div>
  
    <button type="submit">Submit</button>
  </form>
  
  {#if responseMessage}
    <p>{responseMessage}</p>
  {/if}
  
  <style>
    form {
      max-width: 500px;
      margin: 0 auto;
    }
  
    label {
      display: block;
      margin-top: 1rem;
    }
  
    input {
      width: 100%;
      padding: 0.5rem;
      margin-top: 0.5rem;
    }
  
    button {
      margin-top: 1rem;
      padding: 0.7rem 1.5rem;
      background-color: #4CAF50;
      color: white;
      border: none;
      cursor: pointer;
    }
  
    button:hover {
      background-color: #45a049;
    }
  
    p {
      margin-top: 1rem;
      color: green;
    }
  </style>
   -->