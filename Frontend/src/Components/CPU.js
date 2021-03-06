import {
    BarChart,
    Bar,
    Line,
    LineChart,
    XAxis,
    YAxis,
    Tooltip
} from 'recharts';

import { useEffect, useRef, useState } from 'react';

function CPU() {
    const socket = useRef(null)

    const [datos, setDatos] = useState([])
    const [cantidad, setCantidad] = useState({})
    const prueba = { "Prueba": "Hola" }
    useEffect(() => {

        socket.current = new WebSocket("ws://localhost:9100/CPU")

        socket.current.onopen = () => {
            socket.current.send(JSON.stringify(prueba))
        }
        socket.current.onmessage = (message) => {
            let datitos={name:"cpu",value:parseFloat(message.data)}
            console.log(datitos)
            setDatos(datos=>[...datos,datitos])
            setCantidad(datitos)
            //setDatos(currentData=>[...currentData,message.data])
            //let aux=(JSON.parse(message.data.root))
            //setDatos(aux.root)
            //setProcesos(aux.Num_proce)
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
            <h1>Porcentaje de uso del CPU: {cantidad.value}</h1>
            {datos ? <LineChart width={500} height={300} data={datos}>
                <XAxis dataKey="name" />
                <YAxis />
                <Line dataKey="value" />
                <Tooltip/>
            </LineChart> : null}
            

        </div>
    );
}

export default CPU;