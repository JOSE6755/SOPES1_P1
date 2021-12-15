
import './App.css';
import { Navbar, Container, Nav } from 'react-bootstrap';
import { Table } from 'reactstrap';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from 'react-router-dom';
import { useEffect, useRef, useState } from 'react';

function App() {
  const socket = useRef(null)
  const [datos, setDatos] = useState(null)
  const [procesos, setProcesos] = useState(null)
  const prue = { "Prueba": "Hola" }
  useEffect(() => {

    socket.current = new WebSocket("ws://localhost:9100/")

    socket.current.onopen= () => {
      socket.current.send(JSON.stringify({"Prueba":"Hola"}))
    }
    socket.current.onmessage = (message) => {
      let aux=JSON.parse(message.data)
      //let aux=(JSON.parse(message.data.root))
      setDatos(aux.root)
      setProcesos(aux.Num_proce)
    }
    



    /*axios.get("http://localhost:9101/")
      .then((response) => {
        console.log(response.data)
        if(datos!=null){
        setDatos(datos=>[...datos,response.data.root])
        }else{
          setDatos(response.data.root)
        }
        setProcesos(response.data.Num_proce)
      })*/
  }, [])



  return (
    <div>
      <div>
        <Navbar bg="dark" variant="dark">
          <Container>
            <Navbar.Brand href="#home">Navbar</Navbar.Brand>
            <Nav className="me-auto">
              <Nav.Link href="#home">Home</Nav.Link>
              <Nav.Link href="#features">Features</Nav.Link>
              <Nav.Link href="#pricing">Pricing</Nav.Link>
            </Nav>
          </Container>
        </Navbar>
      </div>
      <br />
      <br />
      <div>

        <Table dark>
          <thead>
            <tr>
              <th>
                Ejecucion
              </th>
              <th>
                Suspendidos
              </th>
              <th>
                Detenidos
              </th>
              <th>
                Zombies
              </th>
              <th>
                Total
              </th>
            </tr>
          </thead>
          <tbody>
            {procesos ? <tr>
              <td>
                {procesos.Ejecucion}
              </td>
              <td>
                {procesos.Suspendidos}
              </td>
              <td>
                {procesos.Detenidos}
              </td>
              <td>
                {procesos.Zombies}
              </td>
              <td>
                {procesos.Total}
              </td>
            </tr> : null}




          </tbody>
        </Table>
      </div>

      <div>

        <Table dark>
          <thead>
            <tr>
              <th>
                PID
              </th>
              <th>
                Proceso
              </th>
              <th>
                Usuario
              </th>
              <th>
                Estado
              </th>
              <th>
                RAM
              </th>
            </tr>
          </thead>
          <tbody>
            {datos ? datos.map((items) => {
              return (
                <tr key={items.PID}>
                  <td>
                    {items.PID}
                  </td>
                  <td>
                    {items.Proceso}
                  </td>
                  <td>
                    {items.Usuario}
                  </td>
                  <td>
                    {items.Estado}
                  </td>
                  <td>
                    {items.RAM}%
                  </td>
                </tr>
              )
            }) : null}


          </tbody>
        </Table>
      </div>


    </div>
  );
}

export default App;
