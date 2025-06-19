import { useEffect, useState } from 'react';

function App() {
  const [message, setMessage] = useState('Loading...');

  useEffect(() => {
    fetch('http://localhost:8080/hello')
      .then((res) => res.text())
      .then(setMessage)
      .catch((err) => {
        setMessage('Error: ' + err.message);
      });
  }, []);

  return (
    <div style={{ textAlign: 'center', marginTop: '4rem', fontFamily: 'sans-serif' }}>
      <h1>{message}</h1>
      <p>React Frontend 🚀</p>
    </div>
  );
}

export default App;
