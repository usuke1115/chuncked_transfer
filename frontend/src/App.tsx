import { useState } from 'react'
import './App.css'

function App() {
  const [text, setText] = useState<string>("");

  const start = async () => {
    const response = await fetch("http://localhost:8080/stream");

    if (!response.body) {
      throw new Error("Response body is null");
    }

    const reader = response.body.getReader();
    const decoder = new TextDecoder();

    while (true) {
      const result: ReadableStreamReadResult<Uint8Array> = await reader.read();

      if (result.done) {
        break;
      }

      setText(prev => prev + decoder.decode(result.value));
    }
  }

  return (
    <div className='container'>
      <button className='start-button' onClick={start}>start</button>
      <pre className='text'>{text}</pre>
    </div>
  )
}

export default App
