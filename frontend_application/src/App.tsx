import React, { useState } from 'react';
import axios from 'axios';

const App: React.FC = () => {
  const [content, setContent] = useState<string>('');
  const [response, setResponse] = useState<string>('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const formData = new FormData();
    formData.append('content', content);

    try {
        const result = await axios.post('http://localhost:8080/submit', new URLSearchParams({ content }), {
          headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        });
        setResponse(result.data);
    } catch (error) {
        console.error(error);
        setResponse('Error submitting content');
    }
};



  return (
    <div style={{ padding: '20px' }}>
      <h1>Submit Content</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="content">Content:</label>
          <textarea
            id="content"
            value={content}
            onChange={(e) => setContent(e.target.value)}
            rows={4}
            cols={50}
            style={{ display: 'block', marginBottom: '10px' }}
          />
        </div>
        <button type="submit">Submit</button>
      </form>
      {response && <p>{response}</p>}
    </div>
  );
};

export default App;
