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

function RAM() {
    const socket = useRef(null)

    const [datos, setDatos] = useState([])
    //const [procesos, setProcesos] = useState(null)
    const prueba = { "Prueba": "Hola" }
    useEffect(() => {

        socket.current = new WebSocket("ws://localhost:9100/RAM")

        socket.current.onopen = () => {
            socket.current.send(JSON.stringify(prueba))
        }
        socket.current.onmessage = (message) => {
            let aux=JSON.parse(message.data)
            let datitos={name:"ram",value:parseFloat(aux.ConsumidaP)}
            console.log(datitos)
            setDatos(datos=>[...datos,datitos])
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
            <h1>Porcentaje de uso del CPU</h1>
            {datos ? <LineChart width={500} height={300} data={datos}>
                <XAxis dataKey="name" />
                <YAxis />
                <Line dataKey="value" />
                <Tooltip/>
            </LineChart> : null}
            

        </div>
    );
}

export default RAM;