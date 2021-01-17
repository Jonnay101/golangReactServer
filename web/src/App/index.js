import React, { useState } from 'react';

const App = () => {
  const [state, setState] = useState("not clicked")

  return <button onClick={() => setState("clicked")}>{state}</button>
}

export default App