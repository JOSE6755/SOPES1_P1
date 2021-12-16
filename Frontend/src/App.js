import Tabla from "./Components/Tabla";
import CPU from "./Components/CPU";
import RAM from "./Components/RAM";
import {
  BrowserRouter as Router,
  Routes,
  Route
} from "react-router-dom";
import { useEffect, useRef, useState } from 'react';

function App() {
  return (
    <div>
      <Router>
        <Routes>
          <Route exact path="/" element={<Tabla/>}/>
           
          <Route exact path="/CPU" element={<CPU/>} />
            
          <Route exact path="/RAM" element={<RAM/>}/>
            

        </Routes>
      </Router>
    </div>
  );
}

export default App;
